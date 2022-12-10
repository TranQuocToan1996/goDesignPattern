package behavioralPatterns

import "strings"

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

// Adapter pattern for AnonymousTemplate can work with Template interface
type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}
func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{myFunc: f}
}
