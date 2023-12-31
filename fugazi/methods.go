package sllist

type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

func main() {
	node := NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value: ", node.GetValue())
}

type SingleLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}
