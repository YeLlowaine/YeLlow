package main

import (
	"fmt"
	"net/http"

	"github.com/Hallelujah1025/Stroke-Survivors/pkg/setting"
	"github.com/Hallelujah1025/Stroke-Survivors/routers"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
