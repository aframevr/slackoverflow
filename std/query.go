package std

import "fmt"

// QueryParams builds a request query
type QueryParams map[string]string

// Add adds a new parameter given by the key and value
func (p QueryParams) Add(key string, value interface{}) {
	p[key] = fmt.Sprintf("%v", value)
}

// Set change the value of the given param
func (p QueryParams) Set(key string, value interface{}) {
	p.Add(key, value)
}

// Del deletes the given parameter
func (p QueryParams) Del(key string) {
	delete(p, key)
}
