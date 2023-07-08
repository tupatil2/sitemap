package Sitemap

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(item string) {
	q.elements = append(q.elements, item)
}

func (q *Queue) Dequeue() string {
	if len(q.elements) == 0 {
		return ""
	}
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

func (q *Queue) Peek() string {
	if len(q.elements) == 0 {
		return ""
	}
	return q.elements[0]
}

func (q *Queue) Size() int {
	return len(q.elements)
}
