package main

//exercises from https://gobyexample.com

import (
	"errors"
	"fmt"
	"math"
	"time"
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

	for j := 7; j <= 9; j++ {
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
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
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

func _arrays() {
	fmt.Println("-- Array")
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{111, 222, 333, 444, 555}
	fmt.Println("dcl:", b)
	fmt.Println("dcl:", b[2])

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
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
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "aaa"
	s[1] = "bbb"
	s[2] = "ccc"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)

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
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func _range() {
	fmt.Println("-- Range")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "abcefghijklmnopqrstuvxz" {
		fmt.Println(i, c)
	}

}

func plus(a int, b int) int {
	return a + b
}

func _functions() {
	fmt.Println(plus(50, 70))
}

func vals() (int, int) {
	return 3, 7
}

func _multi_return_values() {
	fmt.Println("-- Multiple return values")
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func _variadic() {
	fmt.Println("-- Variadic")
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func _closures() {
	fmt.Println("-- Closures")
	nextInt := intSeq()
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3
	newInts := intSeq()
	fmt.Println(newInts()) //1
}

func fact(n int) int {
	fmt.Println("-- Recursion")
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func _recursion() {
	fmt.Println(fact(7))
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func _pointers() {
	fmt.Println("-- Pointers")
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

type person struct {
	name string
	age  int
}

func _structs() {
	fmt.Println(person{"BoB", 20})
	fmt.Println(person{age: 23, name: "Alice"})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 50})

	s := person{name: "Sean", age: 22}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 233
	fmt.Println(sp.age)
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func _methods() {
	fmt.Println("-- Methods")
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func _interfaces() {
	fmt.Println("-- Interfaces")
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	measure(s)
	measure(c)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with 42...."}
	}
	return arg + 3, nil
}

func _errors() {
	fmt.Println("-- Errors")
	for _, i := range []int{7, 42} {
		fmt.Println("--->", i)
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		fmt.Println("--->", i)
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func _goroutines() {
	fmt.Println("-- Go Routines")
	f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("Type in something: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done: ", input)
}

func _channels() {
	fmt.Println("-- Channels")
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}

func _channel_buffering() {
	fmt.Println("Channel Buffering")
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("printed...")
	messages <- "channel3"
	fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func _channel_synchronization() {
	fmt.Println("-- Channel Synchronization")
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func _channel_directions() {
	fmt.Println("-- Channel Directions")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func _select() {
	fmt.Println("-- Select")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println(" i = ", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func _timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

func main() {
	values()
	variables()
	constants()
	loops()
	if_else()
	switch_example()
	_arrays()
	_slices()
	_maps()
	_range()
	_functions()
	_multi_return_values()
	_variadic()
	_closures()
	_recursion()
	_pointers()
	_structs()
	_methods()
	_interfaces()
	_errors()
	_goroutines()
	_channels()
	_channel_buffering()
	_channel_synchronization()
	_channel_directions()
	_select()
	_timeouts()
}
