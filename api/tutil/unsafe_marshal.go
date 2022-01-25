package tutil

import "encoding/json"

// UnsafeMarshal unmarshalls stuff.  You get punished with a panic if the error is real.
func UnsafeMarshal(v interface{}) []byte {
	output, err := json.Marshal(v)
	PanicIfErr(err)
	return output
}
