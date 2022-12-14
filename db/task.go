package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		id := int(id64)
		key := intToByte(id)
		return bucket.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, err
}

func Alltasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		c := bucket.Cursor()

		for key, val := c.First(); key != nil; key, val = c.Next() {
			tasks = append(tasks, Task{
				Key:   byteToInt(key),
				Value: string(val),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		return bucket.Delete(intToByte(key))
	})
}

func intToByte(v int) []byte {
	//BigEndian as oldest entries will be the first ones
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
