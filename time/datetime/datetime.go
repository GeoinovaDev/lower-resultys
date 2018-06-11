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

// New ...
func New(date string, format string) (d *Datetime) {
	defer func() {
		err := recover()
		if err != nil {
			d = Parse("0001-01-01 00:00:00")
		}
	}()

	d = createDateFromFormat(date, format)

	return
}

// IsEqualDate ...
func (d *Datetime) IsEqualDate(date *Datetime) bool {
	if date.time.Day() == d.time.Day() && date.time.Month() == d.time.Month() && date.time.Year() == d.time.Year() {
		return true
	}

	return false
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

// TotalHoras ...
func (d *Datetime) TotalHoras(dt *Datetime) int {
	diff := d.time.Sub(dt.time)
	t := int(diff.Hours())

	if t < 0 {
		return t * -1
	}

	return t
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

func createDateFromFormat(date string, format string) *Datetime {
	y := extractNumber(date, "y", format)
	M := extractNumber(date, "M", format)
	d := extractNumber(date, "d", format)
	h := extractNumber(date, "h", format)
	m := extractNumber(date, "m", format)
	s := extractNumber(date, "s", format)

	if len(h) == 0 {
		h = "00"
	}

	if len(m) == 0 {
		m = "00"
	}

	if len(s) == 0 {
		s = "00"
	}

	return Parse(str.Format("{0}-{1}-{2} {3}:{4}:{5}", y, M, d, h, m, s))
}

func extractNumber(date string, token string, format string) string {
	t := len(strings.Split(format, token)) - 1
	x := -1
	n := []string{}

	for i := 0; i < t; i++ {
		y := strings.Index(string(format[x+1:]), token)
		n = append(n, string(date[y+i]))
	}

	return strings.Join(n, "")
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
