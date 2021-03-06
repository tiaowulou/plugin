// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unfreeze.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	unfreeze.proto

It has these top-level messages:
	Unfreeze
	FixAmount
	LeftProportion
	UnfreezeAction
	UnfreezeCreate
	UnfreezeWithdraw
	UnfreezeTerminate
	ReceiptUnfreeze
	ReplyQueryUnfreezeWithdraw
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types1 "github.com/33cn/chain33/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Unfreeze struct {
	// 解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	// 开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	// 币种
	AssetExec   string `protobuf:"bytes,3,opt,name=assetExec" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,4,opt,name=assetSymbol" json:"assetSymbol,omitempty"`
	// 冻结总额
	TotalCount int64 `protobuf:"varint,5,opt,name=totalCount" json:"totalCount,omitempty"`
	// 发币人地址
	Initiator string `protobuf:"bytes,6,opt,name=initiator" json:"initiator,omitempty"`
	// 收币人地址
	Beneficiary string `protobuf:"bytes,7,opt,name=beneficiary" json:"beneficiary,omitempty"`
	// 解冻剩余币数
	Remaining int64 `protobuf:"varint,8,opt,name=remaining" json:"remaining,omitempty"`
	// 解冻方式（百分比；固额）
	Means string `protobuf:"bytes,9,opt,name=means" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*Unfreeze_FixAmount
	//	*Unfreeze_LeftProportion
	MeansOpt isUnfreeze_MeansOpt `protobuf_oneof:"meansOpt"`
}

func (m *Unfreeze) Reset()                    { *m = Unfreeze{} }
func (m *Unfreeze) String() string            { return proto.CompactTextString(m) }
func (*Unfreeze) ProtoMessage()               {}
func (*Unfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isUnfreeze_MeansOpt interface {
	isUnfreeze_MeansOpt()
}

type Unfreeze_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,10,opt,name=fixAmount,oneof"`
}
type Unfreeze_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,11,opt,name=leftProportion,oneof"`
}

func (*Unfreeze_FixAmount) isUnfreeze_MeansOpt()      {}
func (*Unfreeze_LeftProportion) isUnfreeze_MeansOpt() {}

func (m *Unfreeze) GetMeansOpt() isUnfreeze_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *Unfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *Unfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Unfreeze) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *Unfreeze) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *Unfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Unfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Unfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *Unfreeze) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *Unfreeze) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

func (m *Unfreeze) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*Unfreeze_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *Unfreeze) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*Unfreeze_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Unfreeze) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Unfreeze_OneofMarshaler, _Unfreeze_OneofUnmarshaler, _Unfreeze_OneofSizer, []interface{}{
		(*Unfreeze_FixAmount)(nil),
		(*Unfreeze_LeftProportion)(nil),
	}
}

func _Unfreeze_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Unfreeze)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *Unfreeze_FixAmount:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixAmount); err != nil {
			return err
		}
	case *Unfreeze_LeftProportion:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LeftProportion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Unfreeze.MeansOpt has unexpected type %T", x)
	}
	return nil
}

