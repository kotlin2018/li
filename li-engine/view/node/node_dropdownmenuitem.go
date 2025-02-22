package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func DropdownMenuItem(name string) *dropdownmenuitemBuilder {
	return &dropdownmenuitemBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentDropdownMenuItem,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type dropdownmenuitemBuilder struct {
	*NodeBuilder
}

func (b *dropdownmenuitemBuilder) AC(f ac.AC) *dropdownmenuitemBuilder {
	b.schema.AC = f
	return b
}

func (b *dropdownmenuitemBuilder) Title(title string) *dropdownmenuitemBuilder {
	b.schema.Title = title
	return b
}
