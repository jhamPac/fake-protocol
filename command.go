package main

// ID type for commands as int
type ID int

// all possible commands avaiable for the fake protocol
const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
)

type command struct {
	id        ID
	recipient string
	sender    string
	body      string
}
