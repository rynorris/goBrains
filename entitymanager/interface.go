package entitymanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/locationmanager"
)

type Manager interface {
	LocationManager() locationmanager.LM
	Reset()
	Spin()
	Entities() []entity.Entity
}
