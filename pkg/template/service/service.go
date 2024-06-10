package service

var SERVICE_TPL = `package service

type Probe struct{}

func NewProbe() *Probe {
	return &Probe{}
}

func (p *Probe) Liveness() error {
	// TODO: add liveness check
	return nil
}
`
