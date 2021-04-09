package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	handlers "github.com/gospodinbodurov/ports-rest-client/handlers"
	readerjob "github.com/gospodinbodurov/ports-rest-client/readerjob"

	client "github.com/gospodinbodurov/ports-port-domain-service/client"
	clients "github.com/gospodinbodurov/ports-rest-client/clients"
)

func startReaderJob(filename string) {
	job := readerjob.ReaderJob{
		Filename: filename,
	}

	job.Run()
}

func main() {
	flag.String("httpAddress", ":10000", "Provide an url for starting the http service")
	flag.String("serviceAddress", "localhost:6666", "Provide the grpc service url")
	flag.String("databaseFilename", "./database.json", "Provide the absolute path to the json database")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	grpcAddress := viper.GetString("serviceAddress")
	httpAddress := viper.GetString("httpAddress")
	databaseFilename := viper.GetString("databaseFilename")

	clients.ServiceClient = client.DomainPortClient{}

	clients.ServiceClient.Init(grpcAddress)
	defer clients.ServiceClient.Close()

	go func() {
		startReaderJob(databaseFilename)
	}()

	http.HandleFunc("/get-port", handlers.GetPort)

	log.Fatal(http.ListenAndServe(httpAddress, nil))
}
