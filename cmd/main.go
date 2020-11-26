package main

import "os"

func main() {
	string inputFile = os.Args[0]
	accountmanager.NewCommand()
	ProcessTransactionFile(inputFile)
	
}
