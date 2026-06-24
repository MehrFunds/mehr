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

// ── Feed query types ──────────────────────────────────────────────────────────

type QueryFeedRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryFeedRequest) Reset()         { *m = QueryFeedRequest{} }
func (m *QueryFeedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeedRequest) ProtoMessage()    {}
func (m *QueryFeedRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeedRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFeedRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFeedRequest.Merge(m, src) }
func (m *QueryFeedRequest) XXX_Size() int                  { return m.Size() }
func (m *QueryFeedRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFeedRequest.DiscardUnknown(m) }

var xxx_messageInfo_QueryFeedRequest proto.InternalMessageInfo

func (m *QueryFeedRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryFeedResponse struct {
	Feed *Feed `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
}

func (m *QueryFeedResponse) Reset()         { *m = QueryFeedResponse{} }
func (m *QueryFeedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeedResponse) ProtoMessage()    {}
func (m *QueryFeedResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeedResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFeedResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFeedResponse.Merge(m, src) }
func (m *QueryFeedResponse) XXX_Size() int                  { return m.Size() }
func (m *QueryFeedResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFeedResponse.DiscardUnknown(m) }

var xxx_messageInfo_QueryFeedResponse proto.InternalMessageInfo

func (m *QueryFeedResponse) GetFeed() *Feed {
	if m != nil {
		return m.Feed
	}
	return nil
}

type QueryAllFeedsRequest struct{}

func (m *QueryAllFeedsRequest) Reset()         { *m = QueryAllFeedsRequest{} }
func (m *QueryAllFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllFeedsRequest) ProtoMessage()    {}
func (m *QueryAllFeedsRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllFeedsRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllFeedsRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryAllFeedsRequest.Merge(m, src) }
func (m *QueryAllFeedsRequest) XXX_Size() int                  { return m.Size() }
func (m *QueryAllFeedsRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QueryAllFeedsRequest.DiscardUnknown(m) }

var xxx_messageInfo_QueryAllFeedsRequest proto.InternalMessageInfo

type QueryAllFeedsResponse struct {
	Feeds []*Feed `protobuf:"bytes,1,rep,name=feeds,proto3" json:"feeds,omitempty"`
}

func (m *QueryAllFeedsResponse) Reset()         { *m = QueryAllFeedsResponse{} }
func (m *QueryAllFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllFeedsResponse) ProtoMessage()    {}
func (m *QueryAllFeedsResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllFeedsResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllFeedsResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryAllFeedsResponse.Merge(m, src) }
func (m *QueryAllFeedsResponse) XXX_Size() int                  { return m.Size() }
func (m *QueryAllFeedsResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QueryAllFeedsResponse.DiscardUnknown(m) }

var xxx_messageInfo_QueryAllFeedsResponse proto.InternalMessageInfo

func (m *QueryAllFeedsResponse) GetFeeds() []*Feed {
	if m != nil {
		return m.Feeds
	}
	return nil
}

// ── Fact query types ──────────────────────────────────────────────────────────

type QueryFactRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryFactRequest) Reset()         { *m = QueryFactRequest{} }
func (m *QueryFactRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFactRequest) ProtoMessage()    {}
func (m *QueryFactRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFactRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFactRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFactRequest.Merge(m, src) }
func (m *QueryFactRequest) XXX_Size() int                  { return m.Size() }
func (m *QueryFactRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFactRequest.DiscardUnknown(m) }

var xxx_messageInfo_QueryFactRequest proto.InternalMessageInfo

func (m *QueryFactRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryFactResponse struct {
	Fact *Fact `protobuf:"bytes,1,opt,name=fact,proto3" json:"fact,omitempty"`
}

func (m *QueryFactResponse) Reset()         { *m = QueryFactResponse{} }
func (m *QueryFactResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFactResponse) ProtoMessage()    {}
func (m *QueryFactResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFactResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFactResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFactResponse.Merge(m, src) }
func (m *QueryFactResponse) XXX_Size() int                  { return m.Size() }
func (m *QueryFactResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFactResponse.DiscardUnknown(m) }

var xxx_messageInfo_QueryFactResponse proto.InternalMessageInfo

func (m *QueryFactResponse) GetFact() *Fact {
	if m != nil {
		return m.Fact
	}
	return nil
}

type QueryFactsByFeedRequest struct {
	FeedId string `protobuf:"bytes,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
}

func (m *QueryFactsByFeedRequest) Reset()         { *m = QueryFactsByFeedRequest{} }
func (m *QueryFactsByFeedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFactsByFeedRequest) ProtoMessage()    {}
func (m *QueryFactsByFeedRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFactsByFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFactsByFeedRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFactsByFeedRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFactsByFeedRequest.Merge(m, src) }
func (m *QueryFactsByFeedRequest) XXX_Size() int                  { return m.Size() }
func (m *QueryFactsByFeedRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFactsByFeedRequest.DiscardUnknown(m) }

var xxx_messageInfo_QueryFactsByFeedRequest proto.InternalMessageInfo

func (m *QueryFactsByFeedRequest) GetFeedId() string {
	if m != nil {
		return m.FeedId
	}
	return ""
}

type QueryFactsByFeedResponse struct {
	Facts []*Fact `protobuf:"bytes,1,rep,name=facts,proto3" json:"facts,omitempty"`
}

func (m *QueryFactsByFeedResponse) Reset()         { *m = QueryFactsByFeedResponse{} }
func (m *QueryFactsByFeedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFactsByFeedResponse) ProtoMessage()    {}
func (m *QueryFactsByFeedResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryFactsByFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFactsByFeedResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryFactsByFeedResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryFactsByFeedResponse.Merge(m, src) }
func (m *QueryFactsByFeedResponse) XXX_Size() int                  { return m.Size() }
func (m *QueryFactsByFeedResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QueryFactsByFeedResponse.DiscardUnknown(m) }

var xxx_messageInfo_QueryFactsByFeedResponse proto.InternalMessageInfo

func (m *QueryFactsByFeedResponse) GetFacts() []*Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

// ── Subscription query types ──────────────────────────────────────────────────

type QuerySubscriptionRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QuerySubscriptionRequest) Reset()         { *m = QuerySubscriptionRequest{} }
func (m *QuerySubscriptionRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriptionRequest) ProtoMessage()    {}
func (m *QuerySubscriptionRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QuerySubscriptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriptionRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QuerySubscriptionRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QuerySubscriptionRequest.Merge(m, src) }
func (m *QuerySubscriptionRequest) XXX_Size() int                  { return m.Size() }
func (m *QuerySubscriptionRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QuerySubscriptionRequest.DiscardUnknown(m) }

