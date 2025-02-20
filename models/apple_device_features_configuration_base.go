package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleDeviceFeaturesConfigurationBase 
type AppleDeviceFeaturesConfigurationBase struct {
    DeviceConfiguration
}
// NewAppleDeviceFeaturesConfigurationBase instantiates a new AppleDeviceFeaturesConfigurationBase and sets the default values.
func NewAppleDeviceFeaturesConfigurationBase()(*AppleDeviceFeaturesConfigurationBase) {
    m := &AppleDeviceFeaturesConfigurationBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    typeValue := "#microsoft.graph.appleDeviceFeaturesConfigurationBase";
    m.SetType(&typeValue);
    return m
}
// CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosDeviceFeaturesConfiguration":
                        return NewIosDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.macOSDeviceFeaturesConfiguration":
                        return NewMacOSDeviceFeaturesConfiguration(), nil
                }
            }
        }
    }
    return NewAppleDeviceFeaturesConfigurationBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleDeviceFeaturesConfigurationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AppleDeviceFeaturesConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
