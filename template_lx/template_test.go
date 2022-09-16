package template_lx

import (
	"bytes"
	"fmt"
	"geekbang_lx/web"
	"html/template"
	"testing"
)

func TestServer(t *testing.T) {
	s := web.NewHTTPServer()
	s.Get("/", func(ctx *web.Context) {
		_, _ = ctx.Resp.Write([]byte("hello,world"))
	})
	s.Get("/user", func(ctx *web.Context) {
		_, _ = ctx.Resp.Write([]byte("hello , user"))
	})

	s.Post("/form", func(ctx *web.Context) {
		err := ctx.Req.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
	})

	tpl := template.New("login")
	tpl, err := tpl.Parse(`
<html>
	<body>
		<form>
			// 在这里继续写页面
		<form>
	</body>
</html>
`)
	if err != nil {
		t.Fatal(err)
	}
	s.Get("/login", func(ctx *web.Context) {
		page := &bytes.Buffer{}
		err = tpl.Execute(page, nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx.RespStatusCode = 200
		ctx.RespData = page.Bytes()
	})

	_ = s.Start(":8081")
}

func TestServerWithRenderEngine(t *testing.T) {
	tpl, err := template.ParseGlob("testdata/tpls/*.gohtml")
	if err != nil {
		t.Fatal(err)
	}
	s := web.NewHTTPServer(web.ServerWithTemplateEngine(&GoTemplateEngine{T: tpl}))
	s.Get("/login", func(ctx *web.Context) {
		er := ctx.Render("login.gohtml", nil)
		if er != nil {
			t.Fatal(er)
		}
	})
	err = s.Start(":8081")
	if err != nil {
		t.Fatal(err)
	}
}
