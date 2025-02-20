package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type RatingAustraliaTelevisionType int

const (
    // Default value, allow all TV shows content
    ALLALLOWED_RATINGAUSTRALIATELEVISIONTYPE RatingAustraliaTelevisionType = iota
    // Do not allow any TV shows content
    ALLBLOCKED_RATINGAUSTRALIATELEVISIONTYPE
    // The P classification is intended for preschoolers
    PRESCHOOLERS_RATINGAUSTRALIATELEVISIONTYPE
    // The C classification is intended for children under 14
    CHILDREN_RATINGAUSTRALIATELEVISIONTYPE
    // The G classification is suitable for all ages
    GENERAL_RATINGAUSTRALIATELEVISIONTYPE
    // The PG classification is recommended for young viewers
    PARENTALGUIDANCE_RATINGAUSTRALIATELEVISIONTYPE
    // The M classification is recommended for viewers over 15
    MATURE_RATINGAUSTRALIATELEVISIONTYPE
    // The MA15+ classification is not suitable for viewers under 15
    AGESABOVE15_RATINGAUSTRALIATELEVISIONTYPE
    // The AV15+ classification is not suitable for viewers under 15, adult violence-specific
    AGESABOVE15ADULTVIOLENCE_RATINGAUSTRALIATELEVISIONTYPE
)

func (i RatingAustraliaTelevisionType) String() string {
    return []string{"allAllowed", "allBlocked", "preschoolers", "children", "general", "parentalGuidance", "mature", "agesAbove15", "agesAbove15AdultViolence"}[i]
}
func ParseRatingAustraliaTelevisionType(v string) (interface{}, error) {
    result := ALLALLOWED_RATINGAUSTRALIATELEVISIONTYPE
    switch v {
        case "allAllowed":
            result = ALLALLOWED_RATINGAUSTRALIATELEVISIONTYPE
        case "allBlocked":
            result = ALLBLOCKED_RATINGAUSTRALIATELEVISIONTYPE
        case "preschoolers":
            result = PRESCHOOLERS_RATINGAUSTRALIATELEVISIONTYPE
        case "children":
            result = CHILDREN_RATINGAUSTRALIATELEVISIONTYPE
        case "general":
            result = GENERAL_RATINGAUSTRALIATELEVISIONTYPE
        case "parentalGuidance":
            result = PARENTALGUIDANCE_RATINGAUSTRALIATELEVISIONTYPE
        case "mature":
            result = MATURE_RATINGAUSTRALIATELEVISIONTYPE
        case "agesAbove15":
            result = AGESABOVE15_RATINGAUSTRALIATELEVISIONTYPE
        case "agesAbove15AdultViolence":
            result = AGESABOVE15ADULTVIOLENCE_RATINGAUSTRALIATELEVISIONTYPE
        default:
            return 0, errors.New("Unknown RatingAustraliaTelevisionType value: " + v)
    }
    return &result, nil
}
func SerializeRatingAustraliaTelevisionType(values []RatingAustraliaTelevisionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
