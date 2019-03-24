// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/sem/types/types.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SemanticType int32

const (
	BOOL           SemanticType = 0
	INT            SemanticType = 1
	FLOAT          SemanticType = 2
	DECIMAL        SemanticType = 3
	DATE           SemanticType = 4
	TIMESTAMP      SemanticType = 5
	INTERVAL       SemanticType = 6
	STRING         SemanticType = 7
	BYTES          SemanticType = 8
	TIMESTAMPTZ    SemanticType = 9
	COLLATEDSTRING SemanticType = 10
	NAME           SemanticType = 11
	OID            SemanticType = 12
	// NULL is not supported as a table column type, however it can be
	// transferred through distsql streams.
	NULL       SemanticType = 13
	UUID       SemanticType = 14
	ARRAY      SemanticType = 15
	INET       SemanticType = 16
	TIME       SemanticType = 17
	JSONB      SemanticType = 18
	TUPLE      SemanticType = 20
	BIT        SemanticType = 21
	INT2VECTOR SemanticType = 200
	OIDVECTOR  SemanticType = 201
	// Special type used during static analysis. This should never be present
	// at execution time.
	ANY SemanticType = 100
)

var SemanticType_name = map[int32]string{
	0:   "BOOL",
	1:   "INT",
	2:   "FLOAT",
	3:   "DECIMAL",
	4:   "DATE",
	5:   "TIMESTAMP",
	6:   "INTERVAL",
	7:   "STRING",
	8:   "BYTES",
	9:   "TIMESTAMPTZ",
	10:  "COLLATEDSTRING",
	11:  "NAME",
	12:  "OID",
	13:  "NULL",
	14:  "UUID",
	15:  "ARRAY",
	16:  "INET",
	17:  "TIME",
	18:  "JSONB",
	20:  "TUPLE",
	21:  "BIT",
	200: "INT2VECTOR",
	201: "OIDVECTOR",
	100: "ANY",
}
var SemanticType_value = map[string]int32{
	"BOOL":           0,
	"INT":            1,
	"FLOAT":          2,
	"DECIMAL":        3,
	"DATE":           4,
	"TIMESTAMP":      5,
	"INTERVAL":       6,
	"STRING":         7,
	"BYTES":          8,
	"TIMESTAMPTZ":    9,
	"COLLATEDSTRING": 10,
	"NAME":           11,
	"OID":            12,
	"NULL":           13,
	"UUID":           14,
	"ARRAY":          15,
	"INET":           16,
	"TIME":           17,
	"JSONB":          18,
	"TUPLE":          20,
	"BIT":            21,
	"INT2VECTOR":     200,
	"OIDVECTOR":      201,
	"ANY":            100,
}

func (x SemanticType) Enum() *SemanticType {
	p := new(SemanticType)
	*p = x
	return p
}
func (x SemanticType) String() string {
	return proto.EnumName(SemanticType_name, int32(x))
}
func (x *SemanticType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SemanticType_value, data, "SemanticType")
	if err != nil {
		return err
	}
	*x = SemanticType(value)
	return nil
}
func (SemanticType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_77fe5a0620518a56, []int{0}
}

type VisibleType int32

const (
	VisibleType_NONE             VisibleType = 0
	VisibleType_INTEGER          VisibleType = 1
	VisibleType_SMALLINT         VisibleType = 2
	VisibleType_BIGINT           VisibleType = 3
	VisibleType_REAL             VisibleType = 5
	VisibleType_DOUBLE_PRECISION VisibleType = 6
	VisibleType_VARCHAR          VisibleType = 7
	VisibleType_CHAR             VisibleType = 8
	VisibleType_QCHAR            VisibleType = 9
	VisibleType_VARBIT           VisibleType = 10
)

var VisibleType_name = map[int32]string{
	0:  "NONE",
	1:  "INTEGER",
	2:  "SMALLINT",
	3:  "BIGINT",
	5:  "REAL",
	6:  "DOUBLE_PRECISION",
	7:  "VARCHAR",
	8:  "CHAR",
	9:  "QCHAR",
	10: "VARBIT",
}
var VisibleType_value = map[string]int32{
	"NONE":             0,
	"INTEGER":          1,
	"SMALLINT":         2,
	"BIGINT":           3,
	"REAL":             5,
	"DOUBLE_PRECISION": 6,
	"VARCHAR":          7,
	"CHAR":             8,
	"QCHAR":            9,
	"VARBIT":           10,
}

