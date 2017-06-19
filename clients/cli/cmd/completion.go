package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var shell string

var completionCmd = &cobra.Command{
	Use:     "completion",
	Example: `  source <(todo completion --shell zsh)`,
	Short:   "Generates shell completion",
	PreRun:  helpers.CheckFlags,
	Run:     runCompletionCmd,
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().StringVar(&shell, "shell", os.Getenv("SHELL"), "zsh/bash")
	helpers.MarkFlagRequired(completionCmd, "shell")
}

func runCompletionCmd(cmd *cobra.Command, args []string) {
	if shell == "bash" {
		err := cmd.Root().GenBashCompletion(os.Stdout)
		if err != nil {
			glog.Error(err)
			os.Exit(1)
		}
		return
	}
	err := zshCompletion(os.Stdout, cmd.Root())
	if err != nil {
		glog.Error(err)
	}
}

// https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/completion.go
// sed -e 's/kubectl/todo/g'
func zshCompletion(out io.Writer, todo *cobra.Command) error {
	zshInitialization := `
__todo_bash_source() {
	alias shopt=':'
	alias _expand=_bash_expand
	alias _complete=_bash_comp
	emulate -L sh
	setopt kshglob noshglob braceexpand
	source "$@"
}
__todo_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift
		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__todo_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}
__todo_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?
	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}
__todo_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}
__todo_declare() {
	if [ "$1" == "-F" ]; then
		whence -w "$@"
	else
		builtin declare "$@"
	fi
}
__todo_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}
__todo_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}
__todo_filedir() {
	local RET OLD_IFS w qw
	__debug "_filedir $@ cur=$cur"
	if [[ "$1" = \~* ]]; then
		# somehow does not work. Maybe, zsh does not call this at all
		eval echo "$1"
		return 0
	fi
	OLD_IFS="$IFS"
	IFS=$'\n'
	if [ "$1" = "-d" ]; then
		shift
		RET=( $(compgen -d) )
	else
		RET=( $(compgen -f) )
	fi
	IFS="$OLD_IFS"
	IFS="," __debug "RET=${RET[@]} len=${#RET[@]}"
	for w in ${RET[@]}; do
		if [[ ! "${w}" = "${cur}"* ]]; then
			continue
		fi
		if eval "[[ \"\${w}\" = *.$1 || -d \"\${w}\" ]]"; then
			qw="$(__todo_quote "${w}")"
			if [ -d "${w}" ]; then
				COMPREPLY+=("${qw}/")
			else
				COMPREPLY+=("${qw}")
			fi
		fi
	done
}
__todo_quote() {
    if [[ $1 == \'* || $1 == \"* ]]; then
        # Leave out first character
        printf %q "${1:1}"
    else
    	printf %q "$1"
    fi
}
autoload -U +X bashcompinit && bashcompinit
# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q GNU; then
	LWORD='\<'
	RWORD='\>'
fi
__todo_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__todo_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__todo_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__todo_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__todo_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__todo_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/__todo_declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__todo_type/g" \
	<<'BASH_COMPLETION_EOF'
`
	_, err := out.Write([]byte(zshInitialization))
	if err != nil {
		glog.Error(err)
		return err
	}
	buf := new(bytes.Buffer)
	err = todo.GenBashCompletion(buf)
	if err != nil {
		glog.Error(err)
		return err
	}
	_, err = out.Write(buf.Bytes())
	if err != nil {
		glog.Error(err)
		return err
	}

	zshTail := `
BASH_COMPLETION_EOF
}
__todo_bash_source <(__todo_convert_bash_to_zsh)
`
	_, err = out.Write([]byte(zshTail))
	if err != nil {
		glog.Error(err)
	}
	return err
}
