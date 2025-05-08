package bridge

import (
	"sync"
)

type Bridge struct {
	ports []*Port
	macTable map[string]*Port
	mu sync.Mutex
}

func NewBridge() *Bridge {
	return &Bridge{
		ports: []*Port{},
		macTable: make(map[string]*Port),
	}
}

func (b *Bridge) AddPort(name string) error {

}

func (b *Bridge) Run() {
	for _, port := range b.ports {
	}
}
