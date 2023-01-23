func in(input []byte, char byte) bool {
	q := false

	for i := 0; i < len(input); i++ {
		if input[i] == char {
			q = true
			break
		}
	}
	return q
}

func lengthOfLongestSubstring(s string) int {
	var list []byte
	var i, max, now int

	startind := -1

	for {

		if len([]rune(s)) == 0 {
			return 0
		}

		if in(list, s[i]) == false {
			list = append(list, s[i])
			now++
		} else if in(list, s[i]) == true {
			list = nil
			startind++
			i = startind
			now = 0
		}

		if max < now {
			max = now
		}

		if len([]rune(s))-1 == i {
			break
		}

		i++

		//fmt.Println(i, max, now, list)
		//fmt.Println(s[i], in(list, s[i]))

	}
	return max
}
