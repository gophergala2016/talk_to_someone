package ws

import (
	//"math/rand"
	//"time"

	"github.com/streamrail/concurrent-map"
)

var (
	store = cmap.New()
	ready = cmap.New()
)

func Save(id string, user *User) {
	store.Set(id, user)
	ready.Set(id, user)
}

func GetNumberOfAvailableUsers() int {
	return ready.Count()
}

func Get(id string) (*User, bool) {
	value, exist := store.Get(id)
	return value.(*User), exist
}

func deleteFromAvailable(id string) {
	ready.Remove(id)
}

/*func GetRandomFreeUser() (*User, bool) {
	ids := make([]string, ready.Count())
	for id := range ready.Iter() {
		append(ids, id.Key)
	}
	user, exist := ready.Get(ids[random(0, len(ids)-1)])
	return user.(*User), exist
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}*/

func Clear() {
	store = cmap.New()
	ready = cmap.New()
}

func Size() int {
	return store.Count()
}
