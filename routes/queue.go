package routes

import (
	"sync"

	"github.com/judewood/routeDistances/models"
)

type Queue struct {
	Items []models.RouteSection
	Lock  sync.RWMutex
}

// Enqueue adds an Item to the end of the queue
func (s *Queue) Enqueue(item models.RouteSection) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, item)
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if item.CumulativeDistance < v.CumulativeDistance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = item
				insertFlag = true
			} else {
				s.Items = append([]models.RouteSection{item}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, item)
	}
}

// Dequeue removes an Item from the start of the queue
func (s *Queue) Dequeue() *models.RouteSection {
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
	s.Items = []models.RouteSection{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *Queue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of items in the queue
func (s *Queue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}
