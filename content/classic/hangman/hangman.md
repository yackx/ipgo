## Hangman

**Guess a secret word by suggesting letters.**

### The game

In this classic game, one player picks a random word, and the other player attempts to guess it. Each letter of the word to guess is hidden and replaced by a dash. At each turn, the guessing player suggests a letter. If it occurs in the word, the other player writes it at the correct position. If the letter does not occur in the word, the other player draws one limb of a hanged man stick figure on a gallow. The guessing player wins if they find out the word before the hangman is drawn. The other player wins otherwise.

There are many variations of the game ; we will stick to a simple one. In our implementation:

- The computer will pick the word, and the human player will try to guess it.
- We play with uppercase letters only.
- If a letter occuring several times is guessed, it will be revealed at all the places it occurs.
- The program terminates after the word is guessed, or after the stickman is complete.

![Hangman](content/classic/hangman/hangman.jpeg)

> _Note_: Constructing a hanged stickman limb by limb is somewhat unsettling.

### Building blocks

The game is straightforward to implement with the basic programming blocks described so far. There is no complex algorithm or data structure involved. Our first step is to decompose the game in smaller tasks we can eventually assemble:

- Pick a random word from a dictionary
- Ask player for their guess
- Represent the game's state
- Start the game
- Process the guess
- Draw the game state
- Check if game is over

Even though the entire task may be daunting, each individual function is simple enough. By taking a divide and conquer approach, a seemingly difficult task becomes simpler.

### Pick a random word

We will pick a random word from a dictionary stored as a file, *one word per line*. We will read and store the dictionary in memory in order to choose a random entry. This approach is not memory-efficient, but even thousands of words should not cause performance concerns on a modern machine.

> _Note_: There is no such thing as a ``len(file)`` function that would return the number of lines in the file without reading it. This is not a Go specific limitation. From the operating system point of view, a file is made of bytes and there is no concept of "lines" or "words". The linebreak is just a byte (or two bytes), like any other. We cannot know how many "lines" the file contains without reading it entirely.

A different approach would be to first count the number of words in the file (by scanning the file entierely), generate a random number smaller than the number of lines, and then rescan the file to choose the corresponding word. It would avoid loading the whole dictionary in memory, at the cost of two file scans instead of one. So the efficiency gained on memory would be mitigated by the double file scan (worst case, if we pick the last word in the file).

A more effcient but also fairly involved technique would improve the first scanning phase by reading large chuncks of bytes, rather than relying on the slower `scanner.Scan()` and `scanner.Text()` to read lines that will will use.

These optimizations are however beyond the scope of our little program.

Here is a simple implementation to read a dictionary.

```go
const dictionaryPath = "dict.txt"

// Choose a random word from a dictionary
func randomWord(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	word := lines[rand.Intn(len(lines))]
	return strings.ToUpper(word)
}
```

### Ask player's guess

We leniently read from the standard input, ignoring errors. Incorrect input handling and error handling can take a fair amount of work. Don't make a habit of ignoring errors. It is fine in a toy program, but not adequate for a production system.

We then check that the input length is ``1``, that it is a letter, and convert it to upper case.

```go
// Ask the player to enter their guess.
// Single letter only. Will be upper cased
func askGuess() string {
	for {
		fmt.Print("Your guess? ")
		var guess string
		_, _ = fmt.Scan(&guess)
		if len(guess) != 1 {
			continue
		}
		for _, c := range guess {
			if !unicode.IsLetter(c) {
				continue
			}
		}
		return strings.ToUpper(guess)
	}
}
```

Notice the endless ``for`` loop. In structured programming, one would declare a stop condition on the iteration. Arguably the extra variable and nesting in conditional statements would not improve readability. Go favors explicit exits via ``break`` (or in this case ``continue``). The function ``askGuess`` is short, so there is no risk of readability issues.

### Game state

A ``struct`` holds the game state. It is better than having free variables "floating" around.

```go
type Hangman struct {
	secret  string		// Secret word to guess
	word    string		// Current player's word.
						// Letters not found yet will be '-'
	guesses []string	// Player's guesses
	limb    Limb		// Current limb
}
```

The `Limb` will be an enumeration. In Go, enums are composed of:

1. A custom type declaration
2. Several constants with ``iota``
3. A string representation (for humans)

