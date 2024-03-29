# Hash Tables

A **hash table** implements the _dictionary_ (or _map_, _associative array_)
abstract data type. A dictionary maintains a set of key/value pairs. These data
structures support four operations.

- _Insert_ a pair into the collection.
- _Delete_ a pair from the collection.
- _Modify_ the value associated with a given key.
- _Search_ for a given key's value.

Hash tables have the best theoretical and real-world performance for these
operations at $O(1)$. The $O(1)$ time complexity for insertion is for the
average case - a single insert can take $O(n)$ if the hash table has to be
resized.

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

## Implementation

We could implement the operations above with an array (keys would be indexes).
If we have a large number of potential keys, this array would be massive (or
even infinitely large given string keys). But lookup would be really fast.
Another implementation could store each key/value as a node in a linked list.
This collection would take up far less space. But lookup would be quite slow.
The **hash table** tries to have both advantages (fast lookup and space
efficiency). All operations on a hash table have constant average complexity
($O(1)$).

A hash table maintains an array. The data structure converts a given key to a
position in the array (done by a **hash function**). A collision occurs when two
keys go to the same location in the array. Collisions happen because the number
of keys is usually greater than the number of array slots. Many methods exist
for resolving collisions.

### Collision Resolution

Two distinct keys will occasionally hash to the same value. This is a
**collision**. **Chaining** is the easiest approach to collision resolution.
Represent the hash table as an array of $m$ linked lists (each called a
_bucket_). After hashing the key, we search through the linked list to find the
value. Chaining devotes a considerable amount of memory to pointers. **Open
addressing** is an alternative to chaining. The hash table is maintained as an
array of elements, each initialized to null. On an insertion, we check to see if
the desired position is empty. If so, we insert it. If not, we must find some
other place to insert it instead. The simplest possibility (called **sequential
probing**) inserts the item in the next open spot in the table.

## Hash Function

Hash tables use **hash functions** to map keys to integers. A **hash function**
is a function that maps a large data set to a smaller data set of a fixed
length. Hash functions return a hash code (or hash value).

Every hash function must:

- Consistently produce the same location for a key (a pure function).
- Uniformly distribute values over a table (to reduce collisions).
- Be efficient to compute.

The first step of the hash function is usually to map each key to a big integer.
Let $\alpha$ be the size of the alphabet on which a given string $S$ is written.
Let `char(c)` be a function that maps each symbol of the alphabet to a unique
integer from 0 to $\alpha - 1$. The function

$$
H(S) = \sum_{i = 0}^{|S| - 1} \alpha^{|S| - (i + 1)} \times char(s_i)
$$

maps each string to a unique (but large) integer by treating the characters of
the strings as "digits" in a base-$\alpha$ number system.

The result is unique identifier numbers, but they are so large they will quickly
exceed the number of slots in our hash table (denoted by $m$). We must reduce
this number to an integer between $0$ and $m - 1$, by taking the remainder of
$H(S)$ mod $m$.

### Hash Function Applications

#### String Matching via Hashing

The **Rabin-Karp algorithm** gives a linear-time solution to substring search.
Substring search asks if string $t$ contains the pattern $p$ as a substring, and
if so where. In the Rabin-Karp algorithm, we compute a given hash function on
both the pattern string $p$ and the $|p|$-character substring starting from the
\$i\$th position of $t$. If these two strings are identical, clearly the
resulting hash values must be the same. If the two strings are different, the
hash values will _almost certainly be different_ (we can check). Note that we
need our hashing function to be constant for this algorithm to be $O(n)$ instead
of $O(mn)$. This may be accomplished with a rolling hash function.

#### Duplicate Detection via Hashing

The key idea of hashing is to represent a large object using a single number.
Hashing can be applied to duplicate detection. Suppose we're looking to find if
a given document is contained in a corpus. Explicitly comparing the new document
$D$ to all $n$ documents is hopelessly inefficient. But we can hash $D$ to an
integer, and compare it to the hash codes of the rest of the corpus.

## Libraries

Java's standard library provides the `HashMap` class. Python has built in hash
tables with the `dict` type. Python also includes `defaultdict` and `Counter` in
the `collections` module.
