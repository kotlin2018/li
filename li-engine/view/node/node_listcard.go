package node

import "github.com/BeanWei/li/li-engine/view/ui"

func ListCard(name string) *listcardBuilder {
	return &listcardBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
		XComponent:      ui.ComponentListCard,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type listcardBuilder struct {
	schema *ui.Schema
}

func (b *listcardBuilder) Schema() *ui.Schema {
	return b.schema
}
