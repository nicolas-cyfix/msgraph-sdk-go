package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AudioRoutingGroup provides operations to manage the cloudCommunications singleton.
type AudioRoutingGroup struct {
    Entity
    // List of receiving participant ids.
    receivers []string
    // The routingMode property
    routingMode *RoutingMode
    // List of source participant ids.
    sources []string
}
// NewAudioRoutingGroup instantiates a new audioRoutingGroup and sets the default values.
func NewAudioRoutingGroup()(*AudioRoutingGroup) {
    m := &AudioRoutingGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAudioRoutingGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAudioRoutingGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAudioRoutingGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AudioRoutingGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["receivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReceivers(res)
        }
        return nil
    }
    res["routingMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingMode(val.(*RoutingMode))
        }
        return nil
    }
    res["sources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSources(res)
        }
        return nil
    }
    return res
}
// GetReceivers gets the receivers property value. List of receiving participant ids.
func (m *AudioRoutingGroup) GetReceivers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.receivers
    }
}
// GetRoutingMode gets the routingMode property value. The routingMode property
func (m *AudioRoutingGroup) GetRoutingMode()(*RoutingMode) {
    if m == nil {
        return nil
    } else {
        return m.routingMode
    }
}
// GetSources gets the sources property value. List of source participant ids.
func (m *AudioRoutingGroup) GetSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// Serialize serializes information the current object
func (m *AudioRoutingGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetReceivers() != nil {
        err = writer.WriteCollectionOfStringValues("receivers", m.GetReceivers())
        if err != nil {
            return err
        }
    }
    if m.GetRoutingMode() != nil {
        cast := (*m.GetRoutingMode()).String()
        err = writer.WriteStringValue("routingMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        err = writer.WriteCollectionOfStringValues("sources", m.GetSources())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReceivers sets the receivers property value. List of receiving participant ids.
func (m *AudioRoutingGroup) SetReceivers(value []string)() {
    if m != nil {
        m.receivers = value
    }
}
// SetRoutingMode sets the routingMode property value. The routingMode property
func (m *AudioRoutingGroup) SetRoutingMode(value *RoutingMode)() {
    if m != nil {
        m.routingMode = value
    }
}
// SetSources sets the sources property value. List of source participant ids.
func (m *AudioRoutingGroup) SetSources(value []string)() {
    if m != nil {
        m.sources = value
    }
}
