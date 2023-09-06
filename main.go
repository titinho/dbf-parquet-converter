package main

import (
	"fmt"
	"os"
	"strings"
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
	if inputFilePathParts[1] != "dbf" {
		err := fmt.Errorf(
			"the input file must have dbf extension - %w",
			ErrInvalidFileExtension,
		)
		return "", "", err
	}

	// Check the output file extension
	outputFilePathParts := strings.Split(outputFilePath, ".")
	if outputFilePathParts[1] != "parquet" {
		err := fmt.Errorf(
			"the output file must have parquet extension - %w",
			ErrInvalidFileExtension,
		)
		return "", "", err
	}

	return inputFilePath, outputFilePath, nil
}
