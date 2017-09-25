package main

import (
	"crypto/tls"
	"time"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	target := os.Args[1]
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatalln(err)
	}

	var dnsStartAt int64
	var dnsDoneAt int64

	var connStartAt int64
	var connDoneAt int64

	var TLSHandshakeStartAt int64
	var TLSHandshakeEndAt int64

	var gotFirstResponseByteAt int64

	trace := &httptrace.ClientTrace{
		DNSStart: func(dnsinfo httptrace.DNSStartInfo) {
			dnsStartAt = time.Now().UnixNano()
			fmt.Printf("[DNS START] %+v\n", dnsinfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDoneAt = time.Now().UnixNano()
			cost := (dnsDoneAt - dnsStartAt) / int64(time.Millisecond)
			fmt.Printf("[DNS INFO] %+v, COST: %vms\n", dnsInfo, cost)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			connStartAt = time.Now().UnixNano()
			fmt.Printf("[Got Conn] %+v\n", connInfo)
		},
		ConnectStart: func(network, addr string) {
			connStartAt = time.Now().UnixNano()
			fmt.Printf("[ConnectStart] Network: %s, Addr: %s\n", network, addr)
		},
		ConnectDone: func(network, add string, err error){
			if err != nil {
				log.Fatalln(err)
			}
			connDoneAt = time.Now().UnixNano()
			cost := (connDoneAt - connStartAt) / int64(time.Millisecond)
			fmt.Printf("[ConnectDone] Network: %s, Addr: %s, COST: %vms\n", network, add, cost)
		},
		TLSHandshakeStart: func() {
			TLSHandshakeStartAt = time.Now().UnixNano()
			fmt.Printf("[TLSHandshakeStart]\n")						
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, err error) {
			if err != nil {
				log.Fatalln(err)
			}
			TLSHandshakeEndAt = time.Now().UnixNano()
			cost := (TLSHandshakeEndAt - TLSHandshakeStartAt) / int64(time.Millisecond)
			fmt.Printf("[TLSHandshakeDone] COST: %vms\n", cost)			
		},
		GotFirstResponseByte: func() {
			gotFirstResponseByteAt = time.Now().UnixNano()
			cost := (gotFirstResponseByteAt - dnsStartAt) / int64(time.Millisecond)
			fmt.Printf("[GotFirstResponseByte]: Total %vms\n", cost)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
}