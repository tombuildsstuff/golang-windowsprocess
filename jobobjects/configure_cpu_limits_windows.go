//+build windows

package jobobjects

import (
	"os"
	"syscall"
	"unsafe"
)

// https://docs.microsoft.com/en-gb/windows/desktop/api/winnt/ns-winnt-_jobobject_cpu_rate_control_information
// JOB_OBJECT_CPU_RATE_CONTROL_HARD_CAP
var jobObjectCpuRateControlHardCap = 0x00000004

var jobObjectCpuRateControlEnable = 0x00000001

var jobObjectCpuRateControlInformation = 15

type JOBOBJECT_CPU_RATE_CONTROL_INFORMATION struct {
	ControlFlags uint32
	CpuRate       uint32
}

// ConfigureCPULimits configures the maximum percentage of CPU a given Job Object is allowed to use.
func (j *JobObject) ConfigureCPULimits(percentage int) error {
	// https://docs.microsoft.com/en-gb/windows/desktop/api/winnt/ns-winnt-_jobobject_basic_limit_information
	kernel32 := syscall.MustLoadDLL("Kernel32.dll")
	jobObject := kernel32.MustFindProc("SetInformationJobObject")

	limit := JOBOBJECT_CPU_RATE_CONTROL_INFORMATION{
		ControlFlags: uint32(jobObjectCpuRateControlEnable | jobObjectCpuRateControlHardCap),
		// Set CpuRate to a percentage times 100.
		// For example, to let the job use 20% of the CPU, set CpuRate to 20 times 100, or 2,000.
		CpuRate:        uint32(percentage * 100),
	}
	limitPtr := unsafe.Pointer(&limit)

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms686216(v=vs.85).aspx
	//  = JobObjectExtendedLimitInformation
	r1, _, err := jobObject.Call(uintptr(*j.handle), uintptr(jobObjectCpuRateControlInformation), uintptr(limitPtr), unsafe.Sizeof(limit))
	// a non-zero return value means this is successful
	if int(r1) == 0 {
		return os.NewSyscallError("SetInformationJobObject", err)
	}

	return nil
}
