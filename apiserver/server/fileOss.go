package server

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/user"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/config"
)

// func GetBuckets() {
// 	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5t5pmxBdqjvDfMQbSDvP", "tPVIyp7IXasn0L1qKZjhQwg7PS8LqX")
// 	if err != nil {
// 		fmt.Println("New error:", err)
// 	}

// 	lsRes, err := client.ListBuckets()
// 	if err != nil {
// 		fmt.Println("ListBuckets error:", err)
// 	}

// 	for _, bucket := range lsRes.Buckets {
// 		fmt.Println("Buckets:", bucket.Name)
// 	}
// }
func GetSTSToken(params user.GetSTSTokenParams) middleware.Responder {
	//构建一个阿里云客户端, 用于发起请求。
	//设置调用者（RAM用户或RAM角色）的RegionId, AccessKey ID和AccessKey Secret。
	oss, err := sts.NewClientWithAccessKey("cn-hangzhou", config.AccessKeyID, config.AccessKey)
	if err != nil {
		logrus.Errorln("NewClientWithAccessKey error:", err)
	}
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。
	request.RoleArn = "acs:ram::1968596102742607:role/ramosstest"
	request.RoleSessionName = "SessionTest"

	//发起请求，并得到响应。
	response, err := oss.AssumeRole(request)
	if err != nil {
		logrus.Errorln("AssumeRole error:", err)
	}
	logrus.Infoln("GetSTSToken response:", response)
	return user.NewGetSTSTokenOK().WithPayload(&v1.OSSToken{
		AccessKeyID:     response.Credentials.AccessKeyId,
		AccessKeySecret: response.Credentials.AccessKeySecret,
		Expiration:      response.Credentials.Expiration,
		StsToken:        response.Credentials.SecurityToken,
	})
}
