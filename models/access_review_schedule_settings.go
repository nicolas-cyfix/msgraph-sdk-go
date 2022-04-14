package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewScheduleSettings 
type AccessReviewScheduleSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
    applyActions []AccessReviewApplyActionable
    // Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
    autoApplyDecisionsEnabled *bool
    // Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
    defaultDecision *string
    // Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
    defaultDecisionEnabled *bool
    // Duration of each recurrence of review (accessReviewInstance) in number of days.
    instanceDurationInDays *int32
    // Indicates whether reviewers are required to provide justification with their decision. Default value is false.
    justificationRequiredOnApproval *bool
    // Indicates whether emails are enabled or disabled. Default value is false.
    mailNotificationsEnabled *bool
    // Indicates whether decision recommendations are enabled or disabled.
    recommendationsEnabled *bool
    // Detailed settings for recurrence using the standard Outlook recurrence object.  Note: Only dayOfMonth, interval, and type (weekly, absoluteMonthly) properties are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
    recurrence PatternedRecurrenceable
    // Indicates whether reminders are enabled or disabled. Default value is false.
    reminderNotificationsEnabled *bool
}
// NewAccessReviewScheduleSettings instantiates a new accessReviewScheduleSettings and sets the default values.
func NewAccessReviewScheduleSettings()(*AccessReviewScheduleSettings) {
    m := &AccessReviewScheduleSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewScheduleSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewScheduleSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewScheduleSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewScheduleSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplyActions gets the applyActions property value. Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
func (m *AccessReviewScheduleSettings) GetApplyActions()([]AccessReviewApplyActionable) {
    if m == nil {
        return nil
    } else {
        return m.applyActions
    }
}
// GetAutoApplyDecisionsEnabled gets the autoApplyDecisionsEnabled property value. Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
func (m *AccessReviewScheduleSettings) GetAutoApplyDecisionsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoApplyDecisionsEnabled
    }
}
// GetDefaultDecision gets the defaultDecision property value. Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
func (m *AccessReviewScheduleSettings) GetDefaultDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecision
    }
}
// GetDefaultDecisionEnabled gets the defaultDecisionEnabled property value. Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
func (m *AccessReviewScheduleSettings) GetDefaultDecisionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultDecisionEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewScheduleSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applyActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewApplyActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewApplyActionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewApplyActionable)
            }
            m.SetApplyActions(res)
        }
        return nil
    }
    res["autoApplyDecisionsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoApplyDecisionsEnabled(val)
        }
        return nil
    }
    res["defaultDecision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDecision(val)
        }
        return nil
    }
    res["defaultDecisionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDecisionEnabled(val)
        }
        return nil
    }
    res["instanceDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceDurationInDays(val)
        }
        return nil
    }
    res["justificationRequiredOnApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationRequiredOnApproval(val)
        }
        return nil
    }
    res["mailNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNotificationsEnabled(val)
        }
        return nil
    }
    res["recommendationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationsEnabled(val)
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["reminderNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderNotificationsEnabled(val)
        }
        return nil
    }
    return res
}
// GetInstanceDurationInDays gets the instanceDurationInDays property value. Duration of each recurrence of review (accessReviewInstance) in number of days.
func (m *AccessReviewScheduleSettings) GetInstanceDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.instanceDurationInDays
    }
}
// GetJustificationRequiredOnApproval gets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide justification with their decision. Default value is false.
func (m *AccessReviewScheduleSettings) GetJustificationRequiredOnApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.justificationRequiredOnApproval
    }
}
// GetMailNotificationsEnabled gets the mailNotificationsEnabled property value. Indicates whether emails are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) GetMailNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailNotificationsEnabled
    }
}
// GetRecommendationsEnabled gets the recommendationsEnabled property value. Indicates whether decision recommendations are enabled or disabled.
func (m *AccessReviewScheduleSettings) GetRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recommendationsEnabled
    }
}
// GetRecurrence gets the recurrence property value. Detailed settings for recurrence using the standard Outlook recurrence object.  Note: Only dayOfMonth, interval, and type (weekly, absoluteMonthly) properties are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
func (m *AccessReviewScheduleSettings) GetRecurrence()(PatternedRecurrenceable) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetReminderNotificationsEnabled gets the reminderNotificationsEnabled property value. Indicates whether reminders are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) GetReminderNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.reminderNotificationsEnabled
    }
}
// Serialize serializes information the current object
func (m *AccessReviewScheduleSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApplyActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApplyActions()))
        for i, v := range m.GetApplyActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("applyActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoApplyDecisionsEnabled", m.GetAutoApplyDecisionsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultDecision", m.GetDefaultDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("defaultDecisionEnabled", m.GetDefaultDecisionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("instanceDurationInDays", m.GetInstanceDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("justificationRequiredOnApproval", m.GetJustificationRequiredOnApproval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailNotificationsEnabled", m.GetMailNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recommendationsEnabled", m.GetRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("reminderNotificationsEnabled", m.GetReminderNotificationsEnabled())
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
func (m *AccessReviewScheduleSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplyActions sets the applyActions property value. Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction.
func (m *AccessReviewScheduleSettings) SetApplyActions(value []AccessReviewApplyActionable)() {
    if m != nil {
        m.applyActions = value
    }
}
// SetAutoApplyDecisionsEnabled sets the autoApplyDecisionsEnabled property value. Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false.
func (m *AccessReviewScheduleSettings) SetAutoApplyDecisionsEnabled(value *bool)() {
    if m != nil {
        m.autoApplyDecisionsEnabled = value
    }
}
// SetDefaultDecision sets the defaultDecision property value. Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation.
func (m *AccessReviewScheduleSettings) SetDefaultDecision(value *string)() {
    if m != nil {
        m.defaultDecision = value
    }
}
// SetDefaultDecisionEnabled sets the defaultDecisionEnabled property value. Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
func (m *AccessReviewScheduleSettings) SetDefaultDecisionEnabled(value *bool)() {
    if m != nil {
        m.defaultDecisionEnabled = value
    }
}
// SetInstanceDurationInDays sets the instanceDurationInDays property value. Duration of each recurrence of review (accessReviewInstance) in number of days.
func (m *AccessReviewScheduleSettings) SetInstanceDurationInDays(value *int32)() {
    if m != nil {
        m.instanceDurationInDays = value
    }
}
// SetJustificationRequiredOnApproval sets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide justification with their decision. Default value is false.
func (m *AccessReviewScheduleSettings) SetJustificationRequiredOnApproval(value *bool)() {
    if m != nil {
        m.justificationRequiredOnApproval = value
    }
}
// SetMailNotificationsEnabled sets the mailNotificationsEnabled property value. Indicates whether emails are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) SetMailNotificationsEnabled(value *bool)() {
    if m != nil {
        m.mailNotificationsEnabled = value
    }
}
// SetRecommendationsEnabled sets the recommendationsEnabled property value. Indicates whether decision recommendations are enabled or disabled.
func (m *AccessReviewScheduleSettings) SetRecommendationsEnabled(value *bool)() {
    if m != nil {
        m.recommendationsEnabled = value
    }
}
// SetRecurrence sets the recurrence property value. Detailed settings for recurrence using the standard Outlook recurrence object.  Note: Only dayOfMonth, interval, and type (weekly, absoluteMonthly) properties are supported. Use the property startDate on recurrenceRange to determine the day the review starts.
func (m *AccessReviewScheduleSettings) SetRecurrence(value PatternedRecurrenceable)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetReminderNotificationsEnabled sets the reminderNotificationsEnabled property value. Indicates whether reminders are enabled or disabled. Default value is false.
func (m *AccessReviewScheduleSettings) SetReminderNotificationsEnabled(value *bool)() {
    if m != nil {
        m.reminderNotificationsEnabled = value
    }
}
