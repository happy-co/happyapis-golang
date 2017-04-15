// Code generated by protoc-gen-go.
// source: happyco/inspect/report/v1/report.proto
// DO NOT EDIT!

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	happyco/inspect/report/v1/report.proto

It has these top-level messages:
	Report
	Signatory
	ListReportsRequest
	ListReportsResponse
	ReportsCreatedEventRequest
	ReportsCreatedEvent
	ReportsCreatedEventAck
	ReportsCreatedEventAckResponse
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import happyco_type_v1 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import happyco_type_v11 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import happyco_type_v12 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import happyco_type_v13 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import _ "github.com/happy-co/happyapis-golang/google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// * Reports are the end artifact of one or sometimes two inspections. They are
// generated from inspection(s) at a point in time but do not update if the
// underlying inspection updates. They are self contained once generated. They
// are intended to be similar to a printed pieice of paper in that they can't
// be edited but they can be signed at any time.
type Report struct {
	// * ID of report.
	Id *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// * Public URL can be used to share the inspection. It contains a UUID
	// to provide security.
	PublicUrl string `protobuf:"bytes,2,opt,name=public_url,json=publicUrl" json:"public_url,omitempty"`
	// * Name of the asset that was inspected. This is typically the address
	// of the unit or house that was inspected.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// * Title of the report. This can be overriden by the user in the
	// report settings but typically comes from the template name used to
	// perform the inspection.
	Title string `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	// * The revision of the report. This will increment when a change
	// occurs. Reports are limited to changes being archive/unarchive
	// operations or new signatures being added.
	Revision int32 `protobuf:"varint,5,opt,name=revision" json:"revision,omitempty"`
	// * Account ID that the report belongs to. For custom integrations
	// (i.e. not integration partners) it may be omitted.
	AccountId *happyco_type_v1.IntegrationID `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// * TODO: TBD if we should keep this -- inspection does not have it
	FolderId *happyco_type_v1.IntegrationID `protobuf:"bytes,7,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
	// * Asset ID references the parent that inspection is attached to. This field must
	// reference a non archived asset in the same folder as the template
	// referenced below.
	AssetId *happyco_type_v1.IntegrationID `protobuf:"bytes,8,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	// * Inspection IDs references one or more inspection(s) used to
	// generate the report.
	InspectionIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,9,rep,name=inspection_ids,json=inspectionIds" json:"inspection_ids,omitempty"`
	// * Generated at is the time (as reported by the mobile device) that the
	// report was generated at.
	GeneratedAt *happyco_type_v11.Date `protobuf:"bytes,10,opt,name=generated_at,json=generatedAt" json:"generated_at,omitempty"`
	// * Finalized is the time (as reported by the mobile device) that the
	// report was locked. Until the report is locked the display settings
	// can be changed. Signing a report locks it immediately.
	FinalizedAt *happyco_type_v11.Date `protobuf:"bytes,11,opt,name=finalized_at,json=finalizedAt" json:"finalized_at,omitempty"`
	// * TODO: Inspector ID references the user who created the report. TBD
	// whether to include this (where is it coming from)
	InspectorId *happyco_type_v1.IntegrationID `protobuf:"bytes,12,opt,name=inspector_id,json=inspectorId" json:"inspector_id,omitempty"`
	// * Inspector name taken from on the most recently complete inspection.
	InspectorName string `protobuf:"bytes,13,opt,name=inspector_name,json=inspectorName" json:"inspector_name,omitempty"`
	// * All of the report data (this structure is still subject to change
	// so at the moment this is just a json blob)
	Data string `protobuf:"bytes,14,opt,name=data" json:"data,omitempty"`
	// * A map of signatories added to the report. Each key in the map only
	// needs to be unique to this report.
	Signatories map[string]*Signatory `protobuf:"bytes,15,rep,name=signatories" json:"signatories,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Report) GetId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Report) GetPublicUrl() string {
	if m != nil {
		return m.PublicUrl
	}
	return ""
}

func (m *Report) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Report) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Report) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *Report) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *Report) GetFolderId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.FolderId
	}
	return nil
}

func (m *Report) GetAssetId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *Report) GetInspectionIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.InspectionIds
	}
	return nil
}

func (m *Report) GetGeneratedAt() *happyco_type_v11.Date {
	if m != nil {
		return m.GeneratedAt
	}
	return nil
}

func (m *Report) GetFinalizedAt() *happyco_type_v11.Date {
	if m != nil {
		return m.FinalizedAt
	}
	return nil
}

func (m *Report) GetInspectorId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.InspectorId
	}
	return nil
}

func (m *Report) GetInspectorName() string {
	if m != nil {
		return m.InspectorName
	}
	return ""
}

func (m *Report) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Report) GetSignatories() map[string]*Signatory {
	if m != nil {
		return m.Signatories
	}
	return nil
}

// * Signatory represents a signature box on a report
type Signatory struct {
	// * Name of the signatory. e.g. Mafalda Hopkirk.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// * TODO: should this be included? Its not included in the report.
	Role string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	// * Generated is the time when the signatory box was generated and
	// attached to the report.
	GeneratedAt *happyco_type_v11.Date `protobuf:"bytes,3,opt,name=generated_at,json=generatedAt" json:"generated_at,omitempty"`
	// * Signed at is the time (as reported by the mobile device) that the
	// user physically signed the report.
	SignedAt *happyco_type_v11.Date `protobuf:"bytes,4,opt,name=signed_at,json=signedAt" json:"signed_at,omitempty"`
	// * User signed at is a user adjustable signed at time (similar to the
	// date field often found on paper reports).
	UserSignedAt *happyco_type_v11.Date `protobuf:"bytes,5,opt,name=user_signed_at,json=userSignedAt" json:"user_signed_at,omitempty"`
}

func (m *Signatory) Reset()                    { *m = Signatory{} }
func (m *Signatory) String() string            { return proto.CompactTextString(m) }
func (*Signatory) ProtoMessage()               {}
func (*Signatory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Signatory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Signatory) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Signatory) GetGeneratedAt() *happyco_type_v11.Date {
	if m != nil {
		return m.GeneratedAt
	}
	return nil
}

func (m *Signatory) GetSignedAt() *happyco_type_v11.Date {
	if m != nil {
		return m.SignedAt
	}
	return nil
}

func (m *Signatory) GetUserSignedAt() *happyco_type_v11.Date {
	if m != nil {
		return m.UserSignedAt
	}
	return nil
}

// * List report request is used for fetching reports. It can be
// filtered to specific reports, assets or folders. If the intent is to be
// notified when an report undergoes a status transition please use the
// ReportEventService.
type ListReportsRequest struct {
	// * Account ID that the request applies to. For custom integrations
	// (i.e. not integration partners) it can be omitted.
	AccountId *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// * Folder IDs is an optional filter that lists reports only for
	// specific folders.
	FolderIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,2,rep,name=folder_ids,json=folderIds" json:"folder_ids,omitempty"`
	// * Asset IDs is an optional filter that lists reports only for
	// specific assets.
	AssetIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,3,rep,name=asset_ids,json=assetIds" json:"asset_ids,omitempty"`
	// * Inspections IDs is an optional filter that lists reports only for
	// specific inspections.
	InspectionIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,4,rep,name=inspection_ids,json=inspectionIds" json:"inspection_ids,omitempty"`
	// * Report IDs is an optional filter that lists reports only for
	// specific reports.
	ReportIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,5,rep,name=report_ids,json=reportIds" json:"report_ids,omitempty"`
	// * Archived when set to true returns archived reports only.
	Archived bool `protobuf:"varint,6,opt,name=archived" json:"archived,omitempty"`
	// * Paging message that contains information about the page being
	// requested. This message is usually copied from a previous list
	// request's response.
	Paging *happyco_type_v12.Paging `protobuf:"bytes,7,opt,name=paging" json:"paging,omitempty"`
}

