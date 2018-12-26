package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/maria-robobug/dogfacts/dogapi"
)

func main() {
	c := &http.Client{
		Timeout: time.Second * 30,
	}

	dc, err := dogapi.NewDogClient(c)
	if err != nil {
		log.Fatalf("could not create dog client:\n%s", err)
	}

	dogInfo, err := dc.GetRandomDogInfo()
	if err != nil {
		log.Fatalf("could not connect to dog api:\n%s", err)
	}

	fmt.Printf("%+v\n", dogInfo)
}
