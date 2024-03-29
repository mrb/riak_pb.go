// Code generated by protoc-gen-go from "riak_search.proto"
// DO NOT EDIT!

package riak_search

import proto "code.google.com/p/goprotobuf/proto"
import "math"
import riak "riak.pb"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf

type RpbSearchDoc struct {
	Fields           []*riak.RpbPair `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbSearchDoc) Reset()         { *this = RpbSearchDoc{} }
func (this *RpbSearchDoc) String() string { return proto.CompactTextString(this) }

type RpbSearchQueryReq struct {
	Q                []byte   `protobuf:"bytes,1,req,name=q" json:"q,omitempty"`
	Index            []byte   `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Rows             *uint32  `protobuf:"varint,3,opt,name=rows" json:"rows,omitempty"`
	Start            *uint32  `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	Sort             []byte   `protobuf:"bytes,5,opt,name=sort" json:"sort,omitempty"`
	Filter           []byte   `protobuf:"bytes,6,opt,name=filter" json:"filter,omitempty"`
	Df               []byte   `protobuf:"bytes,7,opt,name=df" json:"df,omitempty"`
	Op               []byte   `protobuf:"bytes,8,opt,name=op" json:"op,omitempty"`
	Fl               [][]byte `protobuf:"bytes,9,rep,name=fl" json:"fl,omitempty"`
	Presort          []byte   `protobuf:"bytes,10,opt,name=presort" json:"presort,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *RpbSearchQueryReq) Reset()         { *this = RpbSearchQueryReq{} }
func (this *RpbSearchQueryReq) String() string { return proto.CompactTextString(this) }

type RpbSearchQueryResp struct {
	Docs             []*RpbSearchDoc `protobuf:"bytes,1,rep,name=docs" json:"docs,omitempty"`
	MaxScore         *float32        `protobuf:"fixed32,2,opt,name=max_score" json:"max_score,omitempty"`
	NumFound         *uint32         `protobuf:"varint,3,opt,name=num_found" json:"num_found,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *RpbSearchQueryResp) Reset()         { *this = RpbSearchQueryResp{} }
func (this *RpbSearchQueryResp) String() string { return proto.CompactTextString(this) }

func init() {
}
