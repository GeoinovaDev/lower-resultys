package cmdline

import (
	"os"
	"strconv"
	"strings"
)

// ExistParam ...
func ExistParam(name string) bool {
	token := getToken(name)

	if len(token) == 0 {
		return false
	}

	return true
}

// GetParam ...
func GetParam(name string) string {
	return GetParamString(name)
}

// GetParamString ...
func GetParamString(name string) string {
	return getToken(name)
}

// GetParamInt ...
func GetParamInt(name string) int {
	token := getToken(name)

	if len(token) == 0 {
		return 0
	}

	v, err := strconv.Atoi(token)
	if err != nil {
		return 0
	}

	return v
}

// GetParamFloat ...
func GetParamFloat(name string) float64 {
	token := getToken(name)

	if len(token) == 0 {
		return 0
	}

	v, err := strconv.ParseFloat(token, 10)
	if err != nil {
		return 0
	}

	return v
}

func getToken(name string) string {
	args := extractTokens(os.Args[1:])

	if val, ok := args[name]; ok {
		return val
	}

	return ""
}

func extractTokens(args []string) map[string]string {
	tokens := make(map[string]string)

	for _, arg := range args {
		p := strings.Split(arg, "=")
		if len(p) != 2 {
			continue
		}

		p[0] = strings.Replace(p[0], "-", "", -1)
		tokens[p[0]] = p[1]
	}

	return tokens
}
