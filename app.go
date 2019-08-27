package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {

	apiReverseProxy := buildReverseProxyFromVarOrPanic("apiAddr")
	uiReverseProxy := buildReverseProxyFromVarOrPanic("uiAddr")

	listenAddr := os.Getenv("listenAddr")
	if listenAddr == "" {
		logrus.Fatalln("'", listenAddr, "'", "not a valid listen address")
	}
	http.Handle("/api", apiReverseProxy)
	http.Handle("/", uiReverseProxy)
	logrus.Fatalln(http.ListenAndServe(listenAddr, nil))

}

func buildReverseProxyFromVarOrPanic(v string) *httputil.ReverseProxy {
	vs := os.Getenv(v)
	vURL, err := url.Parse(vs)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln(os.Getenv(vs))
		panic(err)
	}
	rp := httputil.NewSingleHostReverseProxy(vURL)
	return rp
}
