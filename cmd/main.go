package main

import (
	"flag"
	"log"
	"os"

	"github.com/StefanUA/AccountManager/internal/service"
	accountmanager "github.com/StefanUA/AccountManager/pkg/accountManager"
)

func main() {
	inputFilePtr := flag.String("input", "", "Input file to process (Required)")
	outputFilePtr := flag.String("output", "", "Output file to write results (Required)")
	flag.Parse()

	if *inputFilePtr == "" {
		log.Fatal("empty value passed for required field 'input'")
		flag.PrintDefaults()
		os.Exit(1)
	}

	accountManager := accountmanager.NewCommand(service.TransactionService{})
	err := accountManager.Execute(*inputFilePtr, *outputFilePtr)
	if err != nil {
		log.Fatalf("Error processing: %v", err)
		os.Exit(1)
	}
}
