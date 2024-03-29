# Sorting Algorithms

Sorting algorithms put the elements of a collection into a certain order.
Sorting is a basic building block that many other algorithms are built around.
Sorting is used to preprocess the collection to make searching faster (e.g.
[[binary-search|binary search]]), as well as identify items that are similar
(e.g. students are sorted on test scores).

A number of sorting algorithms run in $O(n\logn)$ time - heapsort, merge sort,
and quicksort are examples. Each has its advantages and disadvantages: for
example, heapsort is in-place but not stable; merge sort is stable but not
in-place; quicksort runs $O(n^2)$ time in worst-case. In practice, a
well-implemented quicksort is usually the best choice for sorting.

## Algorithm Properties

Items may be places in **ascending** or **descending** order. The sorting
algorithm may reference a **key** or the entire record. These properties are
often encapsulated in a **comparison function**.

A **stable** sorting algorithm maintains the relative order of the input when
keys are equal. For example, suppose one has an array of people. Each person has
an age and a name. Using a stable sorting algorithm, the array is sorted first
on age, then on name. When you look at the cluster of people named "John,"
you'll notice that they are still organized by age.

**In-place** refers to memory usage. An in-place algorithms needs only constant
($O(1)$) memory beyond the items being sorted. In-place sorting algorithms
usually work within and overwrite the input array. Note that sometimes
$O(\log{}n)$ additional memory is considered in-place.

## Algorithms

### Bubble Sort

#### Overview

Bubble sort compares adjacent neighbors, and if out of order, exchanges them. At
the end of the first pass, the largest is at the end. Sorting usually requires
passes proportional to the input.

|                    |          |
| ------------------ | -------- |
| Best Complexity    | $O(n)$   |
| Average Complexity | $O(n^2)$ |
| Worst Complexity   | $O(n^2)$ |
| Memory Usage       | In-place |
| Stability          | Stable   |

#### Implementation

```python
def bubble_sort(array):
    while True:
        modified = False
        for i in range(len(array) - 1):
            if array[i] > array[i + 1]:
                array[i], array[i + 1] = array[i + 1], array[i]
                modified = True
        if not modified:
            return array
```

### Selection Sort

#### Overview

Selection sort maintains a sorted and an unsorted portion of the list. Initially
the sorted portion is of length zero. Each iteration of the algorithm _selects_
the smallest element in the unsorted portion and places it at the end of the
sorted portion.

|                    |          |
| ------------------ | -------- |
| Best Complexity    | $O(n^2)$ |
| Average Complexity | $O(n^2)$ |
| Worst Complexity   | $O(n^2)$ |
| Memory Usage       | In-place |
| Stability          | Stable   |

#### Implementation

```python
def selection_sort(array):
    for i in range(len(array)):
        for j in range(i, len(array)):
            if array[i] > array[j]:
                array[i], array[j] = array[j], array[i]
    return list
```

### Insertion Sort

#### Overview

Just like selection sort, insertion sort maintains a sorted and an unsorted
portion of the list. Initially the sorted portion is of length zero. Each
iteration of the algorithm takes the first element in the unsorted portion and
_inserts_ it into the correct location in the sorted portion.

|                    |          |
| ------------------ | -------- |
| Best Complexity    | $O(n)$   |
| Average Complexity | $O(n^2)$ |
| Worst Complexity   | $O(n^2)$ |
| Memory Usage       | In-place |
| Stability          | Stable   |

#### Implementation

```python
def insertion_sort(array):
    for i in range(len(array)):
        for j in range(i, 0, -1):
            if array[j] < array[j - 1]:
                array[j], array[j - 1] = array[j - 1], array[j]
            else:
                break
    return array
```

### Shellsort

#### Overview

Shellsort repeatedly applies another sorting algorithm (usually insertion sort)
to subsections of the list. The subsections are defined by a _gap sequence_.
Let's say our list is \[x1, x2, x3, x4, x5, x6, x7\] and our gap sequence is
\[3, 1\]. In the first pass, Shellsort would have a gap of 3 and therefore sort
\[x1, x4, x7\]. In the second pass, Shellsort would have a gap of 1 and sort
\[x1, x2, x3, x4, x5, x6, x7\] (the entire list).

|                    |                         |
| ------------------ | ----------------------- |
| Best Complexity    | Depends on gap sequence |
| Average Complexity | Depends on gap sequence |
| Worst Complexity   | Depends on gap sequence |
| Memory Usage       | In-place                |
| Stability          | Unstable                |

#### Implementation

The following function uses a $2^k - 1$ gap sequence. This gap sequence gives
the algorithm a complexity of $O(n)$.

```python
def shellsort(array):
    gaps = [2**k - 1 for k in range(1, int(math.log(len(array) + 1) /
                                           math.log(2)))]
    for gap in reversed(gaps):
        for i in range(0, len(array), gap):
            for j in range(i, 0, -gap):
                if array[j] < array[j - gap]:
                    array[j], array[j - gap] = array[j - gap], array[j]
                else:
                    break

    return list
```

### Merge Sort

#### Overview

Merge sort is a classic divide-and-conquer algorithm. The algorithm first
divides the input into smaller and smaller lists. At the base case (list length
= 1) the list is sorted. These sorted sublists are progressively _merged_ until
we have sorted the original list.

