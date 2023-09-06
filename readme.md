# DBF-Parquet Converter

This CLI tool helps you to convert DBF database files to parquet and back.

## Usage

There are some flags to use:

- --input/-i (required): the path to the input file
- --output/-o (required): the path to the output file

Example

```bash
go run main.go --input=Product.dbf --output=Product.parquet
```
