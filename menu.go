package main

type Menu struct {
	items                         []MenuItem
	selected                      int
	top, left, gap, width, height float64
}

type MenuItem struct {
	Name string
	Func func()
}

func (m *Menu) Mouseover(x, y float64) {
	//x -= left
	//y -= top
	// determine button over and set
}

func (m *Menu) Mouseout() {
	m.selected = -1
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
	if m.selected >= 0 {
		m.items[m.selected].Func()
	}
}
