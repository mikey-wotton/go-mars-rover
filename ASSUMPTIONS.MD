#How to run
To run the exe:
```
cd go-mars-rover/
go get ./...
go generate ./...
go build
go-mars-rover.exe
```
To run the code:
```
cd go-mars-rover/
go get ./...
go generate ./...
go run main.go
```
To run unit tests:
```
cd go-mars-rover/
go get ./...
go generate ./...
go test ./...
```
To generate and view code coverage:
```
cd go-mars-rover/
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```


##Packages
###Main
Just a quick executable to demonstrate the code running, printing out the start and end locations of each rover.
Only demonstrates the example provided, you can find further examples inside /rover/rover_test.go.
###Rover
Contains the Rover struct and receiver functions for Rover behaviour, namely turn or move. 
* Rovers cannot crash into each other if stopping on the same (X,Y)
* Rovers cannot leave the boundaries provided through any direction

###Parser
Takes in a string and produces a slice of Rovers or an error. 
* Rovers must be parsed in a valid state
    * Not nil
    * Within boundaries
    * Has a valid Direction (North, East, South, West)
    * Has at least one valid command (L, M, R)
* If any rover produces an error, parsing will stop and return a nil slice and the error.
* Expects exactly 2 Boundary values. Top right coordinates of zone (X, Y)
* Expects exactly 3 Rover initialisation values, representing the Rover position.
* Expects exactly 1 Rover commands string, which must not be empty.
