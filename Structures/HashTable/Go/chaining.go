package main

import (
	"fmt"
	list "github.com/sinamna/DSA/Structures/LinkedList/Go/armin"
)

type item struct {
	key   int
	value int
}

type HashTable struct {
	size  int
	array []list.LinkedList
}

func (table *HashTable) hashCode(key int) int {
	return key % table.size
}

func NewHashTable(size int) *HashTable {
	table := &HashTable{
		size:  size,
		array: make([]list.LinkedList, size),
	}
	return table
}

/*struct DataItem *search(int key) {
   //get the hash
   int hashIndex = hashCode(key);

   //move in array until an empty
   while(hashArray[hashIndex] != NULL) {

      if(hashArray[hashIndex]->key == key)
         return hashArray[hashIndex];

      //go to next cell
      ++hashIndex;

      //wrap around the table
      hashIndex %= SIZE;
   }

   return NULL;
}*/

func (table *HashTable) Insert(key int, value int) {
	item := item{}
	item.key = key
	item.value = value
	//get the hash
	hashIndex := table.hashCode(key)
	table.array[hashIndex].AddBack(item)
}

/*struct DataItem* delete(struct DataItem* item) {
   int key = item->key;

   //get the hash
   int hashIndex = hashCode(key);

   //move in array until an empty
   while(hashArray[hashIndex] != NULL) {

      if(hashArray[hashIndex]->key == key) {
         struct DataItem* temp = hashArray[hashIndex];

         //assign a dummy item at deleted position
         hashArray[hashIndex] = dummyItem;
         return temp;
      }

      //go to next cell
      ++hashIndex;

      //wrap around the table
      hashIndex %= SIZE;
   }

   return NULL;
}*/

func (table *HashTable) Display() {
	for i := 0; i < table.size; i++ {
    fmt.Print("for slot ")
    fmt.Print(i)
    fmt.Println(" :")
    table.array[i].Display()
    fmt.Println("###############################################")
	}
}

func main() {
	table := NewHashTable(10)
	table.Insert(2, 3)
  table.Display()
}
