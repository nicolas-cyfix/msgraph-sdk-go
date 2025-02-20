package models
import (
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type BroadcastMeetingAudience int

const (
    ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE BroadcastMeetingAudience = iota
    ORGANIZATION_BROADCASTMEETINGAUDIENCE
    EVERYONE_BROADCASTMEETINGAUDIENCE
    UNKNOWNFUTUREVALUE_BROADCASTMEETINGAUDIENCE
)

func (i BroadcastMeetingAudience) String() string {
    return []string{"roleIsAttendee", "organization", "everyone", "unknownFutureValue"}[i]
}
func ParseBroadcastMeetingAudience(v string) (interface{}, error) {
    result := ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE
    switch v {
        case "roleIsAttendee":
            result = ROLEISATTENDEE_BROADCASTMEETINGAUDIENCE
        case "organization":
            result = ORGANIZATION_BROADCASTMEETINGAUDIENCE
        case "everyone":
            result = EVERYONE_BROADCASTMEETINGAUDIENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BROADCASTMEETINGAUDIENCE
        default:
            return 0, errors.New("Unknown BroadcastMeetingAudience value: " + v)
    }
    return &result, nil
}
func SerializeBroadcastMeetingAudience(values []BroadcastMeetingAudience) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
