package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func UploadAvatar(name string) *uploadavatarBuilder {
	return &uploadavatarBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeString,
			XComponent: ui.ComponentUploadAvatar,
			XComponentProps: map[string]interface{}{
				"action": "/api/liql",
				"data": map[string]string{
					"operation": "@uploadFile",
				},
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type uploadavatarBuilder struct {
	*NodeBuilder
}

func (b *uploadavatarBuilder) AC(f ac.AC) *uploadavatarBuilder {
	b.schema.AC = f
	return b
}

func (b *uploadavatarBuilder) Title(title string) *uploadavatarBuilder {
	b.schema.Title = title
	return b
}

func (b *uploadavatarBuilder) Description(description string) *uploadavatarBuilder {
	b.schema.Description = description
	return b
}

func (b *uploadavatarBuilder) Default(value interface{}) *uploadavatarBuilder {
	b.schema.Default = value
	return b
}

func (b *uploadavatarBuilder) Accept(accept string) *uploadavatarBuilder {
	b.schema.XComponentProps["accept"] = accept
	return b
}

func (b *uploadavatarBuilder) Tip(tip string) *uploadavatarBuilder {
	b.schema.XComponentProps["tip"] = tip
	return b
}
