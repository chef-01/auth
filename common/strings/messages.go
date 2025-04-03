// File: commons/strings/messages.go
package strings

// ----------------------------
// Generic Success Messages
// ----------------------------
const (
	// Used generically to indicate that a create operation succeeded.
	CreateSuccessMsg = "created successfully"
	// Used generically to indicate that an update operation succeeded.
	UpdateSuccessMsg = "updated successfully"
	// Used generically to indicate that a delete operation succeeded.
	DeleteSuccessMsg = "deleted successfully"
	// Used generically to indicate that a retrieve operation succeeded.
	RetrieveSuccessMsg = "retrieved successfully"
	// Used when a generic operation has completed without issues.
	OperationSuccessMsg = "operation completed successfully"
)

// ----------------------------
// Generic Error Messages
// ----------------------------
const (
	// When validation fails.
	RequestValidationFailedMsg = "request validation failed"
	// When data provided is not valid.
	InvalidDataMsg = "invalid data provided"
	// When a resource cannot be found.
	ResourceNotFoundMsg = "resource not found"
	// When an unexpected error occurs.
	InternalServerErrorMsg = "internal server error"
)

// ----------------------------
// Authentication & Authorization Errors
// ----------------------------
const (
	// Missing the Authorization header.
	MissingAuthHeaderMsg = "missing authorization header"
	// The token does not have the correct format.
	InvalidTokenFormatMsg = "invalid token format"
	// The token is either invalid or has expired.
	InvalidOrExpiredTokenMsg = "invalid or expired token"
	// The user does not have the required permission.
	InsufficientPermissionsMsg = "insufficient permissions"
	// When token claims are missing or invalid.
	MissingOrInvalidTokenClaimsMsg = "missing or invalid token claims"
)

// ----------------------------
// JSON Parsing & Field Validation Errors
// ----------------------------
const (
	// When the JSON in the request cannot be parsed.
	InvalidJSONFormatMsg = "invalid json format"
	// Generic invalid ID message. Append with specific context (e.g., "Ticket", "Community", etc.) as needed.
	InvalidIDMsg = "invalid id"
	InvalidValueMsg = "value invalid"
)

// ----------------------------
// Data Authorization Errors
// ----------------------------
const (
	// When the resource expected in the request context is missing.
	ResourceNotFoundInContextMsg = "resource not found in context"
	// When extracting required fields from the resource fails.
	ResourceExtractionFailedMsg = "resource extraction failed"
	// When the authenticated user's access level does not allow access to the resource.
	AccessDeniedMsg = "access denied"
)

// ----------------------------
// Specific Creation Errors
// ----------------------------
const (
	// For creation operations: when a user tries to create a resource for another community.
	CannotCreateForOtherCommunityMsg = "cannot create for a community other than your own"
	// For creation operations: when a user tries to create a resource on behalf of another user.
	CannotCreateForAnotherUserMsg = "cannot create on behalf of another user"
)

// ----------------------------
// Database & Retrieval Errors
// ----------------------------
const (
	// When there is an error retrieving a resource from the database.
	ErrorRetrievingResourceMsg = "error retrieving resource"
)

// ----------------------------
// Health Check Messages
// ----------------------------
const (
	// Generic health check title/message.
	HealthCheckMsg = "health Check"
	// Used to indicate that a component is running.
	APIRunningMsg = "running"
)