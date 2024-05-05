# Description

Uses the input data in `input/Tracks.csv' to determine the shortest route by distance 
for a sample set of pairs of  Timing point locations (TIPLOCs)
The results are output to `output/sample-output.csv'
Uses an implemetation pf Dijkstra's algorithm based on this [repo](https://github.com/rishabh625/graphs)

# Pre-requisites

[Install](https://go.dev/doc/install) the Go version specified in `go.mod` file.
Have a suitable IDE eg VSCode


# Build, Test and Run
Open a ternimal and navugate to the root foldet of this project
Use the follwing commands

Get dependencies
```
go get ./...
```

build 
```
go build ./...
```

run 
```
go run ./...
```

run unit tests
```
go test ./...
```


