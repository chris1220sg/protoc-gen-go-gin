var _ = io.EOF // 防止 "imported and not used" 报错

var marshaler = protojson.MarshalOptions{
	UseProtoNames:   true,
	EmitUnpopulated: true,
}
var unmarshaler = protojson.UnmarshalOptions{
	AllowPartial:   false,
	DiscardUnknown: true,
}

type {{ $.InterfaceName }} interface {
{{range .MethodSet}}
	{{.Name}}(*gin.Context, *{{.Request}}) (*{{.Reply}}, error)
{{end}}
}
func Register{{ $.InterfaceName }}(r gin.IRouter, srv {{ $.InterfaceName }}) {
	s := {{.Name}}{
		server: srv,
		router:     r,
		resp: default{{$.Name}}Resp{},
	}
	s.RegisterService()
}

type {{$.Name}} struct{
	server {{ $.InterfaceName }}
	router gin.IRouter
	resp  interface {
		Error(ctx *gin.Context, err error)
		ParamsError (ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type default{{$.Name}}Resp struct {}

func (resp default{{$.Name}}Resp) response(ctx *gin.Context, status, code int, msg string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code,
		"msg": msg,
		"data": data,
	})
}

// Error 返回错误信息
func (resp default{{$.Name}}Resp) Error(ctx *gin.Context, err error) {
	code := -1
	status := 500
	msg := "Unknown error"
	
	if err == nil {
		msg += ", err is nil"
		resp.response(ctx, status, code, msg, nil)
		return
	}

	type iCode interface{
		HTTPCode() int
		Message() string
		Code() int
	}

	var c iCode
	if errors.As(err, &c) {
		status = c.HTTPCode()
		code = c.Code()
		msg = c.Message()
	}

	type iError interface{
        GetCode() int32
        GetMessage() string
    }

    var e iError
    if errors.As(err, &e) {
        status = 200
        code = int(e.GetCode())
        msg = e.GetMessage()
    }

	_ = ctx.Error(err)

	resp.response(ctx, status, code, msg, nil)
}

// ParamsError 参数错误
func (resp default{{$.Name}}Resp) ParamsError (ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.response(ctx, 400, 400, "Parameter error", nil)
}

// Success 返回成功信息
func (resp default{{$.Name}}Resp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "Success", data)
}


{{range .Methods}}
func (s *{{$.Name}}) {{ .HandlerName }} (ctx *gin.Context) {
	var in {{.Request}}
{{if .HasPathParams }}
	if err := ctx.ShouldBindUri(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}
{{end}}
{{if eq .Method "GET" "DELETE" }}
	if err := ctx.ShouldBindQuery(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}
{{else if eq .Method "POST" "PUT" }}
	reqRaw, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}
	if err := unmarshaler.Unmarshal(reqRaw, &in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}
{{else}}
	if err := ctx.ShouldBind(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}
{{end}}
	//md := metadata.New(nil)
	//for k, v := range ctx.Request.Header {
	//	md.Set(k, v...)
	//}
	//newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.({{ $.InterfaceName }}).{{.Name}}(ctx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	// 如果有返回值，输出 JSON
	if out != nil {
		s.resp.Success(ctx, out)
		return
	}

	// 如果返回 nil，检查是否已经写入响应（如文件下载、代理转发等）
	if !ctx.Writer.Written() {
		// 严格模式：开发者忘记处理响应，返回错误
		s.resp.Error(ctx, errors.New("handler returned nil but no response was written"))
	}
}
{{end}}

func (s *{{$.Name}}) RegisterService() {
{{range .Methods}}
		s.router.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}
