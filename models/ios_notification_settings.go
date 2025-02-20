package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosNotificationSettings an item describing notification setting.
type IosNotificationSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Notification Settings Alert Type.
    alertType *IosNotificationAlertType
    // Application name to be associated with the bundleID.
    appName *string
    // Indicates whether badges are allowed for this app.
    badgesEnabled *bool
    // Bundle id of app to which to apply these notification settings.
    bundleID *string
    // Indicates whether notifications are allowed for this app.
    enabled *bool
    // Publisher to be associated with the bundleID.
    publisher *string
    // Indicates whether notifications can be shown in notification center.
    showInNotificationCenter *bool
    // Indicates whether notifications can be shown on the lock screen.
    showOnLockScreen *bool
    // Indicates whether sounds are allowed for this app.
    soundsEnabled *bool
}
// NewIosNotificationSettings instantiates a new iosNotificationSettings and sets the default values.
func NewIosNotificationSettings()(*IosNotificationSettings) {
    m := &IosNotificationSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosNotificationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosNotificationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosNotificationSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosNotificationSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlertType gets the alertType property value. Notification Settings Alert Type.
func (m *IosNotificationSettings) GetAlertType()(*IosNotificationAlertType) {
    if m == nil {
        return nil
    } else {
        return m.alertType
    }
}
// GetAppName gets the appName property value. Application name to be associated with the bundleID.
func (m *IosNotificationSettings) GetAppName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appName
    }
}
// GetBadgesEnabled gets the badgesEnabled property value. Indicates whether badges are allowed for this app.
func (m *IosNotificationSettings) GetBadgesEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.badgesEnabled
    }
}
// GetBundleID gets the bundleID property value. Bundle id of app to which to apply these notification settings.
func (m *IosNotificationSettings) GetBundleID()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleID
    }
}
// GetEnabled gets the enabled property value. Indicates whether notifications are allowed for this app.
func (m *IosNotificationSettings) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosNotificationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosNotificationAlertType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertType(val.(*IosNotificationAlertType))
        }
        return nil
    }
    res["appName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["badgesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBadgesEnabled(val)
        }
        return nil
    }
    res["bundleID"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleID(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["showInNotificationCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInNotificationCenter(val)
        }
        return nil
    }
    res["showOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowOnLockScreen(val)
        }
        return nil
    }
    res["soundsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoundsEnabled(val)
        }
        return nil
    }
    return res
}
// GetPublisher gets the publisher property value. Publisher to be associated with the bundleID.
func (m *IosNotificationSettings) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetShowInNotificationCenter gets the showInNotificationCenter property value. Indicates whether notifications can be shown in notification center.
func (m *IosNotificationSettings) GetShowInNotificationCenter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showInNotificationCenter
    }
}
// GetShowOnLockScreen gets the showOnLockScreen property value. Indicates whether notifications can be shown on the lock screen.
func (m *IosNotificationSettings) GetShowOnLockScreen()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showOnLockScreen
    }
}
// GetSoundsEnabled gets the soundsEnabled property value. Indicates whether sounds are allowed for this app.
func (m *IosNotificationSettings) GetSoundsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.soundsEnabled
    }
}
// Serialize serializes information the current object
func (m *IosNotificationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAlertType() != nil {
        cast := (*m.GetAlertType()).String()
        err := writer.WriteStringValue("alertType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("badgesEnabled", m.GetBadgesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bundleID", m.GetBundleID())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showInNotificationCenter", m.GetShowInNotificationCenter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showOnLockScreen", m.GetShowOnLockScreen())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("soundsEnabled", m.GetSoundsEnabled())
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
func (m *IosNotificationSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlertType sets the alertType property value. Notification Settings Alert Type.
func (m *IosNotificationSettings) SetAlertType(value *IosNotificationAlertType)() {
    if m != nil {
        m.alertType = value
    }
}
// SetAppName sets the appName property value. Application name to be associated with the bundleID.
func (m *IosNotificationSettings) SetAppName(value *string)() {
    if m != nil {
        m.appName = value
    }
}
// SetBadgesEnabled sets the badgesEnabled property value. Indicates whether badges are allowed for this app.
func (m *IosNotificationSettings) SetBadgesEnabled(value *bool)() {
    if m != nil {
        m.badgesEnabled = value
    }
}
// SetBundleID sets the bundleID property value. Bundle id of app to which to apply these notification settings.
func (m *IosNotificationSettings) SetBundleID(value *string)() {
    if m != nil {
        m.bundleID = value
    }
}
// SetEnabled sets the enabled property value. Indicates whether notifications are allowed for this app.
func (m *IosNotificationSettings) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetPublisher sets the publisher property value. Publisher to be associated with the bundleID.
func (m *IosNotificationSettings) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetShowInNotificationCenter sets the showInNotificationCenter property value. Indicates whether notifications can be shown in notification center.
func (m *IosNotificationSettings) SetShowInNotificationCenter(value *bool)() {
    if m != nil {
        m.showInNotificationCenter = value
    }
}
// SetShowOnLockScreen sets the showOnLockScreen property value. Indicates whether notifications can be shown on the lock screen.
func (m *IosNotificationSettings) SetShowOnLockScreen(value *bool)() {
    if m != nil {
        m.showOnLockScreen = value
    }
}
// SetSoundsEnabled sets the soundsEnabled property value. Indicates whether sounds are allowed for this app.
func (m *IosNotificationSettings) SetSoundsEnabled(value *bool)() {
    if m != nil {
        m.soundsEnabled = value
    }
}
