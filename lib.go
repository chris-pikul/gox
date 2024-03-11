// Package Gox provides useful utilities to extend the standard Go library.
package gox

import "encoding/json"

// Provides a way to declare the JSON marshaling function for the entire library.
// It defaults to the STD [json.Marshal].
var JSONMarshaler = json.Marshal

// Provides a way to declare the JSON unmarshal-ing function for the entire
// library. it defaults to the STD [json.Unmarshal].
var JSONUnmarshaler = json.Unmarshal
