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

// These mirror the types supported by sql/coltypes.
//
// Note: when adding constants to this list or renaming constants,
// verify with PostgreSQL what the type name should be in
// information_schema.columns.data_type, and modify
// (*ColumnType).InformationSchemaVisibleType() accordingly.
//
type ColumnType_SemanticType int32

const (
	ColumnType_BOOL           ColumnType_SemanticType = 0
	ColumnType_INT            ColumnType_SemanticType = 1
	ColumnType_FLOAT          ColumnType_SemanticType = 2
	ColumnType_DECIMAL        ColumnType_SemanticType = 3
	ColumnType_DATE           ColumnType_SemanticType = 4
	ColumnType_TIMESTAMP      ColumnType_SemanticType = 5
	ColumnType_INTERVAL       ColumnType_SemanticType = 6
	ColumnType_STRING         ColumnType_SemanticType = 7
	ColumnType_BYTES          ColumnType_SemanticType = 8
	ColumnType_TIMESTAMPTZ    ColumnType_SemanticType = 9
	ColumnType_COLLATEDSTRING ColumnType_SemanticType = 10
	ColumnType_NAME           ColumnType_SemanticType = 11
	ColumnType_OID            ColumnType_SemanticType = 12
	// NULL is not supported as a table column type, however it can be
	// transferred through distsql streams.
	ColumnType_NULL       ColumnType_SemanticType = 13
	ColumnType_UUID       ColumnType_SemanticType = 14
	ColumnType_ARRAY      ColumnType_SemanticType = 15
	ColumnType_INET       ColumnType_SemanticType = 16
	ColumnType_TIME       ColumnType_SemanticType = 17
	ColumnType_JSONB      ColumnType_SemanticType = 18
	ColumnType_TUPLE      ColumnType_SemanticType = 20
	ColumnType_BIT        ColumnType_SemanticType = 21
	ColumnType_INT2VECTOR ColumnType_SemanticType = 200
	ColumnType_OIDVECTOR  ColumnType_SemanticType = 201
)

var ColumnType_SemanticType_name = map[int32]string{
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
}
var ColumnType_SemanticType_value = map[string]int32{
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
}

func (x ColumnType_SemanticType) Enum() *ColumnType_SemanticType {
	p := new(ColumnType_SemanticType)
	*p = x
	return p
}
func (x ColumnType_SemanticType) String() string {
	return proto.EnumName(ColumnType_SemanticType_name, int32(x))
}
func (x *ColumnType_SemanticType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ColumnType_SemanticType_value, data, "ColumnType_SemanticType")
	if err != nil {
		return err
	}
	*x = ColumnType_SemanticType(value)
	return nil
}
func (ColumnType_SemanticType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_6c941a8efcdf234e, []int{0, 0}
}

type ColumnType_VisibleType int32

const (
	ColumnType_NONE             ColumnType_VisibleType = 0
	ColumnType_INTEGER          ColumnType_VisibleType = 1
	ColumnType_SMALLINT         ColumnType_VisibleType = 2
	ColumnType_BIGINT           ColumnType_VisibleType = 3
	ColumnType_REAL             ColumnType_VisibleType = 5
	ColumnType_DOUBLE_PRECISION ColumnType_VisibleType = 6
	ColumnType_VARCHAR          ColumnType_VisibleType = 7
	ColumnType_CHAR             ColumnType_VisibleType = 8
	ColumnType_QCHAR            ColumnType_VisibleType = 9
	ColumnType_VARBIT           ColumnType_VisibleType = 10
)

