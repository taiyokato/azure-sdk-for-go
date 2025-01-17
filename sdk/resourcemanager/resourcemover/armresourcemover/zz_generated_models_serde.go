//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AvailabilitySetResourceSettings.
func (a AvailabilitySetResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "faultDomain", a.FaultDomain)
	objectMap["resourceType"] = "Microsoft.Compute/availabilitySets"
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "targetResourceName", a.TargetResourceName)
	populate(objectMap, "updateDomain", a.UpdateDomain)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AvailabilitySetResourceSettings.
func (a *AvailabilitySetResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "faultDomain":
			err = unpopulate(val, "FaultDomain", &a.FaultDomain)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &a.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &a.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &a.TargetResourceName)
			delete(rawMsg, key)
		case "updateDomain":
			err = unpopulate(val, "UpdateDomain", &a.UpdateDomain)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BulkRemoveRequest.
func (b BulkRemoveRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", b.MoveResourceInputType)
	populate(objectMap, "moveResources", b.MoveResources)
	populate(objectMap, "validateOnly", b.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CommitRequest.
func (c CommitRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", c.MoveResourceInputType)
	populate(objectMap, "moveResources", c.MoveResources)
	populate(objectMap, "validateOnly", c.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiscardRequest.
func (d DiscardRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", d.MoveResourceInputType)
	populate(objectMap, "moveResources", d.MoveResources)
	populate(objectMap, "validateOnly", d.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DiskEncryptionSetResourceSettings.
func (d DiskEncryptionSetResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Compute/diskEncryptionSets"
	populate(objectMap, "targetResourceName", d.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiskEncryptionSetResourceSettings.
func (d *DiskEncryptionSetResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &d.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &d.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KeyVaultResourceSettings.
func (k KeyVaultResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.KeyVault/vaults"
	populate(objectMap, "targetResourceName", k.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyVaultResourceSettings.
func (k *KeyVaultResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", k, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &k.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &k.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", k, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LoadBalancerResourceSettings.
func (l LoadBalancerResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "backendAddressPools", l.BackendAddressPools)
	populate(objectMap, "frontendIPConfigurations", l.FrontendIPConfigurations)
	objectMap["resourceType"] = "Microsoft.Network/loadBalancers"
	populate(objectMap, "sku", l.SKU)
	populate(objectMap, "tags", l.Tags)
	populate(objectMap, "targetResourceName", l.TargetResourceName)
	populate(objectMap, "zones", l.Zones)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LoadBalancerResourceSettings.
func (l *LoadBalancerResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backendAddressPools":
			err = unpopulate(val, "BackendAddressPools", &l.BackendAddressPools)
			delete(rawMsg, key)
		case "frontendIPConfigurations":
			err = unpopulate(val, "FrontendIPConfigurations", &l.FrontendIPConfigurations)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &l.ResourceType)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &l.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &l.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &l.TargetResourceName)
			delete(rawMsg, key)
		case "zones":
			err = unpopulate(val, "Zones", &l.Zones)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MoveCollection.
func (m MoveCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", m.Etag)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveResourceErrorBody.
func (m MoveResourceErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", m.Code)
	populate(objectMap, "details", m.Details)
	populate(objectMap, "message", m.Message)
	populate(objectMap, "target", m.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MoveResourceProperties.
func (m MoveResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dependsOn", m.DependsOn)
	populate(objectMap, "dependsOnOverrides", m.DependsOnOverrides)
	populate(objectMap, "errors", m.Errors)
	populate(objectMap, "existingTargetId", m.ExistingTargetID)
	populate(objectMap, "isResolveRequired", m.IsResolveRequired)
	populate(objectMap, "moveStatus", m.MoveStatus)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "resourceSettings", m.ResourceSettings)
	populate(objectMap, "sourceId", m.SourceID)
	populate(objectMap, "sourceResourceSettings", m.SourceResourceSettings)
	populate(objectMap, "targetId", m.TargetID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MoveResourceProperties.
func (m *MoveResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dependsOn":
			err = unpopulate(val, "DependsOn", &m.DependsOn)
			delete(rawMsg, key)
		case "dependsOnOverrides":
			err = unpopulate(val, "DependsOnOverrides", &m.DependsOnOverrides)
			delete(rawMsg, key)
		case "errors":
			err = unpopulate(val, "Errors", &m.Errors)
			delete(rawMsg, key)
		case "existingTargetId":
			err = unpopulate(val, "ExistingTargetID", &m.ExistingTargetID)
			delete(rawMsg, key)
		case "isResolveRequired":
			err = unpopulate(val, "IsResolveRequired", &m.IsResolveRequired)
			delete(rawMsg, key)
		case "moveStatus":
			err = unpopulate(val, "MoveStatus", &m.MoveStatus)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &m.ProvisioningState)
			delete(rawMsg, key)
		case "resourceSettings":
			m.ResourceSettings, err = unmarshalResourceSettingsClassification(val)
			delete(rawMsg, key)
		case "sourceId":
			err = unpopulate(val, "SourceID", &m.SourceID)
			delete(rawMsg, key)
		case "sourceResourceSettings":
			m.SourceResourceSettings, err = unmarshalResourceSettingsClassification(val)
			delete(rawMsg, key)
		case "targetId":
			err = unpopulate(val, "TargetID", &m.TargetID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkInterfaceResourceSettings.
func (n NetworkInterfaceResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "enableAcceleratedNetworking", n.EnableAcceleratedNetworking)
	populate(objectMap, "ipConfigurations", n.IPConfigurations)
	objectMap["resourceType"] = "Microsoft.Network/networkInterfaces"
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "targetResourceName", n.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkInterfaceResourceSettings.
func (n *NetworkInterfaceResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enableAcceleratedNetworking":
			err = unpopulate(val, "EnableAcceleratedNetworking", &n.EnableAcceleratedNetworking)
			delete(rawMsg, key)
		case "ipConfigurations":
			err = unpopulate(val, "IPConfigurations", &n.IPConfigurations)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &n.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &n.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &n.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityGroupResourceSettings.
func (n NetworkSecurityGroupResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Network/networkSecurityGroups"
	populate(objectMap, "securityRules", n.SecurityRules)
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "targetResourceName", n.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkSecurityGroupResourceSettings.
func (n *NetworkSecurityGroupResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &n.ResourceType)
			delete(rawMsg, key)
		case "securityRules":
			err = unpopulate(val, "SecurityRules", &n.SecurityRules)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &n.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &n.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NicIPConfigurationResourceSettings.
func (n NicIPConfigurationResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "loadBalancerBackendAddressPools", n.LoadBalancerBackendAddressPools)
	populate(objectMap, "loadBalancerNatRules", n.LoadBalancerNatRules)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "primary", n.Primary)
	populate(objectMap, "privateIpAddress", n.PrivateIPAddress)
	populate(objectMap, "privateIpAllocationMethod", n.PrivateIPAllocationMethod)
	populate(objectMap, "publicIp", n.PublicIP)
	populate(objectMap, "subnet", n.Subnet)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrepareRequest.
func (p PrepareRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", p.MoveResourceInputType)
	populate(objectMap, "moveResources", p.MoveResources)
	populate(objectMap, "validateOnly", p.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PublicIPAddressResourceSettings.
func (p PublicIPAddressResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainNameLabel", p.DomainNameLabel)
	populate(objectMap, "fqdn", p.Fqdn)
	populate(objectMap, "publicIpAllocationMethod", p.PublicIPAllocationMethod)
	objectMap["resourceType"] = "Microsoft.Network/publicIPAddresses"
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "targetResourceName", p.TargetResourceName)
	populate(objectMap, "zones", p.Zones)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublicIPAddressResourceSettings.
func (p *PublicIPAddressResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "domainNameLabel":
			err = unpopulate(val, "DomainNameLabel", &p.DomainNameLabel)
			delete(rawMsg, key)
		case "fqdn":
			err = unpopulate(val, "Fqdn", &p.Fqdn)
			delete(rawMsg, key)
		case "publicIpAllocationMethod":
			err = unpopulate(val, "PublicIPAllocationMethod", &p.PublicIPAllocationMethod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &p.ResourceType)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &p.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &p.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &p.TargetResourceName)
			delete(rawMsg, key)
		case "zones":
			err = unpopulate(val, "Zones", &p.Zones)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceGroupResourceSettings.
func (r ResourceGroupResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "resourceGroups"
	populate(objectMap, "targetResourceName", r.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceGroupResourceSettings.
func (r *ResourceGroupResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &r.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &r.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceMoveRequest.
func (r ResourceMoveRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "moveResourceInputType", r.MoveResourceInputType)
	populate(objectMap, "moveResources", r.MoveResources)
	populate(objectMap, "validateOnly", r.ValidateOnly)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SQLDatabaseResourceSettings.
func (s SQLDatabaseResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers/databases"
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	populate(objectMap, "zoneRedundant", s.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLDatabaseResourceSettings.
func (s *SQLDatabaseResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &s.TargetResourceName)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, "ZoneRedundant", &s.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLElasticPoolResourceSettings.
func (s SQLElasticPoolResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers/elasticPools"
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	populate(objectMap, "zoneRedundant", s.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLElasticPoolResourceSettings.
func (s *SQLElasticPoolResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &s.TargetResourceName)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, "ZoneRedundant", &s.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SQLServerResourceSettings.
func (s SQLServerResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Sql/servers"
	populate(objectMap, "targetResourceName", s.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SQLServerResourceSettings.
func (s *SQLServerResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &s.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateMoveCollectionRequest.
func (u UpdateMoveCollectionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", u.Identity)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualMachineResourceSettings.
func (v VirtualMachineResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["resourceType"] = "Microsoft.Compute/virtualMachines"
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "targetAvailabilitySetId", v.TargetAvailabilitySetID)
	populate(objectMap, "targetAvailabilityZone", v.TargetAvailabilityZone)
	populate(objectMap, "targetResourceName", v.TargetResourceName)
	populate(objectMap, "targetVmSize", v.TargetVMSize)
	populate(objectMap, "userManagedIdentities", v.UserManagedIdentities)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualMachineResourceSettings.
func (v *VirtualMachineResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
			err = unpopulate(val, "ResourceType", &v.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		case "targetAvailabilitySetId":
			err = unpopulate(val, "TargetAvailabilitySetID", &v.TargetAvailabilitySetID)
			delete(rawMsg, key)
		case "targetAvailabilityZone":
			err = unpopulate(val, "TargetAvailabilityZone", &v.TargetAvailabilityZone)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &v.TargetResourceName)
			delete(rawMsg, key)
		case "targetVmSize":
			err = unpopulate(val, "TargetVMSize", &v.TargetVMSize)
			delete(rawMsg, key)
		case "userManagedIdentities":
			err = unpopulate(val, "UserManagedIdentities", &v.UserManagedIdentities)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkResourceSettings.
func (v VirtualNetworkResourceSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "addressSpace", v.AddressSpace)
	populate(objectMap, "dnsServers", v.DNSServers)
	populate(objectMap, "enableDdosProtection", v.EnableDdosProtection)
	objectMap["resourceType"] = "Microsoft.Network/virtualNetworks"
	populate(objectMap, "subnets", v.Subnets)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "targetResourceName", v.TargetResourceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkResourceSettings.
func (v *VirtualNetworkResourceSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "addressSpace":
			err = unpopulate(val, "AddressSpace", &v.AddressSpace)
			delete(rawMsg, key)
		case "dnsServers":
			err = unpopulate(val, "DNSServers", &v.DNSServers)
			delete(rawMsg, key)
		case "enableDdosProtection":
			err = unpopulate(val, "EnableDdosProtection", &v.EnableDdosProtection)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &v.ResourceType)
			delete(rawMsg, key)
		case "subnets":
			err = unpopulate(val, "Subnets", &v.Subnets)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		case "targetResourceName":
			err = unpopulate(val, "TargetResourceName", &v.TargetResourceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
