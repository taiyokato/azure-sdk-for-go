package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AbsoluteMarker enumerates the values for absolute marker.
type AbsoluteMarker string

const (
	// AbsoluteMarkerAllBackup ...
	AbsoluteMarkerAllBackup AbsoluteMarker = "AllBackup"
	// AbsoluteMarkerFirstOfDay ...
	AbsoluteMarkerFirstOfDay AbsoluteMarker = "FirstOfDay"
	// AbsoluteMarkerFirstOfMonth ...
	AbsoluteMarkerFirstOfMonth AbsoluteMarker = "FirstOfMonth"
	// AbsoluteMarkerFirstOfWeek ...
	AbsoluteMarkerFirstOfWeek AbsoluteMarker = "FirstOfWeek"
	// AbsoluteMarkerFirstOfYear ...
	AbsoluteMarkerFirstOfYear AbsoluteMarker = "FirstOfYear"
)

// PossibleAbsoluteMarkerValues returns an array of possible values for the AbsoluteMarker const type.
func PossibleAbsoluteMarkerValues() []AbsoluteMarker {
	return []AbsoluteMarker{AbsoluteMarkerAllBackup, AbsoluteMarkerFirstOfDay, AbsoluteMarkerFirstOfMonth, AbsoluteMarkerFirstOfWeek, AbsoluteMarkerFirstOfYear}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// CurrentProtectionState enumerates the values for current protection state.
type CurrentProtectionState string

const (
	// CurrentProtectionStateBackupSchedulesSuspended ...
	CurrentProtectionStateBackupSchedulesSuspended CurrentProtectionState = "BackupSchedulesSuspended"
	// CurrentProtectionStateConfiguringProtection ...
	CurrentProtectionStateConfiguringProtection CurrentProtectionState = "ConfiguringProtection"
	// CurrentProtectionStateConfiguringProtectionFailed ...
	CurrentProtectionStateConfiguringProtectionFailed CurrentProtectionState = "ConfiguringProtectionFailed"
	// CurrentProtectionStateInvalid ...
	CurrentProtectionStateInvalid CurrentProtectionState = "Invalid"
	// CurrentProtectionStateNotProtected ...
	CurrentProtectionStateNotProtected CurrentProtectionState = "NotProtected"
	// CurrentProtectionStateProtectionConfigured ...
	CurrentProtectionStateProtectionConfigured CurrentProtectionState = "ProtectionConfigured"
	// CurrentProtectionStateProtectionError ...
	CurrentProtectionStateProtectionError CurrentProtectionState = "ProtectionError"
	// CurrentProtectionStateProtectionStopped ...
	CurrentProtectionStateProtectionStopped CurrentProtectionState = "ProtectionStopped"
	// CurrentProtectionStateRetentionSchedulesSuspended ...
	CurrentProtectionStateRetentionSchedulesSuspended CurrentProtectionState = "RetentionSchedulesSuspended"
	// CurrentProtectionStateSoftDeleted ...
	CurrentProtectionStateSoftDeleted CurrentProtectionState = "SoftDeleted"
	// CurrentProtectionStateSoftDeleting ...
	CurrentProtectionStateSoftDeleting CurrentProtectionState = "SoftDeleting"
	// CurrentProtectionStateUpdatingProtection ...
	CurrentProtectionStateUpdatingProtection CurrentProtectionState = "UpdatingProtection"
)

// PossibleCurrentProtectionStateValues returns an array of possible values for the CurrentProtectionState const type.
func PossibleCurrentProtectionStateValues() []CurrentProtectionState {
	return []CurrentProtectionState{CurrentProtectionStateBackupSchedulesSuspended, CurrentProtectionStateConfiguringProtection, CurrentProtectionStateConfiguringProtectionFailed, CurrentProtectionStateInvalid, CurrentProtectionStateNotProtected, CurrentProtectionStateProtectionConfigured, CurrentProtectionStateProtectionError, CurrentProtectionStateProtectionStopped, CurrentProtectionStateRetentionSchedulesSuspended, CurrentProtectionStateSoftDeleted, CurrentProtectionStateSoftDeleting, CurrentProtectionStateUpdatingProtection}
}

// DataStoreTypes enumerates the values for data store types.
type DataStoreTypes string

const (
	// DataStoreTypesArchiveStore ...
	DataStoreTypesArchiveStore DataStoreTypes = "ArchiveStore"
	// DataStoreTypesOperationalStore ...
	DataStoreTypesOperationalStore DataStoreTypes = "OperationalStore"
	// DataStoreTypesVaultStore ...
	DataStoreTypesVaultStore DataStoreTypes = "VaultStore"
)

// PossibleDataStoreTypesValues returns an array of possible values for the DataStoreTypes const type.
func PossibleDataStoreTypesValues() []DataStoreTypes {
	return []DataStoreTypes{DataStoreTypesArchiveStore, DataStoreTypesOperationalStore, DataStoreTypesVaultStore}
}

// DayOfWeek enumerates the values for day of week.
type DayOfWeek string

const (
	// DayOfWeekFriday ...
	DayOfWeekFriday DayOfWeek = "Friday"
	// DayOfWeekMonday ...
	DayOfWeekMonday DayOfWeek = "Monday"
	// DayOfWeekSaturday ...
	DayOfWeekSaturday DayOfWeek = "Saturday"
	// DayOfWeekSunday ...
	DayOfWeekSunday DayOfWeek = "Sunday"
	// DayOfWeekThursday ...
	DayOfWeekThursday DayOfWeek = "Thursday"
	// DayOfWeekTuesday ...
	DayOfWeekTuesday DayOfWeek = "Tuesday"
	// DayOfWeekWednesday ...
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns an array of possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{DayOfWeekFriday, DayOfWeekMonday, DayOfWeekSaturday, DayOfWeekSunday, DayOfWeekThursday, DayOfWeekTuesday, DayOfWeekWednesday}
}

// FeatureSupportStatus enumerates the values for feature support status.
type FeatureSupportStatus string

const (
	// FeatureSupportStatusAlphaPreview ...
	FeatureSupportStatusAlphaPreview FeatureSupportStatus = "AlphaPreview"
	// FeatureSupportStatusGenerallyAvailable ...
	FeatureSupportStatusGenerallyAvailable FeatureSupportStatus = "GenerallyAvailable"
	// FeatureSupportStatusInvalid ...
	FeatureSupportStatusInvalid FeatureSupportStatus = "Invalid"
	// FeatureSupportStatusNotSupported ...
	FeatureSupportStatusNotSupported FeatureSupportStatus = "NotSupported"
	// FeatureSupportStatusPrivatePreview ...
	FeatureSupportStatusPrivatePreview FeatureSupportStatus = "PrivatePreview"
	// FeatureSupportStatusPublicPreview ...
	FeatureSupportStatusPublicPreview FeatureSupportStatus = "PublicPreview"
)

// PossibleFeatureSupportStatusValues returns an array of possible values for the FeatureSupportStatus const type.
func PossibleFeatureSupportStatusValues() []FeatureSupportStatus {
	return []FeatureSupportStatus{FeatureSupportStatusAlphaPreview, FeatureSupportStatusGenerallyAvailable, FeatureSupportStatusInvalid, FeatureSupportStatusNotSupported, FeatureSupportStatusPrivatePreview, FeatureSupportStatusPublicPreview}
}

// FeatureType enumerates the values for feature type.
type FeatureType string

const (
	// FeatureTypeDataSourceType ...
	FeatureTypeDataSourceType FeatureType = "DataSourceType"
	// FeatureTypeInvalid ...
	FeatureTypeInvalid FeatureType = "Invalid"
)

// PossibleFeatureTypeValues returns an array of possible values for the FeatureType const type.
func PossibleFeatureTypeValues() []FeatureType {
	return []FeatureType{FeatureTypeDataSourceType, FeatureTypeInvalid}
}

// Month enumerates the values for month.
type Month string

const (
	// MonthApril ...
	MonthApril Month = "April"
	// MonthAugust ...
	MonthAugust Month = "August"
	// MonthDecember ...
	MonthDecember Month = "December"
	// MonthFebruary ...
	MonthFebruary Month = "February"
	// MonthJanuary ...
	MonthJanuary Month = "January"
	// MonthJuly ...
	MonthJuly Month = "July"
	// MonthJune ...
	MonthJune Month = "June"
	// MonthMarch ...
	MonthMarch Month = "March"
	// MonthMay ...
	MonthMay Month = "May"
	// MonthNovember ...
	MonthNovember Month = "November"
	// MonthOctober ...
	MonthOctober Month = "October"
	// MonthSeptember ...
	MonthSeptember Month = "September"
)

// PossibleMonthValues returns an array of possible values for the Month const type.
func PossibleMonthValues() []Month {
	return []Month{MonthApril, MonthAugust, MonthDecember, MonthFebruary, MonthJanuary, MonthJuly, MonthJune, MonthMarch, MonthMay, MonthNovember, MonthOctober, MonthSeptember}
}

// ObjectType enumerates the values for object type.
type ObjectType string

const (
	// ObjectTypeAuthCredentials ...
	ObjectTypeAuthCredentials ObjectType = "AuthCredentials"
	// ObjectTypeSecretStoreBasedAuthCredentials ...
	ObjectTypeSecretStoreBasedAuthCredentials ObjectType = "SecretStoreBasedAuthCredentials"
)

// PossibleObjectTypeValues returns an array of possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{ObjectTypeAuthCredentials, ObjectTypeSecretStoreBasedAuthCredentials}
}

