package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/muizidn/movieboomgo/thrift/gen-go/movie"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client *movie.MovieServiceClient) (err error) {
	fmt.Println("Client...")
	res, err := client.Get(defaultCtx, "8e43b175c34985bfc0da05ed11b3eca9", "en")
	fmt.Println(res, err)
	if err != nil {
		return err
	}
	fmt.Println("Get result...")
	fmt.Println(res.GetOriginalTitle())
	fmt.Println("End")
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(movie.NewMovieServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
