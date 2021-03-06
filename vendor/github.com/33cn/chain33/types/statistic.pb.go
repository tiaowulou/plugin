// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statistic.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

//手续费
type TotalFee struct {
	Fee                  int64    `protobuf:"varint,1,opt,name=fee,proto3" json:"fee,omitempty"`
	TxCount              int64    `protobuf:"varint,2,opt,name=txCount,proto3" json:"txCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalFee) Reset()         { *m = TotalFee{} }
func (m *TotalFee) String() string { return proto.CompactTextString(m) }
func (*TotalFee) ProtoMessage()    {}
func (*TotalFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{0}
}

func (m *TotalFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalFee.Unmarshal(m, b)
}
func (m *TotalFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalFee.Marshal(b, m, deterministic)
}
func (m *TotalFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalFee.Merge(m, src)
}
func (m *TotalFee) XXX_Size() int {
	return xxx_messageInfo_TotalFee.Size(m)
}
func (m *TotalFee) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalFee.DiscardUnknown(m)
}

var xxx_messageInfo_TotalFee proto.InternalMessageInfo

func (m *TotalFee) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TotalFee) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

//查询symbol代币总额
type ReqGetTotalCoins struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	StateHash            []byte   `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	StartKey             []byte   `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Execer               string   `protobuf:"bytes,5,opt,name=execer,proto3" json:"execer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGetTotalCoins) Reset()         { *m = ReqGetTotalCoins{} }
func (m *ReqGetTotalCoins) String() string { return proto.CompactTextString(m) }
func (*ReqGetTotalCoins) ProtoMessage()    {}
func (*ReqGetTotalCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{1}
}

func (m *ReqGetTotalCoins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGetTotalCoins.Unmarshal(m, b)
}
func (m *ReqGetTotalCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGetTotalCoins.Marshal(b, m, deterministic)
}
func (m *ReqGetTotalCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGetTotalCoins.Merge(m, src)
}
func (m *ReqGetTotalCoins) XXX_Size() int {
	return xxx_messageInfo_ReqGetTotalCoins.Size(m)
}
func (m *ReqGetTotalCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGetTotalCoins.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGetTotalCoins proto.InternalMessageInfo

func (m *ReqGetTotalCoins) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetTotalCoins) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetTotalCoins) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *ReqGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqGetTotalCoins) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

//查询symbol代币总额应答
type ReplyGetTotalCoins struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Num                  int64    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	NextKey              []byte   `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyGetTotalCoins) Reset()         { *m = ReplyGetTotalCoins{} }
func (m *ReplyGetTotalCoins) String() string { return proto.CompactTextString(m) }
func (*ReplyGetTotalCoins) ProtoMessage()    {}
func (*ReplyGetTotalCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{2}
}

func (m *ReplyGetTotalCoins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyGetTotalCoins.Unmarshal(m, b)
}
func (m *ReplyGetTotalCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyGetTotalCoins.Marshal(b, m, deterministic)
}
func (m *ReplyGetTotalCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyGetTotalCoins.Merge(m, src)
}
func (m *ReplyGetTotalCoins) XXX_Size() int {
	return xxx_messageInfo_ReplyGetTotalCoins.Size(m)
}
func (m *ReplyGetTotalCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyGetTotalCoins.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyGetTotalCoins proto.InternalMessageInfo

func (m *ReplyGetTotalCoins) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetTotalCoins) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

//迭代查询symbol代币总额
type IterateRangeByStateHash struct {
	StateHash            []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start                []byte   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  []byte   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IterateRangeByStateHash) Reset()         { *m = IterateRangeByStateHash{} }
func (m *IterateRangeByStateHash) String() string { return proto.CompactTextString(m) }
func (*IterateRangeByStateHash) ProtoMessage()    {}
func (*IterateRangeByStateHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{3}
}

func (m *IterateRangeByStateHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IterateRangeByStateHash.Unmarshal(m, b)
}
func (m *IterateRangeByStateHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IterateRangeByStateHash.Marshal(b, m, deterministic)
}
func (m *IterateRangeByStateHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IterateRangeByStateHash.Merge(m, src)
}
func (m *IterateRangeByStateHash) XXX_Size() int {
	return xxx_messageInfo_IterateRangeByStateHash.Size(m)
}
func (m *IterateRangeByStateHash) XXX_DiscardUnknown() {
	xxx_messageInfo_IterateRangeByStateHash.DiscardUnknown(m)
}

var xxx_messageInfo_IterateRangeByStateHash proto.InternalMessageInfo

func (m *IterateRangeByStateHash) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *IterateRangeByStateHash) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IterateRangeByStateHash) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *IterateRangeByStateHash) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TicketStatistic struct {
	//当前在挖的ticket
	CurrentOpenCount int64 `protobuf:"varint,1,opt,name=currentOpenCount,proto3" json:"currentOpenCount,omitempty"`
	//一共挖到的ticket
	TotalMinerCount int64 `protobuf:"varint,2,opt,name=totalMinerCount,proto3" json:"totalMinerCount,omitempty"`
	//一共取消的ticket
	TotalCancleCount     int64    `protobuf:"varint,3,opt,name=totalCancleCount,proto3" json:"totalCancleCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TicketStatistic) Reset()         { *m = TicketStatistic{} }
