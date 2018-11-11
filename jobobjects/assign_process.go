//+build !windows

package jobobjects

import "fmt"

// AssignProcess assigns the specified Process ID to the specified Job Object
// once a process has been assigned to a Job Object, it's not possible to unassign it
func (j *JobObject) AssignProcess(processId int) error {
	return fmt.Errorf("Job Objects are not supported on this Operating System.")
}
