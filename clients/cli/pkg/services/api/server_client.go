package api

import (
	"bytes"

	"fmt"

	"net/http"

	"io/ioutil"

	"encoding/json"

	"github.com/TsvetanMilanov/todo/clients/cli/pkg/constants"
	"github.com/TsvetanMilanov/todo/clients/cli/pkg/types"
)

// ServerClient server api request methods.
type ServerClient struct {
	ServerConfigManager types.IServerConfigManager `inject:"serverConfigManager"`
	Helpers             types.IHelpers             `inject:"helpers"`
}

// Post makes POST request to the server api.
func (server *ServerClient) Post(urlPath string, body interface{}, headers map[string]string, result interface{}) error {
	return server.sendRequest(constants.HTTPPost, urlPath, body, headers, result)
}

func (server *ServerClient) sendRequest(method string,
	urlPath string,
	body interface{},
	headers map[string]string,
	result interface{}) error {
	serverConfig, err := server.ServerConfigManager.GetServerConfig(server.Helpers.GetEnv())
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s://%s/%s/%s",
		serverConfig.Proto,
		serverConfig.Host,
		"api",
		urlPath,
	)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bodyContent := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest(method, fullURL, bodyContent)
	if err != nil {
		return err
	}

	req.Header.Set(constants.ContentTypeHeader, constants.ApplicationJSONContentType)

	// TODO: Add authorization header.
	for h, v := range headers {
		req.Header.Set(h, v)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, result)

	if err != nil {
		return err
	}

	return nil
}
