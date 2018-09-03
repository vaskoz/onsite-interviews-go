package closestmanager

import "fmt"

// Employee represents an employee of a company.
type Employee struct {
	id      int
	name    string
	title   string
	level   int
	manager *Employee
}

var errNoCommonManager = fmt.Errorf("No common manager found")

// NoCommonManagerError returns the errNoCommonManager error owned by the package.
func NoCommonManagerError() error {
	return errNoCommonManager
}

// FindClosestManager returns the manager of two given employees.
func FindClosestManager(empMgrMap map[Employee]*Employee, e1, e2 Employee) (Employee, error) {
	managers := make(map[Employee]struct{})
	for mgr, found := empMgrMap[e1]; found && mgr != nil; mgr, found = empMgrMap[e1] {
		managers[*mgr] = struct{}{}
		e1 = *mgr
	}
	for mgr, found := empMgrMap[e2]; found && mgr != nil; mgr, found = empMgrMap[e2] {
		if _, found := managers[*mgr]; found {
			return *mgr, nil
		}
		e2 = *mgr
	}
	return Employee{}, NoCommonManagerError()
}
