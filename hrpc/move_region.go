// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/tsuna/gohbase/pb"
)

// MoveRegion represents a MoveRegion HBase call
type MoveRegion struct {
	base
}

// NewEnableTable creates a new EnableTable request that will enable the
// given table in HBase. For use by the admin client.a

func NewMoveRegion(ctx context.Context, encodedRegionName []byte) *MoveRegion {
	return &MoveRegion{
		base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of this RPC call.
func (mr *MoveRegion) Name() string {
	return "MoveRegion"
}

// ToProto converts the RPC into a protobuf message
func (mr *MoveRegion) ToProto() proto.Message {
	req := &pb.MoveRegionRequest{}

	return req
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (mr *MoveRegion) NewResponse() proto.Message {
	return &pb.MoveRegionResponse{}
}
