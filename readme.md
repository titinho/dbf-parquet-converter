# DBF-Parquet Converter

[![codecov](https://codecov.io/gh/titinho/dbf-parquet-converter/graph/badge.svg?token=5ADKWR0EDF)](https://codecov.io/gh/titinho/dbf-parquet-converter)

This CLI tool helps you to convert DBF database files to parquet and back.

## Installation

## Usage

There are some flags to use:

- --input/-i (required): the path to the input file
- --output/-o (required): the path to the output file

Example

```bash
go build
./dbf-parquet-converter --input=Product.dbf --output=Product.csv
```
