package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pikastore/pikastore/core"
	"github.com/pikastore/pikastore/router"
)

func Serve(addr string) {
	r := router.New(&router.Config{
		PORT: addr,
		// SSL:  useSSL,
		// Certs: router.Certs{
		// 	SSLFile: sslCert,
		// 	SSLKey:  sslKey,
		// },
	})

	// Test route
	r.Get("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "haiiii :3",
		})
	})
	core.Log("Server started successfully on 0.0.0.0%s\n" + addr)
	if err := r.Run(addr); err != nil {
		core.Error(err.Error(), true)
		return
	}
}
