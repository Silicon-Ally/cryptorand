_This package brought to you by [Adventure
Scientists](https://adventurescientists.org). Read more about [our open source
policy here](https://siliconally.org/policies/open-source/)._

# Cryptorand

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
