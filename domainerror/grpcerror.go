package domainerror

import "google.golang.org/grpc/codes"

// gRPCStatus converts an error class to a gRPC status code.
func gRPCStatus(class ErrorClass) codes.Code {
	switch class {
	case Invalid:
		return codes.InvalidArgument
	case NotFound:
		return codes.NotFound
	case PermissionDenied:
		return codes.PermissionDenied
	case Conflict:
		return codes.AlreadyExists
	case Unauthenticated:
		return codes.Unauthenticated
	case Unimplemented:
		return codes.Unimplemented
	case Unavailable:
		return codes.Unavailable
	case Internal:
		return codes.Internal
	default:
		return codes.Unknown
	}
}
