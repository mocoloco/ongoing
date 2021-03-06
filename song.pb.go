// Code generated by protoc-gen-go.
// source: song.proto
// DO NOT EDIT!

/*
Package song is a generated protocol buffer package.

option go_package = "song";

It is generated from these files:
	song.proto

It has these top-level messages:
	SongObj
	Source
	AudioProperties
	FileInfo
	GroomedSong
	Stats
	Tags
	SongQuery
	SongResponce
*/
package song

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SongObj_MediaType int32

const (
	// TODO - Check in taglib code if there are any additional types/suffixes
	SongObj_Unknown   SongObj_MediaType = 0
	SongObj_Asf       SongObj_MediaType = 1
	SongObj_Flac      SongObj_MediaType = 2
	SongObj_Mp4       SongObj_MediaType = 3
	SongObj_Mpc       SongObj_MediaType = 4
	SongObj_Mpeg      SongObj_MediaType = 5
	SongObj_OggFlac   SongObj_MediaType = 6
	SongObj_OggSpeex  SongObj_MediaType = 7
	SongObj_OggVorbis SongObj_MediaType = 8
	SongObj_Aiff      SongObj_MediaType = 9
	SongObj_Wav       SongObj_MediaType = 10
	SongObj_TrueAudio SongObj_MediaType = 11
	SongObj_Cdda      SongObj_MediaType = 12
	SongObj_OggOpus   SongObj_MediaType = 13
	SongObj_Stream    SongObj_MediaType = 99
)

var SongObj_MediaType_name = map[int32]string{
	0:  "Unknown",
	1:  "Asf",
	2:  "Flac",
	3:  "Mp4",
	4:  "Mpc",
	5:  "Mpeg",
	6:  "OggFlac",
	7:  "OggSpeex",
	8:  "OggVorbis",
	9:  "Aiff",
	10: "Wav",
	11: "TrueAudio",
	12: "Cdda",
	13: "OggOpus",
	99: "Stream",
}
var SongObj_MediaType_value = map[string]int32{
	"Unknown":   0,
	"Asf":       1,
	"Flac":      2,
	"Mp4":       3,
	"Mpc":       4,
	"Mpeg":      5,
	"OggFlac":   6,
	"OggSpeex":  7,
	"OggVorbis": 8,
	"Aiff":      9,
	"Wav":       10,
	"TrueAudio": 11,
	"Cdda":      12,
	"OggOpus":   13,
	"Stream":    99,
}

