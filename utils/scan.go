package utils

import (
	"bufio"
	"log"
	"os"
)

func Scan(fileName string, directory string) []string{
	// open the file using Open() function from os library
	file, err := os.Open("../" + directory + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	var result []string
	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
    // check for the error that occurred during the scanning
    
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    
	// Close the file at the end of the program
	file.Close()

	return result
}
