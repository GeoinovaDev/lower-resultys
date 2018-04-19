package array

// Concat concatena todos os arrays no primeiro parametro
func Concat(arr1 []string, arrs ...[]string) []string {
	for i := 0; i < len(arrs); i++ {
		arr2 := arrs[i]
		for _, v2 := range arr2 {
			arr1 = append(arr1, v2)
		}
	}

	return arr1
}

// Unique remove elementos duplicados
func Unique(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}
