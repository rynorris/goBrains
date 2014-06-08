package web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/food"
	"github.com/DiscoViking/goBrains/iomanager"
)

// Turns a list of DrawSpecs into a json string to be sent over the wire.
// We have to do some work ourselves since we need to represent different
// entity types differently. However we use the json package to do some of the
// work.
func marshal(data []iomanager.DrawSpec) string {
	output := fmt.Sprintf(`{"scene":{"width":"%v","height":"%v","entities":[`,
		config.Global.General.ScreenWidth,
		config.Global.General.ScreenHeight)

	for i, spec := range data {
		s, err := marshalOne(spec)

		if err != nil {
			log.Println("Failed to marshal entity!")
			continue
		}

		output += s

		if i < len(data)-1 {
			output += ","
		}
	}

	output += "]}}"

	return output
}

// Marshal a single DrawSpec into json.
func marshalOne(spec iomanager.DrawSpec) (string, error) {
	var output []byte
	var err error

	switch e := spec.E.(type) {
	case *creature.Creature:
		s := creatureSpec{
			entitySpec{"creature", int(spec.Loc.X), int(spec.Loc.Y)},
			e.Color(),
			spec.Loc.Orient,
		}
		output, err = json.Marshal(s)

	case *food.Food:
		s := foodSpec{
			entitySpec{"food", int(spec.Loc.X), int(spec.Loc.Y)},
			int(e.Radius()),
		}
		output, err = json.Marshal(s)

	default:
		s := entitySpec{"unknown", int(spec.Loc.X), int(spec.Loc.Y)}
		output, err = json.Marshal(s)
	}

	return string(output), err
}
