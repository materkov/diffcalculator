package diffcalculator

import "testing"

func TestCalculate(t *testing.T) {
	posts := []Post{
		{
			Id:         1,
			Message:    "1",
			ExternalId: "1-1",
		},
		{
			Id:         2,
			Message:    "2",
			ExternalId: "1-2",
		},
	}
	err := Calculate("twitch", posts)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}

	posts = []Post{
		{
			Id:         2,
			Message:    "2",
			ExternalId: "1-2",
		},
		{
			Id:         2,
			Message:    "2",
			ExternalId: "1-3",
		},
	}
	err = Calculate("twitch", posts)
	if err != nil {
		t.Fatalf("error not expected: %s", err)
	}
}
