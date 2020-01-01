package main

type HashSet struct {
	data map[interface{}]struct{}
}

func (hs *HashSet) Add(obj interface{}) {
	if hs.data == nil {
		hs.data = make(map[interface{}]struct{}, 0)
	}
	hs.data[obj] = struct{}{}
}

func (hs *HashSet) Contains(obj interface{}) bool {
	_, ok := hs.data[obj]
	return ok
}

func (hs *HashSet) Delete(obj interface{}) {
	delete(hs.data, obj)
}

func (hs *HashSet) Count() int {
	return len(hs.data)
}

func (hs *HashSet) Union(other HashSet) HashSet {
	newSet := HashSet{}
	for k, _ := range hs.data {
		newSet.Add(k)
	}
	for k, _ := range other.data {
		newSet.Add(k)
	}
	return newSet
}
