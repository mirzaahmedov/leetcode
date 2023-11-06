package solution

func furthestDistanceFromOrigin(moves string) int {
	position := 0
	blanks := 0

	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case '_':
			blanks += 1
		case 'L':
			position -= 1
		case 'R':
			position += 1
		}
	}

	if position < 0 {
		position = -position
	}

	return position + blanks
}
