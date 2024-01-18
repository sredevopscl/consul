// Code generated by protoc-gen-resource-types. DO NOT EDIT.

package catalogv2beta1

import (
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName = "catalog"
	Version   = "v2beta1"

	ComputedFailoverPolicyKind = "ComputedFailoverPolicy"
	FailoverPolicyKind         = "FailoverPolicy"
	HealthChecksKind           = "HealthChecks"
	HealthStatusKind           = "HealthStatus"
	NodeKind                   = "Node"
	NodeHealthStatusKind       = "NodeHealthStatus"
	ServiceKind                = "Service"
	ServiceEndpointsKind       = "ServiceEndpoints"
	VirtualIPsKind             = "VirtualIPs"
	WorkloadKind               = "Workload"
)

var (
	ComputedFailoverPolicyType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ComputedFailoverPolicyKind,
	}

	FailoverPolicyType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         FailoverPolicyKind,
	}

	HealthChecksType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         HealthChecksKind,
	}

	HealthStatusType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         HealthStatusKind,
	}

	NodeType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         NodeKind,
	}

	NodeHealthStatusType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         NodeHealthStatusKind,
	}

	ServiceType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ServiceKind,
	}

	ServiceEndpointsType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         ServiceEndpointsKind,
	}

	VirtualIPsType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         VirtualIPsKind,
	}

	WorkloadType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         WorkloadKind,
	}
)
