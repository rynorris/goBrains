package web

import (
	"testing"

	"github.com/DiscoViking/goBrains/config"
	"github.com/DiscoViking/goBrains/creature"
	"github.com/DiscoViking/goBrains/entity"
	"github.com/DiscoViking/goBrains/food"
	io "github.com/DiscoViking/goBrains/iomanager"
	lm "github.com/DiscoViking/goBrains/locationmanager"
)

func TestMarshalOne(t *testing.T) {
	l := lm.NewLocationManager(200, 200)
	testcases := map[io.DrawSpec]string{
		io.DrawSpec{
			creature.NewSimple(l), lm.Combination{10, 20, 2.5},
		}: `{"Type":"creature","X":10,"Y":20,"Colour":{"R":200,"G":50,"B":50,"A":255},"Angle":2.5}`,

		io.DrawSpec{
			food.New(l, 100), lm.Combination{10, 20, 2.5},
		}: `{"Type":"food","X":10,"Y":20,"Size":10}`,

		io.DrawSpec{
			&entity.TestEntity{}, lm.Combination{10, 20, 2.5},
		}: `{"Type":"unknown","X":10,"Y":20}`,
	}

	for in, out := range testcases {
		marshaled, err := marshalOne(in)

		if err != nil {
			t.Errorf("Error marshaling %v.\nMessage: %v.\n", in, err)
		}

		if marshaled != out {
			t.Errorf("Expected: %v\nGot: %V\n", out, marshaled)
		}
	}
}

func TestMarshal(t *testing.T) {
	config.Load("../../config/test_config.gcfg")
	l := lm.New()

	data := []io.DrawSpec{
		io.DrawSpec{
			creature.NewSimple(l), lm.Combination{10, 20, 2.5},
		},

		io.DrawSpec{
			food.New(l, 100), lm.Combination{10, 20, 2.5},
		},
	}

	exp := `{"scene":{"width":"800","height":"800","entities":[{"Type":"creature","X":10,"Y":20,"Colour":{"R":200,"G":50,"B":50,"A":255},"Angle":2.5},{"Type":"food","X":10,"Y":20,"Size":10}]}}`

	got := marshal(data)

	if got != exp {
		t.Errorf("Expected: %v\nGot: %v\n", exp, got)
	}
}
