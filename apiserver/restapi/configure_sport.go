// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	mws "github.com/yushk/sport_backend/apiserver/middlewares"
	"github.com/yushk/sport_backend/apiserver/restapi/operations"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/oauth"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/user"
	"github.com/yushk/sport_backend/apiserver/server"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

//go:generate swagger generate server --target ../../apiserver --name Sport --spec ../swagger/swagger.yaml --model-package v1 --principal v1.Principal

func configureFlags(api *operations.SportAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SportAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.OAuth2Auth = func(token string, scopes []string) (*v1.Principal, error) {
		return server.RoleAuth(token, scopes)
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.LoginMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.ModifyUserPasswordMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// user.ResetPasswordMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// oauth.TokenMaxParseMemory = 32 << 20

	api.UserChangeUserPasswordHandler = user.ChangeUserPasswordHandlerFunc(func(params user.ChangeUserPasswordParams, principal *v1.Principal) middleware.Responder {
		return server.ChangeUserPassword(params, principal)
	})
	api.MonitorCreateCaseDataHandler = monitor.CreateCaseDataHandlerFunc(func(params monitor.CreateCaseDataParams, principal *v1.Principal) middleware.Responder {
		return server.CreateCaseData(params, principal)
	})
	api.MonitorCreateClassHandler = monitor.CreateClassHandlerFunc(func(params monitor.CreateClassParams, principal *v1.Principal) middleware.Responder {
		return server.CreateClass(params, principal)
	})
	api.MonitorCreateCommentHandler = monitor.CreateCommentHandlerFunc(func(params monitor.CreateCommentParams, principal *v1.Principal) middleware.Responder {
		return server.CreateComment(params, principal)
	})
	api.MonitorCreateCourseHandler = monitor.CreateCourseHandlerFunc(func(params monitor.CreateCourseParams, principal *v1.Principal) middleware.Responder {
		return server.CreateCourse(params, principal)
	})
	api.MonitorCreateEvaluateInfoHandler = monitor.CreateEvaluateInfoHandlerFunc(func(params monitor.CreateEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
		return server.CreateEvaluateInfo(params, principal)
	})
	api.MonitorCreateHomeWorkHandler = monitor.CreateHomeWorkHandlerFunc(func(params monitor.CreateHomeWorkParams, principal *v1.Principal) middleware.Responder {
		return server.CreateHomeWork(params, principal)
	})
	api.MonitorCreateMeasureDetailHandler = monitor.CreateMeasureDetailHandlerFunc(func(params monitor.CreateMeasureDetailParams, principal *v1.Principal) middleware.Responder {
		return server.CreateMeasureDetail(params, principal)
	})
	api.MonitorCreateMeasureResultHandler = monitor.CreateMeasureResultHandlerFunc(func(params monitor.CreateMeasureResultParams, principal *v1.Principal) middleware.Responder {
		return server.CreateMeasureResult(params, principal)
	})
	api.MonitorCreateMoveDataHandler = monitor.CreateMoveDataHandlerFunc(func(params monitor.CreateMoveDataParams, principal *v1.Principal) middleware.Responder {
		return server.CreateMoveData(params, principal)
	})
	api.MonitorCreateMovePrescriptionHandler = monitor.CreateMovePrescriptionHandlerFunc(func(params monitor.CreateMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
		return server.CreateMovePrescription(params, principal)
	})
	api.UserCreatePersonalHandler = user.CreatePersonalHandlerFunc(func(params user.CreatePersonalParams, principal *v1.Principal) middleware.Responder {
		return server.CreatePersonal(params, principal)
	})
	api.MonitorCreateSolutionHandler = monitor.CreateSolutionHandlerFunc(func(params monitor.CreateSolutionParams, principal *v1.Principal) middleware.Responder {
		return server.CreateSolution(params, principal)
	})
	api.MonitorCreateSportItemHandler = monitor.CreateSportItemHandlerFunc(func(params monitor.CreateSportItemParams, principal *v1.Principal) middleware.Responder {
		return server.CreateSportItem(params, principal)
	})
	api.MonitorCreateSportTypeHandler = monitor.CreateSportTypeHandlerFunc(func(params monitor.CreateSportTypeParams, principal *v1.Principal) middleware.Responder {
		return server.CreateSportType(params, principal)
	})
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *v1.Principal) middleware.Responder {
		return server.CreateUser(params, principal)
	})
	api.MonitorCreateWorkHandler = monitor.CreateWorkHandlerFunc(func(params monitor.CreateWorkParams, principal *v1.Principal) middleware.Responder {
		return server.CreateWork(params, principal)
	})
	api.MonitorCreateWorkSubmitHandler = monitor.CreateWorkSubmitHandlerFunc(func(params monitor.CreateWorkSubmitParams, principal *v1.Principal) middleware.Responder {
		return server.CreateWorkSubmit(params, principal)
	})
	api.MonitorDeleteCaseDataHandler = monitor.DeleteCaseDataHandlerFunc(func(params monitor.DeleteCaseDataParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteCaseData(params, principal)
	})
	api.MonitorDeleteClassHandler = monitor.DeleteClassHandlerFunc(func(params monitor.DeleteClassParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteClass(params, principal)
	})
	api.MonitorDeleteCommentHandler = monitor.DeleteCommentHandlerFunc(func(params monitor.DeleteCommentParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteComment(params, principal)
	})
	api.MonitorDeleteCourseHandler = monitor.DeleteCourseHandlerFunc(func(params monitor.DeleteCourseParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteCourse(params, principal)
	})
	api.MonitorDeleteEvaluateInfoHandler = monitor.DeleteEvaluateInfoHandlerFunc(func(params monitor.DeleteEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteEvaluateInfo(params, principal)
	})
	api.MonitorDeleteHomeWorkHandler = monitor.DeleteHomeWorkHandlerFunc(func(params monitor.DeleteHomeWorkParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteHomeWork(params, principal)
	})
	api.MonitorDeleteMeasureDetailHandler = monitor.DeleteMeasureDetailHandlerFunc(func(params monitor.DeleteMeasureDetailParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteMeasureDetail(params, principal)
	})
	api.MonitorDeleteMeasureResultHandler = monitor.DeleteMeasureResultHandlerFunc(func(params monitor.DeleteMeasureResultParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteMeasureResult(params, principal)
	})
	api.MonitorDeleteMoveDataHandler = monitor.DeleteMoveDataHandlerFunc(func(params monitor.DeleteMoveDataParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteMoveData(params, principal)
	})
	api.MonitorDeleteMovePrescriptionHandler = monitor.DeleteMovePrescriptionHandlerFunc(func(params monitor.DeleteMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteMovePrescription(params, principal)
	})
	api.UserDeletePersonalHandler = user.DeletePersonalHandlerFunc(func(params user.DeletePersonalParams, principal *v1.Principal) middleware.Responder {
		return server.DeletePersonal(params, principal)
	})
	api.MonitorDeleteSolutionHandler = monitor.DeleteSolutionHandlerFunc(func(params monitor.DeleteSolutionParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteSolution(params, principal)
	})
	api.MonitorDeleteSportItemHandler = monitor.DeleteSportItemHandlerFunc(func(params monitor.DeleteSportItemParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteSportItem(params, principal)
	})
	api.MonitorDeleteSportTypeHandler = monitor.DeleteSportTypeHandlerFunc(func(params monitor.DeleteSportTypeParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteSportType(params, principal)
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteUser(params, principal)
	})
	api.MonitorDeleteWorkHandler = monitor.DeleteWorkHandlerFunc(func(params monitor.DeleteWorkParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteWork(params, principal)
	})
	api.MonitorDeleteWorkSubmitHandler = monitor.DeleteWorkSubmitHandlerFunc(func(params monitor.DeleteWorkSubmitParams, principal *v1.Principal) middleware.Responder {
		return server.DeleteWorkSubmit(params, principal)
	})
	api.MonitorGetCaseDataHandler = monitor.GetCaseDataHandlerFunc(func(params monitor.GetCaseDataParams, principal *v1.Principal) middleware.Responder {
		return server.GetCaseData(params, principal)
	})
	api.MonitorGetCaseDatasHandler = monitor.GetCaseDatasHandlerFunc(func(params monitor.GetCaseDatasParams, principal *v1.Principal) middleware.Responder {
		return server.GetCaseDatas(params, principal)
	})
	api.MonitorGetClassHandler = monitor.GetClassHandlerFunc(func(params monitor.GetClassParams, principal *v1.Principal) middleware.Responder {
		return server.GetClass(params, principal)
	})
	api.MonitorGetClassesHandler = monitor.GetClassesHandlerFunc(func(params monitor.GetClassesParams, principal *v1.Principal) middleware.Responder {
		return server.GetClasses(params, principal)
	})
	api.MonitorGetCommentHandler = monitor.GetCommentHandlerFunc(func(params monitor.GetCommentParams, principal *v1.Principal) middleware.Responder {
		return server.GetComment(params, principal)
	})
	api.MonitorGetCommentsHandler = monitor.GetCommentsHandlerFunc(func(params monitor.GetCommentsParams, principal *v1.Principal) middleware.Responder {
		return server.GetComments(params, principal)
	})
	api.MonitorGetCourseHandler = monitor.GetCourseHandlerFunc(func(params monitor.GetCourseParams, principal *v1.Principal) middleware.Responder {
		return server.GetCourse(params, principal)
	})
	api.MonitorGetCoursesHandler = monitor.GetCoursesHandlerFunc(func(params monitor.GetCoursesParams, principal *v1.Principal) middleware.Responder {
		return server.GetCourses(params, principal)
	})
	api.MonitorGetEvaluateInfosHandler = monitor.GetEvaluateInfosHandlerFunc(func(params monitor.GetEvaluateInfosParams, principal *v1.Principal) middleware.Responder {
		return server.GetEvaluateInfos(params, principal)
	})
	api.MonitorGetHomeWorkHandler = monitor.GetHomeWorkHandlerFunc(func(params monitor.GetHomeWorkParams, principal *v1.Principal) middleware.Responder {
		return server.GetHomeWork(params, principal)
	})
	api.MonitorGetHomeWorksHandler = monitor.GetHomeWorksHandlerFunc(func(params monitor.GetHomeWorksParams, principal *v1.Principal) middleware.Responder {
		return server.GetHomeWorks(params, principal)
	})
	api.MonitorGetMeasureDetailHandler = monitor.GetMeasureDetailHandlerFunc(func(params monitor.GetMeasureDetailParams, principal *v1.Principal) middleware.Responder {
		return server.GetMeasureDetail(params, principal)
	})
	api.MonitorGetMeasureDetailsHandler = monitor.GetMeasureDetailsHandlerFunc(func(params monitor.GetMeasureDetailsParams, principal *v1.Principal) middleware.Responder {
		return server.GetMeasureDetails(params, principal)
	})
	api.MonitorGetMeasureResultHandler = monitor.GetMeasureResultHandlerFunc(func(params monitor.GetMeasureResultParams, principal *v1.Principal) middleware.Responder {
		return server.GetMeasureResult(params, principal)
	})
	api.MonitorGetMeasureResultsHandler = monitor.GetMeasureResultsHandlerFunc(func(params monitor.GetMeasureResultsParams, principal *v1.Principal) middleware.Responder {
		return server.GetMeasureResults(params, principal)
	})
	api.MonitorGetMoveDataHandler = monitor.GetMoveDataHandlerFunc(func(params monitor.GetMoveDataParams, principal *v1.Principal) middleware.Responder {
		return server.GetMoveData(params, principal)
	})
	api.MonitorGetMoveDatasHandler = monitor.GetMoveDatasHandlerFunc(func(params monitor.GetMoveDatasParams, principal *v1.Principal) middleware.Responder {
		return server.GetMoveDatas(params, principal)
	})
	api.MonitorGetMovePrescriptionHandler = monitor.GetMovePrescriptionHandlerFunc(func(params monitor.GetMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
		return server.GetMovePrescription(params, principal)
	})
	api.MonitorGetMovePrescriptionsHandler = monitor.GetMovePrescriptionsHandlerFunc(func(params monitor.GetMovePrescriptionsParams, principal *v1.Principal) middleware.Responder {
		return server.GetMovePrescriptions(params, principal)
	})
	api.UserGetPersonalHandler = user.GetPersonalHandlerFunc(func(params user.GetPersonalParams, principal *v1.Principal) middleware.Responder {
		return server.GetPersonal(params, principal)
	})
	api.UserGetPersonalsHandler = user.GetPersonalsHandlerFunc(func(params user.GetPersonalsParams, principal *v1.Principal) middleware.Responder {
		return server.GetPersonals(params, principal)
	})
	api.UserGetSTSTokenHandler = user.GetSTSTokenHandlerFunc(func(params user.GetSTSTokenParams) middleware.Responder {
		return server.GetSTSToken(params)
	})
	api.MonitorGetSolutionHandler = monitor.GetSolutionHandlerFunc(func(params monitor.GetSolutionParams, principal *v1.Principal) middleware.Responder {
		return server.GetSolution(params, principal)
	})
	api.MonitorGetSolutionsHandler = monitor.GetSolutionsHandlerFunc(func(params monitor.GetSolutionsParams, principal *v1.Principal) middleware.Responder {
		return server.GetSolutions(params, principal)
	})
	api.MonitorGetSportItemHandler = monitor.GetSportItemHandlerFunc(func(params monitor.GetSportItemParams, principal *v1.Principal) middleware.Responder {
		return server.GetSportItem(params, principal)
	})
	api.MonitorGetSportItemsHandler = monitor.GetSportItemsHandlerFunc(func(params monitor.GetSportItemsParams, principal *v1.Principal) middleware.Responder {
		return server.GetSportItems(params, principal)
	})
	api.MonitorGetSportTypeHandler = monitor.GetSportTypeHandlerFunc(func(params monitor.GetSportTypeParams, principal *v1.Principal) middleware.Responder {
		return server.GetSportType(params, principal)
	})
	api.MonitorGetSportTypesHandler = monitor.GetSportTypesHandlerFunc(func(params monitor.GetSportTypesParams, principal *v1.Principal) middleware.Responder {
		return server.GetSportTypes(params, principal)
	})
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, principal *v1.Principal) middleware.Responder {
		return server.GetUser(params, principal)
	})
	api.UserGetUserInfoHandler = user.GetUserInfoHandlerFunc(func(params user.GetUserInfoParams) middleware.Responder {
		return server.GetUserInfo(params)
	})
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(func(params user.GetUsersParams, principal *v1.Principal) middleware.Responder {
		return server.GetUsers(params, principal)
	})
	api.MonitorGetWorkHandler = monitor.GetWorkHandlerFunc(func(params monitor.GetWorkParams, principal *v1.Principal) middleware.Responder {
		return server.GetWork(params, principal)
	})
	api.MonitorGetWorkSubmitHandler = monitor.GetWorkSubmitHandlerFunc(func(params monitor.GetWorkSubmitParams, principal *v1.Principal) middleware.Responder {
		return server.GetWorkSubmit(params, principal)
	})
	api.MonitorGetWorkSubmitsHandler = monitor.GetWorkSubmitsHandlerFunc(func(params monitor.GetWorkSubmitsParams, principal *v1.Principal) middleware.Responder {
		return server.GetWorkSubmits(params, principal)
	})
	api.MonitorGetWorksHandler = monitor.GetWorksHandlerFunc(func(params monitor.GetWorksParams, principal *v1.Principal) middleware.Responder {
		return server.GetWorks(params, principal)
	})
	api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
		return server.Login(params)
	})
	api.UserLogoutHandler = user.LogoutHandlerFunc(func(params user.LogoutParams) middleware.Responder {
		return server.Logout(params)
	})
	api.UserModifyUserPasswordHandler = user.ModifyUserPasswordHandlerFunc(func(params user.ModifyUserPasswordParams, principal *v1.Principal) middleware.Responder {
		return server.ModifyUserPassword(params, principal)
	})
	api.OauthTokenHandler = oauth.TokenHandlerFunc(func(params oauth.TokenParams) middleware.Responder {
		return server.Token(params)
	})
	api.MonitorUpdateCaseDataHandler = monitor.UpdateCaseDataHandlerFunc(func(params monitor.UpdateCaseDataParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateCaseData(params, principal)
	})
	api.MonitorUpdateClassHandler = monitor.UpdateClassHandlerFunc(func(params monitor.UpdateClassParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateClass(params, principal)
	})
	api.MonitorUpdateCommentHandler = monitor.UpdateCommentHandlerFunc(func(params monitor.UpdateCommentParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateComment(params, principal)
	})
	api.MonitorUpdateCourseHandler = monitor.UpdateCourseHandlerFunc(func(params monitor.UpdateCourseParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateCourse(params, principal)
	})
	api.MonitorUpdateEvaluateInfoHandler = monitor.UpdateEvaluateInfoHandlerFunc(func(params monitor.UpdateEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateEvaluateInfo(params, principal)
	})
	api.MonitorUpdateHomeWorkHandler = monitor.UpdateHomeWorkHandlerFunc(func(params monitor.UpdateHomeWorkParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateHomeWork(params, principal)
	})
	api.MonitorUpdateMeasureDetailHandler = monitor.UpdateMeasureDetailHandlerFunc(func(params monitor.UpdateMeasureDetailParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateMeasureDetail(params, principal)
	})
	api.MonitorUpdateMeasureResultHandler = monitor.UpdateMeasureResultHandlerFunc(func(params monitor.UpdateMeasureResultParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateMeasureResult(params, principal)
	})
	api.MonitorUpdateMoveDataHandler = monitor.UpdateMoveDataHandlerFunc(func(params monitor.UpdateMoveDataParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateMoveData(params, principal)
	})
	api.MonitorUpdateMovePrescriptionHandler = monitor.UpdateMovePrescriptionHandlerFunc(func(params monitor.UpdateMovePrescriptionParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateMovePrescription(params, principal)
	})
	api.UserUpdatePersonalHandler = user.UpdatePersonalHandlerFunc(func(params user.UpdatePersonalParams, principal *v1.Principal) middleware.Responder {
		return server.UpdatePersonal(params, principal)
	})
	api.MonitorUpdateSolutionHandler = monitor.UpdateSolutionHandlerFunc(func(params monitor.UpdateSolutionParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateSolution(params, principal)
	})
	api.MonitorUpdateSportItemHandler = monitor.UpdateSportItemHandlerFunc(func(params monitor.UpdateSportItemParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateSportItem(params, principal)
	})
	api.MonitorUpdateSportTypeHandler = monitor.UpdateSportTypeHandlerFunc(func(params monitor.UpdateSportTypeParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateSportType(params, principal)
	})
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateUser(params, principal)
	})
	api.MonitorUpdateWorkHandler = monitor.UpdateWorkHandlerFunc(func(params monitor.UpdateWorkParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateWork(params, principal)
	})
	api.MonitorUpdateWorkSubmitHandler = monitor.UpdateWorkSubmitHandlerFunc(func(params monitor.UpdateWorkSubmitParams, principal *v1.Principal) middleware.Responder {
		return server.UpdateWorkSubmit(params, principal)
	})
	api.UserWXLoginHandler = user.WXLoginHandlerFunc(func(params user.WXLoginParams) middleware.Responder {
		return server.WXLogin(params)
	})
	api.UserGetPhoneNumberHandler = user.GetPhoneNumberHandlerFunc(func(params user.GetPhoneNumberParams, principal *v1.Principal) middleware.Responder {
		return server.GetPhoneNumber(params, principal)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	server.RegisterClients()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return mws.Limiter(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return mws.HandlePanic(mws.RedocUI(mws.LogViaLogrus(mws.Cross(handler))))
}
