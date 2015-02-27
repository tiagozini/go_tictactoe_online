// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tRefresh struct {}
var Refresh tRefresh


func (_ tRefresh) Room(
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("Refresh.Room", args).Url
}

func (_ tRefresh) Play(
		userId int,
		x int,
		y int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	revel.Unbind(args, "x", x)
	revel.Unbind(args, "y", y)
	return revel.MainRouter.Reverse("Refresh.Play", args).Url
}

func (_ tRefresh) AskRematch(
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("Refresh.AskRematch", args).Url
}

func (_ tRefresh) RejectRematch(
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("Refresh.RejectRematch", args).Url
}

func (_ tRefresh) RefreshSession(
		userId int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("Refresh.RefreshSession", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Welcome(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Welcome", args).Url
}

func (_ tApp) EnterGame(
		nickname string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nickname", nickname)
	return revel.MainRouter.Reverse("App.EnterGame", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
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
	return revel.MainRouter.Reverse("Static.Serve", args).Url
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
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


