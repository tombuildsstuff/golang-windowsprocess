//+build !windows

package jobobjects

import "fmt"

// ConfigureMemoryLimits configures the Memory Limits for a given Job Object
func (j *JobObject) ConfigureMemoryLimits(minWorkingSetKb uint32, maxWorkingSetKb uint32) error {
	return fmt.Errorf("Job Objects are not supported on this Operating System.")
}
