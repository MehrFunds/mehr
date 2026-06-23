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

// ── FeederDelegation ─────────────────────────────────────────────────────────

// FeederDelegation stores the delegator → feeder relationship.
type FeederDelegation struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Feeder    string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
}

func (m *FeederDelegation) Reset()         { *m = FeederDelegation{} }
func (m *FeederDelegation) String() string { return proto.CompactTextString(m) }
func (*FeederDelegation) ProtoMessage()    {}

func (m *FeederDelegation) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *FeederDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeederDelegation.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FeederDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeederDelegation.Merge(m, src)
}
func (m *FeederDelegation) XXX_Size() int {
	return m.Size()
}
func (m *FeederDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeederDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_FeederDelegation proto.InternalMessageInfo

func (m *FeederDelegation) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *FeederDelegation) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

// ── MsgDelegateFeeder ─────────────────────────────────────────────────────────

// MsgDelegateFeeder allows a delegator to authorise a feeder address.
type MsgDelegateFeeder struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Feeder    string `protobuf:"bytes,2,opt,name=feeder,proto3" json:"feeder,omitempty"`
}

func (m *MsgDelegateFeeder) Reset()         { *m = MsgDelegateFeeder{} }
func (m *MsgDelegateFeeder) String() string { return proto.CompactTextString(m) }
func (*MsgDelegateFeeder) ProtoMessage()    {}

func (m *MsgDelegateFeeder) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgDelegateFeeder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDelegateFeeder.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgDelegateFeeder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDelegateFeeder.Merge(m, src)
}
func (m *MsgDelegateFeeder) XXX_Size() int {
	return m.Size()
}
func (m *MsgDelegateFeeder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDelegateFeeder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDelegateFeeder proto.InternalMessageInfo

func (m *MsgDelegateFeeder) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *MsgDelegateFeeder) GetFeeder() string {
	if m != nil {
		return m.Feeder
	}
	return ""
}

// ValidateBasic checks that Delegator and Feeder are non-empty valid bech32 addresses.
func (m *MsgDelegateFeeder) ValidateBasic() error {
	if m.Delegator == "" {
		return fmt.Errorf("delegator cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Delegator); err != nil {
		return fmt.Errorf("invalid delegator address: %w", err)
	}
	if m.Feeder == "" {
		return fmt.Errorf("feeder cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Feeder); err != nil {
		return fmt.Errorf("invalid feeder address: %w", err)
	}
	return nil
}

// GetSigners returns the signer addresses for the message (the delegator).
func (m *MsgDelegateFeeder) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ── MsgDelegateFeederResponse ─────────────────────────────────────────────────

// MsgDelegateFeederResponse is the response type for MsgDelegateFeeder.
type MsgDelegateFeederResponse struct{}

func (m *MsgDelegateFeederResponse) Reset()         { *m = MsgDelegateFeederResponse{} }
func (m *MsgDelegateFeederResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDelegateFeederResponse) ProtoMessage()    {}

func (m *MsgDelegateFeederResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgDelegateFeederResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDelegateFeederResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgDelegateFeederResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDelegateFeederResponse.Merge(m, src)
}
func (m *MsgDelegateFeederResponse) XXX_Size() int       { return m.Size() }
func (m *MsgDelegateFeederResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_MsgDelegateFeederResponse proto.InternalMessageInfo

// ── MsgRevokeDelegation ───────────────────────────────────────────────────────

// MsgRevokeDelegation allows a delegator to revoke a previously granted feeder delegation.
type MsgRevokeDelegation struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *MsgRevokeDelegation) Reset()         { *m = MsgRevokeDelegation{} }
func (m *MsgRevokeDelegation) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeDelegation) ProtoMessage()    {}

func (m *MsgRevokeDelegation) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgRevokeDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeDelegation.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgRevokeDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeDelegation.Merge(m, src)
}
func (m *MsgRevokeDelegation) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeDelegation proto.InternalMessageInfo

func (m *MsgRevokeDelegation) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

// ValidateBasic checks that the Delegator field is non-empty.
func (m *MsgRevokeDelegation) ValidateBasic() error {
	if m.Delegator == "" {
		return fmt.Errorf("delegator cannot be empty")
	}
	return nil
}

// GetSigners returns the signer addresses for the message (the delegator).
func (m *MsgRevokeDelegation) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ── MsgRevokeDelegationResponse ───────────────────────────────────────────────

// MsgRevokeDelegationResponse is the response type for MsgRevokeDelegation.
type MsgRevokeDelegationResponse struct{}

func (m *MsgRevokeDelegationResponse) Reset()         { *m = MsgRevokeDelegationResponse{} }
func (m *MsgRevokeDelegationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeDelegationResponse) ProtoMessage()    {}

