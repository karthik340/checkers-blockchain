package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NextGameKey = "NextGame-value-"
)

const (
	StoredGameKey = "StoredGame-value-"
)

const (
	StoredGameEventKey     = "NewGameCreated"
	StoredGameEventCreator = "Creator"
	StoredGameEventIndex   = "Index"
	StoredGameEventRed     = "Red"
	StoredGameEventBlack   = "Black"
)

const (
	PlayMoveEventKey       = "MovePlayed"
	PlayMoveEventCreator   = "Creator"
	PlayMoveEventIdValue   = "IdValue"
	PlayMoveEventCapturedX = "CapturedX"
	PlayMoveEventCapturedY = "CapturedY"
	PlayMoveEventWinner    = "Winner"
)

const (
	RejectGameEventKey     = "GameRejected"
	RejectGameEventCreator = "Creator"
	RejectGameEventIdValue = "IdValue"
)

const (
	NoFifoIdKey = "-1"
)

const (
	MaxTurnDurationInSeconds = time.Duration(24 * 3_600 * 1000_000_000)
	DeadlineLayout           = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	ForfeitGameEventKey     = "GameForfeited"
	ForfeitGameEventIdValue = "IdValue"
	ForfeitGameEventWinner  = "Winner"
)

const (
	StoredGameEventWager = "Wager"
)

const (
	CreateGameGas = 10
	PlayMoveGas   = 10
	RejectGameGas = 0
)

const (
	StoredGameEventToken = "Token"
)
