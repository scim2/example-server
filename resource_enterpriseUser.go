// Do not edit. This file is auto-generated.
package server

// Enterprise User
type EnterpriseUser struct {
    EmployeeNumber string
    CostCenter     string
    Organization   string
    Division       string
    Department     string
    Manager        EnterpriseUserManager
}

// The User's manager. A complex type that optionally allows service providers to represent organizational hierarchy by
// the 'id' attribute of another User.
type EnterpriseUserManager struct {
    Value       string
    Ref         string
    DisplayName string
}
