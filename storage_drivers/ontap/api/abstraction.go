// Copyright 2021 NetApp, Inc. All Rights Reserved.

package api

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	. "github.com/netapp/trident/logger"
)

////////////////////////////////////////////////////////////////////////////////////////////
///             _____________________
///            |   <<Interface>>    |
///            |       ONTAPI       |
///            |____________________|
///                ^             ^
///     Implements |             | Implements
///   ____________________    ____________________
///  |  ONTAPAPIREST     |   |  ONTAPAPIZAPI     |
///  |___________________|   |___________________|
///  | +API: RestClient  |   | +API: *Client     |
///  |___________________|   |___________________|
///
////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////
// Drivers that offer dual support are to call ONTAP REST or ZAPI's
// via abstraction layer (ONTAPI interface)
////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////
// Abstraction layer
////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	APIResponsePassed = NewAPIResponse("passed", "", "")
)

type OntapAPI interface {
	APIVersion(context.Context) (string, error)

	EmsAutosupportLog(ctx context.Context, driverName string, appVersion string, autoSupport bool, category string,
		computerName string, eventDescription string, eventID int, eventSource string, logLevel int)

	ExportPolicyCreate(ctx context.Context, policy string) error
	ExportPolicyDestroy(ctx context.Context, policy string) error
	ExportPolicyExists(ctx context.Context, policyName string) (bool, error)
	ExportRuleCreate(ctx context.Context, policyName, desiredPolicyRule string) error
	ExportRuleDestroy(ctx context.Context, policyName string, ruleIndex int) error
	ExportRuleList(ctx context.Context, policyName string) (map[string]int, error)

	FlexgroupCreate(ctx context.Context, volume Volume) error
	FlexgroupExists(ctx context.Context, volumeName string) (bool, error)
	FlexgroupInfo(ctx context.Context, volumeName string) (*Volume, error)
	FlexgroupDisableSnapshotDirectoryAccess(ctx context.Context, volumeName string) error
	FlexgroupSetComment(ctx context.Context, volumeNameInternal, volumeNameExternal, comment string) error
	FlexgroupModifyUnixPermissions(ctx context.Context, volumeNameInternal, volumeNameExternal, unixPermissions string) error
	FlexgroupMount(ctx context.Context, name, junctionPath string) error
	FlexgroupListByPrefix(ctx context.Context, prefix string) (Volumes, error)
	FlexgroupDestroy(ctx context.Context, volumeName string, force bool) error
	FlexgroupSetSize(ctx context.Context, name, newSize string) error
	FlexgroupSize(ctx context.Context, volumeName string) (uint64, error)
	FlexgroupUnmount(ctx context.Context, name string, force bool) error
	FlexgroupUsedSize(ctx context.Context, volumeName string) (int, error)
	FlexgroupModifyExportPolicy(ctx context.Context, volumeName, policyName string) error
	FlexgroupSnapshotCreate(ctx context.Context, snapshotName, sourceVolume string) error
	FlexgroupSetQosPolicyGroupName(ctx context.Context, name string, qos QosPolicyGroup) error
	FlexgroupCloneSplitStart(ctx context.Context, cloneName string) error
	FlexgroupSnapshotList(ctx context.Context, sourceVolume string) (Snapshots, error)
	FlexgroupSnapshotDelete(ctx context.Context, snapshotName, sourceVolume string) error
	GetSVMAggregateAttributes(ctx context.Context) (map[string]string, error)
	GetSVMAggregateNames(ctx context.Context) ([]string, error)
	GetSVMAggregateSpace(ctx context.Context, aggregate string) ([]SVMAggregateSpace, error)

	NetInterfaceGetDataLIFs(ctx context.Context, protocol string) ([]string, error)
	NodeListSerialNumbers(ctx context.Context) ([]string, error)

	SnapmirrorDeleteViaDestination(localFlexvolName, localSVMName string) error
	IsSVMDRCapable(ctx context.Context) (bool, error)

	SnapshotRestoreVolume(ctx context.Context, snapshotName, sourceVolume string) error
	SnapshotRestoreFlexgroup(ctx context.Context, snapshotName, sourceVolume string) error

	SupportsFeature(ctx context.Context, feature feature) bool
	ValidateAPIVersion(ctx context.Context) error

	VolumeCloneCreate(ctx context.Context, cloneName, sourceName, snapshot string, async bool) error
	VolumeCloneSplitStart(ctx context.Context, cloneName string) error

	VolumeCreate(ctx context.Context, volume Volume) error
	VolumeDestroy(ctx context.Context, volumeName string, force bool) error
	VolumeDisableSnapshotDirectoryAccess(ctx context.Context, name string) error
	VolumeExists(ctx context.Context, volumeName string) (bool, error)
	VolumeInfo(ctx context.Context, volumeName string) (*Volume, error)
	VolumeListByPrefix(ctx context.Context, prefix string) (Volumes, error)
	VolumeListBySnapshotParent(ctx context.Context, snapshotName, sourceVolume string) (VolumeNameList, error)
	VolumeModifyExportPolicy(ctx context.Context, volumeName, policyName string) error
	VolumeModifyUnixPermissions(ctx context.Context, volumeNameInternal, volumeNameExternal, unixPermissions string) error
	VolumeMount(ctx context.Context, name, junctionPath string) error
	VolumeRename(ctx context.Context, originalName, newName string) error
	VolumeSetComment(ctx context.Context, volumeNameInternal, volumeNameExternal, comment string) error
	VolumeSetQosPolicyGroupName(ctx context.Context, name string, qos QosPolicyGroup) error
	VolumeSetSize(ctx context.Context, name, newSize string) error
	VolumeSize(ctx context.Context, volumeName string) (uint64, error)
	VolumeUsedSize(ctx context.Context, volumeName string) (int, error)
	VolumeSnapshotCreate(ctx context.Context, snapshotName, sourceVolume string) error
	VolumeSnapshotList(ctx context.Context, sourceVolume string) (Snapshots, error)
	VolumeSnapshotDelete(ctx context.Context, snapshotName, sourceVolume string) error

	TieringPolicyValue(ctx context.Context) string
}

