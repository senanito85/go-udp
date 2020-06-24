go-udp
======

Simpe UDP server and client written in go.

To get the repo:
$ go get github.com/senanito85/go-udp/

$ cd ~/go/src/github.com/senanito85/go-udp/

$ make

binaries will be created under target directory

To Start the server: 
$ client-darwin-amd64
or
$ client-linux-amd64


To Start the Client:
echo "test message" ./client-darwin-amd64
or
echo "test message" ./client-linux-amd64

alternatively just ./client-linux-amd64 or ./client-darwin-amd64
then type the message to prompt

Modifications
=====
Server and client are configured to send and listen to port 2323 on 127.0.0.1. So it will work on same machine.
When two servers used IP address on both client and server should be updated accordingly.
