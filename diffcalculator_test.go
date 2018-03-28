package diffcalculator

import "testing"

func TestCalculate(t *testing.T) {
	items := []Item{
		{
			ID: "1",
			Data: map[string]interface{}{
				"Foo": "1",
			},
		},
		{
			ID: "2",
			Data: map[string]interface{}{
				"Bar": "1",
			},
		},
	}
	err := Calculate("twitch", items)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}

	items = []Item{
		{
			ID: "2",
			Data: map[string]interface{}{
				"Foo": "1",
			},
		},
		{
			ID: "3",
			Data: map[string]interface{}{
				"Bar": "1",
			},
		},
	}
	err = Calculate("twitch", items)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}
}
