package leetcode

func licenseKeyFormatting(S string, K int) string {
	length := len(S)
	bytes := make([]byte, 0, length)

	numItems := 0
	for i := length - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}

		if (S[i] >= '0' && S[i] <= '9') ||
			(S[i] >= 'A' && S[i] <= 'Z') {
			bytes = append(bytes, S[i])
		} else {
			bytes = append(bytes, S[i]-32)
		}
		numItems++

		if numItems == K {
			numItems = 0
			bytes = append(bytes, '-')
		}
	}

	length = len(bytes)
	if length == 0 {
		return ""
	}

	if bytes[length-1] == '-' {
		bytes = bytes[:length-1]
		length--
	}

	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
	return string(bytes)
}
