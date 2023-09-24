## Blackjack (guided exercise)

This is a guide exercise. We will explain the rules, suggest an algorithm, and you will take care of the implementation.
### Rules

Blackjack is a casino banking game. We will play a simplified version of the game: single player, single deck of cards, no bets, no hole card.

At the table, the dealer faces the player. The card deck is shuffled. The dealer deals 2 cards to the player, face-down. The dealer's hand is also two cards face-down.

The player's goal is to create a card total higher than those of the dealer's hand but not exceeding 21.

On their turn, the player chooses to "hit" (take a card) or "stand" (end their turn and stop without taking a card). Number cards count as their number. Jack,  queen and king count as 10. Aces count as either 1 or 11 according to the player's choice. If the total exceeds 21 points, it busts, and the player or dealer immediately loses. The winner scores one point.

Cards are left aside after each hand. After seven hands, the dealer grabs all the cards and re-shuffle them.

### Play

The player will be a human interacting with the program. The dealer will be played by the computer.

The dealer's decision process is simple: if their total is below the player's, they hit. If it's above, they stand. If it's equal, to make things simple, the dealer stands.

### Design

Make sure you have completed the Hangman chapter before starting this exercise. You first task is to decompose the problem in smaller, manageable chuncks. The main design ideas are identical or very similar.

- Accept player's input when it's their turn.
- Validate it.
- Perform dealer's logic.
- Keep track of the score.
- Detect a bust.
- Detect the end of the hand.
- Compute total.
- Handle aces (1 or 11).
- ...

The game runs indefinitely. It is an endless loop. We can allow the player to enter `'q'` to quit the game at any time. Inside that loop live **nested loops** running for each hand, handling the player's and the dealer's decisions.

Here is a simple suggested algoirthm, where the deck is shuffled before each hand (rather than after 7 hands).

```
scores = 0, 0
while not quit {
    shuffle
    deal
    while player not standing and player not bust {
        ask decision
        if hit {
            add to hand
        }
    }
    while dealer score < player score {
        dealder hit
        add to hand
    }
    decide winner
}
display scores
```

To make our life easier, we will ignore the aces double value in your first version, counting them as 11.

Good luck!

![A table of blackjack](content/classic/blackjack/blackjack.jpg)
