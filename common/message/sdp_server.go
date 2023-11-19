package message

type BundleLoadParams struct {
	Bundle string `json:"bundle"`
}

type RenderingParams struct {
	Value bool `json:"value"`
}

func (c *Command) BuildBundleLoadCommand(params BundleLoadParams) {
	c.Type = "command"
	c.Command = "LoadBundle"
	c.Params = params
}

func (c *Command) BuildBundleUnloadCommand() {
	c.Type = "command"
	c.Command = "unload"
}

func (c *Command) BuildRenderingCommand(params *RenderingParams) {
	c.Type = "command"
	c.Command = "rendering"
	c.Params = params
}
