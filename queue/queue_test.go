package queue

import "testing"

func TestNew(t *testing.T) {
  q := New().Push(1)

  if q.head.Value != 1 {
    t.Errorf("q.head.Value == %v, want 1", q.head.Value)
  }

  if q.tail.Value != nil {
    t.Errorf("q.tail.Value == %v, want nil", q.tail.Value)
  }

  v := q.Pop()
  if v != 1 {
    t.Errorf("q.Pop() == %v, want 1", v)
  }
  // if New() != true {
  //   t.Errorf("New() != true")
  // }
}
