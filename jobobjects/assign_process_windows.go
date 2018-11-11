//+build windows

package jobobjects

import (
	"fmt"
	"os"
	"syscall"
)

// AssignProcess assigns the specified Process ID to the specified Job Object
// once a process has been assigned to a Job Object, it's not possible to unassign it
func (j *JobObject) AssignProcess(processId int) error {
	// https://msdn.microsoft.com/en-gb/f5d7a39f-6afe-4e4a-a802-e7f875ea6e5b
	processHandle, err := getHandle(processId, true)
	if err != nil {
		return fmt.Errorf("Error obtaining handle for process %d: %+v", processId, err)
	}

	kernel32 := syscall.MustLoadDLL("Kernel32.dll")
	jobObject := kernel32.MustFindProc("AssignProcessToJobObject")

	r1, _, err := jobObject.Call(uintptr(*j.handle), uintptr(*processHandle))
	// a non-zero return value means this is successful
	if int(r1) == 0 {
		return os.NewSyscallError("AssignProcessToJobObject", err)
	}

	return nil
}
