Levenshtein Distance
====================

[Go](http://golang.org) package to calculate the [Levenshtein Distance](http://en.wikipedia.org/wiki/Levenshtein_distance)

Install
-------

    go get github.com/xqdoo00o/levenshtein

Example
-------

```go
package main

import (
        "fmt"
        "github.com/xqdoo00o/levenshtein"
)

func main() {
        str1 := "kitten"
        str2 := "sitting"
        dis, sim := levenshtein.Calc(str1, str2)
        fmt.Printf("The distance between %v and %v is %v\n",
                str1, str2, dis)
        fmt.Printf("The similarity between %v and %v is %v\n",
                str1, str2, sim)
}

```

Documentation
-------------
  for JavaScript version see [js-levenshtein](https://github.com/gustf/js-levenshtein)
