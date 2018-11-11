//+build !windows

package jobobjects

import "fmt"

// Close ensures the handle associated with this Job Object is closed
func (j *JobObject) Close() error {
	return fmt.Errorf("Job Objects are not supported on this Operating System.")
}
