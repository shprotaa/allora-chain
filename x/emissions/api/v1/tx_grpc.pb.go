// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: emissions/v1/tx.proto

package emissionsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateParams_FullMethodName                     = "/emissions.v1.Msg/UpdateParams"
	Msg_InsertInferences_FullMethodName                 = "/emissions.v1.Msg/InsertInferences"
	Msg_InsertForecasts_FullMethodName                  = "/emissions.v1.Msg/InsertForecasts"
	Msg_CreateNewTopic_FullMethodName                   = "/emissions.v1.Msg/CreateNewTopic"
	Msg_ReactivateTopic_FullMethodName                  = "/emissions.v1.Msg/ReactivateTopic"
	Msg_Register_FullMethodName                         = "/emissions.v1.Msg/Register"
	Msg_AddNewRegistration_FullMethodName               = "/emissions.v1.Msg/AddNewRegistration"
	Msg_RemoveRegistration_FullMethodName               = "/emissions.v1.Msg/RemoveRegistration"
	Msg_InsertLosses_FullMethodName                     = "/emissions.v1.Msg/InsertLosses"
	Msg_AddStake_FullMethodName                         = "/emissions.v1.Msg/AddStake"
	Msg_StartRemoveStake_FullMethodName                 = "/emissions.v1.Msg/StartRemoveStake"
	Msg_ConfirmRemoveStake_FullMethodName               = "/emissions.v1.Msg/ConfirmRemoveStake"
	Msg_StartRemoveAllStake_FullMethodName              = "/emissions.v1.Msg/StartRemoveAllStake"
	Msg_DelegateStake_FullMethodName                    = "/emissions.v1.Msg/DelegateStake"
	Msg_StartRemoveDelegatedStake_FullMethodName        = "/emissions.v1.Msg/StartRemoveDelegatedStake"
	Msg_RequestInference_FullMethodName                 = "/emissions.v1.Msg/RequestInference"
	Msg_AddToWhitelistAdmin_FullMethodName              = "/emissions.v1.Msg/AddToWhitelistAdmin"
	Msg_RemoveFromWhitelistAdmin_FullMethodName         = "/emissions.v1.Msg/RemoveFromWhitelistAdmin"
	Msg_AddToTopicCreationWhitelist_FullMethodName      = "/emissions.v1.Msg/AddToTopicCreationWhitelist"
	Msg_RemoveFromTopicCreationWhitelist_FullMethodName = "/emissions.v1.Msg/RemoveFromTopicCreationWhitelist"
	Msg_AddToReputerWhitelist_FullMethodName            = "/emissions.v1.Msg/AddToReputerWhitelist"
	Msg_RemoveFromReputerWhitelist_FullMethodName       = "/emissions.v1.Msg/RemoveFromReputerWhitelist"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	InsertInferences(ctx context.Context, in *MsgInsertInferences, opts ...grpc.CallOption) (*MsgInsertInferencesResponse, error)
	InsertForecasts(ctx context.Context, in *MsgInsertForecasts, opts ...grpc.CallOption) (*MsgInsertForecastsResponse, error)
	CreateNewTopic(ctx context.Context, in *MsgCreateNewTopic, opts ...grpc.CallOption) (*MsgCreateNewTopicResponse, error)
	ReactivateTopic(ctx context.Context, in *MsgReactivateTopic, opts ...grpc.CallOption) (*MsgReactivateTopicResponse, error)
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error)
	AddNewRegistration(ctx context.Context, in *MsgAddNewRegistration, opts ...grpc.CallOption) (*MsgAddNewRegistrationResponse, error)
	RemoveRegistration(ctx context.Context, in *MsgRemoveRegistration, opts ...grpc.CallOption) (*MsgRemoveRegistrationResponse, error)
	InsertLosses(ctx context.Context, in *MsgInsertLosses, opts ...grpc.CallOption) (*MsgInsertLossesResponse, error)
	AddStake(ctx context.Context, in *MsgAddStake, opts ...grpc.CallOption) (*MsgAddStakeResponse, error)
	StartRemoveStake(ctx context.Context, in *MsgStartRemoveStake, opts ...grpc.CallOption) (*MsgStartRemoveStakeResponse, error)
	ConfirmRemoveStake(ctx context.Context, in *MsgConfirmRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveStakeResponse, error)
	StartRemoveAllStake(ctx context.Context, in *MsgStartRemoveAllStake, opts ...grpc.CallOption) (*MsgStartRemoveAllStakeResponse, error)
	DelegateStake(ctx context.Context, in *MsgDelegateStake, opts ...grpc.CallOption) (*MsgDelegateStakeResponse, error)
	StartRemoveDelegatedStake(ctx context.Context, in *MsgStartRemoveDelegatedStake, opts ...grpc.CallOption) (*MsgStartRemoveDelegatedStakeResponse, error)
	RequestInference(ctx context.Context, in *MsgRequestInference, opts ...grpc.CallOption) (*MsgRequestInferenceResponse, error)
	AddToWhitelistAdmin(ctx context.Context, in *MsgAddToWhitelistAdmin, opts ...grpc.CallOption) (*MsgAddToWhitelistAdminResponse, error)
	RemoveFromWhitelistAdmin(ctx context.Context, in *MsgRemoveFromWhitelistAdmin, opts ...grpc.CallOption) (*MsgRemoveFromWhitelistAdminResponse, error)
	AddToTopicCreationWhitelist(ctx context.Context, in *MsgAddToTopicCreationWhitelist, opts ...grpc.CallOption) (*MsgAddToTopicCreationWhitelistResponse, error)
	RemoveFromTopicCreationWhitelist(ctx context.Context, in *MsgRemoveFromTopicCreationWhitelist, opts ...grpc.CallOption) (*MsgRemoveFromTopicCreationWhitelistResponse, error)
	AddToReputerWhitelist(ctx context.Context, in *MsgAddToReputerWhitelist, opts ...grpc.CallOption) (*MsgAddToReputerWhitelistResponse, error)
	RemoveFromReputerWhitelist(ctx context.Context, in *MsgRemoveFromReputerWhitelist, opts ...grpc.CallOption) (*MsgRemoveFromReputerWhitelistResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InsertInferences(ctx context.Context, in *MsgInsertInferences, opts ...grpc.CallOption) (*MsgInsertInferencesResponse, error) {
	out := new(MsgInsertInferencesResponse)
	err := c.cc.Invoke(ctx, Msg_InsertInferences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InsertForecasts(ctx context.Context, in *MsgInsertForecasts, opts ...grpc.CallOption) (*MsgInsertForecastsResponse, error) {
	out := new(MsgInsertForecastsResponse)
	err := c.cc.Invoke(ctx, Msg_InsertForecasts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateNewTopic(ctx context.Context, in *MsgCreateNewTopic, opts ...grpc.CallOption) (*MsgCreateNewTopicResponse, error) {
	out := new(MsgCreateNewTopicResponse)
	err := c.cc.Invoke(ctx, Msg_CreateNewTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ReactivateTopic(ctx context.Context, in *MsgReactivateTopic, opts ...grpc.CallOption) (*MsgReactivateTopicResponse, error) {
	out := new(MsgReactivateTopicResponse)
	err := c.cc.Invoke(ctx, Msg_ReactivateTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgRegisterResponse, error) {
	out := new(MsgRegisterResponse)
	err := c.cc.Invoke(ctx, Msg_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddNewRegistration(ctx context.Context, in *MsgAddNewRegistration, opts ...grpc.CallOption) (*MsgAddNewRegistrationResponse, error) {
	out := new(MsgAddNewRegistrationResponse)
	err := c.cc.Invoke(ctx, Msg_AddNewRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveRegistration(ctx context.Context, in *MsgRemoveRegistration, opts ...grpc.CallOption) (*MsgRemoveRegistrationResponse, error) {
	out := new(MsgRemoveRegistrationResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InsertLosses(ctx context.Context, in *MsgInsertLosses, opts ...grpc.CallOption) (*MsgInsertLossesResponse, error) {
	out := new(MsgInsertLossesResponse)
	err := c.cc.Invoke(ctx, Msg_InsertLosses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddStake(ctx context.Context, in *MsgAddStake, opts ...grpc.CallOption) (*MsgAddStakeResponse, error) {
	out := new(MsgAddStakeResponse)
	err := c.cc.Invoke(ctx, Msg_AddStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartRemoveStake(ctx context.Context, in *MsgStartRemoveStake, opts ...grpc.CallOption) (*MsgStartRemoveStakeResponse, error) {
	out := new(MsgStartRemoveStakeResponse)
	err := c.cc.Invoke(ctx, Msg_StartRemoveStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConfirmRemoveStake(ctx context.Context, in *MsgConfirmRemoveStake, opts ...grpc.CallOption) (*MsgConfirmRemoveStakeResponse, error) {
	out := new(MsgConfirmRemoveStakeResponse)
	err := c.cc.Invoke(ctx, Msg_ConfirmRemoveStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartRemoveAllStake(ctx context.Context, in *MsgStartRemoveAllStake, opts ...grpc.CallOption) (*MsgStartRemoveAllStakeResponse, error) {
	out := new(MsgStartRemoveAllStakeResponse)
	err := c.cc.Invoke(ctx, Msg_StartRemoveAllStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelegateStake(ctx context.Context, in *MsgDelegateStake, opts ...grpc.CallOption) (*MsgDelegateStakeResponse, error) {
	out := new(MsgDelegateStakeResponse)
	err := c.cc.Invoke(ctx, Msg_DelegateStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartRemoveDelegatedStake(ctx context.Context, in *MsgStartRemoveDelegatedStake, opts ...grpc.CallOption) (*MsgStartRemoveDelegatedStakeResponse, error) {
	out := new(MsgStartRemoveDelegatedStakeResponse)
	err := c.cc.Invoke(ctx, Msg_StartRemoveDelegatedStake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestInference(ctx context.Context, in *MsgRequestInference, opts ...grpc.CallOption) (*MsgRequestInferenceResponse, error) {
	out := new(MsgRequestInferenceResponse)
	err := c.cc.Invoke(ctx, Msg_RequestInference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToWhitelistAdmin(ctx context.Context, in *MsgAddToWhitelistAdmin, opts ...grpc.CallOption) (*MsgAddToWhitelistAdminResponse, error) {
	out := new(MsgAddToWhitelistAdminResponse)
	err := c.cc.Invoke(ctx, Msg_AddToWhitelistAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveFromWhitelistAdmin(ctx context.Context, in *MsgRemoveFromWhitelistAdmin, opts ...grpc.CallOption) (*MsgRemoveFromWhitelistAdminResponse, error) {
	out := new(MsgRemoveFromWhitelistAdminResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveFromWhitelistAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToTopicCreationWhitelist(ctx context.Context, in *MsgAddToTopicCreationWhitelist, opts ...grpc.CallOption) (*MsgAddToTopicCreationWhitelistResponse, error) {
	out := new(MsgAddToTopicCreationWhitelistResponse)
	err := c.cc.Invoke(ctx, Msg_AddToTopicCreationWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveFromTopicCreationWhitelist(ctx context.Context, in *MsgRemoveFromTopicCreationWhitelist, opts ...grpc.CallOption) (*MsgRemoveFromTopicCreationWhitelistResponse, error) {
	out := new(MsgRemoveFromTopicCreationWhitelistResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveFromTopicCreationWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToReputerWhitelist(ctx context.Context, in *MsgAddToReputerWhitelist, opts ...grpc.CallOption) (*MsgAddToReputerWhitelistResponse, error) {
	out := new(MsgAddToReputerWhitelistResponse)
	err := c.cc.Invoke(ctx, Msg_AddToReputerWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveFromReputerWhitelist(ctx context.Context, in *MsgRemoveFromReputerWhitelist, opts ...grpc.CallOption) (*MsgRemoveFromReputerWhitelistResponse, error) {
	out := new(MsgRemoveFromReputerWhitelistResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveFromReputerWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	InsertInferences(context.Context, *MsgInsertInferences) (*MsgInsertInferencesResponse, error)
	InsertForecasts(context.Context, *MsgInsertForecasts) (*MsgInsertForecastsResponse, error)
	CreateNewTopic(context.Context, *MsgCreateNewTopic) (*MsgCreateNewTopicResponse, error)
	ReactivateTopic(context.Context, *MsgReactivateTopic) (*MsgReactivateTopicResponse, error)
	Register(context.Context, *MsgRegister) (*MsgRegisterResponse, error)
	AddNewRegistration(context.Context, *MsgAddNewRegistration) (*MsgAddNewRegistrationResponse, error)
	RemoveRegistration(context.Context, *MsgRemoveRegistration) (*MsgRemoveRegistrationResponse, error)
	InsertLosses(context.Context, *MsgInsertLosses) (*MsgInsertLossesResponse, error)
	AddStake(context.Context, *MsgAddStake) (*MsgAddStakeResponse, error)
	StartRemoveStake(context.Context, *MsgStartRemoveStake) (*MsgStartRemoveStakeResponse, error)
	ConfirmRemoveStake(context.Context, *MsgConfirmRemoveStake) (*MsgConfirmRemoveStakeResponse, error)
	StartRemoveAllStake(context.Context, *MsgStartRemoveAllStake) (*MsgStartRemoveAllStakeResponse, error)
	DelegateStake(context.Context, *MsgDelegateStake) (*MsgDelegateStakeResponse, error)
	StartRemoveDelegatedStake(context.Context, *MsgStartRemoveDelegatedStake) (*MsgStartRemoveDelegatedStakeResponse, error)
	RequestInference(context.Context, *MsgRequestInference) (*MsgRequestInferenceResponse, error)
	AddToWhitelistAdmin(context.Context, *MsgAddToWhitelistAdmin) (*MsgAddToWhitelistAdminResponse, error)
	RemoveFromWhitelistAdmin(context.Context, *MsgRemoveFromWhitelistAdmin) (*MsgRemoveFromWhitelistAdminResponse, error)
	AddToTopicCreationWhitelist(context.Context, *MsgAddToTopicCreationWhitelist) (*MsgAddToTopicCreationWhitelistResponse, error)
	RemoveFromTopicCreationWhitelist(context.Context, *MsgRemoveFromTopicCreationWhitelist) (*MsgRemoveFromTopicCreationWhitelistResponse, error)
	AddToReputerWhitelist(context.Context, *MsgAddToReputerWhitelist) (*MsgAddToReputerWhitelistResponse, error)
	RemoveFromReputerWhitelist(context.Context, *MsgRemoveFromReputerWhitelist) (*MsgRemoveFromReputerWhitelistResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) InsertInferences(context.Context, *MsgInsertInferences) (*MsgInsertInferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertInferences not implemented")
}
func (UnimplementedMsgServer) InsertForecasts(context.Context, *MsgInsertForecasts) (*MsgInsertForecastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertForecasts not implemented")
}
func (UnimplementedMsgServer) CreateNewTopic(context.Context, *MsgCreateNewTopic) (*MsgCreateNewTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTopic not implemented")
}
func (UnimplementedMsgServer) ReactivateTopic(context.Context, *MsgReactivateTopic) (*MsgReactivateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReactivateTopic not implemented")
}
func (UnimplementedMsgServer) Register(context.Context, *MsgRegister) (*MsgRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMsgServer) AddNewRegistration(context.Context, *MsgAddNewRegistration) (*MsgAddNewRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewRegistration not implemented")
}
func (UnimplementedMsgServer) RemoveRegistration(context.Context, *MsgRemoveRegistration) (*MsgRemoveRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRegistration not implemented")
}
func (UnimplementedMsgServer) InsertLosses(context.Context, *MsgInsertLosses) (*MsgInsertLossesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertLosses not implemented")
}
func (UnimplementedMsgServer) AddStake(context.Context, *MsgAddStake) (*MsgAddStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStake not implemented")
}
func (UnimplementedMsgServer) StartRemoveStake(context.Context, *MsgStartRemoveStake) (*MsgStartRemoveStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRemoveStake not implemented")
}
func (UnimplementedMsgServer) ConfirmRemoveStake(context.Context, *MsgConfirmRemoveStake) (*MsgConfirmRemoveStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRemoveStake not implemented")
}
func (UnimplementedMsgServer) StartRemoveAllStake(context.Context, *MsgStartRemoveAllStake) (*MsgStartRemoveAllStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRemoveAllStake not implemented")
}
func (UnimplementedMsgServer) DelegateStake(context.Context, *MsgDelegateStake) (*MsgDelegateStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateStake not implemented")
}
func (UnimplementedMsgServer) StartRemoveDelegatedStake(context.Context, *MsgStartRemoveDelegatedStake) (*MsgStartRemoveDelegatedStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRemoveDelegatedStake not implemented")
}
func (UnimplementedMsgServer) RequestInference(context.Context, *MsgRequestInference) (*MsgRequestInferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestInference not implemented")
}
func (UnimplementedMsgServer) AddToWhitelistAdmin(context.Context, *MsgAddToWhitelistAdmin) (*MsgAddToWhitelistAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhitelistAdmin not implemented")
}
func (UnimplementedMsgServer) RemoveFromWhitelistAdmin(context.Context, *MsgRemoveFromWhitelistAdmin) (*MsgRemoveFromWhitelistAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWhitelistAdmin not implemented")
}
func (UnimplementedMsgServer) AddToTopicCreationWhitelist(context.Context, *MsgAddToTopicCreationWhitelist) (*MsgAddToTopicCreationWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToTopicCreationWhitelist not implemented")
}
func (UnimplementedMsgServer) RemoveFromTopicCreationWhitelist(context.Context, *MsgRemoveFromTopicCreationWhitelist) (*MsgRemoveFromTopicCreationWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromTopicCreationWhitelist not implemented")
}
func (UnimplementedMsgServer) AddToReputerWhitelist(context.Context, *MsgAddToReputerWhitelist) (*MsgAddToReputerWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToReputerWhitelist not implemented")
}
func (UnimplementedMsgServer) RemoveFromReputerWhitelist(context.Context, *MsgRemoveFromReputerWhitelist) (*MsgRemoveFromReputerWhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromReputerWhitelist not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InsertInferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertInferences)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertInferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InsertInferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertInferences(ctx, req.(*MsgInsertInferences))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InsertForecasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertForecasts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertForecasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InsertForecasts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertForecasts(ctx, req.(*MsgInsertForecasts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateNewTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateNewTopic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateNewTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateNewTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateNewTopic(ctx, req.(*MsgCreateNewTopic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ReactivateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReactivateTopic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ReactivateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ReactivateTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ReactivateTopic(ctx, req.(*MsgReactivateTopic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddNewRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddNewRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddNewRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddNewRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddNewRegistration(ctx, req.(*MsgAddNewRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveRegistration(ctx, req.(*MsgRemoveRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InsertLosses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertLosses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertLosses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InsertLosses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertLosses(ctx, req.(*MsgInsertLosses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddStake(ctx, req.(*MsgAddStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartRemoveStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRemoveStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartRemoveStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartRemoveStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartRemoveStake(ctx, req.(*MsgStartRemoveStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConfirmRemoveStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConfirmRemoveStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConfirmRemoveStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConfirmRemoveStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConfirmRemoveStake(ctx, req.(*MsgConfirmRemoveStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartRemoveAllStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRemoveAllStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartRemoveAllStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartRemoveAllStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartRemoveAllStake(ctx, req.(*MsgStartRemoveAllStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelegateStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegateStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelegateStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DelegateStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelegateStake(ctx, req.(*MsgDelegateStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartRemoveDelegatedStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRemoveDelegatedStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartRemoveDelegatedStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartRemoveDelegatedStake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartRemoveDelegatedStake(ctx, req.(*MsgStartRemoveDelegatedStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestInference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestInference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestInference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestInference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestInference(ctx, req.(*MsgRequestInference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToWhitelistAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToWhitelistAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToWhitelistAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddToWhitelistAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToWhitelistAdmin(ctx, req.(*MsgAddToWhitelistAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveFromWhitelistAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveFromWhitelistAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveFromWhitelistAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveFromWhitelistAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveFromWhitelistAdmin(ctx, req.(*MsgRemoveFromWhitelistAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToTopicCreationWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToTopicCreationWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToTopicCreationWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddToTopicCreationWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToTopicCreationWhitelist(ctx, req.(*MsgAddToTopicCreationWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveFromTopicCreationWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveFromTopicCreationWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveFromTopicCreationWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveFromTopicCreationWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveFromTopicCreationWhitelist(ctx, req.(*MsgRemoveFromTopicCreationWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToReputerWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToReputerWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToReputerWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddToReputerWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToReputerWhitelist(ctx, req.(*MsgAddToReputerWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveFromReputerWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveFromReputerWhitelist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveFromReputerWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveFromReputerWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveFromReputerWhitelist(ctx, req.(*MsgRemoveFromReputerWhitelist))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emissions.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "InsertInferences",
			Handler:    _Msg_InsertInferences_Handler,
		},
		{
			MethodName: "InsertForecasts",
			Handler:    _Msg_InsertForecasts_Handler,
		},
		{
			MethodName: "CreateNewTopic",
			Handler:    _Msg_CreateNewTopic_Handler,
		},
		{
			MethodName: "ReactivateTopic",
			Handler:    _Msg_ReactivateTopic_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
		{
			MethodName: "AddNewRegistration",
			Handler:    _Msg_AddNewRegistration_Handler,
		},
		{
			MethodName: "RemoveRegistration",
			Handler:    _Msg_RemoveRegistration_Handler,
		},
		{
			MethodName: "InsertLosses",
			Handler:    _Msg_InsertLosses_Handler,
		},
		{
			MethodName: "AddStake",
			Handler:    _Msg_AddStake_Handler,
		},
		{
			MethodName: "StartRemoveStake",
			Handler:    _Msg_StartRemoveStake_Handler,
		},
		{
			MethodName: "ConfirmRemoveStake",
			Handler:    _Msg_ConfirmRemoveStake_Handler,
		},
		{
			MethodName: "StartRemoveAllStake",
			Handler:    _Msg_StartRemoveAllStake_Handler,
		},
		{
			MethodName: "DelegateStake",
			Handler:    _Msg_DelegateStake_Handler,
		},
		{
			MethodName: "StartRemoveDelegatedStake",
			Handler:    _Msg_StartRemoveDelegatedStake_Handler,
		},
		{
			MethodName: "RequestInference",
			Handler:    _Msg_RequestInference_Handler,
		},
		{
			MethodName: "AddToWhitelistAdmin",
			Handler:    _Msg_AddToWhitelistAdmin_Handler,
		},
		{
			MethodName: "RemoveFromWhitelistAdmin",
			Handler:    _Msg_RemoveFromWhitelistAdmin_Handler,
		},
		{
			MethodName: "AddToTopicCreationWhitelist",
			Handler:    _Msg_AddToTopicCreationWhitelist_Handler,
		},
		{
			MethodName: "RemoveFromTopicCreationWhitelist",
			Handler:    _Msg_RemoveFromTopicCreationWhitelist_Handler,
		},
		{
			MethodName: "AddToReputerWhitelist",
			Handler:    _Msg_AddToReputerWhitelist_Handler,
		},
		{
			MethodName: "RemoveFromReputerWhitelist",
			Handler:    _Msg_RemoveFromReputerWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emissions/v1/tx.proto",
}
