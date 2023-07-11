# Bitwise Operations

**Bitwise operations** act on the individual bits of their operands. For
example, a bitwise AND (`&`) will perform a logical AND (`&&`) on each pair of
bits. `14 & 9` evaluates to `8`:

```
   1110
&  1001
--------
   1000
```

**Logical operators** treat each operand as having only one value.

## Operators

C has a number of operators that work directly on bits.

| Operator | Name            |
| -------- | --------------- |
| `&`      | Bitwise AND     |
| `│`      | Bitwise OR      |
| `^`      | Bitwise XOR     |
| `~`      | Bitwise NOT     |
| `<<`     | Bit Shift Left  |
| `>>`     | Bit Shift Right |

The first three operators are explained with truth table below.

|     |     | `&` | `│` | `^` |
| --- | --- | --- | --- | --- |
| 0   | 0   | 0   | 0   | 0   |
| 0   | 1   | 0   | 1   | 1   |
| 1   | 0   | 0   | 1   | 1   |
| 1   | 1   | 1   | 1   | 0   |

The right shift (`>>`) can be **logical** or **arithmetic**. Logical right shift
pads with zeros, but arithmetic right shift pads with the most significant bit.

## Examples

| Code         | Description          |
| ------------ | -------------------- |
| `x&(1<<i)`   | Retrieve the ith bit |
| `x│=(1<<i)`  | Set the ith bit      |
| `x&=~(1<<i)` | Clear the ith bit    |
| `x^=(1<<i)`  | Flip the ith bit     |

Bit-level operations often use masking. A **mask** is a bit pattern that
indicates a selected set of bits within a word. For example, the bit-level
operation `x & 0xFF` yields a value consisting of the least significant byte of
x, but with all other bytes set to 0. The expression `~0` will yield a mask of
all ones, regardless of the word size of the machine.

### Sum Integers

```c
int add(int x, int y)
{
  if (y == 0) {
    return x;
  } else {
    int carry = (x & y) << 1;
    return add(x ^ y, carry);
  }
}
```

### Count 1 Bits

```c
int count(int bits)
{
  int total = 0;
  while (bits) {
    total += bits & 1;
    bits >>= 1;
  }
  return total;
}
```

### Erase Lowest Bit

`x&(x - 1)` equals `x` with its lowest set bit erased. For example, if
`x = 00101100`, then `x - 1 = 00101011`, so
`x&(x - 1) = 00101100 & 00101011 = 00101000`.

This fact can be used to reduce the time complexity of counting 1 bits.

```c
int count(int bits)
{
  int total = 0;
  while (bits) {
    total += 1;
    bits &= bits - 1;
  }
  return total;
}
```

### Parity

NOTE: This section draws heavily, often verbatim, from _Elements of Programming
Interviews_ by Aziz, Lee, and Prakash.

The parity of a binary word is 1 if the number of 1s in the word is odd;
otherwise, it is 0. For example, the parity of 1011 is 1, and the parity of
10001000 is 0.

Parity could be computed with brute-force by counting 1 bits (see above) and
modding by 2. XOR provides a more efficient solution. Note that the XOR of a
group of bits is its parity. For example `1^0^1^1` is 1. Since XOR is
associative and commutative, we can process multiple bits at a time.

```c
int parity(int bits)
{
  bits ^= bits >> 32;
  bits ^= bits >> 16;
  bits ^= bits >> 8;
  bits ^= bits >> 4;
  bits ^= bits >> 2;
  bits ^= bits >> 1;
  return bits & 1;  /* extract last bit */
}
```
