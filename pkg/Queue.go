package Sitemap

type Queue struct {
	elements []string
}

func (q *Queue) enqueue(item string) {
	q.elements = append(q.elements, item)
}

func (q *Queue) dequeue() string {
	if len(q.elements) == 0 {
		return ""
	}
	item := q.elements[0]
	q.elements = q.elements[1:]
	return item
}

func (q *Queue) size() int {
	return len(q.elements)
}