func (m *ListReportsRequest) Reset()                    { *m = ListReportsRequest{} }
func (m *ListReportsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListReportsRequest) ProtoMessage()               {}
func (*ListReportsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListReportsRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *ListReportsRequest) GetFolderIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.FolderIds
	}
	return nil
}

func (m *ListReportsRequest) GetAssetIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AssetIds
	}
	return nil
}

func (m *ListReportsRequest) GetInspectionIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.InspectionIds
	}
	return nil
}

func (m *ListReportsRequest) GetReportIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.ReportIds
	}
	return nil
}

func (m *ListReportsRequest) GetArchived() bool {
	if m != nil {
		return m.Archived
	}
	return false
}

func (m *ListReportsRequest) GetPaging() *happyco_type_v12.Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

// * List reports response is a page of fetched report results. It
// includes a previous and next page token that can be passed into subsequent
// list report requests.
type ListReportsResponse struct {
	// * Reports contains the results for this page.
	Reports []*Report `protobuf:"bytes,1,rep,name=reports" json:"reports,omitempty"`
	// * Previous page is an token that will be present for all pages after
	// the first.
	PreviousPage *happyco_type_v12.Paging `protobuf:"bytes,2,opt,name=previous_page,json=previousPage" json:"previous_page,omitempty"`
	// * Next page is an token that will be present for all pages except the
	// last.
	NextPage *happyco_type_v12.Paging `protobuf:"bytes,3,opt,name=next_page,json=nextPage" json:"next_page,omitempty"`
}

