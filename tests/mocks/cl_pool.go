// Code generated by MockGen. DO NOT EDIT.
// Source: x/concentrated-liquidity/types/cl_pool_extensionI.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types0 "github.com/four4two/merlin/v17/x/poolmanager/types"
)

// MockConcentratedPoolExtension is a mock of ConcentratedPoolExtension interface.
type MockConcentratedPoolExtension struct {
	ctrl     *gomock.Controller
	recorder *MockConcentratedPoolExtensionMockRecorder
}

// MockConcentratedPoolExtensionMockRecorder is the mock recorder for MockConcentratedPoolExtension.
type MockConcentratedPoolExtensionMockRecorder struct {
	mock *MockConcentratedPoolExtension
}

// NewMockConcentratedPoolExtension creates a new mock instance.
func NewMockConcentratedPoolExtension(ctrl *gomock.Controller) *MockConcentratedPoolExtension {
	mock := &MockConcentratedPoolExtension{ctrl: ctrl}
	mock.recorder = &MockConcentratedPoolExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConcentratedPoolExtension) EXPECT() *MockConcentratedPoolExtensionMockRecorder {
	return m.recorder
}

// ApplySwap mocks base method.
func (m *MockConcentratedPoolExtension) ApplySwap(newLiquidity types.Dec, newCurrentTick int64, newCurrentSqrtPrice types.Dec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplySwap", newLiquidity, newCurrentTick, newCurrentSqrtPrice)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplySwap indicates an expected call of ApplySwap.
func (mr *MockConcentratedPoolExtensionMockRecorder) ApplySwap(newLiquidity, newCurrentTick, newCurrentSqrtPrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplySwap", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).ApplySwap), newLiquidity, newCurrentTick, newCurrentSqrtPrice)
}

// AsSerializablePool mocks base method.
func (m *MockConcentratedPoolExtension) AsSerializablePool() types0.PoolI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsSerializablePool")
	ret0, _ := ret[0].(types0.PoolI)
	return ret0
}

// AsSerializablePool indicates an expected call of AsSerializablePool.
func (mr *MockConcentratedPoolExtensionMockRecorder) AsSerializablePool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsSerializablePool", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).AsSerializablePool))
}

// CalcActualAmounts mocks base method.
func (m *MockConcentratedPoolExtension) CalcActualAmounts(ctx types.Context, lowerTick, upperTick int64, liquidityDelta types.Dec) (types.Dec, types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcActualAmounts", ctx, lowerTick, upperTick, liquidityDelta)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(types.Dec)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CalcActualAmounts indicates an expected call of CalcActualAmounts.
func (mr *MockConcentratedPoolExtensionMockRecorder) CalcActualAmounts(ctx, lowerTick, upperTick, liquidityDelta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcActualAmounts", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).CalcActualAmounts), ctx, lowerTick, upperTick, liquidityDelta)
}

// GetAddress mocks base method.
func (m *MockConcentratedPoolExtension) GetAddress() types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress")
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetAddress))
}

// GetCurrentSqrtPrice mocks base method.
func (m *MockConcentratedPoolExtension) GetCurrentSqrtPrice() types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSqrtPrice")
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// GetCurrentSqrtPrice indicates an expected call of GetCurrentSqrtPrice.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetCurrentSqrtPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSqrtPrice", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetCurrentSqrtPrice))
}

// GetCurrentTick mocks base method.
func (m *MockConcentratedPoolExtension) GetCurrentTick() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTick")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCurrentTick indicates an expected call of GetCurrentTick.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetCurrentTick() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTick", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetCurrentTick))
}

// GetExponentAtPriceOne mocks base method.
func (m *MockConcentratedPoolExtension) GetExponentAtPriceOne() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExponentAtPriceOne")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetExponentAtPriceOne indicates an expected call of GetExponentAtPriceOne.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetExponentAtPriceOne() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExponentAtPriceOne", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetExponentAtPriceOne))
}

// GetId mocks base method.
func (m *MockConcentratedPoolExtension) GetId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetId))
}

// GetIncentivesAddress mocks base method.
func (m *MockConcentratedPoolExtension) GetIncentivesAddress() types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncentivesAddress")
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetIncentivesAddress indicates an expected call of GetIncentivesAddress.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetIncentivesAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncentivesAddress", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetIncentivesAddress))
}

// GetLastLiquidityUpdate mocks base method.
func (m *MockConcentratedPoolExtension) GetLastLiquidityUpdate() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastLiquidityUpdate")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastLiquidityUpdate indicates an expected call of GetLastLiquidityUpdate.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetLastLiquidityUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastLiquidityUpdate", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetLastLiquidityUpdate))
}

// GetLiquidity mocks base method.
func (m *MockConcentratedPoolExtension) GetLiquidity() types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiquidity")
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// GetLiquidity indicates an expected call of GetLiquidity.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetLiquidity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiquidity", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetLiquidity))
}

// GetSpreadFactor mocks base method.
func (m *MockConcentratedPoolExtension) GetSpreadFactor(ctx types.Context) types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpreadFactor", ctx)
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// GetSpreadFactor indicates an expected call of GetSpreadFactor.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetSpreadFactor(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpreadFactor", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetSpreadFactor), ctx)
}

// GetSpreadRewardsAddress mocks base method.
func (m *MockConcentratedPoolExtension) GetSpreadRewardsAddress() types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpreadRewardsAddress")
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetSpreadRewardsAddress indicates an expected call of GetSpreadRewardsAddress.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetSpreadRewardsAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpreadRewardsAddress", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetSpreadRewardsAddress))
}

