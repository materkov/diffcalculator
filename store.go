package diffcalculator

// Store is interface for storage
type Store interface {
	Get(sourceID string) (map[string]interface{}, error)
	Save(sourceID string, items map[string]interface{}) error
}

// StdStore is standard storage
var StdStore = NewBoltStore()