func (m *MsgRevokeDelegationResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *MsgRevokeDelegationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeDelegationResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgRevokeDelegationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeDelegationResponse.Merge(m, src)
}
func (m *MsgRevokeDelegationResponse) XXX_Size() int       { return m.Size() }
func (m *MsgRevokeDelegationResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_MsgRevokeDelegationResponse proto.InternalMessageInfo

// ── init ──────────────────────────────────────────────────────────────────────

func init() {
	proto.RegisterType((*FeederDelegation)(nil), "mehr.v1.FeederDelegation")
	proto.RegisterType((*MsgDelegateFeeder)(nil), "mehr.v1.MsgDelegateFeeder")
	proto.RegisterType((*MsgDelegateFeederResponse)(nil), "mehr.v1.MsgDelegateFeederResponse")
	proto.RegisterType((*MsgRevokeDelegation)(nil), "mehr.v1.MsgRevokeDelegation")
	proto.RegisterType((*MsgRevokeDelegationResponse)(nil), "mehr.v1.MsgRevokeDelegationResponse")
}

// ── FeederDelegation marshal/unmarshal ────────────────────────────────────────

func (m *FeederDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeederDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeederDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 2: Feeder, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintFeederDelegation(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Delegator, string, wire type 2: (1<<3)|2 = 10 = 0x0a
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintFeederDelegation(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeederDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovFeederDelegation(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovFeederDelegation(uint64(l))
	}
	return n
}

func (m *FeederDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeederDelegation
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
			return fmt.Errorf("proto: FeederDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeederDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeederDelegation
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
				return ErrInvalidLengthFeederDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeederDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeederDelegation
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
				return ErrInvalidLengthFeederDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeederDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeederDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeederDelegation
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

// ── MsgDelegateFeeder marshal/unmarshal ───────────────────────────────────────

func (m *MsgDelegateFeeder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDelegateFeeder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDelegateFeeder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 2: Feeder, string, wire type 2: (2<<3)|2 = 18 = 0x12
	if len(m.Feeder) > 0 {
		i -= len(m.Feeder)
		copy(dAtA[i:], m.Feeder)
		i = encodeVarintFeederDelegation(dAtA, i, uint64(len(m.Feeder)))
		i--
		dAtA[i] = 0x12
	}
	// field 1: Delegator, string, wire type 2: (1<<3)|2 = 10 = 0x0a
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintFeederDelegation(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDelegateFeeder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovFeederDelegation(uint64(l))
	}
	l = len(m.Feeder)
	if l > 0 {
		n += 1 + l + sovFeederDelegation(uint64(l))
	}
	return n
}

func (m *MsgDelegateFeeder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeederDelegation
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
			return fmt.Errorf("proto: MsgDelegateFeeder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDelegateFeeder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeederDelegation
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
				return ErrInvalidLengthFeederDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeederDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeederDelegation
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
				return ErrInvalidLengthFeederDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeederDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeederDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeederDelegation
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

// ── MsgDelegateFeederResponse marshal/unmarshal ───────────────────────────────

func (m *MsgDelegateFeederResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDelegateFeederResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDelegateFeederResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}

func (m *MsgDelegateFeederResponse) Size() (n int) { return 0 }

func (m *MsgDelegateFeederResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeederDelegation
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
			return fmt.Errorf("proto: MsgDelegateFeederResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDelegateFeederResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFeederDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeederDelegation
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

// ── MsgRevokeDelegation marshal/unmarshal ─────────────────────────────────────

func (m *MsgRevokeDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	// field 1: Delegator, string, wire type 2: (1<<3)|2 = 10 = 0x0a
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintFeederDelegation(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovFeederDelegation(uint64(l))
	}
	return n
}

func (m *MsgRevokeDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeederDelegation
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
			return fmt.Errorf("proto: MsgRevokeDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeederDelegation
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
				return ErrInvalidLengthFeederDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeederDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeederDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeederDelegation
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

// ── MsgRevokeDelegationResponse marshal/unmarshal ─────────────────────────────

func (m *MsgRevokeDelegationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeDelegationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeDelegationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}

func (m *MsgRevokeDelegationResponse) Size() (n int) { return 0 }

func (m *MsgRevokeDelegationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeederDelegation
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
			return fmt.Errorf("proto: MsgRevokeDelegationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeDelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFeederDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeederDelegation
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

func encodeVarintFeederDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeederDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovFeederDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeederDelegation(x uint64) (n int) {
	return sovFeederDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func skipFeederDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeederDelegation
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
					return 0, ErrIntOverflowFeederDelegation
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
					return 0, ErrIntOverflowFeederDelegation
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
				return 0, ErrInvalidLengthFeederDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeederDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeederDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeederDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeederDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeederDelegation = fmt.Errorf("proto: unexpected end of group")
)
