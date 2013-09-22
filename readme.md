# Goshawk

A package for high performance scientific computing in Go.

## Install

To install simply run:

```
go get github.com/rwl/goshawk
```

## Test

To run the tests execute:

```
go test github.com/rwl/goshawk
```

<a href="https://travis-ci.org/rwl/goshawk" target="_blank">
  <img src="https://api.travis-ci.org/rwl/goshawk.png" alt="Build Status">
</a>

## Benchmark

To benchmark Goshawk run:

```
export GOMAXPROCS=2
go test -run=ZZZ -bench=. github.com/rwl/goshawk
```
