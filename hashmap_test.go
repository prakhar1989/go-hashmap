package hashmap

import "testing"

func TestNewHashMap(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{10, 10},
		{20, 20},
		{0, 0},
	}

	for _, c := range cases {
		got := NewHashMap(c.in).size
		if got != c.want {
			t.Errorf("NewHashMap(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
