# Data Structures

Classes of **abstract data types** such as containers, dictionaries, and
priority queues, have many different but functionally equivalent **data
structures** that implement them. These different data structures realize
different tradeoffs in the time to execute various operations.

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

## Contiguous vs. Linked

Data structures are either **contiguous** or **linked**, depending upon whether
they are based on arrays or pointers.

### Arrays

The **array** is the fundamental contiguously-allocated data structures. These
single slabs of memory have constant access given the index and space
efficiency. **Dynamic arrays** enable resizing. First, an initial size is
allocated. If we run out of space, a larger array (usually 2x) is allocated and
the elements are copied over. Insertion amortizes to $O(1)$.

### Lists

The **list** is the simplest linked structure. Each node in the list has data
and pointer fields. **Pointers** are the connections that hold the pieces of
together. Pointers represent the address of a location in memory. List don't
incur overflow, but require extra space for pointer fields and don't given
efficient random access to items.

## Abstract Data Types

### Containers

A **container** denotes a data structure that permits storage and retrieval of
data items independent of content. Containers are distinguished by the
particular retrieval order they support. **Stacks** support retrieval by
last-in, first-out (LIFO) order. **Queues** support retrieval in first in, first
out (FIFO) order.

See [Bags, Stacks, and Queues](bags-stacks-queues.org) for more information.

### Dictionaries

The **dictionary** data type permits access to data items by content. Operations
include _search_ (when given a _key_), _insert_, and _delete_. Dictionaries are
also called **maps** and **associative arrays**.

[Hash tables](./hash-tables.org) are a very practical way to maintain a
dictionary.

### Trees

See [Trees](./trees.org) for more information.

### Priority Queues

**Priority queues** are data structures that provide more flexibility than
simple sorting, because they allow new elements to enter a system at arbitrary
intervals. The basic priority queue supports three primary operations: _insert_,
_find-minimum/maximum_, and _delete-minimum/maximum_. Priority queues can be
implemented with arrays or BSTs, but a particularly nice implementation is the
[heap](./heaps.org).
