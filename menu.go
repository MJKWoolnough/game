package main

type Menu struct {
	items          []MenuItem
	selected       int
	top, left, gap int
}

type MenuItem struct {
	Name string
	Func func()
}

func (m *Menu) Mouseover(x, y int) {
	x -= left
	y -= top
	// determine button over and set
}

func (m *Menu) Next() {
	m.selected++
	if m.selected >= len(m.items) {
		m.selected = 0
	}
}

func (m *Menu) Prev() {
	if m.selected <= 0 {
		m.selected = len(m.items)
	}
	m.selected--
}

func (m *Menu) Select() {
	if over >= 0 {
		m.items[m.over].Func()
	}
	return nil
}
