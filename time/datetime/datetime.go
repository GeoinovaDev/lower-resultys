package datetime

import (
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

// ParseUTC ...
func ParseUTC(utc string) *Datetime {
	// 2018-05-16 03:01:00 +0000 UTC
	// 2018-05-16 09:29:55.8957322 -0300 -03 m=+0.005003101
	p := strings.Split(utc, " ")
	if len(p) < 2 {
		return nil
	}

	data := strings.Split(p[0], "-")
	y := data[0]
	M := data[1]
	d := data[2]

	hora := strings.Split(p[1], ".")
	hora = strings.Split(hora[0], ":")
	h := hora[0]
	m := hora[1]
	s := hora[2]

	time, _ := time.Parse("2006-01-02 15:04:05", str.Format("{0}-{1}-{2} {3}:{4}:{5}", y, M, d, h, m, s))
	return &Datetime{time: time}
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
