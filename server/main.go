package main

import (
	"github.com/go-terminal-server/setup"
)

func main() {

	err := setup.Start()
	if err != nil {
		panic(err.(any))
	}

}
