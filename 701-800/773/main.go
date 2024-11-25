package main

type data struct {
	mp              map[int]interface{}
	boardsToProcess []*boarder
	iteration       int
}

func slidingPuzzle(board [][]int) int {
	inputBoarder := boarder(board)

	storage := data{
		mp: map[int]interface{}{
			inputBoarder.getHashNum(): nil,
		},
		boardsToProcess: []*boarder{&inputBoarder},
		iteration:       0,
	}

	if inputBoarder.isValid() {
		return storage.iteration
	}
	return storage.storageProcessor()
}

func (s *data) storageProcessor() int {
	for {
		s.iteration++

		if len(s.boardsToProcess) == 0 {
			return -1
		}

		newBoards := make([]*boarder, 0)

		for i := range s.boardsToProcess {
			left, valid := s.boardsToProcess[i].getLeft()
			if valid {
				if left.isValid() {
					return s.iteration
				}
				_, ok := s.mp[left.getHashNum()]
				if !ok {
					s.mp[left.getHashNum()] = nil
					newBoards = append(newBoards, left)
				}
			}

			right, valid := s.boardsToProcess[i].getRight()
			if valid {
				if right.isValid() {
					return s.iteration
				}
				_, ok := s.mp[right.getHashNum()]
				if !ok {
					s.mp[right.getHashNum()] = nil
					newBoards = append(newBoards, right)
				}
			}

			up, valid := s.boardsToProcess[i].getUp()
			if valid {
				if up.isValid() {
					return s.iteration
				}
				_, ok := s.mp[up.getHashNum()]
				if !ok {
					s.mp[up.getHashNum()] = nil
					newBoards = append(newBoards, up)
				}
			}

			down, valid := s.boardsToProcess[i].getDown()
			if valid {
				if down.isValid() {
					return s.iteration
				}
				_, ok := s.mp[down.getHashNum()]
				if !ok {
					s.mp[down.getHashNum()] = nil
					newBoards = append(newBoards, down)
				}
			}
		}
		s.boardsToProcess = newBoards
	}
}

type boarder [][]int

func (b *boarder) isValid() bool {
	if b.getHashNum() == 123450 {
		return true
	}
	return false
}

func (b *boarder) getHashNum() (out int) {
	multiplyNum := 100000
	for i := 0; i < 6; i++ {
		out += b.getVal(i) * multiplyNum
		multiplyNum /= 10
	}
	return out
}

func (b *boarder) getVal(i int) int {
	if i >= 0 && i <= 2 {
		return (*b)[0][i]
	}
	return (*b)[1][i-3]
}

func (b *boarder) getZeroIdx() (pos1, pos2 int) {
	for i := range *b {
		for j := range (*b)[i] {
			if (*b)[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b *boarder) swap(pos1, pos2 []int) {
	pos1Val := (*b)[pos1[0]][pos1[1]]
	(*b)[pos1[0]][pos1[1]] = (*b)[pos2[0]][pos2[1]]
	(*b)[pos2[0]][pos2[1]] = pos1Val
}

func (b *boarder) copy() *boarder {
	out := make(boarder, len(*b))
	out[0] = make([]int, len((*b)[0]))
	out[1] = make([]int, len((*b)[1]))
	for i := range *b {
		for j := range (*b)[i] {
			out[i][j] = (*b)[i][j]
		}
	}
	return &out
}

func (b *boarder) getLeft() (out *boarder, valid bool) {
	pos1, pos2 := b.getZeroIdx()
	if pos2 == 0 {
		return nil, false
	}
	out = b.copy()
	out.swap([]int{pos1, pos2}, []int{pos1, pos2 - 1})
	return out, true
}

func (b *boarder) getRight() (out *boarder, valid bool) {
	pos1, pos2 := b.getZeroIdx()
	if pos2 == 2 {
		return nil, false
	}
	out = b.copy()
	out.swap([]int{pos1, pos2}, []int{pos1, pos2 + 1})
	return out, true
}

func (b *boarder) getUp() (out *boarder, valid bool) {
	pos1, pos2 := b.getZeroIdx()
	if pos1 == 0 {
		return nil, false
	}
	out = b.copy()
	out.swap([]int{pos1, pos2}, []int{pos1 - 1, pos2})
	return out, true
}

func (b *boarder) getDown() (out *boarder, valid bool) {
	pos1, pos2 := b.getZeroIdx()
	if pos1 == 1 {
		return nil, false
	}
	out = b.copy()
	out.swap([]int{pos1, pos2}, []int{pos1 + 1, pos2})
	return out, true
}
