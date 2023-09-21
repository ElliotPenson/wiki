# Bags, Stacks, and Queues

**Bags**, **stacks**, and **queues** are **containers** with different retrieval
orders.

## Bags

A bag (also called a **multiset**) is a collection that stores items in an
unordered fashion without support for removal. The data structure is usually
used to collect items and then iterate through them.

### Operations

- `add(item)`
- `isEmpty()`
- `size()`

The container would also be iterable.

### Implementation

A bag can be implemented in a similar fashion to a stack (see below). The _pop_
operation needs to be removed and _push_ needs to be changed to _add_.

## Stacks

A (pushdown) stack is a collection based on the last-in, first-out (LIFO)
policy. For example, plates are _stacked_ in the kitchen. The final plate added
to the stack is the first one to be removed.

### Operations

- `push(item)`
- `pop()`
- `isEmpty()`
- `size()`

### Implementation

A stack can be implemented with an array. The program should maintain an array
and a size (integer). Pushing an element increases `size` and adds the element
to `array[size]`. Popping returns the element at `array[size]` and reduces
`size` by one. If the array is dynamically resized, the amortized time for both
push and pop is `O(1)`.

A linked list can also be used. This method of course requires references. We
only need to maintain the head of the linked list for a stack implementation. A
push appends to the front, a pop removes the first element.

#### Python

```python
stack = []

stack.append('item')  # push
stack[-1]             # peek
stacks.pop()          # pop
len(stack) == 0       # isEmpty
```

## Queues

A queue is a collection based on the first-in, first-out (FIFO) policy. A real
world example is the checkout line at a store. The next person added to the line
has to wait until all of the people in front of him/her have left the line.

### Operations

- `enqueue(item)`
- `dequeue()`
- `isEmpty()`
- `size()`

### Implementation

Just like a stack, a queue can be implemented with an array. This time, we must
maintain a head index and a tail index. This structure is sometimes referred to
as a **circular queue**.

A linked list can also be used. The algorithm maintains a reference to the first
and last elements in the linked list. Dequeuing means removing and returning the
first element and updating the head reference. Enqueuing means changing the next
node in the last element and then updating the reference to last.

#### Python

Python provides queues with `collections.deque`. A _deque_, also sometimes
called a double-ended queue, is a doubly linked list in which all insertions and
deletions are from two ends of the list.

```python
from collections import deque

queue = deque()

queue.append('item')   # add
queue[0]               # peek
queue.popleft('item')  # remote
len(queue) == 0        # isEmpty
```
