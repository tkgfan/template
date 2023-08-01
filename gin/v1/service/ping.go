// author gmfan
// date 2023/8/1
package service

import (
	"acsupport/v1/model"
	"context"
)

func Pong(ctx context.Context) (resp any, err error) {
	return &model.PongVO{
		Pong: "pong",
	}, nil
}
