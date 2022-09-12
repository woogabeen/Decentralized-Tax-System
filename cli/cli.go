package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/WoodoCoin/explorer"
	"github.com/WoodoCoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Woodo\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:   Set the PORT of the Server\n")
	fmt.Printf("-mode:   Choose between 'HTML' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the Server")
	mode := flag.String("mode", "rest", "Choose between 'HTML' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
