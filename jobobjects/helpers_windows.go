//+build windows

package jobobjects

import (
	"os"
	"syscall"
)

// https://docs.microsoft.com/en-gb/windows/desktop/ProcThread/process-security-and-access-rights
var processSetInformation = 0x0200

// https://stackoverflow.com/questions/15237357/windll-kernel32-openprocessprocess-all-access-pid-false-process-all-access
var processAllAccess = 0x2035711

func getHandle(processId int, inheritProcess bool) (*uint32, error) {
	// https://docs.microsoft.com/en-us/windows/desktop/api/processthreadsapi/nf-processthreadsapi-openprocess
	inheritHandle := 0
	if inheritProcess {
		inheritHandle = 1
	}
	// TODO: determine how to do least privilege
	desiredAccessPtr := uintptr(uint32(processAllAccess))
	inheritHandlePtr := uintptr(uint32(inheritHandle))
	processIdPtr := uintptr(uint32(processId))

	kernel32 := syscall.MustLoadDLL("Kernel32.dll")
	jobObject := kernel32.MustFindProc("OpenProcess")

	r1, _, err := jobObject.Call(desiredAccessPtr, inheritHandlePtr, processIdPtr)
	// a non-zero return value means this is successful
	if int(r1) == 0 {
		return nil, os.NewSyscallError("OpenProcess", err)
	}

	pid := uint32(r1)
	return &pid, nil
}