func (m *ListReportsResponse) Reset()                    { *m = ListReportsResponse{} }
func (m *ListReportsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListReportsResponse) ProtoMessage()               {}
func (*ListReportsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListReportsResponse) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

func (m *ListReportsResponse) GetPreviousPage() *happyco_type_v12.Paging {
	if m != nil {
		return m.PreviousPage
	}
	return nil
}

func (m *ListReportsResponse) GetNextPage() *happyco_type_v12.Paging {
	if m != nil {
		return m.NextPage
	}
	return nil
}

type ReportsCreatedEventRequest struct {
	// * Account ID that the request applies to. For custom integrations
	// (i.e. not integration partners) it can be omitted.
	AccountId *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// * Options the client can set to configure the returned event stream.
	// See EventHandlerOptions for more details.
	Options *happyco_type_v13.EventHandlerOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *ReportsCreatedEventRequest) Reset()                    { *m = ReportsCreatedEventRequest{} }
func (m *ReportsCreatedEventRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportsCreatedEventRequest) ProtoMessage()               {}
func (*ReportsCreatedEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReportsCreatedEventRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *ReportsCreatedEventRequest) GetOptions() *happyco_type_v13.EventHandlerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// * Reports created events are emitted from streams requested with
// OnReportCreated.
type ReportsCreatedEvent struct {
	Event   *happyco_type_v13.Event              `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Reports []*ReportsCreatedEvent_ReportCreated `protobuf:"bytes,2,rep,name=reports" json:"reports,omitempty"`
}

func (m *ReportsCreatedEvent) Reset()                    { *m = ReportsCreatedEvent{} }
func (m *ReportsCreatedEvent) String() string            { return proto.CompactTextString(m) }
func (*ReportsCreatedEvent) ProtoMessage()               {}
func (*ReportsCreatedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReportsCreatedEvent) GetEvent() *happyco_type_v13.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ReportsCreatedEvent) GetReports() []*ReportsCreatedEvent_ReportCreated {
	if m != nil {
		return m.Reports
	}
	return nil
}

type ReportsCreatedEvent_ReportCreated struct {
	// * Report ID can be used to fetch the report with a filtered
	// list reports request
	ReportId *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=report_id,json=reportId" json:"report_id,omitempty"`
	// * TODO: should we remove this? I feel like is only useful if we have an updated event stream too
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *ReportsCreatedEvent_ReportCreated) Reset()         { *m = ReportsCreatedEvent_ReportCreated{} }
func (m *ReportsCreatedEvent_ReportCreated) String() string { return proto.CompactTextString(m) }
func (*ReportsCreatedEvent_ReportCreated) ProtoMessage()    {}
func (*ReportsCreatedEvent_ReportCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *ReportsCreatedEvent_ReportCreated) GetReportId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.ReportId
	}
	return nil
}

