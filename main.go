package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zeucxb/financial_control_api/repository"
	"github.com/zeucxb/financial_control_api/routes"
)

var port *int

func init() {
	log.Println("Initializing...")

	parseFlags()

	repository.Init()
}

func parseFlags() {
	port = flag.Int("port", 8080, "server port")

	flag.Parse()
}

func main() {
	router := routes.Init()

	log.Println("Server running at port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
