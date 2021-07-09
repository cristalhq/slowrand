# slowrand

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Slow deterministic random generator.

## Rationale

Some problems might require random deterministic generator, but this might be problematic due to brute-force attacks.
This (deterministic) random generator addresses this problem by increasing it's work time.

## Note

Inspired by [cornfeedhobo/ssh-keydgen](https://github.com/cornfeedhobo/ssh-keydgen)

## Features

- Slow
- Deterministic
- Random
- Dependency-free

## Install

Go version 1.15+

```
go get github.com/cristalhq/slowrand
```

## Example

```go
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
if n != 42 {
    panic("not 42")
}
_ = err // is always nil
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/slowrand/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/slowrand/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/slowrand
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/slowrand
[reportcard-img]: https://goreportcard.com/badge/cristalhq/slowrand
[reportcard-url]: https://goreportcard.com/report/cristalhq/slowrand
[coverage-img]: https://codecov.io/gh/cristalhq/slowrand/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/slowrand
