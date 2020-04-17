// Code generated by "stringer -type=MMSmsCdmaServiceCategory -trimprefix=MmSmsCdmaServiceCategory"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsCdmaServiceCategoryUnknown-0]
	_ = x[MmSmsCdmaServiceCategoryEmergencyBroadcast-1]
	_ = x[MmSmsCdmaServiceCategoryAdministrative-2]
	_ = x[MmSmsCdmaServiceCategoryMaintenance-3]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsLocal-4]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsRegional-5]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsNational-6]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsInternational-7]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsLocal-8]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsRegional-9]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsNational-10]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsInternational-11]
	_ = x[MmSmsCdmaServiceCategorySportsNewsLocal-12]
	_ = x[MmSmsCdmaServiceCategorySportsNewsRegional-13]
	_ = x[MmSmsCdmaServiceCategorySportsNewsNational-14]
	_ = x[MmSmsCdmaServiceCategorySportsNewsInternational-15]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsLocal-16]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsRegional-17]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsNational-18]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsInternational-19]
	_ = x[MmSmsCdmaServiceCategoryLocalWeather-20]
	_ = x[MmSmsCdmaServiceCategoryTrafficReport-21]
	_ = x[MmSmsCdmaServiceCategoryFlightSchedules-22]
	_ = x[MmSmsCdmaServiceCategoryRestaurants-23]
	_ = x[MmSmsCdmaServiceCategoryLodgings-24]
	_ = x[MmSmsCdmaServiceCategoryRetailDirectory-25]
	_ = x[MmSmsCdmaServiceCategoryAdvertisements-26]
	_ = x[MmSmsCdmaServiceCategoryStockQuotes-27]
	_ = x[MmSmsCdmaServiceCategoryEmployment-28]
	_ = x[MmSmsCdmaServiceCategoryHospitals-29]
	_ = x[MmSmsCdmaServiceCategoryTechnologyNews-30]
	_ = x[MmSmsCdmaServiceCategoryMulticategory-31]
	_ = x[MmSmsCdmaServiceCategoryCmasPresidentialAlert-4096]
	_ = x[MmSmsCdmaServiceCategoryCmasExtremeThreat-4097]
	_ = x[MmSmsCdmaServiceCategoryCmasSevereThreat-4098]
	_ = x[MmSmsCdmaServiceCategoryCmasChildAbductionEmergency-4099]
	_ = x[MmSmsCdmaServiceCategoryCmasTest-4100]
}

const (
	_MMSmsCdmaServiceCategory_name_0 = "UnknownEmergencyBroadcastAdministrativeMaintenanceGeneralNewsLocalGeneralNewsRegionalGeneralNewsNationalGeneralNewsInternationalBusinessNewsLocalBusinessNewsRegionalBusinessNewsNationalBusinessNewsInternationalSportsNewsLocalSportsNewsRegionalSportsNewsNationalSportsNewsInternationalEntertainmentNewsLocalEntertainmentNewsRegionalEntertainmentNewsNationalEntertainmentNewsInternationalLocalWeatherTrafficReportFlightSchedulesRestaurantsLodgingsRetailDirectoryAdvertisementsStockQuotesEmploymentHospitalsTechnologyNewsMulticategory"
	_MMSmsCdmaServiceCategory_name_1 = "CmasPresidentialAlertCmasExtremeThreatCmasSevereThreatCmasChildAbductionEmergencyCmasTest"
)

var (
	_MMSmsCdmaServiceCategory_index_0 = [...]uint16{0, 7, 25, 39, 50, 66, 85, 104, 128, 145, 165, 185, 210, 225, 243, 261, 284, 306, 331, 356, 386, 398, 411, 426, 437, 445, 460, 474, 485, 495, 504, 518, 531}
	_MMSmsCdmaServiceCategory_index_1 = [...]uint8{0, 21, 38, 54, 81, 89}
)

func (i MMSmsCdmaServiceCategory) String() string {
	switch {
	case i <= 31:
		return _MMSmsCdmaServiceCategory_name_0[_MMSmsCdmaServiceCategory_index_0[i]:_MMSmsCdmaServiceCategory_index_0[i+1]]
	case 4096 <= i && i <= 4100:
		i -= 4096
		return _MMSmsCdmaServiceCategory_name_1[_MMSmsCdmaServiceCategory_index_1[i]:_MMSmsCdmaServiceCategory_index_1[i+1]]
	default:
		return "MMSmsCdmaServiceCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
