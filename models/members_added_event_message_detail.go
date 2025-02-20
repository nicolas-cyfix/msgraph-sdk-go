package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembersAddedEventMessageDetail 
type MembersAddedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // List of members added.
    members []TeamworkUserIdentityable
    // The timestamp denoting how far back a conversation's history is shared with the conversation members.
    visibleHistoryStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewMembersAddedEventMessageDetail instantiates a new MembersAddedEventMessageDetail and sets the default values.
func NewMembersAddedEventMessageDetail()(*MembersAddedEventMessageDetail) {
    m := &MembersAddedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    return m
}
// CreateMembersAddedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMembersAddedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMembersAddedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MembersAddedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkUserIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkUserIdentityable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["visibleHistoryStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibleHistoryStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *MembersAddedEventMessageDetail) GetInitiator()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.initiator
    }
}
// GetMembers gets the members property value. List of members added.
func (m *MembersAddedEventMessageDetail) GetMembers()([]TeamworkUserIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetVisibleHistoryStartDateTime gets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation members.
func (m *MembersAddedEventMessageDetail) GetVisibleHistoryStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.visibleHistoryStartDateTime
    }
}
// Serialize serializes information the current object
func (m *MembersAddedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("visibleHistoryStartDateTime", m.GetVisibleHistoryStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *MembersAddedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    if m != nil {
        m.initiator = value
    }
}
// SetMembers sets the members property value. List of members added.
func (m *MembersAddedEventMessageDetail) SetMembers(value []TeamworkUserIdentityable)() {
    if m != nil {
        m.members = value
    }
}
// SetVisibleHistoryStartDateTime sets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation members.
func (m *MembersAddedEventMessageDetail) SetVisibleHistoryStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.visibleHistoryStartDateTime = value
    }
}
