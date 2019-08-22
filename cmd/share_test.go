package cmd

import (
	"testing"
)

func Test_0PartsNotBetween2And255(t *testing.T) {
	got := validParams(0, 1, false, "")
	if got != ErrInvalidPartsSize {
		t.Errorf("Expected parts == 0 to fail validation")
	}
}

func Test_1PartNotBetween2And255(t *testing.T) {
	got := validParams(1, 1, false, "")
	if got != ErrInvalidPartsSize {
		t.Errorf("Expected parts == 1 to fail validation")
	}

}

func Test_256PartsNotBetween2And255(t *testing.T) {
	got := validParams(256, 1, false, "")
	if got != ErrInvalidPartsSize {
		t.Errorf("Expected parts == 256 to fail validation")
	}
}

func Test_2PartsBetween2And255(t *testing.T) {
	got := validParams(2, 1, false, "")
	if got != nil {
		t.Errorf("Expected parts == 2 to pass validation")
	}
}

func Test_255PartsBetween2And255(t *testing.T) {
	got := validParams(255, 1, false, "")
	if got != nil {
		t.Errorf("Expected parts == 255 to pass validation")
	}
}

func Test_Threshold0NotBetween1And255(t *testing.T) {
	got := validParams(10, 0, false, "")
	if got != ErrInvalidThresholdSize {
		t.Errorf("Expected threshold == 0 to fail validation")
	}
}

func Test_Threshold256NotBetween1And255(t *testing.T) {
	got := validParams(10, 256, false, "")
	if got != ErrInvalidThresholdSize {
		t.Errorf("Expected threshold == 256 to fail validation")
	}
}

func Test_Threshold1Between1And255(t *testing.T) {
	got := validParams(10, 1, false, "")
	if got != nil {
		t.Errorf("Expected threshold == 1 to pass validation")
	}
}

func Test_Threshold255Between1And255(t *testing.T) {
	got := validParams(255, 255, false, "")
	if got != nil {
		t.Errorf("Expected threshold == 255 to pass validation")
	}
}

func Test_Threshold9LessThanParts10(t *testing.T) {
	got := validParams(10, 9, false, "")
	if got != nil {
		t.Errorf("Expected threshold == 9 and parts == 10 to pass validation")
	}
}

func Test_Threshold10LessThanParts10(t *testing.T) {
	got := validParams(10, 10, false, "")
	if got != nil {
		t.Errorf("Expected threshold == 10 and parts == 10 to pass validation")
	}
}

func Test_Threshold3NotLessThanParts2(t *testing.T) {
	got := validParams(2, 3, false, "")
	if got != ErrThresholdGreaterThanParts {
		t.Errorf("Expected threshold == 3 and parts == 2 to fail validation")
	}
}

func Test_Threshold11NotLessThanParts10(t *testing.T) {
	got := validParams(10, 11, false, "")
	if got != ErrThresholdGreaterThanParts {
		t.Errorf("Expected threshold == 11 and parts == 10 to fail validation")
	}
}

func Test_QuietAndNoFileParts(t *testing.T) {
	got := validParams(2, 2, true, "")
	if got != ErrNoOutputOptions {
		t.Errorf("Expected quiet and no fileparts to fail validation")
	}
}

func Test_QuietAndFileParts(t *testing.T) {
	got := validParams(2, 2, true, "XXX")
	if got != nil {
		t.Errorf("Expected quiet and fileparts to pass validation")
	}
}

func Test_NotQuietAndFileParts(t *testing.T) {
	got := validParams(2, 2, false, "XXX")
	if got != nil {
		t.Errorf("Expected not quiet and fileparts to pass validation")
	}
}

