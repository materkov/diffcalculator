package diffcalculator

type memoryStore struct {
	posts map[string][]Post
}

// NewMemoryStore creates new store
func NewMemoryStore() Store {
	return &memoryStore{
		posts: make(map[string][]Post, 0),
	}
}

// Get return posts
func (s *memoryStore) Get(sourceID string) ([]Post, error) {
	return s.posts[sourceID], nil
}

// Save saves posts
func (s *memoryStore) Save(sourceID string, posts []Post) error {
	s.posts[sourceID] = posts
	return nil
}