// ObjectTypeBasicAzureBackupRecoveryPoint enumerates the values for object type basic azure backup recovery
// point.
type ObjectTypeBasicAzureBackupRecoveryPoint string

const (
	// ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupDiscreteRecoveryPoint ...
	ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupDiscreteRecoveryPoint ObjectTypeBasicAzureBackupRecoveryPoint = "AzureBackupDiscreteRecoveryPoint"
	// ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupRecoveryPoint ...
	ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupRecoveryPoint ObjectTypeBasicAzureBackupRecoveryPoint = "AzureBackupRecoveryPoint"
)

// PossibleObjectTypeBasicAzureBackupRecoveryPointValues returns an array of possible values for the ObjectTypeBasicAzureBackupRecoveryPoint const type.
func PossibleObjectTypeBasicAzureBackupRecoveryPointValues() []ObjectTypeBasicAzureBackupRecoveryPoint {
	return []ObjectTypeBasicAzureBackupRecoveryPoint{ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupDiscreteRecoveryPoint, ObjectTypeBasicAzureBackupRecoveryPointObjectTypeAzureBackupRecoveryPoint}
}

// ObjectTypeBasicAzureBackupRestoreRequest enumerates the values for object type basic azure backup restore
// request.
type ObjectTypeBasicAzureBackupRestoreRequest string