func _Unfreeze_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Unfreeze)
	switch tag {
	case 10: // meansOpt.fixAmount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixAmount)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &Unfreeze_FixAmount{msg}
		return true, err
	case 11: // meansOpt.leftProportion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeftProportion)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &Unfreeze_LeftProportion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Unfreeze_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Unfreeze)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *Unfreeze_FixAmount:
		s := proto.Size(x.FixAmount)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Unfreeze_LeftProportion:
		s := proto.Size(x.LeftProportion)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 按时间固定额度解冻
type FixAmount struct {
	Period int64 `protobuf:"varint,1,opt,name=period" json:"period,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *FixAmount) Reset()                    { *m = FixAmount{} }
func (m *FixAmount) String() string            { return proto.CompactTextString(m) }
func (*FixAmount) ProtoMessage()               {}
func (*FixAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FixAmount) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *FixAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 固定时间间隔按余量百分比解冻
type LeftProportion struct {
	Period        int64 `protobuf:"varint,1,opt,name=period" json:"period,omitempty"`
	TenThousandth int64 `protobuf:"varint,2,opt,name=tenThousandth" json:"tenThousandth,omitempty"`
}

func (m *LeftProportion) Reset()                    { *m = LeftProportion{} }
func (m *LeftProportion) String() string            { return proto.CompactTextString(m) }
func (*LeftProportion) ProtoMessage()               {}
func (*LeftProportion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LeftProportion) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *LeftProportion) GetTenThousandth() int64 {
	if m != nil {
		return m.TenThousandth
	}
	return 0
}

// message for execs.unfreeze
type UnfreezeAction struct {
	// Types that are valid to be assigned to Value:
	//	*UnfreezeAction_Create
	//	*UnfreezeAction_Withdraw
	//	*UnfreezeAction_Terminate
	Value isUnfreezeAction_Value `protobuf_oneof:"value"`
	Ty    int32                  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *UnfreezeAction) Reset()                    { *m = UnfreezeAction{} }
func (m *UnfreezeAction) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeAction) ProtoMessage()               {}
func (*UnfreezeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isUnfreezeAction_Value interface {
	isUnfreezeAction_Value()
}

type UnfreezeAction_Create struct {
	Create *UnfreezeCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type UnfreezeAction_Withdraw struct {
	Withdraw *UnfreezeWithdraw `protobuf:"bytes,2,opt,name=withdraw,oneof"`
}
type UnfreezeAction_Terminate struct {
	Terminate *UnfreezeTerminate `protobuf:"bytes,3,opt,name=terminate,oneof"`
}

func (*UnfreezeAction_Create) isUnfreezeAction_Value()    {}
func (*UnfreezeAction_Withdraw) isUnfreezeAction_Value()  {}
func (*UnfreezeAction_Terminate) isUnfreezeAction_Value() {}

func (m *UnfreezeAction) GetValue() isUnfreezeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UnfreezeAction) GetCreate() *UnfreezeCreate {
	if x, ok := m.GetValue().(*UnfreezeAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UnfreezeAction) GetWithdraw() *UnfreezeWithdraw {
	if x, ok := m.GetValue().(*UnfreezeAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *UnfreezeAction) GetTerminate() *UnfreezeTerminate {
	if x, ok := m.GetValue().(*UnfreezeAction_Terminate); ok {
		return x.Terminate
	}
	return nil
}

func (m *UnfreezeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeAction_OneofMarshaler, _UnfreezeAction_OneofUnmarshaler, _UnfreezeAction_OneofSizer, []interface{}{
		(*UnfreezeAction_Create)(nil),
		(*UnfreezeAction_Withdraw)(nil),
		(*UnfreezeAction_Terminate)(nil),
	}
}

func _UnfreezeAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *UnfreezeAction_Withdraw:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *UnfreezeAction_Terminate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Terminate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeAction.Value has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeCreate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Create{msg}
		return true, err
	case 2: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Withdraw{msg}
		return true, err
	case 3: // value.terminate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnfreezeTerminate)
		err := b.DecodeMessage(msg)
		m.Value = &UnfreezeAction_Terminate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeAction)
	// value
	switch x := m.Value.(type) {
	case *UnfreezeAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeAction_Terminate:
		s := proto.Size(x.Terminate)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// action
type UnfreezeCreate struct {
	StartTime   int64  `protobuf:"varint,1,opt,name=startTime" json:"startTime,omitempty"`
	AssetExec   string `protobuf:"bytes,2,opt,name=assetExec" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,3,opt,name=assetSymbol" json:"assetSymbol,omitempty"`
	TotalCount  int64  `protobuf:"varint,4,opt,name=totalCount" json:"totalCount,omitempty"`
	Beneficiary string `protobuf:"bytes,5,opt,name=beneficiary" json:"beneficiary,omitempty"`
	Means       string `protobuf:"bytes,6,opt,name=means" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*UnfreezeCreate_FixAmount
	//	*UnfreezeCreate_LeftProportion
	MeansOpt isUnfreezeCreate_MeansOpt `protobuf_oneof:"meansOpt"`
}

func (m *UnfreezeCreate) Reset()                    { *m = UnfreezeCreate{} }
func (m *UnfreezeCreate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeCreate) ProtoMessage()               {}
func (*UnfreezeCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isUnfreezeCreate_MeansOpt interface {
	isUnfreezeCreate_MeansOpt()
}

type UnfreezeCreate_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,7,opt,name=fixAmount,oneof"`
}
type UnfreezeCreate_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,8,opt,name=leftProportion,oneof"`
}

