package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAgent 
type UserAgent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifies the version of application software used by this endpoint.
    applicationVersion *string
    // User-agent header value reported by this endpoint.
    headerValue *string
    // The type property
    type_escaped *string
}
// NewUserAgent instantiates a new userAgent and sets the default values.
func NewUserAgent()(*UserAgent) {
    m := &UserAgent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    typeValue := "#microsoft.graph.callRecords.userAgent";
    m.SetType(&typeValue);
    return m
}
// CreateUserAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.callRecords.clientUserAgent":
                        return NewClientUserAgent(), nil
                    case "#microsoft.graph.callRecords.serviceUserAgent":
                        return NewServiceUserAgent(), nil
                }
            }
        }
    }
    return NewUserAgent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAgent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationVersion gets the applicationVersion property value. Identifies the version of application software used by this endpoint.
func (m *UserAgent) GetApplicationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationVersion(val)
        }
        return nil
    }
    res["headerValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderValue(val)
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
    return res
}
// GetHeaderValue gets the headerValue property value. User-agent header value reported by this endpoint.
func (m *UserAgent) GetHeaderValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerValue
    }
}
// GetType gets the type property value. The type property
func (m *UserAgent) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *UserAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationVersion", m.GetApplicationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("headerValue", m.GetHeaderValue())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAgent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationVersion sets the applicationVersion property value. Identifies the version of application software used by this endpoint.
func (m *UserAgent) SetApplicationVersion(value *string)() {
    if m != nil {
        m.applicationVersion = value
    }
}
// SetHeaderValue sets the headerValue property value. User-agent header value reported by this endpoint.
func (m *UserAgent) SetHeaderValue(value *string)() {
    if m != nil {
        m.headerValue = value
    }
}
// SetType sets the type property value. The type property
func (m *UserAgent) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
