package entitymanager

import "github.com/DiscoViking/goBrains/entity"

type Manager interface {
	LocationManager()
	Reset()
	Spin()
	Entities() []entity.Entity
}
