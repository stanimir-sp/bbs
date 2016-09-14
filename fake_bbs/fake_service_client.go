// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"code.cloudfoundry.org/bbs"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/ifrit"
)

type FakeServiceClient struct {
	CellByIdStub        func(logger lager.Logger, cellId string) (*models.CellPresence, error)
	cellByIdMutex       sync.RWMutex
	cellByIdArgsForCall []struct {
		logger lager.Logger
		cellId string
	}
	cellByIdReturns struct {
		result1 *models.CellPresence
		result2 error
	}
	CellsStub        func(logger lager.Logger) (models.CellSet, error)
	cellsMutex       sync.RWMutex
	cellsArgsForCall []struct {
		logger lager.Logger
	}
	cellsReturns struct {
		result1 models.CellSet
		result2 error
	}
	CellEventsStub        func(logger lager.Logger) <-chan models.CellEvent
	cellEventsMutex       sync.RWMutex
	cellEventsArgsForCall []struct {
		logger lager.Logger
	}
	cellEventsReturns struct {
		result1 <-chan models.CellEvent
	}
	NewCellPresenceRunnerStub        func(logger lager.Logger, cellPresence *models.CellPresence, retryInterval, lockTTL time.Duration) ifrit.Runner
	newCellPresenceRunnerMutex       sync.RWMutex
	newCellPresenceRunnerArgsForCall []struct {
		logger        lager.Logger
		cellPresence  *models.CellPresence
		retryInterval time.Duration
		lockTTL       time.Duration
	}
	newCellPresenceRunnerReturns struct {
		result1 ifrit.Runner
	}
	NewBBSLockRunnerStub        func(logger lager.Logger, bbsPresence *models.BBSPresence, retryInterval, lockTTL time.Duration) (ifrit.Runner, error)
	newBBSLockRunnerMutex       sync.RWMutex
	newBBSLockRunnerArgsForCall []struct {
		logger        lager.Logger
		bbsPresence   *models.BBSPresence
		retryInterval time.Duration
		lockTTL       time.Duration
	}
	newBBSLockRunnerReturns struct {
		result1 ifrit.Runner
		result2 error
	}
	CurrentBBSStub        func(logger lager.Logger) (*models.BBSPresence, error)
	currentBBSMutex       sync.RWMutex
	currentBBSArgsForCall []struct {
		logger lager.Logger
	}
	currentBBSReturns struct {
		result1 *models.BBSPresence
		result2 error
	}
	CurrentBBSURLStub        func(logger lager.Logger) (string, error)
	currentBBSURLMutex       sync.RWMutex
	currentBBSURLArgsForCall []struct {
		logger lager.Logger
	}
	currentBBSURLReturns struct {
		result1 string
		result2 error
	}
}

func (fake *FakeServiceClient) CellById(logger lager.Logger, cellId string) (*models.CellPresence, error) {
	fake.cellByIdMutex.Lock()
	fake.cellByIdArgsForCall = append(fake.cellByIdArgsForCall, struct {
		logger lager.Logger
		cellId string
	}{logger, cellId})
	fake.cellByIdMutex.Unlock()
	if fake.CellByIdStub != nil {
		return fake.CellByIdStub(logger, cellId)
	} else {
		return fake.cellByIdReturns.result1, fake.cellByIdReturns.result2
	}
}

func (fake *FakeServiceClient) CellByIdCallCount() int {
	fake.cellByIdMutex.RLock()
	defer fake.cellByIdMutex.RUnlock()
	return len(fake.cellByIdArgsForCall)
}

func (fake *FakeServiceClient) CellByIdArgsForCall(i int) (lager.Logger, string) {
	fake.cellByIdMutex.RLock()
	defer fake.cellByIdMutex.RUnlock()
	return fake.cellByIdArgsForCall[i].logger, fake.cellByIdArgsForCall[i].cellId
}

