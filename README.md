# Golang Streams

* Inspired by Java8 streams
* Has `Filter`, `Map`, `Collect` and `First`
    * `Filter` filters all values
    * `Map` maps values into another type
    * `Collect` or `First` collects the data and returns
* Iterators:
    * Array iterator
    * Channel iterator
    * Generator
* Benefits
    * No need to write a lot of `for`'s
    * All logic is `stream.Filter(*).Map(*).Filter(*).Collect(*)`

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

## We can go compact

```go
// Let's assume we have: someArray[10, 20, 30, 40, 50, 60, 70]

arr := StreamFromArray(someArray)
    .Fiter(func (data interface{}) bool {
        return data.(int) < 30
    })
    .Map(func (data interface{}) interface{} {
        return data.(int) / 10
    })
    .Collect(NewArrayCollector()).([]interface{})
```
* Or even better
```go
arr := StreamFromArray(someArray)
    .Filter(numbersFilter)
    .Map(numbersMapper)
    .Collect(NewArrayCollector()).([]interface{})
```

# Generators
* Instead of iterate over array of values, it can iterate over the GENERATED data on the fly
```go
// Create generator which will generate: 0, 1000, 2000, 3000, 4000
gen := NewGenerator(5, func(i int) interface{} {
    return i * 1000
})

// It can be used in streams
oddNumbers := NewStream(gen)
    .Filter(onlyOddNumbers)
    .Map(toString)
    .Collect(NewArrayCollector()).([]interface{})
```

# Listen to a channel `chan interface{}`
```go
// Let's assume that we have an requestsChannel
// It will have requests data with user profiles

users := StreamFromChannel(requestsChannel)
    .Filter(onlyApiURLs)
    .Map(GetRequestBody)
    .Map(JsonToUserProfile)
    .Collect(NewMapCollector(ByUserId)).(map[interface{}]interface{})
```

# ForEach operation
```go
// Let's assume we have: someArray[1, 2, 3, 4, 5]

StreamFromArray(someArray)
    .Filter(isOdd)
    .ForEach(func (data interface{}) {
        println(data) // Will print only odd numbers
    })
```
```go
// Let's assume that we have an requestsChannel
// It will have requests data with user profiles

StreamFromChannel(requestsChannel)
    .Filter(onlyApiURLs)
    .Map(GetRequestBody)
    .Map(JsonToUserProfile)
    .Filter(OnlyBannedUsers)
    .ForEach(LogToJsonFile)
```