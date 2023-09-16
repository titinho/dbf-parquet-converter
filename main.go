package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LindsayBradford/go-dbf/godbf"
)

func ProcessFlags() (string, string, error) {
	var inputFilePath, outputFilePath string

	// Process the flags
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		switch parts[0] {
		case "-i", "--input":
			inputFilePath = parts[1]
		case "-o", "--output":
			outputFilePath = parts[1]
		default:
			err := fmt.Errorf("%s - %w", arg, ErrUnknownFlag)
			return "", "", err
		}
	}

	// Check if the input flag is given
	if inputFilePath == "" {
		err := fmt.Errorf("the --input/-i flag is missing - %w", ErrMissingRequiredFlag)
		return "", "", err
	}

	// Check if the output flag is given
	if outputFilePath == "" {
		err := fmt.Errorf("the --output/-o flag is missing - %w", ErrMissingRequiredFlag)
		return "", "", err
	}

	// Check the input file extension
	inputFilePathParts := strings.Split(inputFilePath, ".")
	if (inputFilePathParts[1] != "dbf") && (inputFilePathParts[1] != "DBF") {
		err := fmt.Errorf(
			"the input file must have dbf or DBF extension - %w",
			ErrInvalidFileExtension,
		)
		return "", "", err
	}

	// Check the output file extension
	outputFilePathParts := strings.Split(outputFilePath, ".")
	if outputFilePathParts[1] != "csv" {
		err := fmt.Errorf(
			"the output file must have csv extension - %w",
			ErrInvalidFileExtension,
		)
		return "", "", err
	}

	return inputFilePath, outputFilePath, nil
}
func main() {
	inputFilePath, outputFilePath, err := ProcessFlags()
	if err != nil {
		panic(err)
	}

	dbf, err := godbf.NewFromFile(inputFilePath, "852")
	if err != nil {
		log.Fatal(err.Error())
	}


	// Write file
	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	// Headers
	fields := dbf.FieldNames()
	for i, field := range fields {
		if i != len(fields)-1 {
			file.WriteString(fmt.Sprintf("%s,", field))
		} else {
			file.WriteString(fmt.Sprintf("%s\n", field))
		}
	}

	// Table
	for row := 0; row < dbf.NumberOfRecords(); row++ {
		if dbf.HasRecord(row) {
			for i, field := range fields {
				valueStr, err := dbf.FieldValueByName(row, field)
				if err != nil {
					err := fmt.Errorf("in the %d. row there is no name %s - %w", i, field, err)
					log.Fatalln(err.Error())
				}

				if i != len(fields)-1 {
					file.WriteString(fmt.Sprintf("%s,", valueStr))
				} else {
					file.WriteString(fmt.Sprintf("%s\n", valueStr))
				}
			}
		} else {
			err := fmt.Errorf("%d. row - record does not exist", row)
			log.Fatalln(err.Error())
		}
	}
}