const (
	// ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryPointBasedRestoreRequest ...
	ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryPointBasedRestoreRequest ObjectTypeBasicAzureBackupRestoreRequest = "AzureBackupRecoveryPointBasedRestoreRequest"
	// ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryTimeBasedRestoreRequest ...
	ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryTimeBasedRestoreRequest ObjectTypeBasicAzureBackupRestoreRequest = "AzureBackupRecoveryTimeBasedRestoreRequest"
	// ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreRequest ...
	ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreRequest ObjectTypeBasicAzureBackupRestoreRequest = "AzureBackupRestoreRequest"
	// ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreWithRehydrationRequest ...
	ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreWithRehydrationRequest ObjectTypeBasicAzureBackupRestoreRequest = "AzureBackupRestoreWithRehydrationRequest"
)

// PossibleObjectTypeBasicAzureBackupRestoreRequestValues returns an array of possible values for the ObjectTypeBasicAzureBackupRestoreRequest const type.
func PossibleObjectTypeBasicAzureBackupRestoreRequestValues() []ObjectTypeBasicAzureBackupRestoreRequest {
	return []ObjectTypeBasicAzureBackupRestoreRequest{ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryPointBasedRestoreRequest, ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRecoveryTimeBasedRestoreRequest, ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreRequest, ObjectTypeBasicAzureBackupRestoreRequestObjectTypeAzureBackupRestoreWithRehydrationRequest}
}

// ObjectTypeBasicBackupCriteria enumerates the values for object type basic backup criteria.
type ObjectTypeBasicBackupCriteria string

const (
	// ObjectTypeBasicBackupCriteriaObjectTypeBackupCriteria ...
	ObjectTypeBasicBackupCriteriaObjectTypeBackupCriteria ObjectTypeBasicBackupCriteria = "BackupCriteria"
	// ObjectTypeBasicBackupCriteriaObjectTypeScheduleBasedBackupCriteria ...
	ObjectTypeBasicBackupCriteriaObjectTypeScheduleBasedBackupCriteria ObjectTypeBasicBackupCriteria = "ScheduleBasedBackupCriteria"
)

// PossibleObjectTypeBasicBackupCriteriaValues returns an array of possible values for the ObjectTypeBasicBackupCriteria const type.
func PossibleObjectTypeBasicBackupCriteriaValues() []ObjectTypeBasicBackupCriteria {
	return []ObjectTypeBasicBackupCriteria{ObjectTypeBasicBackupCriteriaObjectTypeBackupCriteria, ObjectTypeBasicBackupCriteriaObjectTypeScheduleBasedBackupCriteria}
}

// ObjectTypeBasicBackupParameters enumerates the values for object type basic backup parameters.
type ObjectTypeBasicBackupParameters string

const (
	// ObjectTypeBasicBackupParametersObjectTypeAzureBackupParams ...
	ObjectTypeBasicBackupParametersObjectTypeAzureBackupParams ObjectTypeBasicBackupParameters = "AzureBackupParams"
	// ObjectTypeBasicBackupParametersObjectTypeBackupParameters ...
	ObjectTypeBasicBackupParametersObjectTypeBackupParameters ObjectTypeBasicBackupParameters = "BackupParameters"
)

// PossibleObjectTypeBasicBackupParametersValues returns an array of possible values for the ObjectTypeBasicBackupParameters const type.
func PossibleObjectTypeBasicBackupParametersValues() []ObjectTypeBasicBackupParameters {
	return []ObjectTypeBasicBackupParameters{ObjectTypeBasicBackupParametersObjectTypeAzureBackupParams, ObjectTypeBasicBackupParametersObjectTypeBackupParameters}
}

// ObjectTypeBasicBaseBackupPolicy enumerates the values for object type basic base backup policy.
type ObjectTypeBasicBaseBackupPolicy string

