package gocolors

//HSL ... converts R, G, B into H, S, L.
func HSL(c Newcolor) ColorHSL {
	var h, s, l, most, least, diff int
	r := c[0]
	g := c[1]
	b := c[2]
	if r > g {
		if b > r {
			most, least = b, g
		} else if b > g {
			most, least = r, g
		} else {
			most, least = r, b
		}
	} else {
		if b > g {
			most, least = b, r
		} else if b > r {
			most, least = g, r
		} else {
			most, least = g, b
		}
	}

	l = int((most + least) / 2)

	if most == least {
		s = 0
		h = 0
	} else {
		diff = most - least
		if l > 127 {
			s = int(diff * 255 / (510 - most - least))
		} else {
			s = int(diff * 255 / (most + least))
		}

		switch most {
		case r:
			if g < b {
				h = 1530 + int((g-b)*255/diff)
			} else {
				h = int((g - b) * 255 / diff)
			}
		case g:
			h = 510 + int((b-r)*255/diff)
		default:
			h = 1020 + int((r-g)*255/diff)
		}
		h = int(h / 6)
	}
	cHSL := ColorHSL{h, s, l}
	return cHSL
}
