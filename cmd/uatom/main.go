package main

import (
	"log"
	"os"
	"strconv"
)

const multiplier int = 1000000

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Println("usage: uatom [number of atoms]")
		os.Exit(1)
	}

	atoms, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	uatoms := atoms * multiplier
	log.Printf("%d atoms = %d uatoms", atoms, uatoms)

	os.Exit(0)
}
