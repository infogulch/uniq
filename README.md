uniq
====

A Golang package to get the first unique elements of slices or user-defined collections from an already sorted list. Uses your existing sort.Interface

Usage
=====

The list *must* already be sorted, otherwise the behavior is undefined. Since you just sorted it with `sort.Sort` (you did, right?) you can call `uniq.Ue` with the `sort.Interface` that you just created, because uniq uses it verbatim. `uniq.Ue` will call `Len` once, `Less` up to `n+1` times, and `Swap` up to `n-1` times and return the length of the unique portion of the array. The remaining part of the array will contain all the duplicates but is no longer sorted.

Examples
========

```
todo
```
