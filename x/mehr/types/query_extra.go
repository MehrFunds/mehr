package types

import (
	fmt "fmt"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ── QueryAllWatchesRequest ────────────────────────────────────────────────────

type QueryAllWatchesRequest struct{}

func (m *QueryAllWatchesRequest) Reset()         { *m = QueryAllWatchesRequest{} }
func (m *QueryAllWatchesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllWatchesRequest) ProtoMessage()    {}

func (m *QueryAllWatchesRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllWatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWatchesRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllWatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWatchesRequest.Merge(m, src)
}
func (m *QueryAllWatchesRequest) XXX_Size() int                { return m.Size() }
func (m *QueryAllWatchesRequest) XXX_DiscardUnknown()          {}

var xxx_messageInfo_QueryAllWatchesRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryAllWatchesRequest)(nil), "mehr.v1.QueryAllWatchesRequest")
}

func (m *QueryAllWatchesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWatchesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWatchesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}

func (m *QueryAllWatchesRequest) Size() (n int) { return 0 }

func (m *QueryAllWatchesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryAllWatchesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWatchesRequest: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryAllWatchesResponse ───────────────────────────────────────────────────

type QueryAllWatchesResponse struct {
	Watches []*Watch `protobuf:"bytes,1,rep,name=watches,proto3" json:"watches,omitempty"`
}

func (m *QueryAllWatchesResponse) Reset()         { *m = QueryAllWatchesResponse{} }
func (m *QueryAllWatchesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllWatchesResponse) ProtoMessage()    {}

func (m *QueryAllWatchesResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllWatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWatchesResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllWatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWatchesResponse.Merge(m, src)
}
func (m *QueryAllWatchesResponse) XXX_Size() int       { return m.Size() }
func (m *QueryAllWatchesResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryAllWatchesResponse proto.InternalMessageInfo

func (m *QueryAllWatchesResponse) GetWatches() []*Watch {
	if m != nil {
		return m.Watches
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllWatchesResponse)(nil), "mehr.v1.QueryAllWatchesResponse")
}

func (m *QueryAllWatchesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllWatchesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllWatchesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if len(m.Watches) > 0 {
		for iNdEx := len(m.Watches) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Watches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryExtra(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllWatchesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	for _, e := range m.Watches {
		l := e.Size()
		n += 1 + sovQueryExtra(uint64(l)) + l
	}
	return n
}

func (m *QueryAllWatchesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryAllWatchesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWatchesResponse: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
				return ErrInvalidLengthQueryExtra
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Watches = append(m.Watches, &Watch{})
			if err := m.Watches[len(m.Watches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryAllWebhooksRequest ───────────────────────────────────────────────────

type QueryAllWebhooksRequest struct{}

func (m *QueryAllWebhooksRequest) Reset()         { *m = QueryAllWebhooksRequest{} }
func (m *QueryAllWebhooksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllWebhooksRequest) ProtoMessage()    {}

func (m *QueryAllWebhooksRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllWebhooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWebhooksRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllWebhooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWebhooksRequest.Merge(m, src)
}
func (m *QueryAllWebhooksRequest) XXX_Size() int       { return m.Size() }
func (m *QueryAllWebhooksRequest) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryAllWebhooksRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryAllWebhooksRequest)(nil), "mehr.v1.QueryAllWebhooksRequest")
}

func (m *QueryAllWebhooksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllWebhooksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllWebhooksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}
func (m *QueryAllWebhooksRequest) Size() (n int) { return 0 }
func (m *QueryAllWebhooksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryAllWebhooksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWebhooksRequest: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryAllWebhooksResponse ──────────────────────────────────────────────────

type QueryAllWebhooksResponse struct {
	Webhooks []*Webhook `protobuf:"bytes,1,rep,name=webhooks,proto3" json:"webhooks,omitempty"`
}

func (m *QueryAllWebhooksResponse) Reset()         { *m = QueryAllWebhooksResponse{} }
func (m *QueryAllWebhooksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllWebhooksResponse) ProtoMessage()    {}

func (m *QueryAllWebhooksResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllWebhooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllWebhooksResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllWebhooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllWebhooksResponse.Merge(m, src)
}
func (m *QueryAllWebhooksResponse) XXX_Size() int       { return m.Size() }
func (m *QueryAllWebhooksResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryAllWebhooksResponse proto.InternalMessageInfo

func (m *QueryAllWebhooksResponse) GetWebhooks() []*Webhook {
	if m != nil {
		return m.Webhooks
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllWebhooksResponse)(nil), "mehr.v1.QueryAllWebhooksResponse")
}

func (m *QueryAllWebhooksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllWebhooksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllWebhooksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if len(m.Webhooks) > 0 {
		for iNdEx := len(m.Webhooks) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Webhooks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryExtra(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QueryAllWebhooksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	for _, e := range m.Webhooks {
		l := e.Size()
		n += 1 + sovQueryExtra(uint64(l)) + l
	}
	return n
}
func (m *QueryAllWebhooksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryAllWebhooksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllWebhooksResponse: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Webhooks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
				return ErrInvalidLengthQueryExtra
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Webhooks = append(m.Webhooks, &Webhook{})
			if err := m.Webhooks[len(m.Webhooks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryEventsRequest ────────────────────────────────────────────────────────

type QueryEventsRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryEventsRequest) Reset()         { *m = QueryEventsRequest{} }
func (m *QueryEventsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEventsRequest) ProtoMessage()    {}

func (m *QueryEventsRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEventsRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEventsRequest.Merge(m, src)
}
func (m *QueryEventsRequest) XXX_Size() int       { return m.Size() }
func (m *QueryEventsRequest) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryEventsRequest proto.InternalMessageInfo

func (m *QueryEventsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryEventsRequest)(nil), "mehr.v1.QueryEventsRequest")
}

func (m *QueryEventsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryEventsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryEventsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQueryExtra(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryEventsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	l := len(m.Address)
	if l > 0 {
		n += 1 + sovQueryExtra(uint64(l)) + l
	}
	return n
}
func (m *QueryEventsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryEventsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEventsRequest: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
				return ErrInvalidLengthQueryExtra
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryEventsResponse ───────────────────────────────────────────────────────

type QueryEventsResponse struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *QueryEventsResponse) Reset()         { *m = QueryEventsResponse{} }
func (m *QueryEventsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEventsResponse) ProtoMessage()    {}

func (m *QueryEventsResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEventsResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEventsResponse.Merge(m, src)
}
func (m *QueryEventsResponse) XXX_Size() int       { return m.Size() }
func (m *QueryEventsResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryEventsResponse proto.InternalMessageInfo

func (m *QueryEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryEventsResponse)(nil), "mehr.v1.QueryEventsResponse")
}

func (m *QueryEventsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryEventsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryEventsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryExtra(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QueryEventsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	for _, e := range m.Events {
		l := e.Size()
		n += 1 + sovQueryExtra(uint64(l)) + l
	}
	return n
}
func (m *QueryEventsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryEventsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEventsResponse: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
				return ErrInvalidLengthQueryExtra
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryEventRequest ─────────────────────────────────────────────────────────

type QueryEventRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryEventRequest) Reset()         { *m = QueryEventRequest{} }
func (m *QueryEventRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEventRequest) ProtoMessage()    {}

func (m *QueryEventRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEventRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEventRequest.Merge(m, src)
}
func (m *QueryEventRequest) XXX_Size() int       { return m.Size() }
func (m *QueryEventRequest) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryEventRequest proto.InternalMessageInfo

func (m *QueryEventRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryEventRequest)(nil), "mehr.v1.QueryEventRequest")
}

func (m *QueryEventRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryEventRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryEventRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Id != 0 {
		i = encodeVarintQueryExtra(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}
func (m *QueryEventRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	if m.Id != 0 {
		n += 1 + sovQueryExtra(uint64(m.Id))
	}
	return n
}
func (m *QueryEventRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryEventRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEventRequest: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
		default:
			iNdEx = preIndex
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

// ── QueryEventResponse ────────────────────────────────────────────────────────

type QueryEventResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (m *QueryEventResponse) Reset()         { *m = QueryEventResponse{} }
func (m *QueryEventResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEventResponse) ProtoMessage()    {}

func (m *QueryEventResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEventResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEventResponse.Merge(m, src)
}
func (m *QueryEventResponse) XXX_Size() int       { return m.Size() }
func (m *QueryEventResponse) XXX_DiscardUnknown() {}

var xxx_messageInfo_QueryEventResponse proto.InternalMessageInfo

func (m *QueryEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryEventResponse)(nil), "mehr.v1.QueryEventResponse")
}

func (m *QueryEventResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryEventResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryEventResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Event != nil {
		size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryExtra(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryEventResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	if m.Event != nil {
		l := m.Event.Size()
		n += 1 + sovQueryExtra(uint64(l)) + l
	}
	return n
}
func (m *QueryEventResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryExtra
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
			return fmt.Errorf("proto: QueryEventResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEventResponse: illegal tag %d (none)", fieldNum)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryExtra
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
				return ErrInvalidLengthQueryExtra
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryExtra
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
			skippy, err := skipQueryExtra(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryExtra
			}
			if iNdEx+skippy > l {
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

func encodeVarintQueryExtra(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryExtra(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovQueryExtra(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func skipQueryExtra(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryExtra
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
					return 0, ErrIntOverflowQueryExtra
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
					return 0, ErrIntOverflowQueryExtra
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
				return 0, ErrInvalidLengthQueryExtra
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryExtra
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryExtra
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryExtra        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryExtra          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryExtra = fmt.Errorf("proto: unexpected end of group")
)
