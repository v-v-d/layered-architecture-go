package application_test

import (
	"errors"
	"layered-arch/internal/application"
	"layered-arch/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseTrxManager_Run_Success(t *testing.T) {
	mockTrx := mocks.NewTrxManager(t)
	trxManager := application.TrxManagerBase{TrxManager: mockTrx}

	mockTrx.On("Commit").Return(nil).Once()
	mockTrx.On("Shutdown").Return(nil).Once()

	err := trxManager.Run(func() error {
		return nil
	})

	assert.NoError(t, err, "Transaction should complete successfully")
}

func TestBaseTrxManager_Run_Failure(t *testing.T) {
	mockTrx := mocks.NewTrxManager(t)
	trxManager := application.TrxManagerBase{TrxManager: mockTrx}

	mockTrx.On("Rollback").Return(nil).Once()
	mockTrx.On("Shutdown").Return(nil).Once()

	err := trxManager.Run(func() error {
		return errors.New("database error")
	})

	assert.Error(t, err, "Expected a transaction failure error")
}

func TestBaseTrxManager_Run_CommitError(t *testing.T) {
	mockTrx := mocks.NewTrxManager(t)
	trxManager := application.TrxManagerBase{TrxManager: mockTrx}

	mockTrx.On("Commit").Return(errors.New("commit failed")).Once()
	mockTrx.On("Shutdown").Return(nil).Once()

	err := trxManager.Run(func() error {
		return nil
	})

	assert.Error(t, err, "Expected a commit error")
	assert.IsType(t, &application.TrxManagerError{}, err, "Expected error type TrxManagerError")
}

func TestBaseTrxManager_Run_RollbackError(t *testing.T) {
	mockTrx := mocks.NewTrxManager(t)
	trxManager := application.TrxManagerBase{TrxManager: mockTrx}

	mockTrx.On("Rollback").Return(errors.New("rollback failed")).Once()
	mockTrx.On("Shutdown").Return(nil).Once()

	err := trxManager.Run(func() error {
		return errors.New("database error")
	})

	assert.Error(t, err, "Expected a rollback error")
	assert.IsType(t, &application.TrxManagerError{}, err, "Expected error type TrxManagerError")
}

func TestBaseTrxManager_Run_ShutdownError(t *testing.T) {
	mockTrx := mocks.NewTrxManager(t)
	trxManager := application.TrxManagerBase{TrxManager: mockTrx}

	mockTrx.On("Commit").Return(nil).Once()
	mockTrx.On("Shutdown").Return(errors.New("shutdown error")).Once()

	err := trxManager.Run(func() error {
		return nil
	})

	assert.NoError(t, err, "Shutdown error should not break the transaction")
}
