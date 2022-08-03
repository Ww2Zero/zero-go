package main

func isIsomorphic(s string, t string) bool {
	strList := []byte(s)
	patternByte := []byte(t)
	ms := map[byte]byte{}
	mt := map[byte]byte{}
	if len(strList) != len(patternByte) {
		return false
	}

	for i := 0; i < len(strList); i++ {
		target := patternByte[i]
		dest := strList[i]
		if v, has := ms[dest]; has {
			if v != target {
				return false
			}
		} else {
			if v, has := mt[target]; has {
				if v != dest {
					return false
				}
			}

		}
		mt[target] = dest
		ms[dest] = target
	}

	return true
}
