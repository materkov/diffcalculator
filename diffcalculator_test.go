package diffcalculator

import "testing"

func TestCalculate(t *testing.T) {
	items := []Item{
		{
			ID: "1-1",
			Data: map[string]interface{}{
				"ID":         1,
				"Message":    "1",
				"ExternalId": "1-1",
			},
		},
		{
			ID: "1-2",
			Data: map[string]interface{}{
				"ID":         2,
				"Message":    "2",
				"ExternalId": "1-2",
			},
		},
	}
	err := Calculate("twitch", items)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}

	items = []Item{
		{
			ID: "1-2",
			Data: map[string]interface{}{
				"ID":         2,
				"Message":    "2",
				"ExternalId": "1-2",
			},
		},
		{
			ID: "1-3",
			Data: map[string]interface{}{
				"Id":         2,
				"Message":    "2",
				"ExternalId": "1-3",
			},
		},
	}

	err = Calculate("twitch", items)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}
}
