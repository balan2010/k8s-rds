// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package alexaforbusinessiface provides an interface to enable mocking the Alexa For Business service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package alexaforbusinessiface

import (
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
)

// AlexaForBusinessAPI provides an interface to enable mocking the
// alexaforbusiness.AlexaForBusiness service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Alexa For Business.
//    func myFunc(svc alexaforbusinessiface.AlexaForBusinessAPI) bool {
//        // Make svc.AssociateContactWithAddressBook request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := alexaforbusiness.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAlexaForBusinessClient struct {
//        alexaforbusinessiface.AlexaForBusinessAPI
//    }
//    func (m *mockAlexaForBusinessClient) AssociateContactWithAddressBook(input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAlexaForBusinessClient{}
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
type AlexaForBusinessAPI interface {
	AssociateContactWithAddressBookRequest(*alexaforbusiness.AssociateContactWithAddressBookInput) alexaforbusiness.AssociateContactWithAddressBookRequest

	AssociateDeviceWithRoomRequest(*alexaforbusiness.AssociateDeviceWithRoomInput) alexaforbusiness.AssociateDeviceWithRoomRequest

	AssociateSkillGroupWithRoomRequest(*alexaforbusiness.AssociateSkillGroupWithRoomInput) alexaforbusiness.AssociateSkillGroupWithRoomRequest

	CreateAddressBookRequest(*alexaforbusiness.CreateAddressBookInput) alexaforbusiness.CreateAddressBookRequest

	CreateContactRequest(*alexaforbusiness.CreateContactInput) alexaforbusiness.CreateContactRequest

	CreateProfileRequest(*alexaforbusiness.CreateProfileInput) alexaforbusiness.CreateProfileRequest

	CreateRoomRequest(*alexaforbusiness.CreateRoomInput) alexaforbusiness.CreateRoomRequest

	CreateSkillGroupRequest(*alexaforbusiness.CreateSkillGroupInput) alexaforbusiness.CreateSkillGroupRequest

	CreateUserRequest(*alexaforbusiness.CreateUserInput) alexaforbusiness.CreateUserRequest

	DeleteAddressBookRequest(*alexaforbusiness.DeleteAddressBookInput) alexaforbusiness.DeleteAddressBookRequest

	DeleteContactRequest(*alexaforbusiness.DeleteContactInput) alexaforbusiness.DeleteContactRequest

	DeleteProfileRequest(*alexaforbusiness.DeleteProfileInput) alexaforbusiness.DeleteProfileRequest

	DeleteRoomRequest(*alexaforbusiness.DeleteRoomInput) alexaforbusiness.DeleteRoomRequest

	DeleteRoomSkillParameterRequest(*alexaforbusiness.DeleteRoomSkillParameterInput) alexaforbusiness.DeleteRoomSkillParameterRequest

	DeleteSkillGroupRequest(*alexaforbusiness.DeleteSkillGroupInput) alexaforbusiness.DeleteSkillGroupRequest

	DeleteUserRequest(*alexaforbusiness.DeleteUserInput) alexaforbusiness.DeleteUserRequest

	DisassociateContactFromAddressBookRequest(*alexaforbusiness.DisassociateContactFromAddressBookInput) alexaforbusiness.DisassociateContactFromAddressBookRequest

	DisassociateDeviceFromRoomRequest(*alexaforbusiness.DisassociateDeviceFromRoomInput) alexaforbusiness.DisassociateDeviceFromRoomRequest

	DisassociateSkillGroupFromRoomRequest(*alexaforbusiness.DisassociateSkillGroupFromRoomInput) alexaforbusiness.DisassociateSkillGroupFromRoomRequest

	GetAddressBookRequest(*alexaforbusiness.GetAddressBookInput) alexaforbusiness.GetAddressBookRequest

	GetContactRequest(*alexaforbusiness.GetContactInput) alexaforbusiness.GetContactRequest

	GetDeviceRequest(*alexaforbusiness.GetDeviceInput) alexaforbusiness.GetDeviceRequest

	GetProfileRequest(*alexaforbusiness.GetProfileInput) alexaforbusiness.GetProfileRequest

	GetRoomRequest(*alexaforbusiness.GetRoomInput) alexaforbusiness.GetRoomRequest

	GetRoomSkillParameterRequest(*alexaforbusiness.GetRoomSkillParameterInput) alexaforbusiness.GetRoomSkillParameterRequest

	GetSkillGroupRequest(*alexaforbusiness.GetSkillGroupInput) alexaforbusiness.GetSkillGroupRequest

	ListDeviceEventsRequest(*alexaforbusiness.ListDeviceEventsInput) alexaforbusiness.ListDeviceEventsRequest

	ListSkillsRequest(*alexaforbusiness.ListSkillsInput) alexaforbusiness.ListSkillsRequest

	ListTagsRequest(*alexaforbusiness.ListTagsInput) alexaforbusiness.ListTagsRequest

	PutRoomSkillParameterRequest(*alexaforbusiness.PutRoomSkillParameterInput) alexaforbusiness.PutRoomSkillParameterRequest

	ResolveRoomRequest(*alexaforbusiness.ResolveRoomInput) alexaforbusiness.ResolveRoomRequest

	RevokeInvitationRequest(*alexaforbusiness.RevokeInvitationInput) alexaforbusiness.RevokeInvitationRequest

	SearchAddressBooksRequest(*alexaforbusiness.SearchAddressBooksInput) alexaforbusiness.SearchAddressBooksRequest

	SearchContactsRequest(*alexaforbusiness.SearchContactsInput) alexaforbusiness.SearchContactsRequest

	SearchDevicesRequest(*alexaforbusiness.SearchDevicesInput) alexaforbusiness.SearchDevicesRequest

	SearchProfilesRequest(*alexaforbusiness.SearchProfilesInput) alexaforbusiness.SearchProfilesRequest

	SearchRoomsRequest(*alexaforbusiness.SearchRoomsInput) alexaforbusiness.SearchRoomsRequest

	SearchSkillGroupsRequest(*alexaforbusiness.SearchSkillGroupsInput) alexaforbusiness.SearchSkillGroupsRequest

	SearchUsersRequest(*alexaforbusiness.SearchUsersInput) alexaforbusiness.SearchUsersRequest

	SendInvitationRequest(*alexaforbusiness.SendInvitationInput) alexaforbusiness.SendInvitationRequest

	StartDeviceSyncRequest(*alexaforbusiness.StartDeviceSyncInput) alexaforbusiness.StartDeviceSyncRequest

	TagResourceRequest(*alexaforbusiness.TagResourceInput) alexaforbusiness.TagResourceRequest

	UntagResourceRequest(*alexaforbusiness.UntagResourceInput) alexaforbusiness.UntagResourceRequest

	UpdateAddressBookRequest(*alexaforbusiness.UpdateAddressBookInput) alexaforbusiness.UpdateAddressBookRequest

	UpdateContactRequest(*alexaforbusiness.UpdateContactInput) alexaforbusiness.UpdateContactRequest

	UpdateDeviceRequest(*alexaforbusiness.UpdateDeviceInput) alexaforbusiness.UpdateDeviceRequest

	UpdateProfileRequest(*alexaforbusiness.UpdateProfileInput) alexaforbusiness.UpdateProfileRequest

	UpdateRoomRequest(*alexaforbusiness.UpdateRoomInput) alexaforbusiness.UpdateRoomRequest

	UpdateSkillGroupRequest(*alexaforbusiness.UpdateSkillGroupInput) alexaforbusiness.UpdateSkillGroupRequest
}

var _ AlexaForBusinessAPI = (*alexaforbusiness.AlexaForBusiness)(nil)
