package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

const URL = "http://localhost:7000/list"

func Testmain() {

	res, err := http.Get(URL)
	if err != nil {
		log.Fatalf("error making http request: %s\n", err)
	}

	fmt.Println(res.StatusCode)

}

func BenchmarkTestmain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		res, err := http.Get(URL)
		if err != nil {
			log.Fatalf("error making http request: %s\n", err)
		}

		fmt.Println(res.StatusCode)
	}

}