const (
	// ObjectTypeBasicBaseBackupPolicyObjectTypeBackupPolicy ...
	ObjectTypeBasicBaseBackupPolicyObjectTypeBackupPolicy ObjectTypeBasicBaseBackupPolicy = "BackupPolicy"
	// ObjectTypeBasicBaseBackupPolicyObjectTypeBaseBackupPolicy ...
	ObjectTypeBasicBaseBackupPolicyObjectTypeBaseBackupPolicy ObjectTypeBasicBaseBackupPolicy = "BaseBackupPolicy"
)

// PossibleObjectTypeBasicBaseBackupPolicyValues returns an array of possible values for the ObjectTypeBasicBaseBackupPolicy const type.
func PossibleObjectTypeBasicBaseBackupPolicyValues() []ObjectTypeBasicBaseBackupPolicy {
	return []ObjectTypeBasicBaseBackupPolicy{ObjectTypeBasicBaseBackupPolicyObjectTypeBackupPolicy, ObjectTypeBasicBaseBackupPolicyObjectTypeBaseBackupPolicy}
}

// ObjectTypeBasicBasePolicyRule enumerates the values for object type basic base policy rule.
type ObjectTypeBasicBasePolicyRule string

const (
	// ObjectTypeBasicBasePolicyRuleObjectTypeAzureBackupRule ...
	ObjectTypeBasicBasePolicyRuleObjectTypeAzureBackupRule ObjectTypeBasicBasePolicyRule = "AzureBackupRule"
	// ObjectTypeBasicBasePolicyRuleObjectTypeAzureRetentionRule ...
	ObjectTypeBasicBasePolicyRuleObjectTypeAzureRetentionRule ObjectTypeBasicBasePolicyRule = "AzureRetentionRule"
	// ObjectTypeBasicBasePolicyRuleObjectTypeBasePolicyRule ...
	ObjectTypeBasicBasePolicyRuleObjectTypeBasePolicyRule ObjectTypeBasicBasePolicyRule = "BasePolicyRule"
)

// PossibleObjectTypeBasicBasePolicyRuleValues returns an array of possible values for the ObjectTypeBasicBasePolicyRule const type.
func PossibleObjectTypeBasicBasePolicyRuleValues() []ObjectTypeBasicBasePolicyRule {
	return []ObjectTypeBasicBasePolicyRule{ObjectTypeBasicBasePolicyRuleObjectTypeAzureBackupRule, ObjectTypeBasicBasePolicyRuleObjectTypeAzureRetentionRule, ObjectTypeBasicBasePolicyRuleObjectTypeBasePolicyRule}
}

// ObjectTypeBasicCopyOption enumerates the values for object type basic copy option.
type ObjectTypeBasicCopyOption string

const (
	// ObjectTypeBasicCopyOptionObjectTypeCopyOnExpiryOption ...
	ObjectTypeBasicCopyOptionObjectTypeCopyOnExpiryOption ObjectTypeBasicCopyOption = "CopyOnExpiryOption"
	// ObjectTypeBasicCopyOptionObjectTypeCopyOption ...
	ObjectTypeBasicCopyOptionObjectTypeCopyOption ObjectTypeBasicCopyOption = "CopyOption"
	// ObjectTypeBasicCopyOptionObjectTypeCustomCopyOption ...
	ObjectTypeBasicCopyOptionObjectTypeCustomCopyOption ObjectTypeBasicCopyOption = "CustomCopyOption"
	// ObjectTypeBasicCopyOptionObjectTypeImmediateCopyOption ...
	ObjectTypeBasicCopyOptionObjectTypeImmediateCopyOption ObjectTypeBasicCopyOption = "ImmediateCopyOption"
)

// PossibleObjectTypeBasicCopyOptionValues returns an array of possible values for the ObjectTypeBasicCopyOption const type.
func PossibleObjectTypeBasicCopyOptionValues() []ObjectTypeBasicCopyOption {
	return []ObjectTypeBasicCopyOption{ObjectTypeBasicCopyOptionObjectTypeCopyOnExpiryOption, ObjectTypeBasicCopyOptionObjectTypeCopyOption, ObjectTypeBasicCopyOptionObjectTypeCustomCopyOption, ObjectTypeBasicCopyOptionObjectTypeImmediateCopyOption}
}

// ObjectTypeBasicDataStoreParameters enumerates the values for object type basic data store parameters.
type ObjectTypeBasicDataStoreParameters string

const (
	// ObjectTypeBasicDataStoreParametersObjectTypeAzureOperationalStoreParameters ...
	ObjectTypeBasicDataStoreParametersObjectTypeAzureOperationalStoreParameters ObjectTypeBasicDataStoreParameters = "AzureOperationalStoreParameters"
	// ObjectTypeBasicDataStoreParametersObjectTypeDataStoreParameters ...
	ObjectTypeBasicDataStoreParametersObjectTypeDataStoreParameters ObjectTypeBasicDataStoreParameters = "DataStoreParameters"
)

