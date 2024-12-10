package main

func canChange(start string, target string) bool {
	if len(start) != len(target) {
		return false
	}

	stData := make([]data, 0, len(start))
	ttData := make([]data, 0, len(start))

	for i := range start {
		if start[i] == '_' {
			continue
		}
		if start[i] == 'L' {
			stData = append(stData, data{
				isL: true,
				pos: i,
			})
		} else {
			stData = append(stData, data{
				isL: false,
				pos: i,
			})
		}
	}

	for i := range target {
		if target[i] == '_' {
			continue
		}
		if target[i] == 'L' {
			ttData = append(ttData, data{
				isL: true,
				pos: i,
			})
		} else {
			ttData = append(ttData, data{
				isL: false,
				pos: i,
			})
		}
	}

	if len(stData) != len(ttData) {
		return false
	}

	for i := range stData {
		if stData[i].isL != ttData[i].isL {
			return false
		}
		if stData[i].pos == ttData[i].pos {
			continue
		}
		if stData[i].pos < ttData[i].pos {
			if stData[i].isL {
				return false
			} else {
				continue
			}
		}
		if stData[i].pos > ttData[i].pos {
			if stData[i].isL {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

type data struct {
	isL bool
	pos int
}
