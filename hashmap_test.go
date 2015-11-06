package hashmap

import (
	"strconv"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 10},
		{20, 20},
		{0, 0},
	}

	for _, c := range cases {
		h, err := NewHashMap(c.in)
		if c.in == 0 && err == nil {
			t.Errorf("Expected error, didn't get it")
		} else {
			if h.size != c.want {
				t.Errorf("NewHashMap(%q) == %q, want %q", c.in, h.size, c.want)
			}
		}
	}
}

func TestSize(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 10},
		{20, 20},
	}

	for _, c := range cases {
		h, _ := NewHashMap(c.in)
		got := h.Size()
		if got != c.want {
			t.Errorf("Size(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestLenAndLoad(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 10},
		{20, 20},
	}

	for _, c := range cases {
		h, _ := NewHashMap(c.in * 10)
		for i := 1; i <= c.in; i++ {
			key := strconv.Itoa(i)
			h.Set(key, c.in*10)
		}
		got := h.Len()
		if got != c.want {
			t.Errorf("Len(%d) == %d, want %d", c.in, got, c.want)
		}

		load := h.Load()
		want := float32(c.in) / float32(c.in*10)
		if load != want {
			t.Errorf("Load(%d) == %f, want %f", c.in, load, want)
		}
	}
}

func TestGetAndSet(t *testing.T) {
	h, _ := NewHashMap(100)
	keys := []string{"alpha", "beta", "charlie", "gamma", "delta"}

	// testing primitives
	for _, key := range keys {
		h.Set(key, len(key))
	}

	for _, key := range keys {
		got, _ := h.Get(key)
		want := len(key)
		if got.Value.(int) != len(key) {
			t.Errorf("want: %q, got: %q", want, got)
		}
	}

	// testing strings
	for _, key := range keys {
		h.Set(key, key+key)
	}

	for _, key := range keys {
		got, _ := h.Get(key)
		want := key + key
		if got.Value.(string) != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	}

	// testing references to compound types
	arr := []int{2, 3, 4}
	h.Set("myArray", arr)
	a, _ := h.Get("myArray")
	k := a.Value.([]int)
	k[0] = 100
	if k[0] != arr[0] {
		t.Errorf("Reference has not been mutated")
	}
}

func TestCollisions(t *testing.T) {
	// a small hashmap that is bound to have collisions
	h, _ := NewHashMap(5)

	keys := []string{"alpha", "beta", "charlie", "gamma", "delta"}

	for _, key := range keys {
		h.Set(key, len(key))
	}

	for _, key := range keys {
		got, _ := h.Get(key)
		want := len(key)
		if got.Value.(int) != len(key) {
			t.Errorf("want: %q, got: %q", want, got)
		}
	}
}

func TestDelete(t *testing.T) {
	// a hashtable with just one elem
	h, _ := NewHashMap(1)
	h.Set("alpha", 10)

	// should not be allowed to add
	status := h.Set("beta", 20)
	if status {
		t.Errorf("Able to add more elements than the allowed size")
	}

	// lets delete first
	_, status = h.Delete("alpha")
	if !status {
		t.Errorf("Unable to delete")
	}

	// and add beta now
	status = h.Set("beta", 20)
	if !status {
		t.Errorf("Unable to add element")
	}

	// lastly, lets delete a non-existent element
	_, status = h.Delete("gamma")
	if status {
		t.Errorf("Deleted a missing key")
	}
}
