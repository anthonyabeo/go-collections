package hashed

type HashSet struct {
	set *HashTable
}

func NewHashSet(cap int) *HashSet {
	return &HashSet{
		&HashTable{
			make([]interface{}, cap),
			0,
			cap,
			0.0,
		},
	}
}

func (hs *HashSet) Size() int {
	return hs.set.numItems
}

func (hs *HashSet) IsEmpty() bool {
	return hs.set.numItems == 0
}

func (hs *HashSet) Contains(key interface{}) bool {
	e, err := hs.set.Get(key)
	if e != nil && err == nil {
		return true
	}

	return false
}

func (hs *HashSet) Add(key interface{}) error {
	return hs.set.Put(key, nil)
}

func (hs *HashSet) Remove(key interface{}) (interface{}, error) {
	return hs.set.Remove(key)
}

func (hs *HashSet) Empty() {
	hs.set = &HashTable{
		make([]interface{}, 13),
		0,
		13,
		0.0,
	}
}
