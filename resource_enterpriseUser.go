// Do not edit. This file is auto-generated.
package server

// Enterprise User
type EnterpriseUser struct {
    CostCenter     string
    Department     string
    Division       string
    EmployeeNumber string
    ExternalID     string
    ID             string
    Manager        EnterpriseUserManager
    Organization   string
}

// The User's manager. A complex type that optionally allows service providers to represent organizational hierarchy by
// the 'id' attribute of another User.
type EnterpriseUserManager struct {
    Value       string
    Ref         string
    DisplayName string
}
