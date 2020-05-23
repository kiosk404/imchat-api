package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["imchat-api/controllers:AttachController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:AttachController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:BasicFuncController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:BasicFuncController"],
        beego.ControllerComments{
            Method: "Base",
            Router: `/ping`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:ContactController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:ContactController"],
        beego.ControllerComments{
            Method: "AddFriend",
            Router: `/addfriend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:ContactController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:ContactController"],
        beego.ControllerComments{
            Method: "CreateCommunity",
            Router: `/createcommunity`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:ContactController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:ContactController"],
        beego.ControllerComments{
            Method: "LoadCommunity",
            Router: `/loadcommunity`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:ContactController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:ContactController"],
        beego.ControllerComments{
            Method: "LoadFriend",
            Router: `/loadfriend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"],
        beego.ControllerComments{
            Method: "Main",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"],
        beego.ControllerComments{
            Method: "ChatIndex",
            Router: `/chat`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"],
        beego.ControllerComments{
            Method: "CreateCommunity",
            Router: `/createcom`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:MainTplController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:UserController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserFind",
            Router: `/find`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:UserController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:UserController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["imchat-api/controllers:UserController"] = append(beego.GlobalControllerRouter["imchat-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserRegister",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
