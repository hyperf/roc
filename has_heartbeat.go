package gomul

const PING = "ping"
const PONG = "pong"

type HasHeartbeat interface {
	IsHeartbeat() bool
}
