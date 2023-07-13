// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: report_types.proto

package v1

import (
	fmt "fmt"
	protocol "github.com/gibson/gophercraft/bnet/bgs/protocol"
	v1 "github.com/gibson/gophercraft/bnet/bgs/protocol/account/v1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpamReport_SpamSource int32

const (
	SpamReport_OTHER             SpamReport_SpamSource = 1
	SpamReport_FRIEND_INVITATION SpamReport_SpamSource = 2
	SpamReport_WHISPER           SpamReport_SpamSource = 3
	SpamReport_CHAT              SpamReport_SpamSource = 4
)

var SpamReport_SpamSource_name = map[int32]string{
	1: "OTHER",
	2: "FRIEND_INVITATION",
	3: "WHISPER",
	4: "CHAT",
}

var SpamReport_SpamSource_value = map[string]int32{
	"OTHER":             1,
	"FRIEND_INVITATION": 2,
	"WHISPER":           3,
	"CHAT":              4,
}

func (x SpamReport_SpamSource) Enum() *SpamReport_SpamSource {
	p := new(SpamReport_SpamSource)
	*p = x
	return p
}

func (x SpamReport_SpamSource) String() string {
	return proto.EnumName(SpamReport_SpamSource_name, int32(x))
}

func (x *SpamReport_SpamSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SpamReport_SpamSource_value, data, "SpamReport_SpamSource")
	if err != nil {
		return err
	}
	*x = SpamReport_SpamSource(value)
	return nil
}

func (SpamReport_SpamSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{2, 0}
}

