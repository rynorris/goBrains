/*
 * Genetics utilities.
 */

package genetics

// Compare two genetic sequences to see if they match.
func CompareSequence(dx, dy *Dna) bool {
	if len(dx.sequence) != len(dy.sequence) {
		return false
	}

	for i := 0; i < len(dx.sequence); i++ {
		if dx.sequence[i].value != dy.sequence[i].value {
			return false
		}
	}

	return true
}
