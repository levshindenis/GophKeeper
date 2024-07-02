package cookie

func (c *Cookie) InCookies(key string) bool {
	if _, ok := c.Data[key]; ok {
		return true
	}
	return false
}
