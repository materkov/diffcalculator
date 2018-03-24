package diffcalculator

type Store interface {
	Get(sourceID string) ([]Post, error)
	Save(sourceID string, posts []Post) error
}

var StdStore = NewBoltStore()
