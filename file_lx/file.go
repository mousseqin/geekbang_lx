package file_lx

import (
	"geekbang_lx/web"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type FileUpload struct {
	// FileField 对应于文件在表单中的字段名字
	FileField string
	// DstPathFunc 用于计算目标路径
	DstPathFunc func(fh *multipart.FileHeader) string
}

func (f *FileUpload) Handle() web.HandleFunc {
	// 这里可以额外做一些检测
	// if f.FileField == "" {
	// 	// 这种方案默认值我其实不是很喜欢
	// 	// 因为我们需要教会用户说，这个 file 是指什么意思
	// 	f.FileField = "file"
	// }
	return func(ctx *web.Context) {
		src, srcHeader, err := ctx.Req.FormFile(f.FileField)
		defer src.Close()
		if err != nil {
			ctx.RespStatusCode = 400
			ctx.RespData = []byte("上传失败，未找到数据")
			log.Fatalln(err)
			return
		}
		dst, err := os.OpenFile(f.DstPathFunc(srcHeader),
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		defer dst.Close()
		if err != nil {
			ctx.RespStatusCode = 500
			ctx.RespData = []byte("上传失败")
			log.Fatalln(err)
			return
		}

		_, err = io.CopyBuffer(dst, src, nil)
		if err != nil {
			ctx.RespStatusCode = 500
			ctx.RespData = []byte("上传失败")
			log.Fatalln(err)
			return
		}
		ctx.RespData = []byte("上传成功")
	}
}
