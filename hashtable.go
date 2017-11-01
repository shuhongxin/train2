package hashtable
//hashtable的go语言实现，使用数组加链表的方法，实现了增加，删除，查询等功能

import (
	"container/list"
)

//键值对的结构体
type Node struct {
	key   int 
	value int
}

//哈希表的结构体
type HashTable struct {
	size  int
	array []*list.List
}	

//产生一个新的哈希表，size参数用于设置哈希表里面链表数组的大小
func CreateHashTable(size int) HashTable {
	hashtable := HashTable{}
	hashtable.size = size
	hashtable.array=make([]*list.List,size)
	for i := 0 ; i <= size ; i++{
		hashtable.array[i] = list.New()
	}
	return hashtable
}

//插入一个键值对
func (hashtable HashTable) Insert(newNode *Node) {
	temp := newNode.key%hashtable.size

	hashtable.array[temp].PushBack(newNode)
}

//删除一个键值对
func (hashtable HashTable) Remove(oldNode *Node) {
	temp := oldNode.key%hashtable.size
	current := hashtable.array[temp].Front()
	for current != nil {
		if(current.Value.(*Node) == oldNode) {
			hashtable.array[temp].Remove(current)
		}
		current = current.Next()
	}
}

//根据key值,查询value值
func (hashtable HashTable) GetValue(key int) *Node {
	temp := key%hashtable.size
	current := hashtable.array[temp].Front()
	for current != nil {
		node := current.Value.(*Node)
		if(node.key == key){
			return node
		}
				current = current.Next()
	}
	return nil
}