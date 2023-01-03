package main

import (
	"log"
	"os"
	"strconv"
)

const divider float64 = 1000000

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Println("usage: atom [number of uatoms]")
		os.Exit(1)
	}

	uatoms, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	atoms := uatoms / divider
	log.Printf("%d uatoms = %f atoms", int(uatoms), atoms)

	os.Exit(0)
}
