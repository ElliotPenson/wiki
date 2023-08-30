# Algorithms

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

A **problem** is specified by describing the complete set of **instances** it
must work on and of its output after running on one of these instances. For
example, an instance of the sorting problem might be `[1, 4, 3]`. An
**algorithm** is a procedure that takes any of the possible input instances and
transforms it to the desired output. There are many different algorithms that
solve the problem of sorting. See [program modeling](program-modeling.org) for
more information.

We seek algorithms that are _correct_, _efficient_, and _easy to implement_.
These goals may not be simultaneously achievable! See
[correctness](correctness.org) and
[[analysis-of-algorithms|analysis of algorithms]] for more information

There is a fundamental difference between algorithms, which always produce a
correct result, and **heuristics**, which may usually do a good job but without
providing any guarantee. For example, the nearest-neighbor heuristic can be
applied to the _traveling salesman problem (TSP)_, but sometimes this heuristic
doesn't even come close to finding the shortest possible tour.
