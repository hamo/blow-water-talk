package main

import (
	"fmt"
	"net/http"
	"weixin"

	"github.com/bmizerany/pat"
)

type HttpApiFunc func(w http.ResponseWriter, r *http.Request)

func makeHttpHandler(localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
		return
	}
}

var handlerMatrix = map[string]map[string]HttpApiFunc{
	"GET": {
		"/blow-water": getVerify,
	},
	"POST": {
		"/blow-water": postMessage,
	},
}

func getVerify(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var signature string
	var timestamp string
	var nonce string
	var echostr string

	signature = params.Get("signature")
	if signature == "" {
		goto badReq
	}

	timestamp = params.Get("timestamp")
	if timestamp == "" {
		goto badReq
	}

	nonce = params.Get("nonce")
	if nonce == "" {
		goto badReq
	}

	echostr = params.Get("echostr")
	if echostr == "" {
		goto badReq
	}

	if weixin.Verify(signature, timestamp, nonce) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(echostr))
		return
	}

badReq:
	w.WriteHeader(http.StatusBadRequest)
	return
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	m := make([]byte, r.ContentLength)
	r.Body.Read(m)

	fmt.Printf("Message received: %s\n", string(m))

	t, p, err := weixin.MessageDecodeReceive(string(m))

	if err != nil {
		goto badReq
	}

	return

badReq:
	w.WriteHeader(http.StatusBadRequest)
	return
}

func newRouter() http.Handler {
	h := pat.New()

	for method, routes := range handlerMatrix {
		for route, fct := range routes {
			f := makeHttpHandler(method, route, fct)
			switch method {
			case "GET":
				h.Get(route, f)
			case "POST":
				h.Post(route, f)
			}
		}
	}

	return h
}

func newAPIServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":80"),
		Handler: newRouter(),
	}
}
