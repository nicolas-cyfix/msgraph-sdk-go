package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetApplicationOwners 
type TargetApplicationOwners struct {
    SubjectSet
}
// NewTargetApplicationOwners instantiates a new TargetApplicationOwners and sets the default values.
func NewTargetApplicationOwners()(*TargetApplicationOwners) {
    m := &TargetApplicationOwners{
        SubjectSet: *NewSubjectSet(),
    }
    return m
}
// CreateTargetApplicationOwnersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetApplicationOwnersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetApplicationOwners(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetApplicationOwners) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TargetApplicationOwners) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
