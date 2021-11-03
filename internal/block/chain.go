package block

type Blocks []Block

func (bs Blocks) LastBlock() (*Block, bool) {
	size := len(bs)

	if size > 0 {
		return &bs[size-1], true
	}

	return nil, false
}

func (bs Blocks) Length() int {
	return len(bs)
}
