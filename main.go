package main

import (
	"github.com/WoodoCoin/cli"
	"github.com/WoodoCoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
