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

// ── Fact ──────────────────────────────────────────────────────────────────────

// Fact is a single data point submitted by a feeder to a feed.
type Fact struct {
	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeedId           string `protobuf:"bytes,2,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Feeder           string `protobuf:"bytes,3,opt,name=feeder,proto3" json:"feeder,omitempty"`
	Payload          []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	DeduplicationKey string `protobuf:"bytes,5,opt,name=deduplication_key,json=deduplicationKey,proto3" json:"deduplication_key,omitempty"`
	SubmittedAt      string `protobuf:"bytes,6,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
}

func (m *Fact) Reset()         { *m = Fact{} }
func (m *Fact) String() string { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()    {}

func (m *Fact) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *Fact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fact.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Fact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fact.Merge(m, src)
}
func (m *Fact) XXX_Size() int {
	return m.Size()
}
func (m *Fact) XXX_DiscardUnknown() {
	xxx_messageInfo_Fact.DiscardUnknown(m)
}

var xxx_messageInfo_Fact proto.InternalMessageInfo

func (m *Fact) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Fact) GetFeedId() string {
	if m != nil {
		return m.FeedId
	}
	return ""
}

func (m *Fact) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *Fact) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Fact) GetDeduplicationKey() string {
	if m != nil {
		return m.DeduplicationKey
	}
	return ""
}

func (m *Fact) GetSubmittedAt() string {
	if m != nil {
		return m.SubmittedAt
	}
	return ""
}

// ── MsgSubmitFact ─────────────────────────────────────────────────────────────

