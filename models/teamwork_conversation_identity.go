package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkConversationIdentity 
type TeamworkConversationIdentity struct {
    Identity
    // Type of conversation. Possible values are: team, channel, and chat.
    conversationIdentityType *TeamworkConversationIdentityType
}
// NewTeamworkConversationIdentity instantiates a new TeamworkConversationIdentity and sets the default values.
func NewTeamworkConversationIdentity()(*TeamworkConversationIdentity) {
    m := &TeamworkConversationIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateTeamworkConversationIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkConversationIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkConversationIdentity(), nil
}
// GetConversationIdentityType gets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, and chat.
func (m *TeamworkConversationIdentity) GetConversationIdentityType()(*TeamworkConversationIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.conversationIdentityType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConversationIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["conversationIdentityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkConversationIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationIdentityType(val.(*TeamworkConversationIdentityType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamworkConversationIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConversationIdentityType() != nil {
        cast := (*m.GetConversationIdentityType()).String()
        err = writer.WriteStringValue("conversationIdentityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConversationIdentityType sets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, and chat.
func (m *TeamworkConversationIdentity) SetConversationIdentityType(value *TeamworkConversationIdentityType)() {
    if m != nil {
        m.conversationIdentityType = value
    }
}
