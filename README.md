# goes

A toy programming language inspired by Go and Rust that transpiles to Go.

Do not use in production.

<!--
## Example

See an [example](/testdata/program.goes).

# Syntax

## Types

```go
// This is a struct.
type Parser ()

// This is a struct with fields.
type Parser (src str, idx int)

// This is a struct with an embedded struct.
type Parser (Scanner)

// This is a trait with no behavior.
type Thing () ()

// This is a trait with behavior.
type Read  (p []byte) (n int, err error)
type Write (p []byte) (n int, err error)

// This is a trait composition.
type ReadWriter Read Write

// This is how to implement a trait.
func (*Parser) Read (p []byte) (n int, err error) { ... }
```

## Functions, Methods

```go
// This is a method.
func (*Parser) Match(string) bool { return false }

// This is a method overload.
func (*Parser) Match(rune) bool { return false }

// This is a function.
func DoNothing() {}
```

## Error handling

```go
func Example() error {

    // Variables are not constrained inside the if scope.
    if v, err := DoSomething(); err != nil {
        return err
    }

    // Syntax sugar for the above.
    v = DoSomething()?

    _ = v
    return nil
}
```

## Logical operations

```go
func Example() {

    if true & true {
        fmt.Println("and")
    }

    if false | true {
        fmt.Println("or")
    }

    fmt.Println(true ? "yes" : "no")
}
```

## Captures

A capture is a variable declaration that happens in its first use.

It can be used to:

1. Capture an output argument
1. Capture the return value of a function
1. Capture the arguments of a lambda/anonymous function

### Output argument

```go
func Example() {

    p := Parser("Goes101")

    if p.Match(WORD, :w) & p.Match(DIGITS, :d) {
        fmt.Println("Word: %s, Digits: %s", w, d)
    }
}

// This is how Match is defined.
func (*Parser) Match(r *regexp.Regexp, out *Token) bool
```

### Return value

```go
func Example() {

    p := Parser("Goes101")

    if p.Token(WORD):w & p.Token(DIGITS):d {
        fmt.Println("Word: %s, Digits: %s", w, d)
    }
}

// This is how Token is defined.
func (*Parser) Token(r *regexp.Regexp) Token
```

If wondering how is a token treated as a boolean, see the `Truthy` trait in the next section.

### Lambda function

```go
func Example() {
    v := [1, 2, 3].Map(:x * 2)
    fmt.Println(v)
}

// This is how Map is defined.
func (*List) Map(f func(int) int) *List
```

## Traits

Traits are behavior contracts added to types.

Goes has a few default traits.

### Truthy

The `Truthy` trait allows a type to be used in a boolean context.

```go
type Token (Text str)

// Token struct now implements the Truthy trait
// and can be used in a boolean context.
func (*Token) Truthy() bool {
    return .Text != ""
}

func Example() {
    a := true
    b := Token("Hi")
    if a & b {
        fmt.Println(a, b)
    }
}
```
