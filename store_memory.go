package diffcalculator

type memoryStore struct {
	items map[string][]Item
}

// NewMemoryStore creates new store
func NewMemoryStore() Store {
	return &memoryStore{
		items: make(map[string][]Item, 0),
	}
}

// Get return items for sourceID
func (s *memoryStore) Get(sourceID string) ([]Item, error) {
	return s.items[sourceID], nil
}

// Save saves items
func (s *memoryStore) Save(sourceID string, items []Item) error {
	s.items[sourceID] = items
	return nil
}
