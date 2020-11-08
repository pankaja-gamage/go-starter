package main

import (
	"fmt"
	"go-starter/internal/common/data"
	"go-starter/internal/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("go")

	e := godotenv.Load()
	if e != nil {
		log.Panicln(e)
	}

	data.InitDb()

	r := routers.SetupRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.RegisterOnShutdown(func() {
		log.Print("cleaning up resources before shutdown")
		data.CleanUp()
	})

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quitHandler := make(chan os.Signal)

	signal.Notify(quitHandler, syscall.SIGTERM, syscall.SIGINT)
	receivedSig := <-quitHandler

	log.Printf("Received signal %s. Server is shutting down..", receivedSig)
}
