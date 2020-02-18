// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ethbridge.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EthBridgeStatus int32

const (
	EthBridgeStatus_PendingStatusText EthBridgeStatus = 0
	EthBridgeStatus_SuccessStatusText EthBridgeStatus = 1
	EthBridgeStatus_FailedStatusText  EthBridgeStatus = 2
)

var EthBridgeStatus_name = map[int32]string{
	0: "PendingStatusText",
	1: "SuccessStatusText",
	2: "FailedStatusText",
}
var EthBridgeStatus_value = map[string]int32{
	"PendingStatusText": 0,
	"SuccessStatusText": 1,
	"FailedStatusText":  2,
}

func (x EthBridgeStatus) String() string {
	return proto.EnumName(EthBridgeStatus_name, int32(x))
}
func (EthBridgeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{0}
}

type Eth2Chain33Action struct {
	// Types that are valid to be assigned to Value:
	//	*Eth2Chain33Action_EthBridgeClaim
	//	*Eth2Chain33Action_MsgBurn
	//	*Eth2Chain33Action_MsgLock
	Value                isEth2Chain33Action_Value `protobuf_oneof:"value"`
	Ty                   int32                     `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Eth2Chain33Action) Reset()         { *m = Eth2Chain33Action{} }
func (m *Eth2Chain33Action) String() string { return proto.CompactTextString(m) }
func (*Eth2Chain33Action) ProtoMessage()    {}
func (*Eth2Chain33Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{0}
}
func (m *Eth2Chain33Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Eth2Chain33Action.Unmarshal(m, b)
}
func (m *Eth2Chain33Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Eth2Chain33Action.Marshal(b, m, deterministic)
}
func (dst *Eth2Chain33Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Eth2Chain33Action.Merge(dst, src)
}
func (m *Eth2Chain33Action) XXX_Size() int {
	return xxx_messageInfo_Eth2Chain33Action.Size(m)
}
func (m *Eth2Chain33Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Eth2Chain33Action.DiscardUnknown(m)
}

var xxx_messageInfo_Eth2Chain33Action proto.InternalMessageInfo

type isEth2Chain33Action_Value interface {
	isEth2Chain33Action_Value()
}

type Eth2Chain33Action_EthBridgeClaim struct {
	EthBridgeClaim *EthBridgeClaim `protobuf:"bytes,1,opt,name=ethBridgeClaim,oneof"`
}
type Eth2Chain33Action_MsgBurn struct {
	MsgBurn *MsgBurn `protobuf:"bytes,2,opt,name=msgBurn,oneof"`
}
type Eth2Chain33Action_MsgLock struct {
	MsgLock *MsgLock `protobuf:"bytes,3,opt,name=msgLock,oneof"`
}

func (*Eth2Chain33Action_EthBridgeClaim) isEth2Chain33Action_Value() {}
func (*Eth2Chain33Action_MsgBurn) isEth2Chain33Action_Value()        {}
func (*Eth2Chain33Action_MsgLock) isEth2Chain33Action_Value()        {}

func (m *Eth2Chain33Action) GetValue() isEth2Chain33Action_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Eth2Chain33Action) GetEthBridgeClaim() *EthBridgeClaim {
	if x, ok := m.GetValue().(*Eth2Chain33Action_EthBridgeClaim); ok {
		return x.EthBridgeClaim
	}
	return nil
}

func (m *Eth2Chain33Action) GetMsgBurn() *MsgBurn {
	if x, ok := m.GetValue().(*Eth2Chain33Action_MsgBurn); ok {
		return x.MsgBurn
	}
	return nil
}

func (m *Eth2Chain33Action) GetMsgLock() *MsgLock {
	if x, ok := m.GetValue().(*Eth2Chain33Action_MsgLock); ok {
		return x.MsgLock
	}
	return nil
}

func (m *Eth2Chain33Action) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Eth2Chain33Action) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Eth2Chain33Action_OneofMarshaler, _Eth2Chain33Action_OneofUnmarshaler, _Eth2Chain33Action_OneofSizer, []interface{}{
		(*Eth2Chain33Action_EthBridgeClaim)(nil),
		(*Eth2Chain33Action_MsgBurn)(nil),
		(*Eth2Chain33Action_MsgLock)(nil),
	}
}

func _Eth2Chain33Action_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Eth2Chain33Action)
	// value
	switch x := m.Value.(type) {
	case *Eth2Chain33Action_EthBridgeClaim:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EthBridgeClaim); err != nil {
			return err
		}
	case *Eth2Chain33Action_MsgBurn:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MsgBurn); err != nil {
			return err
		}
	case *Eth2Chain33Action_MsgLock:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MsgLock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Eth2Chain33Action.Value has unexpected type %T", x)
	}
	return nil
}

func _Eth2Chain33Action_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Eth2Chain33Action)
	switch tag {
	case 1: // value.ethBridgeClaim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EthBridgeClaim)
		err := b.DecodeMessage(msg)
		m.Value = &Eth2Chain33Action_EthBridgeClaim{msg}
		return true, err
	case 2: // value.msgBurn
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MsgBurn)
		err := b.DecodeMessage(msg)
		m.Value = &Eth2Chain33Action_MsgBurn{msg}
		return true, err
	case 3: // value.msgLock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MsgLock)
		err := b.DecodeMessage(msg)
		m.Value = &Eth2Chain33Action_MsgLock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Eth2Chain33Action_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Eth2Chain33Action)
	// value
	switch x := m.Value.(type) {
	case *Eth2Chain33Action_EthBridgeClaim:
		s := proto.Size(x.EthBridgeClaim)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Eth2Chain33Action_MsgBurn:
		s := proto.Size(x.MsgBurn)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Eth2Chain33Action_MsgLock:
		s := proto.Size(x.MsgLock)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// EthBridgeClaim is a structure that contains all the data for a particular bridge claim
type EthBridgeClaim struct {
	EthereumChainID       int64            `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	BridgeContractAddress *EthereumAddress `protobuf:"bytes,2,opt,name=BridgeContractAddress" json:"BridgeContractAddress,omitempty"`
	Nonce                 int64            `protobuf:"varint,3,opt,name=Nonce" json:"Nonce,omitempty"`
	Symbol                string           `protobuf:"bytes,4,opt,name=Symbol" json:"Symbol,omitempty"`
	TokenContractAddress  *EthereumAddress `protobuf:"bytes,5,opt,name=TokenContractAddress" json:"TokenContractAddress,omitempty"`
	EthereumSender        *EthereumAddress `protobuf:"bytes,6,opt,name=EthereumSender" json:"EthereumSender,omitempty"`
	Chain33Receiver       *Chain33Address  `protobuf:"bytes,7,opt,name=Chain33Receiver" json:"Chain33Receiver,omitempty"`
	ValidatorAddress      *Chain33Address  `protobuf:"bytes,8,opt,name=ValidatorAddress" json:"ValidatorAddress,omitempty"`
	Amount                uint64           `protobuf:"varint,9,opt,name=Amount" json:"Amount,omitempty"`
	ClaimType             int64            `protobuf:"varint,10,opt,name=ClaimType" json:"ClaimType,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}
func (*EthBridgeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{1}
}
func (m *EthBridgeClaim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthBridgeClaim.Unmarshal(m, b)
}
func (m *EthBridgeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthBridgeClaim.Marshal(b, m, deterministic)
}
func (dst *EthBridgeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthBridgeClaim.Merge(dst, src)
}
func (m *EthBridgeClaim) XXX_Size() int {
	return xxx_messageInfo_EthBridgeClaim.Size(m)
}
func (m *EthBridgeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EthBridgeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EthBridgeClaim proto.InternalMessageInfo

func (m *EthBridgeClaim) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *EthBridgeClaim) GetBridgeContractAddress() *EthereumAddress {
	if m != nil {
		return m.BridgeContractAddress
	}
	return nil
}

func (m *EthBridgeClaim) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthBridgeClaim) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EthBridgeClaim) GetTokenContractAddress() *EthereumAddress {
	if m != nil {
		return m.TokenContractAddress
	}
	return nil
}

func (m *EthBridgeClaim) GetEthereumSender() *EthereumAddress {
	if m != nil {
		return m.EthereumSender
	}
	return nil
}

func (m *EthBridgeClaim) GetChain33Receiver() *Chain33Address {
	if m != nil {
		return m.Chain33Receiver
	}
	return nil
}

func (m *EthBridgeClaim) GetValidatorAddress() *Chain33Address {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *EthBridgeClaim) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *EthBridgeClaim) GetClaimType() int64 {
	if m != nil {
		return m.ClaimType
	}
	return 0
}

type EthereumAddress struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumAddress) Reset()         { *m = EthereumAddress{} }
func (m *EthereumAddress) String() string { return proto.CompactTextString(m) }
func (*EthereumAddress) ProtoMessage()    {}
func (*EthereumAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{2}
}
func (m *EthereumAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumAddress.Unmarshal(m, b)
}
func (m *EthereumAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumAddress.Marshal(b, m, deterministic)
}
func (dst *EthereumAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAddress.Merge(dst, src)
}
func (m *EthereumAddress) XXX_Size() int {
	return xxx_messageInfo_EthereumAddress.Size(m)
}
func (m *EthereumAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAddress proto.InternalMessageInfo

func (m *EthereumAddress) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type Chain33Address struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chain33Address) Reset()         { *m = Chain33Address{} }
func (m *Chain33Address) String() string { return proto.CompactTextString(m) }
func (*Chain33Address) ProtoMessage()    {}
func (*Chain33Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{3}
}
func (m *Chain33Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chain33Address.Unmarshal(m, b)
}
func (m *Chain33Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chain33Address.Marshal(b, m, deterministic)
}
func (dst *Chain33Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain33Address.Merge(dst, src)
}
func (m *Chain33Address) XXX_Size() int {
	return xxx_messageInfo_Chain33Address.Size(m)
}
func (m *Chain33Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain33Address.DiscardUnknown(m)
}

var xxx_messageInfo_Chain33Address proto.InternalMessageInfo

func (m *Chain33Address) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// OracleClaimContent is the details of how the content of the claim for each validator will be stored in the oracle
type OracleClaimContent struct {
	Chain33Receiver      *Chain33Address `protobuf:"bytes,1,opt,name=Chain33Receiver" json:"Chain33Receiver,omitempty"`
	Amount               uint64          `protobuf:"varint,2,opt,name=Amount" json:"Amount,omitempty"`
	ClaimType            int64           `protobuf:"varint,3,opt,name=ClaimType" json:"ClaimType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OracleClaimContent) Reset()         { *m = OracleClaimContent{} }
func (m *OracleClaimContent) String() string { return proto.CompactTextString(m) }
func (*OracleClaimContent) ProtoMessage()    {}
func (*OracleClaimContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{4}
}
func (m *OracleClaimContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleClaimContent.Unmarshal(m, b)
}
func (m *OracleClaimContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleClaimContent.Marshal(b, m, deterministic)
}
func (dst *OracleClaimContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleClaimContent.Merge(dst, src)
}
func (m *OracleClaimContent) XXX_Size() int {
	return xxx_messageInfo_OracleClaimContent.Size(m)
}
func (m *OracleClaimContent) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleClaimContent.DiscardUnknown(m)
}

