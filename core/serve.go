package core

import "fmt"

type ServeConfig struct {
	//bools
	ShowStartBanner bool
	SSL             bool
	//strings
	HttpAddr  string
	HttpsAddr string
	//etc

}

/*
Serves the http server

	err := core.Serve(core.ServeConfig{
		HttpAddr:        addr,
		HttpsAddr:       addr,
		SSL: 			 false,
		ShowStartBanner: false,
	})
*/
func Serve(config ServeConfig) error {
	//where all the server magic happens
	//
	var baseURL string
	if config.SSL {
		baseURL = "https://" + config.HttpsAddr
	} else {
		baseURL = "http://" + config.HttpAddr
	}
	if config.ShowStartBanner {
		fmt.Printf("REST API started at %s/api/\n", baseURL)
		fmt.Printf("Web Console started at %s\n", baseURL)
	}
	return nil
}
