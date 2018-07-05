// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53iface provides an interface to enable mocking the Amazon Route 53 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53iface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

// Route53API provides an interface to enable mocking the
// route53.Route53 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53.
//    func myFunc(svc route53iface.Route53API) bool {
//        // Make svc.AssociateVPCWithHostedZone request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := route53.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53Client struct {
//        route53iface.Route53API
//    }
//    func (m *mockRoute53Client) AssociateVPCWithHostedZone(input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53Client{}
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
type Route53API interface {
	AssociateVPCWithHostedZoneRequest(*route53.AssociateVPCWithHostedZoneInput) route53.AssociateVPCWithHostedZoneRequest

	ChangeResourceRecordSetsRequest(*route53.ChangeResourceRecordSetsInput) route53.ChangeResourceRecordSetsRequest

	ChangeTagsForResourceRequest(*route53.ChangeTagsForResourceInput) route53.ChangeTagsForResourceRequest

	CreateHealthCheckRequest(*route53.CreateHealthCheckInput) route53.CreateHealthCheckRequest

	CreateHostedZoneRequest(*route53.CreateHostedZoneInput) route53.CreateHostedZoneRequest

	CreateQueryLoggingConfigRequest(*route53.CreateQueryLoggingConfigInput) route53.CreateQueryLoggingConfigRequest

	CreateReusableDelegationSetRequest(*route53.CreateReusableDelegationSetInput) route53.CreateReusableDelegationSetRequest

	CreateTrafficPolicyRequest(*route53.CreateTrafficPolicyInput) route53.CreateTrafficPolicyRequest

	CreateTrafficPolicyInstanceRequest(*route53.CreateTrafficPolicyInstanceInput) route53.CreateTrafficPolicyInstanceRequest

	CreateTrafficPolicyVersionRequest(*route53.CreateTrafficPolicyVersionInput) route53.CreateTrafficPolicyVersionRequest

	CreateVPCAssociationAuthorizationRequest(*route53.CreateVPCAssociationAuthorizationInput) route53.CreateVPCAssociationAuthorizationRequest

	DeleteHealthCheckRequest(*route53.DeleteHealthCheckInput) route53.DeleteHealthCheckRequest

	DeleteHostedZoneRequest(*route53.DeleteHostedZoneInput) route53.DeleteHostedZoneRequest

	DeleteQueryLoggingConfigRequest(*route53.DeleteQueryLoggingConfigInput) route53.DeleteQueryLoggingConfigRequest

	DeleteReusableDelegationSetRequest(*route53.DeleteReusableDelegationSetInput) route53.DeleteReusableDelegationSetRequest

	DeleteTrafficPolicyRequest(*route53.DeleteTrafficPolicyInput) route53.DeleteTrafficPolicyRequest

	DeleteTrafficPolicyInstanceRequest(*route53.DeleteTrafficPolicyInstanceInput) route53.DeleteTrafficPolicyInstanceRequest

	DeleteVPCAssociationAuthorizationRequest(*route53.DeleteVPCAssociationAuthorizationInput) route53.DeleteVPCAssociationAuthorizationRequest

	DisassociateVPCFromHostedZoneRequest(*route53.DisassociateVPCFromHostedZoneInput) route53.DisassociateVPCFromHostedZoneRequest

	GetAccountLimitRequest(*route53.GetAccountLimitInput) route53.GetAccountLimitRequest

	GetChangeRequest(*route53.GetChangeInput) route53.GetChangeRequest

	GetCheckerIpRangesRequest(*route53.GetCheckerIpRangesInput) route53.GetCheckerIpRangesRequest

	GetGeoLocationRequest(*route53.GetGeoLocationInput) route53.GetGeoLocationRequest

	GetHealthCheckRequest(*route53.GetHealthCheckInput) route53.GetHealthCheckRequest

	GetHealthCheckCountRequest(*route53.GetHealthCheckCountInput) route53.GetHealthCheckCountRequest

	GetHealthCheckLastFailureReasonRequest(*route53.GetHealthCheckLastFailureReasonInput) route53.GetHealthCheckLastFailureReasonRequest

	GetHealthCheckStatusRequest(*route53.GetHealthCheckStatusInput) route53.GetHealthCheckStatusRequest

	GetHostedZoneRequest(*route53.GetHostedZoneInput) route53.GetHostedZoneRequest

	GetHostedZoneCountRequest(*route53.GetHostedZoneCountInput) route53.GetHostedZoneCountRequest

	GetHostedZoneLimitRequest(*route53.GetHostedZoneLimitInput) route53.GetHostedZoneLimitRequest

	GetQueryLoggingConfigRequest(*route53.GetQueryLoggingConfigInput) route53.GetQueryLoggingConfigRequest

	GetReusableDelegationSetRequest(*route53.GetReusableDelegationSetInput) route53.GetReusableDelegationSetRequest

	GetReusableDelegationSetLimitRequest(*route53.GetReusableDelegationSetLimitInput) route53.GetReusableDelegationSetLimitRequest

	GetTrafficPolicyRequest(*route53.GetTrafficPolicyInput) route53.GetTrafficPolicyRequest

	GetTrafficPolicyInstanceRequest(*route53.GetTrafficPolicyInstanceInput) route53.GetTrafficPolicyInstanceRequest

	GetTrafficPolicyInstanceCountRequest(*route53.GetTrafficPolicyInstanceCountInput) route53.GetTrafficPolicyInstanceCountRequest

	ListGeoLocationsRequest(*route53.ListGeoLocationsInput) route53.ListGeoLocationsRequest

	ListHealthChecksRequest(*route53.ListHealthChecksInput) route53.ListHealthChecksRequest

	ListHostedZonesRequest(*route53.ListHostedZonesInput) route53.ListHostedZonesRequest

	ListHostedZonesByNameRequest(*route53.ListHostedZonesByNameInput) route53.ListHostedZonesByNameRequest

	ListQueryLoggingConfigsRequest(*route53.ListQueryLoggingConfigsInput) route53.ListQueryLoggingConfigsRequest

	ListResourceRecordSetsRequest(*route53.ListResourceRecordSetsInput) route53.ListResourceRecordSetsRequest

	ListReusableDelegationSetsRequest(*route53.ListReusableDelegationSetsInput) route53.ListReusableDelegationSetsRequest

	ListTagsForResourceRequest(*route53.ListTagsForResourceInput) route53.ListTagsForResourceRequest

	ListTagsForResourcesRequest(*route53.ListTagsForResourcesInput) route53.ListTagsForResourcesRequest

	ListTrafficPoliciesRequest(*route53.ListTrafficPoliciesInput) route53.ListTrafficPoliciesRequest

	ListTrafficPolicyInstancesRequest(*route53.ListTrafficPolicyInstancesInput) route53.ListTrafficPolicyInstancesRequest

	ListTrafficPolicyInstancesByHostedZoneRequest(*route53.ListTrafficPolicyInstancesByHostedZoneInput) route53.ListTrafficPolicyInstancesByHostedZoneRequest

	ListTrafficPolicyInstancesByPolicyRequest(*route53.ListTrafficPolicyInstancesByPolicyInput) route53.ListTrafficPolicyInstancesByPolicyRequest

	ListTrafficPolicyVersionsRequest(*route53.ListTrafficPolicyVersionsInput) route53.ListTrafficPolicyVersionsRequest

	ListVPCAssociationAuthorizationsRequest(*route53.ListVPCAssociationAuthorizationsInput) route53.ListVPCAssociationAuthorizationsRequest

	TestDNSAnswerRequest(*route53.TestDNSAnswerInput) route53.TestDNSAnswerRequest

	UpdateHealthCheckRequest(*route53.UpdateHealthCheckInput) route53.UpdateHealthCheckRequest

	UpdateHostedZoneCommentRequest(*route53.UpdateHostedZoneCommentInput) route53.UpdateHostedZoneCommentRequest

	UpdateTrafficPolicyCommentRequest(*route53.UpdateTrafficPolicyCommentInput) route53.UpdateTrafficPolicyCommentRequest

	UpdateTrafficPolicyInstanceRequest(*route53.UpdateTrafficPolicyInstanceInput) route53.UpdateTrafficPolicyInstanceRequest

	WaitUntilResourceRecordSetsChanged(*route53.GetChangeInput) error
	WaitUntilResourceRecordSetsChangedWithContext(aws.Context, *route53.GetChangeInput, ...aws.WaiterOption) error
}

var _ Route53API = (*route53.Route53)(nil)
