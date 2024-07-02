package void

// Inspired by https://dave.cheney.net/2014/03/25/the-empty-struct

// Void is alias for empty struct{}
type Void struct{}

// Value represents a void value to avoid ugly struct{}{}
var Value struct{}