func (fake *FakeServiceClient) CellByIdReturns(result1 *models.CellPresence, result2 error) {
	fake.CellByIdStub = nil
	fake.cellByIdReturns = struct {
		result1 *models.CellPresence
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceClient) Cells(logger lager.Logger) (models.CellSet, error) {
	fake.cellsMutex.Lock()
	fake.cellsArgsForCall = append(fake.cellsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.cellsMutex.Unlock()
	if fake.CellsStub != nil {
		return fake.CellsStub(logger)
	} else {
		return fake.cellsReturns.result1, fake.cellsReturns.result2
	}
}

func (fake *FakeServiceClient) CellsCallCount() int {
	fake.cellsMutex.RLock()
	defer fake.cellsMutex.RUnlock()
	return len(fake.cellsArgsForCall)
}

func (fake *FakeServiceClient) CellsArgsForCall(i int) lager.Logger {
	fake.cellsMutex.RLock()
	defer fake.cellsMutex.RUnlock()
	return fake.cellsArgsForCall[i].logger
}

func (fake *FakeServiceClient) CellsReturns(result1 models.CellSet, result2 error) {
	fake.CellsStub = nil
	fake.cellsReturns = struct {
		result1 models.CellSet
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceClient) CellEvents(logger lager.Logger) <-chan models.CellEvent {
	fake.cellEventsMutex.Lock()
	fake.cellEventsArgsForCall = append(fake.cellEventsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.cellEventsMutex.Unlock()
	if fake.CellEventsStub != nil {
		return fake.CellEventsStub(logger)
	} else {
		return fake.cellEventsReturns.result1
	}
}

func (fake *FakeServiceClient) CellEventsCallCount() int {
	fake.cellEventsMutex.RLock()
	defer fake.cellEventsMutex.RUnlock()
	return len(fake.cellEventsArgsForCall)
}

func (fake *FakeServiceClient) CellEventsArgsForCall(i int) lager.Logger {
	fake.cellEventsMutex.RLock()
	defer fake.cellEventsMutex.RUnlock()
	return fake.cellEventsArgsForCall[i].logger
}

func (fake *FakeServiceClient) CellEventsReturns(result1 <-chan models.CellEvent) {
	fake.CellEventsStub = nil
	fake.cellEventsReturns = struct {
		result1 <-chan models.CellEvent
	}{result1}
}

func (fake *FakeServiceClient) NewCellPresenceRunner(logger lager.Logger, cellPresence *models.CellPresence, retryInterval time.Duration, lockTTL time.Duration) ifrit.Runner {
	fake.newCellPresenceRunnerMutex.Lock()
	fake.newCellPresenceRunnerArgsForCall = append(fake.newCellPresenceRunnerArgsForCall, struct {
		logger        lager.Logger
		cellPresence  *models.CellPresence
		retryInterval time.Duration
		lockTTL       time.Duration
	}{logger, cellPresence, retryInterval, lockTTL})
	fake.newCellPresenceRunnerMutex.Unlock()
	if fake.NewCellPresenceRunnerStub != nil {
		return fake.NewCellPresenceRunnerStub(logger, cellPresence, retryInterval, lockTTL)
	} else {
		return fake.newCellPresenceRunnerReturns.result1
	}
}

func (fake *FakeServiceClient) NewCellPresenceRunnerCallCount() int {
	fake.newCellPresenceRunnerMutex.RLock()
	defer fake.newCellPresenceRunnerMutex.RUnlock()
	return len(fake.newCellPresenceRunnerArgsForCall)
}

func (fake *FakeServiceClient) NewCellPresenceRunnerArgsForCall(i int) (lager.Logger, *models.CellPresence, time.Duration, time.Duration) {
	fake.newCellPresenceRunnerMutex.RLock()
	defer fake.newCellPresenceRunnerMutex.RUnlock()
	return fake.newCellPresenceRunnerArgsForCall[i].logger, fake.newCellPresenceRunnerArgsForCall[i].cellPresence, fake.newCellPresenceRunnerArgsForCall[i].retryInterval, fake.newCellPresenceRunnerArgsForCall[i].lockTTL
}

func (fake *FakeServiceClient) NewCellPresenceRunnerReturns(result1 ifrit.Runner) {
	fake.NewCellPresenceRunnerStub = nil
	fake.newCellPresenceRunnerReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeServiceClient) NewBBSLockRunner(logger lager.Logger, bbsPresence *models.BBSPresence, retryInterval time.Duration, lockTTL time.Duration) (ifrit.Runner, error) {
	fake.newBBSLockRunnerMutex.Lock()
	fake.newBBSLockRunnerArgsForCall = append(fake.newBBSLockRunnerArgsForCall, struct {
		logger        lager.Logger
		bbsPresence   *models.BBSPresence
		retryInterval time.Duration
		lockTTL       time.Duration
	}{logger, bbsPresence, retryInterval, lockTTL})
	fake.newBBSLockRunnerMutex.Unlock()
	if fake.NewBBSLockRunnerStub != nil {
		return fake.NewBBSLockRunnerStub(logger, bbsPresence, retryInterval, lockTTL)
	} else {
		return fake.newBBSLockRunnerReturns.result1, fake.newBBSLockRunnerReturns.result2
	}
}

func (fake *FakeServiceClient) NewBBSLockRunnerCallCount() int {
	fake.newBBSLockRunnerMutex.RLock()
	defer fake.newBBSLockRunnerMutex.RUnlock()
	return len(fake.newBBSLockRunnerArgsForCall)
}

func (fake *FakeServiceClient) NewBBSLockRunnerArgsForCall(i int) (lager.Logger, *models.BBSPresence, time.Duration, time.Duration) {
	fake.newBBSLockRunnerMutex.RLock()
	defer fake.newBBSLockRunnerMutex.RUnlock()
	return fake.newBBSLockRunnerArgsForCall[i].logger, fake.newBBSLockRunnerArgsForCall[i].bbsPresence, fake.newBBSLockRunnerArgsForCall[i].retryInterval, fake.newBBSLockRunnerArgsForCall[i].lockTTL
}

func (fake *FakeServiceClient) NewBBSLockRunnerReturns(result1 ifrit.Runner, result2 error) {
	fake.NewBBSLockRunnerStub = nil
	fake.newBBSLockRunnerReturns = struct {
		result1 ifrit.Runner
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceClient) CurrentBBS(logger lager.Logger) (*models.BBSPresence, error) {
	fake.currentBBSMutex.Lock()
	fake.currentBBSArgsForCall = append(fake.currentBBSArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.currentBBSMutex.Unlock()
	if fake.CurrentBBSStub != nil {
		return fake.CurrentBBSStub(logger)
	} else {
		return fake.currentBBSReturns.result1, fake.currentBBSReturns.result2
	}
}

func (fake *FakeServiceClient) CurrentBBSCallCount() int {
	fake.currentBBSMutex.RLock()
	defer fake.currentBBSMutex.RUnlock()
	return len(fake.currentBBSArgsForCall)
}

func (fake *FakeServiceClient) CurrentBBSArgsForCall(i int) lager.Logger {
	fake.currentBBSMutex.RLock()
	defer fake.currentBBSMutex.RUnlock()
	return fake.currentBBSArgsForCall[i].logger
}

func (fake *FakeServiceClient) CurrentBBSReturns(result1 *models.BBSPresence, result2 error) {
	fake.CurrentBBSStub = nil
	fake.currentBBSReturns = struct {
		result1 *models.BBSPresence
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceClient) CurrentBBSURL(logger lager.Logger) (string, error) {
	fake.currentBBSURLMutex.Lock()
	fake.currentBBSURLArgsForCall = append(fake.currentBBSURLArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.currentBBSURLMutex.Unlock()
	if fake.CurrentBBSURLStub != nil {
		return fake.CurrentBBSURLStub(logger)
	} else {
		return fake.currentBBSURLReturns.result1, fake.currentBBSURLReturns.result2
	}
}

func (fake *FakeServiceClient) CurrentBBSURLCallCount() int {
	fake.currentBBSURLMutex.RLock()
	defer fake.currentBBSURLMutex.RUnlock()
	return len(fake.currentBBSURLArgsForCall)
}

func (fake *FakeServiceClient) CurrentBBSURLArgsForCall(i int) lager.Logger {
	fake.currentBBSURLMutex.RLock()
	defer fake.currentBBSURLMutex.RUnlock()
	return fake.currentBBSURLArgsForCall[i].logger
}

func (fake *FakeServiceClient) CurrentBBSURLReturns(result1 string, result2 error) {
	fake.CurrentBBSURLStub = nil
	fake.currentBBSURLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ bbs.ServiceClient = new(FakeServiceClient)
