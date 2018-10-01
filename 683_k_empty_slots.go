package leetcode

func kEmptySlots(flowers []int, k int) int {
	pots := make([]bool, len(flowers))

LOOP:
	for i := range flowers {
		pot := flowers[i] // pot is 1-based
		pots[pot-1] = true

		// Left check
		// 1 0 0 pot
		if pot-k-1 > 0 && pots[pot-k-1-1] {
			for j := pot - k; j < pot; j++ {
				if pots[j-1] {
					continue LOOP
				}
			}
			return i + 1
		}

		// Right check
		// pot 0 0 1
		if pot+k+1 <= len(pots) && pots[pot+k] {
			for j := pot + 1; j < pot+k+1; j++ {
				if pots[j-1] {
					continue LOOP
				}
			}
			return i + 1
		}
	}
	return -1
}
