package leetcode

import (
	"fmt"
	"strconv"
)

func nextClosestTime(time string) string {
	digitSet := map[int]bool{}
	for i := range time {
		if i == 2 {
			continue
		}
		digitSet[int(time[i]-'0')] = true
	}

	startHH, _ := strconv.Atoi(time[:2])
	startMM, _ := strconv.Atoi(time[3:])
	startMinutes := startHH*60 + startMM
	endMinutes := startMinutes

LOOP:
	for {
		startMinutes++
		if startMinutes == 1440 {
			startMinutes = 0
		}

		digits := []int{
			startMinutes / 60 / 10,
			startMinutes / 60 % 10,
			startMinutes % 60 / 10,
			startMinutes % 60 % 10,
		}
		for _, digit := range digits {
			if !digitSet[digit] {
				continue LOOP
			}
		}
		endMinutes = startMinutes
		break
	}
	return fmt.Sprintf("%02d:%02d", endMinutes/60, endMinutes%60)
}
