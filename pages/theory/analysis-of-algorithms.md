# Analysis of Algorithms

Information on this page is taken from _The Algorithm Design Manual_ by Steven
S. Skiena.

## RAM Model of Computation

Machine-independent algorithm design depends upon a hypothetical computer called
the **Random Access Machine** or **RAM**. Under this model of computation, we
are confronted with a computer where

- Each _simple_ operation (+, \*, -, =, if, call) takes exactly one time step.
- Loops and subroutines are the composition of many single-step operations.
- Each memory access takes exactly one time step. Further, we have as much
  memory as we need.

Under the RAM model, we measure run time by counting up the number of steps an
algorithm takes on a given problem instance. We consider different time
complexities that define a numerical function, representing time versus problem
size.

- **Worst-case complexity** of the algorithm is the function defined by the
  maximum number of steps taken in any instance of size $n$.
- **Best-case complexity** of the algorithm is the function defined by the
  minimum number of steps taken in any instance of size $n$.
- **Average-case complexity** of the algorithm is the function defined by the
  average number of steps taken in any instance of size $n$.

## Big Oh Notation

**Big Oh** simplifies our analysis by ignoring levels of detail that do not
impact our comparison of algorithms. The formal definitions are as follows:

- $f(n) = O(g(n))$ means $c \cdot g(n)$ is an _upper bound_ on $f(n)$. Thus
  there exists some constant $c$ such that $f(n)$ is always
  $\leq c \cdot
   g(n)$, for large enough $n$.
- $f(n) = \Omega(g(n))$ means $c \cdot g(n)$ is an _lower bound_ on $f(n)$. Thus
  there exists some constant $c$ such that $f(n)$ is always $\geq c \cdot g(n)$,
  for large enough $n$.
- $f(n) = \Theta(g(n))$ means $c_1 \cdot g(n)$ is an upper bound on $f(n)$ and
  $c_2 \cdot g(n)$ is a lower bound on $f(n)$. Thus there exists constants $c_1$
  and $c_2$ such that $f(n) \leq c_1 \cdot g(n)$ and
  $f(n)
   \geq c_2 \cdot g(n)$.

For example, $2n^2 + 100n + 6 = O(n^2)$, because I choose $c = 3$ and
$3n^2
  \geq 2n^2 + 100n + 6$ when $n$ is big enough.

### Big Oh Classes

Big Oh groups functions into a set of classes, such that all the functions in a
particular class are equivalent with respect to the Big Oh. A small variety of
time complexities suffice and account for most algorithms that are widely used
in practice.

| Class Name  | Function        |
| ----------- | --------------- |
| Constant    | $f(n) = 1$      |
| Logarithmic | $f(n) = \log n$ |
| Linear      | $f(n) = n$      |
| Superlinear | $f(n) = n lg n$ |
| Quadratic   | $f(n) = n^2$    |
| Cubic       | $f(n) = n^3$    |
| Exponential | $f(n) = c^n$    |
| Factorial   | $f(n) = n!$     |

We say that a faster-growing function **dominates** a slower-growing one.
Specifically, when $f$ and $g$ belong to different classes (i.e.
$f(n)
   \neq \Theta(g(n))$), we say $g$ dominates $f$ when $f(n) = O(g(n))$,
sometimes written $g >> f$.

### Big Oh Operations

The sum of two functions is governed by the dominant one.

$$
O(f(n)) + O(g(n)) \rightarrow O(max(f(n), g(n)))
$$

Multiplying a function by a constant does not affect its asymptotic behavior.

$$
O(c \cdot f(n)) \rightarrow O(f(n))
$$

When two functions in a product are increasing, both are important.

$$
O(f(n)) * O(g(n)) \rightarrow O(f(n) * g(n))
$$

## Logarithms

A **logarithm** is simply an inverse exponential function. Saying $b^x = y$ is
equivalent to saying that $x = \log_b y$. Exponential functions grow at a
distressingly fast rate. Thus, inverse exponential functions - i.e. logarithms -
grow refreshingly slowly. Logarithms arise in any process where things are
repeatedly halved.

**Binary search** is a good example of an $O(\log n)$ algorithm. If searching
for a particular name $p$ in a telephone book, we start by comparing $p$ against
the middle. Then we discard half the names. Only twenty comparisons suffice to
find any name in the million-name Manhattan phone book!

Logarithms appear in trees (height is $\log_2 n$), bits ($\log_2 n$ bits
required to store a number in binary).

### Logarithm Properties

The $b$ term in $\log_b y$ is the **base** of the logarithm. Three bases are of
importance for mathematical and historical reasons.

- Base $b = 2$: The **binary logarithm**, usually denoted $lg n$, is a base 2
  logarithm. Most algorithm applications of logarithms imply binary logarithms.
- Base $b = e$: The **natural log**, usually denoted $ln x$, is a base \$e =
  2.71828…\$ logarithm.
- Base $b = 10$: Less common today is the base-10 or **common logarithm**,
  usually denoted as $\log x$.

$$
\log_x(xy) = \log_a(x) + \log_a(y)
$$

It is easy to convert a logarithm from one base to another. This is a
consequence of the formula:

$$
\log_a b = \frac{\log_c b}{\log_c a}
$$

Thus, changing the base of $\log b$ from base-a to base-c simply involves
dividing by $\log_c a$.

The base of the logarithm has no real impact on the growth rate. We are usually
justified in ignoring the base of the logarithm when analyzing algorithms.
