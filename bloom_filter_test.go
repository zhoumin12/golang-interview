package main

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	filter := &BloomFilter{}
	filter.init(10)
	filter.initDB()
	keys := []string{"aaa", "bbb"}
	for _, key := range keys {
		if filter.isHit(key) {
			fmt.Println(key, ":通过")
		} else {
			fmt.Println(key, ":未通过")
		}
	}
}
