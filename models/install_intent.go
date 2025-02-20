package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type InstallIntent int

const (
    // Available install intent.
    AVAILABLE_INSTALLINTENT InstallIntent = iota
    // Required install intent.
    REQUIRED_INSTALLINTENT
    // Uninstall install intent.
    UNINSTALL_INSTALLINTENT
    // Available without enrollment install intent.
    AVAILABLEWITHOUTENROLLMENT_INSTALLINTENT
)

func (i InstallIntent) String() string {
    return []string{"available", "required", "uninstall", "availableWithoutEnrollment"}[i]
}
func ParseInstallIntent(v string) (interface{}, error) {
    result := AVAILABLE_INSTALLINTENT
    switch v {
        case "available":
            result = AVAILABLE_INSTALLINTENT
        case "required":
            result = REQUIRED_INSTALLINTENT
        case "uninstall":
            result = UNINSTALL_INSTALLINTENT
        case "availableWithoutEnrollment":
            result = AVAILABLEWITHOUTENROLLMENT_INSTALLINTENT
        default:
            return 0, errors.New("Unknown InstallIntent value: " + v)
    }
    return &result, nil
}
func SerializeInstallIntent(values []InstallIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
