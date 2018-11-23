# Vector Clocks
`MakeProcesses(n int) []process` returns an array of processes.  
`IncrEvent(p process, event string)` creates an event in process p, with an id of `event`. The timestamp for this event is saved.  
`SendMessage(p1 process, p2 process, event1 string, event2 string)` sends a message from `p1` to `p2`. The event id in `p1` is `event1` and for `p2` is `event2`.  

For each event, the timestamp is kept in a map. `IsConcurrent(event1 string, event2 string)` and `IsBefore(event1 string, event2 string)` can be used to compare past events.
## Testing
To run the tests `go test` in the test directory
The vector clock test is based on this example:
[![VECTOR CLOCK](https://i.imgur.com/HEqYy8X.png)](http://www.youtube.com/watch?v=jD4ECsieFbE)

Each timestamp is tested, as well as `IsBefore` and `IsConcurrent` relationships