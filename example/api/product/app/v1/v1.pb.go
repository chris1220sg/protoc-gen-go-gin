// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: product/app/v1/v1.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetArticlesReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: form:"title"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title"`
	// @inject_tag: form:"page"
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	// @inject_tag: form:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId      int32 `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticlesReq) Reset() {
	*x = GetArticlesReq{}
	mi := &file_product_app_v1_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticlesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesReq) ProtoMessage() {}

func (x *GetArticlesReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_app_v1_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesReq.ProtoReflect.Descriptor instead.
func (*GetArticlesReq) Descriptor() ([]byte, []int) {
	return file_product_app_v1_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetArticlesReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetArticlesReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetArticlesReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetArticlesReq) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type GetArticlesResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Articles      []*Article             `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticlesResp) Reset() {
	*x = GetArticlesResp{}
	mi := &file_product_app_v1_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticlesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesResp) ProtoMessage() {}

func (x *GetArticlesResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_app_v1_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesResp.ProtoReflect.Descriptor instead.
func (*GetArticlesResp) Descriptor() ([]byte, []int) {
	return file_product_app_v1_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticlesResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetArticlesResp) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type Article struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Title   string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId      int32 `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_product_app_v1_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_product_app_v1_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_product_app_v1_v1_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

var File_product_app_v1_v1_proto protoreflect.FileDescriptor

const file_product_app_v1_v1_proto_rawDesc = "" +
	"\n" +
	"\x17product/app/v1/v1.proto\x12\x0eproduct.app.v1\x1a\x1cgoogle/api/annotations.proto\"t\n" +
	"\x0eGetArticlesReq\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x05R\bpageSize\x12\x1b\n" +
	"\tauthor_id\x18\x04 \x01(\x05R\bauthorId\"\\\n" +
	"\x0fGetArticlesResp\x12\x14\n" +
	"\x05total\x18\x01 \x01(\x03R\x05total\x123\n" +
	"\barticles\x18\x02 \x03(\v2\x17.product.app.v1.ArticleR\barticles\"V\n" +
	"\aArticle\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1b\n" +
	"\tauthor_id\x18\x03 \x01(\x05R\bauthorId2\x83\x02\n" +
	"\vBlogService\x12\x87\x01\n" +
	"\vGetArticles\x12\x1e.product.app.v1.GetArticlesReq\x1a\x1f.product.app.v1.GetArticlesResp\"7\x82\xd3\xe4\x93\x021Z!\x12\x1f/v1/author/{author_id}/articles\x12\f/v1/articles\x12j\n" +
	"\rCreateArticle\x12\x17.product.app.v1.Article\x1a\x17.product.app.v1.Article\"'\x82\xd3\xe4\x93\x02!\"\x1f/v1/author/{author_id}/articlesBCZAgithub.com/mohuishou/protoc-gen-go-gin/example/api/product/app/v1b\x06proto3"

var (
	file_product_app_v1_v1_proto_rawDescOnce sync.Once
	file_product_app_v1_v1_proto_rawDescData []byte
)

func file_product_app_v1_v1_proto_rawDescGZIP() []byte {
	file_product_app_v1_v1_proto_rawDescOnce.Do(func() {
		file_product_app_v1_v1_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_app_v1_v1_proto_rawDesc), len(file_product_app_v1_v1_proto_rawDesc)))
	})
	return file_product_app_v1_v1_proto_rawDescData
}

var file_product_app_v1_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_app_v1_v1_proto_goTypes = []any{
	(*GetArticlesReq)(nil),  // 0: product.app.v1.GetArticlesReq
	(*GetArticlesResp)(nil), // 1: product.app.v1.GetArticlesResp
	(*Article)(nil),         // 2: product.app.v1.Article
}
var file_product_app_v1_v1_proto_depIdxs = []int32{
	2, // 0: product.app.v1.GetArticlesResp.articles:type_name -> product.app.v1.Article
	0, // 1: product.app.v1.BlogService.GetArticles:input_type -> product.app.v1.GetArticlesReq
	2, // 2: product.app.v1.BlogService.CreateArticle:input_type -> product.app.v1.Article
	1, // 3: product.app.v1.BlogService.GetArticles:output_type -> product.app.v1.GetArticlesResp
	2, // 4: product.app.v1.BlogService.CreateArticle:output_type -> product.app.v1.Article
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_app_v1_v1_proto_init() }
func file_product_app_v1_v1_proto_init() {
	if File_product_app_v1_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_app_v1_v1_proto_rawDesc), len(file_product_app_v1_v1_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_app_v1_v1_proto_goTypes,
		DependencyIndexes: file_product_app_v1_v1_proto_depIdxs,
		MessageInfos:      file_product_app_v1_v1_proto_msgTypes,
	}.Build()
	File_product_app_v1_v1_proto = out.File
	file_product_app_v1_v1_proto_goTypes = nil
	file_product_app_v1_v1_proto_depIdxs = nil
}
