_This package was developed by [Silicon Ally](https://siliconally.org) while
working on a project for  [Adventure Scientists](https://adventurescientists.org).
Many thanks to Adventure Scientists for supporting [our open source
mission](https://siliconally.org/policies/open-source/)!_

# cryptorand

[![GoDoc](https://pkg.go.dev/badge/github.com/Silicon-Ally/cryptorand?status.svg)](https://pkg.go.dev/github.com/Silicon-Ally/cryptorand?tab=doc)

`cryptorand` is a simple, zero-dependency Go library that implements the
`math/rand` package's `rand.Source` and `rand.Source64` interfaces backed by
the `crypto/rand` package for cryptographically secure number generation.

This is useful for cases where you don't want a pseudorandom/insecure random
number generator, but would like to use the higher-level interface provided by
`*rand.Rand`

## Usage

```go
package main

import "github.com/Silicon-Ally/cryptorand"

func main() {
  r := cryptorand.New()
  // Will print a number [0,9]
  fmt.Println(r.Intn(10))
}
```

## Panicking Behavior

Generally, Go libraries should avoid panicking except in extreme circumstances,
opting instead to return an error. `*cryptorand.Source` may panic if the
underlying call to `crypto/rand.Read` fails, [which can happen for different
reasons on different
platforms](https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/crypto/rand/).
Because the `math/rand.Source` interface doesn't expose an `error` in the
response, we opt to panic loudly instead of silently failing, as a lack of
randomness can manifest as security vulnerabilities.

## Security

Please report security issues to security@siliconally.org, or by using one of
the contact methods available on our [Contact Us
page](https://siliconally.org/contact/).
