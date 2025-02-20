package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type Win32LobAppDeliveryOptimizationPriority int

const (
    // Not configured or background normal delivery optimization priority.
    NOTCONFIGURED_WIN32LOBAPPDELIVERYOPTIMIZATIONPRIORITY Win32LobAppDeliveryOptimizationPriority = iota
    // Foreground delivery optimization priority.
    FOREGROUND_WIN32LOBAPPDELIVERYOPTIMIZATIONPRIORITY
)

func (i Win32LobAppDeliveryOptimizationPriority) String() string {
    return []string{"notConfigured", "foreground"}[i]
}
func ParseWin32LobAppDeliveryOptimizationPriority(v string) (interface{}, error) {
    result := NOTCONFIGURED_WIN32LOBAPPDELIVERYOPTIMIZATIONPRIORITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPDELIVERYOPTIMIZATIONPRIORITY
        case "foreground":
            result = FOREGROUND_WIN32LOBAPPDELIVERYOPTIMIZATIONPRIORITY
        default:
            return 0, errors.New("Unknown Win32LobAppDeliveryOptimizationPriority value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppDeliveryOptimizationPriority(values []Win32LobAppDeliveryOptimizationPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