// MsgSubmitFact submits a fact to a feed.
type MsgSubmitFact struct {
	Feeder           string `protobuf:"bytes,1,opt,name=feeder,proto3" json:"feeder,omitempty"`
	FeedId           string `protobuf:"bytes,2,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Payload          []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	DeduplicationKey string `protobuf:"bytes,4,opt,name=deduplication_key,json=deduplicationKey,proto3" json:"deduplication_key,omitempty"`
}

func (m *MsgSubmitFact) Reset()         { *m = MsgSubmitFact{} }
func (m *MsgSubmitFact) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFact) ProtoMessage()    {}

func (m *MsgSubmitFact) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgSubmitFact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFact.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgSubmitFact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFact.Merge(m, src)
}
func (m *MsgSubmitFact) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFact) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFact.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFact proto.InternalMessageInfo

func (m *MsgSubmitFact) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

func (m *MsgSubmitFact) GetFeedId() string {
	if m != nil {
		return m.FeedId
	}
	return ""
}

func (m *MsgSubmitFact) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MsgSubmitFact) GetDeduplicationKey() string {
	if m != nil {
		return m.DeduplicationKey
	}
	return ""
}

// ValidateBasic checks that required fields are non-empty and valid.
func (m *MsgSubmitFact) ValidateBasic() error {
	if m.Feeder == "" {
		return fmt.Errorf("feeder cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Feeder); err != nil {
		return fmt.Errorf("invalid feeder address: %w", err)
	}
	if m.FeedId == "" {
		return fmt.Errorf("feed_id cannot be empty")
	}
	if len(m.Payload) == 0 {
		return fmt.Errorf("payload cannot be empty")
	}
	if m.DeduplicationKey == "" {
		return fmt.Errorf("deduplication_key cannot be empty")
	}
	return nil
}

// GetSigners returns the signer addresses for the message (the feeder).
func (m *MsgSubmitFact) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Feeder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ── MsgSubmitFactResponse ─────────────────────────────────────────────────────

// MsgSubmitFactResponse is the response type for MsgSubmitFact.
type MsgSubmitFactResponse struct {
	Fact *Fact `protobuf:"bytes,1,opt,name=fact,proto3" json:"fact,omitempty"`
}

func (m *MsgSubmitFactResponse) Reset()         { *m = MsgSubmitFactResponse{} }
func (m *MsgSubmitFactResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFactResponse) ProtoMessage()    {}

func (m *MsgSubmitFactResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgSubmitFactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFactResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgSubmitFactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFactResponse.Merge(m, src)
}
func (m *MsgSubmitFactResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFactResponse proto.InternalMessageInfo

func (m *MsgSubmitFactResponse) GetFact() *Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

// ── init ──────────────────────────────────────────────────────────────────────

func init() {
	proto.RegisterType((*Fact)(nil), "mehr.v1.Fact")
	proto.RegisterType((*MsgSubmitFact)(nil), "mehr.v1.MsgSubmitFact")
	proto.RegisterType((*MsgSubmitFactResponse)(nil), "mehr.v1.MsgSubmitFactResponse")
}

// ── Fact marshal/unmarshal ────────────────────────────────────────────────────

func (m *Fact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fact) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 6: SubmittedAt, string, wire type 2: (6<<3)|2 = 50 = 0x32
	if len(m.SubmittedAt) > 0 {
		i -= len(m.SubmittedAt)
		copy(dAtA[i:], m.SubmittedAt)
		i = encodeVarintFact(dAtA, i, uint64(len(m.SubmittedAt)))
		i--
		dAtA[i] = 0x32
	}
	// field 5: DeduplicationKey, string, wire type 2: (5<<3)|2 = 42 = 0x2a
	if len(m.DeduplicationKey) > 0 {
		i -= len(m.DeduplicationKey)
		copy(dAtA[i:], m.DeduplicationKey)
		i = encodeVarintFact(dAtA, i, uint64(len(m.DeduplicationKey)))
		i--
		dAtA[i] = 0x2a
	}
	// field 4: Payload, bytes, wire type 2: (4<<3)|2 = 34 = 0x22
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintFact(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	// field 3: Feeder, string, wire type 2: (3<<3)|2 = 26 = 0x1a
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintFact(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x1a
	}
	// field 2: FeedId, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.FeedId) > 0 {
		i -= len(m.FeedId)
		copy(dAtA[i:], m.FeedId)
		i = encodeVarintFact(dAtA, i, uint64(len(m.FeedId)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Id, varint, wire type 0: (1<<3)|0 = 8 = 0x8
	if m.Id != 0 {
		i = encodeVarintFact(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Fact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFact(uint64(m.Id))
	}
	l = len(m.FeedId)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.DeduplicationKey)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.SubmittedAt)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	return n
}

func (m *Fact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFact
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
			return fmt.Errorf("proto: Fact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return fmt.Errorf("proto: wrong wireType = %d for field FeedId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeedId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeduplicationKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeduplicationKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmittedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmittedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFact(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFact
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

// ── MsgSubmitFact marshal/unmarshal ───────────────────────────────────────────

func (m *MsgSubmitFact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFact) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 4: DeduplicationKey, string, wire type 2: (4<<3)|2 = 34 = 0x22
	if len(m.DeduplicationKey) > 0 {
		i -= len(m.DeduplicationKey)
		copy(dAtA[i:], m.DeduplicationKey)
		i = encodeVarintFact(dAtA, i, uint64(len(m.DeduplicationKey)))
		i--
		dAtA[i] = 0x22
	}
	// field 3: Payload, bytes, wire type 2: (3<<3)|2 = 26 = 0x1a
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintFact(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	// field 2: FeedId, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.FeedId) > 0 {
		i -= len(m.FeedId)
		copy(dAtA[i:], m.FeedId)
		i = encodeVarintFact(dAtA, i, uint64(len(m.FeedId)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Feeder, string, wire type 2: (1<<3)|2 = 10 = 0xa
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintFact(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitFact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.FeedId)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	l = len(m.DeduplicationKey)
	if l > 0 {
		n += 1 + l + sovFact(uint64(l))
	}
	return n
}

func (m *MsgSubmitFact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFact
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
			return fmt.Errorf("proto: MsgSubmitFact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeedId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeedId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeduplicationKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeduplicationKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFact(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFact
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

// ── MsgSubmitFactResponse marshal/unmarshal ───────────────────────────────────

func (m *MsgSubmitFactResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFactResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFactResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 1: Fact, message, wire type 2: (1<<3)|2 = 10 = 0xa
	if m.Fact != nil {
		{
			size, err := m.Fact.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFact(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitFactResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fact != nil {
		l = m.Fact.Size()
		n += 1 + l + sovFact(uint64(l))
	}
	return n
}

func (m *MsgSubmitFactResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFact
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
			return fmt.Errorf("proto: MsgSubmitFactResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFactResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fact", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFact
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
				return ErrInvalidLengthFact
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFact
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fact == nil {
				m.Fact = &Fact{}
			}
			if err := m.Fact.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFact(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFact
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

func encodeVarintFact(dAtA []byte, offset int, v uint64) int {
	offset -= sovFact(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovFact(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func skipFact(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFact
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
					return 0, ErrIntOverflowFact
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
					return 0, ErrIntOverflowFact
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
				return 0, ErrInvalidLengthFact
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFact
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFact
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFact        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFact          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFact = fmt.Errorf("proto: unexpected end of group")
)
