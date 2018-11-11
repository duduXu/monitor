package util

import (
	"time"
)

type Pile struct {
	StartAt int64
	EndAt   int64
	ExcTime int64
}

func (p *Pile) Start() {
	p.StartAt = time.Now().UnixNano() / 1e6
	return
}

func (p *Pile) End() {
	p.EndAt = time.Now().UnixNano() / 1e6
	p.ExcTime = p.EndAt - p.StartAt
	return
}
