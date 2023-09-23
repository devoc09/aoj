package main

import (
	"fmt"
)

func reverse(str string, is, ie int) string {
    runes := []rune(str)
    for i, j := is, ie; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func replace(str string, is, ie int, rep string) string {
    return str[:is] + rep + str[ie+1:]
}

func main() {
    var str string
    var q int
    var directive string
    var is, ie int // index start, index end

    fmt.Scan(&str)
    fmt.Scan(&q)

    for i := 0; i < q; i++ {
        fmt.Scan(&directive)
        fmt.Scan(&is)
        fmt.Scan(&ie)

        switch directive {
        case "print":
            fmt.Printf("%s\n", str[is:ie+1])
        case "reverse":
            str = reverse(str, is, ie)
        case "replace":
            var tmp string
            fmt.Scan(&tmp)
            str = replace(str, is, ie, tmp)
        }
    }
}
