package entitymanager

import "github.com/DiscoViking/goBrains/entity"

type entityList map[entity.Entity]struct{}

func (l entityList) Check() {
	for e, _ := range l {
		if e.Check() {
			delete(l, e)
		}
	}
}

func (l entityList) Clear() {
	for e, _ := range l {
		delete(l, e)
	}
}

func (l entityList) Add(e entity.Entity) {
	l[e] = struct{}{}
}

func (l entityList) Slice() []entity.Entity {
	s := make([]entity.Entity, 0, len(l))
	for e, _ := range l {
		s = append(s, e)
	}
	return s
}

func NewList() entityList {
	return map[entity.Entity]struct{}{}
}
