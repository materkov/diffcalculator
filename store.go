package diffcalculator

// Store is interface for storage
type Store interface {
	Get(sourceID string) ([]Item, error)
	Save(sourceID string, items []Item) error
}

// StdStore is standard storage
var StdStore = NewBoltStore()
