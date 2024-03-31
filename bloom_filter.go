package main

import (
	"crypto/sha256"
	"encoding/binary"
)

type BloomFilter struct {
	len    uint32
	filter []byte
}

func (bf *BloomFilter) init(len uint32) {
	bf.len = len
	bf.filter = make([]byte, bf.len)
}

func (bf *BloomFilter) isHit(key string) bool {
	hashValue := sha256.Sum256([]byte(key))
	//注意这里将[Size]byte数组转为切片
	index := binary.BigEndian.Uint32(hashValue[:]) % bf.len
	if bf.filter[index] != 1 {
		return false
	}
	bf.filter[index] = 1
	return true
}

func (bf *BloomFilter) initDB() {
	keys := []string{"aaa", "bbb", "ccc", "ddd"}
	for _, key := range keys {
		hashValue := sha256.Sum256([]byte(key))
		//注意这里将[Size]byte数组转为切片
		index := binary.BigEndian.Uint32(hashValue[:]) % bf.len
		bf.filter[index] = 1
	}
}