// PossibleObjectTypeBasicDataStoreParametersValues returns an array of possible values for the ObjectTypeBasicDataStoreParameters const type.
func PossibleObjectTypeBasicDataStoreParametersValues() []ObjectTypeBasicDataStoreParameters {
	return []ObjectTypeBasicDataStoreParameters{ObjectTypeBasicDataStoreParametersObjectTypeAzureOperationalStoreParameters, ObjectTypeBasicDataStoreParametersObjectTypeDataStoreParameters}
}

// ObjectTypeBasicDeleteOption enumerates the values for object type basic delete option.
type ObjectTypeBasicDeleteOption string

const (
	// ObjectTypeBasicDeleteOptionObjectTypeAbsoluteDeleteOption ...
	ObjectTypeBasicDeleteOptionObjectTypeAbsoluteDeleteOption ObjectTypeBasicDeleteOption = "AbsoluteDeleteOption"
	// ObjectTypeBasicDeleteOptionObjectTypeDeleteOption ...
	ObjectTypeBasicDeleteOptionObjectTypeDeleteOption ObjectTypeBasicDeleteOption = "DeleteOption"
)

// PossibleObjectTypeBasicDeleteOptionValues returns an array of possible values for the ObjectTypeBasicDeleteOption const type.
func PossibleObjectTypeBasicDeleteOptionValues() []ObjectTypeBasicDeleteOption {
	return []ObjectTypeBasicDeleteOption{ObjectTypeBasicDeleteOptionObjectTypeAbsoluteDeleteOption, ObjectTypeBasicDeleteOptionObjectTypeDeleteOption}
}

// ObjectTypeBasicFeatureValidationRequestBase enumerates the values for object type basic feature validation
// request base.
type ObjectTypeBasicFeatureValidationRequestBase string

const (
	// ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequest ...
	ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequest ObjectTypeBasicFeatureValidationRequestBase = "FeatureValidationRequest"
	// ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequestBase ...
	ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequestBase ObjectTypeBasicFeatureValidationRequestBase = "FeatureValidationRequestBase"
)

// PossibleObjectTypeBasicFeatureValidationRequestBaseValues returns an array of possible values for the ObjectTypeBasicFeatureValidationRequestBase const type.
func PossibleObjectTypeBasicFeatureValidationRequestBaseValues() []ObjectTypeBasicFeatureValidationRequestBase {
	return []ObjectTypeBasicFeatureValidationRequestBase{ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequest, ObjectTypeBasicFeatureValidationRequestBaseObjectTypeFeatureValidationRequestBase}
}

// ObjectTypeBasicFeatureValidationResponseBase enumerates the values for object type basic feature validation
// response base.
type ObjectTypeBasicFeatureValidationResponseBase string

const (
	// ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponse ...
	ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponse ObjectTypeBasicFeatureValidationResponseBase = "FeatureValidationResponse"
	// ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponseBase ...
	ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponseBase ObjectTypeBasicFeatureValidationResponseBase = "FeatureValidationResponseBase"
)

// PossibleObjectTypeBasicFeatureValidationResponseBaseValues returns an array of possible values for the ObjectTypeBasicFeatureValidationResponseBase const type.
func PossibleObjectTypeBasicFeatureValidationResponseBaseValues() []ObjectTypeBasicFeatureValidationResponseBase {
	return []ObjectTypeBasicFeatureValidationResponseBase{ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponse, ObjectTypeBasicFeatureValidationResponseBaseObjectTypeFeatureValidationResponseBase}
}

// ObjectTypeBasicItemLevelRestoreCriteria enumerates the values for object type basic item level restore
// criteria.
type ObjectTypeBasicItemLevelRestoreCriteria string

const (
	// ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeItemLevelRestoreCriteria ...
	ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeItemLevelRestoreCriteria ObjectTypeBasicItemLevelRestoreCriteria = "ItemLevelRestoreCriteria"
	// ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeRangeBasedItemLevelRestoreCriteria ...
	ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeRangeBasedItemLevelRestoreCriteria ObjectTypeBasicItemLevelRestoreCriteria = "RangeBasedItemLevelRestoreCriteria"
)

// PossibleObjectTypeBasicItemLevelRestoreCriteriaValues returns an array of possible values for the ObjectTypeBasicItemLevelRestoreCriteria const type.
func PossibleObjectTypeBasicItemLevelRestoreCriteriaValues() []ObjectTypeBasicItemLevelRestoreCriteria {
	return []ObjectTypeBasicItemLevelRestoreCriteria{ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeItemLevelRestoreCriteria, ObjectTypeBasicItemLevelRestoreCriteriaObjectTypeRangeBasedItemLevelRestoreCriteria}
}

// ObjectTypeBasicOperationExtendedInfo enumerates the values for object type basic operation extended info.
type ObjectTypeBasicOperationExtendedInfo string

const (
	// ObjectTypeBasicOperationExtendedInfoObjectTypeOperationExtendedInfo ...
	ObjectTypeBasicOperationExtendedInfoObjectTypeOperationExtendedInfo ObjectTypeBasicOperationExtendedInfo = "OperationExtendedInfo"
	// ObjectTypeBasicOperationExtendedInfoObjectTypeOperationJobExtendedInfo ...
	ObjectTypeBasicOperationExtendedInfoObjectTypeOperationJobExtendedInfo ObjectTypeBasicOperationExtendedInfo = "OperationJobExtendedInfo"
)

