//单链表的go语言实现，具有增加，删除，查询等功能
package link



//定义链表的一个节点结构体
type Node struct{
	next   *Node
	Value  interface{}
}

// 定义整个链表的结构体
type List struct {
	root  *Node
	len    int
}

//产生一个新的链表
func New() *List {
	return new(List)
}

//实现链表的增加操作
func (l *List) Add(v interface{}) {
	 n:=l.root
	 for n.next!=nil{
		 n=n.next
	 }
	 n.next=&Node{Value: v}
}

//实现链表的查询操作
func (l *List) Contains(v interface{}) bool {
	 n := l.root.next
	 for n!=nil{
		 if(n.Value == v){
			 return true
		 }
		 n=n.next
		}
	 return false        
}

//实现链表的删除操作
func (l *List) Remove(v interface{}) bool {
	n := l.root
	m := l.root.next
	for m != nil {
		if m.Value == v {
			n.next = m.next
			return true
		}
		n = m
		m = m.next
	}

	return false
}


