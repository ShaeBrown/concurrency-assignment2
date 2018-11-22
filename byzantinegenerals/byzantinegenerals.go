package byzantinegenerals

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/deckarep/golang-set"
)

type Command int

const (
	ATTACK  Command = 2
	RETREAT Command = 1
	NONE    Command = 0
)

var Result = make(map[string]Command)

func parse_command(Oc string) (Command, error) {
	if Oc == "ATTACK" {
		return ATTACK, nil
	} else if Oc == "RETREAT" {
		return RETREAT, nil
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
	OM(*m, command, loyal)
}

func get_command_array(c []Command) []string {
	c_str := make([]string, len(c))
	for i, command := range c {
		c_str[i] = get_command_string(command)
	}
	return c_str
}

func get_command_string(c Command) string {
	if c == ATTACK {
		return "ATTACK"
	} else if c == RETREAT {
		return "RETREAT"
	}
	return "NONE"
}

func send_command(loyal bool, command Command) Command {
	if loyal {
		return command
	} else {
		return (command + 1) % 2
	}
}

func send_to_lieutenants(command Command, orders map[string]Command, lieu mapset.Set, prefix string) {
	for _, el := range lieu.ToSlice() {
		i := el.(int)
		p := prefix + strconv.Itoa(i)
		orders[p] = command
	}
}

func OM(m int, command Command, loyal []bool) Command {
	Result = make(map[string]Command)
	orders := make(map[string]Command)
	lieu := mapset.NewSet()
	for i := 1; i < len(loyal); i++ {
		lieu.Add(i)
	}
	return om(m, 0, command, loyal, orders, lieu, "0")
}

func om(m int, commander int, command Command, loyal []bool, orders map[string]Command, lieu mapset.Set, prefix string) Command {
	if m == 0 {
		send_to_lieutenants(command, orders, lieu, prefix)
		Result[prefix] = command
		return command
	} else {
		send_to_lieutenants(command, orders, lieu, prefix)
		recursive_step(m, commander, command, loyal, orders, lieu, prefix)
		return get_majority(orders, lieu, prefix)
	}
}

func use_command(n int, command Command) []Command {
	order := make([]Command, n-1)
	for i := range order {
		order[i] = command
	}
	return order
}

func recursive_step(m int, commander int, command Command, loyal []bool, orders map[string]Command, lieu mapset.Set, prefix string) {
	for _, el := range lieu.ToSlice() {
		i := el.(int)
		c := mapset.NewSetFromSlice(lieu.ToSlice())
		c.Remove(i)
		p := prefix + strconv.Itoa(i)
		om(m-1, i, send_command(loyal[i], command), loyal, orders, c, p)
	}
}

func get_majority(orders map[string]Command, lieu mapset.Set, prefix string) Command {
	num_attack := 0
	num_retreat := 0
	for _, el := range lieu.ToSlice() {
		i := el.(int)
		index := prefix + strconv.Itoa(i)
		if Result[index] == ATTACK {
			num_attack += 1
		} else if Result[index] == RETREAT {
			num_retreat += 1
		}
	}
	if num_attack > num_retreat {
		Result[prefix] = ATTACK
		return ATTACK
	} else {
		Result[prefix] = RETREAT
		return RETREAT
	}
}
