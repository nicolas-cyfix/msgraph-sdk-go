package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenotePagePreview 
type OnenotePagePreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The links property
    links OnenotePagePreviewLinksable
    // The previewText property
    previewText *string
}
// NewOnenotePagePreview instantiates a new onenotePagePreview and sets the default values.
func NewOnenotePagePreview()(*OnenotePagePreview) {
    m := &OnenotePagePreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePagePreviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePagePreviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenotePagePreview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePagePreview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenotePagePreviewLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(OnenotePagePreviewLinksable))
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. The links property
func (m *OnenotePagePreview) GetLinks()(OnenotePagePreviewLinksable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *OnenotePagePreview) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// Serialize serializes information the current object
func (m *OnenotePagePreview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
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
func (m *OnenotePagePreview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLinks sets the links property value. The links property
func (m *OnenotePagePreview) SetLinks(value OnenotePagePreviewLinksable)() {
    if m != nil {
        m.links = value
    }
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *OnenotePagePreview) SetPreviewText(value *string)() {
    if m != nil {
        m.previewText = value
    }
}
