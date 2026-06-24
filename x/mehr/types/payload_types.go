package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ── EvmTransferPayload ────────────────────────────────────────────────────────

// EvmTransferPayload is the structured payload for EVM transfer events.
type EvmTransferPayload struct {
	Network     string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	TxHash      string `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	BlockNumber int64  `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	LogIndex    int32  `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	Asset       string `protobuf:"bytes,6,opt,name=asset,proto3" json:"asset,omitempty"`
	Contract    string `protobuf:"bytes,7,opt,name=contract,proto3" json:"contract,omitempty"`
	FromAddress string `protobuf:"bytes,8,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	AmountRaw   string `protobuf:"bytes,9,opt,name=amount_raw,json=amountRaw,proto3" json:"amount_raw,omitempty"`
	Decimals    int32  `protobuf:"varint,10,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Metadata    string `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *EvmTransferPayload) Reset()         { *m = EvmTransferPayload{} }
func (m *EvmTransferPayload) String() string { return proto.CompactTextString(m) }
func (*EvmTransferPayload) ProtoMessage()    {}

func (m *EvmTransferPayload) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *EvmTransferPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvmTransferPayload.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EvmTransferPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvmTransferPayload.Merge(m, src)
}
func (m *EvmTransferPayload) XXX_Size() int {
	return m.Size()
}
func (m *EvmTransferPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EvmTransferPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EvmTransferPayload proto.InternalMessageInfo

func (m *EvmTransferPayload) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}
func (m *EvmTransferPayload) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}
func (m *EvmTransferPayload) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}
func (m *EvmTransferPayload) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}
func (m *EvmTransferPayload) GetLogIndex() int32 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}
func (m *EvmTransferPayload) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}
func (m *EvmTransferPayload) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}
func (m *EvmTransferPayload) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}
func (m *EvmTransferPayload) GetAmountRaw() string {
	if m != nil {
		return m.AmountRaw
	}
	return ""
}
func (m *EvmTransferPayload) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}
func (m *EvmTransferPayload) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// ── PriceDataPayload ──────────────────────────────────────────────────────────

// PriceDataPayload is the structured payload for price data feeds.
type PriceDataPayload struct {
	Pair      string `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Price     string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Source    string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
}

func (m *PriceDataPayload) Reset()         { *m = PriceDataPayload{} }
func (m *PriceDataPayload) String() string { return proto.CompactTextString(m) }
func (*PriceDataPayload) ProtoMessage()    {}

func (m *PriceDataPayload) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *PriceDataPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceDataPayload.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PriceDataPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceDataPayload.Merge(m, src)
}
func (m *PriceDataPayload) XXX_Size() int {
	return m.Size()
}
func (m *PriceDataPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceDataPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PriceDataPayload proto.InternalMessageInfo

func (m *PriceDataPayload) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}
func (m *PriceDataPayload) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}
func (m *PriceDataPayload) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}
func (m *PriceDataPayload) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// ── init ──────────────────────────────────────────────────────────────────────

func init() {
	proto.RegisterType((*EvmTransferPayload)(nil), "mehr.v1.EvmTransferPayload")
	proto.RegisterType((*PriceDataPayload)(nil), "mehr.v1.PriceDataPayload")
}

// ── EvmTransferPayload marshal/unmarshal ──────────────────────────────────────

func (m *EvmTransferPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvmTransferPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvmTransferPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 11: Metadata, string, wire type 2: (11<<3)|2 = 90 = 0x5a
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x5a
	}
	// field 10: Decimals, int32, wire type 0: (10<<3)|0 = 80 = 0x50
	if m.Decimals != 0 {
		i = encodeVarintPayloadTypes(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x50
	}
	// field 9: AmountRaw, string, wire type 2: (9<<3)|2 = 74 = 0x4a
	if len(m.AmountRaw) > 0 {
		i -= len(m.AmountRaw)
		copy(dAtA[i:], m.AmountRaw)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.AmountRaw)))
		i--
		dAtA[i] = 0x4a
	}
	// field 8: FromAddress, string, wire type 2: (8<<3)|2 = 66 = 0x42
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x42
	}
	// field 7: Contract, string, wire type 2: (7<<3)|2 = 58 = 0x3a
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x3a
	}
	// field 6: Asset, string, wire type 2: (6<<3)|2 = 50 = 0x32
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x32
	}
	// field 5: LogIndex, int32, wire type 0: (5<<3)|0 = 40 = 0x28
	if m.LogIndex != 0 {
		i = encodeVarintPayloadTypes(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x28
	}
	// field 4: BlockNumber, int64, wire type 0: (4<<3)|0 = 32 = 0x20
	if m.BlockNumber != 0 {
		i = encodeVarintPayloadTypes(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x20
	}
	// field 3: TxHash, string, wire type 2: (3<<3)|2 = 26 = 0x1a
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x1a
	}
	// field 2: Address, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Network, string, wire type 2: (1<<3)|2 = 10 = 0xa
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EvmTransferPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovPayloadTypes(uint64(m.BlockNumber))
	}
	if m.LogIndex != 0 {
		n += 1 + sovPayloadTypes(uint64(m.LogIndex))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.AmountRaw)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovPayloadTypes(uint64(m.Decimals))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	return n
}

func (m *EvmTransferPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayloadTypes
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
			return fmt.Errorf("proto: EvmTransferPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvmTransferPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountRaw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountRaw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayloadTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayloadTypes
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

// ── PriceDataPayload marshal/unmarshal ────────────────────────────────────────

func (m *PriceDataPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceDataPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceDataPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 4: Source, string, wire type 2: (4<<3)|2 = 34 = 0x22
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x22
	}
	// field 3: Timestamp, int64, wire type 0: (3<<3)|0 = 24 = 0x18
	if m.Timestamp != 0 {
		i = encodeVarintPayloadTypes(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	// field 2: Price, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Pair, string, wire type 2: (1<<3)|2 = 10 = 0xa
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintPayloadTypes(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PriceDataPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPayloadTypes(uint64(m.Timestamp))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovPayloadTypes(uint64(l))
	}
	return n
}

func (m *PriceDataPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayloadTypes
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
			return fmt.Errorf("proto: PriceDataPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceDataPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloadTypes
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
				return ErrInvalidLengthPayloadTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayloadTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayloadTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayloadTypes
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

// ── helpers ───────────────────────────────────────────────────────────────────

func encodeVarintPayloadTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayloadTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovPayloadTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func skipPayloadTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayloadTypes
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
					return 0, ErrIntOverflowPayloadTypes
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
					return 0, ErrIntOverflowPayloadTypes
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
				return 0, ErrInvalidLengthPayloadTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayloadTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayloadTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayloadTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayloadTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayloadTypes = fmt.Errorf("proto: unexpected end of group")
)