func (m *TicketStatistic) String() string { return proto.CompactTextString(m) }
func (*TicketStatistic) ProtoMessage()    {}
func (*TicketStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{4}
}

func (m *TicketStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketStatistic.Unmarshal(m, b)
}
func (m *TicketStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketStatistic.Marshal(b, m, deterministic)
}
func (m *TicketStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketStatistic.Merge(m, src)
}
func (m *TicketStatistic) XXX_Size() int {
	return xxx_messageInfo_TicketStatistic.Size(m)
}
func (m *TicketStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_TicketStatistic proto.InternalMessageInfo

func (m *TicketStatistic) GetCurrentOpenCount() int64 {
	if m != nil {
		return m.CurrentOpenCount
	}
	return 0
}

func (m *TicketStatistic) GetTotalMinerCount() int64 {
	if m != nil {
		return m.TotalMinerCount
	}
	return 0
}

func (m *TicketStatistic) GetTotalCancleCount() int64 {
	if m != nil {
		return m.TotalCancleCount
	}
	return 0
}

type TicketMinerInfo struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	// 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status     int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	PrevStatus int32 `protobuf:"varint,3,opt,name=prevStatus,proto3" json:"prevStatus,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,4,opt,name=isGenesis,proto3" json:"isGenesis,omitempty"`
	//创建ticket时间
	CreateTime int64 `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// ticket挖矿时间
	MinerTime int64 `protobuf:"varint,6,opt,name=minerTime,proto3" json:"minerTime,omitempty"`
	//关闭ticket时间
	CloseTime int64 `protobuf:"varint,7,opt,name=closeTime,proto3" json:"closeTime,omitempty"`
	//挖到的币的数目
	MinerValue           int64    `protobuf:"varint,8,opt,name=minerValue,proto3" json:"minerValue,omitempty"`
	MinerAddress         string   `protobuf:"bytes,9,opt,name=minerAddress,proto3" json:"minerAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TicketMinerInfo) Reset()         { *m = TicketMinerInfo{} }
func (m *TicketMinerInfo) String() string { return proto.CompactTextString(m) }
func (*TicketMinerInfo) ProtoMessage()    {}
func (*TicketMinerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{5}
}

func (m *TicketMinerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketMinerInfo.Unmarshal(m, b)
}
func (m *TicketMinerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketMinerInfo.Marshal(b, m, deterministic)
}
func (m *TicketMinerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketMinerInfo.Merge(m, src)
}
func (m *TicketMinerInfo) XXX_Size() int {
	return xxx_messageInfo_TicketMinerInfo.Size(m)
}
func (m *TicketMinerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketMinerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TicketMinerInfo proto.InternalMessageInfo

func (m *TicketMinerInfo) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *TicketMinerInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TicketMinerInfo) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *TicketMinerInfo) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *TicketMinerInfo) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *TicketMinerInfo) GetCloseTime() int64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *TicketMinerInfo) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

