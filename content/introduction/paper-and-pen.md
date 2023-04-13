## A paper and a pen

… is all you need to write a program. For instance, let’s play hide and seek. By that, we mean to write *pseudo-code* that describes a game of hide and seek:

```
close eyes
count from 20 to 0
while buddy not found:
    seek buddy
```

Pseudo-code

:   is an informal high-level description of the operating principle of a computer program or other algorithm. It uses the structural conventions of a normal programming language, but is intended for human reading rather than machine reading.[^paper-and-pen-1]

[^paper-and-pen-1]: https://en.wikipedia.org/wiki/Pseudocode

You can read this program almost literally:

* Close your eyes
* Count from 20 to 0
* Search for your hidden buddy
* The game ends when you find them

The `while` section is a **loop**. It means to keep looking _while_ƒski you haven’t found your buddy.

It is helpful to describe an algorithm coarsly. But we can take the same algorithm and add more details, for instance when we countdown:

```
counter <- 20
while counter is not 0:
    counter <- counter - 1
```

Here we give more details on how to count down. Instead of just saying “countdown” and leaving the details to the reader, we explicitely say that we need to decrement our “counter” at each step. Notice that we still lack some details, like for instance the delay between each step. Is it a second, should we count as fast as we can? It is not described, but we can be more specific:

```
counter <- 20
while counter is not 0:
    counter <- counter - 1
    wait 1 second
```

Same algorithm, different levels of details. From informal to more formal. The first version is lightweight and easy to read. It leaves out the details. The thrid version leaves less room to interpretation. The second is in-between.

Both have their use. In all cases, we have left many step aside. For instance, think about what would happen if your buddy was too well hidden. As it is described, the game would go on forever as you would keep searching for them endlessly. A compiled program will gleefully execute the sequence of instructions you entered, regardless of their implications, provided that they are syntactically correct, and no matter the aberration they can produce.

Paper and pen are great tools to learn programming, but may not be very motivating. A program run on a computer gives you more feedback, interaction and gratification than paper. This being said, pseudo-code on paper or on a drawing board[^paper-or-tablet] remains a key tool before jumping to your keyboard.

[^paper-or-tablet]: Of course you can use your tablet and wireless pencil.

![Start with a paper and a pen](content/introduction/paper-and-pen.png)
