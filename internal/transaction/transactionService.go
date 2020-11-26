package transaction

import (
	"bufio"
	"log"
	"os"
)

//ProcessTransactionFile receives an input file location
//and executes transactions written in the file
func ProcessTransactionFile(inputFile *string) error {
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		log.Println(line)
	}

	return nil
}
