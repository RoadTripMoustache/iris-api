package enum

import "net/http"

type ErrorKind struct {
	HTTPCode     int
	ErrorMessage string
}

// ----- AUTH ----- //
var (
	// AuthUnauthorized - When a user tries to access a resource without the correct permission.
	AuthUnauthorized = "AUTH001"
	// NotPremium - When a user tries to use a premium feature without being premium.
	NotPremium = "PRE000"
)

// ----- RESOURCE ----- //
var (
	// ResourceNotFound - When a resource is not found.
	ResourceNotFound = "RES000"
	// IDAlreadyUsed - When the id of a new resource is already used.
	IDAlreadyUsed = "RES001"
)

// ----- REQUESTS ----- //
var (
	// BadRequest - When the request is incorrect to process
	BadRequest = "REQ000"
	// UntrustedOrigin - When the request comes from an untrusted origin
	UntrustedOrigin = "REQ001"
	// RequiredFieldMissing - When a required parameter is missing
	RequiredFieldMissing = "REQ002"
)

// ----- SERVER ----- //
var (
	// InternalServerError - When an internal error happened and the API wasn't able to process the request.
	InternalServerError = "SER000"
)

// ----- ADMIN ----- //
var (
	// PermissionsAlreadyGranted - When a permission has already been granted.
	PermissionsAlreadyGranted = "ADM000"
)

// ----- SHOP ----- //
var (
	// NotEnoughCoins - When a user tries to buy something but doesn't enough coins.
	NotEnoughCoins = "SHO000"
)

var ErrorKinds = map[string]ErrorKind{
	// ----- AUTH ----- //
	AuthUnauthorized: {
		HTTPCode:     http.StatusUnauthorized,
		ErrorMessage: "Unauthorized user",
	},
	// > PREMIUM
	NotPremium: {
		HTTPCode:     http.StatusUnauthorized,
		ErrorMessage: "You need premium privileges to do this action",
	},

	// ----- RESOURCE ----- //
	ResourceNotFound: {
		HTTPCode:     http.StatusNotFound,
		ErrorMessage: "Resource not found",
	},
	IDAlreadyUsed: {
		HTTPCode:     http.StatusBadRequest,
		ErrorMessage: "Id already used",
	},

	// ----- REQUESTS ----- //
	BadRequest: {
		HTTPCode:     http.StatusBadRequest,
		ErrorMessage: "There is an issue with your request. Check the logs to have more informations",
	},
	UntrustedOrigin: {
		HTTPCode:     http.StatusUnauthorized,
		ErrorMessage: "Untrusted origin",
	},
	RequiredFieldMissing: {
		HTTPCode:     http.StatusBadRequest,
		ErrorMessage: "A required parameter is missing",
	},

	// ----- SERVER ----- //
	InternalServerError: {
		HTTPCode:     http.StatusInternalServerError,
		ErrorMessage: "An unexpected error happens, please contact support",
	},

	// ----- ADMIN ----- //
	PermissionsAlreadyGranted: {
		HTTPCode:     http.StatusBadRequest,
		ErrorMessage: "Invalid request",
	},

	// ----- SHOP ----- //
	NotEnoughCoins: {
		HTTPCode:     http.StatusBadRequest,
		ErrorMessage: "Not enough coins to do the transaction",
	},
}
