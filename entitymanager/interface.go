package entitymanager

import "github.com/DiscoViking/goBrains/entity"

type Manager interface {
	Reset()
	Spin()
	Entities() []entity.Entity
}
