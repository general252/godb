package main

import (
	"github.com/general252/godb/model"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	model.Build()
}
