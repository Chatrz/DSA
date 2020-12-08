package main

import (
	"errors"
	"fmt"
)

type HashNode struct {
	Key   int
	Value int
}

type HashTable struct {
	Size     int
	Capacity int
	Array    []*HashNode
	Dummy    *HashNode
}

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		Size:     0,
		Capacity: capacity,
		Array:    make([]*HashNode, capacity),
		Dummy: &HashNode{
			Key:   -1,
			Value: -1,
		},
	}
}

func NewHashNode(key, value int) *HashNode {
	return &HashNode{
		Key:   key,
		Value: value,
	}
}

func (table *HashTable) hashCode(key int) int {
	return key % table.Capacity
}

func (table *HashTable) hashCodePrime(hashIndex int) int {
	hashIndex += 1
	hashIndex %= table.Capacity
	return hashIndex
}

func (table *HashTable) Insert(key, value int) error {
	if table.Size == table.Capacity {
		return errors.New("hash table is full !")
	}
	node := NewHashNode(key, value)
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil && table.Array[hashIndex] != table.Dummy {
		hashIndex = table.hashCodePrime(hashIndex)
	}
	if table.Array[hashIndex] == nil || table.Array[hashIndex] == table.Dummy {
		table.Array[hashIndex] = node
		table.Size = table.Size + 1
		return nil
	} else {
		return errors.New("error in finding empty slot !!")
	}
}

func (table *HashTable) Search(key int) *HashNode {
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil {
		if table.Array[hashIndex].Key == key {
			return table.Array[hashIndex]
		}
		hashIndex = table.hashCodePrime(hashIndex)
	}
	return nil
}

func (table *HashTable) Delete(key int) error {
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil {
		if table.Array[hashIndex].Key == key {
			table.Array[hashIndex] = table.Dummy
			return nil
		}
		hashIndex = table.hashCodePrime(hashIndex)
	}
	return errors.New("key not found ! ")
}

func (table *HashTable) Display() {
	for i := 0; i < table.Capacity; i++ {
		fmt.Print("for slot ")
		fmt.Print(i)
		fmt.Print(" : ")
		fmt.Println(table.Array[i])
		fmt.Println("###########################################")
	}
}

/*func main() {
	table := NewHashTable(10)
	table.Insert(10, 20)
	table.Insert(20, 21)
  table.Insert(30, 31)
  table.Insert(40, 41)
  table.Insert(50, 51)
  table.Insert(60, 61)
  table.Insert(70, 71)
  table.Insert(80, 81)
  table.Insert(90, 91)
  table.Insert(100,1001)
	table.Display()
  table.Delete(30)
  table.Display()
	fmt.Println(table.Search(40))
}*/
