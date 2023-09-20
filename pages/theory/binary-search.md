# Binary Search

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

**Binary search** is a fast algorithm for searching in a sorted array. We
compare the key $q$ to the middle item. If $q$ is smaller, it must appear in the
first half; if not it must reside in the second half. By repeating this process
recursively on the correct half, we locate the key in $\lg n$ comparisons.

A disadvantage of binary search is that it requires a sorted array and sorting
an array takes $O(n\log n)$ time. However, if there are many searches to
perform, the time taken to sort is not an issue.

Binary search is the power behind twenty questions!

## Implementation

```python
def binary_search(item, array, start=None, end=None):
    """Precondition: array is sorted"""
    if start is None or end is None:
        start, end = 0, len(array)

    middle_index = (start + end) // 2
    if start >= end:
        return False
    elif item == array[middle_index]:
        return True
    elif item < array[middle_index]:
        return binary_search(item, array, start, middle_index)
    else:
        return binary_search(item, array, middle_index + 1, end)
```

## Libraries

The `bisect` module provides binary search functions in Python.
`bisect_left(array, item)` finds the first element that is not less than a
target value. `bisect_right(array, item)` finds the first element that is
greater than a target value.

```python
from bisect import bisect_left

def binary_search(item, array):
    index = bisect_left(array, item)
    return 0 <= index < len(array) and array[index] == item
```

## Related Algorithms

### Counting Occurrences

Suppose we want to count the number of times a given key $k$ occurs in a given
sorted array. We could use binary search to find the index of an element in the
correct block in $O(\lg n)$ time. Then we sequentially test elements to the left
and right until we find one that differs from the key. The difference between
the boundaries (plus one) gives the count of the number of occurrences of $k$.
This algorithm runs in $O(\lg n + s)$, where $s$ is the number of occurrences of
the key.

A fast algorithm results by modifying binary search to search for the _boundary_
of the block containing $k$, instead of $k$ itself. We perform this search
twice, for a total time of $O(\lg n)$, so we can count the occurrences in
logarithmic time regardless of the size of the block.

### One-Sided Binary Search

Suppose we don't know the bounds of our sorted collection. Binary search can
also proceed from a specific position at repeatedly larger intervals (1, 2, 4,
8, 16) until we find a value greater than our key. We now have a window
containing the target and can proceed with binary search. _One-sided binary
search_ is most useful whenever we are looking for a key that lies close to our
current position.

### Square and Other Roots

Suppose we are searching for the square root $r$ of $n$. Notice that the square
root of $n \geq 1$ must be at least 1 and at most $n$. Consider the midpoint $m$
of this interval. How does $m^2$ compare to $n$? If $n \geq m^2$, then the
square root must be greater than $m$, so the algorithm repeats on a new range of
values. This application of binary search identifies the square root within ±1
after only $\lg n$ rounds. Root-finding algorithms that converge faster are
known, but this is simple, robust and applies to other functions.
