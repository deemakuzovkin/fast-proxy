package main

import (
	"crypto/tls"
	"github.com/deemakuzovkin/https-proxy/pkg/configuration/https"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"net"
	"time"
)

var (
	config          *https.Configuration
	responseTimeout = time.Second
)

func redirect(ctx *fasthttp.RequestCtx) {
	ctx.Request.SetHost(config.RedirectHost)
	err := fasthttp.DoTimeout(&ctx.Request, &ctx.Response, responseTimeout)
	if err != nil {
		println(err.Error())
		return
	}
}

func main() {
	configuration, err := https.LoadConfiguration()
	if err != nil {
		println(err.Error())
		return
	}
	config = configuration
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(config.HostsPolicy...),
		Cache:      autocert.DirCache("./certs"),
		Email:      config.AcmeEmail,
	}
	tlcConfig := &tls.Config{
		GetCertificate: certManager.GetCertificate,
		NextProtos: []string{
			"http/1.1", acme.ALPNProto,
		},
	}
	ln, err := net.Listen("tcp4", "0.0.0.0:443")
	if err != nil {
		panic(err)
	}
	lnTls := tls.NewListener(ln, tlcConfig)
	if err := fasthttp.Serve(lnTls, redirect); err != nil {
		panic(err)
	}
}
