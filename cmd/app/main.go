package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"music_service/internal/config"
	"music_service/internal/handler"
	"net/http"
	"time"
)

func init() {
	//err := mime.AddExtensionType(".css", "text/css")
	//if err != nil {
	//	log.Println(err)
	//}
	//err = mime.AddExtensionType(".js", "application/javascript")
	//if err != nil {
	//	log.Println(err)
	//}
	//err = mime.AddExtensionType(".png", "image/png")
	//if err != nil {
	//	log.Println(err)
	//}
}

func main() {
	cfg := config.Get()

	router := mux.NewRouter()

	staticFileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static")))
	router.PathPrefix("/static/").Handler(handlers.CompressHandler(staticFileHandler))

	router.PathPrefix("/").HandlerFunc(handler.MainHandler)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		WriteTimeout: cfg.Server.Timeout * time.Second,
		ReadTimeout:  cfg.Server.Timeout * time.Second,
	}

	log.Println("Server started on " + cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
