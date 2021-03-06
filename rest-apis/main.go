package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"main.go/config"
	"main.go/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
	gohandlers "github.com/gorilla/handlers"
)

func loadEnv() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var environment config.Environment

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&environment)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("port for this application is %d", environment.Server.Port)
}

func main() {

	loadEnv()


	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddleWareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddleWareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	//CORS header
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))


	s := http.Server{
		Addr: ":8080",
		Handler: ch(sm),
		IdleTimeout: 120*time.Second,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate ,graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown( tc)

}
