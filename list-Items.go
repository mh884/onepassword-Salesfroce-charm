package main

import (
	"math/rand"
	"sync"

	"github.com/1Password/connect-sdk-go/onepassword"
)

type onepasswordItems struct {
	titles     []string
	descs      []string
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
	shuffle    *sync.Once
}

func (r *onepasswordItems) reset(onepasswordItem []onepassword.Item) {
	r.mtx = &sync.Mutex{}
	r.shuffle = &sync.Once{}

	size := len(onepasswordItem)
	r.titles = make([]string, size)
	r.descs = make([]string, size)
	for k, v := range onepasswordItem {
		r.titles[k] = v.Title
		r.descs[k] = v.ID
	}

	r.shuffle.Do(func() {
		shuf := func(x []string) {
			rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })
		}
		shuf(r.titles)
		shuf(r.descs)
	})
}

func (r *onepasswordItems) next(onepasswordItem []onepassword.Item) item {
	if r.mtx == nil {
		r.reset(onepasswordItem)
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := item{
		title:       r.titles[r.titleIndex],
		description: r.descs[r.descIndex],
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	r.descIndex++
	if r.descIndex >= len(r.descs) {
		r.descIndex = 0
	}

	return i
}
