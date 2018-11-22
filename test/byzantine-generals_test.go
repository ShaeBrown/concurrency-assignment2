package test

import (
	"concurrency-assignment2/byzantinegenerals"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ATTACK = byzantinegenerals.ATTACK
var RETREAT = byzantinegenerals.RETREAT

func TestByzantine_m2_n7(t *testing.T) {
	loyalty := []bool{true, true, true, true, true, false, false}
	result := byzantinegenerals.OM(2, ATTACK, loyalty)
	assert.Equal(t, ATTACK, result)
	assert.Equal(t, ATTACK, byzantinegenerals.Result["01"])
	assert.Equal(t, ATTACK, byzantinegenerals.Result["02"])
	assert.Equal(t, ATTACK, byzantinegenerals.Result["03"])
	assert.Equal(t, ATTACK, byzantinegenerals.Result["04"])
	assert.Equal(t, RETREAT, byzantinegenerals.Result["05"])
	assert.Equal(t, RETREAT, byzantinegenerals.Result["06"])
}

func TestByzantine_m1_tie(t *testing.T) {
	loyalty := []bool{true, true, false}
	result := byzantinegenerals.OM(1, ATTACK, loyalty)
	fmt.Printf("%v\n", byzantinegenerals.Result)
	assert.Equal(t, RETREAT, result)
}

func TestByzantine_m1_attack(t *testing.T) {
	loyalty := []bool{true, true, true, false}
	result := byzantinegenerals.OM(1, ATTACK, loyalty)
	assert.Equal(t, ATTACK, result)
}

func TestByzantine_m1_retreat(t *testing.T) {
	loyalty := []bool{true, true, true, false}
	result := byzantinegenerals.OM(1, RETREAT, loyalty)
	assert.Equal(t, RETREAT, result)
}

func TestByzantine_m0_attack(t *testing.T) {
	loyalty := []bool{true, true, true}
	result := byzantinegenerals.OM(0, ATTACK, loyalty)
	assert.Equal(t, ATTACK, result)
}

func TestByzantine_m0_retreat(t *testing.T) {
	loyalty := []bool{true, true, true}
	result := byzantinegenerals.OM(0, RETREAT, loyalty)
	assert.Equal(t, RETREAT, result)
}
