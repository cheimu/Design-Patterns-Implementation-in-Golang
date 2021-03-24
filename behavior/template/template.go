package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

// abstract class
type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

// not implement template but just an example
type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

// use adapater to create a concrete class for template
type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
