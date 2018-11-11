//+build windows

package jobobjects

import (
	"os"
	"syscall"
	"unsafe"
)

// Create creates and returns the handle to a Job Object
func Create(name string) (*JobObject, error) {
	// https://docs.microsoft.com/en-gb/windows/desktop/api/winbase/nf-winbase-createjobobjecta
	null := uintptr(unsafe.Pointer(nil))
	nameVal := uintptr(unsafe.Pointer(&name))
	kernel32 := syscall.MustLoadDLL("Kernel32.dll")
	jobObject := kernel32.MustFindProc("CreateJobObjectA")

	r1, _, err := jobObject.Call(null, nameVal)
	// we should receive a handle to the Job Object
	if int(r1) == 0 {
		return nil, os.NewSyscallError("CreateJobObjectA", err)
	}

	ptr := uint32(r1)
	obj := JobObject{
		handle: &ptr,
	}
	return &obj, nil
}