type ReportType struct {
	UserDescription *string `protobuf:"bytes,1,opt,name=user_description,json=userDescription" json:"user_description,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*ReportType_CustomReport
	//	*ReportType_SpamReport
	//	*ReportType_HarassmentReport
	//	*ReportType_RealLifeThreatReport
	//	*ReportType_InappropriateBattleTagReport
	//	*ReportType_HackingReport
	//	*ReportType_BottingReport
	Type                 isReportType_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReportType) Reset()         { *m = ReportType{} }
func (m *ReportType) String() string { return proto.CompactTextString(m) }
func (*ReportType) ProtoMessage()    {}
func (*ReportType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{0}
}

func (m *ReportType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportType.Unmarshal(m, b)
}
func (m *ReportType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportType.Marshal(b, m, deterministic)
}
func (m *ReportType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportType.Merge(m, src)
}
func (m *ReportType) XXX_Size() int {
	return xxx_messageInfo_ReportType.Size(m)
}
func (m *ReportType) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportType.DiscardUnknown(m)
}

var xxx_messageInfo_ReportType proto.InternalMessageInfo

func (m *ReportType) GetUserDescription() string {
	if m != nil && m.UserDescription != nil {
		return *m.UserDescription
	}
	return ""
}

type isReportType_Type interface {
	isReportType_Type()
}

type ReportType_CustomReport struct {
	CustomReport *CustomReport `protobuf:"bytes,10,opt,name=custom_report,json=customReport,oneof"`
}

type ReportType_SpamReport struct {
	SpamReport *SpamReport `protobuf:"bytes,11,opt,name=spam_report,json=spamReport,oneof"`
}

type ReportType_HarassmentReport struct {
	HarassmentReport *HarassmentReport `protobuf:"bytes,12,opt,name=harassment_report,json=harassmentReport,oneof"`
}

type ReportType_RealLifeThreatReport struct {
	RealLifeThreatReport *RealLifeThreatReport `protobuf:"bytes,13,opt,name=real_life_threat_report,json=realLifeThreatReport,oneof"`
}

type ReportType_InappropriateBattleTagReport struct {
	InappropriateBattleTagReport *InappropriateBattleTagReport `protobuf:"bytes,14,opt,name=inappropriate_battle_tag_report,json=inappropriateBattleTagReport,oneof"`
}

type ReportType_HackingReport struct {
	HackingReport *HackingReport `protobuf:"bytes,15,opt,name=hacking_report,json=hackingReport,oneof"`
}

type ReportType_BottingReport struct {
	BottingReport *BottingReport `protobuf:"bytes,16,opt,name=botting_report,json=bottingReport,oneof"`
}

func (*ReportType_CustomReport) isReportType_Type() {}

func (*ReportType_SpamReport) isReportType_Type() {}

func (*ReportType_HarassmentReport) isReportType_Type() {}

func (*ReportType_RealLifeThreatReport) isReportType_Type() {}

func (*ReportType_InappropriateBattleTagReport) isReportType_Type() {}

func (*ReportType_HackingReport) isReportType_Type() {}

func (*ReportType_BottingReport) isReportType_Type() {}

func (m *ReportType) GetType() isReportType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ReportType) GetCustomReport() *CustomReport {
	if x, ok := m.GetType().(*ReportType_CustomReport); ok {
		return x.CustomReport
	}
	return nil
}

func (m *ReportType) GetSpamReport() *SpamReport {
	if x, ok := m.GetType().(*ReportType_SpamReport); ok {
		return x.SpamReport
	}
	return nil
}

func (m *ReportType) GetHarassmentReport() *HarassmentReport {
	if x, ok := m.GetType().(*ReportType_HarassmentReport); ok {
		return x.HarassmentReport
	}
	return nil
}

func (m *ReportType) GetRealLifeThreatReport() *RealLifeThreatReport {
	if x, ok := m.GetType().(*ReportType_RealLifeThreatReport); ok {
		return x.RealLifeThreatReport
	}
	return nil
}

func (m *ReportType) GetInappropriateBattleTagReport() *InappropriateBattleTagReport {
	if x, ok := m.GetType().(*ReportType_InappropriateBattleTagReport); ok {
		return x.InappropriateBattleTagReport
	}
	return nil
}

func (m *ReportType) GetHackingReport() *HackingReport {
	if x, ok := m.GetType().(*ReportType_HackingReport); ok {
		return x.HackingReport
	}
	return nil
}

func (m *ReportType) GetBottingReport() *BottingReport {
	if x, ok := m.GetType().(*ReportType_BottingReport); ok {
		return x.BottingReport
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReportType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReportType_CustomReport)(nil),
		(*ReportType_SpamReport)(nil),
		(*ReportType_HarassmentReport)(nil),
		(*ReportType_RealLifeThreatReport)(nil),
		(*ReportType_InappropriateBattleTagReport)(nil),
		(*ReportType_HackingReport)(nil),
		(*ReportType_BottingReport)(nil),
	}
}

type CustomReport struct {
	Type                 *string               `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	ProgramId            *string               `protobuf:"bytes,2,opt,name=program_id,json=programId" json:"program_id,omitempty"` // Deprecated: Do not use.
	Attribute            []*protocol.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomReport) Reset()         { *m = CustomReport{} }
func (m *CustomReport) String() string { return proto.CompactTextString(m) }
func (*CustomReport) ProtoMessage()    {}
func (*CustomReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{1}
}

func (m *CustomReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomReport.Unmarshal(m, b)
}
func (m *CustomReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomReport.Marshal(b, m, deterministic)
}
func (m *CustomReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomReport.Merge(m, src)
}
func (m *CustomReport) XXX_Size() int {
	return xxx_messageInfo_CustomReport.Size(m)
}
func (m *CustomReport) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomReport.DiscardUnknown(m)
}

var xxx_messageInfo_CustomReport proto.InternalMessageInfo

func (m *CustomReport) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

// Deprecated: Do not use.
func (m *CustomReport) GetProgramId() string {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return ""
}

func (m *CustomReport) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type SpamReport struct {
	Target               *v1.GameAccountHandle  `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Source               *SpamReport_SpamSource `protobuf:"varint,2,opt,name=source,enum=bgs.protocol.report.v1.SpamReport_SpamSource,def=1" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SpamReport) Reset()         { *m = SpamReport{} }
