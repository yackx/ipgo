## struct

A **struct** is a typed collection of fields. It allows to group data together to form records.

### Motivation

Suppose we want to represent a person. What data would we capture? What characteristics represent a person? It depends on the context. The medical department of a hospital may need to collect: first name, last name, date of birth, blood type, gender, weight and height and plenty of other medical information. The administrative department of the same hospital is likely to be interested in data such as the social security number and billing information.

In our examples, we will limit ourselves to name and age.

So far, we have used "atomic" data types, like `string` and `int`. A `string` is suitable to store a person's name. An `int` is suitable to store their age. It would however be inconvenient to use two independent variables to store those two characteristics.

```go
name := "John"
age := 32
```

Arguably, this is not bad, but what if we want to represent two persons?

```go
person1Name := "John"
person1Age := 32
person2Name := "Marge"
person2Age := 43
```

We only have two persons, each described by two fields, and things are already messy. Notice how these four variables are unrelated. Nothing expresses the fact that `person1Name` and `person1Age` are information related to the same person. Sure, they bear a common prefix `person1` that gives the reader some clue. But this alone does not create the representation of a person. We are left with a bunch of unrelated charactersitics. Imagine handling five persons, each with ten fields using this (lack of) technique. It would not be good programming design.

Instead, we can declare a **struct** that acts as a **blueprint** for a person -- any person.

```go
type Person struct {
    Name string
    Age  int
}
```

> _Note_: Using `uint8` instead of `int` will use up less memory. This is not a concern here. Also, in an actual information system, the date of birth will be used instead of the age.

We can instantiate our persons, and assign them to variables. Notice how the variables `john` and `marge` each conveniently represent a "whole" person, with all their relevant characteristic bundled in the object.

```go
john := Person{"John", 32}
marge := Person{"Marge", 43}
```

### Arguments

We can be more explicit, by specifying each field name.

```go
john := Person{Name: "John", Age: 32}
marge := Persom{Name: "Marge", Age: 43}
```

What is the difference between these two ways to declare a struct? The second one does not rely on the field position, which prevents mismatching fields. Sure enough, there is little risk of error when declaring our persons with our trivial struct. Indeed, the compiler would refuse inverted name and age, as they are of different types. But suppose we store the first name and the last name. In that case, which one comes first? We would have to look it up. If we accidentally invert them, the program will still compile. If someone decides to re-arrange the field order in the struct for any reason, the program would also still compile, but the names would be mixed up. Spelling out each field is more verbose but avoids mistakes and makes our program more robust if we design anything non trivial.

### Constructor

This being said, it is idomatic to encaspulate struct creation in a function. This function's name starts with `New` by convention, followed by the struct name. In our case, it would be named `NewPerson`. It is also idomatic to return a **pointer** to the newly created struct, denoted by a `*`. More on pointers later. Here is how the constructor would look like:

```go
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}

robert := NewPerson("Robert", 67)
// robert.Name == "Robert"
// robert.Age == 67
```

Unfortunately, we loose the explict naming of arguments. On the other hand, this _constructor_ can ensure consistency between fields and check for invalid input. For instance, it could check for negative age value, and reject it.

### Behaviors

Functions can be "associated" with structs. For instance, let us say we want to "attach" two functions to `Person`: one to determine if a senior discount is applicable, and another to determine if a junior account is applicable. We could proceed with as follows:

```go
func IsJuniorDiscountApplicable(p *Person) bool {
    return p.age <= 12
}

fmt.Println(IsJuniorDiscountApplicable(robert))
// false
```

There is another way to declare this function, as a "behavior" for a person.

```go
func (p *Person) IsJuniorDiscountApplicable bool {
    return p.age <= 12
}
fmt.Println(robert.IsJuniorDiscountApplicable())
// false
```

Observe the difference. The second version may not seem to bring much to the table, if anything at all, but it will prove useful in the chapter about **interfaces**.

### String representation

What if we try to print a person?

```go
john := &Person{"John", 34}
fmt.Println(john)
```

```
&{John 34}
```

The output is readable enough, but it can customized.

```go
func (p *Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}
```

Let's decompose this function.

- It is called `String`.
- It "acts" on a person[^person-pointer] aliased by `p`, as per `(p *Person)`.
- It returns a `string`.
- We use `fmt.Sprintf` to build our custom representation. The first argument of `Sprintf` is the format. `"%s (%d)"` denoted a `string` as per `%s`, followed by a number as per `%d`. The number is between bracket.

[^person-pointer]: More precisely, a **pointer** to a `Person`.

`fmt.Println(person)` will look for a `String()` function applicable to `person` and will find the one we have just defined. If we print `john`, the output is now formatted as we requested.

```
John (34)
```

Every person will be formatted that way when output is requested.

### Variable naming

We already discussed the need to have meaningful variable names. Excessive abbreviations are hard to read. In a book example, names like `p1` and `p2` may be acceptable, but we took the good habit of properly naming variables, e.g. `john` and `robert`.

Therefore, how come we wrote `(p *Person)` rather than `(person *Person)`? Isn't the second version less vague? The former is more compact and there is no risk of confusion, as the type is defined. The latter version introduces a form a stuttering. Furthermore, this `p` is only used "privately" by the function, as opposed to a possible argument that could be used by the caller. It does not "leak" outside.

It is customary in Go code to use the compact form _in this context_. At any case, neither form should be considered "wrong".

> _Note_: Always be consistent in your coding style. Nothing is more annoying than unconsistent coding conventions. It slows down even seasoned programmers.

### Multiple structs

A struct instance does not have to live alone. It can be combined with slices. Everything we have described in the Arrays and Slices chapter remains applicable to structs. Instead of declaring slices of `string` or `int`, we can manipulate slices of `*Person` (a pointer to a person).

```go
persons := []*Person{NewPerson("Robert", 67), NewPerson("Myla", 8)}
for _, person := range persons {
	fmt.Println(person)
}
```

```
Robert (67)
Myla (8)
```

### Exercices

1. Declare a struct `Bike` that holds the following information: brand, model, color, number of bikes in stock.

2. Ensure that, in the bike representation, the color is between brackets and the number of items in stock is between square brackers. For instance `Cowboy Mk2 (black) [4]`.

3. Declare a slice `inventory` to contain an arbitrary number of bikes. Initiliaze it with the following products: Blue-Grey Trek Powerfly (3 in stock), Black-Orange BTwin Triban 540 (2 in stock).

4. Write a function `countInStock` that receives our inventory and returns the total count of bikes (in our example, 5 in stock).
