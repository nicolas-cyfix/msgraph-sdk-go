package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingWorkTimeSlot 
type BookingWorkTimeSlot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The time of the day when work stops. For example, 17:00:00.0000000.
    endTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The time of the day when work starts. For example, 08:00:00.0000000.
    startTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewBookingWorkTimeSlot instantiates a new bookingWorkTimeSlot and sets the default values.
func NewBookingWorkTimeSlot()(*BookingWorkTimeSlot) {
    m := &BookingWorkTimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingWorkTimeSlotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingWorkTimeSlotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingWorkTimeSlot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkTimeSlot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEndTime gets the endTime property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) GetEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingWorkTimeSlot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    return res
}
// GetStartTime gets the startTime property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) GetStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// Serialize serializes information the current object
func (m *BookingWorkTimeSlot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeOnlyValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("startTime", m.GetStartTime())
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
func (m *BookingWorkTimeSlot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEndTime sets the endTime property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) SetEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.endTime = value
    }
}
// SetStartTime sets the startTime property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) SetStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.startTime = value
    }
}
