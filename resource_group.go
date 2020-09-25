// Do not edit. This file is auto-generated.
package server

// Group
type Group struct {
    DisplayName string
    Members     []GroupMember
}

// A list of members of the Group.
type GroupMember struct {
    Value   string
    Ref     string
    Type    string
    Display string
}
