# Go dash [![GoDoc][doc-img]][doc] [![Coverage Status][cov-img]][cov] [![GoReport][go-report-img]][go-report]

``go-dash`` is a modern Go library based on [reflect](https://golang.org/pkg/reflect/).

Generic helpers rely on [reflect](https://golang.org/pkg/reflect/), be careful, this code runs exclusively at runtime, so you must have a good test suite.

``go-dash`` may look like [lodash](https://lodash.com/) in JavaScript because I write the go-dash base on lodash, but it will have a version upgrade and have its own roadmap.
We always write the code for the go-dash with the coverage of unit tests equal to 100 percent.

[Lodash](https://lodash.com/) is a popular and widely used JavaScript utility library that provides a comprehensive set of functions to work with arrays, objects, strings, functions, and more. It aims to simplify and streamline common programming tasks and data manipulation in JavaScript applications.

Why this name?
--------------

Because it is based on lodash and I write it with Golang, I decided to give it a name that is ``go-dash``. I hope everyone who can contribute to the ``go-dash`` will help make it more popular, flexible, and stable.

<div align="center">
  <img src="https://media.giphy.com/media/ydPWRTNMR3x8mAWTdk/giphy.gif">
</div>

Installation
------------
    go get github.com/warriors-vn/go-dash

Usage
-----

You can import ``go-dash`` using a basic statement:

```go
import "github.com/warriors-vn/go-dash"
```

These examples will be based on the following data model:

```go
type User struct {
    Name string
    Age  int
}
```

There are some examples about go-dash. You can find all function about: array, collection, date, math,lang, number, string in the library.

*Some examples:*

<h3>godash.filter</h3>
The function returns the new filtered slice and an error if any occurs.
<br>
```go
filter([]int{1, 2, 3, 4}, func(n int) bool {
    return n%2 == 0
}) // []int{2, 4}

filter([]string{"a", "b", "c"}, func(n string) bool {
return n > "a"
}) // []string{"b", "c"}

filter([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
    return n.Age%2 == 0
}) // []*User{{Name: "Kakalot", Age: 26}, {Name: "Trunk", Age: 10}}
```

<h3>godash.find</h3>
The function returns the found element and a boolean indicating whether a match was found, or an error if any occurs.
<br>
```go
find([]int{1, 2, 3, 4}, func(n int) bool {
    return n%2 == 0
}) // 2

find([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
    return n.Age%2 == 0
}) // &User{Name: "Kakalot", Age: 26}
```

<h3>godash.chunk</h3>
The function returns a slice of slices, where each inner slice contains elements from the original array.
<br>
```go
chunk([]int{1, 2, 3}, 1) // [][]int{{1}, {2}, {3}}
chunk([]string{"a", "b", "c", "d"}, 3) // [][]string{{"a", "b", "c"}, {"d"}}
chunk([]interface{}{1, "2", true}, 1) // [][]interface{}{{1}, {"2"}, {true}}
```

<h3>godash.intersection</h3>
The function accepts two arrays (slices) and returns a new array with the intersection of unique elements.
<br>
```go
intersection([]int{1, 2, 3}, []int{1}) // []int{1}
intersection([]string{"1.1", "2.2", "3.3"}, []string{"1.1", "3.3"}) // []string{"1.1", "3.3"}
```

<h3>godash.difference</h3>
It returns the index of the first occurrence of the value in the array. If the value is not found, the function returns -1.
<br>
```go
difference([]int{2, 1}, []int{2, 3}) // []int{1}
difference([]string{"2", "1", "2", "2", "5"}, []string{"2", "3"}) // []string{"1", "5"}
```

<h3>godash.indexOf</h3>
The function returns the new slice of different elements and an error if any occurs.
<br>
```go
indexOf([]string{"1", "2", "b", "3", "c", "b"}, "b", 2) // 2
indexOf([]interface{}{1.1, "2.2", 3.3}, 2.2, 0) // -1
indexOf([]float32{1.1, 2.2, 3.3}, float32(2.2), -6) // 1
```

<h3>godash.every</h3>
The function returns true if all elements satisfy the condition, false otherwise, or an error if any occurs.
<br>
```go
every([]int{1, 2, 3, 4}, func(n int) bool {
    return n%2 == 0
}) // false

every([]string{"d", "b", "c"}, func(n string) bool {
    return n > "a"
}) // true

every([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
    return n.Age%2 == 0
}) // false
```

<h3>godash.partition</h3>
The function returns two new slices: one containing elements that satisfy the predicate, and the other containing elements that do not satisfy the predicate.
<br>
```go
partition([]int{1, 2, 3, 4}, func(n int) bool {
    return n%2 == 0
}) // [][]int{{2, 4}, {1, 3}}

partition([]string{"a", "b", "c"}, func(n string) bool {
    return n > "a"
}) // [][]string{{"b", "c"}, {"a"}}
```

<h3>godash.drop</h3>
The function returns a new array with the specified number of elements removed from the beginning.
<br>
```go
drop([]int{1, 2, 3}, 2) // []int{3}
drop([]string{"1", "2", "3"}, 2) // []string{"3"}
drop([]interface{}{1, "2", true}, 2) // []interface{}{true}
```

<h3>godash.initial</h3>
initial returns a new array (slice) containing all elements of the input array except the last one.
<br>
```go
initial([]int{1, 2, 3}) // []int{1, 2}
initial([]string{"a", "b", "c"}) // []string{"a", "b"}
initial([]interface{}{"true", 1, 1.1, true, int32(123), false, true}) // []interface{}{"true", 1, 1.1, true, int32(123), false}
```

<h3>godash.tail</h3>
The function returns the new slice and an error if any occurs.
<br>
```go
tail([]int{1, 2, 3, 3, 2, 1}) // []int{2, 3, 3, 2, 1}
tail([]string{"1.1", "2.2", "3.3", "3.3", "2.2", "1.1"}) // []string{"2.2", "3.3", "3.3", "2.2", "1.1"}
tail([]interface{}{false, "true", true, 1, 2.2}) // []interface{}{"true", true, 1, 2.2}
```

<h3>godash.reverse</h3>
The function returns the modified array and an error if any occurs.
<br>
```go
reverse([]int{1, 2, 3}) // []int{3, 2, 1}
reverse([]string{"1.1", "2.2", "3.3"}) // []string{"3.3", "2.2", "1.1"}
reverse([]interface{}{"true", true, 1, 2.2, false}) // []interface{}{false, 2.2, 1, true, "true"}
```

<h3>godash.findIndex</h3>
The function returns the index of the first occurrence of 'target' and a boolean indicating if the target was found.
<br>
```go
findIndex([]int64{1, 2, 3, 8, 7, 6, 5, 9}, int64(8)) // 3
findIndex([]string{"a", "b", "c"}, "b") // 1
findIndex([]*User{
    {
        Name: "warriors",
        Age:  20,
    },
    {
        Name: "god",
        Age:  10,
    },
    {
        Name: "vegeta",
        Age:  15,
    },
}, &User{
    Name: "vegeta",
    Age:  15,
}) // 2
```

<h3>godash.pullAt</h3>
pullAt removes elements from the input array at specified indexes and returns the modified array.
<br>
```go
pullAt([]float64{1.1, 2.2, 3.3}, []int{0, 1, 2}) // []float64{1.1, 2.2, 3.3}
pullAt([]bool{true, true, false, true, false}, []int{0, 1, 3, 4}) // []bool{true, true, true, false}
pullAt([]interface{}{"true", 1, false, true, 1.1, "false"}, []int{0, 1, 3, 4, 5}) // []interface{}{"true", 1, true, 1.1, "false"}
```

<h3>godash.without</h3>
The function returns the new slice without the specified values and an error if any occurs.
<br>
```go
without([]int{2, 1, 2, 3}, 1, 2) // []int{3}
without([]string{"2", "1", "2", "3"}, "1", 2) // []string{"2", "2", "3"}
without([]interface{}{"true", true, false, 1, 2.2, 3.3}, 2.2, false) // []interface{}{"true", true, 1, 3.3}
```

<h3>godash.join</h3>
join takes an array (slice) of elements and a separator string, and returns a single string by joining the elements with the separator.
<br>
```go
join([]int{1, 2, 3}, "-") // "1-2-3"
join([]bool{true, false}, "-") // "true-false"
join([]interface{}{"1", false, 1.1, 2.2, 2}, "-") // "1-false-1.1-2.2-2"
```

<h3>godash.xor</h3>
xor returns a new slice that contains the elements that appear in an odd number of input arrays.
<br>
```go
xor([]int{4, 4, 2, 1}, []int{2, 3, 4}) // []int{1, 3}
xor([]float64{2, 1}, 3, []float64{2, 3}, []float64{2}) // []float64{1, 3}
xor([]string{"2", "1"}, 3, []string{"2", "3"}, []string{"2"}) // []string{"1", "3"}

```

<h3>godash.camelCase</h3>
camelCase converts a string to camelCase format.
<br>
```go
camelCase("Foo Bar") // fooBar
camelCase("--foo-bar--") // fooBar
camelCase("@a98@#$,.23237@#$hello-world") // a9823237HelloWorld
```

<h3>godash.unescape</h3>
unescape a string that contains escape sequences and returns the unescaped string. If un-escaping fails, the original input string is returned.
<br>
```go
unescape("fred, barney, &amp; pebbles") // fred, barney, & pebbles
unescape("fred &lt; barney") // fred < barney
unescape("fred &gt; barney") // fred > barney
```

<h3>godash.lte</h3>
The function accepts two values of any comparable type (int, float, string, etc.) and returns true if the first value is less than or equal to the second value, false otherwise.
<br>
```go
lte(1, 2) // true
lte(int64(1), int64(2)) // true
lte(3.2, 2.2) // false
```

Performance
-----------

Don't hesitate to participate in the discussion to enhance the generic helpers implementations.

Contributing
------------

* Ping me on instagram [@tuanelnino9](https://www.instagram.com/tuanelnino9) or [Facebook](https://www.facebook.com/tuanelnino9) or Linkedin [Tuan Nguyen Van](https://www.linkedin.com/in/tuan-nguyen-van-555315156) DMs, mentions, whatever, [send email](mailto:nguyenvantuan2391996@gmail.com) :))
* Fork the [project](https://github.com/warriors-vn/go-dash)
* Fix [open issues](https://github.com/warriors-vn/go-dash/issues) or request new features

Don't hesitate :))

Authors
-------

* Tuan Nguyen Van

[cov-img]: https://codecov.io/gh/warriors-vn/go-dash/branch/master/graph/badge.svg?token=AFGX9MQS21
[cov]: https://app.codecov.io/gh/warriors-vn/go-dash
[doc]: https://pkg.go.dev/github.com/warriors-vn/go-dash
[doc-img]: https://pkg.go.dev/badge/github.com/warriors-vn/go-dash
[go-report]: https://goreportcard.com/report/github.com/warriors-vn/go-dash
[go-report-img]: https://goreportcard.com/badge/github.com/thoas/go-funk