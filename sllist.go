package sllist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	Head *Node[T]
	Length int 
}

func MakeSinglyLinkedList[T comparable](value T) *SinglyLinkedList[T] {
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

func (s *SinglyLinkedList[T]) Delete (value T) {
	ptr := s.Head
	if ptr.Data == value {
		s.Head = ptr.Next
		s.Length--
		return
	}
	tmp := ptr
	ptr = ptr.Next
	for ptr != nil {
		if ptr.Data == value {
			tmp.Next = ptr.Next
			s.Length--
			return
		}
		tmp = ptr
		ptr = ptr.Next
	}
}

func (s *SinglyLinkedList[T]) Search (value T) (int, *Node[T]) {
	ptr := s.Head
	index := 0
	for ptr != nil {
		if ptr.Data == value {
			return index, ptr
		}
		ptr = ptr.Next
		index++
	}
	return -1, nil
}

func (s *SinglyLinkedList[T]) InsertAfterTarget(value T, targetValue T) error {
	ptr := s.Head
	for ptr != nil {
		if ptr.Data == targetValue {
			newNode := Node[T]{value, ptr.Next}
			ptr.Next = &newNode
			s.Length++;
			return nil
		}
		ptr = ptr.Next
	}
	return errors.New(fmt.Sprintf("The \"%v\" value was not found", targetValue))
}

func (s *SinglyLinkedList[T]) Traverse() {
	ptr := s.Head
	for ptr != nil {
		fmt.Printf("%v\n", ptr.Data)
		ptr = ptr.Next
	}
}

func (s *SinglyLinkedList[T]) Map(function func(value T) T) *SinglyLinkedList[T] {
	sll := new(SinglyLinkedList[T])
	ptr := s.Head
	for ptr != nil {
		newValue := function(ptr.Data)
		sll.InsertAtEnd(newValue)
		ptr = ptr.Next
	}
	return sll
}
