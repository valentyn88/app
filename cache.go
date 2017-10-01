package app

type Cache struct {
	Users     map[int]*User
	Locations map[int]*Location
	Visits    map[int]*Visit
}

func NewCache() (cache *Cache) {
	cache = &Cache{
		Users:     make(map[int]*User),
		Locations: make(map[int]*Location),
		Visits:    make(map[int]*Visit),
	}
	return
}
