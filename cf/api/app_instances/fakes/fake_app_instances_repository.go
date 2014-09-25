// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/app_instances"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeAppInstancesRepository struct {
	GetInstancesStub        func(appGuid string) (instances []models.AppInstanceFields, apiErr error)
	getInstancesMutex       sync.RWMutex
	getInstancesArgsForCall []struct {
		appGuid string
	}
	getInstancesReturns struct {
		result1 []models.AppInstanceFields
		result2 error
	}
}

func (fake *FakeAppInstancesRepository) GetInstances(appGuid string) (instances []models.AppInstanceFields, apiErr error) {
	fake.getInstancesMutex.Lock()
	fake.getInstancesArgsForCall = append(fake.getInstancesArgsForCall, struct {
		appGuid string
	}{appGuid})
	fake.getInstancesMutex.Unlock()
	if fake.GetInstancesStub != nil {
		return fake.GetInstancesStub(appGuid)
	} else {
		return fake.getInstancesReturns.result1, fake.getInstancesReturns.result2
	}
}

func (fake *FakeAppInstancesRepository) GetInstancesCallCount() int {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return len(fake.getInstancesArgsForCall)
}

func (fake *FakeAppInstancesRepository) GetInstancesArgsForCall(i int) string {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return fake.getInstancesArgsForCall[i].appGuid
}

func (fake *FakeAppInstancesRepository) GetInstancesReturns(result1 []models.AppInstanceFields, result2 error) {
	fake.GetInstancesStub = nil
	fake.getInstancesReturns = struct {
		result1 []models.AppInstanceFields
		result2 error
	}{result1, result2}
}

var _ app_instances.AppInstancesRepository = new(FakeAppInstancesRepository)
