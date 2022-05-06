package main

import (
	"andySantisteban/init/server"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	serverDoneChanell := make(chan os.Signal, 1)
	signal.Notify(serverDoneChanell, os.Interrupt)
	server := server.Server("localhost:8080")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	server.ListenAndServe()
	log.Println("Starting server...")

	<-serverDoneChanell

	server.Shutdown(ctx)
	log.Println("Server stop...")

}
