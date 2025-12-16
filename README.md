# protoc-gen-go-gin

> 修改自 [kratos v2](https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http)

从 protobuf 文件中生成使用 gin 的 http rpc 服务
## 安装

请确保安装了以下依赖:

- [go 1.16](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

注意由于使用 embed 特性，Go 版本必须大于 1.16

```bash
go install github.com/mohuishou/protoc-gen-go-gin@latest
```

## 使用说明

例子见: [example](./example)

### proto 文件约定

默认情况下 rpc method 命名为 方法+资源，使用驼峰方式命名，生成代码时会进行映射

方法映射方式如下所示:

- `"GET", "FIND", "QUERY", "LIST", "SEARCH"`  --> GET
- `"POST", "CREATE"`  --> POST
- `"PUT", "UPDATE"`  --> PUT
- `"DELETE"`  --> DELETE

```protobuf
service BlogService {
  rpc CreateArticle(Article) returns (Article) {}
  // 生成 http 路由为 post: /article
}
```

除此之外还可以使用 google.api.http option 指定路由，可以通过添加 additional_bindings 使一个 rpc 方法对应多个路由

```protobuf
// blog service is a blog demo
service BlogService {
  rpc GetArticles(GetArticlesReq) returns (GetArticlesResp) {
    // 
    // 可以通过添加 additional_bindings 使一个 rpc 方法对应多个路由
    option (google.api.http) = {
      get: "/v1/articles"
      additional_bindings {
        get: "/v1/author/{author_id}/articles"
      }
    };
  }
}
```

### 文件生成

```bash
  protoc -I ./example/api \
  --go_out ./example/api --go_opt=paths=source_relative \
  --go-gin_out ./example/api --go-gin_opt=paths=source_relative \
  example/api/product/app/v1/v1.proto
```

### 响应处理

生成的代码支持灵活的响应处理机制：

#### 1. 标准 JSON 响应（默认）

当 RPC 方法返回数据时，会自动包装成 JSON 格式输出：

```protobuf
service UserService {
  rpc GetUser(GetUserReq) returns (User) {}
}
```

```go
func (s *Service) GetUser(ctx *gin.Context, req *pb.GetUserReq) (*pb.User, error) {
    user := &pb.User{Name: "John", Age: 30}
    return user, nil  // 自动输出 JSON: {"code": 0, "msg": "Success", "data": {...}}
}
```

#### 2. 自定义响应格式（文件下载、XML、代理转发等）

对于需要自定义响应的场景（如文件下载、XML 输出、代理转发等），可以：
- 使用 `google.protobuf.Empty` 作为返回类型
- 在方法内直接使用 `gin.Context` 处理响应
- 返回 `nil, nil`

```protobuf
import "google/protobuf/empty.proto";

service FileService {
  rpc DownloadFile(FileReq) returns (google.protobuf.Empty) {}
  rpc GetXML(DataReq) returns (google.protobuf.Empty) {}
}
```

```go
// 文件下载示例
func (s *Service) DownloadFile(ctx *gin.Context, req *pb.FileReq) (*emptypb.Empty, error) {
    ctx.Header("Content-Type", "application/pdf")
    ctx.Header("Content-Disposition", "attachment; filename=report.pdf")
    ctx.File("/path/to/file.pdf")
    return nil, nil
}

// XML 响应示例
func (s *Service) GetXML(ctx *gin.Context, req *pb.DataReq) (*emptypb.Empty, error) {
    ctx.Header("Content-Type", "application/xml")
    ctx.XML(200, gin.H{"data": "value"})
    return nil, nil
}

// 代理转发示例
func (s *Service) ProxyRequest(ctx *gin.Context, req *pb.ProxyReq) (*emptypb.Empty, error) {
    // 转发到其他服务
    s.forwardToProxy(ctx, req)
    return nil, nil
}
```

#### 3. 严格模式错误检测

生成的代码会自动检测响应处理的完整性：

- ✅ 如果返回数据不为 `nil`，自动调用 `Success()` 输出 JSON
- ✅ 如果返回 `nil` 但已写入响应（如调用了 `ctx.File()`、`ctx.JSON()` 等），正常完成
- ❌ 如果返回 `nil` 且未写入任何响应，返回错误避免客户端卡住

这个机制确保开发者在使用自定义响应时不会忘记处理响应，避免客户端无限等待的问题。

```go
// ❌ 错误示例：返回 nil 但忘记写入响应
func (s *Service) BadHandler(ctx *gin.Context, req *pb.Req) (*emptypb.Empty, error) {
    // 忘记调用 ctx.JSON/ctx.File 等方法
    return nil, nil
}
// 结果：客户端会收到错误响应 "handler returned nil but no response was written"
```

## 相关介绍

> 待发布

- Go工程化(四) API 设计上: 项目结构 & 设计
- Go工程化(五) API 设计下: 基于 protobuf 自动生成 gin 代码
