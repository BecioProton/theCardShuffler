package shufflestyle

// so the overhand takes a number x of cards from the bottom so deckRight[max-rand] to deckRight[max] is taken away
// the new sub array or deckRight is then put back in a hole that opens in a non accurate middle of the deckLeft..
// leftdeck.size()/2 +||- 1~5 cards in that space the leftArray shifts and makes space for rightArray.
// what this means is that a new array is create which has deckleft[0] .. deckLeft[half-1], deckRight[0] .. deckRight[max],deckLeft[half+1] .. deckLeft[max]
