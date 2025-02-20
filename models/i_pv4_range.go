package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IPv4Range 
type IPv4Range struct {
    IpRange
    // Lower address.
    lowerAddress *string
    // Upper address.
    upperAddress *string
}
// NewIPv4Range instantiates a new IPv4Range and sets the default values.
func NewIPv4Range()(*IPv4Range) {
    m := &IPv4Range{
        IpRange: *NewIpRange(),
    }
    return m
}
// CreateIPv4RangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIPv4RangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIPv4Range(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IPv4Range) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IpRange.GetFieldDeserializers()
    res["lowerAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowerAddress(val)
        }
        return nil
    }
    res["upperAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpperAddress(val)
        }
        return nil
    }
    return res
}
// GetLowerAddress gets the lowerAddress property value. Lower address.
func (m *IPv4Range) GetLowerAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lowerAddress
    }
}
// GetUpperAddress gets the upperAddress property value. Upper address.
func (m *IPv4Range) GetUpperAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upperAddress
    }
}
// Serialize serializes information the current object
func (m *IPv4Range) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IpRange.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("lowerAddress", m.GetLowerAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upperAddress", m.GetUpperAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLowerAddress sets the lowerAddress property value. Lower address.
func (m *IPv4Range) SetLowerAddress(value *string)() {
    if m != nil {
        m.lowerAddress = value
    }
}
// SetUpperAddress sets the upperAddress property value. Upper address.
func (m *IPv4Range) SetUpperAddress(value *string)() {
    if m != nil {
        m.upperAddress = value
    }
}
