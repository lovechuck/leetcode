package src

/*
You are playing the following Flip Game with your friend: Given a string that contains only these two characters: + and -, you and your friend take turns to flip twoconsecutive "++" into "--". The game ends when a person can no longer make a move and therefore the other person will be the winner.

Write a function to compute all possible states of the string after one valid move.

For example, given s = "++++", after one move, it may become one of the following states:

[
  "--++",
  "+--+",
  "++--"
]


If there is no valid move, return an empty list [].
*/

func generatePossibleNextMoves(s string) []string {
	var res []string
	for i := 1; i < len(s); i++ {
		if s[i] == '+' && s[i-1] == '+' {
			temp := s
			pre := s[:i]
			suffix := s[i+1:]
			temp = pre + "__" + suffix
			res = append(res, temp)
		}
	}

	return res
}
