package actions

type nameList struct {
	names map[string]bool
}

func newNameList(names []string) nameList {
	nl := nameList{names: make(map[string]bool)}
	for _, n := range names {
		nl.names[n] = false
	}
	return nl
}

// GetNames gets the first num names from available randomized star names
func (nl nameList) GetNames(num int) []string {
	min := num
	if len(nl.names) < num {
		min = len(nl.names)
	}
	names := make([]string, 0, min)
	for n := range nl.names {
		names = append(names, n)
		if len(names) >= min {
			break
		}
	}

	return names
}
