package cookie

func (c *Cookie) Delete(key string) {
	delete(c.Data, key)
}
