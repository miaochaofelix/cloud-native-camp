package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func remoteIp(request *http.Request) string {
	remoteAddr := request.RemoteAddr
	if ip := request.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = request.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Println("this is home page")
	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for key, value := range request.Header {
		for _, vv := range value {
			//fmt.Printf("%s: %s \n", key, vv)
			response.Header().Set(key, vv)
		}
	}
	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "latest")
	response.Header().Set("VERSION", os.Getenv("VERSION"))
	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("client IP: %s, response code: %d\n", remoteIp(request), 200)
	response.Write([]byte("<h1>this is home page</h1>"))
}

func healthz(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("<h1>this is healthz page</h1>"))
}

func main() {
	fmt.Println("starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// 4. 当访问 localhost/healthz 时，应返回 200
	mux.HandleFunc("/healthz", healthz)
	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		fmt.Printf("start http server failed, error = %s\n", error.Error())
	}
}