var xxx_messageInfo_OracleClaimContent proto.InternalMessageInfo

func (m *OracleClaimContent) GetChain33Receiver() *Chain33Address {
	if m != nil {
		return m.Chain33Receiver
	}
	return nil
}

func (m *OracleClaimContent) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OracleClaimContent) GetClaimType() int64 {
	if m != nil {
		return m.ClaimType
	}
	return 0
}

// MsgBurn defines a message for burning coins and triggering a related event
type MsgBurn struct {
	EthereumChainID      int64            `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	TokenContract        *EthereumAddress `protobuf:"bytes,2,opt,name=TokenContract" json:"TokenContract,omitempty"`
	Chain33Sender        *Chain33Address  `protobuf:"bytes,3,opt,name=Chain33Sender" json:"Chain33Sender,omitempty"`
	EthereumReceiver     *EthereumAddress `protobuf:"bytes,4,opt,name=EthereumReceiver" json:"EthereumReceiver,omitempty"`
	Amount               uint64           `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MsgBurn) Reset()         { *m = MsgBurn{} }
func (m *MsgBurn) String() string { return proto.CompactTextString(m) }
func (*MsgBurn) ProtoMessage()    {}
func (*MsgBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{5}
}
func (m *MsgBurn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBurn.Unmarshal(m, b)
}
func (m *MsgBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBurn.Marshal(b, m, deterministic)
}
func (dst *MsgBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurn.Merge(dst, src)
}
func (m *MsgBurn) XXX_Size() int {
	return xxx_messageInfo_MsgBurn.Size(m)
}
func (m *MsgBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurn proto.InternalMessageInfo

func (m *MsgBurn) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *MsgBurn) GetTokenContract() *EthereumAddress {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *MsgBurn) GetChain33Sender() *Chain33Address {
	if m != nil {
		return m.Chain33Sender
	}
	return nil
}

