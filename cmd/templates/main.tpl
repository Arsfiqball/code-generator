package {{.PkgName}}

import (
	"github.com/google/wire"
	{{range .Configs}} "{{.Source}}"
	{{end}}

	{{if .UseFiber}}
	"github.com/gofiber/fiber/v2"
	{{end}}
	{{if .UseWorkJob}}
	"github.com/gocraft/work"
	{{end}}
	{{if .UseWatermillSub}}
	"github.com/ThreeDotsLabs/watermill/message"
	{{end}}
)

type Config struct {
	{{range .Configs}} {{.Name}} {{.Type}}
	{{end}}
}

type {{.StructName}} struct {
	// handler *Handler
}

{{if .UseFiber}}
func (c *{{.StructName}}) Route(router fiber.Router) {
	// router.Post("somewhere", c.handler.DoSomething)
}
{{end}}

{{if .UseWorkJob}}
type JobMap map[string]func(*work.Job) error

func (c *{{.StructName}}) Jobs() JobMap {
	return JobMap{
		// "do_something": c.handler.DoSomething,
	}
}
{{end}}

{{if .UseWatermillSub}}
func (c *{{.StructName}}) PubSubRoute(pub message.Publisher, sub message.Subscriber, router *message.Router) {
	// router.AddNoPublisherHandler(
	// 	"example.PrintOnHello", // name of the listener
	// 	"example.hello", // event name this listener is listening to
	// 	sub,
	// 	c.handler.DoSomething,
	// )
}
{{end}}

var RegisterSet = wire.NewSet(
	wire.Struct(new( {{.StructName}} ), "*"),
	wire.FieldsOf(new(Config) {{range .Configs}} , "{{.Name}}" {{end}} ,),
)