type TotalAmount struct {
	// 统计的总数
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TotalAmount) Reset()         { *m = TotalAmount{} }
func (m *TotalAmount) String() string { return proto.CompactTextString(m) }
func (*TotalAmount) ProtoMessage()    {}
func (*TotalAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{6}
}

func (m *TotalAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TotalAmount.Unmarshal(m, b)
}
func (m *TotalAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TotalAmount.Marshal(b, m, deterministic)
}
func (m *TotalAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalAmount.Merge(m, src)
}
func (m *TotalAmount) XXX_Size() int {
	return xxx_messageInfo_TotalAmount.Size(m)
}
func (m *TotalAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalAmount.DiscardUnknown(m)
}

var xxx_messageInfo_TotalAmount proto.InternalMessageInfo

func (m *TotalAmount) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

//查询symbol在合约中的代币总额，如果execAddr为空，则为查询symbol在所有合约中的代币总额
type ReqGetExecBalance struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	StateHash            []byte   `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Addr                 []byte   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	ExecAddr             []byte   `protobuf:"bytes,4,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Execer               string   `protobuf:"bytes,5,opt,name=execer,proto3" json:"execer,omitempty"`
	Count                int64    `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	NextKey              []byte   `protobuf:"bytes,7,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGetExecBalance) Reset()         { *m = ReqGetExecBalance{} }
func (m *ReqGetExecBalance) String() string { return proto.CompactTextString(m) }
func (*ReqGetExecBalance) ProtoMessage()    {}
func (*ReqGetExecBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{7}
}

func (m *ReqGetExecBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGetExecBalance.Unmarshal(m, b)
}
func (m *ReqGetExecBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGetExecBalance.Marshal(b, m, deterministic)
}
func (m *ReqGetExecBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGetExecBalance.Merge(m, src)
}
func (m *ReqGetExecBalance) XXX_Size() int {
	return xxx_messageInfo_ReqGetExecBalance.Size(m)
}
func (m *ReqGetExecBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGetExecBalance.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGetExecBalance proto.InternalMessageInfo

func (m *ReqGetExecBalance) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqGetExecBalance) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReqGetExecBalance) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ReqGetExecBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ReqGetExecBalance) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqGetExecBalance) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

type ExecBalanceItem struct {
	ExecAddr             []byte   `protobuf:"bytes,1,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Frozen               int64    `protobuf:"varint,2,opt,name=frozen,proto3" json:"frozen,omitempty"`
	Active               int64    `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecBalanceItem) Reset()         { *m = ExecBalanceItem{} }
func (m *ExecBalanceItem) String() string { return proto.CompactTextString(m) }
func (*ExecBalanceItem) ProtoMessage()    {}
func (*ExecBalanceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{8}
}

func (m *ExecBalanceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecBalanceItem.Unmarshal(m, b)
}
func (m *ExecBalanceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecBalanceItem.Marshal(b, m, deterministic)
}
func (m *ExecBalanceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecBalanceItem.Merge(m, src)
}
func (m *ExecBalanceItem) XXX_Size() int {
	return xxx_messageInfo_ExecBalanceItem.Size(m)
}
func (m *ExecBalanceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecBalanceItem.DiscardUnknown(m)
}

var xxx_messageInfo_ExecBalanceItem proto.InternalMessageInfo

func (m *ExecBalanceItem) GetExecAddr() []byte {
	if m != nil {
		return m.ExecAddr
	}
	return nil
}

