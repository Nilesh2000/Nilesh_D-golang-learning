package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func Dice() int {
	return rand.Intn(6) + 1
}

func Game() (int, int) {
	player1Strategy, _ := strconv.ParseInt(os.Args[1], 10, 64)
	player2Strategy, _ := strconv.ParseInt(os.Args[2], 10, 64)
	return int(player1Strategy), int(player2Strategy)

}

func Play(player1Strategy, player2Strategy int) string {
	score1, score2 := 0, 0

	for turn := 1; score1 <= 100 || score2 <= 100; {
		fmt.Println("Turn", turn)

		fmt.Println("Player 1:-")
		turnScore1 := 0
		for turnScore1 < player1Strategy {
			diceRoll := Dice()
			fmt.Printf("Dice:%d ", diceRoll)
			if diceRoll == 1 {
				turnScore1 = 0
				break
			}
			turnScore1 += diceRoll
			if score1+turnScore1 >= 100 {
				break
			}
		}
		fmt.Printf("=> Total Turn Score:%d", turnScore1)
		score1 += turnScore1

		if score1 >= 100 {
			fmt.Println("\nPlayer 1 Score:", score1)
			return "Player 1"
		}

		// -----------------------------------------------------------------
		fmt.Println("\nPlayer 2:-")
		turnScore2 := 0
		for turnScore2 < player2Strategy {
			diceRoll := Dice()
			fmt.Printf("Dice:%d ", diceRoll)
			if diceRoll == 1 {
				turnScore2 = 0
				break
			}
			turnScore2 += diceRoll
			if score2+turnScore2 >= 100 {
				break
			}
		}
		fmt.Printf("=> Total Turn Score:%d", turnScore2)
		score2 += turnScore2

		if score2 >= 100 {
			fmt.Println("\nPlayer 2 Score:", score2)
			return "Player 2"
		}

		fmt.Printf("\nScores at End of Turn %d: Player 1=%d, Player 2=%d\n", turn, score1, score2)
		fmt.Println()

		turn++
	}
	return "Player 2"
}

func main() {
	s1, s2 := Game()
	fmt.Println("Player 1 Strategy:", s1, ", Player 2 Strategy:", s2)
	fmt.Println()

	count1, count2 := 0, 0
	for i := 1; i <= 10; i++ {
		winner := Play(s1, s2)
		if winner == "Player 1" {
			count1++
		} else {
			count2++
		}
	}
	fmt.Printf("Holding at %d vs Holding at %d: wins (%d/10) (%.1f%%), losses: (%d/10) (%.1f%%)\n", s1, s2, count1, float32(count1)*10, count2, float32(count2)*10)
}
