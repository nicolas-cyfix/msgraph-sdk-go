package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomerBase booking entities that provide a display name.
type BookingCustomerBase struct {
    Entity
}
// NewBookingCustomerBase instantiates a new bookingCustomerBase and sets the default values.
func NewBookingCustomerBase()(*BookingCustomerBase) {
    m := &BookingCustomerBase{
        Entity: *NewEntity(),
    }
    typeValue := "#microsoft.graph.bookingCustomerBase";
    m.SetType(&typeValue);
    return m
}
// CreateBookingCustomerBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.bookingCustomer":
                        return NewBookingCustomer(), nil
                }
            }
        }
    }
    return NewBookingCustomerBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomerBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *BookingCustomerBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
