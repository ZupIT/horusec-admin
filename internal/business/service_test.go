// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package business

import (
	"testing"

	"github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/ZupIT/horusec-admin/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConfigService_CreateOrUpdate_When_HorusecManagerExistsButHasNoChanges_Expect_NoCall(t *testing.T) {
	// expected behavior
	svc, client := setup()
	singleResult := &v1alpha1.HorusecManagerList{Items: []v1alpha1.HorusecManager{{}}}
	client.On("List", mock.Anything, mock.Anything).Return(singleResult, nil).Once()

	// state under test
	_ = svc.CreateOrUpdate(new(core.Configuration))

	// assertions
	client.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything)
	client.AssertNotCalled(t, "Create", mock.Anything, mock.Anything, mock.Anything)
}

func TestConfigService_CreateOrUpdate_When_HorusecManagerNotExists_Expect_CreateCall(t *testing.T) {
	// expected behavior
	svc, client := setup()
	emptyList := new(v1alpha1.HorusecManagerList)
	client.On("List", mock.Anything, mock.Anything).Return(emptyList, nil).Once()
	client.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

	// state under test
	_ = svc.CreateOrUpdate(new(core.Configuration))

	// assertions
	client.AssertCalled(t, "Create", mock.Anything, mock.AnythingOfType("*v1alpha1.HorusecManager"), mock.Anything)
	client.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything)
}

func TestConfigService_CreateOrUpdate_When_HorusecManagerExists_Expect_UpdateCall(t *testing.T) {
	// expected behavior
	svc, client := setup()
	singleResult := &v1alpha1.HorusecManagerList{Items: []v1alpha1.HorusecManager{{
		Spec: v1alpha1.HorusecManagerSpec{Global: &v1alpha1.Global{EnableAdmin: true}},
	}}}
	client.On("List", mock.Anything, mock.Anything).Return(singleResult, nil).Once()
	client.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

	// state under test
	_ = svc.CreateOrUpdate(new(core.Configuration))

	// assertions
	client.AssertCalled(t, "Update", mock.Anything, mock.AnythingOfType("*v1alpha1.HorusecManager"), mock.Anything)
	client.AssertNotCalled(t, "Create", mock.Anything, mock.Anything, mock.Anything)
}

func TestConfigService_GetConfig_When_SingleResult_Expect_NoError(t *testing.T) {
	// expected behavior
	svc, client := setup()
	singleResult := &v1alpha1.HorusecManagerList{Items: []v1alpha1.HorusecManager{{
		Spec: v1alpha1.HorusecManagerSpec{Global: &v1alpha1.Global{EnableAdmin: true}},
	}}}
	client.On("List", mock.Anything, mock.Anything).Return(singleResult, nil).Once()
	client.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

	// state under test
	_, err := svc.GetConfig()

	// assertions
	assert.NoError(t, err)
}

func TestConfigService_GetConfig_When_MultipleResults_Expect_Error(t *testing.T) {
	// expected behavior
	svc, client := setup()
	multiResult := &v1alpha1.HorusecManagerList{Items: []v1alpha1.HorusecManager{
		{ObjectMeta: v1.ObjectMeta{Name: "created-by-the-operator"}},
		{ObjectMeta: v1.ObjectMeta{Name: "created-manually"}},
	}}
	client.On("List", mock.Anything, mock.Anything).Return(multiResult, nil).Once()
	client.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

	// state under test
	_, err := svc.GetConfig()

	// assertions
	assert.Error(t, err)
}

func setup() (*ConfigService, *mocks.HorusecManagerInterface) {
	client := new(mocks.HorusecManagerInterface)
	return NewConfigService(client), client
}
