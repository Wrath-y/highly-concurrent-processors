package core

import "time"

type Payload struct {
	Data string
}

func (p *Payload) Handler() error {
	time.Sleep(time.Second)
	println(p.Data)
	return nil
}
