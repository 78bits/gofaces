package gofaces

type VC interface {
	GetChildren() []VC
}

type Text struct {
	Text string
}

type Panel struct {
	Title    string
	Children []VC
}

type Content struct {
	Children []VC
}

func (l *Text) GetChildren() []VC    { return []VC{} }
func (c *Panel) GetChildren() []VC   { return c.Children }
func (c *Content) GetChildren() []VC { return c.Children }
