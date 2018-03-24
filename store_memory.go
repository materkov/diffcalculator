package diffcalculator

type memoryStore struct {
	posts map[string][]Post
}

func NewMemoryStore() Store {
	return &memoryStore{
		posts: make(map[string][]Post, 0),
	}
}

func (s *memoryStore) Get(sourceId string) ([]Post, error) {
	return s.posts[sourceId], nil
}

func (s *memoryStore) Save(sourceId string, posts []Post) error {
	s.posts[sourceId] = posts
	return nil
}
