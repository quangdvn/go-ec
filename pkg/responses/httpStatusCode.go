package responses

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeInvalidParam = 20003 // Invalid email
	ErrCodeInvalidToken = 30001

	// Register code
	ErrCodeUserIsExist = 40001 // User is exist
	// ErrCodeUserIsNotActive                                                                                    = 40002 // User is not active
	// ErrCodeUserIsNotExist                                                                                     = 40003 // User is not exist
	// ErrCodeUserIsNotAdmin                                                                                     = 40004 // User is not admin
	// ErrCodeUserIsNotSuperAdmin                                                                                = 40005 // User is not super admin
	// ErrCodeUserIsNotSuperAdminOrAdmin                                                                         = 40006 // User is not super admin or admin
	// ErrCodeUserIsNotSuperAdminOrAdminOrUser                                                                   = 40007 // User is not super admin or admin or user
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuest                                                            = 40008 // User is not super admin or admin or user or guest
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomer                                                  = 40009 // User is not super admin or admin or user or guest or customer
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplier                                        = 40010 // User is not super admin or admin or user or guest or customer or supplier
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplierOrAgent                                 = 40011 // User is not super admin or admin or user or guest or customer or supplier or agent
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplierOrAgentOrDealer                         = 40012 // User is not super admin or admin or user or guest or customer or supplier or agent or dealer
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplierOrAgentOrDealerOrMember                 = 40013 // User is not super admin or admin or user or guest or customer or supplier or agent or dealer or member
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplierOrAgentOrDealerOrMemberOrPartner        = 40014 // User is not super admin or admin or user or guest or customer or supplier or agent or dealer or member or partner
	// ErrCodeUserIsNotSuperAdminOrAdminOrUserOrGuestOrCustomerOrSupplierOrAgentOrDealerOrMemberOrPartnerOrStaff = 40015 // User is not super admin or admin or user or guest or customer or supplier or agent or dealer or member or partner or staff
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeInvalidParam: "Invalid email",
	ErrCodeInvalidToken: "Invalid token",

	ErrCodeUserIsExist: "User is exist",
}
