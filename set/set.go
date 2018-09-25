package set

import "sync"

//Set
type Set struct {
	set map[interface{}]bool
	sync.RWMutex
}

//Add element
func (s *Set) Add(a interface{}) {
	s.Lock()
	defer s.Unlock()
	s.set[a] = true
}

//Remove element
func (s *Set) Remove(key interface{}) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.set[key]
	if ok {
		delete(s.set, key)
	}
}

//Clear all Set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	for k, _ := range s.set {
		delete(s.set, k)
	}
}

//HasValue in Set?
func (s *Set) HasValue(key interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.set[key]
	return ok
}

//ToList convert to slice
func (s *Set) ToList() interface{} {
	s.RLock()
	defer s.RUnlock()
	var list []interface{}
	for k, _ := range s.set {
		list = append(list, k)
	}
	return list
}

//Len of set
func (s *Set) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.set)
}

// 使用sync map实现
type SyncSet struct {
	set sync.Map
}

func (s *SyncSet) Add(key interface{}) {
	s.set.Store(key, true)
}

func (s *SyncSet) Remove(key interface{}) {
	_, ok := s.set.Load(key)
	if ok {
		s.set.Delete(key)
	}
}

func (s *SyncSet) HasValue(key interface{}) bool {
	_, ok := s.set.Load(key)
	return ok
}

func (s *SyncSet) Clear() {
	s.set.Range(func(k, v interface{}) bool {
		s.set.Delete(k)
		return true
	})
}

func (s *SyncSet) ToList() []interface{} {
	var list []interface{}
	s.set.Range(func(k, _ interface{}) bool {
		list = append(list, k)
		return true
	})
	return list
}

func (s *SyncSet) Len() int {
	var len int
	s.set.Range(func(_, _ interface{}) bool {
		len++
		return true
	})
	return len
}
