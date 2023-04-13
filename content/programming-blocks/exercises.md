# Exercices

## Part 1

The following summary exercies combine a bit of everything we have seen so far: functions, variables, expressions, conditional statements and loops. To avoid recompiling your program over and over, some data will be introduced on the keyboard.


1. Write a function that computes the division of two numbers. Call this function from a program that displays the result with 4 digit precision.


2. Write a function that computes the surface of a rectangle, given its width and height. Write a second function that computes its permiter. Write a program that reads a rectangle dimensions and prints its surface and perimeter.


3. Write a program that reads two numbers and displays their sum and the product. The output must be to the 4th decimal. The program terminates when the user enters `0` as the first number.


4. Write a function that takes two numbers as arguments and checks that one and only one of the following conditions is met: (i) one of the two numbers is divisible by the other; (ii) their product is greater than 10 times their sum.


5. Write a program that reads 3 numbers representing the dimensions of a triangle and that checks if: (i) this triangle is possible; (ii) it is equilateral; (iii) it is isosceles.


1. Write a function that takes 3 numeric arguments and checks exactly two of them are equal.


2. Write a function that takes 4 integer arguments and checks that 2 of them are equal, a third (any of them) being strictly greater than the two equal, and a fourth being strictly smaller than the two equal.

```
4  6  6  9    true
5  4  5  7    true
8  4  1  4    true
1  2  4  5    false
3  3  4  4    false
3  3  3  4    false
2  7  7  4    false
9  7  7  8    false
```

## Part 2

The problems in this section are of equal difficulty than the previous one. They are however a bit more involved. A bit more fun too. You already have all the building blocks at your disposal.


1. Read integer numbers from the keyboard. We stop when the user enters `0`. The program displays the sum of positive numbers, and the sum of negative numbers.


2. Each student has taken several tests. Write a program that reads their scores (the input stops when a negative number is entered), and displays: (i) the average score (ii) the highest score (iii) the lowest score (iv) the average of all scores after you have eliminated the highest and the lowest score.


3. Write a function that takes an integer `n` as a parameter, and prints successively the following patterns. The examples shown below are given for `n` equal to 4. You must work line by line, as a printer would.

```
a) ****    b) *     c) ****   d) ****
   ****       **       ***        ***
   ****       ***      **          **
   ****       ****     *            *
```


1. Write a function that takes two integers `n` and `lines` as parameters, and prints the following triangles:

For n=7 and lines=4

```
      7
    8 9 8
  9 0 1 0 9
0 1 2 3 2 1 0
```

For n=2 and lines=5

```
        2
      3 4 3
    4 5 6 5 4
  5 6 7 8 7 6 5
6 7 8 9 0 9 8 7 6
```


1. Write a program that reads a positive integer from the keyboard, and displays the a triangle (example given if the input is 4):

```
   *
  * *
 *   *
*******
```


1. Write a program that reads a sequence of integers from the keyboard, and that computes the number of raises and decreases. The input stop when `0` is entered. Example: `4 6 8 2 -1 -5 0` will count 2 raises and 3 decreases. `4 6 4 8 9 0` will yield 3 and 1 respectively.

## Multiplication table training


1. Write a program that chooses randomly two integers from 1 to 10 included, and ask the user to find their product. If the user makes a mistake, the program should display an error message and ask the same problem again, until the right answer is provided (at that time the program will terminate).


2. Modify the program so that it proposes a new multiplication after the previous one is correctly answered.


3. Extend the program to give at most 3 attempts for a given question. If the user has not found the right answer after 3 attempts, the program gives it out and proceeds with the next question.


4. Make the program stop when the user gives a total of 5 correct answers. Do not forget to congratulate them on their progress.


5. Same as above, but an answer is valid only if given on the first try.


6. Same as above, but this time the user has to provide 5 correct answers in a row for the program to terminate.
