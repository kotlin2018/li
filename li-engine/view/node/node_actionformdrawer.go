package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ActionFormDrawer(name string) *actionformdrawerBuilder {
	return &actionformdrawerBuilder{
		&NodeBuilder{
			schema: &ui.Schema{
				Name:       name,
				Type:       ui.SchemaTypeVoid,
				XComponent: ui.ComponentActionFormDrawer,
				XComponentProps: map[string]interface{}{
					"drawerProps": make(map[string]interface{}),
					"layoutProps": make(map[string]interface{}),
				},
				HandlerNames: make([]string, 0),
			},
		},
	}
}

type actionformdrawerBuilder struct {
	*NodeBuilder
}

func (b *actionformdrawerBuilder) AC(f ac.AC) *actionformdrawerBuilder {
	b.schema.AC = f
	return b
}

func (b *actionformdrawerBuilder) Title(title string) *actionformdrawerBuilder {
	b.schema.Title = title
	return b
}

func (b *actionformdrawerBuilder) Description(description string) *actionformdrawerBuilder {
	b.schema.Description = description
	return b
}

func (b *actionformdrawerBuilder) IsMenuItem(isMenuItem bool) *actionformdrawerBuilder {
	b.schema.XComponentProps["isMenuItem"] = isMenuItem
	return b
}

// TODO: 支持通过回调函数获取
func (b *actionformdrawerBuilder) InitialValues(initialValues map[string]interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["initialValues"] = initialValues
	return b
}

func (b *actionformdrawerBuilder) ForInit(operation string, handler interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *actionformdrawerBuilder) ButtonStyle(style map[string]interface{}) *actionformdrawerBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *actionformdrawerBuilder) ButtonType(typ string) *actionformdrawerBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *actionformdrawerBuilder) ButtonStatus(status string) *actionformdrawerBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *actionformdrawerBuilder) ButtonSize(size string) *actionformdrawerBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *actionformdrawerBuilder) ButtonShape(shape string) *actionformdrawerBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *actionformdrawerBuilder) ButtonIcon(icon string) *actionformdrawerBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *actionformdrawerBuilder) ButtonIconOnly(iconOnly bool) *actionformdrawerBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *actionformdrawerBuilder) ButtonLong(long bool) *actionformdrawerBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}

func (b *actionformdrawerBuilder) Body(elements ...view.Node) *actionformdrawerBuilder {
	b.Items(elements...)
	return b
}

func (b *actionformdrawerBuilder) Footer(elements ...view.Node) *actionformdrawerBuilder {
	b.Children(elements...)
	return b
}
