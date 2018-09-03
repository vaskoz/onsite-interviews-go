package closestmanager

import (
	"testing"
)

var employeesForMap = map[string]*Employee{
	"ceo":          {100, "big", "boss", 11, nil},
	"vpoe":         {12, "big", "tech", 10, nil},
	"opsDirector":  {11, "ops", "dude", 9, nil},
	"softDirector": {533, "sw", "peeps", 9, nil},
	"swMgr":        {80, "me", "you", 8, nil},
	"otherMgr":     {5, "other", "guy", 8, nil},
	"vpop":         {54, "prod", "uct guy", 10, nil},
}

func createTestCompanyMap() map[Employee]*Employee {
	result := make(map[Employee]*Employee)
	result[*employeesForMap["ceo"]] = nil
	// engineering
	result[*employeesForMap["vpoe"]] = employeesForMap["ceo"]
	result[*employeesForMap["opsDirector"]] = employeesForMap["vpoe"]
	result[*employeesForMap["softDirector"]] = employeesForMap["vpoe"]
	result[*employeesForMap["swMgr"]] = employeesForMap["softDirector"]
	result[*employeesForMap["otherMgr"]] = employeesForMap["softDirector"]
	result[*employeesForMap["otherMgr"]] = employeesForMap["softDirector"]
	// product
	result[*employeesForMap["vpop"]] = employeesForMap["ceo"]
	return result
}

var testcasesForMap = []struct {
	e1, e2      *Employee
	closestMgr  *Employee
	expectedErr error
}{
	{employeesForMap["vpoe"], employeesForMap["vpop"], employeesForMap["ceo"], nil},
	{employeesForMap["opsDirector"], employeesForMap["swMgr"], employeesForMap["vpoe"], nil},
	{employeesForMap["otherMgr"], employeesForMap["swMgr"], employeesForMap["softDirector"], nil},
	{employeesForMap["vpoe"], employeesForMap["ceo"], &Employee{}, NoCommonManagerError()},
}

func TestFindClosestManagerMap(t *testing.T) {
	t.Parallel()
	empMgrMap := createTestCompanyMap()
	for _, tc := range testcasesForMap {
		if mgr, err := FindClosestManagerMap(empMgrMap, *tc.e1, *tc.e2); err != tc.expectedErr {
			t.Errorf("Expected error %v but got error %v", tc.expectedErr, err)
		} else if mgr != *tc.closestMgr {
			t.Errorf("Expected %v but got %v", tc.closestMgr, mgr)
		}
	}
}

func BenchmarkFindClosestManagerMap(b *testing.B) {
	empMgrMap := createTestCompanyMap()
	for i := 0; i < b.N; i++ {
		for _, tc := range testcasesForMap {
			FindClosestManagerMap(empMgrMap, *tc.e1, *tc.e2)
		}
	}
}

func createTestCompanyGraph() {
	ceo := &Employee{100, "big", "boss", 11, nil}
	employeesForGraph["ceo"] = ceo
	vpoe := &Employee{12, "big", "tech", 10, ceo}
	employeesForGraph["vpoe"] = vpoe
	opsDirector := &Employee{11, "ops", "dude", 9, vpoe}
	employeesForGraph["opsDirector"] = opsDirector
	softDirector := &Employee{533, "sw", "peeps", 9, vpoe}
	employeesForGraph["softDirector"] = softDirector
	swMgr := &Employee{80, "me", "you", 8, softDirector}
	employeesForGraph["swMgr"] = swMgr
	otherMgr := &Employee{5, "other", "guy", 8, softDirector}
	employeesForGraph["otherMgr"] = otherMgr
	vpop := &Employee{54, "prod", "uct guy", 10, ceo}
	employeesForGraph["vpop"] = vpop
}

var employeesForGraph = map[string]*Employee{}

type graphTestCase struct {
	e1, e2      *Employee
	closestMgr  *Employee
	expectedErr error
}

var testcasesForGraph []graphTestCase

func getTestCasesForGraph() {
	testcasesForGraph = append(testcasesForGraph, graphTestCase{employeesForGraph["vpoe"], employeesForGraph["vpop"], employeesForGraph["ceo"], nil})
	testcasesForGraph = append(testcasesForGraph, graphTestCase{employeesForGraph["opsDirector"], employeesForGraph["swMgr"], employeesForGraph["vpoe"], nil})
	testcasesForGraph = append(testcasesForGraph, graphTestCase{employeesForGraph["otherMgr"], employeesForGraph["swMgr"], employeesForGraph["softDirector"], nil})
	testcasesForGraph = append(testcasesForGraph, graphTestCase{employeesForGraph["vpoe"], employeesForGraph["ceo"], nil, NoCommonManagerError()})
}

func TestFindClosestManagerGraph(t *testing.T) {
	t.Parallel()
	createTestCompanyGraph()
	getTestCasesForGraph()
	for _, tc := range testcasesForGraph {
		if mgr, err := FindClosestManagerGraph(tc.e1, tc.e2); err != tc.expectedErr {
			t.Errorf("Expected error %v but got error %v", tc.expectedErr, err)
		} else if mgr != tc.closestMgr {
			t.Errorf("Expected %v but got %v", tc.closestMgr, mgr)
		}
	}
}

func BenchmarkFindClosestManagerGraph(b *testing.B) {
	createTestCompanyGraph()
	getTestCasesForGraph()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testcasesForGraph {
			FindClosestManagerGraph(tc.e1, tc.e2)
		}
	}
}
