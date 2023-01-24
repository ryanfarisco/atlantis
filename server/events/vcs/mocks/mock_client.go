// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: Client)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClient(options ...pegomock.Option) *MockClient {
	mock := &MockClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockClient) CreateComment(_param0 models.Repo, _param1 int, _param2 string, _param3 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateComment", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) DiscardReviews(_param0 models.Repo, _param1 models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DiscardReviews", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) GetCloneURL(_param0 models.VCSHostType, _param1 string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCloneURL", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) GetFileContent(_param0 models.PullRequest, _param1 string) (bool, []byte, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetFileContent", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 []byte
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].([]byte)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockClient) GetModifiedFiles(_param0 models.Repo, _param1 models.PullRequest) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetModifiedFiles", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) GetTeamNamesForUser(_param0 models.Repo, _param1 models.User) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetTeamNamesForUser", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) HidePrevCommandComments(_param0 models.Repo, _param1 int, _param2 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HidePrevCommandComments", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) MarkdownPullLink(_param0 models.PullRequest) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MarkdownPullLink", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) MergePull(_param0 models.PullRequest, _param1 models.PullRequestOptions) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MergePull", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) PullIsApproved(_param0 models.Repo, _param1 models.PullRequest) (models.ApprovalStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", params, []reflect.Type{reflect.TypeOf((*models.ApprovalStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.ApprovalStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ApprovalStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) PullIsMergeable(_param0 models.Repo, _param1 models.PullRequest, _param2 string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsMergeable", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) SupportsSingleFileDownload(_param0 models.Repo) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsSingleFileDownload", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClient) UpdateStatus(_param0 models.Repo, _param1 models.PullRequest, _param2 models.CommitStatus, _param3 string, _param4 string, _param5 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) VerifyWasCalledOnce() *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClient) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockClient struct {
	mock                   *MockClient
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockClient) CreateComment(_param0 models.Repo, _param1 int, _param2 string, _param3 string) *MockClient_CreateComment_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateComment", params, verifier.timeout)
	return &MockClient_CreateComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_CreateComment_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_CreateComment_OngoingVerification) GetCapturedArguments() (models.Repo, int, string, string) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockClient_CreateComment_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) DiscardReviews(_param0 models.Repo, _param1 models.PullRequest) *MockClient_DiscardReviews_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DiscardReviews", params, verifier.timeout)
	return &MockClient_DiscardReviews_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_DiscardReviews_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_DiscardReviews_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_DiscardReviews_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetCloneURL(_param0 models.VCSHostType, _param1 string) *MockClient_GetCloneURL_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCloneURL", params, verifier.timeout)
	return &MockClient_GetCloneURL_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetCloneURL_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetCloneURL_OngoingVerification) GetCapturedArguments() (models.VCSHostType, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_GetCloneURL_OngoingVerification) GetAllCapturedArguments() (_param0 []models.VCSHostType, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.VCSHostType, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.VCSHostType)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetFileContent(_param0 models.PullRequest, _param1 string) *MockClient_GetFileContent_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetFileContent", params, verifier.timeout)
	return &MockClient_GetFileContent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetFileContent_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetFileContent_OngoingVerification) GetCapturedArguments() (models.PullRequest, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_GetFileContent_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetModifiedFiles(_param0 models.Repo, _param1 models.PullRequest) *MockClient_GetModifiedFiles_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetModifiedFiles", params, verifier.timeout)
	return &MockClient_GetModifiedFiles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetModifiedFiles_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetTeamNamesForUser(_param0 models.Repo, _param1 models.User) *MockClient_GetTeamNamesForUser_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetTeamNamesForUser", params, verifier.timeout)
	return &MockClient_GetTeamNamesForUser_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetTeamNamesForUser_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetTeamNamesForUser_OngoingVerification) GetCapturedArguments() (models.Repo, models.User) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_GetTeamNamesForUser_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.User) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.User, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.User)
		}
	}
	return
}

func (verifier *VerifierMockClient) HidePrevCommandComments(_param0 models.Repo, _param1 int, _param2 string) *MockClient_HidePrevCommandComments_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HidePrevCommandComments", params, verifier.timeout)
	return &MockClient_HidePrevCommandComments_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_HidePrevCommandComments_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_HidePrevCommandComments_OngoingVerification) GetCapturedArguments() (models.Repo, int, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockClient_HidePrevCommandComments_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) MarkdownPullLink(_param0 models.PullRequest) *MockClient_MarkdownPullLink_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MarkdownPullLink", params, verifier.timeout)
	return &MockClient_MarkdownPullLink_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_MarkdownPullLink_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_MarkdownPullLink_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClient_MarkdownPullLink_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) MergePull(_param0 models.PullRequest, _param1 models.PullRequestOptions) *MockClient_MergePull_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MergePull", params, verifier.timeout)
	return &MockClient_MergePull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_MergePull_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_MergePull_OngoingVerification) GetCapturedArguments() (models.PullRequest, models.PullRequestOptions) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_MergePull_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []models.PullRequestOptions) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]models.PullRequestOptions, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequestOptions)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsApproved(_param0 models.Repo, _param1 models.PullRequest) *MockClient_PullIsApproved_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", params, verifier.timeout)
	return &MockClient_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsApproved_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsMergeable(_param0 models.Repo, _param1 models.PullRequest, _param2 string) *MockClient_PullIsMergeable_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsMergeable", params, verifier.timeout)
	return &MockClient_PullIsMergeable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsMergeable_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) SupportsSingleFileDownload(_param0 models.Repo) *MockClient_SupportsSingleFileDownload_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsSingleFileDownload", params, verifier.timeout)
	return &MockClient_SupportsSingleFileDownload_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_SupportsSingleFileDownload_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_SupportsSingleFileDownload_OngoingVerification) GetCapturedArguments() models.Repo {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClient_SupportsSingleFileDownload_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
	}
	return
}

func (verifier *VerifierMockClient) UpdateStatus(_param0 models.Repo, _param1 models.PullRequest, _param2 models.CommitStatus, _param3 string, _param4 string, _param5 string) *MockClient_UpdateStatus_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateStatus", params, verifier.timeout)
	return &MockClient_UpdateStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_UpdateStatus_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, string, string, string) {
	_param0, _param1, _param2, _param3, _param4, _param5 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1], _param5[len(_param5)-1]
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []string, _param4 []string, _param5 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
		_param5 = make([]string, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(string)
		}
	}
	return
}
