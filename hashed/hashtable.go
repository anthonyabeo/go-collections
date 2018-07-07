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

/// Creates a new Hashtable and initializes its fields to default values.
/// Capacity of the bucket array is initialized to the argument provided.
///
/// Args
///		cap - initial capacity of the bucket array
///
/// Return
///		*HashTable - a pointer to the hash table created
func NewHashTable(cap int) *HashTable {
	return &HashTable{
		make([]interface{}, cap),
		0,
		cap,
		0.0,
	}
}

func (ht *HashTable) IsEmpty() bool {return ht.numItems == 0}
func (ht *HashTable) Size() int {return ht.numItems}
func (ht *HashTable) Capacity() int {return ht.capacity}
func (ht *HashTable) BucketArray() []interface{} {return ht.bucketArray}
func (ht *HashTable) LoadFactor() float32 {return ht.loadFactor}

/// If the hash table, M does not have an entry with key equal to key, then adds entry
/// (k,v) to M and returns null; else, uses double hashing to find the next
/// bucket to insert the entry. Whenever the load factor is greater than
/// 0.5 the entries are rehashed and inserted into a larger bucket
///
/// Args
///		key -  the key of the item to be inserted
///		value - the value of the item to be inserted
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

/// returns the value associated with the specified key.
/// returns nil of there is no such value.
///
/// Args
///		key -  the key whose value is of interest
///
/// Return
///		value - the value associated with the key
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

/// Update the load factor of the hash table after a new item
/// is inserted into the table. This computed as the ratio of
/// the number of entries in the table to the capacity of the table.
func (ht *HashTable) updateLoadFactor() {
	ht.loadFactor = float32(ht.numItems) / float32(ht.capacity)
}

/// computes and returns the hashcode of a given object
///	if the object is a string, we compute its polynomial hash code
/// All other base types are converted to int64 and returned
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

/// compresses a hashcode into the range [0 - N] using the MAD method
/// where N is the capacity of the bucket array
///
/// Args
///		cap -  the capacity of the bucket array
///		hashCode - the hashCode to be compressed
///
/// Return
///
func compress(cap int, hashCode int) int {
	p := nextPrimeAfter(cap)
	a := random(1, p)
	b := random(0, p-1)

	return ((a * hashCode) + b) % cap
}

/// These are Utility functions that are used by other functions

/// finds the next prime after the specified integer, n
func nextPrimeAfter(n int) int {
	for i := n+1; i < (n + 100); i++ {
		if isPrime(i) {
			return i
		}
	}

	return -1
}

/// checks if a number, n is a prime
func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

/// generates a random number between in the range [min, - max]
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

/// Inserts a new entry into the table by searching for a
/// bucket using the double hashing strategy if the index
/// of the entry is already occupied, otherwise just insert
/// into the specified index
///
/// Args
///		ht - a pointer to the hashtable
///		idx - the index generated from compressing the entry's hashcode
///		e -	a pointer to the entry to be inserted
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

/// the second hash function used in double hashing
///
/// Args
///		k - the index of the entry
func hPrime(k int) int {
	return 5 - (k % 5)
}