package convert

import (
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func PersonalPb2V1(pbPersonal *pb.Personal) (v1Personal *v1.Personal) {
	if pbPersonal == nil {
		return
	}
	class := []*v1.Class{}
	if len(pbPersonal.Class) > 0 {
		for _, v := range pbPersonal.Class {
			class = append(class, ClassPb2V1(v))
		}
	}
	v1Personal = &v1.Personal{
		ID:     pbPersonal.Id,
		Class:  class,
		Type:   pbPersonal.Type,
		Name:   pbPersonal.Name,
		Userid: pbPersonal.Userid,
		Gender: pbPersonal.Gender,
		Birth:  pbPersonal.Birth,
	}
	return
}

func UserPb2V1(pbUser *pb.User) (v1User *v1.User) {
	if pbUser == nil {
		return
	}

	v1User = &v1.User{
		ID:   pbUser.Id,
		Name: pbUser.Name,
		// Ps:      pbUser.Ps,
		Role:    pbUser.Role,
		Phone:   pbUser.Phone,
		Email:   pbUser.Email,
		Openid:  pbUser.Openid,
		Unionid: pbUser.Unionid,
		Avatar: pbUser.Avatar,
		NickName: pbUser.NickName,
	}
	return
}

func CaseDataPb2V1(pbCaseData *pb.CaseData) (v1CaseData *v1.CaseData) {
	if pbCaseData == nil {
		return
	}

	v1CaseData = &v1.CaseData{
		ID:     pbCaseData.Id,
		Userid: pbCaseData.Userid,
	}
	return
}

func ClassPb2V1(pbClass *pb.Class) (v1Class *v1.Class) {
	if pbClass == nil {
		return
	}

	v1Class = &v1.Class{
		ID:      pbClass.Id,
		Grade:   pbClass.Grade,
		Faculty: pbClass.Faculty,
		Number:  pbClass.Number,
		Subject: pbClass.Subject,
	}
	return
}

func MeasureDetailPb2V1(pbMeasureDetail *pb.MeasureDetail) (v1MeasureDetail *v1.MeasureDetail) {
	if pbMeasureDetail == nil {
		return
	}

	v1MeasureDetail = &v1.MeasureDetail{
		ID:     pbMeasureDetail.Id,
		Userid: pbMeasureDetail.Userid,
	}
	return
}

func MeasureResultPb2V1(pbMeasureResult *pb.MeasureResult) (v1MeasureResult *v1.MeasureResult) {
	if pbMeasureResult == nil {
		return
	}

	v1MeasureResult = &v1.MeasureResult{
		ID:     pbMeasureResult.Id,
		Userid: pbMeasureResult.Userid,
	}
	return
}

func MoveDataPb2V1(pbMoveData *pb.MoveData) (v1MoveData *v1.MoveData) {
	if pbMoveData == nil {
		return
	}

	v1MoveData = &v1.MoveData{
		ID:     pbMoveData.Id,
		Userid: pbMoveData.Userid,
	}
	return
}

func MovePrescriptionPb2V1(pbMovePrescription *pb.MovePrescription) (v1MovePrescription *v1.MovePrescription) {
	if pbMovePrescription == nil {
		return
	}

	v1MovePrescription = &v1.MovePrescription{
		ID:     pbMovePrescription.Id,
		Userid: pbMovePrescription.Userid,
	}
	return
}

func FileInfoPb2V1(pbFileInfo *pb.FileInfo) (v1FileInfo *v1.FileInfo) {
	v1FileInfo = &v1.FileInfo{
		Name: pbFileInfo.Name,
		URL:  pbFileInfo.Url,
	}
	return
}

func NoteInfoPb2V1(pbNoteInfo *pb.NoteInfo) (v1NoteInfo *v1.NoteInfo) {
	v1NoteInfo = &v1.NoteInfo{
		Userid:  pbNoteInfo.Userid,
		Score:   pbNoteInfo.Score,
		Content: pbNoteInfo.Content,
		Status:  pbNoteInfo.Status,
		Value:   pbNoteInfo.Value,
	}
	return
}

func CoursePb2V1(pbCourse *pb.Course) (v1Course *v1.Course) {
	if pbCourse == nil {
		return
	}

	files := []*v1.FileInfo{}
	if len(pbCourse.Files) > 0 {
		for _, v := range pbCourse.Files {
			files = append(files, FileInfoPb2V1(v))
		}
	}
	v1Course = &v1.Course{
		ID:       pbCourse.Id,
		Userid:   pbCourse.Userid,
		Title:    pbCourse.Title,
		Files:    files,
		Content:  pbCourse.Content,
		Desc:     pbCourse.Desc,
		Type:     pbCourse.Type,
		Create:   pbCourse.Create,
		Update:   pbCourse.Update,
		Username: pbCourse.Username,
	}
	return
}

func WorkPb2V1(pbWork *pb.Work) (v1Work *v1.Work) {
	if pbWork == nil {
		return
	}
	files := []*v1.FileInfo{}
	if len(pbWork.Files) > 0 {
		for _, v := range pbWork.Files {
			files = append(files, FileInfoPb2V1(v))
		}
	}
	noteInfos := []*v1.NoteInfo{}
	if len(pbWork.NoteInfo) > 0 {
		for _, v := range pbWork.NoteInfo {
			noteInfos = append(noteInfos, NoteInfoPb2V1(v))
		}
	}
	noters := []string{}
	if len(pbWork.Noter) > 0 {
		noters = append(noters, pbWork.Noter...)
	}
	v1Work = &v1.Work{
		ID:            pbWork.Id,
		Userid:        pbWork.Userid,
		Course:        pbWork.Course,
		CourseTitle:   pbWork.CourseTitle,
		Files:         files,
		Content:       pbWork.Content,
		Note:          pbWork.Note,
		Score:         pbWork.Score,
		Type:          pbWork.Type,
		Create:        pbWork.Create,
		Update:        pbWork.Update,
		Username:      pbWork.Username,
		CourseCreater: pbWork.CourseCreater,
		Teacherid:     pbWork.Teacherid,
		NoteInfo:      noteInfos,
		Noter:         noters,
	}
	return
}

func WorkSubmitPb2V1(pbWorkSubmit *pb.WorkSubmit) (v1WorkSubmit *v1.WorkSubmit) {
	if pbWorkSubmit == nil {
		return
	}

	v1WorkSubmit = &v1.WorkSubmit{
		ID:        pbWorkSubmit.Id,
		Studentid: pbWorkSubmit.Studentid,
		Taskid:    pbWorkSubmit.Taskid,
		Content:   pbWorkSubmit.Content,
	}
	return
}

func CommentPb2V1(pbComment *pb.Comment) (v1Comment *v1.Comment) {
	if pbComment == nil {
		return
	}

	v1Comment = &v1.Comment{
		ID:          pbComment.Id,
		Courseid:    pbComment.Courseid,
		CommentUser: pbComment.CommentUser,
		Comment:     pbComment.Comment,
		Score:       pbComment.Score,
	}
	return
}

func SportTypePb2V1(pbSportType *pb.SportType) (v1SportType *v1.SportType) {
	v1SportType = &v1.SportType{
		ID:    pbSportType.Id,
		Label: pbSportType.Label,
		Code:  pbSportType.Code,
	}
	return
}

func EvaluatePb2V1(pbEvaluate *pb.Evaluate) (v1Evaluate *v1.Evaluate) {
	if pbEvaluate == nil {
		return
	}

	v1Evaluate = &v1.Evaluate{
		ID:        pbEvaluate.Id,
		Name:      pbEvaluate.Name,
		Excellent: pbEvaluate.Excellent,
		Good:      pbEvaluate.Good,
		Pass:      pbEvaluate.Pass,
		Flunk:     pbEvaluate.Flunk,
	}
	return
}

func SportItemPb2V1(pbSportItem *pb.SportItem) (v1SportItem *v1.SportItem) {
	if pbSportItem == nil {
		return
	}
	pics := []*v1.FileInfo{}
	if len(pbSportItem.Pics) > 0 {
		for _, v := range pbSportItem.Pics {
			pics = append(pics, FileInfoPb2V1(v))
		}
	}
	files := []*v1.FileInfo{}
	if len(pbSportItem.Files) > 0 {
		for _, v := range pbSportItem.Files {
			files = append(files, FileInfoPb2V1(v))
		}
	}
	videos := []*v1.FileInfo{}
	if len(pbSportItem.Videos) > 0 {
		for _, v := range pbSportItem.Videos {
			videos = append(videos, FileInfoPb2V1(v))
		}
	}
	authUsers := []string{}
	if len(pbSportItem.AuthUser) > 0 {
		authUsers = append(authUsers, pbSportItem.AuthUser...)
	}
	teachers := []string{}
	if len(pbSportItem.Teachers) > 0 {
		teachers = append(teachers, pbSportItem.Teachers...)
	}
	evaluates := []*v1.Evaluate{}
	if len(pbSportItem.Evaluates) > 0 {
		for _, v := range pbSportItem.Evaluates {
			evaluates = append(evaluates, EvaluatePb2V1(v))
		}
	}
	v1SportItem = &v1.SportItem{
		ID:         pbSportItem.Id,
		Lable:      pbSportItem.Lable,
		Content:    pbSportItem.Content,
		Pics:       pics,
		Files:      files,
		Videos:     videos,
		Userid:     pbSportItem.Userid,
		AuthUser:   authUsers,
		CreateTime: pbSportItem.CreateTime,
		UpdateTime: pbSportItem.UpdateTime,
		Teachers:   teachers,
		Type:       pbSportItem.Type,
		SportType:  pbSportItem.SportType,
		Version:    pbSportItem.Version,
		Evaluates:  evaluates,
	}
	return
}

func HomeWorkPb2V1(pbHomeWork *pb.HomeWork) (v1HomeWork *v1.HomeWork) {
	if pbHomeWork == nil {
		return
	}
	solutions := []*v1.Solution{}
	if len(pbHomeWork.SolutionsInfo) > 0 {
		for _, v := range pbHomeWork.SolutionsInfo {
			solutions = append(solutions, SolutionPb2V1(v))
		}
	}
	solutionsId := []string{}
	if len(pbHomeWork.SolutionsInfoId) > 0 {
		solutionsId = append(solutionsId, pbHomeWork.SolutionsInfoId...)
	}
	v1HomeWork = &v1.HomeWork{
		ID:              pbHomeWork.Id,
		SportItemID:     pbHomeWork.SportItemId,
		SportItem:       SportItemPb2V1(pbHomeWork.SportItem),
		Teacher:         pbHomeWork.Teacher,
		StartTime:       pbHomeWork.StartTime,
		EndTime:         pbHomeWork.EndTime,
		QRCode:          pbHomeWork.QRCode,
		Number:          pbHomeWork.Number,
		Desc:            pbHomeWork.Desc,
		SolutionsInfo:   solutions,
		SolutionsInfoID: solutionsId,
		Title:           pbHomeWork.Title,
	}
	return
}

func EvaluateInfoPb2V1(pbEvaluateInfo *pb.EvaluateInfo) (v1EvaluateInfo *v1.EvaluateInfo) {
	if pbEvaluateInfo == nil {
		return
	}
	evaluates := []*v1.Evaluate{}
	if len(pbEvaluateInfo.Evaluates) > 0 {
		for _, v := range pbEvaluateInfo.Evaluates {
			evaluates = append(evaluates, EvaluatePb2V1(v))
		}
	}
	v1EvaluateInfo = &v1.EvaluateInfo{
		ID:         pbEvaluateInfo.Id,
		HomeWorkID: pbEvaluateInfo.HomeWorkId,
		SolutionID: pbEvaluateInfo.SolutionId,
		UserID:     pbEvaluateInfo.UserId,
		Name:       pbEvaluateInfo.Name,
		Evaluates:  evaluates,
		Note:       pbEvaluateInfo.Note,
	}
	return
}

func SolutionPb2V1(pbSolution *pb.Solution) (v1Solution *v1.Solution) {
	if pbSolution == nil {
		return
	}
	evaluates := []*v1.EvaluateInfo{}
	if len(pbSolution.EvaluatesInfo) > 0 {
		for _, v := range pbSolution.EvaluatesInfo {
			evaluates = append(evaluates, EvaluateInfoPb2V1(v))
		}
	}
	teachers := []string{}
	if len(pbSolution.Teachers) > 0 {
		teachers = append(teachers, pbSolution.Teachers...)
	}
	v1Solution = &v1.Solution{
		ID:            pbSolution.Id,
		Title:         pbSolution.Title,
		WorkID:        pbSolution.WorkId,
		StudentID:     pbSolution.StudentId,
		StudentName:   pbSolution.StudentName,
		Content:       pbSolution.Content,
		EvaluatesInfo: evaluates,
		CommitTime:    pbSolution.CommitTime,
		Desc:          pbSolution.Desc,
		Teachers:      teachers,
	}
	return
}
