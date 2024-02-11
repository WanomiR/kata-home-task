---
type:
  - Code
language:
  - GO
libraries: 
tags:
---
> [!important] 
> Started working on the test task Friday 9, 2024.
# Packages, variables, and functions
## Packages, imports, exports
By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.

Programs start running in package `main`. The program below is using the packages with import paths `"fmt"` and `"math/rand"`.

In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```
> [!note] 
> You could also have multiple imports, each on its own line, but the way above if preferable.
> ```go
>import "fmt"
>import "math" 
>```

## Functions
### Arguments
A function can take zero or more arguments. In this example, `add` takes two parameters of type `int`. Notice that the type comes _after_ the variable name. (For more about why types look the way they do, see the [article on Go's declaration syntax](https://go.dev/blog/gos-declaration-syntax).)
```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

When two or more consecutive named function parameters share a type, you can omit the type from all but the last. In this example, we shortened:

```go
x int, y int
```
to
```go
x, y int
```
### Multiple return
A function can return any number of results.
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```
### Named return value
Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // <-- return is naked
}

func main() {
	fmt.Println(split(17))
}
```
### Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.
```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```
### Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
### Exercise: Fibonacci closure
```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, c+b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
## Types
### Basic types
The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
The example shows variables of several types, and also that variable declarations may be “factored” into blocks, as with import statements.
```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```
### Zero values
Variables declared without an explicit initial value are given their _zero value_.

The zero value is:
- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

### Type conversions
The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or, put more simply:
```go
i := 42
f := float64(i)
u := uint(f)
```
Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the `float64` or `uint` conversions in the example and see what happens.
### Type inference
When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right-hand side.

When the right-hand side of the declaration is typed, the new variable is of that same type:
```go
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

Try changing the initial value of `v` in the example code and observe how its type is affected.
```go
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

```
## Variables
The `var` statement declares a list of variables; as in function argument lists, the type is last. A `var` statement can be at package or function level. We see both in this example.
```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```
A var declaration can include initializers, one per variable.
```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
> [!note]
> If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
### Short variable declarations
Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
### Constants
Constants are declared like variables, but with the `const` keyword. Constants can be character, string, boolean, or numeric values.
```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```
> [!important] 
Constants cannot be declared using the `:=` syntax.
### Numeric constants
Numeric constants are high-precision _values_. An untyped constant takes the type needed by its context.
```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```
# Flow control
## For
Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.
```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

> [!note] 
Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

### For is Go's "while"
The init and post statements are optional. At that point, you can drop the semicolons:
```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
### Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
```go
package main

func main() {
	for {
	}
}
```
### Range
The `range` form of the `for` loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

```
### Range continued
You can skip the index or value by assigning to `_`.
```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```
## If
Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.
```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```
### If with a short statement
Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the `if`.
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
### If and else
Variables declared inside an `if` short statement are also available inside any of the `else` blocks.
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
### Newton's sqrt method
```go
package main

import (
	"fmt"
	"math"
)

const n = 5

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		zn := z
		zn -= (z*z - x) / (z * 2)
		if zn == z {
			fmt.Println(i)
			return z
		}
		z = zn
	}
	return z
}

func main() {
	fmt.Println(Sqrt(n))
	fmt.Println(math.Sqrt(n))
}

```
The z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method). It works well for many functions but especially well for square root.
## Switch
A `switch` statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression.
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```
> [!note] 
Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

Switch cases evaluate cases from top to bottom, stopping when a case succeeds. For example,
```go
switch i {
case 0:
case f():
}
```
does not call `f` if `i==0`.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```
Slice exercise.
```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	table := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		table[i] = row

		for j := 0; j < dx; j++ {
			table[i][j] = uint8((i + j) / 2)
		}
	}
	return table
}

func main() {
	pic.Show(Pic)
}
```
### Switch with no condition
Switch without a condition is the same as `switch true`. This construct can be a clean way to write long if-then-else chains.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```
## Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
### Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog post](https://go.dev/blog/defer-panic-and-recover).
```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```
# Advanced types
## Pointers
Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
```go
var p *int
```

The `&` operator generates a pointer to its operand.
```go
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value.
```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as “dereferencing” or "indirecting".

Unlike C, Go has no pointer arithmetic.
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

## Struct
A `struct` is a collection of fields
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

Struct fields are accessed using a dot.
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```
### Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```
### Struct literals
A struct literal denotes a newly allocated struct value by listing the values of its fields. You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)

The special prefix `&` returns a pointer to the struct value.
```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```
## Arrays
The type `[n]T` is an array of `n` values of type `T`. The expression
```go
var a [10]int
```
declares a variable `a` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized.
```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```
## Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```go
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:

```go
a[1:4]
```

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

```
### Slices are references to arrays
A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.
```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```
### Slice literals
A slice literal is like an array literal without the length.

This is an array literal:
```go
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:
```go
[]bool{true, true, false}
```

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```
### Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

For the array
```go
var a [10]int
```
these slice expressions are equivalent:
```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

```
### Slice length and capacity
A slice has both a _length_ and a _capacity_.

The length of a slice is the number of elements it contains. The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`. You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

```
### Nil slices
The zero value of a slice is `nil`. A nil slice has a length and capacity of 0 and has no underlying array.
```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```
### Creating a slice with make
Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to `make`:
```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```
```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

```
### Slices of slices
Slices can contain any type, including other slices.
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```
### Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in `append` function. The [documentation](https://go.dev/pkg/builtin/#append) of the built-in package describes `append`.
```go
func append(s []T, vs ...T) []T
```

The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice. The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

> [!note] 
If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

To learn more about slices, read the [Slices: usage and internals](https://go.dev/blog/go-slices-usage-and-internals) article.)
```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
## Maps
A map maps keys to values.  The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```
### Map literals
Map literals are like struct literals, but the keys are required.
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```
### Map literals continued
If the top-level type is just a type name, you can omit it from the elements of the literal.
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```
### Mutating maps
Insert or update an element in map m:
```go
m[key] = elem
```
Retrieve an element:
```go
elem = m[key]
```
Delete an element:
```go
delete(m, key)
```
Test that a key is present with a two-value assignment:
```go
elem, ok = m[key]
```
If key is in m, ok is true. If not, ok is false.  If key is not in the map, then elem is the zero value for the map's element type.

> [!note]  
> If elem or ok have not yet been declared you could use a short declaration form:
> ```go
> elem, ok := m[key]
> ```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

```
### Exercies: Maps
Count up the number of occurences for each word in the dictionary.
```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, w := range strings.Fields(s) {
		res[w]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
```

# Methods and Interfaces
## Methods
Go does not have classes. However, you can define methods on types.  A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
### Methods are functions
Remember: a method is just a function with a receiver argument. Here's Abs written as a regular function with no change in functionality
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```
### Methods continued
You can declare a method on non-struct types, too. In this example we see a numeric type MyFloat with an Abs method.

> [!important]  
> You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int)
```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
## Pointer receivers
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument.) The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```