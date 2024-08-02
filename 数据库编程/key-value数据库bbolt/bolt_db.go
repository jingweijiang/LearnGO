package kvdb

import (
	"errors"
	bolt "go.etcd.io/bbolt"
	"sync/atomic"
)

var NoDataErr = errors.New("no data")

type Bolt struct {
	db     *bolt.DB
	path   string
	bucket []byte
}

// builder模式
func (s *Bolt) WithDataPath(path string) *Bolt {
	s.path = path
	return s
}

func (s *Bolt) WithBucket(bucket string) *Bolt {
	s.bucket = []byte(bucket)
	return s
}

func (s *Bolt) Open() error {
	var err error
	//0o600=384
	s.db, err = bolt.Open(s.path, 0o600, bolt.DefaultOptions)
	if err != nil {
		return err
	}
	err = s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(s.bucket)
		return err
	})

	if err != nil {
		s.db.Close()
		return err
	}
	return nil
}

func (s *Bolt) GetDBPath() string {
	return s.path
}

func (s *Bolt) WALName() string {
	return s.db.Path()

}

func (s *Bolt) Close() error {
	return s.db.Close()
}

func (s *Bolt) Set(k, v []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(s.bucket).Put(k, v)
	})
}

func (s *Bolt) BatchSet(keys, vals [][]byte) error {
	if len(keys) != len(vals) {
		return errors.New("keys and vals length not equal")
	}

	var err error
	s.db.Batch(func(tx *bolt.Tx) error {
		for i, key := range keys {
			value := vals[i]
			err = tx.Bucket(s.bucket).Put(key, value)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (s *Bolt) Get(k []byte) ([]byte, error) {
	var ival []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		ival = tx.Bucket(s.bucket).Get(k)
		return nil

	})
	if len(ival) == 0 || err != nil {
		return ival, NoDataErr
	}
	return ival, err
}

func (s *Bolt) BatchGet(keys [][]byte) ([][]byte, error) {
	var err error
	values := make([][]byte, len(keys))
	s.db.Batch(func(tx *bolt.Tx) error {
		for i, key := range keys {
			ival := tx.Bucket(s.bucket).Get(key)
			values[i] = ival
		}
		return nil

	})
	return values, err
}

func (s *Bolt) Delete(k []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(s.bucket).Delete(k)
	})
}

func (s *Bolt) BatchDelete(keys [][]byte) error {
	var err error
	s.db.Batch(func(tx *bolt.Tx) error {
		for _, key := range keys {
			err = tx.Bucket(s.bucket).Delete(key)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (s *Bolt) Has(k []byte) bool {

	var b []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b = tx.Bucket(s.bucket).Get(k)
		return nil
	})
	if err != nil || len(b) == 0 {
		return false
	}
	return true
}

func (s *Bolt) IterDB(fn func(k, v []byte) error) int64 {
	var total int64
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)
		// 迭代器
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := fn(k, v); err != nil {
				return err
			} else {
				atomic.AddInt64(&total, 1)
			}
		}
		return nil
	})
	return atomic.LoadInt64(&total)
}
