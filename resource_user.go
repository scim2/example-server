// Do not edit. This file is auto-generated.
package server

// User Account
type User struct {
    Active            bool
    Addresses         []UserAddress
    DisplayName       string
    Emails            []UserEmail
    Entitlements      []UserEntitlement
    ExternalID        string
    Groups            []UserGroup
    ID                string
    Ims               []UserIm
    Locale            string
    Name              UserName
    NickName          string
    PassworD          string
    PhoneNumbers      []UserPhoneNumber
    Photos            []UserPhoto
    PreferredLanguage string
    ProfileUrl        string
    Roles             []UserRole
    Timezone          string
    Title             string
    UserName          string
    UserType          string
    X509Certificates  []UserX509Certificate
}

// A physical mailing address for this User. Canonical type values of 'work', 'home', and 'other'. This attribute is a
// type with the following sub-attributes.
type UserAddress struct {
    FormatteD     string
    StreetAddress string
    Locality      string
    Region        string
    PostalCode    string
    Country       string
    Type          string
}

// Email addresses for the user. The value SHOULD be canonicalized by the service provider, e.g., 'bjensen@example.com'
// of 'bjensen@EXAMPLE.COM'. Canonical type values of 'work', 'home', and 'other'.
type UserEmail struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// A list of entitlements for the User that represent a thing the User has.
type UserEntitlement struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// A list of groups to which the user belongs, either through direct membership, through nested groups, or dynamically
type UserGroup struct {
    Value   string
    Ref     string
    Display string
    Type    string
}

// Instant messaging addresses for the User.
type UserIm struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// The components of the user's real name. Providers MAY return just the full name as a single string in the formatted
// or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both.
// both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the
// attributes should be combined.
type UserName struct {
    FormatteD       string
    FamilyName      string
    GivenName       string
    MiddleName      string
    HonorificPrefix string
    HonorificSuffix string
}

// Phone numbers for the User. The value SHOULD be canonicalized by the service provider according to the format
// in RFC 3966, e.g., 'tel:+1-201-555-0123'. Canonical type values of 'work', 'home', 'mobile', 'fax', 'pager', and
type UserPhoneNumber struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// URLs of photos of the User.
type UserPhoto struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// A list of roles for the User that collectively represent who the User is, e.g., 'Student', 'Faculty'.
type UserRole struct {
    Value   string
    Display string
    Type    string
    Primary bool
}

// A list of certificates issued to the User.
type UserX509Certificate struct {
    Value   string
    Display string
    Type    string
    Primary bool
}
