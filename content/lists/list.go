package list

type Element interface{}

type List []Element

func (l *List) Len() int {
	return len(*l)
}

func (l *List) Head() *Element {
	if l.Len() > 0 {
		return &(*l)[0]
	} else {
		return nil
	}
}
