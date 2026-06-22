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

// Event records a detected on-chain token transfer for a watched address.
type Event struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Feeder      string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
	WatchId     uint64 `protobuf:"varint,3,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	Network     string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	TxHash      string `protobuf:"bytes,6,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	BlockNumber int64  `protobuf:"varint,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	LogIndex    int32  `protobuf:"varint,8,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	Asset       string `protobuf:"bytes,9,opt,name=asset,proto3" json:"asset,omitempty"`
	Contract    string `protobuf:"bytes,10,opt,name=contract,proto3" json:"contract,omitempty"`
	FromAddress string `protobuf:"bytes,11,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	AmountRaw   string `protobuf:"bytes,12,opt,name=amount_raw,json=amountRaw,proto3" json:"amount_raw,omitempty"`
	Decimals    int32  `protobuf:"varint,13,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Metadata    string `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`
	SubmittedAt string `protobuf:"bytes,15,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *Event) GetWatchId() uint64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

func (m *Event) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Event) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Event) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Event) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Event) GetLogIndex() int32 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *Event) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Event) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Event) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *Event) GetAmountRaw() string {
	if m != nil {
		return m.AmountRaw
	}
	return ""
}

func (m *Event) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Event) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Event) GetSubmittedAt() string {
	if m != nil {
		return m.SubmittedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "mehr.v1.Event")
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 15: SubmittedAt, string, wire type 2: (15<<3)|2 = 122 = 0x7a
	if len(m.SubmittedAt) > 0 {
		i -= len(m.SubmittedAt)
		copy(dAtA[i:], m.SubmittedAt)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.SubmittedAt)))
		i--
		dAtA[i] = 0x7a
	}
	// field 14: Metadata, string, wire type 2: (14<<3)|2 = 114 = 0x72
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x72
	}
	// field 13: Decimals, int32, wire type 0: (13<<3)|0 = 104 = 0x68
	if m.Decimals != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x68
	}
	// field 12: AmountRaw, string, wire type 2: (12<<3)|2 = 98 = 0x62
	if len(m.AmountRaw) > 0 {
		i -= len(m.AmountRaw)
		copy(dAtA[i:], m.AmountRaw)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.AmountRaw)))
		i--
		dAtA[i] = 0x62
	}
	// field 11: FromAddress, string, wire type 2: (11<<3)|2 = 90 = 0x5a
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x5a
	}
	// field 10: Contract, string, wire type 2: (10<<3)|2 = 82 = 0x52
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x52
	}
	// field 9: Asset, string, wire type 2: (9<<3)|2 = 74 = 0x4a
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x4a
	}
	// field 8: LogIndex, int32, wire type 0: (8<<3)|0 = 64 = 0x40
	if m.LogIndex != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x40
	}
	// field 7: BlockNumber, int64, wire type 0: (7<<3)|0 = 56 = 0x38
	if m.BlockNumber != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x38
	}
	// field 6: TxHash, string, wire type 2: (6<<3)|2 = 50 = 0x32
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x32
	}
	// field 5: Address, string, wire type 2: (5<<3)|2 = 42 = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	// field 4: Network, string, wire type 2: (4<<3)|2 = 34 = 0x22
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	// field 3: WatchId, uint64, wire type 0: (3<<3)|0 = 24 = 0x18
	if m.WatchId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.WatchId))
		i--
		dAtA[i] = 0x18
	}
	// field 2: Feeder, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Id, uint64, wire type 0: (1<<3)|0 = 8 = 0x8
	if m.Id != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	// field 1: Id
	if m.Id != 0 {
		n += 1 + sovEvent(uint64(m.Id))
	}
	// field 2: Feeder
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 3: WatchId
	if m.WatchId != 0 {
		n += 1 + sovEvent(uint64(m.WatchId))
	}
	// field 4: Network
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 5: Address
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 6: TxHash
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 7: BlockNumber
	if m.BlockNumber != 0 {
		n += 1 + sovEvent(uint64(m.BlockNumber))
	}
	// field 8: LogIndex
	if m.LogIndex != 0 {
		n += 1 + sovEvent(uint64(m.LogIndex))
	}
	// field 9: Asset
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 10: Contract
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 11: FromAddress
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 12: AmountRaw
	l = len(m.AmountRaw)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 13: Decimals
	if m.Decimals != 0 {
		n += 1 + sovEvent(uint64(m.Decimals))
	}
	// field 14: Metadata
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	// field 15: SubmittedAt
	l = len(m.SubmittedAt)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountRaw", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountRaw = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmittedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmittedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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

func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
