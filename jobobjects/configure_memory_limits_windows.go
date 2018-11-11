//+build windows

package jobobjects

import (
	"os"
	"syscall"
	"unsafe"
)

// JOB_OBJECT_LIMIT_SILENT_BREAKAWAY_OK
var jobObjectLimitSilentBreakawayOk = 0x00001000

// JOB_OBJECT_LIMIT_WORKING_SET
var jobObjectLimitMemory = 0x00000001

// JobObjectExtendedLimitInformation
var jobObjectExtendedLimitInformation = 0x00000009

// ConfigureMemoryLimits configures the Memory Limits for a given Job Object
func (j *JobObject) ConfigureMemoryLimits(minWorkingSetKb uint32, maxWorkingSetKb uint32) error {
	// https://docs.microsoft.com/en-gb/windows/desktop/api/winnt/ns-winnt-_jobobject_basic_limit_information
	kernel32 := syscall.MustLoadDLL("Kernel32.dll")
	jobObject := kernel32.MustFindProc("SetInformationJobObject")

	minWorkingSetBytes := minWorkingSetKb * 1024
	maxWorkingSetBytes := maxWorkingSetKb * 1024

	limit := JOBOBJECT_EXTENDED_LIMIT_INFORMATION{
		BasicLimitInformation: JOBOBJECT_BASIC_LIMIT_INFORMATION{
			LimitFlags:            uint32(jobObjectLimitMemory | jobObjectLimitSilentBreakawayOk),
			MinimumWorkingSetSize: uintptr(minWorkingSetBytes),
			MaximumWorkingSetSize: uintptr(maxWorkingSetBytes),
		},
	}
	limitPtr := unsafe.Pointer(&limit)

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms686216(v=vs.85).aspx
	//  = JobObjectExtendedLimitInformation
	r1, _, err := jobObject.Call(uintptr(*j.handle), uintptr(jobObjectExtendedLimitInformation), uintptr(limitPtr), unsafe.Sizeof(limit))
	// a non-zero return value means this is successful
	if int(r1) == 0 {
		return os.NewSyscallError("SetInformationJobObject", err)
	}

	return nil
}

type JOBOBJECT_BASIC_LIMIT_INFORMATION struct {
	PerProcessUserTimeLimit int64
	PerJobUserTimeLimit     int64
	LimitFlags              uint32
	MinimumWorkingSetSize   uintptr
	MaximumWorkingSetSize   uintptr
	ActiveProcessLimit      uint32
	Affinity                uintptr
	PriorityClass           uint32
	SchedulingClass         uint32
}

type IO_COUNTERS struct {
	ReadOperationCount  int64
	WriteOperationCount int64
	OtherOperationCount int64
	ReadTransferCount   int64
	WriteTransferCount  int64
	OtherTransferCount  int64
}

type JOBOBJECT_EXTENDED_LIMIT_INFORMATION struct {
	BasicLimitInformation JOBOBJECT_BASIC_LIMIT_INFORMATION
	IoInfo                IO_COUNTERS
	ProcessMemoryLimit    uintptr
	JobMemoryLimit        uintptr
	PeakProcessMemoryUsed uintptr
	PeakJobMemoryUsed     uintptr
}