type AggregateSpace interface {
	Size() int64
	Used() int64
	Footprint() int64
}

type SVMAggregateSpace struct {
	size      int64
	used      int64
	footprint int64
}

func (o SVMAggregateSpace) Size() int64 {
	return o.size
}

func (o SVMAggregateSpace) Used() int64 {
	return o.used
}

func (o SVMAggregateSpace) Footprint() int64 {
	return o.footprint
}

type Response interface {
	APIName() string
	Client() string
	Name() string
	Version() string
	Status() string
	Reason() string
	Errno() string
}

type APIResponse struct {
	apiName string
	status  string
	reason  string
	errno   string
}

// NewAPIResponse factory method to create a new instance of an APIResponse
func NewAPIResponse(status, reason, errno string) *APIResponse {
	result := &APIResponse{
		status: status,
		reason: reason,
		errno:  errno,
	}
	return result
}

func (o APIResponse) APIName() string {
	return o.apiName
}

func (o APIResponse) Status() string {
	return o.status
}

func (o APIResponse) Reason() string {
	return o.reason
}

func (o APIResponse) Errno() string {
	return o.errno
}

// GetErrorAbstraction inspects the supplied *APIResponse and error parameters to determine if an error occurred
func GetErrorAbstraction(ctx context.Context, response *APIResponse, errorIn error) (errorOut error) {
	defer func() {
		if r := recover(); r != nil {
			Logc(ctx).Errorf("Panic in ontap#GetErrorAbstraction. %v\nStack Trace: %v", response, string(debug.Stack()))
			errorOut = ZapiError{}
		}
	}()

	if errorIn != nil {
		errorOut = errorIn
		return errorOut
	}

	if response == nil {
		errorOut = fmt.Errorf("API error: nil response")
		return errorOut
	}

	responseStatus := response.Status()
	if strings.EqualFold(responseStatus, "passed") {
		errorOut = nil
		return errorOut
	}

	errorOut = fmt.Errorf("API error: %v", response)
	return errorOut
}
