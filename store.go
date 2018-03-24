package diffcalculator

type Store interface {
	Get(sourceId string) ([]Post, error)
	Save(sourceId string, posts []Post) error
}

var StdStore = NewBoltStore()
