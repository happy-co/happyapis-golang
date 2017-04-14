// Code generated by protoc-gen-go.
// source: happyco/manage/account/v1/account.proto
// DO NOT EDIT!

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	happyco/manage/account/v1/account.proto

It has these top-level messages:
	Account
	ListAccountsRequest
	ListAccountsResponse
	UpdateAccountsRequest
	UpdateAccountsResponse
	User
	ListUsersRequest
	ListUsersResponse
	AddUsersRequest
	AddUsersResponse
	UpdateUsersRequest
	UpdateUsersResponse
	UpdateUserStatusesRequest
	UpdateUserStatusesResponse
*/
package v1

import proto "github.com/happy-co/happyapis-golang/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import happyco_type_v1 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import happyco_type_v11 "github.com/happy-co/happyapis-golang/happyco/type/v1"
import _ "github.com/happy-co/happyapis-golang/google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Id            *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Inactive      bool                           `protobuf:"varint,2,opt,name=inactive" json:"inactive,omitempty"`
	Name          string                         `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Phone         string                         `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	AddressLine_1 string                         `protobuf:"bytes,5,opt,name=address_line_1,json=addressLine1" json:"address_line_1,omitempty"`
	AddressLine_2 string                         `protobuf:"bytes,6,opt,name=address_line_2,json=addressLine2" json:"address_line_2,omitempty"`
	City          string                         `protobuf:"bytes,7,opt,name=city" json:"city,omitempty"`
	State         string                         `protobuf:"bytes,8,opt,name=state" json:"state,omitempty"`
	PostalCode    string                         `protobuf:"bytes,9,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	Country       string                         `protobuf:"bytes,10,opt,name=country" json:"country,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Account) GetInactive() bool {
	if m != nil {
		return m.Inactive
	}
	return false
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Account) GetAddressLine_1() string {
	if m != nil {
		return m.AddressLine_1
	}
	return ""
}

func (m *Account) GetAddressLine_2() string {
	if m != nil {
		return m.AddressLine_2
	}
	return ""
}

func (m *Account) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Account) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Account) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Account) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type ListAccountsRequest struct {
	AccountIds      []*happyco_type_v1.IntegrationID `protobuf:"bytes,1,rep,name=account_ids,json=accountIds" json:"account_ids,omitempty"`
	IncludeInactive bool                             `protobuf:"varint,2,opt,name=include_inactive,json=includeInactive" json:"include_inactive,omitempty"`
	Paging          *happyco_type_v11.Paging         `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
}

func (m *ListAccountsRequest) Reset()                    { *m = ListAccountsRequest{} }
func (m *ListAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListAccountsRequest) ProtoMessage()               {}
func (*ListAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListAccountsRequest) GetAccountIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountIds
	}
	return nil
}

func (m *ListAccountsRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListAccountsRequest) GetPaging() *happyco_type_v11.Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

type ListAccountsResponse struct {
	Accounts []*Account                       `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	Paging   *happyco_type_v11.PagingResponse `protobuf:"bytes,2,opt,name=paging" json:"paging,omitempty"`
}

func (m *ListAccountsResponse) Reset()                    { *m = ListAccountsResponse{} }
func (m *ListAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListAccountsResponse) ProtoMessage()               {}
func (*ListAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListAccountsResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ListAccountsResponse) GetPaging() *happyco_type_v11.PagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

type UpdateAccountsRequest struct {
	AccountWithAdminIds []*UpdateAccountsRequest_AccountWithAdminId `protobuf:"bytes,1,rep,name=account_with_admin_ids,json=accountWithAdminIds" json:"account_with_admin_ids,omitempty"`
}

func (m *UpdateAccountsRequest) Reset()                    { *m = UpdateAccountsRequest{} }
func (m *UpdateAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccountsRequest) ProtoMessage()               {}
func (*UpdateAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateAccountsRequest) GetAccountWithAdminIds() []*UpdateAccountsRequest_AccountWithAdminId {
	if m != nil {
		return m.AccountWithAdminIds
	}
	return nil
}