func (m *SpamReport) String() string { return proto.CompactTextString(m) }
func (*SpamReport) ProtoMessage()    {}
func (*SpamReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{2}
}

func (m *SpamReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpamReport.Unmarshal(m, b)
}
func (m *SpamReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpamReport.Marshal(b, m, deterministic)
}
func (m *SpamReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpamReport.Merge(m, src)
}
func (m *SpamReport) XXX_Size() int {
	return xxx_messageInfo_SpamReport.Size(m)
}
func (m *SpamReport) XXX_DiscardUnknown() {
	xxx_messageInfo_SpamReport.DiscardUnknown(m)
}

var xxx_messageInfo_SpamReport proto.InternalMessageInfo

const Default_SpamReport_Source SpamReport_SpamSource = SpamReport_OTHER

func (m *SpamReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *SpamReport) GetSource() SpamReport_SpamSource {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return Default_SpamReport_Source
}

type HarassmentReport struct {
	Target               *v1.GameAccountHandle `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Text                 *string               `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HarassmentReport) Reset()         { *m = HarassmentReport{} }
func (m *HarassmentReport) String() string { return proto.CompactTextString(m) }
func (*HarassmentReport) ProtoMessage()    {}
func (*HarassmentReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{3}
}

func (m *HarassmentReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HarassmentReport.Unmarshal(m, b)
}
func (m *HarassmentReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HarassmentReport.Marshal(b, m, deterministic)
}
func (m *HarassmentReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HarassmentReport.Merge(m, src)
}
func (m *HarassmentReport) XXX_Size() int {
	return xxx_messageInfo_HarassmentReport.Size(m)
}
func (m *HarassmentReport) XXX_DiscardUnknown() {
	xxx_messageInfo_HarassmentReport.DiscardUnknown(m)
}

var xxx_messageInfo_HarassmentReport proto.InternalMessageInfo

func (m *HarassmentReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *HarassmentReport) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type RealLifeThreatReport struct {
	Target               *v1.GameAccountHandle `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Text                 *string               `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RealLifeThreatReport) Reset()         { *m = RealLifeThreatReport{} }
func (m *RealLifeThreatReport) String() string { return proto.CompactTextString(m) }
func (*RealLifeThreatReport) ProtoMessage()    {}
func (*RealLifeThreatReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{4}
}

func (m *RealLifeThreatReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealLifeThreatReport.Unmarshal(m, b)
}
func (m *RealLifeThreatReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealLifeThreatReport.Marshal(b, m, deterministic)
}
func (m *RealLifeThreatReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealLifeThreatReport.Merge(m, src)
}
func (m *RealLifeThreatReport) XXX_Size() int {
	return xxx_messageInfo_RealLifeThreatReport.Size(m)
}
func (m *RealLifeThreatReport) XXX_DiscardUnknown() {
	xxx_messageInfo_RealLifeThreatReport.DiscardUnknown(m)
}

var xxx_messageInfo_RealLifeThreatReport proto.InternalMessageInfo

