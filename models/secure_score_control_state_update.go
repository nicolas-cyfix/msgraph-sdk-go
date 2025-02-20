package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScoreControlStateUpdate 
type SecureScoreControlStateUpdate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Assigns the control to the user who will take the action.
    assignedTo *string
    // Provides optional comment about the control.
    comment *string
    // State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
    state *string
    // ID of the user who updated tenant state.
    updatedBy *string
    // Time at which the control state was updated.
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewSecureScoreControlStateUpdate instantiates a new secureScoreControlStateUpdate and sets the default values.
func NewSecureScoreControlStateUpdate()(*SecureScoreControlStateUpdate) {
    m := &SecureScoreControlStateUpdate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecureScoreControlStateUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreControlStateUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScoreControlStateUpdate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecureScoreControlStateUpdate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignedTo gets the assignedTo property value. Assigns the control to the user who will take the action.
func (m *SecureScoreControlStateUpdate) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetComment gets the comment property value. Provides optional comment about the control.
func (m *SecureScoreControlStateUpdate) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScoreControlStateUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["updatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedBy(val)
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetState gets the state property value. State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
func (m *SecureScoreControlStateUpdate) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUpdatedBy gets the updatedBy property value. ID of the user who updated tenant state.
func (m *SecureScoreControlStateUpdate) GetUpdatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.updatedBy
    }
}
// GetUpdatedDateTime gets the updatedDateTime property value. Time at which the control state was updated.
func (m *SecureScoreControlStateUpdate) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.updatedDateTime
    }
}
// Serialize serializes information the current object
func (m *SecureScoreControlStateUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updatedBy", m.GetUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
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
func (m *SecureScoreControlStateUpdate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignedTo sets the assignedTo property value. Assigns the control to the user who will take the action.
func (m *SecureScoreControlStateUpdate) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetComment sets the comment property value. Provides optional comment about the control.
func (m *SecureScoreControlStateUpdate) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetState sets the state property value. State of the control, which can be modified via a PATCH command (for example, ignored, thirdParty).
func (m *SecureScoreControlStateUpdate) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetUpdatedBy sets the updatedBy property value. ID of the user who updated tenant state.
func (m *SecureScoreControlStateUpdate) SetUpdatedBy(value *string)() {
    if m != nil {
        m.updatedBy = value
    }
}
// SetUpdatedDateTime sets the updatedDateTime property value. Time at which the control state was updated.
func (m *SecureScoreControlStateUpdate) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.updatedDateTime = value
    }
}
