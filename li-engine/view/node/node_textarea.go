package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func TextArea(name string) *textareaBuilder {
	return &textareaBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentInputTextArea,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type textareaBuilder struct {
	*NodeBuilder
}

func (b *textareaBuilder) AC(f ac.AC) *textareaBuilder {
	b.schema.AC = f
	return b
}

func (b *textareaBuilder) Title(title string) *textareaBuilder {
	b.schema.Title = title
	return b
}

func (b *textareaBuilder) Description(description string) *textareaBuilder {
	b.schema.Description = description
	return b
}

func (b *textareaBuilder) Default(value interface{}) *textareaBuilder {
	b.schema.Default = value
	return b
}

func (b *textareaBuilder) Style(style map[string]interface{}) *textareaBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *textareaBuilder) AllowClear(allowClear bool) *textareaBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *textareaBuilder) Placeholder(placeholder string) *textareaBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *textareaBuilder) Size(size string) *textareaBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *textareaBuilder) ShowWordLimit(showWordLimit bool) *textareaBuilder {
	b.schema.XComponentProps["showWordLimit"] = showWordLimit
	return b
}

func (b *textareaBuilder) AutoSize(autoSize bool) *textareaBuilder {
	b.schema.XComponentProps["autoSize"] = autoSize
	return b
}
