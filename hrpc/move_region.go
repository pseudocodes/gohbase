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
	encodedRegionName []byte
	destServerName    *pb.ServerName
}

// NewEnableTable creates a new EnableTable request that will enable the
// given table in HBase. For use by the admin client.a

func NewMoveRegion(ctx context.Context, encodedRegionName, destServerName []byte) (*MoveRegion, error) {
	req := &MoveRegion{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		encodedRegionName: encodedRegionName,
		// destServerName:    destServerName,
	}
	if destServerName == nil || len(destServerName) == 0 {
		req.destServerName = nil
	} else {
		serverName, err := buildServerName(string(destServerName))
		if err != nil {
			return nil, err
		}
		req.destServerName = serverName
	}
	return req, nil
}

// Name returns the name of this RPC call.
func (mr *MoveRegion) Name() string {
	return "MoveRegion"
}

// ToProto converts the RPC into a protobuf message
func (mr *MoveRegion) ToProto() proto.Message {
	regionSpecifier := buildRegionSpecifier(pb.RegionSpecifier_ENCODED_REGION_NAME, mr.encodedRegionName)
	msg := &pb.MoveRegionRequest{
		Region:         regionSpecifier,
		DestServerName: mr.destServerName,
	}
	return msg
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (mr *MoveRegion) NewResponse() proto.Message {
	return &pb.MoveRegionResponse{}
}
