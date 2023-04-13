## Interface

An interface is a **named collection of method signatures** that a type can implement.

This definition is very abstract, so let us work out a canonical example: shapes. A triangle, a square, a rectangle, and a circle are shapes. They are geometrically different, but they share similarities. They all have an area and a perimeter, although each is calculated differently, based on the shape's characteristics.


| Shape     | Area      | Perimeter  | Terms
| --------- | --------- | ---------- | ----------------
| Triangle  | $½ b h$   |            | b=base; h=height
| Square    | $a^2$     |            | a=side
| Rectangle | $l w$     |            | l=length; w=width
| Circle    | $\pi r^2$ |            | r=radius

If you are familiar with other object-oriented languages, notice that interfaces in Go do not enforce a type to implement methods.

An interface `Shape` would allow us to treat triangles, squares, rectangles and circles uniformly, and to call their functions to compute area and perimeter without knowing the actual shape.

Let us declare structs for the different shapes. This is nothing different than what we saw earlier, when we used `Person` as an example. To limit the lines of codes, we shall limit ourselves to `Square` and `Rectangle`.

```go
type Square struct {
    Side float64
}

type Rectangle struct {
    Width float64
    Height float64
}
```

We now declare an interface `Shape`. The keyword `type` is used. Notice the similarities and the differences between an `interface` and a `struct`.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

As we know, functions can be added to structs. Guess what. We are going to add `Area()` and `Perimeter()` functions to our two shapes. The actual computation depend on the shape, but the function declarations are the same, only applied to different types.

```go
func (s *Square) Area() float64 {
    return s.Side * s.Side
}

func (s *Square) Perimeter() float64 {
    return 4 * s.Side
}

func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

Do not mind the `*`. They denote _pointers_, e.g. we are manipulating _something that points to a square_ rather than a square. We will come back to pointers later.

For the sake of readability, let us add the much needed `String` functions, otherwise our shapes will be represented as numbers, e.g. `&{5}` or `&{4. 8}`

```go
func (s *Square) String() string {
    return fmt.Sprintf("Square. s=%f", s.Side)
}

func (r *Rectangle) String() string {
    return fmt.Sprintf("Rectangle. w=%f, h=%f", r.Width, r.Height)
}
```

Now that `Square` and `Rectangle` implement the methods defined in `Shape`, they _are_ shapes and can be treated uniformly as such ; for instance in a slice of `[]Shape`. Notice how, after the slice declaration, there are no references to `Square` or `Triangle`. There is no use of the original structs, only of the interface `Shape`. We have a achieved a form of **polymorphism**.

```go
square := &Square{Side: 5}
rectangle := &Rectangle{Width: 4, Height: 8}
shapes := []Shape{square, rectangle}
for _, shape := range shapes {
    fmt.Printf("%s => area=%f, perimeter=%f\n",
        shape, shape.Area(), shape.Perimeter())
}
```

```
Square. s=5.000000 => area=25.000000, perimeter=20.000000
Rectangle. w=4.000000, h=8.000000 => area=32.000000, perimeter=24.000000
```

> _Note_: Make sure you implement the **exact methods** from the interface. A simple typo and you would just be declaring an unrelated function, failing to adhere to the interface.


### The empty interface

The interface type that specifies zero methods is known as the empty interface.

```go
interface{}
```

An empty interface may hold values of any type, because every type implements "at least" zero methods. In other words, all types implement the empty interface _implicitly_.

Consider the following declaration[^tour-14]:

[^tour-14]: Example comes from [A Tour of Go](https://tour.golang.org/methods/14) https://tour.golang.org/methods/14


```go
var i interface{}
fmt.Printf("(%v, %T)\n", i, i)
```

The output is `(<nil>, <nil>)`, meaning both value and type are `nil`. Let us assign something to `i`.

```go
i = 42
fmt.Printf("(%v, %T)\n", i, i)
```

`(42, int)` will be printed, meaning `i` is an `int` and its value is `42`. One last for the road:

```go
i = "hello"
fmt.Printf("(%v, %T)\n", i, i)
```

Prints `(hello, string)` and, as you figured out, `i` is now a `string` and its value is `"hello"`.

> _Note_: These code snippets are for demonstration purpose only. Their appearance in actual code would be dubious.

Empty interfaces are used by code that handles values of **unknown type**. For instance, the well-known `fmt.Println`[^println] takes any number of arguments of type `interface{}`. All that time, you have been using it intuitively. Now you can understand the mechanism that lies behind.

[^println]: https://golang.org/pkg/fmt/#Println

### Exercices

1. Complete the `Shape` example by adding `Triangle`. Implement the requires functions so that it becomes a shape. Make sure the output is readable when printed.

2. Declare an array containing the string `"hello"` and the integer `42`.
