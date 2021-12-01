# Mnemonic

Mnemonic is a Go library which which generates mnemonics, which is a way of 
representing a large randomly-generated number as a sequence of words. A mnemonic 
gives us a convenient way to store this number, these words then may be used to 
create a seed. For example we could represent a 128 bit number as the following 
sequence of words:
```
"window pig filter walk eye damage arena turtle loyal lobster live act"
```
The greater the number the larger the sequence of words will be, a 256 bit number 
would generate a sequence of 24 words.

# Installation

```
$ go get github.com/umahmood/mnemonic
```

# Usage

```
package main

import (
    "fmt"

    "github.com/umahmood/mnemonic"
)

func main() {
    m, err := mnemonic.New(mnemonic.DefaultConfig) // default 128 bits
    if err != nil {
        // ...
    }
    words, err := m.Words()
    if err != nil {
        // ...
    }
    fmt.Println(words)
}
```
Output:
```
[wild cause filter walk eye damage arena turtle loyal lobster live add]
```
Specifying a non-default config.:
```
m, err := mnemonic.New(mnemonic.Config{ 
    Bits: 256, // must be a multiple of 32
})
```
You can specify a passphrase which will be used in generating a seed:
```
m, err := mnemonic.New(mnemonic.Config{ 
    Bits:       256,
    Passphrase: "To the moon!",
})
```

# Documentation

> https://pkg.go.dev/github.com/umahmood/mnemonic

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
