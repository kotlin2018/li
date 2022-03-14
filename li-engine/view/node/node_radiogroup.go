package node

import "github.com/BeanWei/li/li-engine/view/ui"

func RadioGroup(name string) *radiogroupBuilder {
	return &radiogroupBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeAny,
		XComponent: ui.ComponentRadioGroup,
		XDecorator: ui.DecoratorFormItem,
		Enum:       make([]map[string]interface{}, 0),
	}}
}

type radiogroupBuilder struct {
	schema *ui.Schema
}

func (b *radiogroupBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *radiogroupBuilder) Required() *ui.Schema {
	b.schema.Required = true
	return b.schema
}

func (b *radiogroupBuilder) Title(title string) *ui.Schema {
	b.schema.Title = title
	return b.schema
}

func (b *radiogroupBuilder) Option(label string, value ...interface{}) *ui.Schema {
	var val interface{}
	if len(value) > 0 {
		val = value[0]
	} else {
		val = label
	}
	b.schema.Enum = append(b.schema.Enum, map[string]interface{}{
		"label": label,
		"value": val,
	})
	return b.schema
}
