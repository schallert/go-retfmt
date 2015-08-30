# go-retfmt

[![Build Status](https://travis-ci.org/schallert/go-retfmt.svg?branch=master)](https://travis-ci.org/schallert/go-retfmt)

A small tool to check go source code for proper formatting according to `gofmt`, but rather than reformat the code it exits with exit code `0` if the code is properly formatted and `2` if it is not (and `1` if there was some other sort of error). This is useful in the case such as a CI build script if you want to easily check if code complies with `gofmt` standards.

Tested with `go1.3` through `go1.5`, may work with other versions but not necessarily guaranteed.

## Usage

```
Usage of ./go-retfmt:
  -dir string
        Directory to search (default ".")
  -ignore string
        Comma-separated directories to ignore (useful for vendored deps) (default "")
```
