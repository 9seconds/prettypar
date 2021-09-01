package prettyparlib

type textBlocks []textBlock

func (t textBlocks) Map(mapper func(textBlock) textBlock) textBlocks {
	t = t.makeCopy()

	for i, v := range t {
		t[i] = mapper(v)
	}

	return t
}

func (t textBlocks) Join() textBlock {
	joined := textBlock{}

	for _, block := range t {
		joined = append(joined, block...)
	}

	return joined
}

func (t *textBlocks) makeCopy() textBlocks {
	blocks := make(textBlocks, len(*t))

	copy(blocks, *t)

	return blocks
}
