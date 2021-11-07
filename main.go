package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"fmt"
)

type handle struct {
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("https://" + this.host + ":" + this.port+"/")
	r.Header.Add("Host", this.host)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	w.Header().Set("ccc", this.host)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "s3.ladydaily.com", port: "443"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}