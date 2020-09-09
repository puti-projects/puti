package utils

// String convert util
// Example:
// intVar, err := utils.StrTo("123").Int()
// intVar := utils.StrTo("123").MustInt()

import "strconv"

type StrTo string

// String convert back to string
func (s StrTo) String() string {
	return string(s)
}

// Int convert string to int
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt convert string to int without error
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// Int convert string to uint32
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUInt32 convert string to uint32 without error
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

// Int convert string to uint32
func (s StrTo) UInt64() (uint64, error) {
	v, err := strconv.Atoi(s.String())
	return uint64(v), err
}

// MustUInt32 convert string to uint32 without error
func (s StrTo) MustUInt64() uint64 {
	v, _ := s.UInt64()
	return v
}
