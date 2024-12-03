package main


// List represents a singly-linked list that holds
// values of type int.
type List[T any] struct {
    next *List[T]
    val T
}

// NewList creates a new List with the given value.
func NewList[T any](val T) *List[T] {
    return &List[T]{nil, val}
}

// Append adds a new value to the end of the list.
func (l *List[T]) Append(val T) {
    for l.next != nil {
        l = l.next
    }
    l.next = NewList(val)
}

// Prepend adds a new value to the beginning of the list.
func (l *List[T]) Prepend(val T) *List[T] {
    return &List[T]{l, val}
}

// Reverse reverses the list in place.
func (l *List[T]) Reverse() {
    var prev *List[T]
    for l != nil {
        next := l.next
        l.next = prev
        prev = l
        l = next
    }
}

// Values returns a slice of all the values in the list.
func (l *List[T]) Values() []T {
    vals := []T{}
    for l != nil {
        vals = append(vals, l.val)
        l = l.next
    }
    return vals
}

// main is the entry point for the application.
func main() {
    l := NewList(1)
    l.Append(2)
    l.Append(3)
    l = l.Prepend(0)
    l.Reverse()
    for _, val := range l.Values() {
        println(val)
    }
}