func (*UnfreezeCreate_FixAmount) isUnfreezeCreate_MeansOpt()      {}
func (*UnfreezeCreate_LeftProportion) isUnfreezeCreate_MeansOpt() {}

func (m *UnfreezeCreate) GetMeansOpt() isUnfreezeCreate_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *UnfreezeCreate) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *UnfreezeCreate) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *UnfreezeCreate) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *UnfreezeCreate) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UnfreezeCreate) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *UnfreezeCreate) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

func (m *UnfreezeCreate) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *UnfreezeCreate) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UnfreezeCreate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UnfreezeCreate_OneofMarshaler, _UnfreezeCreate_OneofUnmarshaler, _UnfreezeCreate_OneofSizer, []interface{}{
		(*UnfreezeCreate_FixAmount)(nil),
		(*UnfreezeCreate_LeftProportion)(nil),
	}
}

func _UnfreezeCreate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UnfreezeCreate)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *UnfreezeCreate_FixAmount:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixAmount); err != nil {
			return err
		}
	case *UnfreezeCreate_LeftProportion:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LeftProportion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UnfreezeCreate.MeansOpt has unexpected type %T", x)
	}
	return nil
}

func _UnfreezeCreate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UnfreezeCreate)
	switch tag {
	case 7: // meansOpt.fixAmount
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixAmount)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &UnfreezeCreate_FixAmount{msg}
		return true, err
	case 8: // meansOpt.leftProportion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LeftProportion)
		err := b.DecodeMessage(msg)
		m.MeansOpt = &UnfreezeCreate_LeftProportion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UnfreezeCreate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UnfreezeCreate)
	// meansOpt
	switch x := m.MeansOpt.(type) {
	case *UnfreezeCreate_FixAmount:
		s := proto.Size(x.FixAmount)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UnfreezeCreate_LeftProportion:
		s := proto.Size(x.LeftProportion)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UnfreezeWithdraw struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeWithdraw) Reset()                    { *m = UnfreezeWithdraw{} }
func (m *UnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeWithdraw) ProtoMessage()               {}
func (*UnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type UnfreezeTerminate struct {
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
}

func (m *UnfreezeTerminate) Reset()                    { *m = UnfreezeTerminate{} }
func (m *UnfreezeTerminate) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeTerminate) ProtoMessage()               {}
func (*UnfreezeTerminate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnfreezeTerminate) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

// receipt
type ReceiptUnfreeze struct {
	Prev    *Unfreeze `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *Unfreeze `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptUnfreeze) Reset()                    { *m = ReceiptUnfreeze{} }
func (m *ReceiptUnfreeze) String() string            { return proto.CompactTextString(m) }
func (*ReceiptUnfreeze) ProtoMessage()               {}
func (*ReceiptUnfreeze) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReceiptUnfreeze) GetPrev() *Unfreeze {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptUnfreeze) GetCurrent() *Unfreeze {
	if m != nil {
		return m.Current
	}
	return nil
}

// query
type ReplyQueryUnfreezeWithdraw struct {
	UnfreezeID      string `protobuf:"bytes,1,opt,name=unfreezeID" json:"unfreezeID,omitempty"`
	AvailableAmount int64  `protobuf:"varint,2,opt,name=availableAmount" json:"availableAmount,omitempty"`
}

