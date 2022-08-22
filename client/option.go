package client

import "time"

type RecvOption struct {
	Timeout time.Duration
}

func NewDefaultRecvOption() *RecvOption {
	return &RecvOption{Timeout: 3 * time.Second}
}
