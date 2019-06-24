# Conway's Game of Life

Simple demo using Go. It uses a finite grid, anything that goes beyond the grid is ignored.

## Build

```
brew install go # (if you don't already have go installed)
go build
```

## Example
`gameoflife` takes input from stdin. You can seed a grid by typing on the fly, or you can `cat` a given example file to see cool patterns.
```
go build
cat examples/turns_into_a_loaf | ./gameoflife
cat examples/glider | ./gameoflife --maxX 50 --maxY 50
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