func (m *ReportsCreatedEvent_ReportCreated) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// * Reports created event ack is used to acknowledge or request time extensions
// for one or more events.
type ReportsCreatedEventAck struct {
	Ack *happyco_type_v13.EventAck `protobuf:"bytes,1,opt,name=ack" json:"ack,omitempty"`
}

func (m *ReportsCreatedEventAck) Reset()                    { *m = ReportsCreatedEventAck{} }
func (m *ReportsCreatedEventAck) String() string            { return proto.CompactTextString(m) }
func (*ReportsCreatedEventAck) ProtoMessage()               {}
func (*ReportsCreatedEventAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReportsCreatedEventAck) GetAck() *happyco_type_v13.EventAck {
	if m != nil {
		return m.Ack
	}
	return nil
}

// * ReportsCreatedEventAckResponse contains no data but indicates successful
// acknowledgement.
type ReportsCreatedEventAckResponse struct {
}

func (m *ReportsCreatedEventAckResponse) Reset()                    { *m = ReportsCreatedEventAckResponse{} }
func (m *ReportsCreatedEventAckResponse) String() string            { return proto.CompactTextString(m) }
func (*ReportsCreatedEventAckResponse) ProtoMessage()               {}
func (*ReportsCreatedEventAckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Report)(nil), "happyco.inspect.report.v1.Report")
	proto.RegisterType((*Signatory)(nil), "happyco.inspect.report.v1.Signatory")
	proto.RegisterType((*ListReportsRequest)(nil), "happyco.inspect.report.v1.ListReportsRequest")
	proto.RegisterType((*ListReportsResponse)(nil), "happyco.inspect.report.v1.ListReportsResponse")
	proto.RegisterType((*ReportsCreatedEventRequest)(nil), "happyco.inspect.report.v1.ReportsCreatedEventRequest")
	proto.RegisterType((*ReportsCreatedEvent)(nil), "happyco.inspect.report.v1.ReportsCreatedEvent")
	proto.RegisterType((*ReportsCreatedEvent_ReportCreated)(nil), "happyco.inspect.report.v1.ReportsCreatedEvent.ReportCreated")
	proto.RegisterType((*ReportsCreatedEventAck)(nil), "happyco.inspect.report.v1.ReportsCreatedEventAck")
	proto.RegisterType((*ReportsCreatedEventAckResponse)(nil), "happyco.inspect.report.v1.ReportsCreatedEventAckResponse")
}

