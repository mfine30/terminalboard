// This file was generated by counterfeiter
package pipelinesfakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/pipelines"
)

type FakeSyncherDB struct {
	GetAllPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPipelinesMutex       sync.RWMutex
	getAllPipelinesArgsForCall []struct{}
	getAllPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSyncherDB) GetAllPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPipelinesMutex.Lock()
	fake.getAllPipelinesArgsForCall = append(fake.getAllPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPipelines", []interface{}{})
	fake.getAllPipelinesMutex.Unlock()
	if fake.GetAllPipelinesStub != nil {
		return fake.GetAllPipelinesStub()
	} else {
		return fake.getAllPipelinesReturns.result1, fake.getAllPipelinesReturns.result2
	}
}

func (fake *FakeSyncherDB) GetAllPipelinesCallCount() int {
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return len(fake.getAllPipelinesArgsForCall)
}

func (fake *FakeSyncherDB) GetAllPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	fake.getAllPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncherDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSyncherDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pipelines.SyncherDB = new(FakeSyncherDB)