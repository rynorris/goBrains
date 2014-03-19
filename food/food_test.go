/*
 * Food tasting.  Yum.
 */

package food

import (
	"github.com/DiscoViking/goBrains/locationmanager"
	"math"
	"testing"
)

// Basic food verification.
func TestFood(t *testing.T) {

	var food *Food
	cm := locationmanager.NewLocationManager()

	testContents := []float64{
		0, // Empty food object.
		16,
		30,
		64,
	}

	for _, val := range testContents {
		t.Log("Test food with contents:", val)
		food = New(cm, val)

		// Content should be as entered.
		checkContent(t, food, val)

		// Radius is the square root of the content.
		checkRadius(t, food, math.Sqrt(val))

		// Test simple food consumption.
		checkConsumption(t, food, val)

		// Reset.
		food = New(cm, val)

		// Ensure that we cannot eat more food than there is in the instance.
		checkEmptying(t, food, val)
	}
}

// Content checking.
func checkContent(t *testing.T, food *Food, content float64) {
	if food.content != content {
		t.Errorf("Expected content of %v, found %v.", food.content, content)
	}
}

// Radius checking.
func checkRadius(t *testing.T, food *Food, radius float64) {
	if food.GetRadius() != radius {
		t.Errorf("Expected radius of %v, found %v.", food.GetRadius(), radius)
	}
}

// Food consumption.
func checkConsumption(t *testing.T, food *Food, content float64) {
	resp := food.Consume()

	switch {
	case (content > 0) && (resp == 0):
		t.Errorf("Expected consumption of food, but none reported.")
	case (content == 0) && (resp != 0):
		t.Errorf("Expected no comsumption of food, but %v reported.", resp)
	}
}

// Check being greedy.  If we try and eat all the food, does it stop?
func checkEmptying(t *testing.T, food *Food, content float64) {
	var ii float64
	currCont := content

	for ii = 0; ii < (content + 2); ii++ {
		checkConsumption(t, food, currCont)

		// Food consumed.
		if currCont != 0 {
			currCont--
		}

		checkContent(t, food, currCont)
	}

	// Radius should zero now.
	checkRadius(t, food, 0)
}
