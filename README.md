# Golang Streams

* Inspired by Java8 streams
* Has `Filter`, `Map`, `Collect` and `First`
    * `Filter` filters all values
    * `Map` maps values into another type
    * `Collect` or `First` collects the data and returns

# Stream
```go
// Create stream from array.
// For example: someArray[10, 20, 30, 40, 50, 60, 70]
s := StreamFromArray(someArray)

// We don't need values greater than 30
s.Fiter(func (data interface{}) bool {
    return data.(int) < 30
})

// Convert all 10, 20, 30 into 1, 2, 3
s.Map(func (data interface{}) interface{} {
    return data.(int) / 10
})

// Now let's get an array: [1, 2, 3]
arr := s.Collect(NewArrayCollector()).([]interface{})

// ... Or let's get only fist value
// Will return nil when no values found
first := s.FirstOr(nil)

// We can also collect into a map
// keyGetter - is a function which generates a key for map record based on each value in a stream
// Sample:
//     func(data) { return data.(*User).Id }
// That function will get User.Id as an key for a map
MAP := s.Collect(NewMapCollector(keyGetter)).(map[interface{}]interface{})
```