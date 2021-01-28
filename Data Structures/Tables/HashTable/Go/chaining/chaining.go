////////////////////////////////
//    Author : Armin Goodarzi //
//    GitHub :                //
//      github.com/Armingodiz //
//                            //
////////////////////////////////
package chaining

import (
	"fmt"
)

type Item struct {
	key   int
	value int
}

type HashTable struct {
	size  int
	array []LinkedList
}

func (table *HashTable) hashCode(key int) int {
	return key % table.size
}

func NewHashTable(size int) *HashTable {
	table := &HashTable{
		size:  size,
		array: make([]LinkedList, size),
	}
	return table
}

func (table *HashTable) Search(key int) *Item {
	//get the hash
	hashIndex := table.hashCode(key)
	node := table.array[hashIndex].Search(key)
	if node != nil {
		return &node.Key
	}
	return nil
}

func (table *HashTable) Insert(key int, value int) {
	item := Item{}
	item.key = key
	item.value = value
	//get the hash
	hashIndex := table.hashCode(key)
	table.array[hashIndex].AddBack(item)
}
func (table *HashTable) Delete(key int) *Item {
	hashIndex := table.hashCode(key)
	node := table.array[hashIndex].Delete(key)
	if node != nil {
		return &node.Key
	}
	return nil
}

func (table *HashTable) Display() {
	for i := 0; i < table.size; i++ {
		fmt.Print("for slot ")
		fmt.Print(i)
		fmt.Println(" :")
		table.array[i].Display()
		fmt.Println("###############################################")
	}
}
