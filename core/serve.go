package core

import (
	"fmt"
	"net/http"

	"github.com/pikastore/pikastore/api"
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
Serves the http sersver

	err := core.Serve(core.ServeConfig{
		HttpAddr:        "0.0.0.0:80",
		HttpsAddr:       "0.0.0.0:443",
		SSL: 			 false,
		ShowStartBanner: true,
	})
*/
func Serve(config ServeConfig) error {
	//Start server
	router := router.New()

	//routes (I should move this to another file)
	rg := router.Group("/api")
	api.RegisterAuth(rg)

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
		return http.ListenAndServeTLS(config.HttpsAddr, "", "", router)
	} else {
		return http.ListenAndServe(config.HttpAddr, router)
	}
}
