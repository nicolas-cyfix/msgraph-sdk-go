package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageSecurityStateable 
type MessageSecurityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectingIP()(*string)
    GetDeliveryAction()(*string)
    GetDeliveryLocation()(*string)
    GetDirectionality()(*string)
    GetInternetMessageId()(*string)
    GetMessageFingerprint()(*string)
    GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMessageSubject()(*string)
    GetNetworkMessageId()(*string)
    SetConnectingIP(value *string)()
    SetDeliveryAction(value *string)()
    SetDeliveryLocation(value *string)()
    SetDirectionality(value *string)()
    SetInternetMessageId(value *string)()
    SetMessageFingerprint(value *string)()
    SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMessageSubject(value *string)()
    SetNetworkMessageId(value *string)()
}
