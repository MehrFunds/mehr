package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MsgSubmitEvent is submitted by a feeder to record a detected on-chain transfer.
type MsgSubmitEvent struct {
	Feeder      string `protobuf:"bytes,1,opt,name=feeder,proto3" json:"feeder,omitempty"`
	Network     string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	TxHash      string `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	BlockNumber int64  `protobuf:"varint,5,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	LogIndex    int32  `protobuf:"varint,6,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	Asset       string `protobuf:"bytes,7,opt,name=asset,proto3" json:"asset,omitempty"`
	Contract    string `protobuf:"bytes,8,opt,name=contract,proto3" json:"contract,omitempty"`
	FromAddress string `protobuf:"bytes,9,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	AmountRaw   string `protobuf:"bytes,10,opt,name=amount_raw,json=amountRaw,proto3" json:"amount_raw,omitempty"`
	Decimals    int32  `protobuf:"varint,11,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Metadata    string `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *MsgSubmitEvent) Reset()         { *m = MsgSubmitEvent{} }
func (m *MsgSubmitEvent) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEvent) ProtoMessage()    {}

func (m *MsgSubmitEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEvent.Merge(m, src)
}
func (m *MsgSubmitEvent) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEvent proto.InternalMessageInfo

func (m *MsgSubmitEvent) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *MsgSubmitEvent) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *MsgSubmitEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSubmitEvent) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *MsgSubmitEvent) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *MsgSubmitEvent) GetLogIndex() int32 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *MsgSubmitEvent) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *MsgSubmitEvent) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *MsgSubmitEvent) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSubmitEvent) GetAmountRaw() string {
	if m != nil {
		return m.AmountRaw
	}
	return ""
}

func (m *MsgSubmitEvent) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *MsgSubmitEvent) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// ValidateBasic checks that required fields are non-empty.
func (m *MsgSubmitEvent) ValidateBasic() error {
	if m.Feeder == "" {
		return fmt.Errorf("feeder cannot be empty")
	}
	if m.Network == "" {
		return fmt.Errorf("network cannot be empty")
	}
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if m.TxHash == "" {
		return fmt.Errorf("tx_hash cannot be empty")
	}
	if m.Asset == "" {
		return fmt.Errorf("asset cannot be empty")
	}
	if m.AmountRaw == "" {
		return fmt.Errorf("amount_raw cannot be empty")
	}
	return nil
}

// GetSigners returns the signer addresses for the message (the feeder).
func (m *MsgSubmitEvent) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Feeder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgSubmitEventResponse is the response type for MsgSubmitEvent.
type MsgSubmitEventResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (m *MsgSubmitEventResponse) Reset()         { *m = MsgSubmitEventResponse{} }
func (m *MsgSubmitEventResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEventResponse) ProtoMessage()    {}

func (m *MsgSubmitEventResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEventResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEventResponse.Merge(m, src)
}
func (m *MsgSubmitEventResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEventResponse proto.InternalMessageInfo

func (m *MsgSubmitEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSubmitEvent)(nil), "mehr.v1.MsgSubmitEvent")
	proto.RegisterType((*MsgSubmitEventResponse)(nil), "mehr.v1.MsgSubmitEventResponse")
}

// ---------------------------------------------------------------------------
// MsgSubmitEvent marshal/unmarshal
// ---------------------------------------------------------------------------

func (m *MsgSubmitEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 12: Metadata, string, wire type 2: (12<<3)|2 = 98 = 0x62
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x62
	}
	// field 11: Decimals, int32, wire type 0: (11<<3)|0 = 88 = 0x58
	if m.Decimals != 0 {
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x58
	}
	// field 10: AmountRaw, string, wire type 2: (10<<3)|2 = 82 = 0x52
	if len(m.AmountRaw) > 0 {
		i -= len(m.AmountRaw)
		copy(dAtA[i:], m.AmountRaw)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.AmountRaw)))
		i--
		dAtA[i] = 0x52
	}
	// field 9: FromAddress, string, wire type 2: (9<<3)|2 = 74 = 0x4a
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x4a
	}
	// field 8: Contract, string, wire type 2: (8<<3)|2 = 66 = 0x42
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x42
	}
	// field 7: Asset, string, wire type 2: (7<<3)|2 = 58 = 0x3a
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x3a
	}
	// field 6: LogIndex, int32, wire type 0: (6<<3)|0 = 48 = 0x30
	if m.LogIndex != 0 {
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x30
	}
	// field 5: BlockNumber, int64, wire type 0: (5<<3)|0 = 40 = 0x28
	if m.BlockNumber != 0 {
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x28
	}
	// field 4: TxHash, string, wire type 2: (4<<3)|2 = 34 = 0x22
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x22
	}
	// field 3: Address, string, wire type 2: (3<<3)|2 = 26 = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	// field 2: Network, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Feeder, string, wire type 2: (1<<3)|2 = 10 = 0xa
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovMsgSubmitEvent(uint64(m.BlockNumber))
	}
	if m.LogIndex != 0 {
		n += 1 + sovMsgSubmitEvent(uint64(m.LogIndex))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	l = len(m.AmountRaw)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovMsgSubmitEvent(uint64(m.Decimals))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	return n
}

func (m *MsgSubmitEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSubmitEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountRaw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountRaw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSubmitEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

// ---------------------------------------------------------------------------
// MsgSubmitEventResponse marshal/unmarshal
// ---------------------------------------------------------------------------

func (m *MsgSubmitEventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEventResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEventResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 1: Event, message, wire type 2: (1<<3)|2 = 10 = 0xa
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgSubmitEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitEventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovMsgSubmitEvent(uint64(l))
	}
	return n
}

func (m *MsgSubmitEventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSubmitEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitEventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEventResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &Event{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSubmitEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgSubmitEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func encodeVarintMsgSubmitEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSubmitEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovMsgSubmitEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSubmitEvent(x uint64) (n int) {
	return sovMsgSubmitEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func skipMsgSubmitEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSubmitEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgSubmitEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsgSubmitEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSubmitEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSubmitEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSubmitEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSubmitEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSubmitEvent = fmt.Errorf("proto: unexpected end of group")
)
