package main

import (
    "sync"
    "errors"
)

type Queue struct {
  sync.Mutex
  name string
  slots int
  active_jobs int
}

//type QueueMap Queue

func NewQueue(name string, slots int) *Queue {
  q := Queue{
    *new(sync.Mutex),
    name,
    slots,
    0,
  }
  return &q
}

func (q *Queue) SetName(name string) {
  q.Lock()
  q.name = name
  q.Unlock()
}

func (q *Queue) Name() string {
  q.Lock()
  name := q.name
  q.Unlock()
  return name
}

func (q *Queue) SetSlots(num_slots int) {
  q.Lock()
  q.slots = num_slots
  q.Unlock()
}

func (q *Queue) Slots() int {
  q.Lock()
  slots := q.slots
  q.Unlock()
  return slots
}

// Jobs that wish to run in this queue must call Join and have it return true
// signifying that they have successfully joined the queue. If the queue is full,
// Join will return false, and it is up to the job to retry at some later point.
func (q *Queue) Join() bool {
  var joined bool
  q.Lock()
  joined = (q.slots - q.active_jobs) > 0
  if joined {
    q.active_jobs++
  }
  q.Unlock()
  return joined
}

// What a job is done executing, it must relenquish its spot in the queue by
// calling Leave
func (q *Queue) Leave() error {
  q.Lock()
  if q.active_jobs < 1 {
    return errors.New("The queue is already empty")
  } else {
    q.active_jobs--
  }
  q.Unlock()
  return nil
}
