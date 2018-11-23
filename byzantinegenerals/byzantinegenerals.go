package byzantinegenerals

import (
	"strconv"

	"github.com/deckarep/golang-set"
)

type Command int

const (
	ATTACK  Command = 2
	RETREAT Command = 1
	NONE    Command = 0
)

var Result = make(map[string]Command)

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
