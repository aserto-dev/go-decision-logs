package aerr

import (
	"net/http"

	"github.com/aserto-dev/errors"
	"google.golang.org/grpc/codes"
)

var (
	// Means no tenant id was found in the current context.
	ErrNoTenantID = errors.NewAsertoError("E30001", codes.InvalidArgument, http.StatusBadRequest, "no tenant id specified")
	// Means the tenant id is not valid.
	ErrInvalidTenantID = errors.NewAsertoError("E30002", codes.InvalidArgument, http.StatusBadRequest, "invalid tenant id")
	// The asked-for runtime is not yet available, but will likely be in the future.
	ErrRuntimeLoading = errors.NewAsertoError("E30003", codes.Unavailable, http.StatusTooEarly, "runtime has not yet loaded")
	// Returned if a policy id is not found in the database.
	ErrPolicyNotFound = errors.NewAsertoError("E10004", codes.NotFound, http.StatusNotFound, "policy not found")
	// Returned if a policy id is invalid.
	ErrInvalidPolicyID = errors.NewAsertoError("E30005", codes.InvalidArgument, http.StatusBadRequest, "invalid policy id")
	// Return if a user already exists.
	ErrUserAlreadyExists = errors.NewAsertoError("E30006", codes.AlreadyExists, http.StatusConflict, "user already exists")
	// Returned when authentication has failed or is not possible.
	ErrAuthenticationFailed = errors.NewAsertoError("E30007", codes.FailedPrecondition, http.StatusUnauthorized, "authentication failed")
	// Returned when a given parameter is incorrect (wrong format, value or type).
	ErrInvalidArgument = errors.NewAsertoError("E30008", codes.InvalidArgument, http.StatusBadRequest, "invalid argument")
	// Returned when a runtime query has an error.
	ErrBadQuery = errors.NewAsertoError("E30010", codes.InvalidArgument, http.StatusBadRequest, "invalid query")
	// Returned when a decision is invalid.
	ErrInvalidDecision = errors.NewAsertoError("E30011", codes.InvalidArgument, http.StatusBadRequest, "invalid decision")
	// Returned when a runtime failed to load.
	ErrBadRuntime = errors.NewAsertoError("E30012", codes.Unavailable, http.StatusServiceUnavailable, "runtime loading failed")
	// Returned if object id is not found in the directory.
	ErrDirectoryObjectNotFound = errors.NewAsertoError("E30013", codes.NotFound, http.StatusNotFound, "directory object not found")
	// Returned if the loaded policy is invalid.
	ErrInvalidPolicy = errors.NewAsertoError("E30014", codes.Internal, http.StatusInternalServerError, "invalid policy")
)
