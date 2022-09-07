package _592_reorderSpaces

import "strings"

func reorderSpaces(text string) string {
	sz := len(text)
	cnt := 0
	strList := make([]string, 0)
	for i := 0; i < sz; {
		if text[i] == ' ' {
			cnt++
			i++
			continue
		}
		j := i
		for j < sz && text[j] != ' ' {
			j++
		}
		strList = append(strList, text[i:j])
		i = j
	}

	var sb strings.Builder
	m := len(strList)

	t := cnt / max(m-1, 1)

	space := ""
	for t > 0 {
		space = space + " "
		t--
	}
	for i := 0; i < m; i++ {
		sb.WriteString(strList[i])
		if i != m-1 {
			sb.WriteString(space)
		}
	}
	for sb.Len() != sz {
		sb.WriteString(" ")
	}

	return sb.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
