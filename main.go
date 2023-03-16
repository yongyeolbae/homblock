package main

import (
	"github.com/baeyongyeol/explorer"
	"github.com/baeyongyeol/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
