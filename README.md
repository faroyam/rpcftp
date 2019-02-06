# rpcftp

rpcftp is another ftp implementation written in go that uses grpc.

## Installation

Installation requires the Go toolchain, the [`protoc` compiler](https://github.com/google/protobuf), the [`protoc-gen-go` plugin](https://github.com/golang/protobuf).

```bash
go get github.com/faroyam/rpcftp
./build.sh
```

## Usage

```bash
$ ./rpcftp-exe -h
Usage of ./rpcftp-exe:
  -a string
        [client] set remote ip address (default "0.0.0.0")
  -l string
        set local path (default "/share/")
  -p string
        set port (default "50151")
  -r string
        [client] set remote path
  -s    run in server mode
  -u    [client] upload to the server mode
```

## Example

Client
```bash
$ ./rpcftp-exe -u -l rpcftp-exe -r test/test/test.bin
######################################################################
###########################################
Elapsed time: 172.277146ms
Server answer: success
Code: Ok
```

Server
```bash
$ ./rpcftp-exe -s
2019/01/27 16:10:04 creating listener on 50151
2019/01/27 16:10:04 creating grpc server
2019/01/27 16:10:04 serving /share/
2019/01/27 16:10:45 upload rpc called: /share/test/test/test.bin

```