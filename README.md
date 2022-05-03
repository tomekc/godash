# GoDash

Finally! Introduction of Generics in Golang allows implementation of reusable algorithms or functional programming primitives wiht much less hassle.

The project is work in progress: intended to learn and be of practical use.

Please consider performance aspects of Go generics when implementing time-critical systems.

**Godash** is like [Lodash](https://www.lodash.com/)'s  younger cousin.

**Makes use of generics in Go, hence requires Go 1.18.**


## Usage

```go
import "github.com/tomekc/godash"
```

## Functional primitives

### Map

Returns slice of elements transformed by mapper function.

```go
var mapper = func(s string) int {
	return len(s)
}

result := Map([]string{"alice", "bob", "charlie"}, mapper)

// {5, 3, 7}
```

### Filter

Returns array of items matching given predicate function.

```go
func IsOddInteger(i int) bool {
	return i%2 == 1
}

result := godash.Filter([]int{1, 2, 3}, IsOddInteger)
// {1, 3}

```

### ForEach

Execute function on each element of slice.

```go
var urls []string{}

// ...
godash.ForEach(urls, func(url) {
	fetch(url)
})
```
