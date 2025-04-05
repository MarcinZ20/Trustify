package client

import (
	"time"

	fastshot "github.com/opus-domini/fast-shot"
)

type ClientConfig struct {
	Client fastshot.ClientHttpMethods
}

var clientConfig *ClientConfig

func GetClientConfig() *ClientConfig {
	if clientConfig != nil {
		return clientConfig
	}

	clientConfig = &ClientConfig{
		Client: fastshot.NewClient("http://localhost:5000").
			Header().Add("Content-Type", "application/json").
			Config().SetTimeout(15 * time.Second).Build(),
	}

	return clientConfig
}
