package models

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GraphiteTagDelSeries) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Paths":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Paths")
				return
			}
			if cap(z.Paths) >= int(zb0002) {
				z.Paths = (z.Paths)[:zb0002]
			} else {
				z.Paths = make([]string, zb0002)
			}
			for za0001 := range z.Paths {
				z.Paths[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Paths", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GraphiteTagDelSeries) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Paths"
	err = en.Append(0x81, 0xa5, 0x50, 0x61, 0x74, 0x68, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Paths)))
	if err != nil {
		err = msgp.WrapError(err, "Paths")
		return
	}
	for za0001 := range z.Paths {
		err = en.WriteString(z.Paths[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Paths", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GraphiteTagDelSeries) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Paths"
	o = append(o, 0x81, 0xa5, 0x50, 0x61, 0x74, 0x68, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Paths)))
	for za0001 := range z.Paths {
		o = msgp.AppendString(o, z.Paths[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GraphiteTagDelSeries) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Paths":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Paths")
				return
			}
			if cap(z.Paths) >= int(zb0002) {
				z.Paths = (z.Paths)[:zb0002]
			} else {
				z.Paths = make([]string, zb0002)
			}
			for za0001 := range z.Paths {
				z.Paths[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Paths", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GraphiteTagDelSeries) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Paths {
		s += msgp.StringPrefixSize + len(z.Paths[za0001])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GraphiteTagDelSeriesResp) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Count":
			z.Count, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Count")
				return
			}
		case "Peers":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Peers")
				return
			}
			if z.Peers == nil {
				z.Peers = make(map[string]int, zb0002)
			} else if len(z.Peers) > 0 {
				for key := range z.Peers {
					delete(z.Peers, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 int
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Peers")
					return
				}
				za0002, err = dc.ReadInt()
				if err != nil {
					err = msgp.WrapError(err, "Peers", za0001)
					return
				}
				z.Peers[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GraphiteTagDelSeriesResp) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Count"
	err = en.Append(0x82, 0xa5, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Count)
	if err != nil {
		err = msgp.WrapError(err, "Count")
		return
	}
	// write "Peers"
	err = en.Append(0xa5, 0x50, 0x65, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Peers)))
	if err != nil {
		err = msgp.WrapError(err, "Peers")
		return
	}
	for za0001, za0002 := range z.Peers {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Peers")
			return
		}
		err = en.WriteInt(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Peers", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GraphiteTagDelSeriesResp) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Count"
	o = append(o, 0x82, 0xa5, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.Count)
	// string "Peers"
	o = append(o, 0xa5, 0x50, 0x65, 0x65, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Peers)))
	for za0001, za0002 := range z.Peers {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendInt(o, za0002)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GraphiteTagDelSeriesResp) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Count":
			z.Count, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Count")
				return
			}
		case "Peers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Peers")
				return
			}
			if z.Peers == nil {
				z.Peers = make(map[string]int, zb0002)
			} else if len(z.Peers) > 0 {
				for key := range z.Peers {
					delete(z.Peers, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 int
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Peers")
					return
				}
				za0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Peers", za0001)
					return
				}
				z.Peers[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GraphiteTagDelSeriesResp) Msgsize() (s int) {
	s = 1 + 6 + msgp.IntSize + 6 + msgp.MapHeaderSize
	if z.Peers != nil {
		for za0001, za0002 := range z.Peers {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.IntSize
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GraphiteTagTerms) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tags":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tags")
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]string, zb0002)
			}
			for za0001 := range z.Tags {
				z.Tags[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Tags", za0001)
					return
				}
			}
		case "Expr":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Expr")
				return
			}
			if cap(z.Expr) >= int(zb0003) {
				z.Expr = (z.Expr)[:zb0003]
			} else {
				z.Expr = make([]string, zb0003)
			}
			for za0002 := range z.Expr {
				z.Expr[za0002], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Expr", za0002)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GraphiteTagTerms) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Tags"
	err = en.Append(0x82, 0xa4, 0x54, 0x61, 0x67, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Tags)))
	if err != nil {
		err = msgp.WrapError(err, "Tags")
		return
	}
	for za0001 := range z.Tags {
		err = en.WriteString(z.Tags[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Tags", za0001)
			return
		}
	}
	// write "Expr"
	err = en.Append(0xa4, 0x45, 0x78, 0x70, 0x72)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Expr)))
	if err != nil {
		err = msgp.WrapError(err, "Expr")
		return
	}
	for za0002 := range z.Expr {
		err = en.WriteString(z.Expr[za0002])
		if err != nil {
			err = msgp.WrapError(err, "Expr", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GraphiteTagTerms) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Tags"
	o = append(o, 0x82, 0xa4, 0x54, 0x61, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tags)))
	for za0001 := range z.Tags {
		o = msgp.AppendString(o, z.Tags[za0001])
	}
	// string "Expr"
	o = append(o, 0xa4, 0x45, 0x78, 0x70, 0x72)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Expr)))
	for za0002 := range z.Expr {
		o = msgp.AppendString(o, z.Expr[za0002])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GraphiteTagTerms) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tags":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tags")
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]string, zb0002)
			}
			for za0001 := range z.Tags {
				z.Tags[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tags", za0001)
					return
				}
			}
		case "Expr":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Expr")
				return
			}
			if cap(z.Expr) >= int(zb0003) {
				z.Expr = (z.Expr)[:zb0003]
			} else {
				z.Expr = make([]string, zb0003)
			}
			for za0002 := range z.Expr {
				z.Expr[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Expr", za0002)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GraphiteTagTerms) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Tags {
		s += msgp.StringPrefixSize + len(z.Tags[za0001])
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Expr {
		s += msgp.StringPrefixSize + len(z.Expr[za0002])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *GraphiteTagTermsResp) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "TotalSeries":
			z.TotalSeries, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "TotalSeries")
				return
			}
		case "Terms":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Terms")
				return
			}
			if z.Terms == nil {
				z.Terms = make(map[string]map[string]uint32, zb0002)
			} else if len(z.Terms) > 0 {
				for key := range z.Terms {
					delete(z.Terms, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 map[string]uint32
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Terms")
					return
				}
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Terms", za0001)
					return
				}
				if za0002 == nil {
					za0002 = make(map[string]uint32, zb0003)
				} else if len(za0002) > 0 {
					for key := range za0002 {
						delete(za0002, key)
					}
				}
				for zb0003 > 0 {
					zb0003--
					var za0003 string
					var za0004 uint32
					za0003, err = dc.ReadString()
					if err != nil {
						err = msgp.WrapError(err, "Terms", za0001)
						return
					}
					za0004, err = dc.ReadUint32()
					if err != nil {
						err = msgp.WrapError(err, "Terms", za0001, za0003)
						return
					}
					za0002[za0003] = za0004
				}
				z.Terms[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GraphiteTagTermsResp) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "TotalSeries"
	err = en.Append(0x82, 0xab, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.TotalSeries)
	if err != nil {
		err = msgp.WrapError(err, "TotalSeries")
		return
	}
	// write "Terms"
	err = en.Append(0xa5, 0x54, 0x65, 0x72, 0x6d, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Terms)))
	if err != nil {
		err = msgp.WrapError(err, "Terms")
		return
	}
	for za0001, za0002 := range z.Terms {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Terms")
			return
		}
		err = en.WriteMapHeader(uint32(len(za0002)))
		if err != nil {
			err = msgp.WrapError(err, "Terms", za0001)
			return
		}
		for za0003, za0004 := range za0002 {
			err = en.WriteString(za0003)
			if err != nil {
				err = msgp.WrapError(err, "Terms", za0001)
				return
			}
			err = en.WriteUint32(za0004)
			if err != nil {
				err = msgp.WrapError(err, "Terms", za0001, za0003)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GraphiteTagTermsResp) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "TotalSeries"
	o = append(o, 0x82, 0xab, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73)
	o = msgp.AppendUint32(o, z.TotalSeries)
	// string "Terms"
	o = append(o, 0xa5, 0x54, 0x65, 0x72, 0x6d, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Terms)))
	for za0001, za0002 := range z.Terms {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendMapHeader(o, uint32(len(za0002)))
		for za0003, za0004 := range za0002 {
			o = msgp.AppendString(o, za0003)
			o = msgp.AppendUint32(o, za0004)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GraphiteTagTermsResp) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "TotalSeries":
			z.TotalSeries, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TotalSeries")
				return
			}
		case "Terms":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Terms")
				return
			}
			if z.Terms == nil {
				z.Terms = make(map[string]map[string]uint32, zb0002)
			} else if len(z.Terms) > 0 {
				for key := range z.Terms {
					delete(z.Terms, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 map[string]uint32
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Terms")
					return
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Terms", za0001)
					return
				}
				if za0002 == nil {
					za0002 = make(map[string]uint32, zb0003)
				} else if len(za0002) > 0 {
					for key := range za0002 {
						delete(za0002, key)
					}
				}
				for zb0003 > 0 {
					var za0003 string
					var za0004 uint32
					zb0003--
					za0003, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Terms", za0001)
						return
					}
					za0004, bts, err = msgp.ReadUint32Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Terms", za0001, za0003)
						return
					}
					za0002[za0003] = za0004
				}
				z.Terms[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GraphiteTagTermsResp) Msgsize() (s int) {
	s = 1 + 12 + msgp.Uint32Size + 6 + msgp.MapHeaderSize
	if z.Terms != nil {
		for za0001, za0002 := range z.Terms {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.MapHeaderSize
			if za0002 != nil {
				for za0003, za0004 := range za0002 {
					_ = za0004
					s += msgp.StringPrefixSize + len(za0003) + msgp.Uint32Size
				}
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SeriesPickle) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(SeriesPickle, zb0002)
	}
	for zb0001 := range *z {
		err = (*z)[zb0001].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z SeriesPickle) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0003 := range z {
		err = z[zb0003].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SeriesPickle) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o, err = z[zb0003].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SeriesPickle) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(SeriesPickle, zb0002)
	}
	for zb0001 := range *z {
		bts, err = (*z)[zb0001].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SeriesPickle) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += z[zb0003].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SeriesPickleItem) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			z.Path, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Path")
				return
			}
		case "isLeaf":
			z.IsLeaf, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsLeaf")
				return
			}
		case "intervals":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Intervals")
				return
			}
			if cap(z.Intervals) >= int(zb0002) {
				z.Intervals = (z.Intervals)[:zb0002]
			} else {
				z.Intervals = make([][]int64, zb0002)
			}
			for za0001 := range z.Intervals {
				var zb0003 uint32
				zb0003, err = dc.ReadArrayHeader()
				if err != nil {
					err = msgp.WrapError(err, "Intervals", za0001)
					return
				}
				if cap(z.Intervals[za0001]) >= int(zb0003) {
					z.Intervals[za0001] = (z.Intervals[za0001])[:zb0003]
				} else {
					z.Intervals[za0001] = make([]int64, zb0003)
				}
				for za0002 := range z.Intervals[za0001] {
					z.Intervals[za0001][za0002], err = dc.ReadInt64()
					if err != nil {
						err = msgp.WrapError(err, "Intervals", za0001, za0002)
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SeriesPickleItem) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "path"
	err = en.Append(0x83, 0xa4, 0x70, 0x61, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteString(z.Path)
	if err != nil {
		err = msgp.WrapError(err, "Path")
		return
	}
	// write "isLeaf"
	err = en.Append(0xa6, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x66)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsLeaf)
	if err != nil {
		err = msgp.WrapError(err, "IsLeaf")
		return
	}
	// write "intervals"
	err = en.Append(0xa9, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Intervals)))
	if err != nil {
		err = msgp.WrapError(err, "Intervals")
		return
	}
	for za0001 := range z.Intervals {
		err = en.WriteArrayHeader(uint32(len(z.Intervals[za0001])))
		if err != nil {
			err = msgp.WrapError(err, "Intervals", za0001)
			return
		}
		for za0002 := range z.Intervals[za0001] {
			err = en.WriteInt64(z.Intervals[za0001][za0002])
			if err != nil {
				err = msgp.WrapError(err, "Intervals", za0001, za0002)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SeriesPickleItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "path"
	o = append(o, 0x83, 0xa4, 0x70, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	// string "isLeaf"
	o = append(o, 0xa6, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x66)
	o = msgp.AppendBool(o, z.IsLeaf)
	// string "intervals"
	o = append(o, 0xa9, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Intervals)))
	for za0001 := range z.Intervals {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Intervals[za0001])))
		for za0002 := range z.Intervals[za0001] {
			o = msgp.AppendInt64(o, z.Intervals[za0001][za0002])
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SeriesPickleItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Path")
				return
			}
		case "isLeaf":
			z.IsLeaf, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsLeaf")
				return
			}
		case "intervals":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Intervals")
				return
			}
			if cap(z.Intervals) >= int(zb0002) {
				z.Intervals = (z.Intervals)[:zb0002]
			} else {
				z.Intervals = make([][]int64, zb0002)
			}
			for za0001 := range z.Intervals {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Intervals", za0001)
					return
				}
				if cap(z.Intervals[za0001]) >= int(zb0003) {
					z.Intervals[za0001] = (z.Intervals[za0001])[:zb0003]
				} else {
					z.Intervals[za0001] = make([]int64, zb0003)
				}
				for za0002 := range z.Intervals[za0001] {
					z.Intervals[za0001][za0002], bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Intervals", za0001, za0002)
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SeriesPickleItem) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Path) + 7 + msgp.BoolSize + 10 + msgp.ArrayHeaderSize
	for za0001 := range z.Intervals {
		s += msgp.ArrayHeaderSize + (len(z.Intervals[za0001]) * (msgp.Int64Size))
	}
	return
}
