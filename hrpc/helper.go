package hrpc

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"

	"github.com/tsuna/gohbase/pb"
)

var (
	NonStartCode      = -1
	UnknownServerName = "#unknown#"
)

func buildRegionSpecifier(Type pb.RegionSpecifier_RegionSpecifierType, value []byte) *pb.RegionSpecifier {
	regionSpecifier := &pb.RegionSpecifier{
		Type:  &Type,
		Value: value,
	}
	return regionSpecifier
}

func buildServerName(destServerName string) (*pb.ServerName, error) {
	tokens := strings.Split(destServerName, ",")
	if len(tokens) != 3 {
		return nil, fmt.Errorf("error format server name %v", destServerName)
	}
	startCode, err := cast.ToUint64E(tokens[2])
	if err != nil {
		startCode = uint64(NonStartCode)
	}
	uport, err := cast.ToUint32E(tokens[1])
	if err != nil {
		return nil, fmt.Errorf("error server port %v", tokens[1])
	}
	serverName := &pb.ServerName{
		HostName:  &tokens[0],
		Port:      &uport,
		StartCode: &startCode,
	}
	return serverName, nil
}