func (x VisibleType) Enum() *VisibleType {
	p := new(VisibleType)
	*p = x
	return p
}
func (x VisibleType) String() string {
	return proto.EnumName(VisibleType_name, int32(x))
}
func (x *VisibleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VisibleType_value, data, "VisibleType")
	if err != nil {
		return err
	}
	*x = VisibleType(value)
	return nil
}
func (VisibleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_77fe5a0620518a56, []int{1}
}

type ColumnType struct {
	SemanticType SemanticType `protobuf:"varint,1,opt,name=semantic_type,json=semanticType,enum=cockroach.sql.sem.types.SemanticType" json:"semantic_type"`
	// INT, DECIMAL, CHAR and BINARY
	Width int32 `protobuf:"varint,2,opt,name=width" json:"width"`
	// DECIMAL
	// Also FLOAT pre-2.1 (this was incorrect.)
	Precision int32 `protobuf:"varint,3,opt,name=precision" json:"precision"`
	// The length of each dimension in the array. A dimension of -1 means that
	// no bound was specified for that dimension.
	ArrayDimensions []int32 `protobuf:"varint,4,rep,name=array_dimensions,json=arrayDimensions" json:"array_dimensions,omitempty"`
	// Collated STRING, CHAR, and VARCHAR
	Locale *string `protobuf:"bytes,5,opt,name=locale" json:"locale,omitempty"`
	// Alias for any types where our internal representation is different than
	// the user specification. Examples are INT4, FLOAT4, etc. Mostly for Postgres
	// compatibility.
	VisibleType VisibleType `protobuf:"varint,6,opt,name=visible_type,json=visibleType,enum=cockroach.sql.sem.types.VisibleType" json:"visible_type"`
	// Only used if the kind is ARRAY.
	ArrayContents *SemanticType `protobuf:"varint,7,opt,name=array_contents,json=arrayContents,enum=cockroach.sql.sem.types.SemanticType" json:"array_contents,omitempty"`
	// Only used if the kind is TUPLE
	TupleContents        []ColumnType `protobuf:"bytes,8,rep,name=tuple_contents,json=tupleContents" json:"tuple_contents"`
	TupleLabels          []string     `protobuf:"bytes,9,rep,name=tuple_labels,json=tupleLabels" json:"tuple_labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ColumnType) Reset()         { *m = ColumnType{} }
func (m *ColumnType) String() string { return proto.CompactTextString(m) }
func (*ColumnType) ProtoMessage()    {}
func (*ColumnType) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_77fe5a0620518a56, []int{0}
}
func (m *ColumnType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ColumnType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ColumnType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnType.Merge(dst, src)
}
func (m *ColumnType) XXX_Size() int {
	return m.Size()
}
func (m *ColumnType) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnType.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnType proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ColumnType)(nil), "cockroach.sql.sem.types.ColumnType")
	proto.RegisterEnum("cockroach.sql.sem.types.SemanticType", SemanticType_name, SemanticType_value)
	proto.RegisterEnum("cockroach.sql.sem.types.VisibleType", VisibleType_name, VisibleType_value)
}
func (this *ColumnType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ColumnType)
	if !ok {
		that2, ok := that.(ColumnType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SemanticType != that1.SemanticType {
		return false
	}
	if this.Width != that1.Width {
		return false
	}
	if this.Precision != that1.Precision {
		return false
	}
	if len(this.ArrayDimensions) != len(that1.ArrayDimensions) {
		return false
	}
	for i := range this.ArrayDimensions {
		if this.ArrayDimensions[i] != that1.ArrayDimensions[i] {
			return false
		}
	}
	if this.Locale != nil && that1.Locale != nil {
		if *this.Locale != *that1.Locale {
			return false
		}
	} else if this.Locale != nil {
		return false
	} else if that1.Locale != nil {
		return false
	}
	if this.VisibleType != that1.VisibleType {
		return false
	}
	if this.ArrayContents != nil && that1.ArrayContents != nil {
		if *this.ArrayContents != *that1.ArrayContents {
			return false
		}
	} else if this.ArrayContents != nil {
		return false
	} else if that1.ArrayContents != nil {
		return false
	}
	if len(this.TupleContents) != len(that1.TupleContents) {
		return false
	}
	for i := range this.TupleContents {
		if !this.TupleContents[i].Equal(&that1.TupleContents[i]) {
			return false
		}
	}
	if len(this.TupleLabels) != len(that1.TupleLabels) {
		return false
	}
	for i := range this.TupleLabels {
		if this.TupleLabels[i] != that1.TupleLabels[i] {
			return false
		}
	}
	return true
}
func (m *ColumnType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ColumnType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.SemanticType))
	dAtA[i] = 0x10
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.Width))
	dAtA[i] = 0x18
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.Precision))
	if len(m.ArrayDimensions) > 0 {
		for _, num := range m.ArrayDimensions {
			dAtA[i] = 0x20
			i++
			i = encodeVarintTypes(dAtA, i, uint64(num))
		}
	}
	if m.Locale != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTypes(dAtA, i, uint64(len(*m.Locale)))
		i += copy(dAtA[i:], *m.Locale)
	}
	dAtA[i] = 0x30
	i++
	i = encodeVarintTypes(dAtA, i, uint64(m.VisibleType))
	if m.ArrayContents != nil {
		dAtA[i] = 0x38
		i++
		i = encodeVarintTypes(dAtA, i, uint64(*m.ArrayContents))
	}
	if len(m.TupleContents) > 0 {
		for _, msg := range m.TupleContents {
			dAtA[i] = 0x42
			i++
			i = encodeVarintTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.TupleLabels) > 0 {
		for _, s := range m.TupleLabels {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ColumnType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovTypes(uint64(m.SemanticType))
	n += 1 + sovTypes(uint64(m.Width))
	n += 1 + sovTypes(uint64(m.Precision))
	if len(m.ArrayDimensions) > 0 {
		for _, e := range m.ArrayDimensions {
			n += 1 + sovTypes(uint64(e))
		}
	}
	if m.Locale != nil {
		l = len(*m.Locale)
		n += 1 + l + sovTypes(uint64(l))
	}
	n += 1 + sovTypes(uint64(m.VisibleType))
	if m.ArrayContents != nil {
		n += 1 + sovTypes(uint64(*m.ArrayContents))
	}
	if len(m.TupleContents) > 0 {
		for _, e := range m.TupleContents {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.TupleLabels) > 0 {
		for _, s := range m.TupleLabels {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ColumnType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ColumnType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SemanticType", wireType)
			}
			m.SemanticType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SemanticType |= (SemanticType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ArrayDimensions = append(m.ArrayDimensions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ArrayDimensions) == 0 {
					m.ArrayDimensions = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ArrayDimensions = append(m.ArrayDimensions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayDimensions", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locale", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Locale = &s
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VisibleType", wireType)
			}
			m.VisibleType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VisibleType |= (VisibleType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayContents", wireType)
			}
			var v SemanticType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (SemanticType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ArrayContents = &v
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TupleContents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TupleContents = append(m.TupleContents, ColumnType{})
			if err := m.TupleContents[len(m.TupleContents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TupleLabels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TupleLabels = append(m.TupleLabels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipTypes(dAtA[start:])
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

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sql/sem/types/types.proto", fileDescriptor_types_77fe5a0620518a56) }

var fileDescriptor_types_77fe5a0620518a56 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0x12, 0x51,
	0x14, 0x66, 0x60, 0x06, 0x98, 0xc3, 0xdf, 0xf1, 0xb6, 0x2a, 0xb2, 0x98, 0x22, 0x6a, 0x82, 0x5d,
	0xd0, 0xa4, 0x4b, 0x77, 0x33, 0xcc, 0xb5, 0x8e, 0xb9, 0xcc, 0xe0, 0xe5, 0x42, 0xd2, 0x6e, 0x1a,
	0x4a, 0x27, 0x2d, 0x71, 0x60, 0x28, 0x43, 0x6b, 0xfa, 0x06, 0x2e, 0x7d, 0x04, 0x93, 0xfa, 0x0e,
	0x26, 0x3e, 0x01, 0xee, 0x5c, 0xba, 0x32, 0x8a, 0x1b, 0x5f, 0xc1, 0x9d, 0xb9, 0x03, 0x29, 0x6c,
	0xba, 0x70, 0x33, 0xf9, 0xee, 0x77, 0xbe, 0xf3, 0xdd, 0xf3, 0x9d, 0xb9, 0xf0, 0x28, 0xba, 0x08,
	0xf6, 0x22, 0x7f, 0xb4, 0x37, 0xbb, 0x9e, 0xf8, 0xd1, 0xf2, 0xdb, 0x98, 0x4c, 0xc3, 0x59, 0x48,
	0x1e, 0x0e, 0xc2, 0xc1, 0xdb, 0x69, 0xd8, 0x1f, 0x9c, 0x37, 0xa2, 0x8b, 0xa0, 0x11, 0xf9, 0xa3,
	0x46, 0x5c, 0xae, 0x6c, 0x9f, 0x85, 0x67, 0x61, 0xac, 0xd9, 0x93, 0x68, 0x29, 0xaf, 0xfd, 0x4d,
	0x01, 0x34, 0xc3, 0xe0, 0x72, 0x34, 0x16, 0xd7, 0x13, 0x9f, 0xb4, 0xa1, 0x10, 0xf9, 0xa3, 0xfe,
	0x78, 0x36, 0x1c, 0x1c, 0xcb, 0xb6, 0xb2, 0x52, 0x55, 0xea, 0xc5, 0xfd, 0x67, 0x8d, 0x3b, 0x5c,
	0x1b, 0x9d, 0x95, 0x5a, 0x76, 0x5b, 0xea, 0xfc, 0xc7, 0x4e, 0x82, 0xe7, 0xa3, 0x0d, 0x8e, 0x54,
	0x40, 0x7b, 0x37, 0x3c, 0x9d, 0x9d, 0x97, 0x93, 0x55, 0xa5, 0xae, 0xad, 0x24, 0x4b, 0x8a, 0xd4,
	0x40, 0x9f, 0x4c, 0xfd, 0xc1, 0x30, 0x1a, 0x86, 0xe3, 0x72, 0x6a, 0xa3, 0xbe, 0xa6, 0xc9, 0x73,
	0xc0, 0xfe, 0x74, 0xda, 0xbf, 0x3e, 0x3e, 0x1d, 0x8e, 0xfc, 0xb1, 0xa4, 0xa2, 0xb2, 0x5a, 0x4d,
	0xd5, 0x35, 0x5e, 0x8a, 0x79, 0xfb, 0x96, 0x26, 0x0f, 0x20, 0x1d, 0x84, 0x83, 0x7e, 0xe0, 0x97,
	0xb5, 0xaa, 0x52, 0xd7, 0xf9, 0xea, 0x44, 0x5a, 0x90, 0xbf, 0x1a, 0x46, 0xc3, 0x93, 0xc0, 0x5f,
	0x66, 0x4a, 0xc7, 0x99, 0x9e, 0xde, 0x99, 0xa9, 0xb7, 0x14, 0x6f, 0x44, 0xca, 0x5d, 0xad, 0x29,
	0xc2, 0xa0, 0xb8, 0x9c, 0x68, 0x10, 0x8e, 0x67, 0xfe, 0x78, 0x16, 0x95, 0x33, 0xff, 0xb1, 0x24,
	0x5e, 0x88, 0x9b, 0x9b, 0xab, 0x5e, 0xd2, 0x86, 0xe2, 0xec, 0x72, 0x12, 0xf8, 0x6b, 0xb7, 0x6c,
	0x35, 0x55, 0xcf, 0xed, 0x3f, 0xb9, 0xd3, 0x6d, 0xfd, 0xbb, 0x56, 0xd3, 0x15, 0x62, 0x83, 0x5b,
	0xc7, 0xc7, 0x90, 0x5f, 0x3a, 0x06, 0xfd, 0x13, 0x3f, 0x88, 0xca, 0x7a, 0x35, 0x55, 0xd7, 0x79,
	0x2e, 0xe6, 0x58, 0x4c, 0xbd, 0x50, 0xff, 0x7c, 0xdc, 0x51, 0x76, 0x3f, 0x27, 0x21, 0xbf, 0x39,
	0x1a, 0xc9, 0x82, 0x6a, 0x79, 0x1e, 0xc3, 0x04, 0xc9, 0x40, 0xca, 0x71, 0x05, 0x2a, 0x44, 0x07,
	0xed, 0x25, 0xf3, 0x4c, 0x81, 0x49, 0x92, 0x83, 0x8c, 0x4d, 0x9b, 0x4e, 0xcb, 0x64, 0x98, 0x92,
	0x52, 0xdb, 0x14, 0x14, 0x55, 0x52, 0x00, 0x5d, 0x38, 0x2d, 0xda, 0x11, 0x66, 0xab, 0x8d, 0x1a,
	0xc9, 0x43, 0xd6, 0x71, 0x05, 0xe5, 0x3d, 0x93, 0x61, 0x9a, 0x00, 0xa4, 0x3b, 0x82, 0x3b, 0xee,
	0x01, 0x66, 0xa4, 0x95, 0x75, 0x28, 0x68, 0x07, 0xb3, 0xa4, 0x04, 0xb9, 0xdb, 0x1e, 0x71, 0x84,
	0x3a, 0x21, 0x50, 0x6c, 0x7a, 0x8c, 0x99, 0x82, 0xda, 0x2b, 0x3d, 0xc8, 0x2b, 0x5c, 0xb3, 0x45,
	0x31, 0x27, 0xa7, 0xf1, 0x1c, 0x1b, 0xf3, 0x31, 0xd5, 0x65, 0x0c, 0x0b, 0x12, 0x75, 0xbb, 0x8e,
	0x8d, 0x45, 0x69, 0x6b, 0x72, 0x6e, 0x1e, 0x62, 0x49, 0x92, 0x8e, 0x4b, 0x05, 0xa2, 0x44, 0xf2,
	0x02, 0xbc, 0x27, 0xcb, 0xaf, 0x3b, 0x9e, 0x6b, 0x21, 0x91, 0x50, 0x74, 0xdb, 0x8c, 0xe2, 0xb6,
	0x74, 0xb4, 0x1c, 0x81, 0xf7, 0x49, 0x09, 0xc0, 0x71, 0xc5, 0x7e, 0x8f, 0x36, 0x85, 0xc7, 0x71,
	0xae, 0x90, 0x22, 0xe8, 0x9e, 0x63, 0xaf, 0xce, 0x5f, 0x15, 0xa9, 0x34, 0xdd, 0x43, 0x3c, 0xad,
	0xa8, 0xef, 0x6f, 0x8c, 0x44, 0x4d, 0xcd, 0x6e, 0xe1, 0xd6, 0xee, 0x8d, 0x02, 0xb9, 0x8d, 0x57,
	0x12, 0xcf, 0xe5, 0xb9, 0x14, 0x13, 0x72, 0x49, 0x32, 0xfe, 0x01, 0xe5, 0xa8, 0xc8, 0x5d, 0x74,
	0x5a, 0x26, 0x63, 0x72, 0x95, 0x49, 0xb9, 0x0b, 0xcb, 0x39, 0x90, 0x38, 0x5e, 0x1f, 0xa7, 0x26,
	0x43, 0x8d, 0x6c, 0x03, 0xda, 0x5e, 0xd7, 0x62, 0xf4, 0xb8, 0xcd, 0x69, 0xd3, 0xe9, 0x38, 0x9e,
	0x8b, 0x69, 0x69, 0xd3, 0x33, 0x79, 0xf3, 0x95, 0xc9, 0x31, 0x23, 0xc5, 0x31, 0xca, 0xca, 0x04,
	0x6f, 0x62, 0xa8, 0x4b, 0xb7, 0x9e, 0xc9, 0x65, 0x08, 0xa8, 0x94, 0xbe, 0x7c, 0x32, 0x36, 0xe7,
	0xa9, 0xa9, 0x59, 0x15, 0x55, 0x6b, 0x67, 0xfe, 0xcb, 0x48, 0xcc, 0x17, 0x86, 0xf2, 0x6d, 0x61,
	0x28, 0xdf, 0x17, 0x86, 0xf2, 0x73, 0x61, 0x28, 0x1f, 0x7e, 0x1b, 0x89, 0x23, 0x2d, 0x7e, 0x49,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa0, 0x4a, 0x34, 0x47, 0x04, 0x00, 0x00,
}
