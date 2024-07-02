package cookie

func (c *Cookie) GetUserId(key string) string {
	return c.Data[key]
}
