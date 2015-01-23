package main
import (
    "fmt"
    "time"
    "math"
)

func values() {
    fmt.Println("-- Values")
    fmt.Println("Hello " + "go" + "lang")
    fmt.Println("1+1 = ", 1+1)
    fmt.Println("7.0/3.0 = ", 7.0/3.0)
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(false || true)
    fmt.Println(!false)
    fmt.Println("Timenow is: ", time.Now())
}

func variables() {
    fmt.Println("-- Variables")
    var a string = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short"
    fmt.Println(f)

    g := 10
    fmt.Println(g)
}

func constants() {
    fmt.Println("-- Constants")
    const s string = "constant"
    fmt.Println(s)
    const n = 50000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}

func loops() {
    fmt.Println("-- Loops")
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <=9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }
}

func if_else() {
    fmt.Println("-- If/Else")

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {
        fmt.Println( num, "is negative" )
    } else if num < 10 {
        fmt.Println( num, "has 1 digit" )
    } else {
        fmt.Println( num, "has multiple digits" )
    }
}

func switch_example() {
    fmt.Println("-- Switch")
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("Three")
    }

    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("its the weekend")
        default:
            fmt.Println("its a weekday")
    }

    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("its before noon")
        default:
            fmt.Println("its after noon")
    }
}

func arrays() {
    fmt.Println("-- Array")
    var a [5]int
    fmt.Println("emp: ", a)

    a[4] = 100
    fmt.Println("set:",a)
    fmt.Println("get:",a[4])

    fmt.Println("len:", len(a))

    b := [5]int{111,222,333,444,555}
    fmt.Println("dcl:",b)
    fmt.Println("dcl:",b[2])

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3 ; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d", twoD)
}

func _slices() {
    fmt.Println("-- Slices")
    /*
    -- Slices
    emp: [  ]
    set: [aaa bbb ccc]
    get: ccc
    len: 3
    apd: [aaa bbb ccc d e f]
    cpy: [aaa bbb ccc d e f]
    sl1: [ccc d e]
    sl2 [aaa bbb ccc d e]
    sl3 [ccc d e f]
    dcl: [g h i]
    2d:  [[0] [1 2] [2 3 4]]
    */
    s := make([]string,3)
    fmt.Println("emp:", s)
    s[0] = "aaa"
    s[1] = "bbb"
    s[2] = "ccc"
    fmt.Println("set:",s)
    fmt.Println("get:",s[2])
    fmt.Println("len:",len(s))
    s = append(s,"d")
    s = append(s, "e", "f")
    fmt.Println("apd:",s)

    c := make([]string,len(s))
    copy(c,s)
    fmt.Println("cpy:",c)

    l := s[2:5]
    fmt.Println("sl1:",l)

    l = s[:5]
    fmt.Println("sl2",l)

    l = s[2:]
    fmt.Println("sl3", l)

    t := []string{"g","h","i"}
    fmt.Println("dcl:",t)

    twoD := make([][]int,3)

    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

func _maps() {
    fmt.Println("-- Maps")
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:",v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:",m)

    _,prs := m["k2"]
    fmt.Println("prs:",prs)


    n := map[string]int{"foo":1,"bar":2}
    fmt.Println("map:", n)
}

func main() {
    values()
    variables()
    constants()
    loops()
    if_else()
    switch_example()
    arrays()
    _slices()
    _maps()
}


