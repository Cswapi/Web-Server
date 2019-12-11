package database

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var dbName = "Cswapi.db"
var db *bolt.DB

// 打开一个database
func Start(str string) {
	var err error
	dbName = str
	db, err = bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 关闭database
func Stop() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

// 初始化一个database
func Init(str string) {
	if _, err := os.Open(str); err == nil {
		log.Println("database is already exists!")
		return
	}
	Start(str)
	// 创建桶:键值对
	if err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte("users"))
		tx.CreateBucket([]byte("films"))
		tx.CreateBucket([]byte("people"))
		tx.CreateBucket([]byte("planets"))
		tx.CreateBucket([]byte("species"))
		tx.CreateBucket([]byte("starships"))
		tx.CreateBucket([]byte("vehicles"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	Stop()
}

// 更新数据库内容
func Update(bucketName []byte, key []byte, value []byte) {
	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(bucketName).Put(key, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

// 通过桶名和键获得值
func GetValue(bucketName []byte, key []byte) string {
	var result []byte
	if err := db.View(func(tx *bolt.Tx) error {
		keyLen := len(tx.Bucket([]byte(bucketName)).Get(key))
		result = make([]byte, keyLen)
		copy(result[:], tx.Bucket([]byte(bucketName)).Get(key)[:])
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return string(result)
}

// 检查键是否正确
func CheckKey(bucketName []byte, key []byte) bool {
	// 键的字节数
	var keyLen int
	if err := db.View(func(tx *bolt.Tx) error {
		keyLen = len(tx.Bucket([]byte(bucketName)).Get(key))
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return (keyLen != 0)
}

// 获得桶的数量
func GetBucketCount(bucketName []byte) int {
	count := 0
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			count++
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return count
}

// 检查桶
func CheckBucket(bucketName []byte) {
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
