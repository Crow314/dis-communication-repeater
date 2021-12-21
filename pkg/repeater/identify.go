package repeater

type idStore struct {
	storage [][]byte
	size    int
	index   int
}

func newIDStore(size int) *idStore {
	res := new(idStore)
	res.size = size
	res.index = 0
	res.storage = make([][]byte, 0, size)

	return res
}

func (store *idStore) add(id []byte) {
	if len(store.storage) < store.size { // 溢れる前
		store.storage = append(store.storage, id)
	} else { // 溢れた後
		store.storage[store.index] = id
		store.index++
		store.index %= store.size
	}
}
