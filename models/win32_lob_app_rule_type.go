package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type Win32LobAppRuleType int

const (
    // Detection rule.
    DETECTION_WIN32LOBAPPRULETYPE Win32LobAppRuleType = iota
    // Requirement rule.
    REQUIREMENT_WIN32LOBAPPRULETYPE
)

func (i Win32LobAppRuleType) String() string {
    return []string{"detection", "requirement"}[i]
}
func ParseWin32LobAppRuleType(v string) (interface{}, error) {
    result := DETECTION_WIN32LOBAPPRULETYPE
    switch v {
        case "detection":
            result = DETECTION_WIN32LOBAPPRULETYPE
        case "requirement":
            result = REQUIREMENT_WIN32LOBAPPRULETYPE
        default:
            return 0, errors.New("Unknown Win32LobAppRuleType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppRuleType(values []Win32LobAppRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
