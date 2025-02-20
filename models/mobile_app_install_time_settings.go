package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallTimeSettings contains properties used to determine when to offer an app to devices and when to install the app on devices.
type MobileAppInstallTimeSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The time at which the app should be installed.
    deadlineDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time at which the app should be available for installation.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Whether the local device time or UTC time should be used when determining the available and deadline times.
    useLocalTime *bool
}
// NewMobileAppInstallTimeSettings instantiates a new mobileAppInstallTimeSettings and sets the default values.
func NewMobileAppInstallTimeSettings()(*MobileAppInstallTimeSettings) {
    m := &MobileAppInstallTimeSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppInstallTimeSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppInstallTimeSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallTimeSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppInstallTimeSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeadlineDateTime gets the deadlineDateTime property value. The time at which the app should be installed.
func (m *MobileAppInstallTimeSettings) GetDeadlineDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deadlineDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppInstallTimeSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deadlineDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadlineDateTime(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["useLocalTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseLocalTime(val)
        }
        return nil
    }
    return res
}
// GetStartDateTime gets the startDateTime property value. The time at which the app should be available for installation.
func (m *MobileAppInstallTimeSettings) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetUseLocalTime gets the useLocalTime property value. Whether the local device time or UTC time should be used when determining the available and deadline times.
func (m *MobileAppInstallTimeSettings) GetUseLocalTime()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useLocalTime
    }
}
// Serialize serializes information the current object
func (m *MobileAppInstallTimeSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("deadlineDateTime", m.GetDeadlineDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useLocalTime", m.GetUseLocalTime())
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
func (m *MobileAppInstallTimeSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeadlineDateTime sets the deadlineDateTime property value. The time at which the app should be installed.
func (m *MobileAppInstallTimeSettings) SetDeadlineDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deadlineDateTime = value
    }
}
// SetStartDateTime sets the startDateTime property value. The time at which the app should be available for installation.
func (m *MobileAppInstallTimeSettings) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetUseLocalTime sets the useLocalTime property value. Whether the local device time or UTC time should be used when determining the available and deadline times.
func (m *MobileAppInstallTimeSettings) SetUseLocalTime(value *bool)() {
    if m != nil {
        m.useLocalTime = value
    }
}
