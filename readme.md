PearTok("/pɛrtoʊk/") is a peer-to-peer, token transfer app made with Go, Gin and JavaScript. Instead of relying on centralized servers, it uses direct TCP connections to connect with peers.


Disclaimer: This software is provided "as is" without any warranty of any kind. These tokens do not have any monetary value.


## Usage

1. Build the project by `go build -o peartok-go .`


2. Copy `wallet-receiver.tok` and `wallet-sender.tok` into same directory as the executable.  

3. Start the http server by `./peartok-go`. Default is 8080.

From now you can visit http://localhost:8080 on receiver's machine and enter port to start a server and then send tokens. Or use the script below.

-  Review `peartok-curl.sh`. You can edit port, amount, host if you want to.

 -  Run the script, and optionally see the updated balance by printing `*.tok` files to stdout.
