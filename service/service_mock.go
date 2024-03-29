// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/MatsuoTakuro/fcoin-balances-manager/entity"
	"sync"
)

// Ensure, that RegisterUserServiceMock does implement RegisterUserService.
// If this is not the case, regenerate this file with moq.
var _ RegisterUserService = &RegisterUserServiceMock{}

// RegisterUserServiceMock is a mock implementation of RegisterUserService.
//
//	func TestSomethingThatUsesRegisterUserService(t *testing.T) {
//
//		// make and configure a mocked RegisterUserService
//		mockedRegisterUserService := &RegisterUserServiceMock{
//			RegisterUserFunc: func(ctx context.Context, name string) (*entity.User, *entity.Balance, error) {
//				panic("mock out the RegisterUser method")
//			},
//		}
//
//		// use mockedRegisterUserService in code that requires RegisterUserService
//		// and then make assertions.
//
//	}
type RegisterUserServiceMock struct {
	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, name string) (*entity.User, *entity.Balance, error)

	// calls tracks calls to the methods.
	calls struct {
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
	}
	lockRegisterUser sync.RWMutex
}

// RegisterUser calls RegisterUserFunc.
func (mock *RegisterUserServiceMock) RegisterUser(ctx context.Context, name string) (*entity.User, *entity.Balance, error) {
	if mock.RegisterUserFunc == nil {
		panic("RegisterUserServiceMock.RegisterUserFunc: method is nil but RegisterUserService.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, name)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedRegisterUserService.RegisterUserCalls())
func (mock *RegisterUserServiceMock) RegisterUserCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// Ensure, that UpdateBalanceServiceMock does implement UpdateBalanceService.
// If this is not the case, regenerate this file with moq.
var _ UpdateBalanceService = &UpdateBalanceServiceMock{}

// UpdateBalanceServiceMock is a mock implementation of UpdateBalanceService.
//
//	func TestSomethingThatUsesUpdateBalanceService(t *testing.T) {
//
//		// make and configure a mocked UpdateBalanceService
//		mockedUpdateBalanceService := &UpdateBalanceServiceMock{
//			UpdateBalanceFunc: func(ctx context.Context, userID entity.UserID, amount int32) (*entity.BalanceTrans, error) {
//				panic("mock out the UpdateBalance method")
//			},
//		}
//
//		// use mockedUpdateBalanceService in code that requires UpdateBalanceService
//		// and then make assertions.
//
//	}
type UpdateBalanceServiceMock struct {
	// UpdateBalanceFunc mocks the UpdateBalance method.
	UpdateBalanceFunc func(ctx context.Context, userID entity.UserID, amount int32) (*entity.BalanceTrans, error)

	// calls tracks calls to the methods.
	calls struct {
		// UpdateBalance holds details about calls to the UpdateBalance method.
		UpdateBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID entity.UserID
			// Amount is the amount argument value.
			Amount int32
		}
	}
	lockUpdateBalance sync.RWMutex
}

// UpdateBalance calls UpdateBalanceFunc.
func (mock *UpdateBalanceServiceMock) UpdateBalance(ctx context.Context, userID entity.UserID, amount int32) (*entity.BalanceTrans, error) {
	if mock.UpdateBalanceFunc == nil {
		panic("UpdateBalanceServiceMock.UpdateBalanceFunc: method is nil but UpdateBalanceService.UpdateBalance was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID entity.UserID
		Amount int32
	}{
		Ctx:    ctx,
		UserID: userID,
		Amount: amount,
	}
	mock.lockUpdateBalance.Lock()
	mock.calls.UpdateBalance = append(mock.calls.UpdateBalance, callInfo)
	mock.lockUpdateBalance.Unlock()
	return mock.UpdateBalanceFunc(ctx, userID, amount)
}

// UpdateBalanceCalls gets all the calls that were made to UpdateBalance.
// Check the length with:
//
//	len(mockedUpdateBalanceService.UpdateBalanceCalls())
func (mock *UpdateBalanceServiceMock) UpdateBalanceCalls() []struct {
	Ctx    context.Context
	UserID entity.UserID
	Amount int32
} {
	var calls []struct {
		Ctx    context.Context
		UserID entity.UserID
		Amount int32
	}
	mock.lockUpdateBalance.RLock()
	calls = mock.calls.UpdateBalance
	mock.lockUpdateBalance.RUnlock()
	return calls
}

// Ensure, that TransferCoinsServiceMock does implement TransferCoinsService.
// If this is not the case, regenerate this file with moq.
var _ TransferCoinsService = &TransferCoinsServiceMock{}

// TransferCoinsServiceMock is a mock implementation of TransferCoinsService.
//
//	func TestSomethingThatUsesTransferCoinsService(t *testing.T) {
//
//		// make and configure a mocked TransferCoinsService
//		mockedTransferCoinsService := &TransferCoinsServiceMock{
//			TransferCoinsFunc: func(ctx context.Context, fromUser entity.UserID, toUser entity.UserID, amount uint32) (*entity.BalanceTrans, error) {
//				panic("mock out the TransferCoins method")
//			},
//		}
//
//		// use mockedTransferCoinsService in code that requires TransferCoinsService
//		// and then make assertions.
//
//	}
type TransferCoinsServiceMock struct {
	// TransferCoinsFunc mocks the TransferCoins method.
	TransferCoinsFunc func(ctx context.Context, fromUser entity.UserID, toUser entity.UserID, amount uint32) (*entity.BalanceTrans, error)

	// calls tracks calls to the methods.
	calls struct {
		// TransferCoins holds details about calls to the TransferCoins method.
		TransferCoins []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// FromUser is the fromUser argument value.
			FromUser entity.UserID
			// ToUser is the toUser argument value.
			ToUser entity.UserID
			// Amount is the amount argument value.
			Amount uint32
		}
	}
	lockTransferCoins sync.RWMutex
}

// TransferCoins calls TransferCoinsFunc.
func (mock *TransferCoinsServiceMock) TransferCoins(ctx context.Context, fromUser entity.UserID, toUser entity.UserID, amount uint32) (*entity.BalanceTrans, error) {
	if mock.TransferCoinsFunc == nil {
		panic("TransferCoinsServiceMock.TransferCoinsFunc: method is nil but TransferCoinsService.TransferCoins was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		FromUser entity.UserID
		ToUser   entity.UserID
		Amount   uint32
	}{
		Ctx:      ctx,
		FromUser: fromUser,
		ToUser:   toUser,
		Amount:   amount,
	}
	mock.lockTransferCoins.Lock()
	mock.calls.TransferCoins = append(mock.calls.TransferCoins, callInfo)
	mock.lockTransferCoins.Unlock()
	return mock.TransferCoinsFunc(ctx, fromUser, toUser, amount)
}

// TransferCoinsCalls gets all the calls that were made to TransferCoins.
// Check the length with:
//
//	len(mockedTransferCoinsService.TransferCoinsCalls())
func (mock *TransferCoinsServiceMock) TransferCoinsCalls() []struct {
	Ctx      context.Context
	FromUser entity.UserID
	ToUser   entity.UserID
	Amount   uint32
} {
	var calls []struct {
		Ctx      context.Context
		FromUser entity.UserID
		ToUser   entity.UserID
		Amount   uint32
	}
	mock.lockTransferCoins.RLock()
	calls = mock.calls.TransferCoins
	mock.lockTransferCoins.RUnlock()
	return calls
}

// Ensure, that GetBalanceDetailsMock does implement GetBalanceDetails.
// If this is not the case, regenerate this file with moq.
var _ GetBalanceDetails = &GetBalanceDetailsMock{}

// GetBalanceDetailsMock is a mock implementation of GetBalanceDetails.
//
//	func TestSomethingThatUsesGetBalanceDetails(t *testing.T) {
//
//		// make and configure a mocked GetBalanceDetails
//		mockedGetBalanceDetails := &GetBalanceDetailsMock{
//			GetBalanceDetailsFunc: func(ctx context.Context, userID entity.UserID) (*entity.Balance, []*entity.BalanceTrans, error) {
//				panic("mock out the GetBalanceDetails method")
//			},
//		}
//
//		// use mockedGetBalanceDetails in code that requires GetBalanceDetails
//		// and then make assertions.
//
//	}
type GetBalanceDetailsMock struct {
	// GetBalanceDetailsFunc mocks the GetBalanceDetails method.
	GetBalanceDetailsFunc func(ctx context.Context, userID entity.UserID) (*entity.Balance, []*entity.BalanceTrans, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetBalanceDetails holds details about calls to the GetBalanceDetails method.
		GetBalanceDetails []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID entity.UserID
		}
	}
	lockGetBalanceDetails sync.RWMutex
}

// GetBalanceDetails calls GetBalanceDetailsFunc.
func (mock *GetBalanceDetailsMock) GetBalanceDetails(ctx context.Context, userID entity.UserID) (*entity.Balance, []*entity.BalanceTrans, error) {
	if mock.GetBalanceDetailsFunc == nil {
		panic("GetBalanceDetailsMock.GetBalanceDetailsFunc: method is nil but GetBalanceDetails.GetBalanceDetails was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID entity.UserID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockGetBalanceDetails.Lock()
	mock.calls.GetBalanceDetails = append(mock.calls.GetBalanceDetails, callInfo)
	mock.lockGetBalanceDetails.Unlock()
	return mock.GetBalanceDetailsFunc(ctx, userID)
}

// GetBalanceDetailsCalls gets all the calls that were made to GetBalanceDetails.
// Check the length with:
//
//	len(mockedGetBalanceDetails.GetBalanceDetailsCalls())
func (mock *GetBalanceDetailsMock) GetBalanceDetailsCalls() []struct {
	Ctx    context.Context
	UserID entity.UserID
} {
	var calls []struct {
		Ctx    context.Context
		UserID entity.UserID
	}
	mock.lockGetBalanceDetails.RLock()
	calls = mock.calls.GetBalanceDetails
	mock.lockGetBalanceDetails.RUnlock()
	return calls
}
