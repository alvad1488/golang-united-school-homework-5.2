package cache

import (
	"time"
)

type rowCache struct {
	key           string
	value         string
	expiredAt     time.Time
	shouldRemoved bool
}

type Cache struct {
	rowset map[string]rowCache
}

func NewCache() Cache {
	return Cache{make(map[string]rowCache)}
}

//cleaning expired cache
func (inp *Cache) CleanUp() {
	currentDate := time.Now()
	iter := inp.rowset

	for _, row := range iter {
		if row.shouldRemoved {
			if currentDate.After(row.expiredAt) {
				delete(iter, row.key)
			}
		}
	}
}

func (inp *Cache) Get(key string) (string, bool) {

	inp.CleanUp()
	data := inp.rowset
	val, ok := data[key]
	return val.value, ok

}

func (inp *Cache) Put(key, value string) {
	data := rowCache{key: key, value: value, shouldRemoved: false}
	inp.rowset[key] = data
}

func (inp *Cache) Keys() []string {
	var values []string

	inp.CleanUp()
	data := inp.rowset
	for _, row := range data {
		values = append(values, row.value)
	}

	return values

}

func (inp *Cache) PutTill(key, value string, deadline time.Time) {

	inp.CleanUp()
	inp.Put(key, value)

	data := inp.rowset[key]
	data.expiredAt = deadline
	data.shouldRemoved = true
	inp.rowset[key] = data

}