var ColumnType_VisibleType_name = map[int32]string{
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
var ColumnType_VisibleType_value = map[string]int32{
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

func (x ColumnType_VisibleType) Enum() *ColumnType_VisibleType {
	p := new(ColumnType_VisibleType)
	*p = x
	return p
}
func (x ColumnType_VisibleType) String() string {
	return proto.EnumName(ColumnType_VisibleType_name, int32(x))
}
func (x *ColumnType_VisibleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ColumnType_VisibleType_value, data, "ColumnType_VisibleType")
	if err != nil {
		return err
	}
	*x = ColumnType_VisibleType(value)
	return nil
}
func (ColumnType_VisibleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_types_6c941a8efcdf234e, []int{0, 1}
}

type ColumnType struct {
	SemanticType ColumnType_SemanticType `protobuf:"varint,1,opt,name=semantic_type,json=semanticType,enum=cockroach.sql.sem.types.ColumnType_SemanticType" json:"semantic_type"`
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
	VisibleType ColumnType_VisibleType `protobuf:"varint,6,opt,name=visible_type,json=visibleType,enum=cockroach.sql.sem.types.ColumnType_VisibleType" json:"visible_type"`
	// Only used if the kind is ARRAY.
	ArrayContents *ColumnType_SemanticType `protobuf:"varint,7,opt,name=array_contents,json=arrayContents,enum=cockroach.sql.sem.types.ColumnType_SemanticType" json:"array_contents,omitempty"`
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
	return fileDescriptor_types_6c941a8efcdf234e, []int{0}
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
	proto.RegisterEnum("cockroach.sql.sem.types.ColumnType_SemanticType", ColumnType_SemanticType_name, ColumnType_SemanticType_value)
	proto.RegisterEnum("cockroach.sql.sem.types.ColumnType_VisibleType", ColumnType_VisibleType_name, ColumnType_VisibleType_value)
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
				m.SemanticType |= (ColumnType_SemanticType(b) & 0x7F) << shift
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
				m.VisibleType |= (ColumnType_VisibleType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrayContents", wireType)
			}
			var v ColumnType_SemanticType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (ColumnType_SemanticType(b) & 0x7F) << shift
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

func init() { proto.RegisterFile("sql/sem/types/types.proto", fileDescriptor_types_6c941a8efcdf234e) }

var fileDescriptor_types_6c941a8efcdf234e = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x72, 0x12, 0x4d,
	0x14, 0xc5, 0x19, 0x98, 0x01, 0xe6, 0xf2, 0xef, 0x7e, 0x9d, 0x7c, 0x8a, 0x59, 0x10, 0xc4, 0x0d,
	0x6e, 0xc0, 0xca, 0xd2, 0xdd, 0x0c, 0xd3, 0xc6, 0xb1, 0x9a, 0x19, 0x6c, 0x1a, 0x34, 0x71, 0x91,
	0x22, 0x64, 0x2a, 0xa1, 0x1c, 0x18, 0xc2, 0x90, 0x58, 0x79, 0x0b, 0xcb, 0x27, 0x70, 0xe9, 0xa3,
	0xe0, 0xce, 0xa5, 0xe5, 0xc2, 0x52, 0xdc, 0xf8, 0x18, 0x56, 0x37, 0x54, 0x60, 0x63, 0x55, 0xca,
	0xcd, 0xd4, 0xed, 0xd3, 0xf7, 0xfe, 0xe6, 0x9c, 0x5b, 0x0d, 0x0f, 0xe2, 0xcb, 0xb0, 0x19, 0x07,
	0xe3, 0xe6, 0xfc, 0x66, 0x1a, 0xc4, 0xab, 0x6f, 0x63, 0x3a, 0x8b, 0xe6, 0x11, 0xb9, 0x3f, 0x8c,
	0x86, 0x6f, 0x67, 0xd1, 0x60, 0x78, 0xd1, 0x88, 0x2f, 0xc3, 0x46, 0x1c, 0x8c, 0x1b, 0xea, 0x7a,
	0x6f, 0xf7, 0x3c, 0x3a, 0x8f, 0x54, 0x4f, 0x53, 0x56, 0xab, 0xf6, 0xda, 0xb7, 0x0c, 0x40, 0x2b,
	0x0a, 0xaf, 0xc6, 0x13, 0x71, 0x33, 0x0d, 0xc8, 0x1b, 0x28, 0xc4, 0xc1, 0x78, 0x30, 0x99, 0x8f,
	0x86, 0x27, 0x72, 0xac, 0xac, 0x55, 0xb5, 0x7a, 0xf1, 0xe0, 0x49, 0xe3, 0x2f, 0xd4, 0xc6, 0x66,
	0xb6, 0xd1, 0x5d, 0x0f, 0xca, 0x83, 0xad, 0x2f, 0xbe, 0xef, 0x27, 0x78, 0x3e, 0xde, 0xd2, 0xc8,
	0x1e, 0x18, 0xef, 0x46, 0x67, 0xf3, 0x8b, 0x72, 0xb2, 0xaa, 0xd5, 0x8d, 0x75, 0xcb, 0x4a, 0x22,
	0x35, 0x30, 0xa7, 0xb3, 0x60, 0x38, 0x8a, 0x47, 0xd1, 0xa4, 0x9c, 0xda, 0xba, 0xdf, 0xc8, 0xe4,
	0x31, 0xe0, 0x60, 0x36, 0x1b, 0xdc, 0x9c, 0x9c, 0x8d, 0xc6, 0xc1, 0x44, 0x4a, 0x71, 0x59, 0xaf,
	0xa6, 0xea, 0x06, 0x2f, 0x29, 0xdd, 0xb9, 0x95, 0xc9, 0x3d, 0x48, 0x87, 0xd1, 0x70, 0x10, 0x06,
	0x65, 0xa3, 0xaa, 0xd5, 0x4d, 0xbe, 0x3e, 0x91, 0xd7, 0x90, 0xbf, 0x1e, 0xc5, 0xa3, 0xd3, 0x30,
	0x58, 0xc5, 0x4b, 0xab, 0x78, 0xcd, 0xbb, 0xc4, 0xeb, 0xaf, 0xe6, 0xb6, 0xd2, 0xe5, 0xae, 0x37,
	0x12, 0x79, 0x05, 0xc5, 0x95, 0xb9, 0x61, 0x34, 0x99, 0x07, 0x93, 0x79, 0x5c, 0xce, 0xfc, 0xdb,
	0xea, 0x78, 0x41, 0x71, 0x5a, 0x6b, 0x0c, 0xe9, 0x40, 0x71, 0x7e, 0x35, 0x0d, 0x83, 0x0d, 0x38,
	0x5b, 0x4d, 0xd5, 0x73, 0x07, 0x8f, 0xee, 0x00, 0x5e, 0x1b, 0x2d, 0x28, 0xc0, 0x2d, 0xf1, 0x21,
	0xe4, 0x57, 0xc4, 0x70, 0x70, 0x1a, 0x84, 0x71, 0xd9, 0xac, 0xa6, 0xea, 0x26, 0xcf, 0x29, 0x8d,
	0x29, 0xa9, 0xf6, 0x29, 0x09, 0xf9, 0x6d, 0x53, 0x24, 0x0b, 0xba, 0xed, 0xfb, 0x0c, 0x13, 0x24,
	0x03, 0x29, 0xd7, 0x13, 0xa8, 0x11, 0x13, 0x8c, 0x67, 0xcc, 0xb7, 0x04, 0x26, 0x49, 0x0e, 0x32,
	0x0e, 0x6d, 0xb9, 0x6d, 0x8b, 0x61, 0x4a, 0xb6, 0x3a, 0x96, 0xa0, 0xa8, 0x93, 0x02, 0x98, 0xc2,
	0x6d, 0xd3, 0xae, 0xb0, 0xda, 0x1d, 0x34, 0x48, 0x1e, 0xb2, 0xae, 0x27, 0x28, 0xef, 0x5b, 0x0c,
	0xd3, 0x04, 0x20, 0xdd, 0x15, 0xdc, 0xf5, 0x0e, 0x31, 0x23, 0x51, 0xf6, 0x91, 0xa0, 0x5d, 0xcc,
	0x92, 0x12, 0xe4, 0x6e, 0x67, 0xc4, 0x31, 0x9a, 0x84, 0x40, 0xb1, 0xe5, 0x33, 0x66, 0x09, 0xea,
	0xac, 0xfb, 0x41, 0xfe, 0xc2, 0xb3, 0xda, 0x14, 0x73, 0xd2, 0x8d, 0xef, 0x3a, 0x98, 0x57, 0x52,
	0x8f, 0x31, 0x2c, 0xc8, 0xaa, 0xd7, 0x73, 0x1d, 0x2c, 0x4a, 0xac, 0xc5, 0xb9, 0x75, 0x84, 0x25,
	0x29, 0xba, 0x1e, 0x15, 0x88, 0xb2, 0x92, 0x3f, 0xc0, 0xff, 0xe4, 0xf5, 0x8b, 0xae, 0xef, 0xd9,
	0x48, 0x64, 0x29, 0x7a, 0x1d, 0x46, 0x71, 0x57, 0x12, 0x6d, 0x57, 0xe0, 0xff, 0xa4, 0x04, 0xe0,
	0x7a, 0xe2, 0xa0, 0x4f, 0x5b, 0xc2, 0xe7, 0xb8, 0xd0, 0x48, 0x11, 0x4c, 0xdf, 0x75, 0xd6, 0xe7,
	0xcf, 0x5a, 0x4d, 0xcf, 0xee, 0xe0, 0x4e, 0xed, 0x83, 0x06, 0xb9, 0xad, 0xb7, 0xa1, 0x8c, 0xf8,
	0x1e, 0xc5, 0x84, 0xdc, 0x8a, 0xcc, 0x7b, 0x48, 0x39, 0x6a, 0x32, 0x7c, 0xb7, 0x6d, 0x31, 0x26,
	0x77, 0x97, 0x94, 0xe1, 0x6d, 0xf7, 0x50, 0xd6, 0x6a, 0x5f, 0x9c, 0x5a, 0x0c, 0x0d, 0xb2, 0x0b,
	0xe8, 0xf8, 0x3d, 0x9b, 0xd1, 0x93, 0x0e, 0xa7, 0x2d, 0xb7, 0xeb, 0xfa, 0x1e, 0xa6, 0x25, 0xa6,
	0x6f, 0xf1, 0xd6, 0x73, 0x8b, 0x63, 0x46, 0x36, 0xab, 0x2a, 0x2b, 0x2d, 0xbf, 0x54, 0xa5, 0x29,
	0x69, 0x7d, 0x8b, 0x4b, 0xd7, 0x50, 0xd3, 0xb3, 0x3a, 0xea, 0x4f, 0xf5, 0xdf, 0x1f, 0xf7, 0x35,
	0x7b, 0x7f, 0xf1, 0xb3, 0x92, 0x58, 0x2c, 0x2b, 0xda, 0x97, 0x65, 0x45, 0xfb, 0xba, 0xac, 0x68,
	0x3f, 0x96, 0x15, 0xed, 0xfd, 0xaf, 0x4a, 0xe2, 0xd8, 0x50, 0x2f, 0xe5, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x08, 0x97, 0x89, 0x48, 0x04, 0x00, 0x00,
}