func (m *RealLifeThreatReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RealLifeThreatReport) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type InappropriateBattleTagReport struct {
	Target               *v1.GameAccountHandle `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	BattleTag            *string               `protobuf:"bytes,2,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InappropriateBattleTagReport) Reset()         { *m = InappropriateBattleTagReport{} }
func (m *InappropriateBattleTagReport) String() string { return proto.CompactTextString(m) }
func (*InappropriateBattleTagReport) ProtoMessage()    {}
func (*InappropriateBattleTagReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{5}
}

func (m *InappropriateBattleTagReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InappropriateBattleTagReport.Unmarshal(m, b)
}
func (m *InappropriateBattleTagReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InappropriateBattleTagReport.Marshal(b, m, deterministic)
}
func (m *InappropriateBattleTagReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InappropriateBattleTagReport.Merge(m, src)
}
func (m *InappropriateBattleTagReport) XXX_Size() int {
	return xxx_messageInfo_InappropriateBattleTagReport.Size(m)
}
func (m *InappropriateBattleTagReport) XXX_DiscardUnknown() {
	xxx_messageInfo_InappropriateBattleTagReport.DiscardUnknown(m)
}

var xxx_messageInfo_InappropriateBattleTagReport proto.InternalMessageInfo

func (m *InappropriateBattleTagReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *InappropriateBattleTagReport) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

type HackingReport struct {
	Target               *v1.GameAccountHandle `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HackingReport) Reset()         { *m = HackingReport{} }
func (m *HackingReport) String() string { return proto.CompactTextString(m) }
func (*HackingReport) ProtoMessage()    {}
func (*HackingReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{6}
}

func (m *HackingReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HackingReport.Unmarshal(m, b)
}
func (m *HackingReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HackingReport.Marshal(b, m, deterministic)
}
func (m *HackingReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HackingReport.Merge(m, src)
}
func (m *HackingReport) XXX_Size() int {
	return xxx_messageInfo_HackingReport.Size(m)
}
func (m *HackingReport) XXX_DiscardUnknown() {
	xxx_messageInfo_HackingReport.DiscardUnknown(m)
}

var xxx_messageInfo_HackingReport proto.InternalMessageInfo

func (m *HackingReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

type BottingReport struct {
	Target               *v1.GameAccountHandle `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BottingReport) Reset()         { *m = BottingReport{} }
func (m *BottingReport) String() string { return proto.CompactTextString(m) }
func (*BottingReport) ProtoMessage()    {}
func (*BottingReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{7}
}

func (m *BottingReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BottingReport.Unmarshal(m, b)
}
func (m *BottingReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BottingReport.Marshal(b, m, deterministic)
}
func (m *BottingReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BottingReport.Merge(m, src)
}
func (m *BottingReport) XXX_Size() int {
	return xxx_messageInfo_BottingReport.Size(m)
}
func (m *BottingReport) XXX_DiscardUnknown() {
	xxx_messageInfo_BottingReport.DiscardUnknown(m)
}

var xxx_messageInfo_BottingReport proto.InternalMessageInfo

func (m *BottingReport) GetTarget() *v1.GameAccountHandle {
	if m != nil {
		return m.Target
	}
	return nil
}

type Report struct {
	ReportType           *string               `protobuf:"bytes,1,req,name=report_type,json=reportType" json:"report_type,omitempty"`
	Attribute            []*protocol.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	ReportQos            *int32                `protobuf:"varint,3,opt,name=report_qos,json=reportQos,def=0" json:"report_qos,omitempty"`
	ReportingAccount     *protocol.EntityId    `protobuf:"bytes,4,opt,name=reporting_account,json=reportingAccount" json:"reporting_account,omitempty"`
	ReportingGameAccount *protocol.EntityId    `protobuf:"bytes,5,opt,name=reporting_game_account,json=reportingGameAccount" json:"reporting_game_account,omitempty"`
	ReportTimestamp      *uint64               `protobuf:"fixed64,6,opt,name=report_timestamp,json=reportTimestamp" json:"report_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b4aafac95560569, []int{8}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

const Default_Report_ReportQos int32 = 0

func (m *Report) GetReportType() string {
	if m != nil && m.ReportType != nil {
		return *m.ReportType
	}
	return ""
}

func (m *Report) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Report) GetReportQos() int32 {
	if m != nil && m.ReportQos != nil {
		return *m.ReportQos
	}
	return Default_Report_ReportQos
}

func (m *Report) GetReportingAccount() *protocol.EntityId {
	if m != nil {
		return m.ReportingAccount
	}
	return nil
}

func (m *Report) GetReportingGameAccount() *protocol.EntityId {
	if m != nil {
		return m.ReportingGameAccount
	}
	return nil
}

func (m *Report) GetReportTimestamp() uint64 {
	if m != nil && m.ReportTimestamp != nil {
		return *m.ReportTimestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("bgs.protocol.report.v1.SpamReport_SpamSource", SpamReport_SpamSource_name, SpamReport_SpamSource_value)
	proto.RegisterType((*ReportType)(nil), "bgs.protocol.report.v1.ReportType")
	proto.RegisterType((*CustomReport)(nil), "bgs.protocol.report.v1.CustomReport")
	proto.RegisterType((*SpamReport)(nil), "bgs.protocol.report.v1.SpamReport")
	proto.RegisterType((*HarassmentReport)(nil), "bgs.protocol.report.v1.HarassmentReport")
	proto.RegisterType((*RealLifeThreatReport)(nil), "bgs.protocol.report.v1.RealLifeThreatReport")
	proto.RegisterType((*InappropriateBattleTagReport)(nil), "bgs.protocol.report.v1.InappropriateBattleTagReport")
	proto.RegisterType((*HackingReport)(nil), "bgs.protocol.report.v1.HackingReport")
	proto.RegisterType((*BottingReport)(nil), "bgs.protocol.report.v1.BottingReport")
	proto.RegisterType((*Report)(nil), "bgs.protocol.report.v1.Report")
}

func init() { proto.RegisterFile("report_types.proto", fileDescriptor_3b4aafac95560569) }

var fileDescriptor_3b4aafac95560569 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x5d, 0x6f, 0xf3, 0x34,
	0x14, 0xc7, 0x97, 0xbe, 0x41, 0x4f, 0xd7, 0x35, 0x33, 0x7b, 0x89, 0xa6, 0xa1, 0x95, 0x08, 0xa4,
	0x82, 0xa0, 0xdd, 0x2a, 0xb8, 0x99, 0xb8, 0x69, 0xb7, 0x42, 0x23, 0xa6, 0x0e, 0xdc, 0x8a, 0x49,
	0xdc, 0x44, 0x4e, 0xea, 0xa5, 0x81, 0x26, 0x31, 0x8e, 0x3b, 0x31, 0x09, 0x24, 0xf8, 0x4a, 0x7c,
	0x02, 0xbe, 0x0c, 0xdf, 0xe3, 0x51, 0x9c, 0x34, 0x2f, 0xd3, 0xda, 0xe7, 0xb9, 0xe8, 0x73, 0x67,
	0xff, 0x7d, 0xce, 0xef, 0x7f, 0xec, 0xf8, 0xc4, 0x80, 0x38, 0x65, 0x01, 0x17, 0xa6, 0x78, 0x66,
	0x34, 0xec, 0x32, 0x1e, 0x88, 0x00, 0x9d, 0x58, 0x4e, 0x32, 0xb4, 0x83, 0x65, 0x37, 0x0e, 0xe8,
	0x3e, 0x5d, 0x9d, 0x7d, 0x44, 0x6c, 0x3b, 0x58, 0xf9, 0x85, 0xe0, 0xb3, 0x63, 0x22, 0x04, 0x77,
	0xad, 0x95, 0xa0, 0x05, 0x19, 0x51, 0x5f, 0xb8, 0xe2, 0xb9, 0xa0, 0xb5, 0x38, 0xb3, 0xf3, 0x82,
	0xfe, 0x6f, 0x15, 0x00, 0x4b, 0xfc, 0xec, 0x99, 0x51, 0xf4, 0x39, 0xa8, 0xab, 0x90, 0x72, 0x73,
	0x4e, 0x43, 0x9b, 0xbb, 0x4c, 0xb8, 0x81, 0xaf, 0x29, 0x6d, 0xa5, 0x53, 0xc7, 0xad, 0x48, 0xbf,
	0xcd, 0x64, 0xf4, 0x03, 0x34, 0xed, 0x55, 0x28, 0x02, 0xcf, 0x8c, 0xcb, 0xd3, 0xa0, 0xad, 0x74,
	0x1a, 0xfd, 0x4f, 0xbb, 0xaf, 0x97, 0xde, 0xbd, 0x91, 0xc1, 0xb1, 0xd7, 0x78, 0x0f, 0xef, 0xdb,
	0xb9, 0x39, 0x1a, 0x41, 0x23, 0x64, 0x24, 0x45, 0x35, 0x24, 0x4a, 0xdf, 0x84, 0x9a, 0x32, 0x92,
	0x81, 0x20, 0x4c, 0x67, 0xe8, 0x01, 0x0e, 0x17, 0x84, 0x93, 0x30, 0xf4, 0xa8, 0x2f, 0xd6, 0xb0,
	0x7d, 0x09, 0xeb, 0x6c, 0x82, 0x8d, 0xd3, 0x84, 0x14, 0xa9, 0x2e, 0x5e, 0x68, 0x88, 0xc2, 0x29,
	0xa7, 0x64, 0x69, 0x2e, 0xdd, 0x47, 0x6a, 0x8a, 0x05, 0xa7, 0x24, 0xc5, 0x37, 0x25, 0xfe, 0xcb,
	0x4d, 0x78, 0x4c, 0xc9, 0xf2, 0xce, 0x7d, 0xa4, 0x33, 0x99, 0x94, 0x5a, 0x1c, 0xf1, 0x57, 0x74,
	0xf4, 0x17, 0x5c, 0xb8, 0x3e, 0x61, 0x8c, 0x07, 0x8c, 0xbb, 0x44, 0x50, 0xd3, 0x22, 0x42, 0x2c,
	0xa9, 0x29, 0x88, 0xb3, 0xb6, 0x3b, 0x90, 0x76, 0x5f, 0x6f, 0xb2, 0x33, 0xf2, 0xe9, 0x43, 0x99,
	0x3d, 0x23, 0x4e, 0x6a, 0x7b, 0xee, 0x6e, 0x59, 0x47, 0x13, 0x38, 0x58, 0x10, 0xfb, 0x37, 0xd7,
	0x4f, 0xdd, 0x5a, 0xd2, 0xed, 0xb3, 0xcd, 0x67, 0x27, 0xa3, 0x53, 0x7c, 0x73, 0x91, 0x17, 0x22,
	0x9e, 0x15, 0x08, 0x91, 0xe3, 0xa9, 0xdb, 0x79, 0xc3, 0x38, 0x3a, 0xe3, 0x59, 0x79, 0x61, 0x58,
	0x83, 0x4a, 0x74, 0x77, 0xf5, 0x3f, 0x61, 0x3f, 0x7f, 0x9b, 0x10, 0x8a, 0xf5, 0xe4, 0xa6, 0xca,
	0x31, 0xfa, 0x04, 0x80, 0xf1, 0xc0, 0xe1, 0xc4, 0x33, 0xdd, 0xb9, 0x56, 0x8a, 0x56, 0x86, 0x25,
	0x4d, 0xc1, 0xf5, 0x44, 0x35, 0xe6, 0xe8, 0x1b, 0xa8, 0xa7, 0x9d, 0xa3, 0x95, 0xdb, 0xe5, 0x4e,
	0xa3, 0x7f, 0x5a, 0xac, 0x6c, 0xb0, 0x5e, 0xc6, 0x59, 0xa4, 0xfe, 0xbf, 0x02, 0x90, 0xdd, 0x40,
	0x34, 0x84, 0x9a, 0x20, 0xdc, 0xa1, 0x42, 0xda, 0x37, 0xfa, 0x5f, 0x14, 0x11, 0x49, 0xc3, 0x46,
	0xbb, 0xfb, 0x9e, 0x78, 0x74, 0x10, 0x4f, 0xc7, 0xc4, 0x9f, 0x2f, 0x29, 0x4e, 0x32, 0xd1, 0x1d,
	0xd4, 0xc2, 0x60, 0xc5, 0x6d, 0x2a, 0x0b, 0x3d, 0xe8, 0x7f, 0xf5, 0xf6, 0x9b, 0x2f, 0x87, 0x53,
	0x99, 0x74, 0x5d, 0xbd, 0x9f, 0x8d, 0x47, 0x18, 0x27, 0x0c, 0x7d, 0x14, 0xd7, 0x17, 0x2f, 0xa2,
	0x3a, 0xc4, 0xcb, 0xaa, 0x82, 0x8e, 0xe1, 0xf0, 0x3b, 0x6c, 0x8c, 0x26, 0xb7, 0xa6, 0x31, 0xf9,
	0xd9, 0x98, 0x0d, 0x66, 0xc6, 0xfd, 0x44, 0x2d, 0xa1, 0x06, 0x7c, 0xf0, 0x30, 0x36, 0xa6, 0x3f,
	0x8e, 0xb0, 0x5a, 0x46, 0x1f, 0x42, 0xe5, 0x66, 0x3c, 0x98, 0xa9, 0x15, 0xfd, 0x57, 0x50, 0x5f,
	0xf6, 0xc6, 0x4e, 0x36, 0x1b, 0x7d, 0x2d, 0xfa, 0x87, 0x88, 0xbf, 0x09, 0x96, 0x63, 0xdd, 0x87,
	0xa3, 0xd7, 0x1a, 0xe5, 0xbd, 0xf9, 0xfd, 0xa3, 0xc0, 0xf9, 0xb6, 0x56, 0xd9, 0x89, 0xf1, 0xc7,
	0x00, 0x59, 0xff, 0x26, 0xf6, 0x75, 0x6b, 0x6d, 0xa4, 0x4f, 0xa1, 0x59, 0xe8, 0x9f, 0x5d, 0x78,
	0x46, 0xd0, 0x42, 0x13, 0xed, 0x04, 0xfa, 0x5f, 0x09, 0x6a, 0x09, 0xee, 0x02, 0x1a, 0xb9, 0xe7,
	0x4a, 0x53, 0xda, 0xa5, 0x4e, 0x1d, 0x03, 0xcf, 0x5e, 0x90, 0x42, 0x53, 0x95, 0xde, 0xb5, 0xa9,
	0x50, 0x1b, 0x12, 0x88, 0xf9, 0x7b, 0x10, 0x6a, 0xe5, 0xb6, 0xd2, 0xa9, 0x5e, 0x2b, 0x97, 0xb8,
	0x1e, 0x8b, 0x3f, 0x05, 0x21, 0xba, 0x81, 0xc3, 0x78, 0x12, 0xfd, 0x4e, 0x92, 0xb2, 0xb5, 0x8a,
	0xdc, 0xd3, 0x49, 0xd1, 0x60, 0x24, 0xdf, 0x3d, 0x63, 0x8e, 0xd5, 0x34, 0x21, 0xd9, 0x17, 0xba,
	0x83, 0x93, 0x0c, 0xe2, 0x10, 0x8f, 0xa6, 0xa4, 0xea, 0x56, 0xd2, 0x51, 0x9a, 0x95, 0x3b, 0xa5,
	0xe8, 0xb5, 0x5c, 0x1f, 0x86, 0xeb, 0xd1, 0x50, 0x10, 0x8f, 0x69, 0xb5, 0xb6, 0xd2, 0xa9, 0xe1,
	0x56, 0x72, 0x22, 0x6b, 0x79, 0x38, 0xf8, 0xe5, 0x5b, 0xc7, 0x15, 0x8b, 0x95, 0xd5, 0xb5, 0x03,
	0xaf, 0x17, 0xae, 0x18, 0xe5, 0xec, 0xf2, 0x52, 0xf4, 0x9c, 0x80, 0x2d, 0x28, 0xb7, 0x39, 0x79,
	0x14, 0x3d, 0xcb, 0xa7, 0xa2, 0x67, 0x39, 0x61, 0x6f, 0xed, 0xdf, 0x8b, 0x29, 0xbd, 0xa7, 0xab,
	0xbf, 0x95, 0xbd, 0x37, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x63, 0x16, 0x84, 0x28, 0x08, 0x00,
	0x00,
}
