PearTok("/pɛrtoʊk/") is a peer-to-peer, token transfer app made with Go, Gin and TypeScript. Instead of relying on centralized servers, it uses direct TCP connections to connect with peers.

## Note: This version of PearTok is depreciated, and will not receive any updates. I am working on a different version that uses Typescript and Node instead of Go.
Disclaimer: This software is provided "as is" without any warranty of any kind. These tokens do not have any monetary value.


## Usage

This project is deprecated and the TypeScript version is still in development. The frontend code is not yet ready for release. However, a shell script that can be used to send or receive tokens is provided.

1. Build the project by `go build -o peartok-go .`


2. Copy `wallet-receiver.tok` and `wallet-sender.tok` into same directory as the executable.  

3. Start the http server by `./peartok-go`. Default is 8080.

4. Review `peartok-curl.sh`. You can edit port, amount, host if you want to.

5. Run the script, and optionally see the updated balance by printing `*.tok` files to stdout.
