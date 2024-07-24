## Overflow

If you attempt to create a numeric value that is outside of the range that can be represented with a given number of digits, an **overflow** will occur. Typically, the result will “wrap around” the maximum.

(From now on, we shall omit the obvious `package main` and `import "fmt"` from our examples.)

```go
func add(a, b uint8) uint8 {
	return a + b
}

func main() {
	fmt.Println(add(250, 10))
}
```

```
4
```

Why is the result equal to $4$? Remember that `uint8` can hold values from $0$ to $255$. We are trying to add $10$ to $250$. The resulting $260$ exceeds the capacity of `uint8`. When computing, Go reaches $255$ and wraps back to $0$, hence a result of $4$. This can be shown with the following binary addition. It works like a decimal addition, with carry, only with 2 digits. As you can see, the result has a 9th bit on the left that won't fit in `uint8`. Because the data type can only hold 8 bits, the 9th bit is dropped, resulting in `00000100`` or $4$.

```
    --uint8--
    1111 1010 (250)
+   0000 1010 (5)
= 1 0000 0100 (260)
    0000 0100 (4)
```

(Bits are grouped by 4 for readability.)

It is up to the program author to make provisions to avoid such errors. For instance, by modifying the return type so that it can always hold the result.

> Exercise: modify the program to return a larger integer, and test it.

If the overflow was not anticipated, it will lead to an incorrect result. This can have dire consequences.

Notoriously, on June 4th, 1996, an [Ariane 5 rocket bursted into flames 39 seconds after liftoff](https://en.wikipedia.org/wiki/Cluster_(spacecraft)#Launch_failure)[^overflow-1]. The explosion was caused by a **buffer overflow** when a program tried to stuff a 64-bit number into a 16-bit space. Sounds familiar? It is estimated that the explosion resulted in a loss of US$ 370m. Fortunatelly there was no crew on board.

[^overflow-1]: https://en.wikipedia.org/wiki/Cluster_(spacecraft)#Launch_failure

![Ariane explosion was caused by a buffer overflow](content/basic-data-types/ariane-501-explosion.jpg){ width=192px }

Sometimes, the result is less harmful. In 2014, [a popular video clip caused the Youtube view counter to overflow](https://arstechnica.com/information-technology/2014/12/gangnam-style-overflows-int_max-forces-youtube-to-go-64-bit/)[^overflow-2], forcing Google to switch from 32 to 64 bits integer. A number of views greater than 2 billions had not been anticipated.

[^overflow-2]: https://arstechnica.com/information-technology/2014/12/gangnam-style-overflows-int_max-forces-youtube-to-go-64-bit/

The consequences were far greater in the former case than in the latter. Depending on the context, the programmer may need to be extremelly wary, or they can afford to remain relatively casual while designing and testing their program.

Proving the correctness of a program is beyond the scope of this bool. We can at least reason about a the safety of a datatype (or of a data structure) based on its purpose.

An application dealing with “small” amounts may be confortable with `int32`. The numbers supported by `uint64` seem to go even beyond a banker’s wildest dreams, although today sky-high numbers in the financial world cast a doubt on the safest assumptions. Imagine a program that manipulates cents rather than units in order to avoid dealing with decimal numbers. We would multiply every amount by 100. Or even by 10000 to safely manipulate 4 decimals. And suppose we do computations on a currency like japanese yen at 1 JPY for 0.008 EUR, leading to further multiply values by about 1000.

Say the consolidated result is in the billions of euros, converted to yens, counting in cents. How safe is our initial assumption now?

To safely manipulate large numbers, Go has dedicated implementations for [big numbers](https://golang.org/pkg/math/big/)[^overflow-3]. But they come at the cost of convenience, readability and performance. That is why they are not your go-to solution in all contexts.

[^overflow-3]: https://golang.org/pkg/math/big/
