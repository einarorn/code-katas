# Code Katas

## [001] Game of Threes
Given any number, repeatedly divide the number by 3 till you reach 1. Add or subtract 1 whenever division by 3 is not possible. At each stage output the number
### Run tests
```
go test .\game-of-threes -v
```
### Run program
Program takes in an argument which is the number to run against
```
go run .\game-of-threes\cmd\ 100
```

## [002] Rotate
Write a function that rotates a slice by k elements. For example ```[1,2,3,4,5,6]``` rotated by 2 becomes ```[3,4,5,6,1,2]```
### Run tests
```
go test .\rotate -v
```

## [003] Square or Root
Write a method, that will get an integer array as parameter and will process every number from this array.

Return a new array processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example
```[4,3,9,7,2,1]``` -> ```[2,9,3,49,4,1]```

### Run tests
```
go test .\square-or-root -v
```

## [004] Merge lists
Write a function that merges two sorted lists of the same length into a new sorted list.

```[1,4,6]```, ```[2,3,5]``` → ```[1,2,3,4,5,6]```

You can do this quicker than concatenating them followed by a sort.
### Run tests
```
go test .\merge-lists -v
```