// PossibleObjectTypeBasicOperationExtendedInfoValues returns an array of possible values for the ObjectTypeBasicOperationExtendedInfo const type.
func PossibleObjectTypeBasicOperationExtendedInfoValues() []ObjectTypeBasicOperationExtendedInfo {
	return []ObjectTypeBasicOperationExtendedInfo{ObjectTypeBasicOperationExtendedInfoObjectTypeOperationExtendedInfo, ObjectTypeBasicOperationExtendedInfoObjectTypeOperationJobExtendedInfo}
}

// ObjectTypeBasicRestoreTargetInfoBase enumerates the values for object type basic restore target info base.
type ObjectTypeBasicRestoreTargetInfoBase string

const (
	// ObjectTypeBasicRestoreTargetInfoBaseObjectTypeItemLevelRestoreTargetInfo ...
	ObjectTypeBasicRestoreTargetInfoBaseObjectTypeItemLevelRestoreTargetInfo ObjectTypeBasicRestoreTargetInfoBase = "ItemLevelRestoreTargetInfo"
	// ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreFilesTargetInfo ...
	ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreFilesTargetInfo ObjectTypeBasicRestoreTargetInfoBase = "RestoreFilesTargetInfo"
	// ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfo ...
	ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfo ObjectTypeBasicRestoreTargetInfoBase = "RestoreTargetInfo"
	// ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfoBase ...
	ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfoBase ObjectTypeBasicRestoreTargetInfoBase = "RestoreTargetInfoBase"
)

// PossibleObjectTypeBasicRestoreTargetInfoBaseValues returns an array of possible values for the ObjectTypeBasicRestoreTargetInfoBase const type.
func PossibleObjectTypeBasicRestoreTargetInfoBaseValues() []ObjectTypeBasicRestoreTargetInfoBase {
	return []ObjectTypeBasicRestoreTargetInfoBase{ObjectTypeBasicRestoreTargetInfoBaseObjectTypeItemLevelRestoreTargetInfo, ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreFilesTargetInfo, ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfo, ObjectTypeBasicRestoreTargetInfoBaseObjectTypeRestoreTargetInfoBase}
}

// ObjectTypeBasicTriggerContext enumerates the values for object type basic trigger context.
type ObjectTypeBasicTriggerContext string

const (
	// ObjectTypeBasicTriggerContextObjectTypeAdhocBasedTriggerContext ...
	ObjectTypeBasicTriggerContextObjectTypeAdhocBasedTriggerContext ObjectTypeBasicTriggerContext = "AdhocBasedTriggerContext"
	// ObjectTypeBasicTriggerContextObjectTypeScheduleBasedTriggerContext ...
	ObjectTypeBasicTriggerContextObjectTypeScheduleBasedTriggerContext ObjectTypeBasicTriggerContext = "ScheduleBasedTriggerContext"
	// ObjectTypeBasicTriggerContextObjectTypeTriggerContext ...
	ObjectTypeBasicTriggerContextObjectTypeTriggerContext ObjectTypeBasicTriggerContext = "TriggerContext"
)

// PossibleObjectTypeBasicTriggerContextValues returns an array of possible values for the ObjectTypeBasicTriggerContext const type.
func PossibleObjectTypeBasicTriggerContextValues() []ObjectTypeBasicTriggerContext {
	return []ObjectTypeBasicTriggerContext{ObjectTypeBasicTriggerContextObjectTypeAdhocBasedTriggerContext, ObjectTypeBasicTriggerContextObjectTypeScheduleBasedTriggerContext, ObjectTypeBasicTriggerContextObjectTypeTriggerContext}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning ...
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUnknown ...
	ProvisioningStateUnknown ProvisioningState = "Unknown"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateFailed, ProvisioningStateProvisioning, ProvisioningStateSucceeded, ProvisioningStateUnknown, ProvisioningStateUpdating}
}

// RehydrationPriority enumerates the values for rehydration priority.
type RehydrationPriority string

const (
	// RehydrationPriorityHigh ...
	RehydrationPriorityHigh RehydrationPriority = "High"
	// RehydrationPriorityInvalid ...
	RehydrationPriorityInvalid RehydrationPriority = "Invalid"
	// RehydrationPriorityStandard ...
	RehydrationPriorityStandard RehydrationPriority = "Standard"
)

// PossibleRehydrationPriorityValues returns an array of possible values for the RehydrationPriority const type.
func PossibleRehydrationPriorityValues() []RehydrationPriority {
	return []RehydrationPriority{RehydrationPriorityHigh, RehydrationPriorityInvalid, RehydrationPriorityStandard}
}

// RehydrationStatus enumerates the values for rehydration status.
type RehydrationStatus string

