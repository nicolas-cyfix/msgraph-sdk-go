package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplePushNotificationCertificate 
type ApplePushNotificationCertificate struct {
    Entity
    // Apple Id of the account used to create the MDM push certificate.
    appleIdentifier *string
    // Not yet documented
    certificate *string
    // Certificate serial number. This property is read-only.
    certificateSerialNumber *string
    // The expiration date and time for Apple push notification certificate.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last modified date and time for Apple push notification certificate.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Topic Id.
    topicIdentifier *string
}
// NewApplePushNotificationCertificate instantiates a new ApplePushNotificationCertificate and sets the default values.
func NewApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    m := &ApplePushNotificationCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApplePushNotificationCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplePushNotificationCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplePushNotificationCertificate(), nil
}
// GetAppleIdentifier gets the appleIdentifier property value. Apple Id of the account used to create the MDM push certificate.
func (m *ApplePushNotificationCertificate) GetAppleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleIdentifier
    }
}
// GetCertificate gets the certificate property value. Not yet documented
func (m *ApplePushNotificationCertificate) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// GetCertificateSerialNumber gets the certificateSerialNumber property value. Certificate serial number. This property is read-only.
func (m *ApplePushNotificationCertificate) GetCertificateSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificateSerialNumber
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplePushNotificationCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdentifier(val)
        }
        return nil
    }
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["certificateSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSerialNumber(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["topicIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicIdentifier(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetTopicIdentifier gets the topicIdentifier property value. Topic Id.
func (m *ApplePushNotificationCertificate) GetTopicIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topicIdentifier
    }
}
// Serialize serializes information the current object
func (m *ApplePushNotificationCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleIdentifier", m.GetAppleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSerialNumber", m.GetCertificateSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topicIdentifier", m.GetTopicIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleIdentifier sets the appleIdentifier property value. Apple Id of the account used to create the MDM push certificate.
func (m *ApplePushNotificationCertificate) SetAppleIdentifier(value *string)() {
    if m != nil {
        m.appleIdentifier = value
    }
}
// SetCertificate sets the certificate property value. Not yet documented
func (m *ApplePushNotificationCertificate) SetCertificate(value *string)() {
    if m != nil {
        m.certificate = value
    }
}
// SetCertificateSerialNumber sets the certificateSerialNumber property value. Certificate serial number. This property is read-only.
func (m *ApplePushNotificationCertificate) SetCertificateSerialNumber(value *string)() {
    if m != nil {
        m.certificateSerialNumber = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date and time for Apple push notification certificate.
func (m *ApplePushNotificationCertificate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetTopicIdentifier sets the topicIdentifier property value. Topic Id.
func (m *ApplePushNotificationCertificate) SetTopicIdentifier(value *string)() {
    if m != nil {
        m.topicIdentifier = value
    }
}
