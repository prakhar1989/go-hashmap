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
