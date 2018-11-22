package test

import (
	"concurrency-assignment2/byzantinegenerals"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ATTACK = byzantinegenerals.ATTACK
var RETREAT = byzantinegenerals.RETREAT
var UNSURE = byzantinegenerals.UNSURE
var TIE = byzantinegenerals.TIE

func TestByzantineM2n7(t *testing.T) {
	loyalty := []bool{true, true, true, true, true, false, false}
	result := byzantinegenerals.OM(2, ATTACK, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{ATTACK, ATTACK, ATTACK, ATTACK, UNSURE, UNSURE}, result)
}

func TestByzantineM2n6(t *testing.T) {
	loyalty := []bool{true, false, false, true, true, true}
	result := byzantinegenerals.OM(2, ATTACK, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{UNSURE, UNSURE, RETREAT, RETREAT, RETREAT}, result)
}

func TestByzantineM1Tie(t *testing.T) {
	loyalty := []bool{true, true, false}
	result := byzantinegenerals.OM(1, ATTACK, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{TIE, UNSURE}, result)
}

func TestByzantineM1Attack(t *testing.T) {
	loyalty := []bool{true, true, true, false}
	result := byzantinegenerals.OM(1, ATTACK, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{ATTACK, ATTACK, UNSURE}, result)
}

func TestByzantineM1Retreat(t *testing.T) {
	loyalty := []bool{true, true, true, false}
	result := byzantinegenerals.OM(1, RETREAT, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{RETREAT, RETREAT, UNSURE}, result)
}

func TestByzantineM0Attack(t *testing.T) {
	loyalty := []bool{true, true, true}
	result := byzantinegenerals.OM(0, ATTACK, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{ATTACK, ATTACK}, result)
}

func TestByzantineM0Retreat(t *testing.T) {
	loyalty := []bool{true, true, true}
	result := byzantinegenerals.OM(0, RETREAT, loyalty)
	assert.Equal(t, []byzantinegenerals.Command{RETREAT, RETREAT}, result)
}
