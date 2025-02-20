package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataSubjectable 
type DataSubjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetResidency()(*string)
    SetEmail(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetResidency(value *string)()
}
