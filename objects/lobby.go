package objects

import (
	"Proj/golang_pub-sub_observe/observer"
	"slices"
)

type lobby struct {
	players     []*player
	turn        int
	moves       map[int]*player
	subscribers []observer.Observer
}

func (l *lobby) Register(listener observer.Observer) {
	l.subscribers = append(l.subscribers, listener)
}

// Not sure it works
func (l *lobby) Deregister(listener observer.Observer) {
	idx := slices.IndexFunc(l.subscribers, func(o observer.Observer) bool { return o == listener })
	l.subscribers = slices.Delete(l.subscribers, idx, idx)
}

func (l *lobby) Notify(subject any) {
	for _, listener := range l.subscribers {
		listener.Observe(subject)
	}
}
