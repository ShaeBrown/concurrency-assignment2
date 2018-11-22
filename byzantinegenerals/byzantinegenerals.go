package byzantinegenerals

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/deckarep/golang-set"
)

type Command int

const (
	ATTACK  Command = 2
	RETREAT Command = 1
	UNSURE  Command = 0
	TIE     Command = 3
)

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

func create_order_array(n int) [][]Command {
	order := make([][]Command, n)
	for i := range order {
		order[i] = make([]Command, n)
	}
	for i, _ := range order {
		for j, _ := range order[i] {
			order[i][j] = UNSURE
		}
	}
	return order
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
	lieu := OM(*m, command, loyal)
	fmt.Printf("Commands that each lieutenant will make: %v\n", get_command_array(lieu))
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
	} else if c == UNSURE {
		return "UNSURE"
	} else if c == TIE {
		return "TIE"
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

func send_to_lieutenants(commander int, command Command, n int, orders [][]Command, lieu mapset.Set) {
	for _, el := range lieu.ToSlice() {
		orders[el.(int)][commander] = command
	}
}

func OM(m int, command Command, loyal []bool) []Command {
	orders := create_order_array(len(loyal))
	lieu := mapset.NewSet()
	for i := 1; i < len(loyal); i++ {
		lieu.Add(i)
	}
	return om(m, 0, command, loyal, orders, lieu)
}

func om(m int, commander int, command Command, loyal []bool, orders [][]Command, lieu mapset.Set) []Command {
	n := len(loyal)
	if m == 0 {
		send_to_lieutenants(commander, command, n, orders, lieu)
		return use_command(n, command)
	} else {
		send_to_lieutenants(commander, command, n, orders, lieu)
		recursive_step(m, commander, command, loyal, orders, lieu)
		maj := get_majority(orders, commander, loyal)
		return maj
	}
}

func use_command(n int, command Command) []Command {
	order := make([]Command, n-1)
	for i := range order {
		order[i] = command
	}
	return order
}

func recursive_step(m int, commander int, command Command, loyal []bool, orders [][]Command, lieu mapset.Set) {
	for _, el := range lieu.ToSlice() {
		i := el.(int)
		c := mapset.NewSetFromSlice(lieu.ToSlice())
		c.Remove(i)
		om(m-1, i, send_command(loyal[i], command), loyal, orders, c)
	}
}

func get_majority(orders [][]Command, commander int, loyalty []bool) []Command {
	majority := make([]Command, len(orders)-1)
	k := 0
	for i, lieu := range orders {
		if i != commander {
			if !loyalty[i] {
				majority[k] = UNSURE
				k += 1
				continue
			}
			num_attack := 0
			num_retreat := 0
			for j, command := range lieu {
				if i != j {
					if command == ATTACK {
						num_attack += 1
					} else if command == RETREAT {
						num_retreat += 1
					}
				}
			}
			var result Command
			if num_attack > num_retreat {
				result = ATTACK
			} else if num_retreat > num_attack {
				result = RETREAT
			} else {
				result = TIE
			}
			majority[k] = result
			k += 1
		}
	}
	return majority
}
