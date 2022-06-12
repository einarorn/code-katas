# Code Kata's

## [001] Game of Threes
Given any number, repeatedly divide the number by 3 till you reach 1. Add or subtract 1 whenever division by 3 is not possible. At each stage output the number
### Run tests
```
go test ./game-of-threes -v
```
### Run program
Program takes in an argument which is the number to run against
```
go run ./game-of-threes/cmd 100
```

## [002] Rotate
Write a function that rotates a slice by k elements. For example ```[1,2,3,4,5,6]``` rotated by 2 becomes ```[3,4,5,6,1,2]```
### Run tests
```
go test ./rotate -v
```

## [003] Square or Root
Write a method, that will get an integer array as parameter and will process every number from this array.

Return a new array processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example
```[4,3,9,7,2,1]``` -> ```[2,9,3,49,4,1]```

### Run tests
```
go test ./square-or-root -v
```

## [004] Merge lists
Write a function that merges two sorted lists of the same length into a new sorted list.

```[1,4,6]```, ```[2,3,5]``` → ```[1,2,3,4,5,6]```

You can do this quicker than concatenating them followed by a sort.
### Run tests
```
go test ./merge-lists -v
```

## [005] Letter sums
Assign every lowercase letter a value, from ```1``` for ```a``` to ```26``` for ```z```.

Given a string of lowercase letters, find the sum of the values of the letters in the string.

```
lettersum("") => 0
lettersum("a") => 1
lettersum("z") => 26
lettersum("cab") => 6
lettersum("excellent") => 100
lettersum("microspectrophotometries") => 317
```

### Run tests
```
go test ./letter-sum -v
```
### Run program
Program takes in an argument which is the word to sum all letters
```
go run ./letter-sum/cmd settlements
```
## [006] Fibonacci
The Fibonacci sequence, named after mathematician Fibonacci, is a sequence of numbers that looks like this:

```0, 1, 1, 2, 3, 5, 8, 13, ...```

You get first two starting numbers, ```0``` and ```1```, and the next number in the sequence is always the sum of the previous two numbers.

Write a function ```fib```, that takes one parameter ```n```, and returns the nth number from the Fibonacci sequence, counting from ```0```.

For example:

```fib(0)``` returns ```0```

```fib(4)``` returns ```3```

```fib(15)``` returns ```610```

### Run tests
```
go test ./fibonacci -v
```

### Run program
Program takes in an argument which is the word to sum all letters
```
go run ./fibonacci/cmd 100
```