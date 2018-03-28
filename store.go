package diffcalculator

// Store is interface for storage
type Store interface {
	Get(sourceID string) ([]Item, error)
	Save(sourceID string, items []Item) error
}

// StdStore is standart storage
var StdStore = NewBoltStore()
