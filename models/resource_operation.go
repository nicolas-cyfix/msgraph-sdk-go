package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceOperation 
type ResourceOperation struct {
    Entity
    // Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
    actionName *string
    // Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
    description *string
    // Name of the Resource this operation is performed on.
    resourceName *string
}
// NewResourceOperation instantiates a new ResourceOperation and sets the default values.
func NewResourceOperation()(*ResourceOperation) {
    m := &ResourceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateResourceOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceOperation(), nil
}
// GetActionName gets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// GetDescription gets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
// GetResourceName gets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// Serialize serializes information the current object
func (m *ResourceOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionName sets the actionName property value. Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
func (m *ResourceOperation) SetActionName(value *string)() {
    if m != nil {
        m.actionName = value
    }
}
// SetDescription sets the description property value. Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
func (m *ResourceOperation) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetResourceName sets the resourceName property value. Name of the Resource this operation is performed on.
func (m *ResourceOperation) SetResourceName(value *string)() {
    if m != nil {
        m.resourceName = value
    }
}
