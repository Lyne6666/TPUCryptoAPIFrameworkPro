// cmd/tpucryptoapiframeworkpro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"tpucryptoapiframeworkpro/internal/tpucryptoapiframeworkpro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := tpucryptoapiframeworkpro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
