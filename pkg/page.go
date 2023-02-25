package gofaces

import "golang.org/x/text/language"

type Page interface {
	GetTitle() string
	SetMenu(menu *Menu)
	GetMenu() *Menu
	SetContent(content *Content)
	GetContent() *Content
	GetLanguage() language.Tag
	Render() string
}

type page struct {
	menu      *Menu
	content   *Content
	skin      Skin
	titlei18n string
}

func NewPage(title string) Page {
	p := page{
		titlei18n: title,
		skin:      NewBootstrapprSkin(),
	}
	return &p
}

func (page *page) Render() string {
	return page.skin.Render(page)
}

func (page *page) GetMenu() *Menu {
	return page.menu
}

func (page *page) SetMenu(menu *Menu) {
	page.menu = menu
}

func (page *page) GetLanguage() language.Tag {
	return language.German
}

func (page *page) GetTitle() string {
	return page.titlei18n
}

func (page *page) SetContent(content *Content) {
	page.content = content
}

func (page *page) GetContent() *Content {
	return page.content
}
