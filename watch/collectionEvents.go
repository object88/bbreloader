package watch

type collectedEvents struct {
	created map[string]bool
	removed map[string]bool
	renamed map[string]bool
	written map[string]bool
}

func newCollectedEvents() *collectedEvents {
	return &collectedEvents{
		created: make(map[string]bool),
		removed: make(map[string]bool),
		renamed: make(map[string]bool),
		written: make(map[string]bool),
	}
}

func (ce *collectedEvents) HasEvents() bool {
	return len(ce.created) > 0 ||
		len(ce.removed) > 0 ||
		len(ce.renamed) > 0 ||
		len(ce.written) > 0
}
