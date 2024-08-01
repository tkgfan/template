// author gmfan
// date 2023/8/1
package service

import (
	"anifun/models"
	"context"
)

func Pong(ctx context.Context) (resp any, err error) {
	return &models.PongVO{
		Pong: "pong",
	}, nil
}
