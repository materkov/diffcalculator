package diffcalculator

type memoryStore struct {
	items map[string]map[string]interface{}
}

// NewMemoryStore creates new store
func NewMemoryStore() Store {
	return &memoryStore{
		items: map[string]map[string]interface{}{},
	}
}

// Get return items for sourceID
func (s *memoryStore) Get(sourceID string) (map[string]interface{}, error) {
	return s.items[sourceID], nil
}

// Save saves items
func (s *memoryStore) Save(sourceID string, items map[string]interface{}) error {
	s.items[sourceID] = items
	return nil
}
