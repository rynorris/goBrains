package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/events"
	"github.com/DiscoViking/goBrains/locationmanager"
)

type IoManager interface {
	Start(lm *locationmanager.Location, done chan struct{}, event chan events.Event) chan []*entity.Entity
}