func (m *ExecBalanceItem) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *ExecBalanceItem) GetActive() int64 {
	if m != nil {
		return m.Active
	}
	return 0
}

//查询symbol在合约中的代币总额应答
type ReplyGetExecBalance struct {
	Amount               int64              `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountFrozen         int64              `protobuf:"varint,2,opt,name=amountFrozen,proto3" json:"amountFrozen,omitempty"`
	AmountActive         int64              `protobuf:"varint,3,opt,name=amountActive,proto3" json:"amountActive,omitempty"`
	NextKey              []byte             `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	Items                []*ExecBalanceItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReplyGetExecBalance) Reset()         { *m = ReplyGetExecBalance{} }
func (m *ReplyGetExecBalance) String() string { return proto.CompactTextString(m) }
func (*ReplyGetExecBalance) ProtoMessage()    {}
func (*ReplyGetExecBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_405f6cee9ed2da7e, []int{9}
}

func (m *ReplyGetExecBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyGetExecBalance.Unmarshal(m, b)
}
func (m *ReplyGetExecBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyGetExecBalance.Marshal(b, m, deterministic)
}
func (m *ReplyGetExecBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyGetExecBalance.Merge(m, src)
}
func (m *ReplyGetExecBalance) XXX_Size() int {
	return xxx_messageInfo_ReplyGetExecBalance.Size(m)
}
func (m *ReplyGetExecBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyGetExecBalance.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyGetExecBalance proto.InternalMessageInfo

func (m *ReplyGetExecBalance) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountFrozen() int64 {
	if m != nil {
		return m.AmountFrozen
	}
	return 0
}

func (m *ReplyGetExecBalance) GetAmountActive() int64 {
	if m != nil {
		return m.AmountActive
	}
	return 0
}

func (m *ReplyGetExecBalance) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *ReplyGetExecBalance) GetItems() []*ExecBalanceItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*TotalFee)(nil), "types.TotalFee")
	proto.RegisterType((*ReqGetTotalCoins)(nil), "types.ReqGetTotalCoins")
	proto.RegisterType((*ReplyGetTotalCoins)(nil), "types.ReplyGetTotalCoins")
	proto.RegisterType((*IterateRangeByStateHash)(nil), "types.IterateRangeByStateHash")
	proto.RegisterType((*TicketStatistic)(nil), "types.TicketStatistic")
	proto.RegisterType((*TicketMinerInfo)(nil), "types.TicketMinerInfo")
	proto.RegisterType((*TotalAmount)(nil), "types.TotalAmount")
	proto.RegisterType((*ReqGetExecBalance)(nil), "types.ReqGetExecBalance")
	proto.RegisterType((*ExecBalanceItem)(nil), "types.ExecBalanceItem")
	proto.RegisterType((*ReplyGetExecBalance)(nil), "types.ReplyGetExecBalance")
}

func init() { proto.RegisterFile("statistic.proto", fileDescriptor_405f6cee9ed2da7e) }

