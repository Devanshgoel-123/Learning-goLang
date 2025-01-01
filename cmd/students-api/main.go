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
	"github.com/Devanshgoel-123/students-api/internal/http/handlers/student"
	"github.com/Devanshgoel-123/students-api/internal/storage/sqlite"
)

func main(){
	fmt.Println("Welcome to the students API");

	//load Config -> this can be done by running mustLoad function inside the internal folder
	cfg := config.MustLoad()
	//database setup

	storage,err:=sqlite.New(cfg)
	if(err!=nil){
		log.Fatal("Error while connecting to the Database",err)
	}

	slog.Info("storage initialized", slog.String("env",cfg.Env), slog.String("version","1.0.0"))

	//setup router

	router:=http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))
	router.HandleFunc("GET /api/students/{id}",student.GetStudentById(storage))
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
	err2:=server.Shutdown(ctx);
	if err2!=nil{
		slog.Error("failed to shutdown server", slog.String("error",err2.Error()))
	}

	slog.Info("server shutdown succesfully")
}