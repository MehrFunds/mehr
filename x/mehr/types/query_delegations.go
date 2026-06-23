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

// ── QueryFeederDelegationRequest ──────────────────────────────────────────────

type QueryFeederDelegationRequest struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *QueryFeederDelegationRequest) Reset()         { *m = QueryFeederDelegationRequest{} }
func (m *QueryFeederDelegationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeederDelegationRequest) ProtoMessage()    {}

func (m *QueryFeederDelegationRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFeederDelegationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeederDelegationRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFeederDelegationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeederDelegationRequest.Merge(m, src)
}
func (m *QueryFeederDelegationRequest) XXX_Size() int       { return m.Size() }
func (m *QueryFeederDelegationRequest) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryFeederDelegationRequest proto.InternalMessageInfo

func (m *QueryFeederDelegationRequest) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

// ── QueryFeederDelegationResponse ─────────────────────────────────────────────

type QueryFeederDelegationResponse struct {
	Delegation *FeederDelegation `protobuf:"bytes,1,opt,name=delegation,proto3" json:"delegation,omitempty"`
}

func (m *QueryFeederDelegationResponse) Reset()         { *m = QueryFeederDelegationResponse{} }
func (m *QueryFeederDelegationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeederDelegationResponse) ProtoMessage()    {}

func (m *QueryFeederDelegationResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFeederDelegationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeederDelegationResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFeederDelegationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeederDelegationResponse.Merge(m, src)
}
func (m *QueryFeederDelegationResponse) XXX_Size() int       { return m.Size() }
func (m *QueryFeederDelegationResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryFeederDelegationResponse proto.InternalMessageInfo

func (m *QueryFeederDelegationResponse) GetDelegation() *FeederDelegation {
	if m != nil {
		return m.Delegation
	}
	return nil
}

// ── QueryAllDelegationsRequest ────────────────────────────────────────────────

type QueryAllDelegationsRequest struct{}

func (m *QueryAllDelegationsRequest) Reset()         { *m = QueryAllDelegationsRequest{} }
func (m *QueryAllDelegationsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllDelegationsRequest) ProtoMessage()    {}

func (m *QueryAllDelegationsRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllDelegationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDelegationsRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllDelegationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDelegationsRequest.Merge(m, src)
}
func (m *QueryAllDelegationsRequest) XXX_Size() int       { return m.Size() }
func (m *QueryAllDelegationsRequest) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryAllDelegationsRequest proto.InternalMessageInfo

// ── QueryAllDelegationsResponse ───────────────────────────────────────────────

type QueryAllDelegationsResponse struct {
	Delegations []*FeederDelegation `protobuf:"bytes,1,rep,name=delegations,proto3" json:"delegations,omitempty"`
}

func (m *QueryAllDelegationsResponse) Reset()         { *m = QueryAllDelegationsResponse{} }
func (m *QueryAllDelegationsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllDelegationsResponse) ProtoMessage()    {}

func (m *QueryAllDelegationsResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllDelegationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllDelegationsResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllDelegationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllDelegationsResponse.Merge(m, src)
}
func (m *QueryAllDelegationsResponse) XXX_Size() int       { return m.Size() }
func (m *QueryAllDelegationsResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryAllDelegationsResponse proto.InternalMessageInfo

func (m *QueryAllDelegationsResponse) GetDelegations() []*FeederDelegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

// ── init ──────────────────────────────────────────────────────────────────────

func init() {
	proto.RegisterType((*QueryFeederDelegationRequest)(nil), "mehr.v1.QueryFeederDelegationRequest")
	proto.RegisterType((*QueryFeederDelegationResponse)(nil), "mehr.v1.QueryFeederDelegationResponse")
	proto.RegisterType((*QueryAllDelegationsRequest)(nil), "mehr.v1.QueryAllDelegationsRequest")
	proto.RegisterType((*QueryAllDelegationsResponse)(nil), "mehr.v1.QueryAllDelegationsResponse")
}

// ── QueryFeederDelegationRequest marshal/unmarshal ────────────────────────────

func (m *QueryFeederDelegationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeederDelegationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeederDelegationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	// field 1: Delegator, string, wire type 2: (1<<3)|2 = 10 = 0x0a
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintQueryDelegations(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeederDelegationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	l := len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovQueryDelegations(uint64(l))
	}
	return n
}

func (m *QueryFeederDelegationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDelegations
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
			return fmt.Errorf("proto: QueryFeederDelegationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeederDelegationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDelegations
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
				return ErrInvalidLengthQueryDelegations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDelegations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDelegations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDelegations
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

// ── QueryFeederDelegationResponse marshal/unmarshal ───────────────────────────

func (m *QueryFeederDelegationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeederDelegationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeederDelegationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	// field 1: Delegation, message, wire type 2: (1<<3)|2 = 10 = 0x0a
	if m.Delegation != nil {
		size, err := m.Delegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryDelegations(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeederDelegationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	if m.Delegation != nil {
		l := m.Delegation.Size()
		n += 1 + l + sovQueryDelegations(uint64(l))
	}
	return n
}

func (m *QueryFeederDelegationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDelegations
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
			return fmt.Errorf("proto: QueryFeederDelegationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeederDelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDelegations
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
				return ErrInvalidLengthQueryDelegations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDelegations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delegation == nil {
				m.Delegation = &FeederDelegation{}
			}
			if err := m.Delegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDelegations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDelegations
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

// ── QueryAllDelegationsRequest marshal/unmarshal ──────────────────────────────

func (m *QueryAllDelegationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDelegationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDelegationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}

func (m *QueryAllDelegationsRequest) Size() (n int) { return 0 }

func (m *QueryAllDelegationsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDelegations
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
			return fmt.Errorf("proto: QueryAllDelegationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDelegations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDelegations
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

// ── QueryAllDelegationsResponse marshal/unmarshal ─────────────────────────────

func (m *QueryAllDelegationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllDelegationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllDelegationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	// field 1: Delegations, repeated message, wire type 2: (1<<3)|2 = 10 = 0x0a
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryDelegations(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllDelegationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	for _, e := range m.Delegations {
		l := e.Size()
		n += 1 + sovQueryDelegations(uint64(l)) + l
	}
	return n
}

func (m *QueryAllDelegationsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDelegations
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
			return fmt.Errorf("proto: QueryAllDelegationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllDelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDelegations
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
				return ErrInvalidLengthQueryDelegations
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDelegations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, &FeederDelegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDelegations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDelegations
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

func encodeVarintQueryDelegations(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryDelegations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovQueryDelegations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryDelegations(x uint64) (n int) {
	return sovQueryDelegations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func skipQueryDelegations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryDelegations
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
					return 0, ErrIntOverflowQueryDelegations
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
					return 0, ErrIntOverflowQueryDelegations
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
				return 0, ErrInvalidLengthQueryDelegations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryDelegations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryDelegations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryDelegations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryDelegations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryDelegations = fmt.Errorf("proto: unexpected end of group")
)
