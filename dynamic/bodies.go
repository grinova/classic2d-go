package dynamic

import "github.com/grinova/classic2d-server/physics"

// BodyList - список тел
type BodyList struct {
	Body *physics.Body
	Next *BodyList
	Prev *BodyList
}

func (b *BodyList) addNext(body *physics.Body) *BodyList {
	list := &BodyList{Body: body, Prev: b, Next: b.Next}
	b.Next = list
	return list
}

func (b *BodyList) addPrev(body *physics.Body) *BodyList {
	list := &BodyList{Body: body, Prev: b.Prev, Next: b}
	b.Prev = list
	return list
}

func (b *BodyList) remove() {
	if b.Prev != nil {
		b.Prev.Next = b.Next
	}
	if b.Next != nil {
		b.Next.Prev = b.Prev
	}
	b.Prev = nil
	b.Next = nil
	b.Body = nil
}

type bodies struct {
	m     map[*physics.Body]*BodyList
	first *BodyList
	last  *BodyList
}

func (b *bodies) add(body *physics.Body) {
	if b.first == nil && b.last == nil {
		if b.m == nil {
			b.m = make(map[*physics.Body]*BodyList)
		}
		list := &BodyList{Body: body}
		b.first = list
		b.last = list
	} else {
		b.last = b.last.addNext(body)
	}
	b.m[body] = b.last
}

func (b *bodies) remove(body *physics.Body) {
	if list, ok := b.m[body]; ok {
		if b.first == list {
			b.first = list.Next
		}
		if b.last == list {
			b.last = list.Prev
		}
		list.remove()
		delete(b.m, body)
	}
}
