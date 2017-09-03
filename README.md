gtm - Go Tierra memory allocator
================================

This is a Go port of the memory allocator of the Tierra artificial life simulator.

## References

* The original [Tierra](http://life.ou.edu/tierra/) simulation by Tom Ray.
* The OS X port [MacTierra](http://www.smfr.org/work/sfi/mactierra/).

## Analysis of original C code

A `mal` instruction is decoded into a call of the `malchm` function
in `instruct.c`.

`malchm` performs some simple size checks and then calls `mal` in `memalloc.c`.

The signature of `mal` is:
```c
I32s mal(I32s *sug_addr, I32s sug_size, I32s mode)
```

Where the parameters are:

* `sug_addr` is the address of the allocated block. This is an output parameter.
* `sug_size` is the desired size of the block.
* `mode` is the allocation mode.

The return value is the actual allocated size or 0 in case of error.
The allocated size can be different from the requested one because of
flaws.

The main job of `mal` is preparing a number of parameters dependent on the
allocation mode.

In this implementation we will only replicate the _Better fit_ algorithm that
corresponds to mode `1`.

Regardless of the specified mode, `mal` proceeds to call repeatedly the
`MemAlloc` and the `reaper` (in `tierra.c`) functions until one of the following
two conditions is satisfied:

1. `MemAlloc` succeeds and returns a valid address.
1. `reaper` can't find a suitable cell to kill.

`MemAlloc` has the following signature:
```c
I32s MemAlloc(I32s size, I32s pref, I32s tol)
```

Where the parameters are:

* `size` is the requested size in memory slots. (In our case bytes.)
* `pref` is the preferred soup address.
* `tol` is the acceptable tolerance.

When the allocation mode is 1 (_Better fit_) the `pref` parameter is set to the
special value of `-1` and the tolerance is set to `0`.

The body of the `MemAlloc` function is actually split in two separate sections:
one for handling all _Friendly fit_ allocation modes (mode != 1), and one for
the _Better fit_ algorithm (mode == 1).

The _Better fit_ algorithm searches the original data structure for the smallest
free segment that still larger than the requested size. Once found, allocation
is performed on its _left side_ i.e. towards smaller addresses.

## Data structures

The original Tierra relies on an hand-coded [cartesian tree](https://en.wikipedia.org/wiki/Cartesian_tree)
for storing the memory segments.

The best data structure for the _Better fit_ mode is something that keeps the
collection of free segments sorted by size for easily searching the appropriate 
size.

Segments of the same size could be sorted by address or not. When kept sorted,
allocation could happen easily towards one of the boundaries of the soup.
In the other case allocation could happen randomly within the soup.

Go does not have any sorted data structure in its standard library but
implementations of trees that keep their nodes sorted according to a given
comparator are available.
