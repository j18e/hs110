package main

import (
	"flag"
	"fmt"

	"github.com/j18e/hs110"

	log "github.com/sirupsen/logrus"
)

func main() {
	addr := flag.String("address", "", "ip address of the hs110 plug")
	flag.Parse()

	if *addr == "" {
		log.Fatal("required flag -address")
	}

	// log.SetLevel(log.DebugLevel)
	plug, err := hs110.NewPlug(*addr)
	if err != nil {
		log.Fatal(err)
	}

	switch flag.Arg(0) {
	case "on":
		err := plug.On()
		check(err)
	case "off":
		err := plug.Off()
		check(err)
	case "status":
		state, err := plug.Status()
		check(err)
		if state {
			fmt.Println("plug is on")
		} else {
			fmt.Println("plug is off")
		}
	default:
		log.Fatal("valid args [on off]")
	}
	fmt.Println("done")
}

func check(err error) {
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
}