const (
	// RehydrationStatusCOMPLETED ...
	RehydrationStatusCOMPLETED RehydrationStatus = "COMPLETED"
	// RehydrationStatusCREATEINPROGRESS ...
	RehydrationStatusCREATEINPROGRESS RehydrationStatus = "CREATE_IN_PROGRESS"
	// RehydrationStatusDELETED ...
	RehydrationStatusDELETED RehydrationStatus = "DELETED"
	// RehydrationStatusDELETEINPROGRESS ...
	RehydrationStatusDELETEINPROGRESS RehydrationStatus = "DELETE_IN_PROGRESS"
	// RehydrationStatusFAILED ...
	RehydrationStatusFAILED RehydrationStatus = "FAILED"
)

// PossibleRehydrationStatusValues returns an array of possible values for the RehydrationStatus const type.
func PossibleRehydrationStatusValues() []RehydrationStatus {
	return []RehydrationStatus{RehydrationStatusCOMPLETED, RehydrationStatusCREATEINPROGRESS, RehydrationStatusDELETED, RehydrationStatusDELETEINPROGRESS, RehydrationStatusFAILED}
}

// ResourceMoveState enumerates the values for resource move state.
type ResourceMoveState string

const (
	// ResourceMoveStateCommitFailed ...
	ResourceMoveStateCommitFailed ResourceMoveState = "CommitFailed"
	// ResourceMoveStateCommitTimedout ...
	ResourceMoveStateCommitTimedout ResourceMoveState = "CommitTimedout"
	// ResourceMoveStateCriticalFailure ...
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	// ResourceMoveStateFailed ...
	ResourceMoveStateFailed ResourceMoveState = "Failed"
	// ResourceMoveStateInProgress ...
	ResourceMoveStateInProgress ResourceMoveState = "InProgress"
	// ResourceMoveStateMoveSucceeded ...
	ResourceMoveStateMoveSucceeded ResourceMoveState = "MoveSucceeded"
	// ResourceMoveStatePartialSuccess ...
	ResourceMoveStatePartialSuccess ResourceMoveState = "PartialSuccess"
	// ResourceMoveStatePrepareFailed ...
	ResourceMoveStatePrepareFailed ResourceMoveState = "PrepareFailed"
	// ResourceMoveStatePrepareTimedout ...
	ResourceMoveStatePrepareTimedout ResourceMoveState = "PrepareTimedout"
	// ResourceMoveStateUnknown ...
	ResourceMoveStateUnknown ResourceMoveState = "Unknown"
)

// PossibleResourceMoveStateValues returns an array of possible values for the ResourceMoveState const type.
func PossibleResourceMoveStateValues() []ResourceMoveState {
	return []ResourceMoveState{ResourceMoveStateCommitFailed, ResourceMoveStateCommitTimedout, ResourceMoveStateCriticalFailure, ResourceMoveStateFailed, ResourceMoveStateInProgress, ResourceMoveStateMoveSucceeded, ResourceMoveStatePartialSuccess, ResourceMoveStatePrepareFailed, ResourceMoveStatePrepareTimedout, ResourceMoveStateUnknown}
}

// RestoreSourceDataStoreType enumerates the values for restore source data store type.
type RestoreSourceDataStoreType string

const (
	// RestoreSourceDataStoreTypeArchiveStore ...
	RestoreSourceDataStoreTypeArchiveStore RestoreSourceDataStoreType = "ArchiveStore"
	// RestoreSourceDataStoreTypeOperationalStore ...
	RestoreSourceDataStoreTypeOperationalStore RestoreSourceDataStoreType = "OperationalStore"
	// RestoreSourceDataStoreTypeVaultStore ...
	RestoreSourceDataStoreTypeVaultStore RestoreSourceDataStoreType = "VaultStore"
)

// PossibleRestoreSourceDataStoreTypeValues returns an array of possible values for the RestoreSourceDataStoreType const type.
func PossibleRestoreSourceDataStoreTypeValues() []RestoreSourceDataStoreType {
	return []RestoreSourceDataStoreType{RestoreSourceDataStoreTypeArchiveStore, RestoreSourceDataStoreTypeOperationalStore, RestoreSourceDataStoreTypeVaultStore}
}

// RestoreTargetLocationType enumerates the values for restore target location type.
type RestoreTargetLocationType string

const (
	// RestoreTargetLocationTypeAzureBlobs ...
	RestoreTargetLocationTypeAzureBlobs RestoreTargetLocationType = "AzureBlobs"
	// RestoreTargetLocationTypeAzureFiles ...
	RestoreTargetLocationTypeAzureFiles RestoreTargetLocationType = "AzureFiles"
	// RestoreTargetLocationTypeInvalid ...
	RestoreTargetLocationTypeInvalid RestoreTargetLocationType = "Invalid"
)

// PossibleRestoreTargetLocationTypeValues returns an array of possible values for the RestoreTargetLocationType const type.
func PossibleRestoreTargetLocationTypeValues() []RestoreTargetLocationType {
	return []RestoreTargetLocationType{RestoreTargetLocationTypeAzureBlobs, RestoreTargetLocationTypeAzureFiles, RestoreTargetLocationTypeInvalid}
}

