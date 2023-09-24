// Code generated by MockGen. DO NOT EDIT.
// Source: ecom/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen.exe -destination db/mock/mock.go -package mockdb ecom/db/sqlc Store
//
// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sqlc "ecom/db/sqlc"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockStore) AddAddress(arg0 context.Context, arg1 sqlc.AddAddressParams) (sqlc.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockStoreMockRecorder) AddAddress(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockStore)(nil).AddAddress), arg0, arg1)
}

// AddAddressTx mocks base method.
func (m *MockStore) AddAddressTx(arg0 context.Context, arg1 sqlc.AddAddressesTxParams) (sqlc.AddAddressesTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddressTx", arg0, arg1)
	ret0, _ := ret[0].(sqlc.AddAddressesTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddressTx indicates an expected call of AddAddressTx.
func (mr *MockStoreMockRecorder) AddAddressTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddressTx", reflect.TypeOf((*MockStore)(nil).AddAddressTx), arg0, arg1)
}

// CreateCart mocks base method.
func (m *MockStore) CreateCart(arg0 context.Context, arg1 string) (sqlc.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockStoreMockRecorder) CreateCart(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockStore)(nil).CreateCart), arg0, arg1)
}

// CreateCartItem mocks base method.
func (m *MockStore) CreateCartItem(arg0 context.Context, arg1 sqlc.CreateCartItemParams) (sqlc.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCartItem indicates an expected call of CreateCartItem.
func (mr *MockStoreMockRecorder) CreateCartItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartItem", reflect.TypeOf((*MockStore)(nil).CreateCartItem), arg0, arg1)
}

// CreateItem mocks base method.
func (m *MockStore) CreateItem(arg0 context.Context, arg1 sqlc.CreateItemParams) (sqlc.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockStoreMockRecorder) CreateItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockStore)(nil).CreateItem), arg0, arg1)
}

// CreateItemImage mocks base method.
func (m *MockStore) CreateItemImage(arg0 context.Context, arg1 sqlc.CreateItemImageParams) (sqlc.ItemImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItemImage", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ItemImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItemImage indicates an expected call of CreateItemImage.
func (mr *MockStoreMockRecorder) CreateItemImage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItemImage", reflect.TypeOf((*MockStore)(nil).CreateItemImage), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 sqlc.CreateSessionParams) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAddress mocks base method.
func (m *MockStore) DeleteAddress(arg0 context.Context, arg1 int32) (sqlc.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockStoreMockRecorder) DeleteAddress(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockStore)(nil).DeleteAddress), arg0, arg1)
}

// DeleteCartItem mocks base method.
func (m *MockStore) DeleteCartItem(arg0 context.Context, arg1 sqlc.DeleteCartItemParams) (sqlc.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCartItem indicates an expected call of DeleteCartItem.
func (mr *MockStoreMockRecorder) DeleteCartItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartItem", reflect.TypeOf((*MockStore)(nil).DeleteCartItem), arg0, arg1)
}

// DeleteItem mocks base method.
func (m *MockStore) DeleteItem(arg0 context.Context, arg1 int32) (sqlc.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockStoreMockRecorder) DeleteItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockStore)(nil).DeleteItem), arg0, arg1)
}

// DeleteItemImage mocks base method.
func (m *MockStore) DeleteItemImage(arg0 context.Context, arg1 int32) (sqlc.ItemImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItemImage", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ItemImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItemImage indicates an expected call of DeleteItemImage.
func (mr *MockStoreMockRecorder) DeleteItemImage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemImage", reflect.TypeOf((*MockStore)(nil).DeleteItemImage), arg0, arg1)
}

// DeleteSessionById mocks base method.
func (m *MockStore) DeleteSessionById(arg0 context.Context, arg1 uuid.UUID) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionById", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSessionById indicates an expected call of DeleteSessionById.
func (mr *MockStoreMockRecorder) DeleteSessionById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionById", reflect.TypeOf((*MockStore)(nil).DeleteSessionById), arg0, arg1)
}

// DeleteSessionByUsername mocks base method.
func (m *MockStore) DeleteSessionByUsername(arg0 context.Context, arg1 string) ([]sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionByUsername", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSessionByUsername indicates an expected call of DeleteSessionByUsername.
func (mr *MockStoreMockRecorder) DeleteSessionByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionByUsername", reflect.TypeOf((*MockStore)(nil).DeleteSessionByUsername), arg0, arg1)
}

