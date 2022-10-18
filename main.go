package main

import (
	"github.com/WoodoCoin/cli"
	"github.com/WoodoCoin/db"
	"github.com/WoodoCoin/tax"
)

func main() {
	defer db.Close()
	go cli.Start()
	tax.CalcTax()
}
