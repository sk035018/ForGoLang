package main
 
import "fmt"
 
type Node struct {
    prev *Node
    next *Node
    value  interface{}
}
 
type List struct {
    head *Node
    tail *Node
}
 
func (L *List) add(value interface{}) {
    list := &Node{
        next: L.head,
        value:  value,
    }
    if L.head != nil {
        L.head.prev = list
    }
    L.head = list
 
    l := L.head
    for l.next != nil {
        l = l.next
    }
    L.tail = l
}
 
func (l *List) view() {
    list := l.head
    for list != nil {
        fmt.Print(list.value, " >> ")
        list = list.next
    }
    fmt.Println()
}
  
func main() {
    link := List{}
    link.add("Shivam")
    link.add("Gaurav")
    link.add("Ayush")
    link.add("Triyank")     
    
    fmt.Println("Head: " , link.head.value)
    fmt.Println("Tail: " , link.tail.value)
    link.view()
}