type UpdateAccountsRequest_AccountWithAdminId struct {
	Account     *Account                       `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	AdminUserId *happyco_type_v1.IntegrationID `protobuf:"bytes,2,opt,name=admin_user_id,json=adminUserId" json:"admin_user_id,omitempty"`
}

func (m *UpdateAccountsRequest_AccountWithAdminId) Reset() {
	*m = UpdateAccountsRequest_AccountWithAdminId{}
}
func (m *UpdateAccountsRequest_AccountWithAdminId) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountsRequest_AccountWithAdminId) ProtoMessage()    {}
func (*UpdateAccountsRequest_AccountWithAdminId) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *UpdateAccountsRequest_AccountWithAdminId) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *UpdateAccountsRequest_AccountWithAdminId) GetAdminUserId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AdminUserId
	}
	return nil
}

type UpdateAccountsResponse struct {
}

func (m *UpdateAccountsResponse) Reset()                    { *m = UpdateAccountsResponse{} }
func (m *UpdateAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccountsResponse) ProtoMessage()               {}
func (*UpdateAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type User struct {
	Id        *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Inactive  bool                           `protobuf:"varint,2,opt,name=inactive" json:"inactive,omitempty"`
	FirstName string                         `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string                         `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string                         `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Phone     string                         `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *User) GetId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetInactive() bool {
	if m != nil {
		return m.Inactive
	}
	return false
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ListUsersRequest struct {
	AccountId       *happyco_type_v1.IntegrationID   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserIds         []*happyco_type_v1.IntegrationID `protobuf:"bytes,2,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
	IncludeInactive bool                             `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive" json:"include_inactive,omitempty"`
	Paging          *happyco_type_v11.Paging         `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
}

func (m *ListUsersRequest) Reset()                    { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()               {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListUsersRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *ListUsersRequest) GetUserIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ListUsersRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListUsersRequest) GetPaging() *happyco_type_v11.Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

type ListUsersResponse struct {
	Users  []*User                          `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Paging *happyco_type_v11.PagingResponse `protobuf:"bytes,2,opt,name=paging" json:"paging,omitempty"`
}

