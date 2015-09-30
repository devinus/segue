package queue

type Message struct {
  next, prev *Message
  Value interface{}
}

type Queue struct {
  head Message
  tail Message
  size int
}

func (q *Queue) Init() *Queue {
  q.head = q.tail
  q.tail = q.head
  q.head.next = &q.tail
  q.tail.prev = &q.head
  q.size = 0
  return q
}

func New() *Queue {
  return new(Queue)
}

func (q *Queue) Push(v interface{}) *Queue {
  m := &Message{Value: v}
  q.head.prev = m
  m.next = &q.head
  q.head = *m
  return q
}

func (q *Queue) Pop() interface{} {
  m := q.tail

  if m.prev != nil {
    q.tail = *m.prev
    q.tail.next = nil
  }

  return m.Value
}
