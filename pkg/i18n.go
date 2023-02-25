package gofaces

import "golang.org/x/text/language"

type I18n interface {
	Text(key string, language language.Tag) string
}

type i18n struct {
}

func NewI18n() I18n {
	return &i18n{}
}

func (i18n *i18n) Text(key string, language language.Tag) string {
	return key
}
