# Performance Analysis
This document describes the performance analysis of the project using pprof.

## Allocs profile
```shell
go tool pprof http://localhost:8000/debug/pprof/allocs
```
![allocs](./imgs/allocs.png)

## Goroutine profile
```shell
go tool pprof http://localhost:8000/debug/pprof/goroutine
```
![goroutine](./imgs/goroutine_profile.png)

## Heap profile
```shell
go tool pprof http://localhost:8000/debug/pprof/heap
```
![heap](./imgs/heap.png)