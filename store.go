package diffcalculator

// Store is interface for storage
type Store interface {
	Get(sourceID string) ([]Post, error)
	Save(sourceID string, posts []Post) error
}

// StdStore is standart storage
var StdStore = NewBoltStore()
