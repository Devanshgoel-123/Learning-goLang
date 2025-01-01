package main

//entry point of the application

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Devanshgoel-123/students-api/internal/config"
)

func main(){
	fmt.Println("Welcome to the students API");

	//load Config -> this can be done by running mustLoad function inside the internal folder
	cfg := config.MustLoad()
	//database setup
	//setup router

	router:=http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	//setup server

	server:=http.Server{
		Addr: cfg.Address,
		Handler: router,
	}
	fmt.Printf("Started the server at Port : %s",cfg.Address)
	done:=make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//to close server gracefully, creating the closing cuntion inside go routine
	go func(){
		err:=server.ListenAndServe();
	if(err!=nil){
		log.Fatal("Error running the server")
	}
	}()

	<-done //code bl0cked until done channel reads some value
	//server stopping code

	slog.Info("Shutting down the server")
	ctx, cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err:=server.Shutdown(ctx);
	if err!=nil{
		slog.Error("failed to shutdown server", slog.String("error",err.Error()))
	}

	slog.Info("server shutdown succesfully")
}