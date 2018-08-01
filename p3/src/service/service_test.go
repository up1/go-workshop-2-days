package service_test

import (
	. "service"
	"testing"
)

func Test_FormatPatientID_Input_Year_2018_PatientCount_1_Shoule_Be_2018_0001(t *testing.T) {
	year := 2018
	patientCount := 1
	expectedPatientID := "2018-0001"

	actualPatientID := FormatPatientID(year, patientCount)
	if actualPatientID != expectedPatientID {
		t.Errorf("expected %s but it got %s", expectedPatientID, actualPatientID)
	}
}
