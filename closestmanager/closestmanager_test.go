package closestmanager

import (
	"testing"
)

var employees = map[string]*Employee{
	"ceo":          &Employee{100, "big", "boss", 11, nil},
	"vpoe":         &Employee{12, "big", "tech", 10, nil},
	"opsDirector":  &Employee{11, "ops", "dude", 9, nil},
	"softDirector": &Employee{533, "sw", "peeps", 9, nil},
	"swMgr":        &Employee{80, "me", "you", 8, nil},
	"otherMgr":     &Employee{5, "other", "guy", 8, nil},
	"vpop":         &Employee{54, "prod", "uct guy", 10, nil},
}

func createTestCompanyMap() map[Employee]*Employee {
	result := make(map[Employee]*Employee)
	result[*employees["ceo"]] = nil
	// engineering
	result[*employees["vpoe"]] = employees["ceo"]
	result[*employees["opsDirector"]] = employees["vpoe"]
	result[*employees["softDirector"]] = employees["vpoe"]
	result[*employees["swMgr"]] = employees["softDirector"]
	result[*employees["otherMgr"]] = employees["softDirector"]
	result[*employees["otherMgr"]] = employees["softDirector"]
	// product
	result[*employees["vpop"]] = employees["ceo"]
	return result
}

var testcases = []struct {
	e1, e2      *Employee
	closestMgr  *Employee
	expectedErr error
}{
	{employees["vpoe"], employees["vpop"], employees["ceo"], nil},
	{employees["opsDirector"], employees["swMgr"], employees["vpoe"], nil},
	{employees["otherMgr"], employees["swMgr"], employees["softDirector"], nil},
	{employees["vpoe"], employees["ceo"], &Employee{}, getNoCommonManagerError()},
}

func TestFindClosestManagerMap(t *testing.T) {
	t.Parallel()
	empMgrMap := createTestCompanyMap()
	for _, tc := range testcases {
		if mgr, err := FindClosestManager(empMgrMap, *tc.e1, *tc.e2); err != tc.expectedErr {
			t.Errorf("Expected error %v but got error %v", tc.expectedErr, err)
		} else if mgr != *tc.closestMgr {
			t.Errorf("Expected %v but got %v", tc.closestMgr, mgr)
		}
	}
}

func BenchmarkFindClosestManagerMap(b *testing.B) {
	empMgrMap := createTestCompanyMap()
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindClosestManager(empMgrMap, *tc.e1, *tc.e2)
		}
	}
}
