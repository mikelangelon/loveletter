package main

type state string

const (
	startGame  state = "start"
	cpuTurn    state = "cpu"
	playerTurn state = "player"
	endTurn    state = "turnEnd"
	endGame    state = "end"
)
