package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteEntityBaseModel provides operations to manage the admin singleton.
type OnenoteEntityBaseModel struct {
    Entity
    // The endpoint where you can get details about the page. Read-only.
    self *string
}
// NewOnenoteEntityBaseModel instantiates a new onenoteEntityBaseModel and sets the default values.
func NewOnenoteEntityBaseModel()(*OnenoteEntityBaseModel) {
    m := &OnenoteEntityBaseModel{
        Entity: *NewEntity(),
    }
    typeValue := "#microsoft.graph.onenoteEntityBaseModel";
    m.SetType(&typeValue);
    return m
}
// CreateOnenoteEntityBaseModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteEntityBaseModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.onenoteEntitySchemaObjectModel":
                        return NewOnenoteEntitySchemaObjectModel(), nil
                    case "#microsoft.graph.onenoteResource":
                        return NewOnenoteResource(), nil
                }
            }
        }
    }
    return NewOnenoteEntityBaseModel(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteEntityBaseModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
        }
        return nil
    }
    return res
}
// GetSelf gets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
// Serialize serializes information the current object
func (m *OnenoteEntityBaseModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("self", m.GetSelf())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSelf sets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) SetSelf(value *string)() {
    if m != nil {
        m.self = value
    }
}
