package leetcode

func lengthOfLongestSubstring(s string) int {
	set := map[byte]bool{}
	start := 0
	end := 0
	longest := 0
	for start < len(s) && end < len(s) {
		if !set[s[end]] {
			set[s[end]] = true
			end++
			if end-start > longest {
				longest = end - start
			}
			continue
		}

		delete(set, s[start])
		start++
	}
	return longest
}
