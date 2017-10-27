package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
)

type Proxy struct {
	authorizationHeader string
	upstreamURI *url.URL
}

func (p *Proxy) New() http.Handler {
	return p.upstreamHost(httputil.NewSingleHostReverseProxy(p.upstreamURI))
}

func (p *Proxy) upstreamHost(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Authorization", p.authorizationHeader)
		r.Host = r.URL.Host
		handler.ServeHTTP(w, r)
	})
}

func main() {
	accessToken := flag.String("token", "", "API Bearer token")
	upstreamHost := flag.String("host", "", "Upstream host")
	port := flag.String("port", "8080", "Port")
	flag.Parse()

	if *accessToken == "" || *upstreamHost == "" {
		log.Fatal("host or token as argument not provided")
	}

	upstreamURI, err := url.Parse(*upstreamHost)
	if err != nil {
		log.Fatal("URL failed to parse")
	}

	authorizationHeader := fmt.Sprintf("Bearer %s", *accessToken)
	reverseProxy := Proxy{authorizationHeader, upstreamURI}

	fmt.Printf("Server listens on :%s\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), reverseProxy.New())
}
