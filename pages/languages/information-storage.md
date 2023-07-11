# Information Storage

Rather than accessing individual bits in memory, most computers use blocks of
eight bits, or **bytes**. A machine-level program views memory as a very large
array of bytes, referred to as **virtual memory**. Every byte of memory is
identified by a unique number, known as its **address**. **Virtual address
space** is the set of all possible addresses.

Every computer has a **word size**. A word size indicates the normal size of
integers and pointer data. Since a virtual address is encoded by such a word,
word size determines the maximum size of the virtual address space. That is, for
a machine with a w-bit word size, virtual addresses can range from $0$ to
$2^{(w-1)}$, giving the program access to at most $2^w$ bytes. Thus, 32-bit
computers have a limited virtual address space of 4GB.

We generally write bit patterns as base-16, or **hexadecimal**. Hexadecimal
(hex) uses digits 0 through 9 along with characters A through F to represent 16
possible values. In C, number constants starting with `0x` are interpreted as
hexadecimal.

## Byte Ordering

In virtually all machines, a multi-byte object is stored as a contiguous
sequence of bytes, with the address of the object given by the smallest address
of the bytes used. Two common conventions exist for ordering the bytes of an
object. **Little endian** puts the least significant byte first. **Big endian**
puts the most significant byte first.

## Integer Storage

Integers may be **unsigned** or **signed**. Unsigned integers are nonnegative
numbers. Signed integers may be negative, zero, or positive numbers.

### Unsigned

Unsigned binary representation encodes every number between $0$ and
$2^{(w - 1)}$ with a unique w-bit value. …0000 is 0, …0001 is 1, etc. In C,
unsigned literals have a `u` suffix.

### Signed

Signed integers are most commonly represented with the **two's complement**
encoding. The two's complement of an N-bit number is defined as its complement
with respect to $2^N$. For example, the three-bit number 010 has the two's
complement 110 because $010 + 110 = 1000$. The two's complement is calculated by
inverting the digits and adding one.

In the two's complement _encoding_, if the binary number 010 encodes the decimal
2, then its two's complement, 110, encodes the inverse -2. In other words, to
reverse the sign of any integer in this scheme, you can take the two's
complement of its binary representation.

Two's complement encoding can also be understood as interpreting the most
significant bit of a word as a negative weight. For example, $1011$ will give
$-8 + 2 + 1 = -5$.

In the two's complement encoding, The least representable value is $-2^{(w-1)}$
and the greatest value is $2^{(w-1)} - 1$ (the $-1$ accounts for $0$).
