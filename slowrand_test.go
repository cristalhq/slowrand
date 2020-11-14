package slowrand_test

import (
	"testing"

	"github.com/cristalhq/slowrand"
)

func TestSlowRand(t *testing.T) {
	seed := []byte("test")
	rounds := 3
	time := uint32(5)
	memory := uint32(7)
	threads := uint8(9)

	r, err := slowrand.New(seed, rounds, time, memory, threads)
	if err != nil {
		t.Fatal(err)
	}

	const size = 4
	var buf [size]byte

	for range [4]struct{}{} {
		n, err := r.Read(buf[:])
		if err != nil {
			t.Fatal(err)
		}

		if size != n {
			t.Fatalf("got %v, want %v", n, size)
		}
	}
}

func Test(t *testing.T) {
	seed := []byte("some-secure-seed")
	rounds := 3
	time := uint32(5)
	memory := uint32(7)
	threads := uint8(11)

	r, err := slowrand.New(seed, rounds, time, memory, threads)
	if err != nil {
		t.Fatal(err)
	}

	var buf [42]byte
	n, err := r.Read(buf[:])
	t.Log(n)
	t.Log(err)
}
