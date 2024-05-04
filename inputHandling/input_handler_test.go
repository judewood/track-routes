package inputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/models"
)

func testApplyFilter(t *testing.T) {
	var inputData = []models.RouteSection{
		{
			Start:        "abc  ",
			End:          "XYz",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ //duplicate of above
			Start:        "ABC",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but with start and end swapped
			Start:        "XYZ",
			End:          "ABC",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but not a passenger route
			Start:        "XYZ",
			End:          "ABC",
			Distance:     123,
			PassengerUse: "",
		},
		{
			Start:        "ABC",
			End:          "XYZ",
			Distance:     99,
			PassengerUse: "Y",
		},
		{
			Start:        "DEF",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{
			Start:        "ABC",
			End:          "RST",
			Distance:     123,
			PassengerUse: "not y",
		},
	}

	var expectedOutputData = []models.RouteSection{
		{
			Start:        "ABC",
			End:          "XYZ",
			Distance:     99,
			PassengerUse: "Y",
		},
		{
			Start:        "DEF",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{
			Start:        "ABC",
			End:          "RST",
			Distance:     123,
			PassengerUse: "not y",
		},
	}

	filter := &InputStruct{}

	res1 := filter.ApplyFilter(&inputData)
	res := *res1
	for i, v := range expectedOutputData {
		if v != res[i] {
			fmt.Println("filter test failed")
			t.Fail()
			return
		}
	}
}
