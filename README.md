# go-retfmt

A small tool to check go source code for proper formatting according to `gofmt`, but rather than reformat the code it exits with exit code `0` if the code is properly formatted and `1` if it is not. This is useful in the case such as a CI build script if you want to easily check if code complies with `gofmt` standards.