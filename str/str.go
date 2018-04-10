package str

import (
	"strconv"
	"strings"
)

// Format injeta os parametros args[1]... no parametro args[0]
// Retorna a string formatada
func Format(arr ...string) string {
	if len(arr) == 0 {
		return ""
	}

	if len(arr) == 1 {
		return ""
	}

	formatador := arr[0]
	for i := 1; i < len(arr); i++ {
		index := "{" + strconv.Itoa(i-1) + "}"
		formatador = strings.Replace(formatador, index, arr[i], -1)
	}

	return formatador
}
