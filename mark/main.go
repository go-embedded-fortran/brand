// Command mark writes the go-embedded-fortran mark, svg/color/go-embedded-fortran.svg:
//
//	cd mark && go run . > ../svg/color/go-embedded-fortran.svg
//
// go-embedded-fortran is a Fortran implementation. The mark is Fortran's
// double colon, ::, the separator of a declaration (integer, parameter :: n = 8):
// two colons side by side. It is a glyph of the language's own syntax, not the
// language's logo.
//
// Nothing is placed by hand. The program takes the dot radius, the spacing of
// the dots inside a colon and the gap between the two colons from the constants
// below and centres the whole glyph, ink included, on the badge. Change a
// constant and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient and a white glyph.
// logos/brandgen renders the format matrix from the file this writes; the kit
// workflow runs this again and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#F472B6" // the gradient's light stop
	dark  = "#9D174D" // its dark stop

	dotR = 14.0 // radius of a dot
	rise = 40.0 // each colon's dots sit at the middle line, plus and minus this
	gap  = 56.0 // between the centres of the two colons
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	const midX, midY = 128.0, 128.0
	var dots []string
	for _, dx := range []float64{-gap / 2, gap / 2} {
		for _, dy := range []float64{-rise, rise} {
			dots = append(dots, fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`,
				num(midX+dx), num(midY+dy), num(dotR)))
		}
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-embedded-fortran">
  <defs><linearGradient id="efgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#efgrad)"/>
  <g fill="#FFFFFF">
    %s
  </g>
</svg>
`, light, dark, strings.Join(dots, "\n    "))
}
