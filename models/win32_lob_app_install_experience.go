package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppInstallExperience contains installation experience properties for a Win32 App
type Win32LobAppInstallExperience struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates the type of restart action.
    deviceRestartBehavior *Win32LobAppRestartBehavior
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
}
// NewWin32LobAppInstallExperience instantiates a new win32LobAppInstallExperience and sets the default values.
func NewWin32LobAppInstallExperience()(*Win32LobAppInstallExperience) {
    m := &Win32LobAppInstallExperience{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWin32LobAppInstallExperienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppInstallExperienceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppInstallExperience(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppInstallExperience) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceRestartBehavior gets the deviceRestartBehavior property value. Indicates the type of restart action.
func (m *Win32LobAppInstallExperience) GetDeviceRestartBehavior()(*Win32LobAppRestartBehavior) {
    if m == nil {
        return nil
    } else {
        return m.deviceRestartBehavior
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppInstallExperience) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceRestartBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRestartBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRestartBehavior(val.(*Win32LobAppRestartBehavior))
        }
        return nil
    }
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    return res
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppInstallExperience) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// Serialize serializes information the current object
func (m *Win32LobAppInstallExperience) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceRestartBehavior() != nil {
        cast := (*m.GetDeviceRestartBehavior()).String()
        err := writer.WriteStringValue("deviceRestartBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err := writer.WriteStringValue("runAsAccount", &cast)
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
func (m *Win32LobAppInstallExperience) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceRestartBehavior sets the deviceRestartBehavior property value. Indicates the type of restart action.
func (m *Win32LobAppInstallExperience) SetDeviceRestartBehavior(value *Win32LobAppRestartBehavior)() {
    if m != nil {
        m.deviceRestartBehavior = value
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppInstallExperience) SetRunAsAccount(value *RunAsAccountType)() {
    if m != nil {
        m.runAsAccount = value
    }
}
