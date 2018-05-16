package datetime

import (
	"strconv"
	"strings"
	"time"

	"git.resultys.com.br/lib/lower/str"
)

// Datetime struct
type Datetime struct {
	time time.Time
}

// Now ...
func Now() *Datetime {
	return &Datetime{time: time.Now()}
}

// Parse ...
func Parse(datetime string) *Datetime {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)

	return &Datetime{time: t}
}

// String ...
func (d *Datetime) String() string {
	tokens := d.extractTokens()

	return str.Format("{0}-{1}-{2} {3}:{4}:{5}",
		formatNumber(tokens["y"]),
		formatNumber(tokens["M"]),
		formatNumber(tokens["d"]),
		formatNumber(tokens["h"]),
		formatNumber(tokens["m"]),
		formatNumber(tokens["s"]),
	)
}

// TotalDias ...
func (d *Datetime) TotalDias(n *Datetime) int {
	diff := d.time.Sub(n.time)
	t := int(diff.Hours() / 24)

	if t < 0 {
		return t * -1
	}

	return t
}

func formatNumber(n int) string {
	s := strconv.Itoa(n)

	if n < 10 {
		return "0" + s
	}

	return s
}

func (d *Datetime) extractTokens() map[string]int {
	tokens := map[string]int{}

	p := strings.Split(d.time.String(), " ")
	if len(p) < 2 {
		return tokens
	}

	data := strings.Split(p[0], "-")
	tokens["y"], _ = strconv.Atoi(data[0])
	tokens["M"], _ = strconv.Atoi(data[1])
	tokens["d"], _ = strconv.Atoi(data[2])

	hora := strings.Split(p[1], ".")
	hora = strings.Split(hora[0], ":")
	tokens["h"], _ = strconv.Atoi(hora[0])
	tokens["m"], _ = strconv.Atoi(hora[1])
	tokens["s"], _ = strconv.Atoi(hora[2])

	return tokens
}
