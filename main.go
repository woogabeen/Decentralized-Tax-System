package main

import (
	"github.com/WoodoCoin/cli"
	"github.com/WoodoCoin/db"
	"github.com/WoodoCoin/explorer"
)

func main() {
	defer db.Close()
	cli.Start()
	explorer.Start(7070)
}
