# product-microservice
Product microservice using Golang, GRPC and Postgresql

Install protobuf
```
$ brew install protobuf
```
The protocol buffer compiler requires a plugin to generate Go code.
```
$ go get github.com/golang/protobuf/protoc-gen-go
```
To compile the protocol buffer definition, run protoc with the --go_out parameter set to the directory you want to output the Go code to.
```
$ protoc --go_out=. *.proto
```
