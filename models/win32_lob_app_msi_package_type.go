package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type Win32LobAppMsiPackageType int

const (
    // Indicates a per-machine app package.
    PERMACHINE_WIN32LOBAPPMSIPACKAGETYPE Win32LobAppMsiPackageType = iota
    // Indicates a per-user app package.
    PERUSER_WIN32LOBAPPMSIPACKAGETYPE
    // Indicates a dual-purpose app package.
    DUALPURPOSE_WIN32LOBAPPMSIPACKAGETYPE
)

func (i Win32LobAppMsiPackageType) String() string {
    return []string{"perMachine", "perUser", "dualPurpose"}[i]
}
func ParseWin32LobAppMsiPackageType(v string) (interface{}, error) {
    result := PERMACHINE_WIN32LOBAPPMSIPACKAGETYPE
    switch v {
        case "perMachine":
            result = PERMACHINE_WIN32LOBAPPMSIPACKAGETYPE
        case "perUser":
            result = PERUSER_WIN32LOBAPPMSIPACKAGETYPE
        case "dualPurpose":
            result = DUALPURPOSE_WIN32LOBAPPMSIPACKAGETYPE
        default:
            return 0, errors.New("Unknown Win32LobAppMsiPackageType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppMsiPackageType(values []Win32LobAppMsiPackageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
