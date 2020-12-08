package main

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
	return &node.Key
}

func (table *HashTable) Insert(key int, value int) {
	item := Item{}
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
	table.Insert(12, 3)
	table.Insert(2, 4)
	table.Insert(15, 5)
	table.Insert(21, 17)
	table.Insert(13, 9)
	table.Insert(14, 1)
	table.Insert(4, 31)
	table.Display()
  itemm := table.Search(12)
	fmt.Println(itemm.value)
}
