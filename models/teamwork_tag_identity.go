package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTagIdentity 
type TeamworkTagIdentity struct {
    Identity
}
// NewTeamworkTagIdentity instantiates a new TeamworkTagIdentity and sets the default values.
func NewTeamworkTagIdentity()(*TeamworkTagIdentity) {
    m := &TeamworkTagIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateTeamworkTagIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTagIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTagIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTagIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TeamworkTagIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