func (m *ListUsersResponse) Reset()                    { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()               {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetPaging() *happyco_type_v11.PagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

type AddUsersRequest struct {
	AccountId *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Users     []*User                        `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *AddUsersRequest) Reset()                    { *m = AddUsersRequest{} }
func (m *AddUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUsersRequest) ProtoMessage()               {}
func (*AddUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddUsersRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *AddUsersRequest) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type AddUsersResponse struct {
	UserIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,1,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
}

func (m *AddUsersResponse) Reset()                    { *m = AddUsersResponse{} }
func (m *AddUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUsersResponse) ProtoMessage()               {}
func (*AddUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddUsersResponse) GetUserIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type UpdateUsersRequest struct {
	AccountId *happyco_type_v1.IntegrationID `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Users     []*User                        `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *UpdateUsersRequest) Reset()                    { *m = UpdateUsersRequest{} }
func (m *UpdateUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUsersRequest) ProtoMessage()               {}
func (*UpdateUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateUsersRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *UpdateUsersRequest) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateUsersResponse struct {
	UserIds []*happyco_type_v1.IntegrationID `protobuf:"bytes,1,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
}

func (m *UpdateUsersResponse) Reset()                    { *m = UpdateUsersResponse{} }
func (m *UpdateUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateUsersResponse) ProtoMessage()               {}
func (*UpdateUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateUsersResponse) GetUserIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type UpdateUserStatusesRequest struct {
	AccountId *happyco_type_v1.IntegrationID   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserIds   []*happyco_type_v1.IntegrationID `protobuf:"bytes,2,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
	Active    bool                             `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
}

func (m *UpdateUserStatusesRequest) Reset()                    { *m = UpdateUserStatusesRequest{} }
func (m *UpdateUserStatusesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserStatusesRequest) ProtoMessage()               {}
func (*UpdateUserStatusesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateUserStatusesRequest) GetAccountId() *happyco_type_v1.IntegrationID {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *UpdateUserStatusesRequest) GetUserIds() []*happyco_type_v1.IntegrationID {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *UpdateUserStatusesRequest) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type UpdateUserStatusesResponse struct {
}

func (m *UpdateUserStatusesResponse) Reset()                    { *m = UpdateUserStatusesResponse{} }
func (m *UpdateUserStatusesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserStatusesResponse) ProtoMessage()               {}
func (*UpdateUserStatusesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Account)(nil), "happyco.manage.account.v1.Account")
	proto.RegisterType((*ListAccountsRequest)(nil), "happyco.manage.account.v1.ListAccountsRequest")
	proto.RegisterType((*ListAccountsResponse)(nil), "happyco.manage.account.v1.ListAccountsResponse")
	proto.RegisterType((*UpdateAccountsRequest)(nil), "happyco.manage.account.v1.UpdateAccountsRequest")
	proto.RegisterType((*UpdateAccountsRequest_AccountWithAdminId)(nil), "happyco.manage.account.v1.UpdateAccountsRequest.AccountWithAdminId")
	proto.RegisterType((*UpdateAccountsResponse)(nil), "happyco.manage.account.v1.UpdateAccountsResponse")
	proto.RegisterType((*User)(nil), "happyco.manage.account.v1.User")
	proto.RegisterType((*ListUsersRequest)(nil), "happyco.manage.account.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "happyco.manage.account.v1.ListUsersResponse")
	proto.RegisterType((*AddUsersRequest)(nil), "happyco.manage.account.v1.AddUsersRequest")
	proto.RegisterType((*AddUsersResponse)(nil), "happyco.manage.account.v1.AddUsersResponse")
	proto.RegisterType((*UpdateUsersRequest)(nil), "happyco.manage.account.v1.UpdateUsersRequest")
	proto.RegisterType((*UpdateUsersResponse)(nil), "happyco.manage.account.v1.UpdateUsersResponse")
	proto.RegisterType((*UpdateUserStatusesRequest)(nil), "happyco.manage.account.v1.UpdateUserStatusesRequest")
	proto.RegisterType((*UpdateUserStatusesResponse)(nil), "happyco.manage.account.v1.UpdateUserStatusesResponse")
}

func init() { proto.RegisterFile("happyco/manage/account/v1/account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 975 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x97, 0x4d, 0x6f, 0x1c, 0x35,
	0x18, 0xc7, 0xe5, 0x49, 0xb2, 0x2f, 0xcf, 0x96, 0x34, 0x38, 0x25, 0x4c, 0xa6, 0xa1, 0xd9, 0x4e,
	0x0b, 0x4d, 0x93, 0x32, 0xd3, 0x5d, 0x28, 0x88, 0x8a, 0x17, 0xa5, 0xe5, 0xb2, 0x52, 0xa9, 0xaa,
	0xad, 0x2a, 0x24, 0x2e, 0x23, 0xb3, 0x36, 0x1b, 0x4b, 0x9b, 0x99, 0x61, 0xec, 0x5d, 0xc8, 0x15,
	0x0e, 0x48, 0x5c, 0x40, 0x02, 0x21, 0x84, 0x38, 0x70, 0x46, 0x42, 0x42, 0xdc, 0xe0, 0x43, 0x70,
	0xe1, 0x2b, 0x70, 0xe0, 0x63, 0xa0, 0xf1, 0xd8, 0xb3, 0xb3, 0x2f, 0xd9, 0xcc, 0x42, 0x0f, 0xb9,
	0x8d, 0xed, 0xe7, 0xf1, 0xf3, 0xf3, 0xdf, 0xfe, 0xdb, 0xbb, 0x70, 0xe3, 0x88, 0xc4, 0xf1, 0x49,
	0x2f, 0xf2, 0x8f, 0x49, 0x48, 0xfa, 0xcc, 0x27, 0xbd, 0x5e, 0x34, 0x0c, 0xa5, 0x3f, 0x6a, 0x99,
	0x4f, 0x2f, 0x4e, 0x22, 0x19, 0xe1, 0x6d, 0x1d, 0xe8, 0x65, 0x81, 0x9e, 0x19, 0x1d, 0xb5, 0x9c,
	0xeb, 0x66, 0x0e, 0x79, 0x12, 0xb3, 0x34, 0x93, 0x87, 0x92, 0xf5, 0x13, 0x22, 0x79, 0x14, 0x06,
	0x9c, 0x66, 0x13, 0x38, 0x3b, 0xd3, 0x51, 0x31, 0xe9, 0xf3, 0xb0, 0x6f, 0x46, 0xfb, 0x51, 0xd4,
	0x1f, 0x30, 0x9f, 0xc4, 0xdc, 0x27, 0x61, 0x18, 0x49, 0x95, 0x2d, 0xb2, 0x51, 0xf7, 0x57, 0x0b,
	0xaa, 0x87, 0x59, 0x41, 0xec, 0x81, 0xc5, 0xa9, 0x8d, 0x9a, 0x68, 0xaf, 0xd1, 0xbe, 0xe2, 0x19,
	0xaa, 0x74, 0x52, 0x6f, 0xd4, 0xf2, 0x3a, 0xe3, 0xd2, 0x9d, 0x77, 0xbb, 0x16, 0xa7, 0xd8, 0x81,
	0x1a, 0x0f, 0x49, 0x4f, 0xf2, 0x11, 0xb3, 0xad, 0x26, 0xda, 0xab, 0x75, 0xf3, 0x36, 0xc6, 0xb0,
	0x1a, 0x92, 0x63, 0x66, 0xaf, 0x34, 0xd1, 0x5e, 0xbd, 0xab, 0xbe, 0xf1, 0x25, 0x58, 0x8b, 0x8f,
	0xa2, 0x90, 0xd9, 0xab, 0xaa, 0x33, 0x6b, 0xe0, 0xeb, 0xb0, 0x4e, 0x28, 0x4d, 0x98, 0x10, 0xc1,
	0x80, 0x87, 0x2c, 0x68, 0xd9, 0x6b, 0x6a, 0xf8, 0x82, 0xee, 0x7d, 0xc0, 0x43, 0xd6, 0x9a, 0x89,
	0x6a, 0xdb, 0x95, 0x99, 0xa8, 0x76, 0x5a, 0xb5, 0xc7, 0xe5, 0x89, 0x5d, 0xcd, 0xaa, 0xa6, 0xdf,
	0x69, 0x55, 0x21, 0x89, 0x64, 0x76, 0x2d, 0xab, 0xaa, 0x1a, 0x78, 0x17, 0x1a, 0x71, 0x24, 0x24,
	0x19, 0x04, 0xbd, 0x88, 0x32, 0xbb, 0xae, 0xc6, 0x20, 0xeb, 0xba, 0x1f, 0x51, 0x86, 0x6d, 0xa8,
	0x2a, 0x55, 0x92, 0x13, 0x1b, 0xd4, 0xa0, 0x69, 0xba, 0xbf, 0x21, 0xd8, 0x7c, 0xc0, 0x85, 0xd4,
	0xb2, 0x89, 0x2e, 0xfb, 0x78, 0xc8, 0x84, 0xc4, 0xef, 0x40, 0x43, 0x6f, 0x5d, 0xc0, 0xa9, 0xb0,
	0x51, 0x73, 0xa5, 0x84, 0x8e, 0xa0, 0x53, 0x3a, 0x54, 0xe0, 0x9b, 0xb0, 0xc1, 0xc3, 0xde, 0x60,
	0x48, 0x59, 0x30, 0xa5, 0xeb, 0x45, 0xdd, 0xdf, 0x31, 0xf2, 0xfa, 0x50, 0xc9, 0x36, 0x59, 0x09,
	0xdc, 0x68, 0x3f, 0x3f, 0x53, 0xe6, 0x91, 0x1a, 0xee, 0xea, 0x30, 0xf7, 0x2b, 0x04, 0x97, 0x26,
	0xa1, 0x45, 0x1c, 0x85, 0x82, 0xe1, 0xb7, 0xa1, 0xa6, 0x11, 0x0c, 0xb2, 0xeb, 0x9d, 0x7a, 0x20,
	0x3d, 0x9d, 0xde, 0xcd, 0x73, 0xf0, 0xeb, 0x39, 0x89, 0xa5, 0x48, 0x76, 0x4f, 0x23, 0xd1, 0x05,
	0x73, 0xa2, 0x9f, 0x2d, 0x78, 0xee, 0x49, 0x4c, 0x89, 0x64, 0xd3, 0x42, 0x7e, 0x0a, 0x5b, 0x46,
	0xc8, 0x4f, 0xb8, 0x3c, 0x0a, 0x08, 0x3d, 0xe6, 0x61, 0x41, 0xd3, 0xfb, 0x0b, 0x00, 0xe7, 0xce,
	0x68, 0xb0, 0xdf, 0xe7, 0xf2, 0xe8, 0x30, 0x9d, 0xac, 0x43, 0xbb, 0x9b, 0x64, 0xa6, 0x4f, 0x38,
	0xdf, 0x21, 0xc0, 0xb3, 0xb1, 0xf8, 0x4d, 0xa8, 0xea, 0x68, 0xed, 0x8e, 0x32, 0x12, 0x99, 0x14,
	0x7c, 0x0f, 0x9e, 0xc9, 0x56, 0x30, 0x14, 0x2c, 0x09, 0x38, 0xd5, 0x42, 0x9d, 0x75, 0x32, 0x1a,
	0x2a, 0xe9, 0x89, 0x60, 0x49, 0x87, 0xba, 0x36, 0x6c, 0x4d, 0xaf, 0x2c, 0x93, 0xd3, 0xfd, 0x03,
	0xc1, 0x6a, 0x1a, 0xf4, 0x54, 0xdd, 0xfb, 0x02, 0xc0, 0x47, 0x3c, 0x11, 0x32, 0x28, 0x78, 0xb8,
	0xae, 0x7a, 0x1e, 0xa6, 0x46, 0xbe, 0x0c, 0xf5, 0x01, 0x31, 0xa3, 0x99, 0x99, 0x6b, 0x69, 0xc7,
	0x43, 0xed, 0x72, 0x76, 0x4c, 0xf8, 0x40, 0xdb, 0x38, 0x6b, 0x8c, 0xbd, 0x5f, 0x29, 0x78, 0xdf,
	0xfd, 0x07, 0xc1, 0x46, 0x7a, 0x2a, 0xd3, 0x05, 0xe4, 0xdb, 0xff, 0x16, 0xc0, 0xd8, 0x47, 0x25,
	0x17, 0x54, 0xcf, 0x6d, 0x84, 0xdf, 0x80, 0x9a, 0x16, 0x5a, 0xd8, 0x56, 0x29, 0x0f, 0x56, 0x87,
	0x4a, 0xe4, 0xf9, 0x06, 0x5c, 0x39, 0xcb, 0x80, 0xab, 0xe5, 0x0c, 0xf8, 0x39, 0x82, 0x67, 0x0b,
	0x4b, 0xd5, 0xee, 0xbb, 0x03, 0x6b, 0x69, 0x71, 0x73, 0xb2, 0x77, 0x17, 0x9d, 0x6c, 0xc1, 0x92,
	0x6e, 0x16, 0xfd, 0xdf, 0x4d, 0xf7, 0x05, 0x82, 0x8b, 0x87, 0x94, 0x3e, 0x4d, 0xbd, 0xf3, 0x25,
	0x58, 0xcb, 0x2c, 0xc1, 0x7d, 0x0f, 0x36, 0xc6, 0x20, 0x5a, 0x8d, 0xe2, 0xd6, 0xa1, 0xa5, 0xb6,
	0xce, 0xfd, 0x12, 0x01, 0xce, 0x1c, 0x72, 0x0e, 0xd6, 0xf6, 0x08, 0x36, 0x27, 0x58, 0xfe, 0xff,
	0xf2, 0x7e, 0x41, 0xb0, 0x3d, 0x9e, 0xf2, 0xb1, 0x24, 0x72, 0x28, 0xd8, 0x39, 0x70, 0xcc, 0x16,
	0x54, 0x26, 0x7c, 0xa2, 0x5b, 0xee, 0x0e, 0x38, 0xf3, 0x70, 0x33, 0x21, 0xda, 0x7f, 0x5a, 0xb0,
	0xae, 0x2f, 0xb2, 0xc7, 0x2c, 0x19, 0xf1, 0x1e, 0xc3, 0xdf, 0x23, 0xb8, 0x50, 0x7c, 0x9f, 0xb0,
	0xb7, 0x40, 0xeb, 0x39, 0xaf, 0xaf, 0xe3, 0x97, 0x8e, 0xd7, 0x17, 0xe7, 0x8d, 0xcf, 0xfe, 0xfa,
	0xfb, 0x1b, 0xeb, 0xaa, 0xbb, 0xa3, 0x7e, 0x19, 0x8d, 0x5a, 0x53, 0xbf, 0xd3, 0x84, 0x3f, 0xe0,
	0x42, 0xde, 0x45, 0xfb, 0xf8, 0x27, 0x04, 0xeb, 0x93, 0x97, 0x2f, 0xbe, 0xbd, 0xec, 0x0b, 0xe4,
	0xb4, 0x96, 0xc8, 0xd0, 0x80, 0x37, 0x15, 0xe0, 0x35, 0xf7, 0xca, 0x69, 0x80, 0x43, 0x95, 0x77,
	0x17, 0xed, 0xb7, 0x7f, 0x58, 0x83, 0x4d, 0x9d, 0xaf, 0x8e, 0x9c, 0x51, 0xf5, 0x5b, 0x04, 0xf5,
	0xfc, 0xd2, 0xc1, 0x07, 0x67, 0x48, 0x54, 0x74, 0x8e, 0x73, 0xab, 0x5c, 0xb0, 0x66, 0xbd, 0xa5,
	0x58, 0x5f, 0x72, 0xaf, 0xce, 0x67, 0xf5, 0x95, 0x2f, 0x72, 0x45, 0xbf, 0x46, 0x50, 0x33, 0xe6,
	0xc7, 0xfb, 0x8b, 0xde, 0xd2, 0xc9, 0xab, 0xca, 0x39, 0x28, 0x15, 0xab, 0x99, 0x0e, 0x14, 0xd3,
	0x8b, 0x6e, 0x73, 0x21, 0x13, 0xa1, 0x34, 0x45, 0xfa, 0x11, 0x41, 0xa3, 0xe0, 0x59, 0xfc, 0xf2,
	0x99, 0xfb, 0x35, 0x01, 0xe6, 0x95, 0x0d, 0xd7, 0x6c, 0x9e, 0x62, 0xdb, 0x73, 0xae, 0x2d, 0x64,
	0xcb, 0x37, 0x18, 0xff, 0x3e, 0x71, 0xbd, 0x19, 0x43, 0xe1, 0x57, 0x4b, 0x95, 0x9d, 0xba, 0x2e,
	0x9c, 0x3b, 0x4b, 0x66, 0x69, 0xe6, 0xd7, 0x14, 0xf3, 0x6d, 0xe7, 0x60, 0x01, 0xb3, 0x2f, 0x74,
	0xd6, 0x98, 0xfd, 0xde, 0xe5, 0x0f, 0xb6, 0x4f, 0xfd, 0x2b, 0xf4, 0x61, 0x45, 0xfd, 0x0d, 0x79,
	0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xc9, 0x64, 0xe5, 0x2e, 0x0d, 0x00, 0x00,
}