func (x SongObj_MediaType) String() string {
	return proto.EnumName(SongObj_MediaType_name, int32(x))
}
func (SongObj_MediaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type SongQuery_SearchType int32

const (
	SongQuery_Regex    SongQuery_SearchType = 0
	SongQuery_Wildcard SongQuery_SearchType = 1
	SongQuery_Contains SongQuery_SearchType = 2
	SongQuery_Match    SongQuery_SearchType = 3
)

var SongQuery_SearchType_name = map[int32]string{
	0: "Regex",
	1: "Wildcard",
	2: "Contains",
	3: "Match",
}
var SongQuery_SearchType_value = map[string]int32{
	"Regex":    0,
	"Wildcard": 1,
	"Contains": 2,
	"Match":    3,
}

func (x SongQuery_SearchType) String() string {
	return proto.EnumName(SongQuery_SearchType_name, int32(x))
}
func (SongQuery_SearchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type SongQuery_SearchFiled int32

const (
	SongQuery_Title  SongQuery_SearchFiled = 0
	SongQuery_Artist SongQuery_SearchFiled = 1
	SongQuery_Album  SongQuery_SearchFiled = 2
)

var SongQuery_SearchFiled_name = map[int32]string{
	0: "Title",
	1: "Artist",
	2: "Album",
}
var SongQuery_SearchFiled_value = map[string]int32{
	"Title":  0,
	"Artist": 1,
	"Album":  2,
}

func (x SongQuery_SearchFiled) String() string {
	return proto.EnumName(SongQuery_SearchFiled_name, int32(x))
}
func (SongQuery_SearchFiled) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

// The request message containing the user's name.
type SongObj struct {
	Id              string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Uri             string            `protobuf:"bytes,2,opt,name=uri" json:"uri,omitempty"`
	Source          *Source           `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	MediaType       SongObj_MediaType `protobuf:"varint,4,opt,name=media_type,enum=song.SongObj_MediaType" json:"media_type,omitempty"`
	Unavailable     bool              `protobuf:"varint,5,opt,name=unavailable" json:"unavailable,omitempty"`
	Valid           bool              `protobuf:"varint,6,opt,name=valid" json:"valid,omitempty"`
	Tags            *Tags             `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	AudioProperties *AudioProperties  `protobuf:"bytes,8,opt,name=audio_properties" json:"audio_properties,omitempty"`
	FileInfo        *FileInfo         `protobuf:"bytes,9,opt,name=file_info" json:"file_info,omitempty"`
	TagStats        *Stats            `protobuf:"bytes,10,opt,name=tag_stats" json:"tag_stats,omitempty"`
	Data            []byte            `protobuf:"bytes,11,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SongObj) Reset()                    { *m = SongObj{} }
func (m *SongObj) String() string            { return proto.CompactTextString(m) }
func (*SongObj) ProtoMessage()               {}
func (*SongObj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SongObj) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *SongObj) GetTags() *Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SongObj) GetAudioProperties() *AudioProperties {
	if m != nil {
		return m.AudioProperties
	}
	return nil
}

func (m *SongObj) GetFileInfo() *FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *SongObj) GetTagStats() *Stats {
	if m != nil {
		return m.TagStats
	}
	return nil
}

type Source struct {
	Uri  string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Type int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AudioProperties struct {
	Bitrate    int32 `protobuf:"varint,1,opt,name=bitrate" json:"bitrate,omitempty"`
	Samplerate int32 `protobuf:"varint,2,opt,name=samplerate" json:"samplerate,omitempty"`
	Channels   int32 `protobuf:"varint,3,opt,name=channels" json:"channels,omitempty"`
	Length     int32 `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
}

func (m *AudioProperties) Reset()                    { *m = AudioProperties{} }
func (m *AudioProperties) String() string            { return proto.CompactTextString(m) }
func (*AudioProperties) ProtoMessage()               {}
func (*AudioProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type FileInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Suffix   string `protobuf:"bytes,3,opt,name=suffix" json:"suffix,omitempty"`
	Filesize int64  `protobuf:"varint,4,opt,name=filesize" json:"filesize,omitempty"`
	Mtime    string `protobuf:"bytes,5,opt,name=mtime" json:"mtime,omitempty"`
	Ctime    string `protobuf:"bytes,6,opt,name=ctime" json:"ctime,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GroomedSong struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Artist  string   `protobuf:"bytes,3,opt,name=artist" json:"artist,omitempty"`
	Album   string   `protobuf:"bytes,4,opt,name=album" json:"album,omitempty"`
	Songids []string `protobuf:"bytes,5,rep,name=songids" json:"songids,omitempty"`
}

func (m *GroomedSong) Reset()                    { *m = GroomedSong{} }
func (m *GroomedSong) String() string            { return proto.CompactTextString(m) }
func (*GroomedSong) ProtoMessage()               {}
func (*GroomedSong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Stats struct {
	Playcount  int32  `protobuf:"varint,1,opt,name=playcount" json:"playcount,omitempty"`
	Lastplayed string `protobuf:"bytes,2,opt,name=lastplayed" json:"lastplayed,omitempty"`
	Skipcount  int32  `protobuf:"varint,3,opt,name=skipcount" json:"skipcount,omitempty"`
	Rating     int32  `protobuf:"varint,4,opt,name=rating" json:"rating,omitempty"`
	Score      int32  `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Tags struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Artist  string `protobuf:"bytes,2,opt,name=artist" json:"artist,omitempty"`
	Album   string `protobuf:"bytes,3,opt,name=album" json:"album,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	Genre   string `protobuf:"bytes,5,opt,name=genre" json:"genre,omitempty"`
	Year    int32  `protobuf:"varint,6,opt,name=year" json:"year,omitempty"`
	Track   int32  `protobuf:"varint,7,opt,name=track" json:"track,omitempty"`
}

func (m *Tags) Reset()                    { *m = Tags{} }
func (m *Tags) String() string            { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()               {}
func (*Tags) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type SongQuery struct {
	Search      string                `protobuf:"bytes,1,opt,name=search" json:"search,omitempty"`
	SearchType  SongQuery_SearchType  `protobuf:"varint,2,opt,name=search_type,enum=song.SongQuery_SearchType" json:"search_type,omitempty"`
	SearchField SongQuery_SearchFiled `protobuf:"varint,3,opt,name=search_field,enum=song.SongQuery_SearchFiled" json:"search_field,omitempty"`
}

func (m *SongQuery) Reset()                    { *m = SongQuery{} }
func (m *SongQuery) String() string            { return proto.CompactTextString(m) }
func (*SongQuery) ProtoMessage()               {}
func (*SongQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// The response message containing the song action status
type SongResponce struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SongResponce) Reset()                    { *m = SongResponce{} }
func (m *SongResponce) String() string            { return proto.CompactTextString(m) }
func (*SongResponce) ProtoMessage()               {}
func (*SongResponce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*SongObj)(nil), "song.SongObj")
	proto.RegisterType((*Source)(nil), "song.Source")
	proto.RegisterType((*AudioProperties)(nil), "song.AudioProperties")
	proto.RegisterType((*FileInfo)(nil), "song.FileInfo")
	proto.RegisterType((*GroomedSong)(nil), "song.GroomedSong")
	proto.RegisterType((*Stats)(nil), "song.Stats")
	proto.RegisterType((*Tags)(nil), "song.Tags")
	proto.RegisterType((*SongQuery)(nil), "song.SongQuery")
	proto.RegisterType((*SongResponce)(nil), "song.SongResponce")
	proto.RegisterEnum("song.SongObj_MediaType", SongObj_MediaType_name, SongObj_MediaType_value)
	proto.RegisterEnum("song.SongQuery_SearchType", SongQuery_SearchType_name, SongQuery_SearchType_value)
	proto.RegisterEnum("song.SongQuery_SearchFiled", SongQuery_SearchFiled_name, SongQuery_SearchFiled_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for SongSrv service

type SongSrvClient interface {
	Add(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error)
	Adds(ctx context.Context, opts ...grpc.CallOption) (SongSrv_AddsClient, error)
	Modify(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error)
	Delete(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error)
	Get(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongObj, error)
	Filter(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (SongSrv_FilterClient, error)
}

type songSrvClient struct {
	cc *grpc.ClientConn
}

func NewSongSrvClient(cc *grpc.ClientConn) SongSrvClient {
	return &songSrvClient{cc}
}

func (c *songSrvClient) Add(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error) {
	out := new(SongResponce)
	err := grpc.Invoke(ctx, "/song.SongSrv/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songSrvClient) Adds(ctx context.Context, opts ...grpc.CallOption) (SongSrv_AddsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SongSrv_serviceDesc.Streams[0], c.cc, "/song.SongSrv/Adds", opts...)
	if err != nil {
		return nil, err
	}
	x := &songSrvAddsClient{stream}
	return x, nil
}

type SongSrv_AddsClient interface {
	Send(*SongObj) error
	Recv() (*SongResponce, error)
	grpc.ClientStream
}

type songSrvAddsClient struct {
	grpc.ClientStream
}

func (x *songSrvAddsClient) Send(m *SongObj) error {
	return x.ClientStream.SendMsg(m)
}

func (x *songSrvAddsClient) Recv() (*SongResponce, error) {
	m := new(SongResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *songSrvClient) Modify(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error) {
	out := new(SongResponce)
	err := grpc.Invoke(ctx, "/song.SongSrv/Modify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songSrvClient) Delete(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongResponce, error) {
	out := new(SongResponce)
	err := grpc.Invoke(ctx, "/song.SongSrv/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songSrvClient) Get(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongObj, error) {
	out := new(SongObj)
	err := grpc.Invoke(ctx, "/song.SongSrv/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songSrvClient) Filter(ctx context.Context, in *SongQuery, opts ...grpc.CallOption) (SongSrv_FilterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SongSrv_serviceDesc.Streams[1], c.cc, "/song.SongSrv/Filter", opts...)
	if err != nil {
		return nil, err
	}
	x := &songSrvFilterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SongSrv_FilterClient interface {
	Recv() (*SongObj, error)
	grpc.ClientStream
}

type songSrvFilterClient struct {
	grpc.ClientStream
}

func (x *songSrvFilterClient) Recv() (*SongObj, error) {
	m := new(SongObj)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SongSrv service

type SongSrvServer interface {
	Add(context.Context, *SongObj) (*SongResponce, error)
	Adds(SongSrv_AddsServer) error
	Modify(context.Context, *SongObj) (*SongResponce, error)
	Delete(context.Context, *SongObj) (*SongResponce, error)
	Get(context.Context, *SongObj) (*SongObj, error)
	Filter(*SongQuery, SongSrv_FilterServer) error
}

func RegisterSongSrvServer(s *grpc.Server, srv SongSrvServer) {
	s.RegisterService(&_SongSrv_serviceDesc, srv)
}

func _SongSrv_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SongObj)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SongSrvServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SongSrv_Adds_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SongSrvServer).Adds(&songSrvAddsServer{stream})
}

type SongSrv_AddsServer interface {
	Send(*SongResponce) error
	Recv() (*SongObj, error)
	grpc.ServerStream
}

type songSrvAddsServer struct {
	grpc.ServerStream
}

func (x *songSrvAddsServer) Send(m *SongResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *songSrvAddsServer) Recv() (*SongObj, error) {
	m := new(SongObj)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SongSrv_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SongObj)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SongSrvServer).Modify(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SongSrv_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SongObj)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SongSrvServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SongSrv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SongObj)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SongSrvServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _SongSrv_Filter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SongQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SongSrvServer).Filter(m, &songSrvFilterServer{stream})
}

type SongSrv_FilterServer interface {
	Send(*SongObj) error
	grpc.ServerStream
}

type songSrvFilterServer struct {
	grpc.ServerStream
}

func (x *songSrvFilterServer) Send(m *SongObj) error {
	return x.ServerStream.SendMsg(m)
}

var _SongSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "song.SongSrv",
	HandlerType: (*SongSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _SongSrv_Add_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _SongSrv_Modify_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SongSrv_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SongSrv_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Adds",
			Handler:       _SongSrv_Adds_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Filter",
			Handler:       _SongSrv_Filter_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 974 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xce, 0xd8, 0x1e, 0xdb, 0x53, 0x76, 0x92, 0x51, 0x8b, 0x9f, 0x51, 0x40, 0xab, 0xd5, 0x08,
	0xa4, 0x88, 0x9f, 0xec, 0x2a, 0x20, 0x2e, 0x48, 0x2b, 0xc2, 0xa2, 0xac, 0x38, 0x44, 0x81, 0x76,
	0x60, 0x8f, 0x56, 0x7b, 0xa6, 0xc7, 0x69, 0x32, 0x7f, 0x9a, 0xee, 0x09, 0x6b, 0x2e, 0x3c, 0xc3,
	0x72, 0xe3, 0xce, 0x33, 0xf0, 0x0a, 0xbc, 0x16, 0x55, 0xdd, 0x63, 0x3b, 0xf1, 0xe6, 0x90, 0x5b,
	0x7d, 0x5f, 0x7f, 0x55, 0x53, 0x5d, 0x55, 0x5d, 0x03, 0xa0, 0xab, 0x72, 0x79, 0x52, 0x37, 0x95,
	0xa9, 0xd8, 0x80, 0xec, 0xf8, 0xbf, 0x01, 0x8c, 0x66, 0x68, 0x5c, 0x2e, 0x7e, 0x63, 0x07, 0xd0,
	0x53, 0x69, 0xe4, 0x3d, 0xf5, 0x8e, 0x03, 0x8e, 0x16, 0x0b, 0xa1, 0xdf, 0x36, 0x2a, 0xea, 0x59,
	0x82, 0x4c, 0xf6, 0x09, 0x0c, 0x75, 0xd5, 0x36, 0x89, 0x8c, 0xfa, 0x48, 0x4e, 0x4e, 0xa7, 0x27,
	0x36, 0xe0, 0xcc, 0x72, 0xbc, 0x3b, 0x63, 0xdf, 0x00, 0x14, 0x32, 0x55, 0x62, 0x6e, 0x56, 0xb5,
	0x8c, 0x06, 0xa8, 0x3c, 0x38, 0xfd, 0x70, 0xad, 0xb4, 0x9f, 0x3a, 0xb9, 0xa0, 0xf3, 0x2b, 0x3c,
	0xe6, 0x41, 0xb1, 0x36, 0xd9, 0x53, 0x98, 0xb4, 0xa5, 0xb8, 0x15, 0x2a, 0x17, 0x8b, 0x5c, 0x46,
	0x3e, 0x3a, 0x8e, 0xf9, 0x5d, 0x8a, 0xbd, 0x07, 0xfe, 0xad, 0xc8, 0x31, 0xc9, 0xa1, 0x3d, 0x73,
	0x80, 0x3d, 0x81, 0x81, 0x11, 0x4b, 0x1d, 0x8d, 0x6c, 0x4e, 0xe0, 0xbe, 0x74, 0x85, 0x0c, 0xb7,
	0x3c, 0xfb, 0x0e, 0x42, 0xd1, 0xa6, 0xaa, 0x9a, 0xe3, 0xc5, 0x6b, 0xd9, 0x18, 0x25, 0x75, 0x34,
	0xb6, 0xda, 0xf7, 0x9d, 0xf6, 0x8c, 0x4e, 0x7f, 0xda, 0x1c, 0xf2, 0x43, 0x71, 0x9f, 0x60, 0x9f,
	0x43, 0x90, 0xa9, 0x5c, 0xce, 0x55, 0x99, 0x55, 0x51, 0x60, 0x5d, 0x0f, 0x9c, 0xeb, 0x39, 0xd2,
	0x3f, 0x22, 0xcb, 0xc7, 0x59, 0x67, 0xb1, 0x63, 0x08, 0xf0, 0xb3, 0x73, 0x6d, 0x84, 0xd1, 0x11,
	0x58, 0xf1, 0xa4, 0xbb, 0x3d, 0x51, 0x7c, 0x8c, 0xa7, 0xd6, 0x62, 0x0c, 0x06, 0xa9, 0x30, 0x22,
	0x9a, 0xa0, 0x68, 0xca, 0xad, 0x1d, 0xff, 0xeb, 0x41, 0xb0, 0xa9, 0x0e, 0x9b, 0xc0, 0xe8, 0x97,
	0xf2, 0xa6, 0xac, 0x7e, 0x2f, 0xc3, 0x3d, 0x36, 0x82, 0xfe, 0x99, 0xce, 0x42, 0x8f, 0x8d, 0x61,
	0x70, 0x9e, 0x8b, 0x24, 0xec, 0x11, 0x75, 0x51, 0x7f, 0x1d, 0xf6, 0x9d, 0x91, 0x84, 0x03, 0x3a,
	0xbb, 0xa8, 0xe5, 0x32, 0xf4, 0xc9, 0xf7, 0x72, 0xb9, 0xb4, 0xc2, 0x21, 0x9b, 0xc2, 0x18, 0xc1,
	0xac, 0x96, 0xf2, 0x4d, 0x38, 0x62, 0xfb, 0x10, 0x20, 0xfa, 0xb5, 0x6a, 0x16, 0x4a, 0x87, 0x63,
	0xf2, 0x39, 0x53, 0x59, 0x16, 0x06, 0x14, 0xe6, 0xb5, 0xb8, 0x0d, 0x81, 0x14, 0x57, 0x4d, 0x2b,
	0x6d, 0x65, 0xc2, 0x09, 0x29, 0x5e, 0xa6, 0xa9, 0x08, 0xa7, 0x5d, 0xd4, 0xcb, 0xba, 0xd5, 0xe1,
	0x3e, 0x03, 0x18, 0xce, 0x4c, 0x23, 0x45, 0x11, 0x26, 0xf1, 0xf7, 0x68, 0xbb, 0xfe, 0x77, 0x73,
	0xe3, 0x6d, 0xe7, 0x06, 0x2f, 0x6a, 0x67, 0x81, 0x46, 0xc9, 0xe7, 0xd6, 0x26, 0xae, 0x14, 0x85,
	0x9b, 0xa4, 0x80, 0x5b, 0x3b, 0xfe, 0x13, 0x0e, 0x77, 0x7a, 0xc1, 0x22, 0x18, 0x2d, 0x94, 0x69,
	0x84, 0x91, 0x36, 0xa0, 0xcf, 0xd7, 0x10, 0xdb, 0x0e, 0x5a, 0x14, 0x75, 0x2e, 0xed, 0xa1, 0x0b,
	0x7d, 0x87, 0x61, 0x47, 0x30, 0x4e, 0xae, 0x45, 0x59, 0xca, 0x5c, 0xdb, 0x8f, 0xf8, 0x7c, 0x83,
	0xd9, 0x07, 0x30, 0xcc, 0x65, 0xb9, 0x34, 0xd7, 0x76, 0x3c, 0x7d, 0xde, 0xa1, 0xf8, 0x2f, 0x0f,
	0xc6, 0xeb, 0x96, 0x6e, 0x32, 0xf4, 0xb6, 0x19, 0x12, 0x57, 0x0b, 0x74, 0x73, 0x8f, 0xc2, 0xda,
	0x14, 0x4c, 0xb7, 0x59, 0xa6, 0xde, 0x74, 0x77, 0xe9, 0x10, 0x25, 0x40, 0x43, 0xa1, 0xd5, 0x1f,
	0xee, 0x15, 0xf4, 0xf9, 0x06, 0xd3, 0x24, 0x17, 0x46, 0x15, 0x6e, 0xca, 0x03, 0xee, 0x00, 0xb1,
	0x89, 0x65, 0x87, 0x8e, 0xb5, 0x00, 0xab, 0x32, 0x79, 0xd5, 0x54, 0x15, 0x3e, 0x14, 0x7a, 0x3e,
	0x0f, 0xa6, 0x85, 0x8e, 0x46, 0x99, 0x5c, 0x76, 0x79, 0x39, 0x40, 0x89, 0x09, 0xac, 0xa2, 0x36,
	0xeb, 0xc4, 0x1c, 0x22, 0xb5, 0xc8, 0x17, 0x6d, 0x61, 0xb3, 0x42, 0xb5, 0x05, 0x54, 0x69, 0x9a,
	0x52, 0x95, 0x6a, 0x4c, 0xaa, 0x8f, 0xfc, 0x1a, 0xc6, 0x6f, 0x3d, 0xf0, 0xdd, 0xc4, 0x7e, 0x0c,
	0x41, 0x9d, 0x8b, 0x55, 0x52, 0xb5, 0xa5, 0xe9, 0xfa, 0xb1, 0x25, 0xa8, 0x23, 0xb9, 0xd0, 0x86,
	0x08, 0x99, 0x76, 0xa9, 0xdc, 0x61, 0xc8, 0x5b, 0xdf, 0xa8, 0xda, 0x79, 0xbb, 0x96, 0x6c, 0x09,
	0xca, 0x16, 0xfb, 0xa6, 0xca, 0xe5, 0xba, 0x27, 0x0e, 0x51, 0xb6, 0x3a, 0xa9, 0x1a, 0x57, 0x2a,
	0x9f, 0x3b, 0x10, 0xff, 0xe3, 0xc1, 0x80, 0xde, 0xf8, 0xf6, 0xea, 0xde, 0xc3, 0x57, 0xef, 0x3d,
	0x7c, 0xf5, 0xfe, 0xce, 0xd5, 0x93, 0xaa, 0x28, 0x24, 0xa6, 0xe5, 0x4a, 0xb2, 0x86, 0xa4, 0x5f,
	0xca, 0xb2, 0xd9, 0xf4, 0xc9, 0x02, 0x6a, 0xc1, 0x4a, 0x8a, 0xc6, 0xb6, 0x09, 0xe7, 0x99, 0x6c,
	0x9b, 0x47, 0x23, 0x92, 0x1b, 0xbb, 0x86, 0x30, 0x4d, 0x0b, 0xe2, 0xb7, 0x3d, 0x08, 0xa8, 0x6b,
	0x3f, 0xb7, 0xb2, 0x59, 0xd9, 0x49, 0x41, 0x6d, 0x72, 0xdd, 0x25, 0xdb, 0x21, 0xf6, 0x2d, 0x4c,
	0x9c, 0x35, 0xdf, 0x3c, 0x93, 0x83, 0xd3, 0xa3, 0xed, 0xca, 0xb4, 0xde, 0x27, 0x33, 0x2b, 0xb1,
	0x5b, 0x13, 0xf4, 0xc6, 0x66, 0x2f, 0x60, 0xda, 0x39, 0x67, 0x4a, 0xe6, 0xa9, 0xbd, 0xd9, 0xc1,
	0xe9, 0x47, 0x0f, 0x7b, 0xd3, 0x70, 0xa7, 0xbc, 0xfb, 0xda, 0x39, 0xe9, 0xe3, 0x17, 0x00, 0xdb,
	0xc8, 0x2c, 0x00, 0x9f, 0xcb, 0x25, 0x6e, 0x89, 0x3d, 0xda, 0x19, 0xaf, 0x55, 0x9e, 0x26, 0xa2,
	0x49, 0x71, 0xe9, 0x20, 0x7a, 0x59, 0x95, 0x46, 0xa8, 0x52, 0xe3, 0xe2, 0x41, 0xd9, 0x85, 0x30,
	0xc9, 0x75, 0xd8, 0x8f, 0x9f, 0xc1, 0xe4, 0x4e, 0x6c, 0x3a, 0xb9, 0xa2, 0x16, 0x60, 0x00, 0x5c,
	0x0f, 0x67, 0xb6, 0xec, 0xe8, 0x8e, 0xf4, 0x19, 0xd5, 0x3a, 0xec, 0xc5, 0x4f, 0x60, 0x4a, 0x69,
	0x71, 0xa9, 0xeb, 0xaa, 0xc4, 0x7d, 0xb1, 0xf3, 0xdf, 0x39, 0xfd, 0xbb, 0xe7, 0xfe, 0x49, 0xb3,
	0xe6, 0x96, 0x7d, 0x86, 0x3b, 0x2f, 0x4d, 0xd9, 0xfe, 0xbd, 0xdf, 0xc7, 0x11, 0xdb, 0xc2, 0x75,
	0x94, 0x78, 0x8f, 0x3d, 0xc3, 0x35, 0x96, 0xa6, 0xfa, 0x51, 0xe2, 0x63, 0xef, 0xb9, 0xc7, 0xbe,
	0x84, 0xe1, 0x45, 0x95, 0xaa, 0x6c, 0xf5, 0xb8, 0xf8, 0x28, 0xff, 0x41, 0xe6, 0x12, 0x57, 0xcb,
	0xa3, 0xe4, 0x9f, 0x42, 0xff, 0x95, 0x34, 0xbb, 0xda, 0xfb, 0x10, 0x65, 0x5f, 0xc0, 0x10, 0x0b,
	0x67, 0x64, 0xc3, 0x0e, 0x77, 0x5a, 0xf6, 0x8e, 0xf6, 0xb9, 0xb7, 0x18, 0xda, 0x9f, 0xf7, 0x57,
	0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x03, 0xcf, 0x39, 0x8a, 0xca, 0x07, 0x00, 0x00,
}
