// Code generated by protoc-gen-gogo.
// source: modification_tag.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ModificationTag struct {
	Epoch *string `protobuf:"bytes,1,opt,name=epoch" json:"epoch,omitempty"`
	Index *uint32 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
}

func (m *ModificationTag) Reset()         { *m = ModificationTag{} }
func (m *ModificationTag) String() string { return proto.CompactTextString(m) }
func (*ModificationTag) ProtoMessage()    {}

func (m *ModificationTag) GetEpoch() string {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return ""
}

func (m *ModificationTag) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func init() {
}
func (m *ModificationTag) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Epoch = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Index = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipModificationTag(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipModificationTag(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipModificationTag(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}
func (m *ModificationTag) Size() (n int) {
	var l int
	_ = l
	if m.Epoch != nil {
		l = len(*m.Epoch)
		n += 1 + l + sovModificationTag(uint64(l))
	}
	if m.Index != nil {
		n += 1 + sovModificationTag(uint64(*m.Index))
	}
	return n
}

func sovModificationTag(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModificationTag(x uint64) (n int) {
	return sovModificationTag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ModificationTag) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ModificationTag) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Epoch != nil {
		data[i] = 0xa
		i++
		i = encodeVarintModificationTag(data, i, uint64(len(*m.Epoch)))
		i += copy(data[i:], *m.Epoch)
	}
	if m.Index != nil {
		data[i] = 0x10
		i++
		i = encodeVarintModificationTag(data, i, uint64(*m.Index))
	}
	return i, nil
}

func encodeFixed64ModificationTag(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ModificationTag(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintModificationTag(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *ModificationTag) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ModificationTag)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Epoch != nil && that1.Epoch != nil {
		if *this.Epoch != *that1.Epoch {
			return false
		}
	} else if this.Epoch != nil {
		return false
	} else if that1.Epoch != nil {
		return false
	}
	if this.Index != nil && that1.Index != nil {
		if *this.Index != *that1.Index {
			return false
		}
	} else if this.Index != nil {
		return false
	} else if that1.Index != nil {
		return false
	}
	return true
}