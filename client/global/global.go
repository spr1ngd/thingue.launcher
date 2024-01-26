package global

import (
	"gorm.io/gorm"
	pb "thingue-launcher/common/gen/proto/go/apis/v1"
)

var (
	AppDB      *gorm.DB
	ClientId   uint32
	GrpcClient pb.ServerInstanceServiceClient
)
