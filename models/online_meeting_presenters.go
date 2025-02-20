package models
import (
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type OnlineMeetingPresenters int

const (
    EVERYONE_ONLINEMEETINGPRESENTERS OnlineMeetingPresenters = iota
    ORGANIZATION_ONLINEMEETINGPRESENTERS
    ROLEISPRESENTER_ONLINEMEETINGPRESENTERS
    ORGANIZER_ONLINEMEETINGPRESENTERS
    UNKNOWNFUTUREVALUE_ONLINEMEETINGPRESENTERS
)

func (i OnlineMeetingPresenters) String() string {
    return []string{"everyone", "organization", "roleIsPresenter", "organizer", "unknownFutureValue"}[i]
}
func ParseOnlineMeetingPresenters(v string) (interface{}, error) {
    result := EVERYONE_ONLINEMEETINGPRESENTERS
    switch v {
        case "everyone":
            result = EVERYONE_ONLINEMEETINGPRESENTERS
        case "organization":
            result = ORGANIZATION_ONLINEMEETINGPRESENTERS
        case "roleIsPresenter":
            result = ROLEISPRESENTER_ONLINEMEETINGPRESENTERS
        case "organizer":
            result = ORGANIZER_ONLINEMEETINGPRESENTERS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGPRESENTERS
        default:
            return 0, errors.New("Unknown OnlineMeetingPresenters value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingPresenters(values []OnlineMeetingPresenters) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