// GetAddresses mocks base method.
func (m *MockStore) GetAddresses(arg0 context.Context, arg1 string) ([]sqlc.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockStoreMockRecorder) GetAddresses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockStore)(nil).GetAddresses), arg0, arg1)
}

// GetCart mocks base method.
func (m *MockStore) GetCart(arg0 context.Context, arg1 string) (sqlc.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockStoreMockRecorder) GetCart(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockStore)(nil).GetCart), arg0, arg1)
}

// GetCartItemFromCartID mocks base method.
func (m *MockStore) GetCartItemFromCartID(arg0 context.Context, arg1 int32) ([]sqlc.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItemFromCartID", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItemFromCartID indicates an expected call of GetCartItemFromCartID.
func (mr *MockStoreMockRecorder) GetCartItemFromCartID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItemFromCartID", reflect.TypeOf((*MockStore)(nil).GetCartItemFromCartID), arg0, arg1)
}

// GetItemById mocks base method.
func (m *MockStore) GetItemById(arg0 context.Context, arg1 int32) (sqlc.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemById", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemById indicates an expected call of GetItemById.
func (mr *MockStoreMockRecorder) GetItemById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemById", reflect.TypeOf((*MockStore)(nil).GetItemById), arg0, arg1)
}

// GetItemImagesFromItemId mocks base method.
func (m *MockStore) GetItemImagesFromItemId(arg0 context.Context, arg1 int32) ([]sqlc.ItemImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemImagesFromItemId", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.ItemImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemImagesFromItemId indicates an expected call of GetItemImagesFromItemId.
func (mr *MockStoreMockRecorder) GetItemImagesFromItemId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemImagesFromItemId", reflect.TypeOf((*MockStore)(nil).GetItemImagesFromItemId), arg0, arg1)
}

// GetSessionFromId mocks base method.
func (m *MockStore) GetSessionFromId(arg0 context.Context, arg1 uuid.UUID) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionFromId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionFromId indicates an expected call of GetSessionFromId.
func (mr *MockStoreMockRecorder) GetSessionFromId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionFromId", reflect.TypeOf((*MockStore)(nil).GetSessionFromId), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// UpdateAddress mocks base method.
func (m *MockStore) UpdateAddress(arg0 context.Context, arg1 sqlc.UpdateAddressParams) (sqlc.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockStoreMockRecorder) UpdateAddress(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockStore)(nil).UpdateAddress), arg0, arg1)
}

// UpdateCartAmount mocks base method.
func (m *MockStore) UpdateCartAmount(arg0 context.Context, arg1 sqlc.UpdateCartAmountParams) (sqlc.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartAmount", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartAmount indicates an expected call of UpdateCartAmount.
func (mr *MockStoreMockRecorder) UpdateCartAmount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartAmount", reflect.TypeOf((*MockStore)(nil).UpdateCartAmount), arg0, arg1)
}

// UpdateCartItem mocks base method.
func (m *MockStore) UpdateCartItem(arg0 context.Context, arg1 sqlc.UpdateCartItemParams) (sqlc.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartItem indicates an expected call of UpdateCartItem.
func (mr *MockStoreMockRecorder) UpdateCartItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartItem", reflect.TypeOf((*MockStore)(nil).UpdateCartItem), arg0, arg1)
}

// UpdateItem mocks base method.
func (m *MockStore) UpdateItem(arg0 context.Context, arg1 sqlc.UpdateItemParams) (sqlc.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockStoreMockRecorder) UpdateItem(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockStore)(nil).UpdateItem), arg0, arg1)
}

// UpdateItemImageURL mocks base method.
func (m *MockStore) UpdateItemImageURL(arg0 context.Context, arg1 sqlc.UpdateItemImageURLParams) (sqlc.ItemImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemImageURL", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ItemImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItemImageURL indicates an expected call of UpdateItemImageURL.
func (mr *MockStoreMockRecorder) UpdateItemImageURL(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemImageURL", reflect.TypeOf((*MockStore)(nil).UpdateItemImageURL), arg0, arg1)
}

// UpdateSession mocks base method.
func (m *MockStore) UpdateSession(arg0 context.Context, arg1 sqlc.UpdateSessionParams) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockStoreMockRecorder) UpdateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 sqlc.UpdateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}