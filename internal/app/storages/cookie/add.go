package cookie

func (c *Cookie) Add(key string, value string) {
	c.Data[key] = value
}
