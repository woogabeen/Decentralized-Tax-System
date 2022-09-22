package main

import (
	"github.com/WoodoCoin/cli"
	"github.com/WoodoCoin/db"
	"github.com/WoodoCoin/explorer"
	"github.com/WoodoCoin/tax"
)

func main() {
	defer db.Close()
	go tax.CalcTax()
	cli.Start()
	explorer.Start(7070)
}
