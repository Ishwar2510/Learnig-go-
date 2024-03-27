package main

import "fmt"

type Set struct {
	mymap map[interface{}]interface{}
}

func newSet() *Set {
	return &Set{
		mymap: make(map[interface{}]interface{}),
	}
}
func (s *Set) Add(data interface{}) {
	if _, isAvailable := s.mymap[data]; !isAvailable {
		s.mymap[data] = struct{}{}
	}
}
func (s *Set) Remove(data interface{}) {
	if _, isAvailable := s.mymap[data]; isAvailable {
		delete(s.mymap, data)
	}
}

func (s *Set) Contains(data interface{}) bool {
	_, isAvailable := s.mymap[data]
	return isAvailable
}
func (s *Set) Printset() {
	var set []interface{}
	for key, _ := range s.mymap {
		set = append(set, key)
	}
	fmt.Println(set)
}
