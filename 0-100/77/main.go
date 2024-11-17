package main

func combine(n int, k int) [][]int {
	if n < 1 {
		return nil
	}
	if k < 1 {
		return nil
	}

	prefix := make([]int, 0, k)
	return combineRecursion(prefix, 1, n, k)
}

func combineRecursion(prefix []int, nStart, nEnd, k int) (out [][]int) {
	if k == 0 {
		return [][]int{prefix}
	}
	if nStart == nEnd {
		if k == 1 {
			newPrefix := make([]int, len(prefix), cap(prefix))
			copy(newPrefix, prefix)
			newPrefix = append(newPrefix, nStart)
			return [][]int{
				newPrefix,
			}
		}
		return nil
	}
	if (nEnd-nStart)+1 < k {
		return nil
	}
	if (nEnd-nStart)+1 == k {
		newPrefix := make([]int, len(prefix), cap(prefix))
		copy(newPrefix, prefix)
		for i := nStart; i <= nEnd; i++ {
			newPrefix = append(newPrefix, i)
		}
		return [][]int{
			newPrefix,
		}
	}

	for i := nStart; i <= nEnd; i++ {
		newPrefix := make([]int, len(prefix), cap(prefix))
		copy(newPrefix, prefix)
		newPrefix = append(newPrefix, i)
		subData := combineRecursion(newPrefix, i+1, nEnd, k-1)
		out = append(out, subData...)
	}

	return out
}
