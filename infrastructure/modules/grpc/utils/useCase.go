package grpcUtils

import (
	"encoding/json"
	"go-boilerplate/api"
	"go-boilerplate/infrastructure/global/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UseCase[req any, res any] func(params *req) *api.UseCaseResponse[res]

func CallUseCaseAsProto[req any, res any, input any, output any](useCase UseCase[req, res], i *input, o *output) (*output, error) {
	var r = new(req)
	ijs, _ := json.Marshal(i)
	_ = json.Unmarshal(ijs, r)

	result := useCase(r)
	if result.ErrorCode != nil {
		switch *result.ErrorCode {
		case errors.Internal:
			return nil, status.Errorf(codes.Internal, *result.ErrorMsg)

		case errors.Validation:
			return nil, status.Errorf(codes.InvalidArgument, *result.ErrorMsg)

		case errors.ItemRequestedNotFound:
			return nil, status.Errorf(codes.NotFound, *result.ErrorMsg)

		default:
			return nil, status.Errorf(codes.Internal, "Cannot handle the error")
		}
	}

	ojs, _ := json.Marshal(*result.Response)
	_ = json.Unmarshal(ojs, o)
	return o, nil
}
