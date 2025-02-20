package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentWindowsHelloForBusinessConfiguration 
type DeviceEnrollmentWindowsHelloForBusinessConfiguration struct {
    DeviceEnrollmentConfiguration
    // The enhancedBiometricsState property
    enhancedBiometricsState *Enablement
    // Controls the period of time (in days) that a PIN can be used before the system requires the user to change it. This must be set between 0 and 730, inclusive. If set to 0, the user's PIN will never expire
    pinExpirationInDays *int32
    // Windows Hello for Business pin usage options
    pinLowercaseCharactersUsage *WindowsHelloForBusinessPinUsage
    // Controls the maximum number of characters allowed for the Windows Hello for Business PIN. This value must be between 4 and 127, inclusive. This value must be greater than or equal to the value set for the minimum PIN.
    pinMaximumLength *int32
    // Controls the minimum number of characters required for the Windows Hello for Business PIN.  This value must be between 4 and 127, inclusive, and less than or equal to the value set for the maximum PIN.
    pinMinimumLength *int32
    // Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset.
    pinPreviousBlockCount *int32
    // Windows Hello for Business pin usage options
    pinSpecialCharactersUsage *WindowsHelloForBusinessPinUsage
    // Windows Hello for Business pin usage options
    pinUppercaseCharactersUsage *WindowsHelloForBusinessPinUsage
    // Controls the use of Remote Windows Hello for Business. Remote Windows Hello for Business provides the ability for a portable, registered device to be usable as a companion for desktop authentication. The desktop must be Azure AD joined and the companion device must have a Windows Hello for Business PIN.
    remotePassportEnabled *bool
    // Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
    securityDeviceRequired *bool
    // The state property
    state *Enablement
    // Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
    unlockWithBiometricsEnabled *bool
}
// NewDeviceEnrollmentWindowsHelloForBusinessConfiguration instantiates a new DeviceEnrollmentWindowsHelloForBusinessConfiguration and sets the default values.
func NewDeviceEnrollmentWindowsHelloForBusinessConfiguration()(*DeviceEnrollmentWindowsHelloForBusinessConfiguration) {
    m := &DeviceEnrollmentWindowsHelloForBusinessConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    return m
}
// CreateDeviceEnrollmentWindowsHelloForBusinessConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentWindowsHelloForBusinessConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentWindowsHelloForBusinessConfiguration(), nil
}
// GetEnhancedBiometricsState gets the enhancedBiometricsState property value. The enhancedBiometricsState property
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetEnhancedBiometricsState()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.enhancedBiometricsState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["enhancedBiometricsState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedBiometricsState(val.(*Enablement))
        }
        return nil
    }
    res["pinExpirationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinExpirationInDays(val)
        }
        return nil
    }
    res["pinLowercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsHelloForBusinessPinUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinLowercaseCharactersUsage(val.(*WindowsHelloForBusinessPinUsage))
        }
        return nil
    }
    res["pinMaximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMaximumLength(val)
        }
        return nil
    }
    res["pinMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinMinimumLength(val)
        }
        return nil
    }
    res["pinPreviousBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinPreviousBlockCount(val)
        }
        return nil
    }
    res["pinSpecialCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsHelloForBusinessPinUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinSpecialCharactersUsage(val.(*WindowsHelloForBusinessPinUsage))
        }
        return nil
    }
    res["pinUppercaseCharactersUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsHelloForBusinessPinUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinUppercaseCharactersUsage(val.(*WindowsHelloForBusinessPinUsage))
        }
        return nil
    }
    res["remotePassportEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemotePassportEnabled(val)
        }
        return nil
    }
    res["securityDeviceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDeviceRequired(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*Enablement))
        }
        return nil
    }
    res["unlockWithBiometricsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlockWithBiometricsEnabled(val)
        }
        return nil
    }
    return res
}
// GetPinExpirationInDays gets the pinExpirationInDays property value. Controls the period of time (in days) that a PIN can be used before the system requires the user to change it. This must be set between 0 and 730, inclusive. If set to 0, the user's PIN will never expire
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinExpirationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinExpirationInDays
    }
}
// GetPinLowercaseCharactersUsage gets the pinLowercaseCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinLowercaseCharactersUsage()(*WindowsHelloForBusinessPinUsage) {
    if m == nil {
        return nil
    } else {
        return m.pinLowercaseCharactersUsage
    }
}
// GetPinMaximumLength gets the pinMaximumLength property value. Controls the maximum number of characters allowed for the Windows Hello for Business PIN. This value must be between 4 and 127, inclusive. This value must be greater than or equal to the value set for the minimum PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMaximumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinMaximumLength
    }
}
// GetPinMinimumLength gets the pinMinimumLength property value. Controls the minimum number of characters required for the Windows Hello for Business PIN.  This value must be between 4 and 127, inclusive, and less than or equal to the value set for the maximum PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinMinimumLength
    }
}
// GetPinPreviousBlockCount gets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinPreviousBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pinPreviousBlockCount
    }
}
// GetPinSpecialCharactersUsage gets the pinSpecialCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinSpecialCharactersUsage()(*WindowsHelloForBusinessPinUsage) {
    if m == nil {
        return nil
    } else {
        return m.pinSpecialCharactersUsage
    }
}
// GetPinUppercaseCharactersUsage gets the pinUppercaseCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinUppercaseCharactersUsage()(*WindowsHelloForBusinessPinUsage) {
    if m == nil {
        return nil
    } else {
        return m.pinUppercaseCharactersUsage
    }
}
// GetRemotePassportEnabled gets the remotePassportEnabled property value. Controls the use of Remote Windows Hello for Business. Remote Windows Hello for Business provides the ability for a portable, registered device to be usable as a companion for desktop authentication. The desktop must be Azure AD joined and the companion device must have a Windows Hello for Business PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetRemotePassportEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remotePassportEnabled
    }
}
// GetSecurityDeviceRequired gets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetSecurityDeviceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityDeviceRequired
    }
}
// GetState gets the state property value. The state property
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetState()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUnlockWithBiometricsEnabled gets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetUnlockWithBiometricsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.unlockWithBiometricsEnabled
    }
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEnhancedBiometricsState() != nil {
        cast := (*m.GetEnhancedBiometricsState()).String()
        err = writer.WriteStringValue("enhancedBiometricsState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinExpirationInDays", m.GetPinExpirationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetPinLowercaseCharactersUsage() != nil {
        cast := (*m.GetPinLowercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinLowercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMaximumLength", m.GetPinMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinMinimumLength", m.GetPinMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pinPreviousBlockCount", m.GetPinPreviousBlockCount())
        if err != nil {
            return err
        }
    }
    if m.GetPinSpecialCharactersUsage() != nil {
        cast := (*m.GetPinSpecialCharactersUsage()).String()
        err = writer.WriteStringValue("pinSpecialCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPinUppercaseCharactersUsage() != nil {
        cast := (*m.GetPinUppercaseCharactersUsage()).String()
        err = writer.WriteStringValue("pinUppercaseCharactersUsage", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("remotePassportEnabled", m.GetRemotePassportEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDeviceRequired", m.GetSecurityDeviceRequired())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unlockWithBiometricsEnabled", m.GetUnlockWithBiometricsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnhancedBiometricsState sets the enhancedBiometricsState property value. The enhancedBiometricsState property
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetEnhancedBiometricsState(value *Enablement)() {
    if m != nil {
        m.enhancedBiometricsState = value
    }
}
// SetPinExpirationInDays sets the pinExpirationInDays property value. Controls the period of time (in days) that a PIN can be used before the system requires the user to change it. This must be set between 0 and 730, inclusive. If set to 0, the user's PIN will never expire
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinExpirationInDays(value *int32)() {
    if m != nil {
        m.pinExpirationInDays = value
    }
}
// SetPinLowercaseCharactersUsage sets the pinLowercaseCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinLowercaseCharactersUsage(value *WindowsHelloForBusinessPinUsage)() {
    if m != nil {
        m.pinLowercaseCharactersUsage = value
    }
}
// SetPinMaximumLength sets the pinMaximumLength property value. Controls the maximum number of characters allowed for the Windows Hello for Business PIN. This value must be between 4 and 127, inclusive. This value must be greater than or equal to the value set for the minimum PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMaximumLength(value *int32)() {
    if m != nil {
        m.pinMaximumLength = value
    }
}
// SetPinMinimumLength sets the pinMinimumLength property value. Controls the minimum number of characters required for the Windows Hello for Business PIN.  This value must be between 4 and 127, inclusive, and less than or equal to the value set for the maximum PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMinimumLength(value *int32)() {
    if m != nil {
        m.pinMinimumLength = value
    }
}
// SetPinPreviousBlockCount sets the pinPreviousBlockCount property value. Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinPreviousBlockCount(value *int32)() {
    if m != nil {
        m.pinPreviousBlockCount = value
    }
}
// SetPinSpecialCharactersUsage sets the pinSpecialCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinSpecialCharactersUsage(value *WindowsHelloForBusinessPinUsage)() {
    if m != nil {
        m.pinSpecialCharactersUsage = value
    }
}
// SetPinUppercaseCharactersUsage sets the pinUppercaseCharactersUsage property value. Windows Hello for Business pin usage options
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinUppercaseCharactersUsage(value *WindowsHelloForBusinessPinUsage)() {
    if m != nil {
        m.pinUppercaseCharactersUsage = value
    }
}
// SetRemotePassportEnabled sets the remotePassportEnabled property value. Controls the use of Remote Windows Hello for Business. Remote Windows Hello for Business provides the ability for a portable, registered device to be usable as a companion for desktop authentication. The desktop must be Azure AD joined and the companion device must have a Windows Hello for Business PIN.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetRemotePassportEnabled(value *bool)() {
    if m != nil {
        m.remotePassportEnabled = value
    }
}
// SetSecurityDeviceRequired sets the securityDeviceRequired property value. Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetSecurityDeviceRequired(value *bool)() {
    if m != nil {
        m.securityDeviceRequired = value
    }
}
// SetState sets the state property value. The state property
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetState(value *Enablement)() {
    if m != nil {
        m.state = value
    }
}
// SetUnlockWithBiometricsEnabled sets the unlockWithBiometricsEnabled property value. Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
func (m *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetUnlockWithBiometricsEnabled(value *bool)() {
    if m != nil {
        m.unlockWithBiometricsEnabled = value
    }
}