// SecretStoreType enumerates the values for secret store type.
type SecretStoreType string

const (
	// SecretStoreTypeAzureKeyVault ...
	SecretStoreTypeAzureKeyVault SecretStoreType = "AzureKeyVault"
	// SecretStoreTypeInvalid ...
	SecretStoreTypeInvalid SecretStoreType = "Invalid"
)

// PossibleSecretStoreTypeValues returns an array of possible values for the SecretStoreType const type.
func PossibleSecretStoreTypeValues() []SecretStoreType {
	return []SecretStoreType{SecretStoreTypeAzureKeyVault, SecretStoreTypeInvalid}
}

// SourceDataStoreType enumerates the values for source data store type.
type SourceDataStoreType string

const (
	// SourceDataStoreTypeArchiveStore ...
	SourceDataStoreTypeArchiveStore SourceDataStoreType = "ArchiveStore"
	// SourceDataStoreTypeSnapshotStore ...
	SourceDataStoreTypeSnapshotStore SourceDataStoreType = "SnapshotStore"
	// SourceDataStoreTypeVaultStore ...
	SourceDataStoreTypeVaultStore SourceDataStoreType = "VaultStore"
)

// PossibleSourceDataStoreTypeValues returns an array of possible values for the SourceDataStoreType const type.
func PossibleSourceDataStoreTypeValues() []SourceDataStoreType {
	return []SourceDataStoreType{SourceDataStoreTypeArchiveStore, SourceDataStoreTypeSnapshotStore, SourceDataStoreTypeVaultStore}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusConfiguringProtection ...
	StatusConfiguringProtection Status = "ConfiguringProtection"
	// StatusConfiguringProtectionFailed ...
	StatusConfiguringProtectionFailed Status = "ConfiguringProtectionFailed"
	// StatusProtectionConfigured ...
	StatusProtectionConfigured Status = "ProtectionConfigured"
	// StatusProtectionStopped ...
	StatusProtectionStopped Status = "ProtectionStopped"
	// StatusSoftDeleted ...
	StatusSoftDeleted Status = "SoftDeleted"
	// StatusSoftDeleting ...
	StatusSoftDeleting Status = "SoftDeleting"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusConfiguringProtection, StatusConfiguringProtectionFailed, StatusProtectionConfigured, StatusProtectionStopped, StatusSoftDeleted, StatusSoftDeleting}
}

// StorageSettingStoreTypes enumerates the values for storage setting store types.
type StorageSettingStoreTypes string

const (
	// StorageSettingStoreTypesArchiveStore ...
	StorageSettingStoreTypesArchiveStore StorageSettingStoreTypes = "ArchiveStore"
	// StorageSettingStoreTypesSnapshotStore ...
	StorageSettingStoreTypesSnapshotStore StorageSettingStoreTypes = "SnapshotStore"
	// StorageSettingStoreTypesVaultStore ...
	StorageSettingStoreTypesVaultStore StorageSettingStoreTypes = "VaultStore"
)

// PossibleStorageSettingStoreTypesValues returns an array of possible values for the StorageSettingStoreTypes const type.
func PossibleStorageSettingStoreTypesValues() []StorageSettingStoreTypes {
	return []StorageSettingStoreTypes{StorageSettingStoreTypesArchiveStore, StorageSettingStoreTypesSnapshotStore, StorageSettingStoreTypesVaultStore}
}

// StorageSettingTypes enumerates the values for storage setting types.
type StorageSettingTypes string

const (
	// StorageSettingTypesGeoRedundant ...
	StorageSettingTypesGeoRedundant StorageSettingTypes = "GeoRedundant"
	// StorageSettingTypesLocallyRedundant ...
	StorageSettingTypesLocallyRedundant StorageSettingTypes = "LocallyRedundant"
)

// PossibleStorageSettingTypesValues returns an array of possible values for the StorageSettingTypes const type.
func PossibleStorageSettingTypesValues() []StorageSettingTypes {
	return []StorageSettingTypes{StorageSettingTypesGeoRedundant, StorageSettingTypesLocallyRedundant}
}

// WeekNumber enumerates the values for week number.
type WeekNumber string

const (
	// WeekNumberFirst ...
	WeekNumberFirst WeekNumber = "First"
	// WeekNumberFourth ...
	WeekNumberFourth WeekNumber = "Fourth"
	// WeekNumberLast ...
	WeekNumberLast WeekNumber = "Last"
	// WeekNumberSecond ...
	WeekNumberSecond WeekNumber = "Second"
	// WeekNumberThird ...
	WeekNumberThird WeekNumber = "Third"
)

// PossibleWeekNumberValues returns an array of possible values for the WeekNumber const type.
func PossibleWeekNumberValues() []WeekNumber {
	return []WeekNumber{WeekNumberFirst, WeekNumberFourth, WeekNumberLast, WeekNumberSecond, WeekNumberThird}
}