package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/user"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

// 小程序id 小程序密钥
var MpAppid = "wx812ac13612c65659"
var MpSerect = "0e8120143ce980b6195f88d7df9acf8c"
var AccessToken = ""
var AccessTokenExpires int64 = 0

func CreateUser(params user.CreateUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().CreateUser(ctx, &pb.CreateUserRequest{
		User: convert.UserV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return user.NewCreateUserOK().WithPayload(convert.UserPb2V1(reply))
}

func DeleteUser(params user.DeleteUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewDeleteUserOK().WithPayload(convert.UserPb2V1(reply))
}

func GetUser(params user.GetUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().GetUser(ctx, &pb.GetUserRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewGetUserOK().WithPayload(convert.UserPb2V1(reply))
}

func Login(params user.LoginParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().Login(ctx, &pb.LoginRequest{
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		return Error(err)
	}

	accessToken := reply.Token["access_token"]
	tokenType := reply.Token["token_type"]
	expiresIn := reply.Token["expires_in"]
	expiresAt := reply.Token["expires_at"]

	payload := &v1.Token{
		AccessToken: accessToken,
		TokenType:   &tokenType,
		ExpiresIn:   expiresIn,
		ExpiresAt:   expiresAt,
	}
	return user.NewLoginOK().WithPayload(payload)
}

func Logout(params user.LogoutParams) middleware.Responder {
	return user.NewLogoutOK()
}

func UpdateUser(params user.UpdateUserParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.User().UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:   params.ID,
		User: convert.UserV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return user.NewUpdateUserOK().WithPayload(convert.UserPb2V1(reply))
}

func ModifyUserPassword(params user.ModifyUserPasswordParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	_, err := client.User().ModifyUserPassword(ctx, &pb.ModifyUserPasswordRequest{
		Username: *params.Username,
		NewPwd:   params.NewPassword,
		OldPwd:   params.OldPassword,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewModifyUserPasswordOK()
}

// ChangeUserPassword patch /user
func ChangeUserPassword(params user.ChangeUserPasswordParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	openid := ""
	if params.Openid != nil {
		openid = *params.Openid
	}
	_, err := client.User().ChangePassword(ctx, &pb.ChangePasswordRequest{
		Username: params.Username,
		Password: params.Password,
		Openid:   openid,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewChangeUserPasswordOK()
}

// ListUser
func GetUsers(params user.GetUsersParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	limit := int64(0)
	if params.Limit != nil {
		limit = *params.Limit
	}
	skip := int64(0)
	if params.Skip != nil {
		skip = *params.Skip
	}
	query := `{}`
	if params.Query != nil {
		query = *params.Query
	}
	reply, err := client.User().GetUsers(ctx, &pb.GetUsersRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.User{}
	for _, v := range reply.Items {
		items = append(items, convert.UserPb2V1(v))
	}

	payload := &v1.Users{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return user.NewGetUsersOK().WithPayload(payload)
}

func GetUserInfo(params user.GetUserInfoParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	headerContent := params.HTTPRequest.Header.Get("authorization")
	token := strings.Split(headerContent, " ")[1]

	authenticate, err := client.User().Authenticate(ctx, &pb.AuthenticateRequest{
		Token: token,
	})
	if err != nil {
		logrus.Errorln("Authenticate error:", err)
		return Error(err)
	}

	logrus.Info("authenticate:", authenticate.Id)
	reply, err := client.User().GetUser(ctx, &pb.GetUserRequest{
		Id: authenticate.Id,
	})
	if err != nil {
		return Error(err)
	}
	return user.NewGetUserOK().WithPayload(convert.UserPb2V1(reply))
}

// 小程序登录 获取openid  用户唯一标识 先用openid 代替用户名和密码 生成token
func WXLogin(params user.WXLoginParams) middleware.Responder {
	code := params.Code
	logrus.Info("WXLogin code:", code)

	//向微信服务器发送获取用户openid的get请求
	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/sns/jscode2session?appid=",
		MpAppid,
		"&secret=",
		MpSerect,
		"&js_code=",
		code,
		"&grant_type=authorization_code",
	}, "")

	logrus.Info(requestLine)

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error(err)
	}
	logrus.Info("body", string(body))
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Errorln("解析微信数据失败", err)
		return Error(err)
	}
	ctx := params.HTTPRequest.Context()
	logrus.Info("openid", data["openid"])
	// get user
	hasuser, err := client.User().GetUser(ctx, &pb.GetUserRequest{
		Id: data["openid"],
	})
	if err != nil {
		logrus.Errorln("查找是否注册过失败：", err)
		// return Error(err)
	}
	logrus.Info("hasuser", hasuser)
	if hasuser == nil {
		// creater user
		hasuser, err = client.User().CreateUser(ctx, &pb.CreateUserRequest{
			User: &pb.User{
				Id:      data["openid"],
				Openid:  data["openid"],
				Unionid: data["openid"],
				Name:    data["openid"],
				Ps:      data["openid"],
				Role:    "",
				Phone:   "",
				Email:   "",
			},
		})
		if err != nil {
			logrus.Errorln("初始化小程序用户失败：", err)
			return Error(err)
		}
	}
	reply, err := client.User().Login(ctx, &pb.LoginRequest{
		Username: hasuser.Name,
		Password: hasuser.Ps,
	})
	if err != nil {
		logrus.Errorln("生成小程序用户 token失败：", err)
		return Error(err)
	}

	accessToken := reply.Token["access_token"]
	tokenType := reply.Token["token_type"]
	expiresIn := reply.Token["expires_in"]
	expiresAt := reply.Token["expires_at"]

	payload := &v1.Token{
		AccessToken: accessToken,
		TokenType:   &tokenType,
		ExpiresIn:   expiresIn,
		ExpiresAt:   expiresAt,
	}
	return user.NewLoginOK().WithPayload(payload)

}

// 获取 access_token，需要自己进行妥善保存
func GetAccessToken() (string, middleware.Responder) {
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	requestLine := strings.Join([]string{"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=",
		MpAppid,
		"&secret=",
		MpSerect,
	}, "")

	logrus.Info(requestLine)

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", Error(err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	AccessToken = data["access_token"].(string)
	AccessTokenExpires = time.Now().Unix() + data["expires_in"].(int64)

	return AccessToken, nil
}

// 获取用户手机号，作为登录账号
func GetPhoneNumber(params user.GetPhoneNumberParams, principal *v1.Principal) middleware.Responder {
	code := ""
	logrus.Info("GetPhoneNumber code:", code)
	if params.Code == "" {
		return Error(errors.New("code 不能为空"))
	}
	if AccessTokenExpires-time.Now().Unix() > 60 {
		GetAccessToken()
	}

	req := make(map[string]string)
	req["code"] = params.Code
	reqStr, err := json.Marshal(req)
	if err != nil {
		return Error(err)
	}
	// POST https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN
	requestLine := strings.Join([]string{"POST https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=",
		AccessToken,
	}, "")

	logrus.Info(requestLine)

	resp, err := http.Post(requestLine, "application/json", bytes.NewReader([]byte(reqStr)))
	if err != nil || resp.StatusCode != http.StatusOK {
		return Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error(err)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Error(err)
	}
	payload := &v1.Phone{
		Phone: data["phone_info"].(pb.PhoneInfo).PurePhoneNumber,
	}

	return user.NewGetPhoneNumberOK().WithPayload(payload)
}
