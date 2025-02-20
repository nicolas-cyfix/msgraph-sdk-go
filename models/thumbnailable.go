package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Thumbnailable 
type Thumbnailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetHeight()(*int32)
    GetSourceItemId()(*string)
    GetUrl()(*string)
    GetWidth()(*int32)
    SetContent(value []byte)()
    SetHeight(value *int32)()
    SetSourceItemId(value *string)()
    SetUrl(value *string)()
    SetWidth(value *int32)()
}
