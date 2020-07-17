// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/user.proto

package mu_micro_book_user_srv_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	// user
	GetFromUid(ctx context.Context, in *GetFromUidReq, opts ...client.CallOption) (*GetFromUidRsp, error)
	GetFromPhone(ctx context.Context, in *GetFromPhoneReq, opts ...client.CallOption) (*GetFromPhoneRsp, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRsp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...client.CallOption) (*UpdateUserRsp, error)
	UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...client.CallOption) (*UpdatePwdRsp, error)
	VerifyPwd(ctx context.Context, in *VerifyPwdReq, opts ...client.CallOption) (*VerifyPwdRsp, error)
	// token
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...client.CallOption) (*GenerateTokenRsp, error)
	ParseToken(ctx context.Context, in *ParseTokenReq, opts ...client.CallOption) (*ParseTokenRsp, error)
	UserOauthLogin(ctx context.Context, in *UserOauthLoginReq, opts ...client.CallOption) (*UserOauthLoginRsp, error)
	// FriendHelp 好友助力
	AddFriendHelp(ctx context.Context, in *AddFriendHelpReq, opts ...client.CallOption) (*AddFriendHelpRsp, error)
	GetFriendHelpListByUser(ctx context.Context, in *GetFriendHelpListByUserReq, opts ...client.CallOption) (*GetFriendHelpListByUserRsp, error)
	// 每日任务
	GetDailyTaskList(ctx context.Context, in *GetDailyTaskListReq, opts ...client.CallOption) (*GetDailyTaskListRsp, error)
	// 发送验证码
	SendCodeSms(ctx context.Context, in *SendCodeSmsReq, opts ...client.CallOption) (*SendCodeSmsRsp, error)
	VerifyCodeSms(ctx context.Context, in *VerifyCodeSmsReq, opts ...client.CallOption) (*VerifyCodeSmsRsp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mu.micro.book.user.srv.service"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetFromUid(ctx context.Context, in *GetFromUidReq, opts ...client.CallOption) (*GetFromUidRsp, error) {
	req := c.c.NewRequest(c.name, "User.GetFromUid", in)
	out := new(GetFromUidRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetFromPhone(ctx context.Context, in *GetFromPhoneReq, opts ...client.CallOption) (*GetFromPhoneRsp, error) {
	req := c.c.NewRequest(c.name, "User.GetFromPhone", in)
	out := new(GetFromPhoneRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRsp, error) {
	req := c.c.NewRequest(c.name, "User.AddUser", in)
	out := new(AddUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...client.CallOption) (*UpdateUserRsp, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUser", in)
	out := new(UpdateUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...client.CallOption) (*UpdatePwdRsp, error) {
	req := c.c.NewRequest(c.name, "User.UpdatePwd", in)
	out := new(UpdatePwdRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) VerifyPwd(ctx context.Context, in *VerifyPwdReq, opts ...client.CallOption) (*VerifyPwdRsp, error) {
	req := c.c.NewRequest(c.name, "User.VerifyPwd", in)
	out := new(VerifyPwdRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...client.CallOption) (*GenerateTokenRsp, error) {
	req := c.c.NewRequest(c.name, "User.GenerateToken", in)
	out := new(GenerateTokenRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...client.CallOption) (*ParseTokenRsp, error) {
	req := c.c.NewRequest(c.name, "User.ParseToken", in)
	out := new(ParseTokenRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserOauthLogin(ctx context.Context, in *UserOauthLoginReq, opts ...client.CallOption) (*UserOauthLoginRsp, error) {
	req := c.c.NewRequest(c.name, "User.UserOauthLogin", in)
	out := new(UserOauthLoginRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddFriendHelp(ctx context.Context, in *AddFriendHelpReq, opts ...client.CallOption) (*AddFriendHelpRsp, error) {
	req := c.c.NewRequest(c.name, "User.AddFriendHelp", in)
	out := new(AddFriendHelpRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetFriendHelpListByUser(ctx context.Context, in *GetFriendHelpListByUserReq, opts ...client.CallOption) (*GetFriendHelpListByUserRsp, error) {
	req := c.c.NewRequest(c.name, "User.GetFriendHelpListByUser", in)
	out := new(GetFriendHelpListByUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetDailyTaskList(ctx context.Context, in *GetDailyTaskListReq, opts ...client.CallOption) (*GetDailyTaskListRsp, error) {
	req := c.c.NewRequest(c.name, "User.GetDailyTaskList", in)
	out := new(GetDailyTaskListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SendCodeSms(ctx context.Context, in *SendCodeSmsReq, opts ...client.CallOption) (*SendCodeSmsRsp, error) {
	req := c.c.NewRequest(c.name, "User.SendCodeSms", in)
	out := new(SendCodeSmsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) VerifyCodeSms(ctx context.Context, in *VerifyCodeSmsReq, opts ...client.CallOption) (*VerifyCodeSmsRsp, error) {
	req := c.c.NewRequest(c.name, "User.VerifyCodeSms", in)
	out := new(VerifyCodeSmsRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// user
	GetFromUid(context.Context, *GetFromUidReq, *GetFromUidRsp) error
	GetFromPhone(context.Context, *GetFromPhoneReq, *GetFromPhoneRsp) error
	AddUser(context.Context, *AddUserReq, *AddUserRsp) error
	UpdateUser(context.Context, *UpdateUserReq, *UpdateUserRsp) error
	UpdatePwd(context.Context, *UpdatePwdReq, *UpdatePwdRsp) error
	VerifyPwd(context.Context, *VerifyPwdReq, *VerifyPwdRsp) error
	// token
	GenerateToken(context.Context, *GenerateTokenReq, *GenerateTokenRsp) error
	ParseToken(context.Context, *ParseTokenReq, *ParseTokenRsp) error
	UserOauthLogin(context.Context, *UserOauthLoginReq, *UserOauthLoginRsp) error
	// FriendHelp 好友助力
	AddFriendHelp(context.Context, *AddFriendHelpReq, *AddFriendHelpRsp) error
	GetFriendHelpListByUser(context.Context, *GetFriendHelpListByUserReq, *GetFriendHelpListByUserRsp) error
	// 每日任务
	GetDailyTaskList(context.Context, *GetDailyTaskListReq, *GetDailyTaskListRsp) error
	// 发送验证码
	SendCodeSms(context.Context, *SendCodeSmsReq, *SendCodeSmsRsp) error
	VerifyCodeSms(context.Context, *VerifyCodeSmsReq, *VerifyCodeSmsRsp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		GetFromUid(ctx context.Context, in *GetFromUidReq, out *GetFromUidRsp) error
		GetFromPhone(ctx context.Context, in *GetFromPhoneReq, out *GetFromPhoneRsp) error
		AddUser(ctx context.Context, in *AddUserReq, out *AddUserRsp) error
		UpdateUser(ctx context.Context, in *UpdateUserReq, out *UpdateUserRsp) error
		UpdatePwd(ctx context.Context, in *UpdatePwdReq, out *UpdatePwdRsp) error
		VerifyPwd(ctx context.Context, in *VerifyPwdReq, out *VerifyPwdRsp) error
		GenerateToken(ctx context.Context, in *GenerateTokenReq, out *GenerateTokenRsp) error
		ParseToken(ctx context.Context, in *ParseTokenReq, out *ParseTokenRsp) error
		UserOauthLogin(ctx context.Context, in *UserOauthLoginReq, out *UserOauthLoginRsp) error
		AddFriendHelp(ctx context.Context, in *AddFriendHelpReq, out *AddFriendHelpRsp) error
		GetFriendHelpListByUser(ctx context.Context, in *GetFriendHelpListByUserReq, out *GetFriendHelpListByUserRsp) error
		GetDailyTaskList(ctx context.Context, in *GetDailyTaskListReq, out *GetDailyTaskListRsp) error
		SendCodeSms(ctx context.Context, in *SendCodeSmsReq, out *SendCodeSmsRsp) error
		VerifyCodeSms(ctx context.Context, in *VerifyCodeSmsReq, out *VerifyCodeSmsRsp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) GetFromUid(ctx context.Context, in *GetFromUidReq, out *GetFromUidRsp) error {
	return h.UserHandler.GetFromUid(ctx, in, out)
}

func (h *userHandler) GetFromPhone(ctx context.Context, in *GetFromPhoneReq, out *GetFromPhoneRsp) error {
	return h.UserHandler.GetFromPhone(ctx, in, out)
}

func (h *userHandler) AddUser(ctx context.Context, in *AddUserReq, out *AddUserRsp) error {
	return h.UserHandler.AddUser(ctx, in, out)
}

func (h *userHandler) UpdateUser(ctx context.Context, in *UpdateUserReq, out *UpdateUserRsp) error {
	return h.UserHandler.UpdateUser(ctx, in, out)
}

func (h *userHandler) UpdatePwd(ctx context.Context, in *UpdatePwdReq, out *UpdatePwdRsp) error {
	return h.UserHandler.UpdatePwd(ctx, in, out)
}

func (h *userHandler) VerifyPwd(ctx context.Context, in *VerifyPwdReq, out *VerifyPwdRsp) error {
	return h.UserHandler.VerifyPwd(ctx, in, out)
}

func (h *userHandler) GenerateToken(ctx context.Context, in *GenerateTokenReq, out *GenerateTokenRsp) error {
	return h.UserHandler.GenerateToken(ctx, in, out)
}

func (h *userHandler) ParseToken(ctx context.Context, in *ParseTokenReq, out *ParseTokenRsp) error {
	return h.UserHandler.ParseToken(ctx, in, out)
}

func (h *userHandler) UserOauthLogin(ctx context.Context, in *UserOauthLoginReq, out *UserOauthLoginRsp) error {
	return h.UserHandler.UserOauthLogin(ctx, in, out)
}

func (h *userHandler) AddFriendHelp(ctx context.Context, in *AddFriendHelpReq, out *AddFriendHelpRsp) error {
	return h.UserHandler.AddFriendHelp(ctx, in, out)
}

func (h *userHandler) GetFriendHelpListByUser(ctx context.Context, in *GetFriendHelpListByUserReq, out *GetFriendHelpListByUserRsp) error {
	return h.UserHandler.GetFriendHelpListByUser(ctx, in, out)
}

func (h *userHandler) GetDailyTaskList(ctx context.Context, in *GetDailyTaskListReq, out *GetDailyTaskListRsp) error {
	return h.UserHandler.GetDailyTaskList(ctx, in, out)
}

func (h *userHandler) SendCodeSms(ctx context.Context, in *SendCodeSmsReq, out *SendCodeSmsRsp) error {
	return h.UserHandler.SendCodeSms(ctx, in, out)
}

func (h *userHandler) VerifyCodeSms(ctx context.Context, in *VerifyCodeSmsReq, out *VerifyCodeSmsRsp) error {
	return h.UserHandler.VerifyCodeSms(ctx, in, out)
}
