package hashmap

import (
	"strconv"

	"crypto/sha256"
	"encoding/hex"

	"github.com/zhuzeyu/structure/linkedlist"
)

const Length = 16

type KV struct {
	Key   string
	Value string
}

type HashMap struct {
	Buckets [Length]*linkedlist.LinkedList
}

func NewHashMap() *HashMap {
	hashMap := &HashMap{}
	for i := 0; i < Length; i++ {
		hashMap.Buckets[i] = linkedlist.NewLinkedList()
	}
	return hashMap
}

func Hash(key string) int64 {
	b := []byte(key)
	h := sha256.New()
	h.Write(b)
	sr := h.Sum(nil)
	res, err := strconv.ParseInt(hex.EncodeToString(sr[:]), 16, 64)
	if err != nil {
		return res % Length
	}
	return 0
}

func (hashMap *HashMap) Set(key string, value string) bool {
	index := Hash(key)
	cur := hashMap.Buckets[index].Head
	for cur != nil {
		if cur.Data.(KV).Key == key {
			cur.Data.(KV).Value = value
			return true
		}
		cur = cur.Next
	}
	hashMap.Buckets[index].Append(&KV{Key: key, Value: value})
	return true
}

func (hashMap *HashMap) Get(key string) (string, bool) {
	index := Hash(key)
	cur := hashMap.Buckets[index].Head
	for cur != nil {
		if cur.Data.(KV).Key == key {
			return cur.Data.(KV).Value, true
		}
		cur = cur.Next
	}
	return "", false
}

func (hashMap *HashMap) Remove(key string) bool {
	index := Hash(key)
	cur := hashMap.Buckets[index].Head
	for cur != nil {
		if cur.Data.(KV).Key == key {
			hashMap.Buckets[index].Remove(cur)
			return true
		}
		cur = cur.Next
	}
	return false
}
