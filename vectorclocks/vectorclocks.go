package vectorclocks

import "sync"

var Timestamps = make(map[string][]int)
var num_processes = 0
var mutex = sync.Mutex{}

type process struct {
	i              int
	curr_timestamp []int
	mux            *sync.Mutex
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MakeProcesses(n int) []process {
	num_processes = n
	p := make([]process, n)
	for i := 0; i < n; i++ {
		p[i] = process{
			i:              i,
			curr_timestamp: make([]int, n),
			mux:            &sync.Mutex{},
		}
	}
	return p
}

func IncrEvent(p process, event string) {
	p.mux.Lock()
	p.curr_timestamp[p.i] = p.curr_timestamp[p.i] + 1
	p.mux.Unlock()
	update_timestamp(event, p.curr_timestamp)
}

func SendMessage(p1 process, p2 process, event1 string, event2 string) {
	IncrEvent(p1, event1)
	receive_message(p2, event2, p1.curr_timestamp)
}

func update_timestamp(event string, timestamp []int) {
	dst := make([]int, len(timestamp))
	copy(dst, timestamp)
	mutex.Lock()
	Timestamps[event] = dst
	mutex.Unlock()
}

func receive_message(p process, event string, timestamp []int) {
	IncrEvent(p, event)
	p.mux.Lock()
	for i := 0; i < num_processes; i++ {
		p.curr_timestamp[i] = max(p.curr_timestamp[i], timestamp[i])
	}
	p.mux.Unlock()
	update_timestamp(event, p.curr_timestamp)
}

func IsBefore(event1 string, event2 string) bool {
	mutex.Lock()
	t1 := Timestamps[event1]
	t2 := Timestamps[event2]
	mutex.Unlock()
	for i := 0; i < num_processes; i++ {
		if t1[i] > t2[i] {
			return false
		}
	}
	for i := 0; i < num_processes; i++ {
		if t1[i] < t2[i] {
			return true
		}
	}
	return false
}

func IsConcurrent(event1 string, event2 string) bool {
	return !less_than_equal(event1, event2) && !less_than_equal(event2, event1)
}

func is_equal(event1 string, event2 string) bool {
	mutex.Lock()
	t1 := Timestamps[event1]
	t2 := Timestamps[event2]
	mutex.Unlock()
	for i := 0; i < num_processes; i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
func less_than_equal(event1 string, event2 string) bool {
	mutex.Lock()
	t1 := Timestamps[event1]
	t2 := Timestamps[event2]
	mutex.Unlock()
	for i := 0; i < num_processes; i++ {
		if t1[i] > t2[i] {
			return false
		}
	}
	return true
}
