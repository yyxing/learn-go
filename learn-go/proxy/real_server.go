package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RealServer struct {
	Addr string
}

func (server RealServer) Run() {
	log.Println("Starting at:", server.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		uPath := fmt.Sprintf("http://%s%s\n", server.Addr, request.URL.Path)
		log.Println(io.WriteString(writer, uPath))
	})
	mux.HandleFunc("/base/error", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(io.WriteString(writer, "Error Handler"))
	})
	httpServer := &http.Server{
		Addr:         server.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		log.Fatal(httpServer.ListenAndServe())
	}()
}
func main() {
	realServer1 := RealServer{Addr: "127.0.0.1:2003"}
	realServer2 := RealServer{Addr: "127.0.0.1:2004"}
	realServer1.Run()
	realServer2.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
