// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

// Package sysinfo is a pure Go library providing Linux OS / kernel / hardware system information.
package sysinfo

// SysInfo struct encapsulates all other information structs.
type SysInfo struct {
	Meta    Meta            `json:"sysinfo"`
	Node    Node            `json:"node"`
	OS      OS              `json:"os"`
	Kernel  Kernel          `json:"kernel"`
	Product Product         `json:"product"`
	CPU     CPU             `json:"cpu"`
	Memory  Memory          `json:"memory"`
	Storage []StorageDevice `json:"storage"`
	Network []NetworkDevice `json:"network"`
}

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	// Meta info
	si.getMetaInfo()

	// Software info
	si.getNodeInfo()
	si.getOSInfo()
	si.getKernelInfo()

	// Hardware info
	si.getProductInfo()
	si.getCPUInfo()
	si.getMemoryInfo()
	si.getStorageInfo()
	si.getNetworkInfo()
}
