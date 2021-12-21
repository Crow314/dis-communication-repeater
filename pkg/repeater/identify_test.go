package repeater

import (
	"testing"
)

func TestAddSize(t *testing.T) {
	size := 4
	store := newIDStore(size)

	if len(store.storage) != 0 || cap(store.storage) != size {
		t.Fatal("initialize error")
	}

	for i := 0; i < size; i++ {
		store.add([]byte{0x43, 0x72, 0x6F, 0x77, byte(i)})

		if len(store.storage) > size {
			t.Fatalf("size limit over\n%v", store.storage)
		}

		if cap(store.storage) != size {
			t.Fatalf("capacity limit over\n%v", store.storage)
		}

		if store.index >= size {
			t.Fatalf("index range is wrong\n%v", store.storage)
		}
	}

	for i := 0; i < size+1; i++ {
		store.add([]byte{0x43, 0x72, 0x6F, 0x77, byte(i + 8)})

		if len(store.storage) > size {
			t.Fatalf("size limit over\n%v", store.storage)
		}

		if cap(store.storage) != size {
			t.Fatalf("capacity limit over\n%v", store.storage)
		}

		if store.index >= size {
			t.Fatalf("index range is wrong\n%v", store.storage)
		}
	}
}
