// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGormController struct {}
var GormController tGormController


func (_ tGormController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Begin", args).URL
}

func (_ tGormController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Commit", args).URL
}

func (_ tGormController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Rollback", args).URL
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tTorrentController struct {}
var TorrentController tTorrentController


func (_ tTorrentController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TorrentController.List", args).URL
}

func (_ tTorrentController) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TorrentController.Add", args).URL
}

func (_ tTorrentController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TorrentController.Create", args).URL
}


type tMemberController struct {}
var MemberController tMemberController


func (_ tMemberController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberController.Login", args).URL
}

func (_ tMemberController) Sign(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberController.Sign", args).URL
}

func (_ tMemberController) Create(
		member interface{},
		userPwConfirm string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "member", member)
	revel.Unbind(args, "userPwConfirm", userPwConfirm)
	return revel.MainRouter.Reverse("MemberController.Create", args).URL
}

func (_ tMemberController) Get(
		member interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "member", member)
	return revel.MainRouter.Reverse("MemberController.Get", args).URL
}


