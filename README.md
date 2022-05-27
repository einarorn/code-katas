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
rite a function that rotates a slice by k elements. For example ```[1,2,3,4,5,6]``` rotated by 2 becomes ```[3,4,5,6,1,2]```
### Run tests
```
go test .\rotate -v
```