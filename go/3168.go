package leetcode

func minimumChairs(s string) int {
	m := map[string]int{
		"person": 0,
		"chairs": 0,
	}

	for _, val := range s {
		if string(val) == "E" {
			if m["chairs"]-m["person"] >= 1 {
				m["person"] += 1
			} else {
				m["person"] += 1
				m["chairs"] += 1
			}
		} else {
			m["person"] -= 1
		}
	}
	return m["chairs"]
}