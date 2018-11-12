//+build !windows

package jobobjects

import "fmt"

// ConfigureCPULimits configures the maximum percentage of CPU a given Job Object is allowed to use.
func (j *JobObject) ConfigureCPULimits(percentage int) error {
	return fmt.Errorf("Job Objects are not supported on this Operating System.")
}
