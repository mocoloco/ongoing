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
	SongQuery_Name    SongQuery_SearchFiled = 0
	SongQuery_Title   SongQuery_SearchFiled = 2
	SongQuery_Atrists SongQuery_SearchFiled = 3
	SongQuery_Album   SongQuery_SearchFiled = 4
)

var SongQuery_SearchFiled_name = map[int32]string{
	0: "Name",
	2: "Title",
	3: "Atrists",
	4: "Album",
}
var SongQuery_SearchFiled_value = map[string]int32{
	"Name":    0,
	"Title":   2,
	"Atrists": 3,
	"Album":   4,
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
	// rpc Modify (SongObj) returns (SongResponce) {}
	// rpc Delete (SongObj) returns (SongResponce) {}
	Get(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongObj, error)
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

func (c *songSrvClient) Get(ctx context.Context, in *SongObj, opts ...grpc.CallOption) (*SongObj, error) {
	out := new(SongObj)
	err := grpc.Invoke(ctx, "/song.SongSrv/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SongSrv service

type SongSrvServer interface {
	Add(context.Context, *SongObj) (*SongResponce, error)
	Adds(SongSrv_AddsServer) error
	// rpc Modify (SongObj) returns (SongResponce) {}
	// rpc Delete (SongObj) returns (SongResponce) {}
	Get(context.Context, *SongObj) (*SongObj, error)
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

var _SongSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "song.SongSrv",
	HandlerType: (*SongSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _SongSrv_Add_Handler,
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
	},
}

var fileDescriptor0 = []byte{
	// 950 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xce, 0xd8, 0x1e, 0xdb, 0x53, 0xf6, 0x7a, 0x5b, 0x2d, 0x7e, 0x46, 0x01, 0xad, 0x56, 0x23,
	0x90, 0x22, 0x90, 0x02, 0x32, 0x88, 0x03, 0x48, 0x2b, 0xcc, 0x4a, 0x59, 0x71, 0x08, 0x81, 0x76,
	0x60, 0x8f, 0x56, 0x7b, 0xa6, 0xed, 0x34, 0x99, 0x3f, 0x4d, 0xf7, 0x84, 0x0d, 0x17, 0x9e, 0x80,
	0x03, 0x9c, 0x39, 0xf2, 0x0c, 0xbc, 0x02, 0xaf, 0x45, 0x55, 0xf7, 0xd8, 0x4e, 0xa2, 0x1c, 0xb8,
	0xd5, 0xf7, 0xf5, 0x57, 0x3d, 0x5f, 0x77, 0x55, 0xd7, 0x00, 0x98, 0xaa, 0xdc, 0x9e, 0xd6, 0x4d,
	0x65, 0x2b, 0x3e, 0xa0, 0x38, 0xf9, 0x77, 0x00, 0xa3, 0x25, 0x06, 0x17, 0xeb, 0x9f, 0xf9, 0x0c,
	0x7a, 0x3a, 0x8b, 0x83, 0xe7, 0xc1, 0x49, 0x24, 0x30, 0xe2, 0x0c, 0xfa, 0x6d, 0xa3, 0xe3, 0x9e,
	0x23, 0x28, 0xe4, 0x1f, 0xc0, 0xd0, 0x54, 0x6d, 0x93, 0xaa, 0xb8, 0x8f, 0xe4, 0x64, 0x3e, 0x3d,
	0x75, 0x1b, 0x2e, 0x1d, 0x27, 0xba, 0x35, 0xfe, 0x05, 0x40, 0xa1, 0x32, 0x2d, 0x57, 0xf6, 0xb6,
	0x56, 0xf1, 0x00, 0x95, 0xb3, 0xf9, 0xbb, 0x3b, 0xa5, 0xfb, 0xd4, 0xe9, 0x39, 0xad, 0x5f, 0xe2,
	0xb2, 0x88, 0x8a, 0x5d, 0xc8, 0x9f, 0xc3, 0xa4, 0x2d, 0xe5, 0x8d, 0xd4, 0xb9, 0x5c, 0xe7, 0x2a,
	0x0e, 0x31, 0x71, 0x2c, 0xee, 0x52, 0xfc, 0x2d, 0x08, 0x6f, 0x64, 0x8e, 0x26, 0x87, 0x6e, 0xcd,
	0x03, 0xfe, 0x0c, 0x06, 0x56, 0x6e, 0x4d, 0x3c, 0x72, 0x9e, 0xc0, 0x7f, 0xe9, 0x12, 0x19, 0xe1,
	0x78, 0xfe, 0x35, 0x30, 0xd9, 0x66, 0xba, 0x5a, 0xe1, 0xc1, 0x6b, 0xd5, 0x58, 0xad, 0x4c, 0x3c,
	0x76, 0xda, 0xb7, 0xbd, 0x76, 0x41, 0xab, 0xdf, 0xef, 0x17, 0xc5, 0x53, 0x79, 0x9f, 0xe0, 0x1f,
	0x43, 0xb4, 0xd1, 0xb9, 0x5a, 0xe9, 0x72, 0x53, 0xc5, 0x91, 0x4b, 0x9d, 0xf9, 0xd4, 0x33, 0xa4,
	0xbf, 0x45, 0x56, 0x8c, 0x37, 0x5d, 0xc4, 0x4f, 0x20, 0xc2, 0xcf, 0xae, 0x8c, 0x95, 0xd6, 0xc4,
	0xe0, 0xc4, 0x93, 0xee, 0xf4, 0x44, 0x89, 0x31, 0xae, 0xba, 0x88, 0x73, 0x18, 0x64, 0xd2, 0xca,
	0x78, 0x82, 0xa2, 0xa9, 0x70, 0x71, 0xf2, 0x4f, 0x00, 0xd1, 0xfe, 0x76, 0xf8, 0x04, 0x46, 0x3f,
	0x96, 0xd7, 0x65, 0xf5, 0x4b, 0xc9, 0x8e, 0xf8, 0x08, 0xfa, 0x0b, 0xb3, 0x61, 0x01, 0x1f, 0xc3,
	0xe0, 0x2c, 0x97, 0x29, 0xeb, 0x11, 0x75, 0x5e, 0x7f, 0xce, 0xfa, 0x3e, 0x48, 0xd9, 0x80, 0xd6,
	0xce, 0x6b, 0xb5, 0x65, 0x21, 0xe5, 0x5e, 0x6c, 0xb7, 0x4e, 0x38, 0xe4, 0x53, 0x18, 0x23, 0x58,
	0xd6, 0x4a, 0xbd, 0x61, 0x23, 0xfe, 0x04, 0x22, 0x44, 0x3f, 0x55, 0xcd, 0x5a, 0x1b, 0x36, 0xa6,
	0x9c, 0x85, 0xde, 0x6c, 0x58, 0x44, 0xdb, 0xbc, 0x96, 0x37, 0x0c, 0x48, 0x71, 0xd9, 0xb4, 0xca,
	0xdd, 0x0c, 0x9b, 0x90, 0xe2, 0x65, 0x96, 0x49, 0x36, 0xed, 0x76, 0xbd, 0xa8, 0x5b, 0xc3, 0x9e,
	0x70, 0x80, 0xe1, 0xd2, 0x36, 0x4a, 0x16, 0x2c, 0x4d, 0xbe, 0xc1, 0xd8, 0xd7, 0xbf, 0xeb, 0x9b,
	0xe0, 0xd0, 0x37, 0x78, 0x50, 0xd7, 0x0b, 0xd4, 0x4a, 0xa1, 0x70, 0x31, 0x71, 0xa5, 0x2c, 0x7c,
	0x27, 0x45, 0xc2, 0xc5, 0xc9, 0x6f, 0xf0, 0xf4, 0x41, 0x2d, 0x78, 0x0c, 0xa3, 0xb5, 0xb6, 0x8d,
	0xb4, 0xca, 0x6d, 0x18, 0x8a, 0x1d, 0xc4, 0xb2, 0x83, 0x91, 0x45, 0x9d, 0x2b, 0xb7, 0xe8, 0xb7,
	0xbe, 0xc3, 0xf0, 0x63, 0x18, 0xa7, 0x57, 0xb2, 0x2c, 0x55, 0x6e, 0xdc, 0x47, 0x42, 0xb1, 0xc7,
	0xfc, 0x1d, 0x18, 0xe6, 0xaa, 0xdc, 0xda, 0x2b, 0xd7, 0x9e, 0xa1, 0xe8, 0x50, 0xf2, 0x67, 0x00,
	0xe3, 0x5d, 0x49, 0xf7, 0x0e, 0x83, 0x83, 0x43, 0xe2, 0x6a, 0x89, 0x69, 0xfe, 0x51, 0xb8, 0x98,
	0x36, 0x33, 0xed, 0x66, 0xa3, 0xdf, 0x74, 0x67, 0xe9, 0x10, 0x19, 0xa0, 0xa6, 0x30, 0xfa, 0x57,
	0xff, 0x0a, 0xfa, 0x62, 0x8f, 0xa9, 0x93, 0x0b, 0xab, 0x0b, 0xdf, 0xe5, 0x91, 0xf0, 0x80, 0xd8,
	0xd4, 0xb1, 0x43, 0xcf, 0x3a, 0x80, 0xb7, 0x32, 0x79, 0xd5, 0x54, 0x15, 0x3e, 0x14, 0x7a, 0x3e,
	0x8f, 0xda, 0xc2, 0x44, 0xab, 0x6d, 0xae, 0x3a, 0x5f, 0x1e, 0x90, 0x31, 0x89, 0xb7, 0x68, 0xec,
	0xce, 0x98, 0x47, 0xa4, 0x96, 0xf9, 0xba, 0x2d, 0x9c, 0x2b, 0x54, 0x3b, 0x40, 0x37, 0x4d, 0x5d,
	0xaa, 0x33, 0x83, 0xa6, 0xfa, 0xc8, 0xef, 0x60, 0xf2, 0x47, 0x00, 0xa1, 0xef, 0xd8, 0xf7, 0x21,
	0xaa, 0x73, 0x79, 0x9b, 0x56, 0x6d, 0x69, 0xbb, 0x7a, 0x1c, 0x08, 0xaa, 0x48, 0x2e, 0x8d, 0x25,
	0x42, 0x65, 0x9d, 0x95, 0x3b, 0x0c, 0x65, 0x9b, 0x6b, 0x5d, 0xfb, 0x6c, 0x5f, 0x92, 0x03, 0x41,
	0x6e, 0xb1, 0x6e, 0xba, 0xdc, 0xee, 0x6a, 0xe2, 0x11, 0xb9, 0x35, 0x69, 0xd5, 0xf8, 0xab, 0x0a,
	0x85, 0x07, 0xc9, 0xdf, 0x01, 0x0c, 0xe8, 0x8d, 0x1f, 0x8e, 0x1e, 0x3c, 0x7e, 0xf4, 0xde, 0xe3,
	0x47, 0xef, 0x3f, 0x38, 0x7a, 0x5a, 0x15, 0x85, 0x42, 0x5b, 0xfe, 0x4a, 0x76, 0x90, 0xf4, 0x5b,
	0x55, 0x36, 0xfb, 0x3a, 0x39, 0x40, 0x25, 0xb8, 0x55, 0xb2, 0x71, 0x65, 0xc2, 0x7e, 0xa6, 0xd8,
	0xf9, 0x68, 0x64, 0x7a, 0xed, 0xc6, 0x10, 0xda, 0x74, 0x20, 0xf9, 0xab, 0x07, 0x11, 0x55, 0xed,
	0x87, 0x56, 0x35, 0xb7, 0xae, 0x53, 0x50, 0x9b, 0x5e, 0x75, 0x66, 0x3b, 0xc4, 0xbf, 0x82, 0x89,
	0x8f, 0x56, 0xfb, 0x67, 0x32, 0x9b, 0x1f, 0x1f, 0x46, 0xa6, 0xcb, 0x3e, 0x5d, 0x3a, 0x89, 0x9b,
	0x9a, 0x60, 0xf6, 0x31, 0x7f, 0x01, 0xd3, 0x2e, 0x79, 0xa3, 0x55, 0x9e, 0xb9, 0x93, 0xcd, 0xe6,
	0xef, 0x3d, 0x9e, 0x4d, 0xcd, 0x9d, 0x89, 0xee, 0x6b, 0x67, 0xa4, 0x4f, 0x5e, 0x00, 0x1c, 0x76,
	0xe6, 0x11, 0x84, 0x42, 0x6d, 0x71, 0x4a, 0x1c, 0xd1, 0xcc, 0x78, 0xad, 0xf3, 0x2c, 0x95, 0x4d,
	0x86, 0x43, 0x07, 0xd1, 0xcb, 0xaa, 0xb4, 0x52, 0x97, 0x06, 0x07, 0x0f, 0xca, 0xce, 0xa5, 0x4d,
	0xaf, 0x58, 0x3f, 0xf9, 0x12, 0x26, 0x77, 0xf6, 0xa6, 0x51, 0xf1, 0x1d, 0xb6, 0x24, 0xe6, 0xa3,
	0xe6, 0x92, 0x8a, 0x81, 0x72, 0x9c, 0x1a, 0x0b, 0xdb, 0x60, 0x01, 0x0c, 0xce, 0x2a, 0xe4, 0x17,
	0x74, 0xed, 0x6c, 0x90, 0x3c, 0x83, 0x29, 0x39, 0x14, 0xca, 0xd4, 0x55, 0x89, 0xa3, 0xe3, 0xc1,
	0x2f, 0x68, 0xfe, 0x7b, 0xe0, 0x7f, 0x4f, 0xcb, 0xe6, 0x86, 0x7f, 0x84, 0xe3, 0x2f, 0xcb, 0xf8,
	0x93, 0x7b, 0x7f, 0x92, 0x63, 0x7e, 0x80, 0xbb, 0x5d, 0x92, 0x23, 0xfe, 0x09, 0x4e, 0xb4, 0x2c,
	0x33, 0xff, 0x4b, 0x7c, 0x12, 0x7c, 0x1a, 0xf0, 0x0f, 0xa1, 0xff, 0x4a, 0xd9, 0x87, 0xfa, 0xfb,
	0x30, 0x39, 0x5a, 0x0f, 0xdd, 0xbf, 0xf3, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x35,
	0xb7, 0xf4, 0x49, 0x07, 0x00, 0x00,
}
