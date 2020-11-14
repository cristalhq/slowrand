package slowrand

import (
	"crypto/sha512"
	"errors"
	"sync"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ripemd160"
)

// Reader represents a deterministic rand source from seed.
// Slowness is a feature to prevent brute-force attack.
type Reader struct {
	rounds  int
	time    uint32
	memory  uint32
	threads uint8

	mu    *sync.RWMutex
	seed  []byte
	salt  []byte
	key   []byte
	reads int
}

// New instance of slow reader.
// Seed as a start of the pseudorandom sequence.
// Rounds, time, memory and threads are params to the underlying PBKDF2 & Argon2 algorithms.
//
func New(seed []byte, rounds int, time, memory uint32, threads uint8) (*Reader, error) {
	switch {
	case len(seed) == 0:
		return nil, errors.New("slowrand: seed cannoe be empty")

	case rounds < 1:
		return nil, errors.New("slowrand: rounds must be greater than 0")

	case time < 1:
		return nil, errors.New("slowrand: time must be greater than 0")

	case memory < 1:
		return nil, errors.New("slowrand: memory must be greater than 0")

	case threads < 1:
		return nil, errors.New("slowrand: threads must be greater than 0")
	}

	r := &Reader{
		seed:    seed,
		rounds:  rounds,
		time:    time,
		memory:  memory,
		threads: threads,
		mu:      &sync.RWMutex{},
	}
	return r, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seed = pbkdf2.Key(r.seed, r.key, r.rounds, sha512.Size, sha512.New)
	r.salt = pbkdf2.Key(r.salt, r.key, r.reads, ripemd160.Size, ripemd160.New)
	r.key = argon2.Key(r.seed, r.salt, r.time, r.memory, r.threads, uint32(len(b)))

	n := copy(b, r.key)
	r.reads += n

	return n, nil
}

// Len returns number of read bytes.
func (r *Reader) Len() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.reads
}
