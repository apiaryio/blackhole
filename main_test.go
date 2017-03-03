package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestConn(t *testing.T) {
	os.Setenv("PORT", "8888")
	go SetupServer()
	res, err := http.Get("http://0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
	}
	message, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(message))
}
