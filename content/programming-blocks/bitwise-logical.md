## Bitwise logical operators

Bitwise logical operators perform logical operations on individual bits of binary numbers. These operators work at the binary level, manipulating individual bits. In binary, each bit can have one of two possible values, represented as `1` and `0`, or `I` and `O`, or `true` or `false`.

Before we work with numbers, let’s first work on single bits. Let’s have two of them: `p` and `q`. The operators are:

--- --------------------
&    bitwise AND
|    bitwise OR
^    bitwise XOR
&^   bit clear (AND NOT)
--- --------------------

### Tables of truth

We can create a _table of truth_ that covers all combinations of `p` and `q`.

|  p  |  q  | p & q | p \| q | p ^ q | p &^ q |
|-----|-----|-------|--------|-------|--------|
|  0  |  0  |   0   |  0     |   0   |   0    |
|  0  |  1  |   0   |  1     |   1   |   0    |
|  1  |  0  |   0   |  1     |   1   |   1    |
|  1  |  1  |   1   |  1     |   0   |   0    |

Now that you understand how binary operators work on single bits, it is easy to see how they operate on integers once you represent the integer in their binary form.

### Binary numbers

Given 12 and 10 in decimal, represented in binary, one above the other:

```
1 1 0 0 (12)
1 0 1 0 (10)
```

We can perform the bitwise operations on each pair of bits taken vertically. We will say that two bits “have the same place” if they have the same significance or weight, that is, if they both are on the same vertical imaginary line.

Bitwise AND `&` will look for bits with the same significance (same place) only to keep bits that are “1” in both integers:

```
  1 1 0 0
& 1 0 1 0
= 1 0 0 0
```

The binary result `1000` is `8` in decimal.

Bitwise OR will keep bits that are “1” *at the same place* in either integer, or in both:

```
  1 1 0 0
| 1 0 1 0
= 1 1 1 0
```

The binary result `1110` is `14` in decimal.

Bitwise XOR will keep bits that are “1” *at the same place* in either integers, but not in both:

```
  1 1 0 0
^ 1 0 1 0
= 0 1 1 0
```

Bitwise AND NOT will keep bits that are “1” for the first bit and “0” for the second:

```
   1 1 0 0
&^ 1 0 1 0
=  0 1 0 0
```

### Shift

```
<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

Shifting is simply a matter of moving bits to the left (`<<`) or to the right (`>>`) a number of times.

`5<<2` means:

* Take the binary representation of `5`: `101`
* Move the bits to the left, two times, filling with `0` on the right as you go: `10100`
* The result is `10100` in binary or `20` in decimal.

`5>>2` means:

* Take the binary representation of `5`: `101`
* Move the bits to the right, two times, dropping the bit on the right each time: `10` then `1`
* The result is `1` in binary or `1` in decimal.

Attempting to shift a negative number of times results in an error.

An interresting property of shifting `<<1` is that it multiplies the original value by 2. Conversely, `>>1` divides by 2.

### Exercises

Calculate the following:

- `128^255`
- `128>>4`
- `64|32`
