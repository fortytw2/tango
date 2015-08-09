package dota

// LeaverStatus is the endgame status of a player in a match
type LeaverStatus int

const (
	// Finished the game, no problem
	Finished LeaverStatus = iota
	// DisconnectedNoAbandon - disconncted, but didn't abandon
	DisconnectedNoAbandon
	// DisconnectedAbandon - disconnected too long, abandoned
	DisconnectedAbandon
	// Abandoned - manually clicked Abandon
	Abandoned
	// NeverConnected - didn't connect
	NeverConnected
	// TooLongToConnect - took too long to connect
	TooLongToConnect
)
