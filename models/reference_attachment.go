package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceAttachment 
type ReferenceAttachment struct {
    Attachment
}
// NewReferenceAttachment instantiates a new ReferenceAttachment and sets the default values.
func NewReferenceAttachment()(*ReferenceAttachment) {
    m := &ReferenceAttachment{
        Attachment: *NewAttachment(),
    }
    return m
}
// CreateReferenceAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferenceAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferenceAttachment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferenceAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Attachment.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ReferenceAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Attachment.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
