package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type StateManagementSetting int

const (
    // Not configured.
    NOTCONFIGURED_STATEMANAGEMENTSETTING StateManagementSetting = iota
    // Blocked.
    BLOCKED_STATEMANAGEMENTSETTING
    // Allowed.
    ALLOWED_STATEMANAGEMENTSETTING
)

func (i StateManagementSetting) String() string {
    return []string{"notConfigured", "blocked", "allowed"}[i]
}
func ParseStateManagementSetting(v string) (interface{}, error) {
    result := NOTCONFIGURED_STATEMANAGEMENTSETTING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_STATEMANAGEMENTSETTING
        case "blocked":
            result = BLOCKED_STATEMANAGEMENTSETTING
        case "allowed":
            result = ALLOWED_STATEMANAGEMENTSETTING
        default:
            return 0, errors.New("Unknown StateManagementSetting value: " + v)
    }
    return &result, nil
}
func SerializeStateManagementSetting(values []StateManagementSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
