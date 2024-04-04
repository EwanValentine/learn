package limiters

type Adapter interface {
	IsAllowed(id string) bool
}

func NewLimiter(adapter Adapter) *Limiter {
	return &Limiter{adapter: adapter}
}

type Limiter struct {
	adapter Adapter
}

func (l *Limiter) IsAllowed(id string) bool {
	return l.adapter.IsAllowed(id)
}

