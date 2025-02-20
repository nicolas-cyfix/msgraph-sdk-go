package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IncomingCallOptions 
type IncomingCallOptions struct {
    CallOptions
}
// NewIncomingCallOptions instantiates a new IncomingCallOptions and sets the default values.
func NewIncomingCallOptions()(*IncomingCallOptions) {
    m := &IncomingCallOptions{
        CallOptions: *NewCallOptions(),
    }
    return m
}
// CreateIncomingCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIncomingCallOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIncomingCallOptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IncomingCallOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CallOptions.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IncomingCallOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CallOptions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
