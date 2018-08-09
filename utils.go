package websocket

// hasIntersected s1 中的数据是否在 s2 中存在
func hasIntersected(s1, s2 []string) bool {
	if s1 == nil || s2 == nil {
		return false
	}

	for _, i := range s1 {
		for _, j := range s2 {
			if i == j {
				return true
			}
		}
	}

	return false
}
