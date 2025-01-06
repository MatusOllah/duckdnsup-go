# duckdnsup-go

**duckdnsup-go** is a Go tool for updating the IP address of your Duck DNS domain(s).

## Building & Installing from source

1. Install Go
2. Run `go install github.com/MatusOllah/duckdnsup-go@latest`

## Usage

`duckdnsup-go -d domain -t duck-dns-token`

### Flags

* `-d`, `--domain` - The domain(s) to be updated
* `-i`, `--ip` - The IP address, auto-detect when empty
* `-t`, `--token` - The token
* `-v`, `--verbose` - Print verbose information
