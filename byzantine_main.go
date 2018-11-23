package main

import (
	"concurrency-assignment2/byzantinegenerals"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parse_command(Oc string) (byzantinegenerals.Command, error) {
	if Oc == "ATTACK" {
		return byzantinegenerals.ATTACK, nil
	} else if Oc == "RETREAT" {
		return byzantinegenerals.RETREAT, nil
	}
	return 0, errors.New("Oc should be ATTACK or RETREAT")
}

func parse_loyalty(G string) ([]bool, error) {
	l := strings.Split(G, ",")
	if len(l) < 2 {
		return nil, errors.New("Less than 2 generals." +
			"Provide list in comma seperated values of \"L\" or \"T\"")
	}
	result := make([]bool, len(l))
	for i, loyalty := range l {
		if loyalty == "L" {
			result[i] = true
		} else if loyalty == "T" {
			result[i] = false
		} else {
			return nil, errors.New("Invalid value." +
				"Provide list in comma seperated values of \"L\" or \"T\"")
		}
	}
	return result, nil
}

func main() {
	m := flag.Int("m", 0, "The level of recursion")
	G := flag.String("G", "L,L,L", "The loyalty of the generals."+
		"Provide list in comma seperated values of \"L\" or \"T\"")
	Oc := flag.String("Oc", "ATTACK", "The order the commander gives: ATTACK or RETREAT")
	flag.Parse()
	if *m < 0 {
		fmt.Println("m should be greater than 0")
		os.Exit(1)
	}
	command, err := parse_command(*Oc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	loyal, err := parse_loyalty(*G)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	consensus := byzantinegenerals.OM(*m, command, loyal)
	fmt.Printf("The general has come to concensus that they will %s\n", get_command_string(consensus))
}

func get_command_string(c byzantinegenerals.Command) string {
	if c == byzantinegenerals.ATTACK {
		return "ATTACK"
	} else if c == byzantinegenerals.RETREAT {
		return "RETREAT"
	}
	return "NONE"
}
