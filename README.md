# Conway's Game of Life

Simple demo using Go. It uses a finite grid, anything that goes beyond the grid is ignored.

## Build

```
brew install go # (if you don't already have go installed)
go build
```

Alternatively, download the binary with `go get` (nifty!)
```
go get github.com/cjlint/gameoflifedemo
$(go env GOPATH)/bin/gameoflifedemo
```

## Example
`gameoflife` takes input from stdin. You can `cat` a given example file to see cool patterns.
```
cat examples/turns_into_a_loaf | ./gameoflife
cat examples/glider | ./gameoflife --maxX 50 --maxY 50
```

Or you can seed a grid by typing 0s and 1s directly to stdin, followed by an empty line
```
./gameoflife
0010010011
1001001
01010111001
10
100
10010

```

## Command line args
You can tinker with the command line arguments if you want to try a larger grid or a faster/slower interval
```
./gameoflife -help
```

## Tests

There are a few unit tests but not full coverage
```
go test
```
