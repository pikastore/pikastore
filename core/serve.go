package core

import (
	"fmt"
	"net/http"

	"github.com/pikastore/pikastore/router"
)

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
		HttpAddr:        "0.0.0.0:80",
		HttpsAddr:       "0.0.0.0:443",
		SSL: 			 false,
		ShowStartBanner: true,
	})
*/
func Serve(config ServeConfig) error {
	//Start server
	server := router.New()
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

	if config.SSL {
		return http.ListenAndServeTLS(config.HttpsAddr, "", "", server)
	} else {
		return http.ListenAndServe(config.HttpAddr, server)
	}
}
