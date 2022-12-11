package sllist

import "testing"

func TestSinglyLinkedList(t *testing.T) {
	list := new(SinglyLinkedList[string])

	t.Run("Insert items at the beginning of the list", func(t *testing.T) {
		list.InsertAtBeginning("Test")
		list.InsertAtBeginning("a")
		list.InsertAtBeginning("am")
		list.InsertAtBeginning("I")
		ptr := list.Head
		var got [4]string
		index := 0
		for ptr != nil {
			got[index] = ptr.Data
			ptr = ptr.Next
			index++
		}
		want := [4]string{"I", "am", "a", "Test"}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Insert Items at the end of the list", func(t *testing.T) {
		list.InsertAtEnd(",")
		list.InsertAtEnd("write")
		list.InsertAtEnd("by")
		list.InsertAtEnd("dev4ndy")
		ptr := list.Head
		var got [8]string
		index := 0
		for ptr != nil {
			got[index] = ptr.Data
			ptr = ptr.Next
			index++
		}
		want := [8]string{"I", "am", "a", "Test", ",", "write", "by", "dev4ndy"}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Delete an item given a value", func(t *testing.T) {
		list.Delete("dev4ndy")
		ptr := list.Head
		var got [7]string
		index := 0
		for ptr != nil {
			got[index] = ptr.Data
			ptr = ptr.Next
			index++
		}
		want := [7]string{"I", "am", "a", "Test", ",", "write", "by"}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Search an item given a value", func(t *testing.T) {
		index, node := list.Search("Test")
		got := node.Data
		want := "Test"

		if got != want || index != 3 {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
