// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: game.proto

package coms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_PASS    Action = 0
	Action_MOVE    Action = 1
	Action_ATTACK  Action = 2
	Action_CONNECT Action = 3
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "PASS",
		1: "MOVE",
		2: "ATTACK",
		3: "CONNECT",
	}
	Action_value = map[string]int32{
		"PASS":    0,
		"MOVE":    1,
		"ATTACK":  2,
		"CONNECT": 3,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

type NewPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ServerAddress string `protobuf:"bytes,2,opt,name=serverAddress,proto3" json:"serverAddress,omitempty"`
}

func (x *NewPlayer) Reset() {
	*x = NewPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayer) ProtoMessage() {}

func (x *NewPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayer.ProtoReflect.Descriptor instead.
func (*NewPlayer) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *NewPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewPlayer) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

type MapRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row []int32 `protobuf:"varint,1,rep,packed,name=Row,proto3" json:"Row,omitempty"`
}

func (x *MapRow) Reset() {
	*x = MapRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRow) ProtoMessage() {}

func (x *MapRow) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRow.ProtoReflect.Descriptor instead.
func (*MapRow) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *MapRow) GetRow() []int32 {
	if x != nil {
		return x.Row
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destination.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *Position) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Lighthouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position    *Position   `protobuf:"bytes,1,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Owner       int32       `protobuf:"varint,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Energy      int32       `protobuf:"varint,3,opt,name=Energy,proto3" json:"Energy,omitempty"`
	Connections []*Position `protobuf:"bytes,4,rep,name=Connections,proto3" json:"Connections,omitempty"`
	HaveKey     bool        `protobuf:"varint,5,opt,name=HaveKey,proto3" json:"HaveKey,omitempty"`
}

func (x *Lighthouse) Reset() {
	*x = Lighthouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lighthouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lighthouse) ProtoMessage() {}

func (x *Lighthouse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lighthouse.ProtoReflect.Descriptor instead.
func (*Lighthouse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *Lighthouse) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Lighthouse) GetOwner() int32 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *Lighthouse) GetEnergy() int32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *Lighthouse) GetConnections() []*Position {
	if x != nil {
		return x.Connections
	}
	return nil
}

func (x *Lighthouse) GetHaveKey() bool {
	if x != nil {
		return x.HaveKey
	}
	return false
}

type NewPlayerInitialState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerNum   int32         `protobuf:"varint,1,opt,name=PlayerNum,proto3" json:"PlayerNum,omitempty"`
	PlayerCount int32         `protobuf:"varint,2,opt,name=PlayerCount,proto3" json:"PlayerCount,omitempty"`
	Position    *Position     `protobuf:"bytes,3,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Map         []*MapRow     `protobuf:"bytes,4,rep,name=Map,proto3" json:"Map,omitempty"`
	Lighthouses []*Lighthouse `protobuf:"bytes,5,rep,name=Lighthouses,proto3" json:"Lighthouses,omitempty"`
}

func (x *NewPlayerInitialState) Reset() {
	*x = NewPlayerInitialState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlayerInitialState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayerInitialState) ProtoMessage() {}

func (x *NewPlayerInitialState) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayerInitialState.ProtoReflect.Descriptor instead.
func (*NewPlayerInitialState) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *NewPlayerInitialState) GetPlayerNum() int32 {
	if x != nil {
		return x.PlayerNum
	}
	return 0
}

func (x *NewPlayerInitialState) GetPlayerCount() int32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

func (x *NewPlayerInitialState) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *NewPlayerInitialState) GetMap() []*MapRow {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *NewPlayerInitialState) GetLighthouses() []*Lighthouse {
	if x != nil {
		return x.Lighthouses
	}
	return nil
}

type NewTurn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position    *Position     `protobuf:"bytes,1,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Score       int32         `protobuf:"varint,2,opt,name=Score,proto3" json:"Score,omitempty"`
	Energy      int32         `protobuf:"varint,3,opt,name=Energy,proto3" json:"Energy,omitempty"`
	View        []*MapRow     `protobuf:"bytes,4,rep,name=View,proto3" json:"View,omitempty"`
	Lighthouses []*Lighthouse `protobuf:"bytes,5,rep,name=Lighthouses,proto3" json:"Lighthouses,omitempty"`
}

