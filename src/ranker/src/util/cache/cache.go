package cache

type Cache struct {
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item := c.store[key]

	return item, item != nil
}

func (c *Cache) Set(key string, value interface{}) {
	c.store[key] = value
}
