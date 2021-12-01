/*
Package mnemonic a way of representing a large randomly-generated number as a
sequence of words, making it easier for humans to store.

	package main

	import (
		"fmt"

		"github.com/umahmood/mnemonic"
	)

	func main() {
		m, err := mnemonic.New(mnemonic.DefaultConfig)
		if err != nil {
			// ...
		}
		words, err := m.Words()
		if err != nil {
			// ...
		}
		fmt.Println(words)
	}
*/
package mnemonic
