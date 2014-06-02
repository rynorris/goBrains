package iomanager

import (
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/locationmanager"
)

func Locate(lm locationmanager.Location, e entity.Entity) (DrawSpec, bool) {
	ok, comb := lm.GetLocation(e)
	if !ok {
		return DrawSpec{}, false
	}

	return DrawSpec{e, comb}, true
}
