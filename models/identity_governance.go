package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernance 
type IdentityGovernance struct {
    // The accessReviews property
    accessReviews AccessReviewSetable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appConsent property
    appConsent AppConsentApprovalRouteable
    // The entitlementManagement property
    entitlementManagement EntitlementManagementable
    // The termsOfUse property
    termsOfUse TermsOfUseContainerable
}
// NewIdentityGovernance instantiates a new IdentityGovernance and sets the default values.
func NewIdentityGovernance()(*IdentityGovernance) {
    m := &IdentityGovernance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernance(), nil
}
// GetAccessReviews gets the accessReviews property value. The accessReviews property
func (m *IdentityGovernance) GetAccessReviews()(AccessReviewSetable) {
    if m == nil {
        return nil
    } else {
        return m.accessReviews
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppConsent gets the appConsent property value. The appConsent property
func (m *IdentityGovernance) GetAppConsent()(AppConsentApprovalRouteable) {
    if m == nil {
        return nil
    } else {
        return m.appConsent
    }
}
// GetEntitlementManagement gets the entitlementManagement property value. The entitlementManagement property
func (m *IdentityGovernance) GetEntitlementManagement()(EntitlementManagementable) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessReviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviews(val.(AccessReviewSetable))
        }
        return nil
    }
    res["appConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppConsentApprovalRouteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConsent(val.(AppConsentApprovalRouteable))
        }
        return nil
    }
    res["entitlementManagement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlementManagement(val.(EntitlementManagementable))
        }
        return nil
    }
    res["termsOfUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermsOfUseContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfUse(val.(TermsOfUseContainerable))
        }
        return nil
    }
    return res
}
// GetTermsOfUse gets the termsOfUse property value. The termsOfUse property
func (m *IdentityGovernance) GetTermsOfUse()(TermsOfUseContainerable) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
// Serialize serializes information the current object
func (m *IdentityGovernance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviews", m.GetAccessReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("appConsent", m.GetAppConsent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("termsOfUse", m.GetTermsOfUse())
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
// SetAccessReviews sets the accessReviews property value. The accessReviews property
func (m *IdentityGovernance) SetAccessReviews(value AccessReviewSetable)() {
    if m != nil {
        m.accessReviews = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppConsent sets the appConsent property value. The appConsent property
func (m *IdentityGovernance) SetAppConsent(value AppConsentApprovalRouteable)() {
    if m != nil {
        m.appConsent = value
    }
}
// SetEntitlementManagement sets the entitlementManagement property value. The entitlementManagement property
func (m *IdentityGovernance) SetEntitlementManagement(value EntitlementManagementable)() {
    if m != nil {
        m.entitlementManagement = value
    }
}
// SetTermsOfUse sets the termsOfUse property value. The termsOfUse property
func (m *IdentityGovernance) SetTermsOfUse(value TermsOfUseContainerable)() {
    if m != nil {
        m.termsOfUse = value
    }
}
