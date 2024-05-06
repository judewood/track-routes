# Description

Uses the input data in `input/Tracks.csv' to determine the shortest route by distance 
for a sample set of pairs of  Timing point locations (TIPLOCs)
The results are output to `output/sample-output.csv'
Uses an implementation pf Dijkstra's algorithm based on this [repo](https://github.com/rishabh625/graphs)

# Assumptions and Limitations

## Assumptions
The calculation assumes:

1. Only route sections where PASSENGER_USE equates to yes are considered since this calculation is for a railway timetable

2. The calculation is for non-electric trains so the value of ELECTRIC is ignored since electric trains can only travel on electric lines

3. Trains can switch lines at TIPLOC locations using points. I am assuming therefore that the LINE_CODE is not relevant and that trains can switch between LINE_CODES at any TIPLOC. This may be incorrect but I am unable to find what these LINE_CODES are

4. Some pairs of route sections have multiple defined distances and some of these are zero. It seems illogical for zero to be correct and we want the shortest route. For these pairs the code picks the shortest non-zero distance

## Limitations

1. Currently not producing identical sample answers to the given ones. The results are similar so likely due to the above assumptions. I am investigating this

## Further improvements

1. Full unit test coverage
2. Move graph related models to models package
3. Benchmark parallel vs non parallel code to see what the improvement is
4. Cache non-connected TIPLOC pairs to a lookup file to speed up later runs.
5. Create a CSV reader based on the library one that rejects not needed rows before pulling them into memory to save memory usage for large data sets


# Pre-requisites

[Install](https://go.dev/doc/install) the Go version specified in `go.mod` file.
Have a suitable IDE eg VSCode


# Build, Test and Run
Open a terminal and navigate to the root folder of this project
Use the following commands

Get dependencies
```
go get ./...
```

build 
```
go build ./...
```
to create executable file

```
go build github.com/judewood/routeDistances
```

run 
You can run the included routeDistance.exe file
To run  local build
```
go run ./...
```

run unit tests
```
go test ./...
```