var fileDescriptor_405f6cee9ed2da7e = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xeb, 0x3a, 0x4d, 0xa7, 0x95, 0x12, 0x96, 0xaa, 0x58, 0x88, 0x9f, 0xca, 0x5c, 0x2a,
	0x84, 0x52, 0x89, 0x48, 0xdc, 0xdb, 0x88, 0x96, 0x08, 0x21, 0x24, 0xb7, 0xe2, 0x80, 0xc4, 0x61,
	0xbb, 0x99, 0xb6, 0x16, 0xf6, 0x3a, 0x78, 0xc7, 0x55, 0xc2, 0x63, 0xc0, 0x23, 0xf0, 0x1c, 0x1c,
	0x78, 0x33, 0xb4, 0xe3, 0x8d, 0x7f, 0x52, 0x7a, 0xe1, 0xb6, 0xdf, 0x37, 0x3b, 0xbf, 0x3b, 0xdf,
	0xc2, 0xc0, 0x90, 0xa4, 0xc4, 0x50, 0xa2, 0x46, 0xf3, 0x22, 0xa7, 0x5c, 0x04, 0xb4, 0x9c, 0xa3,
	0x89, 0xde, 0x40, 0xff, 0x22, 0x27, 0x99, 0x9e, 0x22, 0x8a, 0x21, 0xf8, 0x57, 0x88, 0xa1, 0x77,
	0xe0, 0x1d, 0xfa, 0xb1, 0x3d, 0x8a, 0x10, 0xb6, 0x68, 0x31, 0xc9, 0x4b, 0x4d, 0xe1, 0x06, 0xb3,
	0x2b, 0x18, 0xfd, 0xf0, 0x60, 0x18, 0xe3, 0xb7, 0x33, 0x24, 0x76, 0x9f, 0xe4, 0x89, 0x36, 0x62,
	0x1f, 0x7a, 0x66, 0x99, 0x5d, 0xe6, 0x29, 0xc7, 0xd8, 0x8e, 0x1d, 0x12, 0x4f, 0x60, 0xdb, 0xa6,
	0xc7, 0x77, 0xd2, 0xdc, 0x70, 0xa0, 0xdd, 0xb8, 0x21, 0xc4, 0x63, 0xe8, 0x1b, 0x92, 0x05, 0xbd,
	0xc7, 0x65, 0xe8, 0xb3, 0xb1, 0xc6, 0x62, 0x0f, 0x02, 0xc5, 0xe9, 0x37, 0x39, 0x7d, 0x05, 0x6c,
	0x1e, 0x5c, 0xa0, 0xc2, 0x22, 0x0c, 0xaa, 0x3c, 0x15, 0x8a, 0x34, 0x88, 0x18, 0xe7, 0xe9, 0xb2,
	0x5b, 0x55, 0x1d, 0xc3, 0x6b, 0xc7, 0x18, 0x82, 0xaf, 0xcb, 0xcc, 0xb5, 0x65, 0x8f, 0x36, 0xaa,
	0xcc, 0xf8, 0xa2, 0xcf, 0xa4, 0x43, 0x76, 0x08, 0x1a, 0x17, 0x5c, 0xde, 0x26, 0x97, 0xb7, 0x82,
	0x51, 0x09, 0x8f, 0xa6, 0x84, 0x85, 0x24, 0x8c, 0xa5, 0xbe, 0xc6, 0x93, 0xe5, 0x79, 0xdd, 0x54,
	0xa7, 0x65, 0x6f, 0xbd, 0xe5, 0x3d, 0x08, 0xb8, 0x45, 0x37, 0x8c, 0x0a, 0xd8, 0x92, 0x50, 0xcf,
	0xdc, 0x0c, 0xec, 0xf1, 0xdf, 0xed, 0x47, 0x3f, 0x3d, 0x18, 0x5c, 0x24, 0xea, 0x2b, 0xd2, 0xf9,
	0xea, 0x51, 0xc5, 0x4b, 0x18, 0xaa, 0xb2, 0x28, 0x50, 0xd3, 0xc7, 0x39, 0xea, 0x49, 0xab, 0xdf,
	0x3b, 0xbc, 0x38, 0x84, 0x01, 0xd9, 0xf1, 0x7c, 0x48, 0x34, 0x16, 0xed, 0xd7, 0x5d, 0xa7, 0x6d,
	0x54, 0xa6, 0x26, 0x52, 0xab, 0x14, 0x27, 0xad, 0xe1, 0xdc, 0xe1, 0xa3, 0x5f, 0x1b, 0xab, 0xaa,
	0x38, 0xc0, 0x54, 0x5f, 0xe5, 0xf6, 0x69, 0x89, 0xa9, 0xe9, 0xcc, 0xad, 0x44, 0x8d, 0x79, 0x59,
	0x48, 0x52, 0x69, 0x38, 0x79, 0x10, 0x3b, 0x24, 0x9e, 0x01, 0xcc, 0x0b, 0xbc, 0x3d, 0xaf, 0x6c,
	0x3e, 0xdb, 0x5a, 0x8c, 0x9d, 0x6c, 0x62, 0xce, 0x50, 0xa3, 0x49, 0x0c, 0xcf, 0xa5, 0x1f, 0x37,
	0x84, 0xf5, 0x56, 0x05, 0x4a, 0xc2, 0x8b, 0x24, 0x43, 0x5e, 0x0f, 0x3f, 0x6e, 0x31, 0xd6, 0x3b,
	0xb3, 0xe5, 0xb1, 0xb9, 0xc7, 0xe6, 0x86, 0xb0, 0x56, 0x95, 0xe6, 0xa6, 0x72, 0xde, 0xaa, 0xac,
	0x35, 0x61, 0x63, 0xf3, 0xd5, 0x4f, 0x32, 0x2d, 0x31, 0xec, 0x57, 0xb1, 0x1b, 0x46, 0x44, 0xb0,
	0xcb, 0xe8, 0x78, 0x36, 0x2b, 0xd0, 0x98, 0x70, 0x9b, 0x3b, 0xee, 0x70, 0xd1, 0x0b, 0xd8, 0xe1,
	0xd5, 0x3c, 0xae, 0x76, 0x6b, 0x0f, 0x02, 0x1e, 0xe4, 0x6a, 0x37, 0x19, 0x44, 0x7f, 0x3c, 0x78,
	0x50, 0x89, 0xeb, 0xed, 0x02, 0xd5, 0x89, 0x4c, 0xa5, 0x56, 0xf8, 0x9f, 0xea, 0x12, 0xb0, 0x29,
	0x67, 0xb3, 0xc2, 0x6d, 0x15, 0x9f, 0xed, 0xb3, 0x58, 0xc5, 0xd8, 0x9a, 0xdc, 0x4a, 0xd7, 0xf8,
	0x3e, 0x6d, 0x35, 0xab, 0xd8, 0x6b, 0xab, 0xa8, 0xa5, 0x8d, 0xad, 0xae, 0x36, 0xbe, 0xc0, 0xa0,
	0x55, 0xfc, 0x94, 0x30, 0xeb, 0xa4, 0xf5, 0xee, 0xa6, 0xbd, 0x2a, 0xf2, 0xef, 0xa8, 0xdd, 0x2a,
	0x3a, 0xc4, 0xa2, 0x54, 0x94, 0xdc, 0x62, 0x2d, 0x4a, 0x46, 0xd1, 0x6f, 0x0f, 0x1e, 0xae, 0xb4,
	0xbe, 0x36, 0x24, 0x27, 0x62, 0xaf, 0x23, 0xe2, 0x08, 0x76, 0xab, 0xd3, 0x69, 0x3b, 0x4b, 0x87,
	0x6b, 0xee, 0x1c, 0xb7, 0x33, 0x76, 0xb8, 0xfb, 0x3f, 0x03, 0xf1, 0x0a, 0x82, 0x84, 0x30, 0x33,
	0x61, 0x70, 0xe0, 0x1f, 0xee, 0xbc, 0xde, 0x1f, 0xf1, 0x07, 0x3b, 0x5a, 0x1b, 0x42, 0x5c, 0x5d,
	0x3a, 0x79, 0xfe, 0xf9, 0xe9, 0x75, 0x42, 0x37, 0xe5, 0xe5, 0x48, 0xe5, 0xd9, 0xd1, 0x78, 0xac,
	0xf4, 0x91, 0xba, 0x91, 0x89, 0x1e, 0x8f, 0x8f, 0xd8, 0xef, 0xb2, 0xc7, 0xdf, 0xf4, 0xf8, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x11, 0x9f, 0xa6, 0xb9, 0x05, 0x00, 0x00,
}
