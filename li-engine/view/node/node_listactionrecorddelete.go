package node

import (
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordDelete(name string) *listactionrecorddeleteBuilder {
	return &listactionrecorddeleteBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordDelete,
			XComponentProps: map[string]interface{}{
				"status":       "danger",
				"confirmProps": make(map[string]interface{}),
			},
		},
	}}
}

type listactionrecorddeleteBuilder struct {
	*NodeBuilder
}

func (b *listactionrecorddeleteBuilder) ForSubmit(operation string, handler interface{}) *listactionrecorddeleteBuilder {
	b.schema.XComponentProps["forSubmit"] = operation
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecorddeleteBuilder) ConfirmTitle(title string) *listactionrecorddeleteBuilder {
	confirmProps, ok := b.schema.XComponentProps["confirmProps"].(map[string]interface{})
	if ok {
		confirmProps["title"] = title
		b.schema.XComponentProps["confirmProps"] = confirmProps
	}
	return b
}
