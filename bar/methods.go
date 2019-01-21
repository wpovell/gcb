package bar

// Return block associated with `x` coordinate
// Returns nil if no block associated
func (bar *Bar) findBlock(x int) Block {
	for b, state := range bar.blocks {
		if state.Contains(x) {
			return b
		}
	}

	return nil
}

// Add block to position on bar
func (bar *Bar) AddBlock(aln Align, blk Block, name string) {
	bar.blocks[blk] = nil
	bar.align[aln] = append(bar.align[aln], blk)
	bar.Names[name] = blk
}
