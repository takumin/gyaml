package helpers

func RemoveDuplicateStrings(l []string) (r []string) {
	k := make(map[string]bool, len(l))
	for _, s := range l {
		if _, ok := k[s]; !ok {
			k[s] = true
			r = append(r, s)
		}
	}
	return r
}
