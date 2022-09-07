package hashtable

import "fmt"

const BUCKETS_NUMBER = 7

//Hashtable
type HashTable struct {
	ArrayOfBuckets [BUCKETS_NUMBER]*Bucket
}

//Bucket
type Bucket struct {
	HeadNode *BucketNode
}

//BucketNodes
type BucketNode struct {
	value       string
	NextNodeRef *BucketNode
}

func Hash(value string) int {
	sum := 0
	for s := range value {
		sum += s
	}

	return sum % BUCKETS_NUMBER
}

func Init() *HashTable {
	hashTables := &HashTable{}

	for i := range hashTables.ArrayOfBuckets {
		hashTables.ArrayOfBuckets[i] = &Bucket{}
	}
	return hashTables
}

func (h *HashTable) Insert(val string) {
	key := Hash(val)
	h.ArrayOfBuckets[key].Insert(val)
}

func (h *HashTable) Search(val string) bool {
	key := Hash(val)
	return h.ArrayOfBuckets[key].Search(val)
}

func (h *HashTable) Delete(val string) {
	key := Hash(val)
	h.ArrayOfBuckets[key].Delete(val)
}

func (b *Bucket) Insert(val string) {
	if !b.Search(val) {
		newNode := &BucketNode{
			value: val,
		}
		newNode.NextNodeRef = b.HeadNode
		b.HeadNode = newNode
	} else {
		fmt.Println(val, " already Exists")
	}

}

func (b *Bucket) Search(val string) bool {
	currentNode := b.HeadNode
	for currentNode != nil {
		if currentNode.value == val {
			return true
		}
		currentNode = currentNode.NextNodeRef
	}
	return false
}

func (b *Bucket) Delete(val string) {

	if b.HeadNode.value == val {
		b.HeadNode = b.HeadNode.NextNodeRef
		return
	}

	prevNode := b.HeadNode
	for prevNode.NextNodeRef != nil {
		if prevNode.NextNodeRef.value == val {
			//delete
			prevNode.NextNodeRef = prevNode.NextNodeRef.NextNodeRef
		}
		prevNode = prevNode.NextNodeRef
	}
}
