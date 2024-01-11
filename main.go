package main

import (
	"github.com/yigithankarabulut/simplemedia/src/apiserver"
	"log"
	"os"
)

func main() {
	if err := apiserver.New(
		apiserver.WithLogLevel(os.Getenv("LOG_LEVEL")),
	); err != nil {
		log.Fatal(err)
	}
}