```go
type Limb int

const (
	Empty = iota
	Head
	Torso
	LeftArm
	RightArm
	LeftLeg
	RightLeg
)

func (limb Limb) String() string {
	return [...]string{
		"Empty", "Head", "Torso", "Left Arm",
		"Right Arm", "Left Leg", "Right Leg"}[limb]
}
```

We have included a fake limb called `Empty` to treat the starting point of the game like the rest of the game.

### Start the game

When starting the game, a `Hangman` is created by performing the following actions:

- Pick a random `secret` word, using the function we declared above
- Initialize the `guesses` to an empty slice
- Set the `word` to dashes
- Set the current `limb` to `Empty`

```go
// Instantiate a new game
func newHangman() *Hangman {
	return &Hangman{
		guesses: make([]string, 0),
		secret:  randomWord(dictionaryPath),
		word:    strings.Repeat("-", len(secret)),
		limb:    Empty}
}
```

The pseudo-random generator must also be intitialized. It has to be done once in the program, no matter how many games are played. Initiliazation has been therefore kept out of `newHangman()`.

```go
func initGame() {
	rand.Seed(time.Now().UnixNano())
}
```

### Process player's guess

The bulk of the game's logic is to process the player's guess.

If the `letter` occurs in the `secret`, a loop will replace the dash(es) to reveal it. Notice how we leverage *slicing* to construct a new `word` in 3 parts: 1. everything before the letter's position `[:i]`, 2. the `letter` itself and 3. what is left after the letter's position `[i+1:]`.

If the `letter` does not occur, we simply increase to the next `limb`.

At this stage, no check is performed to find out if one of the players won, or if the game continues.

```go
// Attempt to guess a letter.
// If the letter exists in the secret, it is revealed in the word.
// Otherwise, limb is incremented.
func (h *Hangman) guess(letter string) {
	h.guesses = append(h.guesses, letter)
	found := false
	for i, c := range h.secret {
		s := fmt.Sprintf("%c", c)
		if s == letter {
			h.word = h.word[:i] + letter + h.word[i+1:] // reveal
			found = true
		}
	}

	if !found {
		h.limb++
	}
}
```

### Display the game state

The reader is free to be creative for the text interface and the hangman representation. We will opt for the utmost sobriety.

```go
func (h *Hangman) draw() {
	fmt.Printf("%s %v %s\n", h.word, h.guesses, h.limb)
}
```

Admitedly, calling this function `draw()` may be an overstatement since it merely displays or prints the game state in a raw form.

```
----E----E [A E P] Torso
```

> *Note:* Displaying a gallow and an actual stickman will require more lines of code on their own than the complete implementation presented here. There is nothing wrong with that fact; we simply prefer to keep things simple.

### Check game over

The guessing player wins if the `word` matches the `secret`. They lose if the `limb` is the last one. Give them the courtesy to reveal the secret in that case. If neither `won()` or `lost()` are `true`, the game goes on!

```go
// Check if player won
func (h *Hangman) won() bool {
	return h.word == h.secret
}

// Check if player lost
func (h *Hangman) lost() bool {
	return h.limb == RightLeg
}

// Display game result
func (h *Hangman) displayResult() {
	if h.won() {
		fmt.Println("You won")
	} else if h.lost() {
		fmt.Printf("You lost. Secret word was: %s\n", h.secret)
	}
}
```

### Main

Finally we assemble our building blocks in neat `main` function.

```go
func main() {
	initGame()
	hangman := newHangman()
	hangman.draw()

	for !hangman.won() && !hangman.lost() {
		letter := askGuess()
		hangman.guess(letter)
		hangman.draw()
	}

	hangman.displayResult()
}
```

### Sample

```
---------- [] -
Your guess? a
---------- [A] Head
Your guess? e
----E----E [A E] Head
Your guess? p
----E----E [A E P] Torso
Your guess? d
----ED---E [A E P D] Torso
Your guess? i
I---EDI--E [A E P D I] Torso
Your guess? n
IN--EDI--E [A E P D I N] Torso
Your guess? c
INC-EDI--E [A E P D I N C] Torso
Your guess? r
INCREDI--E [A E P D I N C R] Torso
Your guess? b
INCREDIB-E [A E P D I N C R B] Torso
Your guess? l
INCREDIBLE [A E P D I N C R B L] Torso
You won
```

### Dictionary

You can easily find and download dictionaries from the internet, in English or in any other language. A short dictionary is however useful for testing purpose, like this one.

```
panda
cactus
dog
cat
incredible
```
