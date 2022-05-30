//go:build linux && !android

package cicada

import (
	"github.com/zcalusic/sysinfo"
)

// RecognizeOs identifies the environment,
// as an endoflife.date product name.
func RecognizeOs() (*string, error) {
	var si sysinfo.SysInfo
	si.GetSysInfo()
	return &si.OS.Vendor, nil
}
