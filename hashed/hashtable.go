package hashed

import (
	"go-collections/graphs/tree"
	"math"
	"math/rand"
	"time"
)

type entry struct {
	key interface{}
	value interface{}
	hashCode int64
}

type HashTable struct {
	bucketArray []interface{}
	numItems int
	capacity int
	loadFactor float32
}

func NewHashTable() *HashTable {
	return &HashTable{
		make([]interface{}, 13),
		0,
		13,
		0.0,
	}
}

func (ht *HashTable) IsEmpty() bool {return ht.numItems == 0}
func (ht *HashTable) Size() int {return ht.numItems}
func (ht *HashTable) Capacity() int {return ht.capacity}
func (ht *HashTable) BucketArray() []interface{} {return ht.bucketArray}
func (ht *HashTable) LoadFactor() float32 {return ht.loadFactor}

func (ht *HashTable) Put(key, value interface{}) {
	hc := hashCode(key)
	newEntry := &entry{key, value, hc}

	if ht.loadFactor <= 0.5 {
		bucketIdx := compress(ht.capacity, int(hc))
		doubleHashInsert(ht, bucketIdx, newEntry)
	} else {
		newCap := 2 * ht.capacity
		largerBucket := make([]interface{}, newCap)

		for i := 0; i < ht.capacity; i++ {
			if ht.bucketArray[i] != nil {
				entry := ht.bucketArray[i].(*entry)
				newIdx := compress(newCap, int(entry.hashCode))
				largerBucket[newIdx] = entry
			}
		}

		ht.bucketArray = largerBucket
		ht.capacity = newCap

		newBucketIdx := compress(newCap, int(hc))
		doubleHashInsert(ht, newBucketIdx, newEntry)
	}


	ht.updateLoadFactor()
	ht.numItems++
}

func (ht *HashTable) Get(key interface{}) interface{} {
	idx := compress(ht.capacity, int(hashCode(key)))
	if ht.bucketArray[idx] == nil {
		return nil
	}

	return ht.bucketArray[idx].(*entry).value
}

func (ht *HashTable) Remove(key interface{}) interface{} {
	return nil
}

func (ht *HashTable) KeySet() tree.BinarySearchTree {
	return tree.BinarySearchTree{}
}

func (ht *HashTable) Values() tree.BinarySearchTree {
	return tree.BinarySearchTree{}
}

func (ht *HashTable) EntrySet() tree.BinarySearchTree {
	return tree.BinarySearchTree{}
}

func (ht *HashTable) updateLoadFactor() {
	ht.loadFactor = float32(ht.numItems) / float32(ht.capacity)
}

func hashCode(k interface{}) int64 {
	var x int64

	switch k.(type) {
	case int:
		x = k.(int64)
	case string:
		for idx, char := range k.(string) {
			x += int64(float64(char) * math.Pow(33, float64(idx)))
		}
	case float32:
		x = int64(k.(float32))
	case uint:
		x = int64(k.(uint))
	case rune:
		x = int64(k.(rune))
	case byte:
		x = int64(k.(byte))
	}
	return x
}
func compress(cap int, hashCode int) int {
	p := nextPrimeAfter(cap)
	a := random(1, p)
	b := random(0, p-1)

	return ((a * hashCode) + b) % cap
}
func nextPrimeAfter(n int) int {
	for i := n+1; i < (n + 100); i++ {
		if isPrime(i) {
			return i
		}
	}

	return -1
}
func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func doubleHashInsert(ht *HashTable, idx int, e *entry)  {
	if ht.bucketArray[idx] == nil {
		ht.bucketArray[idx] = e
	} else {
		for i := 1; i <= 100; i++ {
			idx := (idx + (i * hPrime(idx))) % ht.capacity
			if ht.bucketArray[idx]  == nil {
				ht.bucketArray[idx] = e
				break
			}
		}
	}
}

func hPrime(k int) int {
	return 5 - (k % 5)
}