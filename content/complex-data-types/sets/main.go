package main

import "fmt"

type StringSet struct {
	set map[string]bool
}

func NewStringSet() StringSet {
	return StringSet{make(map[string]bool)}
}

func (ss StringSet) Add(value string) {
	ss.set[value] = true
}

func (ss StringSet) Remove(value string) {
	delete(ss.set, value)
}

func (ss StringSet) Contains(value string) bool {
	_, found := ss.set[value]
	return found
}

func (ss StringSet) Union(otherSet StringSet) StringSet {
	union := NewStringSet()
	for k, _ := range ss.set {
		union.set[k] = true
	}
	for k, _ := range otherSet.set {
		union.set[k] = true
	}
	return union
}

func (ss StringSet) Intersection(otherSet StringSet) StringSet {
	intersection := NewStringSet()
	for k, _ := range ss.set {
		if _, found := otherSet.set[k]; found {
			intersection.set[k] = true
		}
	}
	return intersection
}

func main() {
	set := NewStringSet()
	fmt.Println(set)
	fruits := NewStringSet()
	fruits.Add("blackberry")
	fruits.Contains("orange")
	fruits.Remove("blackberry")
	fmt.Println(fruits)
	vegetables := NewStringSet()
	vegetables.Add("eggplant")
	vegetables.Add("salad")
	union := fruits.Union(vegetables)
	fmt.Println(union)
	fruits.Add("tomato")
	vegetables.Add("tomato")
	interestion := fruits.Intersection(vegetables)
	fmt.Println(interestion)
}