func (x *NewTurn) Reset() {
	*x = NewTurn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTurn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTurn) ProtoMessage() {}

func (x *NewTurn) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTurn.ProtoReflect.Descriptor instead.
func (*NewTurn) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *NewTurn) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *NewTurn) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *NewTurn) GetEnergy() int32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *NewTurn) GetView() []*MapRow {
	if x != nil {
		return x.View
	}
	return nil
}

func (x *NewTurn) GetLighthouses() []*Lighthouse {
	if x != nil {
		return x.Lighthouses
	}
	return nil
}

type NewAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action      Action    `protobuf:"varint,1,opt,name=Action,proto3,enum=Action" json:"Action,omitempty"`
	Destination *Position `protobuf:"bytes,2,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Energy      int32     `protobuf:"varint,3,opt,name=Energy,proto3" json:"Energy,omitempty"`
}

func (x *NewAction) Reset() {
	*x = NewAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAction) ProtoMessage() {}

func (x *NewAction) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAction.ProtoReflect.Descriptor instead.
func (*NewAction) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *NewAction) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_PASS
}

func (x *NewAction) GetDestination() *Position {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *NewAction) GetEnergy() int32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x09,
	0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x77, 0x12, 0x10, 0x0a,
	0x03, 0x52, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x52, 0x6f, 0x77, 0x22,
	0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x59, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x2b, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x76, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x76, 0x65, 0x4b,
	0x65, 0x79, 0x22, 0xc8, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x77, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x2d,
	0x0a, 0x0b, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x0b, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22, 0xaa, 0x01,
	0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x1b,
	0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d,
	0x61, 0x70, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x2d, 0x0a, 0x0b, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x0b, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x09, 0x4e, 0x65,
	0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2a, 0x35, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x10, 0x03, 0x32, 0x5b, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0a, 0x2e, 0x4e, 0x65,
	0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x00, 0x12, 0x1e, 0x0a, 0x04, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x08, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x75, 0x72, 0x6e, 0x1a, 0x0a, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x6e, 0x61, 0x73, 0x64, 0x61, 0x63, 0x72, 0x75, 0x7a, 0x2f, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x69, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_game_proto_goTypes = []interface{}{
	(Action)(0),                   // 0: Action
	(*NewPlayer)(nil),             // 1: NewPlayer
	(*MapRow)(nil),                // 2: MapRow
	(*Position)(nil),              // 3: Destination
	(*Lighthouse)(nil),            // 4: Lighthouse
	(*NewPlayerInitialState)(nil), // 5: NewPlayerInitialState
	(*NewTurn)(nil),               // 6: NewTurn
	(*NewAction)(nil),             // 7: NewAction
}
var file_game_proto_depIdxs = []int32{
	3,  // 0: Lighthouse.Destination:type_name -> Destination
	3,  // 1: Lighthouse.Connections:type_name -> Destination
	3,  // 2: NewPlayerInitialState.Destination:type_name -> Destination
	2,  // 3: NewPlayerInitialState.Map:type_name -> MapRow
	4,  // 4: NewPlayerInitialState.Lighthouses:type_name -> Lighthouse
	3,  // 5: NewTurn.Destination:type_name -> Destination
	2,  // 6: NewTurn.View:type_name -> MapRow
	4,  // 7: NewTurn.Lighthouses:type_name -> Lighthouse
	0,  // 8: NewAction.Action:type_name -> Action
	3,  // 9: NewAction.Destination:type_name -> Destination
	1,  // 10: GameService.Join:input_type -> NewPlayer
	6,  // 11: GameService.Turn:input_type -> NewTurn
	5,  // 12: GameService.Join:output_type -> NewPlayerInitialState
	7,  // 13: GameService.Turn:output_type -> NewAction
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRow); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lighthouse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlayerInitialState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTurn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
