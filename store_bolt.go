package diffcalculator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const (
	envDbPath       = "DB_PATH"
	openFileTimeout = time.Second * 5
	boltBucket      = "items2"
)

type boltStore struct {
	db *bolt.DB
}

// NewBoltStore creates new store
func NewBoltStore() Store {
	s := &boltStore{}
	if err := s.open(); err != nil {
		log.Fatalf("[ERROR] Error opening database: %s", err)
	}
	return s
}

func (s *boltStore) open() error {
	var err error
	s.db, err = bolt.Open(os.Getenv(envDbPath), 0644, &bolt.Options{Timeout: openFileTimeout})
	if err != nil {
		s.db = nil
		return fmt.Errorf("error opening db file: %s", err)
	}

	err = s.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(boltBucket)); err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error initializing db: %s", err)
	}

	return nil
}

// Get return items for this sourceID
func (s *boltStore) Get(sourceID string) (map[string]interface{}, error) {
	items := map[string]interface{}{}
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(boltBucket))
		value := bucket.Get([]byte(sourceID))

		if value != nil {
			if err := json.Unmarshal(value, &items); err != nil {
				return fmt.Errorf("error unmarshaling json: %s", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("tx error: %s", err)
	}

	return items, nil
}

// Save saves items
func (s *boltStore) Save(sourceID string, items map[string]interface{}) error {
	itemsBytes, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("error marshaling: %s", err)
	}

	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(boltBucket))
		err := bucket.Put([]byte(sourceID), itemsBytes)
		if err != nil {
			return fmt.Errorf("error putting to bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("tx error: %s", err)
	}

	return nil
}
