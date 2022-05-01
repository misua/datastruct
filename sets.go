//sets - unordered,unique elements
package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set

}

//add elements to struct
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

//deletes and element from our set if it exist
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in the set")
	}
	delete(s.Elements, elem)
	return nil
}

//List prints out all the elements within our set
func (s *Set) List() {
	for k := range s.Elements {
		fmt.Println(k)
	}
}

//contains -  to check if element exists within the set
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists

}

func main() {

	mySet := NewSet()

	mySet.Add("Earth")
	mySet.Add("Mars")
	mySet.Add("Jupiter")
	mySet.List()
	fmt.Println(mySet.Contains("Earth"))

}
