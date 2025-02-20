package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type IosNotificationAlertType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_IOSNOTIFICATIONALERTTYPE IosNotificationAlertType = iota
    // Banner.
    BANNER_IOSNOTIFICATIONALERTTYPE
    // Modal.
    MODAL_IOSNOTIFICATIONALERTTYPE
    // None.
    NONE_IOSNOTIFICATIONALERTTYPE
)

func (i IosNotificationAlertType) String() string {
    return []string{"deviceDefault", "banner", "modal", "none"}[i]
}
func ParseIosNotificationAlertType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_IOSNOTIFICATIONALERTTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_IOSNOTIFICATIONALERTTYPE
        case "banner":
            result = BANNER_IOSNOTIFICATIONALERTTYPE
        case "modal":
            result = MODAL_IOSNOTIFICATIONALERTTYPE
        case "none":
            result = NONE_IOSNOTIFICATIONALERTTYPE
        default:
            return 0, errors.New("Unknown IosNotificationAlertType value: " + v)
    }
    return &result, nil
}
func SerializeIosNotificationAlertType(values []IosNotificationAlertType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
