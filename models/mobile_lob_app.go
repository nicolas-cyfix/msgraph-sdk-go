package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileLobApp 
type MobileLobApp struct {
    MobileApp
    // The internal committed content version.
    committedContentVersion *string
    // The list of content versions for this app.
    contentVersions []MobileAppContentable
    // The name of the main Lob application file.
    fileName *string
    // The total size, including all uploaded files.
    size *int64
}
// NewMobileLobApp instantiates a new MobileLobApp and sets the default values.
func NewMobileLobApp()(*MobileLobApp) {
    m := &MobileLobApp{
        MobileApp: *NewMobileApp(),
    }
    typeValue := "#microsoft.graph.mobileLobApp";
    m.SetType(&typeValue);
    return m
}
// CreateMobileLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidLobApp":
                        return NewAndroidLobApp(), nil
                    case "#microsoft.graph.iosLobApp":
                        return NewIosLobApp(), nil
                    case "#microsoft.graph.win32LobApp":
                        return NewWin32LobApp(), nil
                    case "#microsoft.graph.windowsMobileMSI":
                        return NewWindowsMobileMSI(), nil
                    case "#microsoft.graph.windowsUniversalAppX":
                        return NewWindowsUniversalAppX(), nil
                }
            }
        }
    }
    return NewMobileLobApp(), nil
}
// GetCommittedContentVersion gets the committedContentVersion property value. The internal committed content version.
func (m *MobileLobApp) GetCommittedContentVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.committedContentVersion
    }
}
// GetContentVersions gets the contentVersions property value. The list of content versions for this app.
func (m *MobileLobApp) GetContentVersions()([]MobileAppContentable) {
    if m == nil {
        return nil
    } else {
        return m.contentVersions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["committedContentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommittedContentVersion(val)
        }
        return nil
    }
    res["contentVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppContentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppContentable)
            }
            m.SetContentVersions(res)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The name of the main Lob application file.
func (m *MobileLobApp) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetSize gets the size property value. The total size, including all uploaded files.
func (m *MobileLobApp) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Serialize serializes information the current object
func (m *MobileLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("committedContentVersion", m.GetCommittedContentVersion())
        if err != nil {
            return err
        }
    }
    if m.GetContentVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentVersions()))
        for i, v := range m.GetContentVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contentVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCommittedContentVersion sets the committedContentVersion property value. The internal committed content version.
func (m *MobileLobApp) SetCommittedContentVersion(value *string)() {
    if m != nil {
        m.committedContentVersion = value
    }
}
// SetContentVersions sets the contentVersions property value. The list of content versions for this app.
func (m *MobileLobApp) SetContentVersions(value []MobileAppContentable)() {
    if m != nil {
        m.contentVersions = value
    }
}
// SetFileName sets the fileName property value. The name of the main Lob application file.
func (m *MobileLobApp) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetSize sets the size property value. The total size, including all uploaded files.
func (m *MobileLobApp) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