func (m *MsgBurn) GetEthereumReceiver() *EthereumAddress {
	if m != nil {
		return m.EthereumReceiver
	}
	return nil
}

func (m *MsgBurn) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgLock defines a message for locking coins and triggering a related event
type MsgLock struct {
	EthereumChainID      int64            `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	TokenContract        *EthereumAddress `protobuf:"bytes,2,opt,name=TokenContract" json:"TokenContract,omitempty"`
	Chain33Sender        *Chain33Address  `protobuf:"bytes,3,opt,name=Chain33Sender" json:"Chain33Sender,omitempty"`
	EthereumReceiver     *EthereumAddress `protobuf:"bytes,4,opt,name=EthereumReceiver" json:"EthereumReceiver,omitempty"`
	Amount               uint64           `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MsgLock) Reset()         { *m = MsgLock{} }
func (m *MsgLock) String() string { return proto.CompactTextString(m) }
func (*MsgLock) ProtoMessage()    {}
func (*MsgLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{6}
}
func (m *MsgLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLock.Unmarshal(m, b)
}
func (m *MsgLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLock.Marshal(b, m, deterministic)
}
func (dst *MsgLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLock.Merge(dst, src)
}
func (m *MsgLock) XXX_Size() int {
	return xxx_messageInfo_MsgLock.Size(m)
}
func (m *MsgLock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLock proto.InternalMessageInfo

func (m *MsgLock) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *MsgLock) GetTokenContract() *EthereumAddress {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

func (m *MsgLock) GetChain33Sender() *Chain33Address {
	if m != nil {
		return m.Chain33Sender
	}
	return nil
}

func (m *MsgLock) GetEthereumReceiver() *EthereumAddress {
	if m != nil {
		return m.EthereumReceiver
	}
	return nil
}

func (m *MsgLock) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type QueryEthProphecyParams struct {
	EthereumChainID       int64            `protobuf:"varint,1,opt,name=EthereumChainID" json:"EthereumChainID,omitempty"`
	BridgeContractAddress *EthereumAddress `protobuf:"bytes,2,opt,name=BridgeContractAddress" json:"BridgeContractAddress,omitempty"`
	Nonce                 int64            `protobuf:"varint,3,opt,name=Nonce" json:"Nonce,omitempty"`
	Symbol                string           `protobuf:"bytes,4,opt,name=Symbol" json:"Symbol,omitempty"`
	TokenContractAddress  *EthereumAddress `protobuf:"bytes,5,opt,name=TokenContractAddress" json:"TokenContractAddress,omitempty"`
	EthereumSender        *EthereumAddress `protobuf:"bytes,6,opt,name=EthereumSender" json:"EthereumSender,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *QueryEthProphecyParams) Reset()         { *m = QueryEthProphecyParams{} }
func (m *QueryEthProphecyParams) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyParams) ProtoMessage()    {}
func (*QueryEthProphecyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{7}
}
func (m *QueryEthProphecyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryEthProphecyParams.Unmarshal(m, b)
}
func (m *QueryEthProphecyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryEthProphecyParams.Marshal(b, m, deterministic)
}
func (dst *QueryEthProphecyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyParams.Merge(dst, src)
}
func (m *QueryEthProphecyParams) XXX_Size() int {
	return xxx_messageInfo_QueryEthProphecyParams.Size(m)
}
func (m *QueryEthProphecyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyParams proto.InternalMessageInfo

func (m *QueryEthProphecyParams) GetEthereumChainID() int64 {
	if m != nil {
		return m.EthereumChainID
	}
	return 0
}

func (m *QueryEthProphecyParams) GetBridgeContractAddress() *EthereumAddress {
	if m != nil {
		return m.BridgeContractAddress
	}
	return nil
}

func (m *QueryEthProphecyParams) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *QueryEthProphecyParams) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *QueryEthProphecyParams) GetTokenContractAddress() *EthereumAddress {
	if m != nil {
		return m.TokenContractAddress
	}
	return nil
}

func (m *QueryEthProphecyParams) GetEthereumSender() *EthereumAddress {
	if m != nil {
		return m.EthereumSender
	}
	return nil
}

type QueryEthProphecyResponse struct {
	ID                   string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status               *ProphecyStatus   `protobuf:"bytes,2,opt,name=Status" json:"Status,omitempty"`
	Claims               []*EthBridgeClaim `protobuf:"bytes,3,rep,name=Claims" json:"Claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryEthProphecyResponse) Reset()         { *m = QueryEthProphecyResponse{} }
func (m *QueryEthProphecyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEthProphecyResponse) ProtoMessage()    {}
func (*QueryEthProphecyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{8}
}
func (m *QueryEthProphecyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryEthProphecyResponse.Unmarshal(m, b)
}
func (m *QueryEthProphecyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryEthProphecyResponse.Marshal(b, m, deterministic)
}
func (dst *QueryEthProphecyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEthProphecyResponse.Merge(dst, src)
}
func (m *QueryEthProphecyResponse) XXX_Size() int {
	return xxx_messageInfo_QueryEthProphecyResponse.Size(m)
}
func (m *QueryEthProphecyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEthProphecyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEthProphecyResponse proto.InternalMessageInfo

func (m *QueryEthProphecyResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *QueryEthProphecyResponse) GetStatus() *ProphecyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryEthProphecyResponse) GetClaims() []*EthBridgeClaim {
	if m != nil {
		return m.Claims
	}
	return nil
}

type ProphecyStatus struct {
	Text                 EthBridgeStatus `protobuf:"varint,1,opt,name=Text,enum=types.EthBridgeStatus" json:"Text,omitempty"`
	FinalClaim           string          `protobuf:"bytes,2,opt,name=FinalClaim" json:"FinalClaim,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProphecyStatus) Reset()         { *m = ProphecyStatus{} }
func (m *ProphecyStatus) String() string { return proto.CompactTextString(m) }
func (*ProphecyStatus) ProtoMessage()    {}
func (*ProphecyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethbridge_506190bf2b0f36e7, []int{9}
}
func (m *ProphecyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProphecyStatus.Unmarshal(m, b)
}
func (m *ProphecyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProphecyStatus.Marshal(b, m, deterministic)
}
func (dst *ProphecyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProphecyStatus.Merge(dst, src)
}
func (m *ProphecyStatus) XXX_Size() int {
	return xxx_messageInfo_ProphecyStatus.Size(m)
}
func (m *ProphecyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ProphecyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ProphecyStatus proto.InternalMessageInfo

func (m *ProphecyStatus) GetText() EthBridgeStatus {
	if m != nil {
		return m.Text
	}
	return EthBridgeStatus_PendingStatusText
}

func (m *ProphecyStatus) GetFinalClaim() string {
	if m != nil {
		return m.FinalClaim
	}
	return ""
}

func init() {
	proto.RegisterType((*Eth2Chain33Action)(nil), "types.Eth2Chain33Action")
	proto.RegisterType((*EthBridgeClaim)(nil), "types.EthBridgeClaim")
	proto.RegisterType((*EthereumAddress)(nil), "types.EthereumAddress")
	proto.RegisterType((*Chain33Address)(nil), "types.Chain33Address")
	proto.RegisterType((*OracleClaimContent)(nil), "types.OracleClaimContent")
	proto.RegisterType((*MsgBurn)(nil), "types.MsgBurn")
	proto.RegisterType((*MsgLock)(nil), "types.MsgLock")
	proto.RegisterType((*QueryEthProphecyParams)(nil), "types.QueryEthProphecyParams")
	proto.RegisterType((*QueryEthProphecyResponse)(nil), "types.QueryEthProphecyResponse")
	proto.RegisterType((*ProphecyStatus)(nil), "types.ProphecyStatus")
	proto.RegisterEnum("types.EthBridgeStatus", EthBridgeStatus_name, EthBridgeStatus_value)
}

func init() { proto.RegisterFile("ethbridge.proto", fileDescriptor_ethbridge_506190bf2b0f36e7) }

var fileDescriptor_ethbridge_506190bf2b0f36e7 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xed, 0x7c, 0x90, 0x81, 0xba, 0xe9, 0xaa, 0xad, 0x7c, 0x40, 0x28, 0xca, 0x29, 0x0a,
	0xa2, 0x87, 0xf6, 0x08, 0xa2, 0x4a, 0xda, 0x54, 0x2d, 0x2a, 0x50, 0x36, 0x85, 0x13, 0x17, 0xd7,
	0x1e, 0xd5, 0x56, 0x9d, 0xdd, 0x68, 0xbd, 0xae, 0xf0, 0x3f, 0x40, 0xf0, 0xbb, 0xb8, 0xf0, 0x9b,
	0x38, 0x20, 0xaf, 0xd7, 0xc4, 0x76, 0x5b, 0xa3, 0x9e, 0xe1, 0x96, 0x1d, 0xbf, 0x37, 0x3b, 0xef,
	0xed, 0xb3, 0x37, 0xb0, 0x81, 0x32, 0xb8, 0x14, 0xa1, 0x7f, 0x85, 0xbb, 0x4b, 0xc1, 0x25, 0x27,
	0x6d, 0x99, 0x2e, 0x31, 0x1e, 0xfe, 0x30, 0x60, 0x73, 0x26, 0x83, 0xbd, 0xc3, 0xc0, 0x0d, 0xd9,
	0xfe, 0xfe, 0xc4, 0x93, 0x21, 0x67, 0xe4, 0x00, 0x6c, 0x94, 0xc1, 0x54, 0xe1, 0x0f, 0x23, 0x37,
	0x5c, 0x38, 0xc6, 0xc0, 0x18, 0x3d, 0xde, 0xdb, 0xde, 0x55, 0xac, 0xdd, 0x59, 0xe5, 0xe1, 0xc9,
	0x1a, 0xad, 0xc1, 0xc9, 0x18, 0xba, 0x8b, 0xf8, 0x6a, 0x9a, 0x08, 0xe6, 0x98, 0x8a, 0x69, 0x6b,
	0xe6, 0xdb, 0xbc, 0x7a, 0xb2, 0x46, 0x0b, 0x80, 0xc6, 0x9e, 0x71, 0xef, 0xda, 0xb1, 0xea, 0xd8,
	0xac, 0xaa, 0xb1, 0xd9, 0x4f, 0x62, 0x83, 0x29, 0x53, 0x07, 0x06, 0xc6, 0xa8, 0x4d, 0x4d, 0x99,
	0x4e, 0xbb, 0xd0, 0xbe, 0x71, 0xa3, 0x04, 0x87, 0xbf, 0x2c, 0xb0, 0xab, 0x53, 0x91, 0x11, 0x6c,
	0xcc, 0x64, 0x80, 0x02, 0x93, 0x85, 0x52, 0x77, 0x7a, 0xa4, 0x54, 0x58, 0xb4, 0x5e, 0x26, 0x67,
	0xb0, 0xad, 0x89, 0x9c, 0x49, 0xe1, 0x7a, 0x72, 0xe2, 0xfb, 0x02, 0xe3, 0x58, 0xcf, 0xbe, 0xb3,
	0x52, 0xad, 0x68, 0xfa, 0x29, 0xbd, 0x9b, 0x44, 0xb6, 0xa0, 0xfd, 0x8e, 0x33, 0x0f, 0x95, 0x1a,
	0x8b, 0xe6, 0x0b, 0xb2, 0x03, 0x9d, 0x79, 0xba, 0xb8, 0xe4, 0x91, 0xd3, 0x1a, 0x18, 0xa3, 0x1e,
	0xd5, 0x2b, 0xf2, 0x06, 0xb6, 0x2e, 0xf8, 0x35, 0xb2, 0xfa, 0xd6, 0xed, 0xc6, 0xad, 0xef, 0xe4,
	0x90, 0xd7, 0xca, 0x03, 0x05, 0x9c, 0x23, 0xf3, 0x51, 0x38, 0x9d, 0xc6, 0x2e, 0x35, 0x34, 0x39,
	0x80, 0x0d, 0x9d, 0x03, 0x8a, 0x1e, 0x86, 0x37, 0x28, 0x9c, 0x6e, 0xe5, 0xdc, 0x8b, 0x94, 0x68,
	0x7e, 0x1d, 0x4d, 0x26, 0xd0, 0xff, 0xe4, 0x46, 0xa1, 0xef, 0x4a, 0x2e, 0x0a, 0x21, 0x8f, 0x9a,
	0x3a, 0xdc, 0x82, 0x67, 0x3e, 0x4d, 0x16, 0x3c, 0x61, 0xd2, 0xe9, 0x0d, 0x8c, 0x51, 0x8b, 0xea,
	0x15, 0x79, 0x0a, 0x3d, 0x75, 0xac, 0x17, 0xe9, 0x12, 0x55, 0x00, 0x2c, 0xba, 0x2a, 0x0c, 0x9f,
	0xaf, 0xce, 0xba, 0x68, 0xe4, 0x40, 0xb7, 0x18, 0x21, 0x3b, 0xf6, 0x27, 0xb4, 0x58, 0x0e, 0xc7,
	0x60, 0x57, 0xc7, 0x68, 0xc0, 0x7e, 0x37, 0x80, 0xbc, 0x17, 0xae, 0x17, 0xe5, 0xa1, 0xca, 0x1c,
	0x47, 0x26, 0xef, 0x72, 0xca, 0x78, 0x90, 0x53, 0x2b, 0x99, 0xe6, 0xfd, 0x32, 0xad, 0xba, 0xcc,
	0x6f, 0x26, 0x74, 0xf5, 0x1b, 0xf4, 0x80, 0x78, 0xbf, 0x82, 0xf5, 0x4a, 0x5c, 0xfe, 0x12, 0xeb,
	0x2a, 0x98, 0xbc, 0x84, 0x75, 0x3d, 0xbc, 0xce, 0x94, 0xd5, 0x24, 0xb4, 0x8a, 0x25, 0x53, 0xe8,
	0x17, 0xed, 0xff, 0x18, 0xd5, 0x6a, 0xdc, 0xfd, 0x16, 0xbe, 0x64, 0x55, 0xbb, 0x6c, 0x55, 0x61,
	0x86, 0xfa, 0x2e, 0xfc, 0xf3, 0x66, 0xfc, 0x34, 0x61, 0xe7, 0x43, 0x82, 0x22, 0x9d, 0xc9, 0xe0,
	0x5c, 0xf0, 0x65, 0x80, 0x5e, 0x7a, 0xee, 0x0a, 0x77, 0x11, 0xff, 0xff, 0x0e, 0x3e, 0xf4, 0x3b,
	0x38, 0xfc, 0x6a, 0x80, 0x53, 0x37, 0x93, 0x62, 0xbc, 0xe4, 0x2c, 0xc6, 0xec, 0x0a, 0xd2, 0x0e,
	0xf6, 0xa8, 0x79, 0x7a, 0x44, 0x5e, 0x40, 0x67, 0x2e, 0x5d, 0x99, 0x14, 0x2e, 0x15, 0x59, 0x28,
	0x88, 0xf9, 0x43, 0xaa, 0x41, 0x19, 0x5c, 0xbd, 0xcf, 0xb1, 0x63, 0x0d, 0xac, 0x7b, 0xaf, 0x54,
	0xaa, 0x41, 0xc3, 0xcf, 0x60, 0x57, 0x1b, 0x91, 0x31, 0xb4, 0x2e, 0xf0, 0x8b, 0x54, 0x13, 0xd8,
	0x65, 0x49, 0x39, 0x5d, 0x6f, 0xa7, 0x30, 0xe4, 0x19, 0xc0, 0x71, 0xc8, 0xdc, 0x28, 0xbf, 0xc3,
	0x4d, 0x35, 0x73, 0xa9, 0x32, 0xfe, 0xa8, 0xa2, 0x51, 0x26, 0x92, 0x6d, 0xd8, 0x3c, 0x47, 0xe6,
	0x87, 0xec, 0x2a, 0x2f, 0x64, 0x7d, 0xfa, 0x6b, 0x59, 0x79, 0x9e, 0x78, 0x1e, 0xc6, 0x71, 0xa9,
	0x6c, 0x90, 0x2d, 0xe8, 0x1f, 0xbb, 0x61, 0x84, 0x7e, 0xa9, 0x6a, 0x5e, 0x76, 0xd4, 0x5f, 0x8c,
	0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x31, 0x2e, 0x41, 0x75, 0x08, 0x00, 0x00,
}