func init() { proto.RegisterFile("happyco/inspect/report/v1/report.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0x97, 0xf3, 0x67, 0x63, 0xbf, 0x24, 0x65, 0x35, 0x0b, 0xc5, 0xeb, 0x96, 0x2a, 0xb5, 0x0a,
	0xac, 0xf8, 0xe3, 0xb0, 0x01, 0xa4, 0x76, 0xb7, 0x8b, 0x14, 0xe8, 0x4a, 0x44, 0x42, 0xb4, 0xf2,
	0x02, 0x07, 0x2e, 0xd1, 0xd4, 0x9e, 0x4d, 0x47, 0x31, 0xb6, 0x99, 0x99, 0x44, 0x04, 0x6e, 0xf4,
	0x23, 0x20, 0xc4, 0x81, 0x1b, 0x17, 0xbe, 0x09, 0x12, 0x67, 0xee, 0x9c, 0xf8, 0x20, 0x68, 0xfe,
	0xd8, 0xd9, 0x4d, 0xb2, 0xa9, 0x17, 0x71, 0x9b, 0x99, 0xf7, 0x7e, 0x6f, 0xde, 0x9b, 0xf7, 0xfb,
	0xcd, 0x0c, 0xbc, 0xf1, 0x0c, 0xe7, 0xf9, 0x22, 0xca, 0xfa, 0x34, 0xe5, 0x39, 0x89, 0x44, 0x9f,
	0x91, 0x3c, 0x63, 0xa2, 0x3f, 0x3f, 0x34, 0xa3, 0x20, 0x67, 0x99, 0xc8, 0xd0, 0xbe, 0xf1, 0x0b,
	0x8c, 0x5f, 0x60, 0xac, 0xf3, 0x43, 0xef, 0x5e, 0x11, 0x42, 0x2c, 0x72, 0x22, 0x81, 0x34, 0x15,
	0x64, 0xc2, 0xb0, 0xa0, 0x59, 0x3a, 0xa6, 0xb1, 0x0e, 0xe0, 0x79, 0xab, 0x5e, 0x31, 0x16, 0xc4,
	0xd8, 0x6e, 0xaf, 0xda, 0x72, 0x3c, 0xa1, 0xe9, 0xc4, 0x58, 0x6f, 0xad, 0x5a, 0xc9, 0x9c, 0xa4,
	0xa2, 0x80, 0x4e, 0xb2, 0x6c, 0x92, 0x90, 0x3e, 0xce, 0x69, 0x1f, 0xa7, 0x69, 0x26, 0xd4, 0xb6,
	0x5c, 0x5b, 0xfd, 0x3f, 0x77, 0x60, 0x27, 0x54, 0x89, 0xa2, 0x00, 0x6a, 0x34, 0x76, 0xad, 0x9e,
	0x75, 0xd0, 0x1e, 0xdc, 0x09, 0x8a, 0x6a, 0x64, 0xc8, 0x60, 0x7e, 0x18, 0x8c, 0x96, 0x29, 0x8f,
	0x1e, 0x85, 0x35, 0x1a, 0xa3, 0xd7, 0x00, 0xf2, 0xd9, 0xd3, 0x84, 0x46, 0xe3, 0x19, 0x4b, 0xdc,
	0x5a, 0xcf, 0x3a, 0x70, 0x42, 0x47, 0xaf, 0x7c, 0xc9, 0x12, 0x84, 0xa0, 0x91, 0xe2, 0x6f, 0x88,
	0x5b, 0x57, 0x06, 0x35, 0x46, 0x2f, 0x43, 0x53, 0x50, 0x91, 0x10, 0xb7, 0xa1, 0x16, 0xf5, 0x04,
	0x79, 0x60, 0x33, 0x32, 0xa7, 0x9c, 0x66, 0xa9, 0xdb, 0xec, 0x59, 0x07, 0xcd, 0xb0, 0x9c, 0xa3,
	0x13, 0x00, 0x1c, 0x45, 0xd9, 0x2c, 0x15, 0x63, 0x1a, 0xbb, 0x3b, 0x95, 0x92, 0x73, 0x0c, 0x62,
	0x14, 0xa3, 0x63, 0x70, 0xce, 0xb3, 0x24, 0x26, 0x4c, 0xa2, 0x5b, 0x95, 0xd0, 0xb6, 0x06, 0x8c,
	0x62, 0xf4, 0x00, 0x6c, 0xcc, 0x39, 0x51, 0x3b, 0xdb, 0x95, 0xb0, 0x2d, 0xe5, 0x3f, 0x8a, 0xd1,
	0x29, 0xdc, 0x30, 0x34, 0xd0, 0x2d, 0xe6, 0xae, 0xd3, 0xab, 0x57, 0x08, 0xd0, 0x5d, 0xa2, 0x46,
	0x31, 0x47, 0xf7, 0xa1, 0x33, 0x21, 0x29, 0x61, 0x58, 0x90, 0x78, 0x8c, 0x85, 0x0b, 0x2a, 0x8b,
	0x57, 0xd6, 0x82, 0x3c, 0xc2, 0x82, 0x84, 0xed, 0xd2, 0x75, 0x28, 0x24, 0xf2, 0x9c, 0xa6, 0x38,
	0xa1, 0xdf, 0x6b, 0x64, 0x7b, 0x2b, 0xb2, 0x74, 0x1d, 0x0a, 0x34, 0x84, 0x8e, 0x49, 0x22, 0x53,
	0xa7, 0xd6, 0xa9, 0x54, 0x79, 0xbb, 0xc4, 0x8c, 0x62, 0xf4, 0x7a, 0x59, 0x7d, 0xc6, 0xc6, 0x8a,
	0x04, 0x5d, 0xd5, 0xef, 0x6e, 0xb9, 0xfa, 0xb9, 0x64, 0x03, 0x82, 0x46, 0x8c, 0x05, 0x76, 0x6f,
	0x68, 0x86, 0xc8, 0x31, 0xfa, 0x02, 0xda, 0x9c, 0x4e, 0x52, 0x2c, 0x32, 0x46, 0x09, 0x77, 0x5f,
	0x52, 0xa7, 0x36, 0x08, 0xae, 0xd4, 0x56, 0xa0, 0xc9, 0x1b, 0x9c, 0x2d, 0x41, 0xa7, 0xa9, 0x60,
	0x8b, 0xf0, 0x62, 0x18, 0x2f, 0x86, 0xdd, 0x55, 0x07, 0xb4, 0x0b, 0xf5, 0x29, 0x59, 0x28, 0xbe,
	0x3b, 0xa1, 0x1c, 0xa2, 0x23, 0x68, 0xce, 0x71, 0x32, 0x23, 0x8a, 0xcb, 0xed, 0xc1, 0xbd, 0x2d,
	0xbb, 0x16, 0xd1, 0x16, 0xa1, 0x86, 0x1c, 0xd5, 0xee, 0x5b, 0xfe, 0xdf, 0x16, 0x38, 0xa5, 0xa1,
	0xe4, 0xbf, 0x75, 0x81, 0xff, 0x08, 0x1a, 0x2c, 0x4b, 0x88, 0x11, 0x8b, 0x1a, 0xaf, 0xf5, 0xb8,
	0x5e, 0xb9, 0xc7, 0x03, 0x70, 0x64, 0x91, 0x1a, 0xd6, 0xd8, 0x06, 0xb3, 0xb5, 0xdf, 0x50, 0xa0,
	0x63, 0xb8, 0x31, 0xe3, 0x84, 0x8d, 0x97, 0xc0, 0xe6, 0x36, 0x60, 0x47, 0x3a, 0x9f, 0x19, 0xb0,
	0xff, 0x5b, 0x1d, 0xd0, 0x67, 0x94, 0x0b, 0x7d, 0xe6, 0x3c, 0x24, 0xdf, 0xce, 0x08, 0x17, 0x2b,
	0x1a, 0xb5, 0xae, 0xab, 0xd1, 0x13, 0x80, 0x52, 0xa3, 0xdc, 0xad, 0x55, 0xd2, 0x89, 0x53, 0x88,
	0x94, 0x4b, 0x89, 0x17, 0x2a, 0xe5, 0x6e, 0xbd, 0x12, 0xda, 0x36, 0x32, 0xe5, 0x1b, 0x74, 0xda,
	0xf8, 0x2f, 0x3a, 0x3d, 0x01, 0xd0, 0xdc, 0x50, 0x21, 0x9a, 0xd5, 0x4a, 0xd0, 0x08, 0x09, 0xf7,
	0xc0, 0xc6, 0x2c, 0x7a, 0x46, 0xe7, 0x44, 0x5f, 0x71, 0x76, 0x58, 0xce, 0x51, 0x1f, 0x76, 0xf4,
	0x5d, 0x6f, 0xae, 0xaf, 0x57, 0xd7, 0xc2, 0x3e, 0x51, 0xe6, 0xd0, 0xb8, 0xf9, 0x7f, 0x58, 0xb0,
	0x77, 0xa9, 0x49, 0x3c, 0xcf, 0x52, 0x4e, 0xd0, 0x31, 0xb4, 0xf4, 0x8e, 0xdc, 0xb5, 0x54, 0x82,
	0x77, 0x5f, 0xa8, 0xaa, 0xb0, 0x40, 0xa0, 0x87, 0xd0, 0xcd, 0xe5, 0x9d, 0x9c, 0xcd, 0xf8, 0x38,
	0xc7, 0x93, 0x42, 0x22, 0x57, 0x26, 0xd3, 0x29, 0xbc, 0x9f, 0xe0, 0x09, 0x41, 0x1f, 0x80, 0x93,
	0x92, 0xef, 0x84, 0x46, 0xd6, 0xb7, 0x23, 0x6d, 0xe9, 0x29, 0x51, 0xfe, 0xaf, 0x16, 0x78, 0xa6,
	0x88, 0x4f, 0x18, 0x91, 0x9c, 0x3f, 0x95, 0xcf, 0xda, 0xff, 0xc4, 0xba, 0x8f, 0xa0, 0x95, 0xe5,
	0xea, 0x25, 0x5c, 0x93, 0x7b, 0x81, 0x55, 0xdb, 0x7d, 0x8a, 0xd3, 0x38, 0x21, 0xec, 0xb1, 0xf6,
	0x0d, 0x0b, 0x90, 0xff, 0xbc, 0x06, 0x7b, 0x1b, 0xb2, 0x43, 0xef, 0x40, 0x53, 0xbd, 0xbe, 0x26,
	0xa3, 0x9b, 0x9b, 0xa3, 0x86, 0xda, 0x09, 0x7d, 0xb5, 0x6c, 0x8a, 0x26, 0xfe, 0xc3, 0x17, 0x36,
	0xe5, 0xd2, 0x76, 0x66, 0xcd, 0x2c, 0x95, 0xfd, 0xf2, 0xce, 0xa1, 0x7b, 0xc9, 0x22, 0x55, 0x52,
	0x32, 0xb4, 0xe2, 0x61, 0xd9, 0x05, 0x41, 0x91, 0x0b, 0xad, 0x39, 0x61, 0xea, 0x7d, 0xae, 0xa9,
	0xf7, 0xb9, 0x98, 0xfa, 0xa7, 0x70, 0x73, 0x43, 0x56, 0xc3, 0x68, 0x8a, 0xde, 0x86, 0x3a, 0x8e,
	0xa6, 0x66, 0xab, 0xfd, 0xcd, 0xa7, 0x30, 0x8c, 0xa6, 0xa1, 0xf4, 0xf2, 0x7b, 0x70, 0x67, 0x73,
	0x98, 0x82, 0xbd, 0x83, 0xdf, 0xad, 0xa2, 0xa2, 0x33, 0xc2, 0xe6, 0x34, 0x22, 0xe8, 0x67, 0x0b,
	0xda, 0x17, 0x78, 0x8e, 0xde, 0xdd, 0x72, 0x72, 0xeb, 0x97, 0x96, 0x17, 0x54, 0x75, 0xd7, 0x09,
	0xf8, 0x6f, 0xfe, 0xf8, 0xd7, 0x3f, 0x3f, 0xd5, 0xee, 0xfa, 0xb7, 0xd5, 0x47, 0x4a, 0xfd, 0xe1,
	0x2e, 0x7e, 0x07, 0x79, 0x3f, 0xa1, 0x5c, 0x1c, 0x59, 0x6f, 0x0d, 0x7e, 0xa9, 0x01, 0xd2, 0x60,
	0x55, 0x44, 0x91, 0xee, 0x0f, 0xb0, 0xfb, 0x38, 0xbd, 0x5c, 0x24, 0xfa, 0xf0, 0x7a, 0xcd, 0xae,
	0x92, 0xfa, 0x06, 0xd8, 0x7b, 0x16, 0x7a, 0x6e, 0xc1, 0xde, 0xea, 0xee, 0xb2, 0x49, 0x87, 0xd7,
	0x8b, 0x34, 0x8c, 0xa6, 0xde, 0x83, 0x6b, 0x43, 0x8a, 0x23, 0xfc, 0xf8, 0xd6, 0xd7, 0xfb, 0x57,
	0xfe, 0xa5, 0x9f, 0xee, 0xa8, 0xff, 0xe8, 0xfb, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x74, 0xa9,
	0x1f, 0xfa, 0x6f, 0x0b, 0x00, 0x00,
}
