//go:generate protoc -I rpcftp --go_out=plugins=grpc:rpcftp rpcftp/rpcftp.proto

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/faroyam/rpcftp/server"

	"github.com/faroyam/rpcftp/client"
)

var remotePath string
var localPath string
var port string
var ipAddr string
var serverMode bool
var upload bool

var servDir = "/share/"

func init() {

	flag.StringVar(&remotePath, "r", "", "[client] set remote path")
	flag.StringVar(&ipAddr, "a", "0.0.0.0", "[client] set remote ip address")
	flag.BoolVar(&upload, "u", false, "[client] upload to the server mode")

	flag.StringVar(&localPath, "l", servDir, "set local path")

	flag.StringVar(&port, "p", "50151", "set port")
	flag.BoolVar(&serverMode, "s", false, "run in server mode")

	flag.Parse()
}

func main() {
	if serverMode {
		server.Serve(port, localPath)
	} else {
		cl, err := client.NewClient(ipAddr, port)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if upload {
			if remotePath == "" || localPath == servDir {
				flag.PrintDefaults()
			}
			cl.UploadFile(localPath, remotePath)
		} else {
			if remotePath == "" || localPath == servDir {
				flag.PrintDefaults()
			}
			cl.DownloadFile(remotePath, localPath)
		}
	}
}
