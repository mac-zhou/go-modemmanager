// Code generated by "stringer -type=MMBearerIpMethod -trimprefix=MmBearerIpMethod"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmBearerIpMethodUnknown-0]
	_ = x[MmBearerIpMethodPpp-1]
	_ = x[MmBearerIpMethodStatic-2]
	_ = x[MmBearerIpMethodDhcp-3]
}

const _MMBearerIpMethod_name = "UnknownPppStaticDhcp"

var _MMBearerIpMethod_index = [...]uint8{0, 7, 10, 16, 20}

func (i MMBearerIpMethod) String() string {
	if i >= MMBearerIpMethod(len(_MMBearerIpMethod_index)-1) {
		return "MMBearerIpMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMBearerIpMethod_name[_MMBearerIpMethod_index[i]:_MMBearerIpMethod_index[i+1]]
}
