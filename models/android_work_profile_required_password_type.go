package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type AndroidWorkProfileRequiredPasswordType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE AndroidWorkProfileRequiredPasswordType = iota
    // Low security biometrics based password required.
    LOWSECURITYBIOMETRIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // Required.
    REQUIRED_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // At least numeric password required.
    ATLEASTNUMERIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // Numeric complex password required.
    NUMERICCOMPLEX_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // At least alphabetic password required.
    ATLEASTALPHABETIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // At least alphanumeric password required.
    ATLEASTALPHANUMERIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    // At least alphanumeric with symbols password required.
    ALPHANUMERICWITHSYMBOLS_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
)

func (i AndroidWorkProfileRequiredPasswordType) String() string {
    return []string{"deviceDefault", "lowSecurityBiometric", "required", "atLeastNumeric", "numericComplex", "atLeastAlphabetic", "atLeastAlphanumeric", "alphanumericWithSymbols"}[i]
}
func ParseAndroidWorkProfileRequiredPasswordType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "lowSecurityBiometric":
            result = LOWSECURITYBIOMETRIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "required":
            result = REQUIRED_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "atLeastNumeric":
            result = ATLEASTNUMERIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "numericComplex":
            result = NUMERICCOMPLEX_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "atLeastAlphabetic":
            result = ATLEASTALPHABETIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "atLeastAlphanumeric":
            result = ATLEASTALPHANUMERIC_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        case "alphanumericWithSymbols":
            result = ALPHANUMERICWITHSYMBOLS_ANDROIDWORKPROFILEREQUIREDPASSWORDTYPE
        default:
            return 0, errors.New("Unknown AndroidWorkProfileRequiredPasswordType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidWorkProfileRequiredPasswordType(values []AndroidWorkProfileRequiredPasswordType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
