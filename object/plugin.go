package object

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"plugin"
)

type Plugin struct {
	plugin *plugin.Plugin
	name   string
}

func NewPlugin(name string) (*Plugin, error) {
	p, err := plugin.Open(fmt.Sprintf("%s.so", name))
	if err != nil {
		return &Plugin{}, err // @todo handle error
	}
	return &Plugin{name: name, plugin: p}, nil
}

func (p *Plugin) Type() ObjectType { return PLUGIN_OBJ }
func (p *Plugin) Inspect() string  { return p.name }

func (p *Plugin) Call(method string, args ...Object) Object {
	caser := cases.Title(language.Und, cases.NoLower)

	sym, err := p.plugin.Lookup(caser.String(p.name))
	if err != nil {
		return NewErrorFormat("method %s not exists on plugin %s object. Got: %s", method, p.name, err.Error())
	}

	return NewBuiltin(sym.(BuiltinFunction))
}
