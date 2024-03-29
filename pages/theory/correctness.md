# Correctness

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

Reasonable-looking [[algorithms]] can easily be incorrect. Algorithm correctness
is a property that must be carefully demonstrated. The primary tool to
distinguish correct algorithms from incorrect ones is a **proof**. A proof has
four main parts.

1. Clear, precise statement of what you are trying to prove.
2. Set of assumptions.
3. Chain of reasoning that takes you from the assumptions to the statement you
   are trying to prove.
4. Little, black square or _QED_.

Computer scientists usually prove things with **proof by induction** or **proof
by contradiction**. We'll see induction below.

To reason about an algorithm, you need a careful description of the sequence of
steps to be performed. The three most common forms of algorithmic notation are
English, pseudocode, or a real programming language. These methods have natural
tradeoffs between expression and precision.

## Incorrectness

The best way to prove that an algorithm is _incorrect_ is to produce an instance
in which it yields an incorrect answer. Such instances are called
**counter-examples**. Good counter-examples are verifiable and simple. Many
tricks exist for finding counter-examples.

- _Think small_: small examples are easier to reason about.
- _Think exhaustively_: consider different types of examples.
- _Go for a tie_: try similar values in the input collection.
- _Seek extremes_: e.g. use values that are far apart or close together.

## Induction

**Mathematical induction** is a common choice for proving correctness. If you're
familiar with recursion, you're familiar with induction; recursion _is_
induction. In both, we have general and boundary conditions. The general
condition breaks the problem into smaller and smaller pieces and the boundary
condition terminates the recursion. Suppose that you are trying to prove that a
statement holds true for all natural numbers (all $n$). Induction would usually
take the form:

1. Show that the statement holds for the base case (usually $P(0)$ or $P(1)$).
2. Assume that $P(k)$ is true.
3. Prove that the statement also holds for $P(k + 1)$.
