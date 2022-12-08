package sllist

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T any] struct {
	Head *Node[T]
	Length int 
}

func MakeSinglyLinkedList[T any](value T) *SinglyLinkedList[T] {
	sll := SinglyLinkedList[T]{}
	sll.Head = &Node[T]{value, nil}
	sll.Length = 1
	return &sll
}

func (s *SinglyLinkedList[T]) InsertAtBeginning(value T) {
	ptr := s.Head
	if(ptr == nil){
		*s = *MakeSinglyLinkedList(value)
		return
	}
	newNode := Node[T]{value, s.Head}
	s.Head = &newNode
	s.Length++
}

func (s *SinglyLinkedList[T]) InsertAtEnd(value T) {
	ptr := s.Head
	if(ptr == nil){
		*s = *MakeSinglyLinkedList(value)
		return
	}
	for {
		if(ptr.Next == nil){
			ptr.Next = &Node[T]{value, nil}
			s.Length++
			return
		}
		ptr = ptr.Next
	}
}