func (m *ReplyQueryUnfreezeWithdraw) Reset()                    { *m = ReplyQueryUnfreezeWithdraw{} }
func (m *ReplyQueryUnfreezeWithdraw) String() string            { return proto.CompactTextString(m) }
func (*ReplyQueryUnfreezeWithdraw) ProtoMessage()               {}
func (*ReplyQueryUnfreezeWithdraw) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReplyQueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReplyQueryUnfreezeWithdraw) GetAvailableAmount() int64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Unfreeze)(nil), "types.Unfreeze")
	proto.RegisterType((*FixAmount)(nil), "types.FixAmount")
	proto.RegisterType((*LeftProportion)(nil), "types.LeftProportion")
	proto.RegisterType((*UnfreezeAction)(nil), "types.UnfreezeAction")
	proto.RegisterType((*UnfreezeCreate)(nil), "types.UnfreezeCreate")
	proto.RegisterType((*UnfreezeWithdraw)(nil), "types.UnfreezeWithdraw")
	proto.RegisterType((*UnfreezeTerminate)(nil), "types.UnfreezeTerminate")
	proto.RegisterType((*ReceiptUnfreeze)(nil), "types.ReceiptUnfreeze")
	proto.RegisterType((*ReplyQueryUnfreezeWithdraw)(nil), "types.ReplyQueryUnfreezeWithdraw")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Unfreeze service

type UnfreezeClient interface {
	GetUnfreezeWithdraw(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*ReplyQueryUnfreezeWithdraw, error)
	QueryUnfreeze(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*Unfreeze, error)
}

type unfreezeClient struct {
	cc *grpc.ClientConn
}

func NewUnfreezeClient(cc *grpc.ClientConn) UnfreezeClient {
	return &unfreezeClient{cc}
}

func (c *unfreezeClient) GetUnfreezeWithdraw(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*ReplyQueryUnfreezeWithdraw, error) {
	out := new(ReplyQueryUnfreezeWithdraw)
	err := grpc.Invoke(ctx, "/types.unfreeze/GetUnfreezeWithdraw", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unfreezeClient) QueryUnfreeze(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*Unfreeze, error) {
	out := new(Unfreeze)
	err := grpc.Invoke(ctx, "/types.unfreeze/QueryUnfreeze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Unfreeze service

type UnfreezeServer interface {
	GetUnfreezeWithdraw(context.Context, *types1.ReqString) (*ReplyQueryUnfreezeWithdraw, error)
	QueryUnfreeze(context.Context, *types1.ReqString) (*Unfreeze, error)
}

func RegisterUnfreezeServer(s *grpc.Server, srv UnfreezeServer) {
	s.RegisterService(&_Unfreeze_serviceDesc, srv)
}

func _Unfreeze_GetUnfreezeWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnfreezeServer).GetUnfreezeWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.unfreeze/GetUnfreezeWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnfreezeServer).GetUnfreezeWithdraw(ctx, req.(*types1.ReqString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Unfreeze_QueryUnfreeze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnfreezeServer).QueryUnfreeze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.unfreeze/QueryUnfreeze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnfreezeServer).QueryUnfreeze(ctx, req.(*types1.ReqString))
	}
	return interceptor(ctx, in, info, handler)
}

var _Unfreeze_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.unfreeze",
	HandlerType: (*UnfreezeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnfreezeWithdraw",
			Handler:    _Unfreeze_GetUnfreezeWithdraw_Handler,
		},
		{
			MethodName: "QueryUnfreeze",
			Handler:    _Unfreeze_QueryUnfreeze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unfreeze.proto",
}

