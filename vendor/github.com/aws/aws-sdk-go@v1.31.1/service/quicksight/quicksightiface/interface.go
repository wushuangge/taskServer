// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package quicksightiface provides an interface to enable mocking the Amazon QuickSight service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package quicksightiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/quicksight"
)

// QuickSightAPI provides an interface to enable mocking the
// quicksight.QuickSight service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon QuickSight.
//    func myFunc(svc quicksightiface.QuickSightAPI) bool {
//        // Make svc.CancelIngestion request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := quicksight.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockQuickSightClient struct {
//        quicksightiface.QuickSightAPI
//    }
//    func (m *mockQuickSightClient) CancelIngestion(input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockQuickSightClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type QuickSightAPI interface {
	CancelIngestion(*quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error)
	CancelIngestionWithContext(aws.Context, *quicksight.CancelIngestionInput, ...request.Option) (*quicksight.CancelIngestionOutput, error)
	CancelIngestionRequest(*quicksight.CancelIngestionInput) (*request.Request, *quicksight.CancelIngestionOutput)

	CreateDashboard(*quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error)
	CreateDashboardWithContext(aws.Context, *quicksight.CreateDashboardInput, ...request.Option) (*quicksight.CreateDashboardOutput, error)
	CreateDashboardRequest(*quicksight.CreateDashboardInput) (*request.Request, *quicksight.CreateDashboardOutput)

	CreateDataSet(*quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error)
	CreateDataSetWithContext(aws.Context, *quicksight.CreateDataSetInput, ...request.Option) (*quicksight.CreateDataSetOutput, error)
	CreateDataSetRequest(*quicksight.CreateDataSetInput) (*request.Request, *quicksight.CreateDataSetOutput)

	CreateDataSource(*quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error)
	CreateDataSourceWithContext(aws.Context, *quicksight.CreateDataSourceInput, ...request.Option) (*quicksight.CreateDataSourceOutput, error)
	CreateDataSourceRequest(*quicksight.CreateDataSourceInput) (*request.Request, *quicksight.CreateDataSourceOutput)

	CreateGroup(*quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *quicksight.CreateGroupInput, ...request.Option) (*quicksight.CreateGroupOutput, error)
	CreateGroupRequest(*quicksight.CreateGroupInput) (*request.Request, *quicksight.CreateGroupOutput)

	CreateGroupMembership(*quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error)
	CreateGroupMembershipWithContext(aws.Context, *quicksight.CreateGroupMembershipInput, ...request.Option) (*quicksight.CreateGroupMembershipOutput, error)
	CreateGroupMembershipRequest(*quicksight.CreateGroupMembershipInput) (*request.Request, *quicksight.CreateGroupMembershipOutput)

	CreateIAMPolicyAssignment(*quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error)
	CreateIAMPolicyAssignmentWithContext(aws.Context, *quicksight.CreateIAMPolicyAssignmentInput, ...request.Option) (*quicksight.CreateIAMPolicyAssignmentOutput, error)
	CreateIAMPolicyAssignmentRequest(*quicksight.CreateIAMPolicyAssignmentInput) (*request.Request, *quicksight.CreateIAMPolicyAssignmentOutput)

	CreateIngestion(*quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error)
	CreateIngestionWithContext(aws.Context, *quicksight.CreateIngestionInput, ...request.Option) (*quicksight.CreateIngestionOutput, error)
	CreateIngestionRequest(*quicksight.CreateIngestionInput) (*request.Request, *quicksight.CreateIngestionOutput)

	CreateTemplate(*quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error)
	CreateTemplateWithContext(aws.Context, *quicksight.CreateTemplateInput, ...request.Option) (*quicksight.CreateTemplateOutput, error)
	CreateTemplateRequest(*quicksight.CreateTemplateInput) (*request.Request, *quicksight.CreateTemplateOutput)

	CreateTemplateAlias(*quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error)
	CreateTemplateAliasWithContext(aws.Context, *quicksight.CreateTemplateAliasInput, ...request.Option) (*quicksight.CreateTemplateAliasOutput, error)
	CreateTemplateAliasRequest(*quicksight.CreateTemplateAliasInput) (*request.Request, *quicksight.CreateTemplateAliasOutput)

	DeleteDashboard(*quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error)
	DeleteDashboardWithContext(aws.Context, *quicksight.DeleteDashboardInput, ...request.Option) (*quicksight.DeleteDashboardOutput, error)
	DeleteDashboardRequest(*quicksight.DeleteDashboardInput) (*request.Request, *quicksight.DeleteDashboardOutput)

	DeleteDataSet(*quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error)
	DeleteDataSetWithContext(aws.Context, *quicksight.DeleteDataSetInput, ...request.Option) (*quicksight.DeleteDataSetOutput, error)
	DeleteDataSetRequest(*quicksight.DeleteDataSetInput) (*request.Request, *quicksight.DeleteDataSetOutput)

	DeleteDataSource(*quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error)
	DeleteDataSourceWithContext(aws.Context, *quicksight.DeleteDataSourceInput, ...request.Option) (*quicksight.DeleteDataSourceOutput, error)
	DeleteDataSourceRequest(*quicksight.DeleteDataSourceInput) (*request.Request, *quicksight.DeleteDataSourceOutput)

	DeleteGroup(*quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *quicksight.DeleteGroupInput, ...request.Option) (*quicksight.DeleteGroupOutput, error)
	DeleteGroupRequest(*quicksight.DeleteGroupInput) (*request.Request, *quicksight.DeleteGroupOutput)

	DeleteGroupMembership(*quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error)
	DeleteGroupMembershipWithContext(aws.Context, *quicksight.DeleteGroupMembershipInput, ...request.Option) (*quicksight.DeleteGroupMembershipOutput, error)
	DeleteGroupMembershipRequest(*quicksight.DeleteGroupMembershipInput) (*request.Request, *quicksight.DeleteGroupMembershipOutput)

	DeleteIAMPolicyAssignment(*quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error)
	DeleteIAMPolicyAssignmentWithContext(aws.Context, *quicksight.DeleteIAMPolicyAssignmentInput, ...request.Option) (*quicksight.DeleteIAMPolicyAssignmentOutput, error)
	DeleteIAMPolicyAssignmentRequest(*quicksight.DeleteIAMPolicyAssignmentInput) (*request.Request, *quicksight.DeleteIAMPolicyAssignmentOutput)

	DeleteTemplate(*quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error)
	DeleteTemplateWithContext(aws.Context, *quicksight.DeleteTemplateInput, ...request.Option) (*quicksight.DeleteTemplateOutput, error)
	DeleteTemplateRequest(*quicksight.DeleteTemplateInput) (*request.Request, *quicksight.DeleteTemplateOutput)

	DeleteTemplateAlias(*quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error)
	DeleteTemplateAliasWithContext(aws.Context, *quicksight.DeleteTemplateAliasInput, ...request.Option) (*quicksight.DeleteTemplateAliasOutput, error)
	DeleteTemplateAliasRequest(*quicksight.DeleteTemplateAliasInput) (*request.Request, *quicksight.DeleteTemplateAliasOutput)

	DeleteUser(*quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *quicksight.DeleteUserInput, ...request.Option) (*quicksight.DeleteUserOutput, error)
	DeleteUserRequest(*quicksight.DeleteUserInput) (*request.Request, *quicksight.DeleteUserOutput)

	DeleteUserByPrincipalId(*quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error)
	DeleteUserByPrincipalIdWithContext(aws.Context, *quicksight.DeleteUserByPrincipalIdInput, ...request.Option) (*quicksight.DeleteUserByPrincipalIdOutput, error)
	DeleteUserByPrincipalIdRequest(*quicksight.DeleteUserByPrincipalIdInput) (*request.Request, *quicksight.DeleteUserByPrincipalIdOutput)

	DescribeDashboard(*quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error)
	DescribeDashboardWithContext(aws.Context, *quicksight.DescribeDashboardInput, ...request.Option) (*quicksight.DescribeDashboardOutput, error)
	DescribeDashboardRequest(*quicksight.DescribeDashboardInput) (*request.Request, *quicksight.DescribeDashboardOutput)

	DescribeDashboardPermissions(*quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error)
	DescribeDashboardPermissionsWithContext(aws.Context, *quicksight.DescribeDashboardPermissionsInput, ...request.Option) (*quicksight.DescribeDashboardPermissionsOutput, error)
	DescribeDashboardPermissionsRequest(*quicksight.DescribeDashboardPermissionsInput) (*request.Request, *quicksight.DescribeDashboardPermissionsOutput)

	DescribeDataSet(*quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error)
	DescribeDataSetWithContext(aws.Context, *quicksight.DescribeDataSetInput, ...request.Option) (*quicksight.DescribeDataSetOutput, error)
	DescribeDataSetRequest(*quicksight.DescribeDataSetInput) (*request.Request, *quicksight.DescribeDataSetOutput)

	DescribeDataSetPermissions(*quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error)
	DescribeDataSetPermissionsWithContext(aws.Context, *quicksight.DescribeDataSetPermissionsInput, ...request.Option) (*quicksight.DescribeDataSetPermissionsOutput, error)
	DescribeDataSetPermissionsRequest(*quicksight.DescribeDataSetPermissionsInput) (*request.Request, *quicksight.DescribeDataSetPermissionsOutput)

	DescribeDataSource(*quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error)
	DescribeDataSourceWithContext(aws.Context, *quicksight.DescribeDataSourceInput, ...request.Option) (*quicksight.DescribeDataSourceOutput, error)
	DescribeDataSourceRequest(*quicksight.DescribeDataSourceInput) (*request.Request, *quicksight.DescribeDataSourceOutput)

	DescribeDataSourcePermissions(*quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error)
	DescribeDataSourcePermissionsWithContext(aws.Context, *quicksight.DescribeDataSourcePermissionsInput, ...request.Option) (*quicksight.DescribeDataSourcePermissionsOutput, error)
	DescribeDataSourcePermissionsRequest(*quicksight.DescribeDataSourcePermissionsInput) (*request.Request, *quicksight.DescribeDataSourcePermissionsOutput)

	DescribeGroup(*quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupWithContext(aws.Context, *quicksight.DescribeGroupInput, ...request.Option) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupRequest(*quicksight.DescribeGroupInput) (*request.Request, *quicksight.DescribeGroupOutput)

	DescribeIAMPolicyAssignment(*quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
	DescribeIAMPolicyAssignmentWithContext(aws.Context, *quicksight.DescribeIAMPolicyAssignmentInput, ...request.Option) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
	DescribeIAMPolicyAssignmentRequest(*quicksight.DescribeIAMPolicyAssignmentInput) (*request.Request, *quicksight.DescribeIAMPolicyAssignmentOutput)

	DescribeIngestion(*quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error)
	DescribeIngestionWithContext(aws.Context, *quicksight.DescribeIngestionInput, ...request.Option) (*quicksight.DescribeIngestionOutput, error)
	DescribeIngestionRequest(*quicksight.DescribeIngestionInput) (*request.Request, *quicksight.DescribeIngestionOutput)

	DescribeTemplate(*quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error)
	DescribeTemplateWithContext(aws.Context, *quicksight.DescribeTemplateInput, ...request.Option) (*quicksight.DescribeTemplateOutput, error)
	DescribeTemplateRequest(*quicksight.DescribeTemplateInput) (*request.Request, *quicksight.DescribeTemplateOutput)

	DescribeTemplateAlias(*quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error)
	DescribeTemplateAliasWithContext(aws.Context, *quicksight.DescribeTemplateAliasInput, ...request.Option) (*quicksight.DescribeTemplateAliasOutput, error)
	DescribeTemplateAliasRequest(*quicksight.DescribeTemplateAliasInput) (*request.Request, *quicksight.DescribeTemplateAliasOutput)

	DescribeTemplatePermissions(*quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error)
	DescribeTemplatePermissionsWithContext(aws.Context, *quicksight.DescribeTemplatePermissionsInput, ...request.Option) (*quicksight.DescribeTemplatePermissionsOutput, error)
	DescribeTemplatePermissionsRequest(*quicksight.DescribeTemplatePermissionsInput) (*request.Request, *quicksight.DescribeTemplatePermissionsOutput)

	DescribeUser(*quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *quicksight.DescribeUserInput, ...request.Option) (*quicksight.DescribeUserOutput, error)
	DescribeUserRequest(*quicksight.DescribeUserInput) (*request.Request, *quicksight.DescribeUserOutput)

	GetDashboardEmbedUrl(*quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetDashboardEmbedUrlWithContext(aws.Context, *quicksight.GetDashboardEmbedUrlInput, ...request.Option) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetDashboardEmbedUrlRequest(*quicksight.GetDashboardEmbedUrlInput) (*request.Request, *quicksight.GetDashboardEmbedUrlOutput)

	ListDashboardVersions(*quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error)
	ListDashboardVersionsWithContext(aws.Context, *quicksight.ListDashboardVersionsInput, ...request.Option) (*quicksight.ListDashboardVersionsOutput, error)
	ListDashboardVersionsRequest(*quicksight.ListDashboardVersionsInput) (*request.Request, *quicksight.ListDashboardVersionsOutput)

	ListDashboardVersionsPages(*quicksight.ListDashboardVersionsInput, func(*quicksight.ListDashboardVersionsOutput, bool) bool) error
	ListDashboardVersionsPagesWithContext(aws.Context, *quicksight.ListDashboardVersionsInput, func(*quicksight.ListDashboardVersionsOutput, bool) bool, ...request.Option) error

	ListDashboards(*quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error)
	ListDashboardsWithContext(aws.Context, *quicksight.ListDashboardsInput, ...request.Option) (*quicksight.ListDashboardsOutput, error)
	ListDashboardsRequest(*quicksight.ListDashboardsInput) (*request.Request, *quicksight.ListDashboardsOutput)

	ListDashboardsPages(*quicksight.ListDashboardsInput, func(*quicksight.ListDashboardsOutput, bool) bool) error
	ListDashboardsPagesWithContext(aws.Context, *quicksight.ListDashboardsInput, func(*quicksight.ListDashboardsOutput, bool) bool, ...request.Option) error

	ListDataSets(*quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error)
	ListDataSetsWithContext(aws.Context, *quicksight.ListDataSetsInput, ...request.Option) (*quicksight.ListDataSetsOutput, error)
	ListDataSetsRequest(*quicksight.ListDataSetsInput) (*request.Request, *quicksight.ListDataSetsOutput)

	ListDataSetsPages(*quicksight.ListDataSetsInput, func(*quicksight.ListDataSetsOutput, bool) bool) error
	ListDataSetsPagesWithContext(aws.Context, *quicksight.ListDataSetsInput, func(*quicksight.ListDataSetsOutput, bool) bool, ...request.Option) error

	ListDataSources(*quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error)
	ListDataSourcesWithContext(aws.Context, *quicksight.ListDataSourcesInput, ...request.Option) (*quicksight.ListDataSourcesOutput, error)
	ListDataSourcesRequest(*quicksight.ListDataSourcesInput) (*request.Request, *quicksight.ListDataSourcesOutput)

	ListDataSourcesPages(*quicksight.ListDataSourcesInput, func(*quicksight.ListDataSourcesOutput, bool) bool) error
	ListDataSourcesPagesWithContext(aws.Context, *quicksight.ListDataSourcesInput, func(*quicksight.ListDataSourcesOutput, bool) bool, ...request.Option) error

	ListGroupMemberships(*quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroupMembershipsWithContext(aws.Context, *quicksight.ListGroupMembershipsInput, ...request.Option) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroupMembershipsRequest(*quicksight.ListGroupMembershipsInput) (*request.Request, *quicksight.ListGroupMembershipsOutput)

	ListGroups(*quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *quicksight.ListGroupsInput, ...request.Option) (*quicksight.ListGroupsOutput, error)
	ListGroupsRequest(*quicksight.ListGroupsInput) (*request.Request, *quicksight.ListGroupsOutput)

	ListIAMPolicyAssignments(*quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
	ListIAMPolicyAssignmentsWithContext(aws.Context, *quicksight.ListIAMPolicyAssignmentsInput, ...request.Option) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
	ListIAMPolicyAssignmentsRequest(*quicksight.ListIAMPolicyAssignmentsInput) (*request.Request, *quicksight.ListIAMPolicyAssignmentsOutput)

	ListIAMPolicyAssignmentsForUser(*quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
	ListIAMPolicyAssignmentsForUserWithContext(aws.Context, *quicksight.ListIAMPolicyAssignmentsForUserInput, ...request.Option) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
	ListIAMPolicyAssignmentsForUserRequest(*quicksight.ListIAMPolicyAssignmentsForUserInput) (*request.Request, *quicksight.ListIAMPolicyAssignmentsForUserOutput)

	ListIngestions(*quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error)
	ListIngestionsWithContext(aws.Context, *quicksight.ListIngestionsInput, ...request.Option) (*quicksight.ListIngestionsOutput, error)
	ListIngestionsRequest(*quicksight.ListIngestionsInput) (*request.Request, *quicksight.ListIngestionsOutput)

	ListIngestionsPages(*quicksight.ListIngestionsInput, func(*quicksight.ListIngestionsOutput, bool) bool) error
	ListIngestionsPagesWithContext(aws.Context, *quicksight.ListIngestionsInput, func(*quicksight.ListIngestionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *quicksight.ListTagsForResourceInput, ...request.Option) (*quicksight.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*quicksight.ListTagsForResourceInput) (*request.Request, *quicksight.ListTagsForResourceOutput)

	ListTemplateAliases(*quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error)
	ListTemplateAliasesWithContext(aws.Context, *quicksight.ListTemplateAliasesInput, ...request.Option) (*quicksight.ListTemplateAliasesOutput, error)
	ListTemplateAliasesRequest(*quicksight.ListTemplateAliasesInput) (*request.Request, *quicksight.ListTemplateAliasesOutput)

	ListTemplateAliasesPages(*quicksight.ListTemplateAliasesInput, func(*quicksight.ListTemplateAliasesOutput, bool) bool) error
	ListTemplateAliasesPagesWithContext(aws.Context, *quicksight.ListTemplateAliasesInput, func(*quicksight.ListTemplateAliasesOutput, bool) bool, ...request.Option) error

	ListTemplateVersions(*quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error)
	ListTemplateVersionsWithContext(aws.Context, *quicksight.ListTemplateVersionsInput, ...request.Option) (*quicksight.ListTemplateVersionsOutput, error)
	ListTemplateVersionsRequest(*quicksight.ListTemplateVersionsInput) (*request.Request, *quicksight.ListTemplateVersionsOutput)

	ListTemplateVersionsPages(*quicksight.ListTemplateVersionsInput, func(*quicksight.ListTemplateVersionsOutput, bool) bool) error
	ListTemplateVersionsPagesWithContext(aws.Context, *quicksight.ListTemplateVersionsInput, func(*quicksight.ListTemplateVersionsOutput, bool) bool, ...request.Option) error

	ListTemplates(*quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error)
	ListTemplatesWithContext(aws.Context, *quicksight.ListTemplatesInput, ...request.Option) (*quicksight.ListTemplatesOutput, error)
	ListTemplatesRequest(*quicksight.ListTemplatesInput) (*request.Request, *quicksight.ListTemplatesOutput)

	ListTemplatesPages(*quicksight.ListTemplatesInput, func(*quicksight.ListTemplatesOutput, bool) bool) error
	ListTemplatesPagesWithContext(aws.Context, *quicksight.ListTemplatesInput, func(*quicksight.ListTemplatesOutput, bool) bool, ...request.Option) error

	ListUserGroups(*quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error)
	ListUserGroupsWithContext(aws.Context, *quicksight.ListUserGroupsInput, ...request.Option) (*quicksight.ListUserGroupsOutput, error)
	ListUserGroupsRequest(*quicksight.ListUserGroupsInput) (*request.Request, *quicksight.ListUserGroupsOutput)

	ListUsers(*quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *quicksight.ListUsersInput, ...request.Option) (*quicksight.ListUsersOutput, error)
	ListUsersRequest(*quicksight.ListUsersInput) (*request.Request, *quicksight.ListUsersOutput)

	RegisterUser(*quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error)
	RegisterUserWithContext(aws.Context, *quicksight.RegisterUserInput, ...request.Option) (*quicksight.RegisterUserOutput, error)
	RegisterUserRequest(*quicksight.RegisterUserInput) (*request.Request, *quicksight.RegisterUserOutput)

	SearchDashboards(*quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error)
	SearchDashboardsWithContext(aws.Context, *quicksight.SearchDashboardsInput, ...request.Option) (*quicksight.SearchDashboardsOutput, error)
	SearchDashboardsRequest(*quicksight.SearchDashboardsInput) (*request.Request, *quicksight.SearchDashboardsOutput)

	SearchDashboardsPages(*quicksight.SearchDashboardsInput, func(*quicksight.SearchDashboardsOutput, bool) bool) error
	SearchDashboardsPagesWithContext(aws.Context, *quicksight.SearchDashboardsInput, func(*quicksight.SearchDashboardsOutput, bool) bool, ...request.Option) error

	TagResource(*quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *quicksight.TagResourceInput, ...request.Option) (*quicksight.TagResourceOutput, error)
	TagResourceRequest(*quicksight.TagResourceInput) (*request.Request, *quicksight.TagResourceOutput)

	UntagResource(*quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *quicksight.UntagResourceInput, ...request.Option) (*quicksight.UntagResourceOutput, error)
	UntagResourceRequest(*quicksight.UntagResourceInput) (*request.Request, *quicksight.UntagResourceOutput)

	UpdateDashboard(*quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error)
	UpdateDashboardWithContext(aws.Context, *quicksight.UpdateDashboardInput, ...request.Option) (*quicksight.UpdateDashboardOutput, error)
	UpdateDashboardRequest(*quicksight.UpdateDashboardInput) (*request.Request, *quicksight.UpdateDashboardOutput)

	UpdateDashboardPermissions(*quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error)
	UpdateDashboardPermissionsWithContext(aws.Context, *quicksight.UpdateDashboardPermissionsInput, ...request.Option) (*quicksight.UpdateDashboardPermissionsOutput, error)
	UpdateDashboardPermissionsRequest(*quicksight.UpdateDashboardPermissionsInput) (*request.Request, *quicksight.UpdateDashboardPermissionsOutput)

	UpdateDashboardPublishedVersion(*quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error)
	UpdateDashboardPublishedVersionWithContext(aws.Context, *quicksight.UpdateDashboardPublishedVersionInput, ...request.Option) (*quicksight.UpdateDashboardPublishedVersionOutput, error)
	UpdateDashboardPublishedVersionRequest(*quicksight.UpdateDashboardPublishedVersionInput) (*request.Request, *quicksight.UpdateDashboardPublishedVersionOutput)

	UpdateDataSet(*quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error)
	UpdateDataSetWithContext(aws.Context, *quicksight.UpdateDataSetInput, ...request.Option) (*quicksight.UpdateDataSetOutput, error)
	UpdateDataSetRequest(*quicksight.UpdateDataSetInput) (*request.Request, *quicksight.UpdateDataSetOutput)

	UpdateDataSetPermissions(*quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error)
	UpdateDataSetPermissionsWithContext(aws.Context, *quicksight.UpdateDataSetPermissionsInput, ...request.Option) (*quicksight.UpdateDataSetPermissionsOutput, error)
	UpdateDataSetPermissionsRequest(*quicksight.UpdateDataSetPermissionsInput) (*request.Request, *quicksight.UpdateDataSetPermissionsOutput)

	UpdateDataSource(*quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error)
	UpdateDataSourceWithContext(aws.Context, *quicksight.UpdateDataSourceInput, ...request.Option) (*quicksight.UpdateDataSourceOutput, error)
	UpdateDataSourceRequest(*quicksight.UpdateDataSourceInput) (*request.Request, *quicksight.UpdateDataSourceOutput)

	UpdateDataSourcePermissions(*quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error)
	UpdateDataSourcePermissionsWithContext(aws.Context, *quicksight.UpdateDataSourcePermissionsInput, ...request.Option) (*quicksight.UpdateDataSourcePermissionsOutput, error)
	UpdateDataSourcePermissionsRequest(*quicksight.UpdateDataSourcePermissionsInput) (*request.Request, *quicksight.UpdateDataSourcePermissionsOutput)

	UpdateGroup(*quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *quicksight.UpdateGroupInput, ...request.Option) (*quicksight.UpdateGroupOutput, error)
	UpdateGroupRequest(*quicksight.UpdateGroupInput) (*request.Request, *quicksight.UpdateGroupOutput)

	UpdateIAMPolicyAssignment(*quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error)
	UpdateIAMPolicyAssignmentWithContext(aws.Context, *quicksight.UpdateIAMPolicyAssignmentInput, ...request.Option) (*quicksight.UpdateIAMPolicyAssignmentOutput, error)
	UpdateIAMPolicyAssignmentRequest(*quicksight.UpdateIAMPolicyAssignmentInput) (*request.Request, *quicksight.UpdateIAMPolicyAssignmentOutput)

	UpdateTemplate(*quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error)
	UpdateTemplateWithContext(aws.Context, *quicksight.UpdateTemplateInput, ...request.Option) (*quicksight.UpdateTemplateOutput, error)
	UpdateTemplateRequest(*quicksight.UpdateTemplateInput) (*request.Request, *quicksight.UpdateTemplateOutput)

	UpdateTemplateAlias(*quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error)
	UpdateTemplateAliasWithContext(aws.Context, *quicksight.UpdateTemplateAliasInput, ...request.Option) (*quicksight.UpdateTemplateAliasOutput, error)
	UpdateTemplateAliasRequest(*quicksight.UpdateTemplateAliasInput) (*request.Request, *quicksight.UpdateTemplateAliasOutput)

	UpdateTemplatePermissions(*quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error)
	UpdateTemplatePermissionsWithContext(aws.Context, *quicksight.UpdateTemplatePermissionsInput, ...request.Option) (*quicksight.UpdateTemplatePermissionsOutput, error)
	UpdateTemplatePermissionsRequest(*quicksight.UpdateTemplatePermissionsInput) (*request.Request, *quicksight.UpdateTemplatePermissionsOutput)

	UpdateUser(*quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *quicksight.UpdateUserInput, ...request.Option) (*quicksight.UpdateUserOutput, error)
	UpdateUserRequest(*quicksight.UpdateUserInput) (*request.Request, *quicksight.UpdateUserOutput)
}

var _ QuickSightAPI = (*quicksight.QuickSight)(nil)
