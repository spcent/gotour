package main

import (
	"fmt"
	"math/rand"
)

const (
	win = 100	// The winning score in a game of Pig
	gamePerSeries = 10 // The number of games per series to simulate
)

// A score includes scores accmulated in previous turns for each players,
// as well as the points scored by the current player in this turn
type score struct {
	player, opponent, thisTurn int
}

// An action transitions stochastically to a resulting score.
type action func(current score) (result score, turnIsOver bool)