var xxx_messageInfo_QuerySubscriptionRequest proto.InternalMessageInfo

func (m *QuerySubscriptionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QuerySubscriptionResponse struct {
	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (m *QuerySubscriptionResponse) Reset()         { *m = QuerySubscriptionResponse{} }
func (m *QuerySubscriptionResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriptionResponse) ProtoMessage()    {}
func (m *QuerySubscriptionResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QuerySubscriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriptionResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QuerySubscriptionResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QuerySubscriptionResponse.Merge(m, src) }
func (m *QuerySubscriptionResponse) XXX_Size() int                  { return m.Size() }
func (m *QuerySubscriptionResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QuerySubscriptionResponse.DiscardUnknown(m) }

var xxx_messageInfo_QuerySubscriptionResponse proto.InternalMessageInfo

func (m *QuerySubscriptionResponse) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type QuerySubscriptionsByOwnerRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *QuerySubscriptionsByOwnerRequest) Reset()         { *m = QuerySubscriptionsByOwnerRequest{} }
func (m *QuerySubscriptionsByOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriptionsByOwnerRequest) ProtoMessage()    {}
func (m *QuerySubscriptionsByOwnerRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QuerySubscriptionsByOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriptionsByOwnerRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QuerySubscriptionsByOwnerRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QuerySubscriptionsByOwnerRequest.Merge(m, src) }
func (m *QuerySubscriptionsByOwnerRequest) XXX_Size() int                  { return m.Size() }
func (m *QuerySubscriptionsByOwnerRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QuerySubscriptionsByOwnerRequest.DiscardUnknown(m) }

var xxx_messageInfo_QuerySubscriptionsByOwnerRequest proto.InternalMessageInfo

func (m *QuerySubscriptionsByOwnerRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type QuerySubscriptionsByOwnerResponse struct {
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (m *QuerySubscriptionsByOwnerResponse) Reset()         { *m = QuerySubscriptionsByOwnerResponse{} }
func (m *QuerySubscriptionsByOwnerResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriptionsByOwnerResponse) ProtoMessage()    {}
func (m *QuerySubscriptionsByOwnerResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QuerySubscriptionsByOwnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriptionsByOwnerResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QuerySubscriptionsByOwnerResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QuerySubscriptionsByOwnerResponse.Merge(m, src) }
func (m *QuerySubscriptionsByOwnerResponse) XXX_Size() int                  { return m.Size() }
func (m *QuerySubscriptionsByOwnerResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QuerySubscriptionsByOwnerResponse.DiscardUnknown(m) }

var xxx_messageInfo_QuerySubscriptionsByOwnerResponse proto.InternalMessageInfo

func (m *QuerySubscriptionsByOwnerResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type QueryAllSubscriptionsRequest struct{}

func (m *QueryAllSubscriptionsRequest) Reset()         { *m = QueryAllSubscriptionsRequest{} }
func (m *QueryAllSubscriptionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSubscriptionsRequest) ProtoMessage()    {}
func (m *QueryAllSubscriptionsRequest) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllSubscriptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSubscriptionsRequest.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllSubscriptionsRequest) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryAllSubscriptionsRequest.Merge(m, src) }
func (m *QueryAllSubscriptionsRequest) XXX_Size() int                  { return m.Size() }
func (m *QueryAllSubscriptionsRequest) XXX_DiscardUnknown()            { xxx_messageInfo_QueryAllSubscriptionsRequest.DiscardUnknown(m) }

var xxx_messageInfo_QueryAllSubscriptionsRequest proto.InternalMessageInfo

type QueryAllSubscriptionsResponse struct {
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (m *QueryAllSubscriptionsResponse) Reset()         { *m = QueryAllSubscriptionsResponse{} }
func (m *QueryAllSubscriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllSubscriptionsResponse) ProtoMessage()    {}
func (m *QueryAllSubscriptionsResponse) XXX_Unmarshal(b []byte) error { return m.Unmarshal(b) }
func (m *QueryAllSubscriptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSubscriptionsResponse.Marshal(b, m, deterministic)
	}
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllSubscriptionsResponse) XXX_Merge(src proto.Message)    { xxx_messageInfo_QueryAllSubscriptionsResponse.Merge(m, src) }
func (m *QueryAllSubscriptionsResponse) XXX_Size() int                  { return m.Size() }
func (m *QueryAllSubscriptionsResponse) XXX_DiscardUnknown()            { xxx_messageInfo_QueryAllSubscriptionsResponse.DiscardUnknown(m) }

var xxx_messageInfo_QueryAllSubscriptionsResponse proto.InternalMessageInfo

func (m *QueryAllSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

// ── init ──────────────────────────────────────────────────────────────────────

func init() {
	proto.RegisterType((*QueryFeedRequest)(nil), "mehr.v1.QueryFeedRequest")
	proto.RegisterType((*QueryFeedResponse)(nil), "mehr.v1.QueryFeedResponse")
	proto.RegisterType((*QueryAllFeedsRequest)(nil), "mehr.v1.QueryAllFeedsRequest")
	proto.RegisterType((*QueryAllFeedsResponse)(nil), "mehr.v1.QueryAllFeedsResponse")
	proto.RegisterType((*QueryFactRequest)(nil), "mehr.v1.QueryFactRequest")
	proto.RegisterType((*QueryFactResponse)(nil), "mehr.v1.QueryFactResponse")
	proto.RegisterType((*QueryFactsByFeedRequest)(nil), "mehr.v1.QueryFactsByFeedRequest")
	proto.RegisterType((*QueryFactsByFeedResponse)(nil), "mehr.v1.QueryFactsByFeedResponse")
	proto.RegisterType((*QuerySubscriptionRequest)(nil), "mehr.v1.QuerySubscriptionRequest")
	proto.RegisterType((*QuerySubscriptionResponse)(nil), "mehr.v1.QuerySubscriptionResponse")
	proto.RegisterType((*QuerySubscriptionsByOwnerRequest)(nil), "mehr.v1.QuerySubscriptionsByOwnerRequest")
	proto.RegisterType((*QuerySubscriptionsByOwnerResponse)(nil), "mehr.v1.QuerySubscriptionsByOwnerResponse")
	proto.RegisterType((*QueryAllSubscriptionsRequest)(nil), "mehr.v1.QueryAllSubscriptionsRequest")
	proto.RegisterType((*QueryAllSubscriptionsResponse)(nil), "mehr.v1.QueryAllSubscriptionsResponse")
}

// ── Marshal/Unmarshal implementations ────────────────────────────────────────

// QueryFeedRequest
func (m *QueryFeedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFeedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFeedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQueryFeeds(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryFeedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QueryFeedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFeedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryFeedResponse
func (m *QueryFeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Feed != nil {
		{
			size, err := m.Feed.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryFeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Feed != nil {
		l = m.Feed.Size()
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QueryFeedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFeedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Feed == nil {
				m.Feed = &Feed{}
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryAllFeedsRequest
func (m *QueryAllFeedsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllFeedsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllFeedsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}
func (m *QueryAllFeedsRequest) Size() (n int) { return 0 }
func (m *QueryAllFeedsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryAllFeedsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllFeedsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryAllFeedsResponse
func (m *QueryAllFeedsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllFeedsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllFeedsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feeds) > 0 {
		for iNdEx := len(m.Feeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Feeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QueryAllFeedsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Feeds) > 0 {
		for _, e := range m.Feeds {
			l = e.Size()
			n += 1 + l + sovQueryFeeds(uint64(l))
		}
	}
	return n
}
func (m *QueryAllFeedsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryAllFeedsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllFeedsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feeds = append(m.Feeds, &Feed{})
			if err := m.Feeds[len(m.Feeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryFactRequest
func (m *QueryFactRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFactRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFactRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQueryFeeds(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}
func (m *QueryFactRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQueryFeeds(uint64(m.Id))
	}
	return n
}
func (m *QueryFactRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFactRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFactRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryFactResponse
func (m *QueryFactResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFactResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFactResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fact != nil {
		{
			size, err := m.Fact.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryFactResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fact != nil {
		l = m.Fact.Size()
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QueryFactResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFactResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFactResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fact", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
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
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryFactsByFeedRequest
func (m *QueryFactsByFeedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFactsByFeedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFactsByFeedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeedId) > 0 {
		i -= len(m.FeedId)
		copy(dAtA[i:], m.FeedId)
		i = encodeVarintQueryFeeds(dAtA, i, uint64(len(m.FeedId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QueryFactsByFeedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeedId)
	if l > 0 {
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QueryFactsByFeedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFactsByFeedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFactsByFeedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeedId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeedId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryFactsByFeedResponse
func (m *QueryFactsByFeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryFactsByFeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryFactsByFeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for iNdEx := len(m.Facts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Facts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QueryFactsByFeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for _, e := range m.Facts {
			l = e.Size()
			n += 1 + l + sovQueryFeeds(uint64(l))
		}
	}
	return n
}
func (m *QueryFactsByFeedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryFactsByFeedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFactsByFeedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Facts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Facts = append(m.Facts, &Fact{})
			if err := m.Facts[len(m.Facts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QuerySubscriptionRequest
func (m *QuerySubscriptionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QuerySubscriptionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QuerySubscriptionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQueryFeeds(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}
func (m *QuerySubscriptionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQueryFeeds(uint64(m.Id))
	}
	return n
}
func (m *QuerySubscriptionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QuerySubscriptionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriptionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QuerySubscriptionResponse
func (m *QuerySubscriptionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QuerySubscriptionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QuerySubscriptionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Subscription != nil {
		{
			size, err := m.Subscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QuerySubscriptionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscription != nil {
		l = m.Subscription.Size()
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QuerySubscriptionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QuerySubscriptionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriptionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subscription == nil {
				m.Subscription = &Subscription{}
			}
			if err := m.Subscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QuerySubscriptionsByOwnerRequest
func (m *QuerySubscriptionsByOwnerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QuerySubscriptionsByOwnerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QuerySubscriptionsByOwnerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQueryFeeds(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *QuerySubscriptionsByOwnerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQueryFeeds(uint64(l))
	}
	return n
}
func (m *QuerySubscriptionsByOwnerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QuerySubscriptionsByOwnerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriptionsByOwnerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QuerySubscriptionsByOwnerResponse
func (m *QuerySubscriptionsByOwnerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QuerySubscriptionsByOwnerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QuerySubscriptionsByOwnerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subscriptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QuerySubscriptionsByOwnerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovQueryFeeds(uint64(l))
		}
	}
	return n
}
func (m *QuerySubscriptionsByOwnerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QuerySubscriptionsByOwnerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriptionsByOwnerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, &Subscription{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryAllSubscriptionsRequest
func (m *QueryAllSubscriptionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllSubscriptionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllSubscriptionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	return len(dAtA) - i, nil
}
func (m *QueryAllSubscriptionsRequest) Size() (n int) { return 0 }
func (m *QueryAllSubscriptionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryAllSubscriptionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSubscriptionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

// QueryAllSubscriptionsResponse
func (m *QueryAllSubscriptionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *QueryAllSubscriptionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}
func (m *QueryAllSubscriptionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subscriptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryFeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}
func (m *QueryAllSubscriptionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovQueryFeeds(uint64(l))
		}
	}
	return n
}
func (m *QueryAllSubscriptionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryFeeds
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
			return fmt.Errorf("proto: QueryAllSubscriptionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSubscriptionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryFeeds
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
				return ErrInvalidLengthQueryFeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryFeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, &Subscription{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryFeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryFeeds
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

func encodeVarintQueryFeeds(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryFeeds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func sovQueryFeeds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func skipQueryFeeds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryFeeds
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
					return 0, ErrIntOverflowQueryFeeds
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
					return 0, ErrIntOverflowQueryFeeds
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
				return 0, ErrInvalidLengthQueryFeeds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryFeeds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryFeeds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryFeeds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryFeeds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryFeeds = fmt.Errorf("proto: unexpected end of group")
)
