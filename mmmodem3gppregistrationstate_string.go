// Code generated by "stringer -type=MMModem3gppRegistrationState -trimprefix=MmModem3gppRegistrationState"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModem3gppRegistrationStateIdle-0]
	_ = x[MmModem3gppRegistrationStateHome-1]
	_ = x[MmModem3gppRegistrationStateSearching-2]
	_ = x[MmModem3gppRegistrationStateDenied-3]
	_ = x[MmModem3gppRegistrationStateUnknown-4]
	_ = x[MmModem3gppRegistrationStateRoaming-5]
	_ = x[MmModem3gppRegistrationStateHomeSmsOnly-6]
	_ = x[MmModem3gppRegistrationStateRoamingSmsOnly-7]
	_ = x[MmModem3gppRegistrationStateEmergencyOnly-8]
	_ = x[MmModem3gppRegistrationStateHomeCsfbNotPreferred-9]
	_ = x[MmModem3gppRegistrationStateRoamingCsfbNotPreferred-10]
}

const _MMModem3gppRegistrationState_name = "IdleHomeSearchingDeniedUnknownRoamingHomeSmsOnlyRoamingSmsOnlyEmergencyOnlyHomeCsfbNotPreferredRoamingCsfbNotPreferred"

var _MMModem3gppRegistrationState_index = [...]uint8{0, 4, 8, 17, 23, 30, 37, 48, 62, 75, 95, 118}

func (i MMModem3gppRegistrationState) String() string {
	if i >= MMModem3gppRegistrationState(len(_MMModem3gppRegistrationState_index)-1) {
		return "MMModem3gppRegistrationState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMModem3gppRegistrationState_name[_MMModem3gppRegistrationState_index[i]:_MMModem3gppRegistrationState_index[i+1]]
}
