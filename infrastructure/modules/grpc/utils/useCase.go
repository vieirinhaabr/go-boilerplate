package grpcUtils

import (
	"encoding/json"
	"go-boilerplate/api"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UseCase[req any, res any] func(params *req) *api.UseCaseResponse[res]

func CallUseCaseAsProto[req any, res any, input any, output any](id string, useCase UseCase[req, res], i *input, o *output) (*output, error) {
	utils.Log.Info("[GRPC] %s called\n", id)

	var r = new(req)
	ijs, _ := json.Marshal(i)
	_ = json.Unmarshal(ijs, r)

	result := useCase(r)
	if result.ErrorCode != nil {
		utils.Log.Error("[GRPC] %s got error\n", id)

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

	utils.Log.Success("[GRPC] %s executed\n", id)
	return o, nil
}
