package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type RatingNewZealandTelevisionType int

const (
    // Default value, allow all TV shows content
    ALLALLOWED_RATINGNEWZEALANDTELEVISIONTYPE RatingNewZealandTelevisionType = iota
    // Do not allow any TV shows content
    ALLBLOCKED_RATINGNEWZEALANDTELEVISIONTYPE
    // The G classification excludes materials likely to harm children under 14
    GENERAL_RATINGNEWZEALANDTELEVISIONTYPE
    // The PGR classification encourages parents and guardians to supervise younger viewers
    PARENTALGUIDANCE_RATINGNEWZEALANDTELEVISIONTYPE
    // The AO classification is not suitable for children
    ADULTS_RATINGNEWZEALANDTELEVISIONTYPE
)

func (i RatingNewZealandTelevisionType) String() string {
    return []string{"allAllowed", "allBlocked", "general", "parentalGuidance", "adults"}[i]
}
func ParseRatingNewZealandTelevisionType(v string) (interface{}, error) {
    result := ALLALLOWED_RATINGNEWZEALANDTELEVISIONTYPE
    switch v {
        case "allAllowed":
            result = ALLALLOWED_RATINGNEWZEALANDTELEVISIONTYPE
        case "allBlocked":
            result = ALLBLOCKED_RATINGNEWZEALANDTELEVISIONTYPE
        case "general":
            result = GENERAL_RATINGNEWZEALANDTELEVISIONTYPE
        case "parentalGuidance":
            result = PARENTALGUIDANCE_RATINGNEWZEALANDTELEVISIONTYPE
        case "adults":
            result = ADULTS_RATINGNEWZEALANDTELEVISIONTYPE
        default:
            return 0, errors.New("Unknown RatingNewZealandTelevisionType value: " + v)
    }
    return &result, nil
}
func SerializeRatingNewZealandTelevisionType(values []RatingNewZealandTelevisionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
