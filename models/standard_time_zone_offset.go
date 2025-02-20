package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StandardTimeZoneOffset 
type StandardTimeZoneOffset struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the nth occurrence of the day of week that the transition from daylight saving time to standard time occurs.
    dayOccurrence *int32
    // Represents the day of the week when the transition from daylight saving time to standard time.
    dayOfWeek *DayOfWeek
    // Represents the month of the year when the transition from daylight saving time to standard time occurs.
    month *int32
    // Represents the time of day when the transition from daylight saving time to standard time occurs.
    time *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The type property
    type_escaped *string
    // Represents how frequently in terms of years the change from daylight saving time to standard time occurs. For example, a value of 0 means every year.
    year *int32
}
// NewStandardTimeZoneOffset instantiates a new standardTimeZoneOffset and sets the default values.
func NewStandardTimeZoneOffset()(*StandardTimeZoneOffset) {
    m := &StandardTimeZoneOffset{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    typeValue := "#microsoft.graph.standardTimeZoneOffset";
    m.SetType(&typeValue);
    return m
}
// CreateStandardTimeZoneOffsetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStandardTimeZoneOffsetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.daylightTimeZoneOffset":
                        return NewDaylightTimeZoneOffset(), nil
                }
            }
        }
    }
    return NewStandardTimeZoneOffset(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StandardTimeZoneOffset) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDayOccurrence gets the dayOccurrence property value. Represents the nth occurrence of the day of week that the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) GetDayOccurrence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dayOccurrence
    }
}
// GetDayOfWeek gets the dayOfWeek property value. Represents the day of the week when the transition from daylight saving time to standard time.
func (m *StandardTimeZoneOffset) GetDayOfWeek()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.dayOfWeek
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StandardTimeZoneOffset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dayOccurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayOccurrence(val)
        }
        return nil
    }
    res["dayOfWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayOfWeek(val.(*DayOfWeek))
        }
        return nil
    }
    res["month"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonth(val)
        }
        return nil
    }
    res["time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTime(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["year"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYear(val)
        }
        return nil
    }
    return res
}
// GetMonth gets the month property value. Represents the month of the year when the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) GetMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.month
    }
}
// GetTime gets the time property value. Represents the time of day when the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) GetTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.time
    }
}
// GetType gets the type property value. The type property
func (m *StandardTimeZoneOffset) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetYear gets the year property value. Represents how frequently in terms of years the change from daylight saving time to standard time occurs. For example, a value of 0 means every year.
func (m *StandardTimeZoneOffset) GetYear()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.year
    }
}
// Serialize serializes information the current object
func (m *StandardTimeZoneOffset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dayOccurrence", m.GetDayOccurrence())
        if err != nil {
            return err
        }
    }
    if m.GetDayOfWeek() != nil {
        cast := (*m.GetDayOfWeek()).String()
        err := writer.WriteStringValue("dayOfWeek", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("month", m.GetMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("time", m.GetTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("year", m.GetYear())
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
func (m *StandardTimeZoneOffset) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDayOccurrence sets the dayOccurrence property value. Represents the nth occurrence of the day of week that the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) SetDayOccurrence(value *int32)() {
    if m != nil {
        m.dayOccurrence = value
    }
}
// SetDayOfWeek sets the dayOfWeek property value. Represents the day of the week when the transition from daylight saving time to standard time.
func (m *StandardTimeZoneOffset) SetDayOfWeek(value *DayOfWeek)() {
    if m != nil {
        m.dayOfWeek = value
    }
}
// SetMonth sets the month property value. Represents the month of the year when the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) SetMonth(value *int32)() {
    if m != nil {
        m.month = value
    }
}
// SetTime sets the time property value. Represents the time of day when the transition from daylight saving time to standard time occurs.
func (m *StandardTimeZoneOffset) SetTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.time = value
    }
}
// SetType sets the type property value. The type property
func (m *StandardTimeZoneOffset) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetYear sets the year property value. Represents how frequently in terms of years the change from daylight saving time to standard time occurs. For example, a value of 0 means every year.
func (m *StandardTimeZoneOffset) SetYear(value *int32)() {
    if m != nil {
        m.year = value
    }
}
