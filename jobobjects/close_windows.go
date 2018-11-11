//+build windows

package jobobjects

import "syscall"

// Close ensures the handle associated with this Job Object is closed
func (j *JobObject) Close() error {
	if j.handle == nil {
		return nil
	}

	return syscall.CloseHandle(syscall.Handle(*j.handle))
}