// GetTickSpacing mocks base method.
func (m *MockConcentratedPoolExtension) GetTickSpacing() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTickSpacing")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetTickSpacing indicates an expected call of GetTickSpacing.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetTickSpacing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTickSpacing", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetTickSpacing))
}

// GetToken0 mocks base method.
func (m *MockConcentratedPoolExtension) GetToken0() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken0")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetToken0 indicates an expected call of GetToken0.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetToken0() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken0", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetToken0))
}

// GetToken1 mocks base method.
func (m *MockConcentratedPoolExtension) GetToken1() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken1")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetToken1 indicates an expected call of GetToken1.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetToken1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken1", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetToken1))
}

// GetType mocks base method.
func (m *MockConcentratedPoolExtension) GetType() types0.PoolType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(types0.PoolType)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockConcentratedPoolExtensionMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).GetType))
}

// IsActive mocks base method.
func (m *MockConcentratedPoolExtension) IsActive(ctx types.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockConcentratedPoolExtensionMockRecorder) IsActive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).IsActive), ctx)
}

// IsCurrentTickInRange mocks base method.
func (m *MockConcentratedPoolExtension) IsCurrentTickInRange(lowerTick, upperTick int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrentTickInRange", lowerTick, upperTick)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrentTickInRange indicates an expected call of IsCurrentTickInRange.
func (mr *MockConcentratedPoolExtensionMockRecorder) IsCurrentTickInRange(lowerTick, upperTick interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrentTickInRange", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).IsCurrentTickInRange), lowerTick, upperTick)
}

// ProtoMessage mocks base method.
func (m *MockConcentratedPoolExtension) ProtoMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProtoMessage")
}

// ProtoMessage indicates an expected call of ProtoMessage.
func (mr *MockConcentratedPoolExtensionMockRecorder) ProtoMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoMessage", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).ProtoMessage))
}

// Reset mocks base method.
func (m *MockConcentratedPoolExtension) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockConcentratedPoolExtensionMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).Reset))
}

// SetCurrentSqrtPrice mocks base method.
func (m *MockConcentratedPoolExtension) SetCurrentSqrtPrice(newSqrtPrice types.Dec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentSqrtPrice", newSqrtPrice)
}

// SetCurrentSqrtPrice indicates an expected call of SetCurrentSqrtPrice.
func (mr *MockConcentratedPoolExtensionMockRecorder) SetCurrentSqrtPrice(newSqrtPrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentSqrtPrice", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).SetCurrentSqrtPrice), newSqrtPrice)
}

// SetCurrentTick mocks base method.
func (m *MockConcentratedPoolExtension) SetCurrentTick(newTick int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentTick", newTick)
}

// SetCurrentTick indicates an expected call of SetCurrentTick.
func (mr *MockConcentratedPoolExtensionMockRecorder) SetCurrentTick(newTick interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentTick", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).SetCurrentTick), newTick)
}

// SetLastLiquidityUpdate mocks base method.
func (m *MockConcentratedPoolExtension) SetLastLiquidityUpdate(newTime time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastLiquidityUpdate", newTime)
}

// SetLastLiquidityUpdate indicates an expected call of SetLastLiquidityUpdate.
func (mr *MockConcentratedPoolExtensionMockRecorder) SetLastLiquidityUpdate(newTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastLiquidityUpdate", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).SetLastLiquidityUpdate), newTime)
}

// SetTickSpacing mocks base method.
func (m *MockConcentratedPoolExtension) SetTickSpacing(newTickSpacing uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTickSpacing", newTickSpacing)
}

// SetTickSpacing indicates an expected call of SetTickSpacing.
func (mr *MockConcentratedPoolExtensionMockRecorder) SetTickSpacing(newTickSpacing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTickSpacing", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).SetTickSpacing), newTickSpacing)
}

// SpotPrice mocks base method.
func (m *MockConcentratedPoolExtension) SpotPrice(ctx types.Context, quoteAssetDenom, baseAssetDenom string) (types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpotPrice", ctx, quoteAssetDenom, baseAssetDenom)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpotPrice indicates an expected call of SpotPrice.
func (mr *MockConcentratedPoolExtensionMockRecorder) SpotPrice(ctx, quoteAssetDenom, baseAssetDenom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpotPrice", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).SpotPrice), ctx, quoteAssetDenom, baseAssetDenom)
}

// String mocks base method.
func (m *MockConcentratedPoolExtension) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockConcentratedPoolExtensionMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).String))
}

// UpdateLiquidity mocks base method.
func (m *MockConcentratedPoolExtension) UpdateLiquidity(newLiquidity types.Dec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateLiquidity", newLiquidity)
}

// UpdateLiquidity indicates an expected call of UpdateLiquidity.
func (mr *MockConcentratedPoolExtensionMockRecorder) UpdateLiquidity(newLiquidity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLiquidity", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).UpdateLiquidity), newLiquidity)
}

// UpdateLiquidityIfActivePosition mocks base method.
func (m *MockConcentratedPoolExtension) UpdateLiquidityIfActivePosition(ctx types.Context, lowerTick, upperTick int64, liquidityDelta types.Dec) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLiquidityIfActivePosition", ctx, lowerTick, upperTick, liquidityDelta)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateLiquidityIfActivePosition indicates an expected call of UpdateLiquidityIfActivePosition.
func (mr *MockConcentratedPoolExtensionMockRecorder) UpdateLiquidityIfActivePosition(ctx, lowerTick, upperTick, liquidityDelta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLiquidityIfActivePosition", reflect.TypeOf((*MockConcentratedPoolExtension)(nil).UpdateLiquidityIfActivePosition), ctx, lowerTick, upperTick, liquidityDelta)
}
