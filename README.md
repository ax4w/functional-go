![image](https://github.com/user-attachments/assets/83629097-4ddd-4500-8837-2c8ca6fb4dc6)


A lightweight library that provides functional programming utilities for Go, leveraging Go's generics support.

## Overview

Functional-Go implements common functional programming patterns and operations such as map, filter, fold, and more. The library is designed to work with Go's type system and makes extensive use of generics to provide type-safe functional operations.

## Features

- **Tuple operations**: Create and manipulate 2-element tuples
- **List operations**: Common list manipulations like `Head`, `Tail`, `Take`, `Drop`, and `Last`
- **Higher-order functions**: `Map`, `Filter`, and function composition
- **Folding**: Left fold (`Foldl`) and right fold (`Foldr`) operations
- **Zipping**: Combine lists with `Zip` and `ZipWith`
- **Predicate functions**: Test elements with `Any` and `All`
- **Comparison**: Generic comparison with `Compare`
- **Numeric operations**: `Sum`, `Product`, `Maximum`, and `Minimum`
- **Map operations**: Convert maps to lists with `Flatten` and `FlattenWith`

## Installation

```bash
go get github.com/ax4w/functional-go
```

## Usage

```go
package main

import (
    "fmt"
    f "github.com/username/functional-go"
)

func main() {
    // List operations
    nums := []int{1, 2, 3, 4, 5}
    
    // Get first element
    head := f.Head(nums) // 1
    
    // Get all elements except the first
    tail := f.Tail(nums) // [2, 3, 4, 5]
    
    // Get the first n elements
    taken := f.Take(nums, 3) // [1, 2, 3]
    
    // Drop the first n elements
    dropped := f.Drop(nums, 2) // [4, 5]
    
    // Get the last element
    last := f.Last(nums) // 5
    
    // Map: Double each number
    doubled := f.Map(func(x int) int { return x * 2 }, nums) 
    // [2, 4, 6, 8, 10]
    
    // Filter: Keep only even numbers
    evens := f.Filter(func(x int) bool { return x%2 == 0 }, nums) 
    // [2, 4]
    
    // Fold left: Sum all numbers
    sum := f.Foldl(func(acc, x int) int { return acc + x }, 0, nums) 
    // 15
    
    // Fold right: String concatenation with proper order
    strs := []string{"A", "B", "C"}
    concat := f.Foldr(func(x string, acc string) string { 
        return x + acc 
    }, "", strs) 
    // "ABC"
    
    // Zip two slices into tuples
    names := []string{"Alice", "Bob", "Charlie"}
    ages := []int{30, 25, 35}
    people := f.Zip(names, ages)
    // [(Alice, 30), (Bob, 25), (Charlie, 35)]
    
    // ZipWith: Combine two slices using a function
    greetings := f.ZipWith(
        func(name string, age int) string {
            return fmt.Sprintf("%s is %d years old", name, age)
        },
        names, ages,
    )
    // ["Alice is 30 years old", "Bob is 25 years old", "Charlie is 35 years old"]
    
    // Function composition
    addOne := func(x int) int { return x + 1 }
    double := func(x int) int { return x * 2 }
    addOneThenDouble := f.Compose(double, addOne)
    result := addOneThenDouble(3) // (3 + 1) * 2 = 8
    
    // Predicate functions
    hasEven := f.Any(nums, func(x int) bool { return x%2 == 0 }) // true
    allPositive := f.All(nums, func(x int) bool { return x > 0 }) // true
    
    // Comparison examples
    if f.Compare(5, 10) == f.LT {
        fmt.Println("5 is less than 10") // will print
    }
    
    if f.Compare("apple", "banana") == f.LT {
        fmt.Println("'apple' comes before 'banana'") // will print
    }
    
    if f.Compare(true, false) == f.GT {
        fmt.Println("true is greater than false") // will print
    }
    
    // Comparison with custom types
    type Person struct {
        Name string
        Age  int
    }
    
    p1 := Person{Name: "Alice", Age: 30}
    p2 := Person{Name: "Alice", Age: 30}
    
    if f.Compare(p1, p2) == f.EQ {
        fmt.Println("p1 and p2 are the same person") // will print
    }
}
```

## API Reference

### Tuple Operations

- `Tuple[A, B]`: A generic struct holding two values of potentially different types
- `Fst[A, B](t Tuple[A, B]) A`: Extract the first element of a tuple
- `Snd[A, B](t Tuple[A, B]) B`: Extract the second element of a tuple

### List Operations

- `Head[A](src []A) A`: Returns the first element of a slice
- `Tail[A](src []A) []A`: Returns all elements except the first
- `Take[A](src []A, num int) []A`: Takes the first `num` elements
- `Drop[A](src []A, num int) []A`: Drops the first `num` elements
- `Last[A](src []A) A`: Returns the last element

### Higher-Order Functions

- `Map[A, B](fn func(A) B, src []A) []B`: Applies function to each element
- `Filter[A](fn func(A) bool, src []A) []A`: Keeps elements that satisfy predicate
- `Compose[A, B, C](fnB func(B) C, fnA func(A) B) func(A) C`: Composes two functions

### Folding

- `Foldl[A, B](fn func(B, A) B, acc B, src []A) B`: Left fold (accumulate from left to right)
- `Foldr[A, B](fn func(A, B) B, acc B, src []A) B`: Right fold (accumulate from right to left)

### Zipping

- `ZipWith[A, B, C](fn func(A, B) C, srcA []A, srcB []B) []C`: Combines elements using function
- `Zip[A, B](srcA []A, srcB []B) []Tuple[A, B]`: Combines elements into tuples

### Predicate Functions

- `Any[A](src []A, fn func(A) bool) bool`: Returns true if any element satisfies the predicate
- `All[A](src []A, fn func(A) bool) bool`: Returns true if all elements satisfy the predicate

### Comparison

- `ComparisonResult`: An enum type representing the result of a comparison (`LT`, `EQ`, or `GT`)
- `Compare[T any](a, b T) ComparisonResult`: Compares two values of the same type and returns:
  - `LT` if `a` is less than `b`
  - `EQ` if `a` equals `b`
  - `GT` if `a` is greater than `b`
  - Works with primitive types (numbers, strings, booleans) and custom types

### Numeric Operations

- `Sum[A numeric](src []A) A`: Returns the sum of all elements in a numeric slice
- `Product[A numeric](src []A) A`: Returns the product of all elements in a numeric slice (Note: currently always returns 0 due to implementation)
- `Maximum[A comparable](src []A) A`: Returns the maximum element in a slice
- `Minimum[A comparable](src []A) A`: Returns the minimum element in a slice

### Map Operations

- `FlattenWith[A, B, C](fn func(A, B) C, src map[A]B) []C`: Converts a map to a slice by applying a function to each key-value pair
- `Flatten[A, B](src map[A]B) []Tuple[A, B]`: Converts a map to a slice of key-value tuples

## Notes

- Most functions that operate on empty slices will return empty slices or the accumulator
- `Head`, `Tail`, `Last`, `Maximum`, and `Minimum` will panic when called on empty slices
- For zipping operations, the result length is determined by the shorter input slice
- `Any` on an empty slice returns false, while `All` on an empty slice returns true
- `Compare` works best with primitive types; for custom types it can determine equality but ordering is arbitrary
- `Product` currently has a bug where it initializes the accumulator to 0, meaning it always returns 0

## License

MIT License