package main

import (
	"flag"
	"os"

	"github.com/TsvetanMilanov/todo/server/pkg/injector"
	"github.com/TsvetanMilanov/todo/server/pkg/services"
	"github.com/golang/glog"
)

func main() {
	err := flag.Set("logtostderr", "true")
	if err != nil {
		glog.Error(err)
	}

	err = flag.CommandLine.Parse([]string{})
	if err != nil {
		glog.Error(err)
	}

	graph, err := injector.CreateInjectorGraph()
	if err != nil {
		glog.Fatal(err)
		os.Exit(1)
	}

	glog.Info("Injector successfully initialized.")

	serverObject, err := injector.Resolve(*graph, "server")

	if err != nil {
		glog.Fatal(err)
		os.Exit(1)
	}

	server, ok := serverObject.(*services.Server)

	if !ok {
		glog.Fatal("Unable to start server.")
		os.Exit(1)
	}

	server.Run()
}