|                    |              |
| ------------------ | ------------ |
| Best Complexity    | $O(n\log n)$ |
| Average Complexity | $O(n\log n)$ |
| Worst Complexity   | $O(n\log n)$ |
| Memory Usage       | $O(n)$       |
| Stability          | Stable       |

#### Implementation

The efficiency of merge sort depends upon how we combine the two sorted halves
into a single sorted list. We need to _merge_ the two lists together. Observe
that the smallest overall item in the two sorted lists must sit at the top of
one of the two lists. To merge, we remove the smallest element, then repeat.
Because the recursion goes $\lg n$ levels deep, and a linear amount of work is
done per level, merge sort takes $O(n
    \log n)$ time in the worst case.

```python
def merge_sort(array):
    if len(array) < 2:
        return array
    else:
        middle = len(array) / 2
        left, right = merge_sort(array[:middle]), merge_sort(array[middle:])
        return merge(left, right)

def merge(array1, array2):
    """Combine two sorted lists into one sorted list."""
    merged = []
    while array1 or array2:
        if not array2 or (array1 and array1[0] < array2[0]):
            merged.append(array1.pop(0))
        else:
            merged.append(array2.pop(0))
    return merged
```

### Quicksort

#### Overview

Like merge sort, quicksort is a divide-and-conquer algorithm. In merge sort, the
hard part is combining the sublists. In quicksort, the hard part is dividing the
list. Quicksort first chooses a _pivot_. The input is then divided into two
parts: one with elements smaller than the pivot and one with elements larger
than the pivot. We place the pivot between the other two piles, and then sort
piles independently.

|                    |                   |
| ------------------ | ----------------- |
| Best Complexity    | $O(n\log n)$      |
| Average Complexity | $O(n\log n)$      |
| Worst Complexity   | $O(n^2)$          |
| Memory Usage       | Extra $O(\log n)$ |
| Stability          | Stable            |

Quicksort runs in $O(n * h)$, where $h$ is the height of the recursion tree.
Suppose, luckily, we always the median element, the subproblems are always half
the size of the previous level. This produces $O(n \log n)$, the best case of
quicksort. Suppose, unluckily, we always choose the biggest or smallest element
in the sub-array. This produces $O(n^2)$, the worst case of quicksort.

Quicksort is typically 2-3 times faster than merge sort or heapsort when
implemented well. All three algorithms are $O(n \log n)$, but experimentation
shows that the simpler operations in the inner loop give quicksort a constant
improvement.

#### Implementation

The following implementation uses the leftmost element as the pivot.
Unfortunately, this choice produces worst-case performance on sorted lists. Most
implementations will therefore select a different pivot.

```python
def quicksort(array):
    if len(array) < 2:
        return array
    else:
        pivot = array[0]
        left = [x for x in array[1:] if x <= pivot]
        right = [x for x in array[1:] if x > pivot]
        return quicksort(left) + [pivot] + quicksort(right)
```

Note that the implementation makes new lists at each logarithmic step. It's
possible to implement quicksort with only $O(\log n)$ extra memory (for the
stack). See the
[Hoare partition scheme](https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme/)
for an approach that progressively switches elements around a central pivot.

#### Randomization

Randomization is a powerful tool to improve algorithms with bad worst-case but
good average-case complexity.

If we randomly choose the pivot in quicksort, we can expect, with high
probability, $O(n \log n)$. The best possible selection for the pivot is the
median. Suppose a key is good enough if it lies in the center half of the sorted
space of keys. Since the expected number of good splits and bad splits is the
same, the bad splits can only double the height of the tree, which still
produces $O(\log n)$ height. This randomization may be done by either shuffling
the array first or by selecting a random index at each step.

### Heapsort

#### Overview

Selection sort (see above) is a simple algorithm that repeatedly extracts the
smallest remaining element from the unsorted part of an array. A computer takes
$O(n)$ time to find the smallest element in an array. This is the operation
supported by a priority queue. What if we improve the data structure?
**Heapsort** is nothing but an implementation of selection sort using the right
data structure. Heapsort uses a [heap](heaps.org).

|                    |              |
| ------------------ | ------------ |
| Best Complexity    | $O(n\log n)$ |
| Average Complexity | $O(n\log n)$ |
| Worst Complexity   | $O(n\log n)$ |
| Memory Usage       | $O(1)$       |
| Stability          | Unstable     |

#### Implementation

Heapsort creates a heap and repeatedly extracts the minimum to give a worst-case
$(O \log n)$ algorithm. Heapsort can be implemented as an in-place sort.

```python
from heapq import heappush, heappop

def heapsort(list):
    heap = []
    for item in list:
        heappush(heap, item)
    return [heappop(heap) for _ in range(len(heap))]
```

### Distribution Sort

Suppose we have a list of names to sort. We could partition them according to
the first letter. This creates 26 different piles, or buckets, or names. Then,
we partition each pile based on the second letter of each name, etc. The names
will be sorted as soon as each bucket contains only a single name. At the end,
we'll be able to simply concatenate the bunch of piles together. This algorithm
is commonly called **bucketsort** or **distribution sort**.

**Bucketing** is a very effective idea whenever we are confident that the
distribution of data will be roughly uniform. It is the idea that underlies hash
tables, kd-trees, and a variety of other practical data structures. The downside
of such techniques is that the performance can be terrible when the data
distribution is not what we expected.

### <span class="todo TODO">TODO</span> Tapesort

See <https://en.wikipedia.org/wiki/External_sorting#External_merge_sort>
