package apiclient

import (
	"context"

	"google.golang.org/grpc"

	workflowpkg "github.com/argoproj/argo/pkg/apiclient/workflow"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

type argoKubeWorkflowServiceClient struct {
	delegate workflowpkg.WorkflowServiceServer
}

func (c argoKubeWorkflowServiceClient) CreateWorkflow(ctx context.Context, req *workflowpkg.WorkflowCreateRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.CreateWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) GetWorkflow(ctx context.Context, req *workflowpkg.WorkflowGetRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.GetWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) ListWorkflows(ctx context.Context, req *workflowpkg.WorkflowListRequest, _ ...grpc.CallOption) (*v1alpha1.WorkflowList, error) {
	return c.delegate.ListWorkflows(ctx, req)
}

func (c argoKubeWorkflowServiceClient) WatchWorkflows(ctx context.Context, req *workflowpkg.WatchWorkflowsRequest, _ ...grpc.CallOption) (workflowpkg.WorkflowService_WatchWorkflowsClient, error) {
	panic("not implemented")
}

func (c argoKubeWorkflowServiceClient) DeleteWorkflow(ctx context.Context, req *workflowpkg.WorkflowDeleteRequest, _ ...grpc.CallOption) (*workflowpkg.WorkflowDeleteResponse, error) {
	return c.delegate.DeleteWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) RetryWorkflow(ctx context.Context, req *workflowpkg.WorkflowRetryRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.RetryWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) ResubmitWorkflow(ctx context.Context, req *workflowpkg.WorkflowResubmitRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.ResubmitWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) ResumeWorkflow(ctx context.Context, req *workflowpkg.WorkflowResumeRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.ResumeWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) SuspendWorkflow(ctx context.Context, req *workflowpkg.WorkflowSuspendRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.SuspendWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) TerminateWorkflow(ctx context.Context, req *workflowpkg.WorkflowTerminateRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.TerminateWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) LintWorkflow(ctx context.Context, req *workflowpkg.WorkflowLintRequest, _ ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	return c.delegate.LintWorkflow(ctx, req)
}

func (c argoKubeWorkflowServiceClient) PodLogs(ctx context.Context, req *workflowpkg.WorkflowLogRequest, _ ...grpc.CallOption) (workflowpkg.WorkflowService_PodLogsClient, error) {
	panic("not implemented")
}
