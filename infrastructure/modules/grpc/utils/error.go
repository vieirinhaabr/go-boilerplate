package grpcUtils

import (
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	grpcTypes "go-boilerplate/infrastructure/modules/grpc/types"
	"go-boilerplate/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ResolveServerError[res any](id string, result *types.UseCaseResponse[res]) error {
	utils.Log.Error("[GRPC] %s got error\n", id)

	switch *result.ErrorCode {
	case errors.Internal:
		return status.Errorf(codes.Internal, *result.ErrorMsg)

	case errors.Validation:
		return status.Errorf(codes.InvalidArgument, *result.ErrorMsg)

	case errors.ItemRequestedNotFound:
		return status.Errorf(codes.NotFound, *result.ErrorMsg)

	default:
		return status.Errorf(codes.Internal, "Cannot handle the error")
	}
}

func ResolveClientError[res any](err error) grpcTypes.ClientResponse[res] {
	st, _ := status.FromError(err)
	msg := st.Message()

	switch st.Code() {
	case codes.Internal:
		return grpcTypes.ClientResponse[res]{
			Error: &grpcTypes.ClientError{
				Code: errors.Internal,
				Msg:  msg,
			},
		}

	case codes.InvalidArgument:
		return grpcTypes.ClientResponse[res]{
			Error: &grpcTypes.ClientError{
				Code: errors.Validation,
				Msg:  msg,
			},
		}

	case codes.NotFound:
		return grpcTypes.ClientResponse[res]{
			Error: &grpcTypes.ClientError{
				Code: errors.ItemRequestedNotFound,
				Msg:  msg,
			},
		}

	default:
		return grpcTypes.ClientResponse[res]{
			Error: &grpcTypes.ClientError{
				Code: errors.Internal,
				Msg:  "got error not defined during request",
			},
		}
	}
}
