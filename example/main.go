package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohuishou/protoc-gen-go-gin/example/api/product/app/ecode"
	v1 "github.com/mohuishou/protoc-gen-go-gin/example/api/product/app/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
}

func (s service) CreateArticle(ctx *gin.Context, article *v1.Article) (*v1.Article, error) {
	if article.AuthorId < 1 {
		return nil, ecode.Errorf(http.StatusBadRequest, 400, "author id must > 0")
	}
	return article, nil
}

func (s service) GetArticles(ctx *gin.Context, req *v1.GetArticlesReq) (*v1.GetArticlesResp, error) {
	if req.AuthorId < 0 {
		return nil, ecode.Errorf(http.StatusBadRequest, 400, "author id must >= 0")
	}
	return &v1.GetArticlesResp{
		Total: 1,
		Articles: []*v1.Article{
			{
				Title:    "test article: " + req.Title,
				Content:  "test",
				AuthorId: 1,
			},
		},
	}, nil
}

// DownloadArticle 自定义响应示例：文件下载
// 返回 nil 表示我们已经自己处理了响应
func (s service) DownloadArticle(ctx *gin.Context, req *v1.DownloadArticleReq) (*emptypb.Empty, error) {
	// 设置文件下载的响应头
	ctx.Header("Content-Type", "text/plain")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=article_%d.txt", req.ArticleId))
	
	// 模拟文件内容
	content := fmt.Sprintf("Article ID: %d\nTitle: Sample Article\nContent: This is a sample article for download.", req.ArticleId)
	
	// 直接写入文件内容
	ctx.String(http.StatusOK, content)
	
	// 返回 nil 表示已自己处理响应
	return nil, nil
}

// GetArticleXML 自定义响应示例：XML 输出
// 返回 nil 表示我们已经自己处理了响应
func (s service) GetArticleXML(ctx *gin.Context, req *v1.GetArticleXMLReq) (*emptypb.Empty, error) {
	// 定义要输出的数据结构
	type XMLArticle struct {
		ID      int32  `xml:"id"`
		Title   string `xml:"title"`
		Content string `xml:"content"`
	}
	
	// 创建示例数据
	article := XMLArticle{
		ID:      req.ArticleId,
		Title:   "Sample Article",
		Content: "This is a sample article in XML format",
	}
	
	// 使用 gin 的 XML 方法输出
	ctx.XML(http.StatusOK, article)
	
	// 返回 nil 表示已自己处理响应
	return nil, nil
}

func main() {
	e := gin.Default()
	v1.RegisterBlogServiceHTTPServer(e, &service{})
	e.Run()
}
