package lru

import "container/list"

type LRUCache interface {
	// Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
	// В случае дублирования ключа вернуть false
	// В случае превышения размера - вытесняется наименее приоритетный элемент
	Add(key, value string) bool

	// Возвращает значение под ключом и флаг его наличия в кеше
	// В случае наличия в кеше элемента повышает его приоритет
	Get(key string) (value string, ok bool)

	// Удаляет элемент из кеша, в случае успеха возврашает true, в случае отсутствия элемента - false
	Remove(key string) (ok bool)
}

type Node struct {
	Data   string
	KeyPrt *list.Element
}

type LRU struct {
	Queue    *list.List
	Items    map[string]*Node
	Capacity int
}

func NewLRUCache(cap int) LRUCache {
	return &LRU{
		Queue:    list.New(),
		Items:    make(map[string]*Node),
		Capacity: cap,
	}
}

func (c *LRU) Add(key string, value string) bool {
	if _, ok := c.Items[key]; !ok {
		c.Items[key] = &Node{
			Data:   value,
			KeyPrt: c.Queue.PushFront(key),
		}

		if len(c.Items) > c.Capacity {
			back := c.Queue.Back()
			c.Queue.Remove(back)
			delete(c.Items, back.Value.(string))
		}

		return !ok
	}

	return false
}

func (c *LRU) Get(key string) (value string, ok bool) {
	if item, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(item.KeyPrt)
		return item.Data, ok
	}

	return value, ok
}

func (c *LRU) Remove(key string) (ok bool) {
	if item, ok := c.Items[key]; ok {
		c.Queue.Remove(item.KeyPrt)
		delete(c.Items, key)
		return ok
	}

	return false
}