func init() { proto.RegisterFile("unfreeze.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x93, 0x26, 0xb1, 0x6f, 0x68, 0x5a, 0x86, 0x97, 0x55, 0x21, 0x14, 0x0c, 0x8b, 0xb0,
	0x29, 0x28, 0x05, 0x09, 0x89, 0x05, 0x6a, 0xcb, 0x23, 0x48, 0x15, 0x8f, 0x69, 0x11, 0xeb, 0x89,
	0x7b, 0xd3, 0x8e, 0x64, 0xcf, 0x98, 0xf1, 0xb8, 0xad, 0xf9, 0x08, 0x7e, 0x80, 0xef, 0x61, 0xc5,
	0x4f, 0x21, 0x8f, 0x1f, 0x49, 0x9c, 0x96, 0x88, 0x2e, 0xe7, 0x9c, 0x7b, 0xce, 0xbd, 0x93, 0x7b,
	0x3c, 0x81, 0x7e, 0x22, 0xa6, 0x0a, 0xf1, 0x07, 0x6e, 0x47, 0x4a, 0x6a, 0x49, 0xda, 0x3a, 0x8d,
	0x30, 0xde, 0xba, 0xe1, 0xcb, 0x30, 0x94, 0x22, 0x07, 0xbd, 0x5f, 0x2d, 0xb0, 0xbf, 0x16, 0x75,
	0xe4, 0x01, 0x40, 0xa9, 0xf9, 0xf0, 0xc6, 0xb5, 0x06, 0xd6, 0xd0, 0xa1, 0x73, 0x08, 0xb9, 0x0f,
	0x4e, 0xac, 0x99, 0xd2, 0x47, 0x3c, 0x44, 0xb7, 0x39, 0xb0, 0x86, 0x2d, 0x3a, 0x03, 0x32, 0x96,
	0xc5, 0x31, 0xea, 0xb7, 0x17, 0xe8, 0xbb, 0x2d, 0x23, 0x9e, 0x01, 0x64, 0x00, 0x3d, 0x73, 0x38,
	0x4c, 0xc3, 0x89, 0x0c, 0xdc, 0x35, 0xc3, 0xcf, 0x43, 0x59, 0x77, 0x2d, 0x35, 0x0b, 0xf6, 0x65,
	0x22, 0xb4, 0xdb, 0x36, 0xf6, 0x73, 0x48, 0xe6, 0xcf, 0x05, 0xd7, 0x9c, 0x69, 0xa9, 0xdc, 0x4e,
	0xee, 0x5f, 0x01, 0x99, 0xff, 0x04, 0x05, 0x4e, 0xb9, 0xcf, 0x99, 0x4a, 0xdd, 0x6e, 0xee, 0x3f,
	0x07, 0x65, 0x7a, 0x85, 0x21, 0xe3, 0x82, 0x8b, 0x13, 0xd7, 0xce, 0xa7, 0xaf, 0x00, 0x72, 0x1b,
	0xda, 0x21, 0x32, 0x11, 0xbb, 0x8e, 0x51, 0xe6, 0x07, 0xf2, 0x0c, 0x9c, 0x29, 0xbf, 0xd8, 0x0d,
	0xcd, 0x48, 0x30, 0xb0, 0x86, 0xbd, 0xd1, 0xe6, 0xb6, 0xf9, 0x1d, 0xb7, 0xdf, 0x95, 0xf8, 0xb8,
	0x41, 0x67, 0x45, 0xe4, 0x35, 0xf4, 0x03, 0x9c, 0xea, 0xcf, 0x4a, 0x46, 0x52, 0x69, 0x2e, 0x85,
	0xdb, 0x33, 0xb2, 0x3b, 0x85, 0xec, 0x60, 0x81, 0x1c, 0x37, 0x68, 0xad, 0x7c, 0x0f, 0xc0, 0x36,
	0xbd, 0x3f, 0x45, 0xda, 0x7b, 0x05, 0x4e, 0xd5, 0x86, 0xdc, 0x85, 0x4e, 0x84, 0x8a, 0xcb, 0x63,
	0xb3, 0x99, 0x16, 0x2d, 0x4e, 0x19, 0xce, 0xf2, 0x01, 0xf3, 0x95, 0x14, 0x27, 0xef, 0x23, 0xf4,
	0x17, 0x9b, 0x5d, 0xe9, 0xf0, 0x18, 0xd6, 0x35, 0x8a, 0xa3, 0x53, 0x99, 0xc4, 0x4c, 0x1c, 0xeb,
	0xd3, 0xc2, 0x68, 0x11, 0xf4, 0xfe, 0x58, 0xd0, 0x2f, 0xa3, 0xb2, 0xeb, 0x1b, 0xc3, 0xa7, 0xd0,
	0xf1, 0x15, 0x32, 0x8d, 0xc6, 0x70, 0x76, 0xc9, 0xb2, 0x6c, 0xdf, 0x90, 0xe3, 0x06, 0x2d, 0xca,
	0xc8, 0x0b, 0xb0, 0xcf, 0xb9, 0x3e, 0x3d, 0x56, 0xec, 0xdc, 0x34, 0xe9, 0x8d, 0xee, 0xd5, 0x24,
	0xdf, 0x0a, 0x7a, 0xdc, 0xa0, 0x55, 0x29, 0x79, 0x09, 0x8e, 0x46, 0x15, 0x72, 0x91, 0xb5, 0x6a,
	0x19, 0x9d, 0x5b, 0xd3, 0x1d, 0x95, 0x7c, 0xb6, 0x8e, 0xaa, 0x98, 0xf4, 0xa1, 0xa9, 0x53, 0x93,
	0xb6, 0x36, 0x6d, 0xea, 0x74, 0xaf, 0x0b, 0xed, 0x33, 0x16, 0x24, 0xe8, 0xfd, 0x6e, 0xce, 0x6e,
	0x93, 0x8f, 0xb9, 0x18, 0x6f, 0xeb, 0x9f, 0xf1, 0x6e, 0xae, 0x88, 0x77, 0x6b, 0x55, 0xbc, 0xd7,
	0x96, 0xe2, 0x5d, 0x0b, 0x70, 0x7b, 0x39, 0xc0, 0x55, 0x44, 0x3b, 0x57, 0x46, 0xb4, 0x7b, 0xbd,
	0x88, 0xda, 0xd7, 0x8f, 0xe8, 0x08, 0x36, 0xeb, 0xab, 0x5b, 0xf5, 0x8e, 0x78, 0x3b, 0x70, 0x73,
	0x69, 0x6d, 0x2b, 0x45, 0x0c, 0x36, 0x28, 0xfa, 0xc8, 0x23, 0x5d, 0xbd, 0x57, 0x8f, 0x60, 0x2d,
	0x52, 0x78, 0x56, 0x84, 0x6f, 0xa3, 0x96, 0x08, 0x6a, 0x48, 0xf2, 0x04, 0xba, 0x7e, 0xa2, 0x14,
	0x16, 0xdf, 0xc7, 0x25, 0x75, 0x25, 0xef, 0x4d, 0x61, 0x8b, 0x62, 0x14, 0xa4, 0x5f, 0x12, 0x54,
	0xe9, 0xff, 0xde, 0x8a, 0x0c, 0x61, 0x83, 0x9d, 0x31, 0x1e, 0xb0, 0x49, 0x80, 0xbb, 0xf3, 0x1f,
	0x64, 0x1d, 0x1e, 0xfd, 0xb4, 0xc0, 0x2e, 0x85, 0xe4, 0x00, 0x6e, 0xbd, 0x47, 0xbd, 0xd4, 0xad,
	0xdc, 0x21, 0xc5, 0xef, 0x87, 0x5a, 0x71, 0x71, 0xb2, 0xf5, 0xb0, 0x42, 0xae, 0x1a, 0xd1, 0x6b,
	0x90, 0xe7, 0xb0, 0xbe, 0x40, 0x5d, 0xe2, 0x53, 0xbf, 0xbf, 0xd7, 0x98, 0x74, 0xcc, 0x9f, 0xc1,
	0xce, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x46, 0xd7, 0x7c, 0x33, 0x06, 0x00, 0x00,
}
