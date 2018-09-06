package set

import "sync"

//Set collection
type Set struct {
	set map[interface{}]bool
	sync.RWMutex
}

//Add element
func (s *Set) Add() {

}

//Remove element
func (s *Set) Remove() {

}

//Clear all Set
func (s *Set) Clear() {

}

//HasValue in Set?
func (s *Set) HasValue() {

}

//ToList convert to slice
func (s *Set) ToList() {

}

//Len of set
func (s *Set) Len() {

}
