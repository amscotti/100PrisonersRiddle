# 100 Prisoners Riddle Simulator

This application simulates many sessions of 100 Prisoners Riddle

Default strategy should return with 31% success rate with the steps of,

1. Each prisoner first opens the drawer labeled with their own number.
1. If this drawer contains their number, they are done and were successful.
1. Otherwise, the drawer contains the number of another prisoner, and they next open the drawer labeled with this number.
1. The prisoner repeats steps 2 and 3 until they find their own number, or fail because the number is not found in the first fifty opened drawers.

## Resources
* [100 prisoners problem - Wikipedia](https://en.wikipedia.org/wiki/100_prisoners_problem)
* [The Riddle That Seems Impossible Even If You Know The Answer - YouTube](https://www.youtube.com/watch?v=iSNsgj1OCLA)

## Build
Run `go build` to build

## Usage
```
Usage of ./100PrisonersRiddle:
  -n int
        Number of sessions to run (default 1000000)
  -s string
        Type of strategy to use, 'loop', 'random', 'true', 'false' (default "loop")
```