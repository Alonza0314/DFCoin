package bboltdb

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

func DbSave(dbPath string, bucket, key, value []byte) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(fmt.Sprintf("DbSave error: %v", err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(fmt.Sprintf("DbSave error: %v", err))
		}
	}()

	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}

func DbLoad(dbPath string, bucket, key []byte) ([]byte, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(fmt.Sprintf("DbLoad error: %v", err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(fmt.Sprintf("DbLoad error: %v", err))
		}
	}()

	var value []byte
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		v := b.Get(key)
		value = make([]byte, len(v))
		copy(value, v)
		return nil
	}); err != nil {
		return nil, err
	}

	return value, nil
}

func DbUpdate(dbPath string, bucket, key, value []byte) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(fmt.Sprintf("DbUpdate error: %v", err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(fmt.Sprintf("DbUpdate error: %v", err))
		}
	}()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.Put(key, value)
	})
}

func DbRemove(dbPath string, bucket, key []byte) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(fmt.Sprintf("DbRemove error: %v", err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(fmt.Sprintf("DbRemove error: %v", err))
		}
	}()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.Delete(key)
	})
}
