# treasure_hunt
## Prerequisite

1. Since I’m using go.mod , to test my API, please supported Go version for go.mod (I am using go 1.17)
2. pull this repo
3. run 
```
go run main.go
```

the prompt will be like this
```
# # # # # # # #
# . . . . $ . #
# . # # # . $ #
# . . . # . # #
# x # . $ . $ #
# # # # # # # #
```

```
Legend
‘x’ : your position
‘#’: the blocker (obstacle)
‘. ‘: the path
‘$’ : the potential treasure location
```

the treasure will be located if you follow this, in order:
	1. move to up 3 steps, then
	2. move right 5 steps, then
	3. move down 1 step 

if you get the treasure, the prompt will be like this
```
# # # # # # # #
# . . . . . . #
# . # # # . W #
# . . . # . # #
# . # . $ . $ #
# # # # # # # #
CONGRATULATION YOU HAVE FOUND THE TREASURE
```
4. Done!
