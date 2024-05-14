package routes

import "sync"

type Queue struct {
	Items []Edge
	Lock  sync.RWMutex
}

// Enqueue adds an Item to the end of the queue
func (s *Queue) Enqueue(t Edge) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Edge{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
}

// Dequeue removes an Item from the start of the queue
func (s *Queue) Dequeue() *Edge {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

// NewQ Creates New Queue
func (s *Queue) NewQ() *Queue {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Items = []Edge{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *Queue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *Queue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}
