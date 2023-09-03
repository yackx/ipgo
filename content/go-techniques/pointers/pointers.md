## Pointers

A variable holds a value. A **pointer** holds the memory address of a value.

### Declaration

```go
// Declare an integer variable with value 10
val := 10
    
// Declare a pointer to the integer variable
ptr := &val
    
// Print the value of the integer variable by dereferencing pointer
fmt.Println(*ptr) // Output: 10
```

The following syntax is more compact. It also shows `ptr` is not a mere `int` and that it requires the use of `new`.

```go
ptr := new(int)
*ptr = 10
```

\begin{center}
\begin{figure}
\centering
\begin{tikzpicture}[baseline=(current bounding box.center)]
  % Draw the memory box for the variable
  \draw (0,0) rectangle (2,1);
  % Draw the value of the variable
  \node at (0.5,0.5) {10};
  % Draw the label for the variable
  \node at (3,0.5) {val};
\end{tikzpicture}
\hspace{2cm}
\begin{tikzpicture}[baseline=(current bounding box.center)]
  % Draw the memory box for the pointer
  \draw (0,0) rectangle (2,1);
  % Draw the value of the pointer
  \node at (0.5,0.5) {10};
  % Draw the pointer arrow
  \draw[<-] (1.5,0.5) -- (2.5,0.5);
  % Draw the label for the pointer
  \node at (3,0.5) {ptr};
\end{tikzpicture}
\caption{Illustration of a variable (left) and a pointer (right)}
\end{figure}
\end{center}

A pointer is represented by an asterisk `*` followed by the variable name. `*ptr` is a _pointer to an int_.

The asterisk is also used to **dereference** the pointer. Dereferencing a pointer gives you access to the value the pointer points to.

### Usage

Pointers in Go are useful for:

- **Modifying values in-place**. When you pass a value to a function, it is passed by value, which means that any changes made to the parameter in the function will not affect the original value. However, if you pass a _pointer_ to the value, you can modify the value in-place, which means that the changes made in the function will also be reflected in the original value.

- **Data structures**. Pointers are essential for implementing many data structures such as linked lists, trees, and graphs. These data structures require dynamic allocation of memory and pointers allow you to allocate and manage memory dynamically.

- **Performance**. When a large data structure such as a struct is passed to a function, it is copied to a new memory location. This copy can be expensive in terms of time and memory. By passing a pointer to the original data structure, you can avoid copying the entire data structure.

### Modify in-place

Consider the following function:

```go
func AddOne(x int) {
  x = x + 1
}

func main() {
  x := 10
  AddOne(x)
  fmt.Println(x) // x is still 10
}
```

The function `AddOne` will **not** modify `x` because it is passed **by value**. Of course, we don't need no fancy pointer to add 1. We don't even need a function. That is just an example. Let's modify our function to actually add 1 to `x`.

```go
func AddOne(x *int) {
  *x = *x + 1
}

func main() {
  x := 10
  AddOne(&x)
  fmt.Println(x) // x is 11
}
```

By using a pointer `(*int)`, the `AddOne` function is able to modify the original `x`.

The `&` operator is used to obtain the **address** of a variable. `&x` returns a `*int` (pointer to an int) because `x` is an `int`. This is what allows us to modify the original variable.

### Data structures

We have already seen an example of pointers with the `struct Person`. Quick refresher:

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}

robert := NewPerson("Robert", 67)
// robert.Name == "Robert"
// robert.Age == 67

func (p *Person) IsJuniorDiscountApplicable bool {
    return p.age <= 12
}
// robert.IsJuniorDiscountApplicable()) == false
```

### Performance

Even if a function won't modify your structure, it is still beneficial to pass a pointer for performance reason when the struct contains a lot of data.

### Safety

Go provides a higher level of safety than C with regard to pointers. This is because Go has built-in garbage collection, which makes it easier to manage memory allocation and deallocation. Go also has a strong type system that helps prevent common pointer-related errors in C, such as dereferencing null pointers or accessing memory that has already been freed.

Go also provides some additional safety features like bounds checking for slices and arrays.

Unlike C, Go has no pointer arithmetic.

![A meme on C pointers](content/go-techniques/pointers/pointer.png){ width=249px }
