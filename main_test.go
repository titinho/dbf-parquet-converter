package main

import (
	"fmt"
	"os"
	"testing"
)

func TestProcessArgsFunc(t *testing.T) {
	type want struct {
		inputFile  string
		outputFile string
		err        error
	}

	scenarios := []struct {
		description string
		args        []string
		want        want
	}{
		{
			description: "Happy case - long flags",
			args:        []string{"--input=Product.dbf", "--output=Product.parquet"},
			want: want{
				inputFile:  "Product.dbf",
				outputFile: "Product.parquet",
				err:        nil,
			},
		},
		{
			description: "Happy case - short flags",
			args:        []string{"-i=Product.dbf", "-o=Product.parquet"},
			want: want{
				inputFile:  "Product.dbf",
				outputFile: "Product.parquet",
				err:        nil,
			},
		},
		{
			description: "Error - unknown flag",
			args:        []string{"--list"},
			want: want{
				inputFile:  "",
				outputFile: "",
				err:        fmt.Errorf("%s - %w", "--list", ErrUnknownFlag),
			},
		},
		{
			description: "Error - input file not given",
			args:        []string{"--output=Product.parquet"},
			want: want{
				inputFile:  "",
				outputFile: "",
				err: fmt.Errorf(
					"the --input/-i flag is missing - %w",
					ErrMissingRequiredFlag,
				),
			},
		},
		{
			description: "Error - output file not given",
			args:        []string{"--input=Product.dbf"},
			want: want{
				inputFile:  "",
				outputFile: "",
				err: fmt.Errorf(
					"the --output/-o flag is missing - %w",
					ErrMissingRequiredFlag,
				),
			},
		},
		{
			description: "Error - input file has invalid extension",
			args:        []string{"--input=Product.xlsx", "--output=Product.parquet"},
			want: want{
				inputFile:  "",
				outputFile: "",
				err: fmt.Errorf(
					"the input file must have dbf extension - %w",
					ErrInvalidFileExtension,
				),
			},
		},
		{
			description: "Error - output file has invalid extension",
			args:        []string{"--input=Product.dbf", "--output=Product.xlsx"},
			want: want{
				inputFile:  "",
				outputFile: "",
				err: fmt.Errorf(
					"the output file must have parquet extension - %w",
					ErrInvalidFileExtension,
				),
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			os.Args = []string{os.Args[0]}
			os.Args = append(os.Args, scenario.args...)

			gotInputFile, gotOutputFile, gotErr := ProcessFlags()

			if gotInputFile != scenario.want.inputFile {
				t.Errorf("got %s want %s", gotInputFile, scenario.want.inputFile)
			}

			if gotOutputFile != scenario.want.outputFile {
				t.Errorf("got %s want %s", gotOutputFile, scenario.want.outputFile)
			}

			if scenario.want.err == nil {
				if gotErr != nil {
					t.Errorf("got %s want nil", gotErr.Error())
				}
			} else {
				if gotErr.Error() != scenario.want.err.Error() {
					t.Errorf("got %s want %s", gotErr.Error(), scenario.want.err.Error())
				}
			}
		})
	}
}
