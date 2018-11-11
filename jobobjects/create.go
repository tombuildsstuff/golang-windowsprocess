//+build !windows

package jobobjects

import "fmt"

// Create creates and returns the handle to a Job Object
func Create(name string) (*JobObject, error) {
	return nil, fmt.Errorf("Job Objects are not supported on this Operating System.")
}
