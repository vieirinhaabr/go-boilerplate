package grpcUtils

import (
	"context"
	"encoding/json"
	"go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/infrastructure/global/types"
	grpcTypes "go-boilerplate/infrastructure/modules/grpc/types"
	"go-boilerplate/utils"
	"google.golang.org/grpc"
)

func CallUseCaseAsProto[req any, res any, input any, output any](id string, useCase types.UseCase[req, res], i *input, o *output) (*output, error) {
	utils.Log.Info("[GRPC] %s called\n", id)

	var r = new(req)
	ijs, _ := json.Marshal(i)
	_ = json.Unmarshal(ijs, r)

	result := useCase(r)
	if result.ErrorCode != nil {
		return nil, ResolveServerError(id, result)
	}

	ojs, _ := json.Marshal(*result.Response)
	_ = json.Unmarshal(ojs, o)

	utils.Log.Success("[GRPC] %s executed\n", id)
	return o, nil
}

func CallGrpcServer[rq any, r any, rs any, p any](params *p, call func(ctx context.Context, in *rq, opts ...grpc.CallOption) (*rs, error)) grpcTypes.ClientResponse[r] {
	input := new(rq)
	err := utils.TranslateStruct(params, input)
	if err != nil {
		return grpcTypes.ClientResponse[r]{
			Error: &grpcTypes.ClientError{
				Code: errors.Internal,
				Msg:  "error when mounting request proto",
			},
		}
	}

	ctx, cancel := utils.CreateContext()
	defer cancel()
	resp, err := call(ctx, input)
	if err != nil {
		return ResolveClientError[r](err)
	}

	output := new(r)
	err = utils.TranslateStruct(resp, output)
	if err != nil {
		return grpcTypes.ClientResponse[r]{
			Error: &grpcTypes.ClientError{
				Code: errors.Internal,
				Msg:  "error when mounting request proto",
			},
		}
	}

	return grpcTypes.ClientResponse[r]{
		Response: output,
		Error:    nil,
	}
}
