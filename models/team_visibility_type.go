package models
import (
    "errors"
)
// Casts the previous resource to group.
type TeamVisibilityType int

const (
    PRIVATE_TEAMVISIBILITYTYPE TeamVisibilityType = iota
    PUBLIC_TEAMVISIBILITYTYPE
    HIDDENMEMBERSHIP_TEAMVISIBILITYTYPE
    UNKNOWNFUTUREVALUE_TEAMVISIBILITYTYPE
)

func (i TeamVisibilityType) String() string {
    return []string{"private", "public", "hiddenMembership", "unknownFutureValue"}[i]
}
func ParseTeamVisibilityType(v string) (interface{}, error) {
    result := PRIVATE_TEAMVISIBILITYTYPE
    switch v {
        case "private":
            result = PRIVATE_TEAMVISIBILITYTYPE
        case "public":
            result = PUBLIC_TEAMVISIBILITYTYPE
        case "hiddenMembership":
            result = HIDDENMEMBERSHIP_TEAMVISIBILITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMVISIBILITYTYPE
        default:
            return 0, errors.New("Unknown TeamVisibilityType value: " + v)
    }
    return &result, nil
}
func SerializeTeamVisibilityType(values []TeamVisibilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
