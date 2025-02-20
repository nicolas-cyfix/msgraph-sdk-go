package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOnPremisesInfo 
type EducationOnPremisesInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Unique identifier for the user object in Active Directory.
    immutableId *string
}
// NewEducationOnPremisesInfo instantiates a new educationOnPremisesInfo and sets the default values.
func NewEducationOnPremisesInfo()(*EducationOnPremisesInfo) {
    m := &EducationOnPremisesInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationOnPremisesInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationOnPremisesInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationOnPremisesInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationOnPremisesInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationOnPremisesInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["immutableId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImmutableId(val)
        }
        return nil
    }
    return res
}
// GetImmutableId gets the immutableId property value. Unique identifier for the user object in Active Directory.
func (m *EducationOnPremisesInfo) GetImmutableId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.immutableId
    }
}
// Serialize serializes information the current object
func (m *EducationOnPremisesInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("immutableId", m.GetImmutableId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationOnPremisesInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetImmutableId sets the immutableId property value. Unique identifier for the user object in Active Directory.
func (m *EducationOnPremisesInfo) SetImmutableId(value *string)() {
    if m != nil {
        m.immutableId = value
    }
}
