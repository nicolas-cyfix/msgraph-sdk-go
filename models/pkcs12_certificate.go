package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Pkcs12Certificate 
type Pkcs12Certificate struct {
    ApiAuthenticationConfigurationBase
    // This is the password for the pfx file. Required. If no password is used, must still provide a value of ''.
    password *string
    // This is the field for sending pfx content. The value should be a base-64 encoded version of the actual certificate content. Required.
    pkcs12Value *string
}
// NewPkcs12Certificate instantiates a new Pkcs12Certificate and sets the default values.
func NewPkcs12Certificate()(*Pkcs12Certificate) {
    m := &Pkcs12Certificate{
        ApiAuthenticationConfigurationBase: *NewApiAuthenticationConfigurationBase(),
    }
    return m
}
// CreatePkcs12CertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePkcs12CertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPkcs12Certificate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Pkcs12Certificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApiAuthenticationConfigurationBase.GetFieldDeserializers()
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["pkcs12Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPkcs12Value(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. This is the password for the pfx file. Required. If no password is used, must still provide a value of ''.
func (m *Pkcs12Certificate) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetPkcs12Value gets the pkcs12Value property value. This is the field for sending pfx content. The value should be a base-64 encoded version of the actual certificate content. Required.
func (m *Pkcs12Certificate) GetPkcs12Value()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pkcs12Value
    }
}
// Serialize serializes information the current object
func (m *Pkcs12Certificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApiAuthenticationConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("pkcs12Value", m.GetPkcs12Value())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPassword sets the password property value. This is the password for the pfx file. Required. If no password is used, must still provide a value of ''.
func (m *Pkcs12Certificate) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
// SetPkcs12Value sets the pkcs12Value property value. This is the field for sending pfx content. The value should be a base-64 encoded version of the actual certificate content. Required.
func (m *Pkcs12Certificate) SetPkcs12Value(value *string)() {
    if m != nil {
        m.pkcs12Value = value
    }
}
