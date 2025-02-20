package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalDomainName provides operations to manage the admin singleton.
type ExternalDomainName struct {
    Entity
}
// NewExternalDomainName instantiates a new externalDomainName and sets the default values.
func NewExternalDomainName()(*ExternalDomainName) {
    m := &ExternalDomainName{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExternalDomainNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalDomainNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalDomainName(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalDomainName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExternalDomainName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
