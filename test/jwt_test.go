package test

import (
	"fmt"
	"gin-go-bl/internal/middlewares"
	uuid "github.com/satori/go.uuid"
	"testing"
)

// sample token is expired.  override time so it parses as valid

func TestOne(t *testing.T) {
	// sample token is expired.  override time so it parses as valid
	//time.Unix(0, 0)
	//fmt.Println(time.Now().Unix() + 60*60*24*7)
	//e()
	//a := middlewares.NewJWT()
	//a.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiTmlja05hbWUiOiLmtYvor5UiLCJVVUlEIjoiMjRlMjgxOTMtMDIyYi00ZTM2LTk4ZmEtNTY0MDI0NzExOGFjIiwiZXhwIjoxNjk0MjY4NTExLCJpYXQiOjE2OTM2NjM3MTEsImlzcyI6InRlc3QiLCJuYmYiOjE2OTM2NjM3MTF9.22yX2WEfDz-bbDS54-uGydc9IforleuC_7LMLxJbRok")
	//fmt.Println(middlewares.CreateToken_t(1, "测试", uuid.NewV4()))
	c := middlewares.CreateToken_t(1, "测试", uuid.NewV4())
	fmt.Println(c)
	//time.Sleep(5 * time.Second)
	//fmt.Println(a.RefreshToken(c))
	//time.Sleep(10 * time.Second)
	//fmt.Println(a.RefreshToken(c))
	//fmt.Println(a.RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiTmlja05hbWUiOiLmtYvor5UiLCJVVUlEIjoiMjRlMjgxOTMtMDIyYi00ZTM2LTk4ZmEtNTY0MDI0NzExOGFjIiwiZXhwIjoxNjk0MjY4NTExLCJpYXQiOjE2OTM2NjM3MTEsImlzcyI6InRlc3QiLCJuYmYiOjE2OTM2NjM3MTF9.22yX2WEfDz-bbDS54-uGydc9IforleuC_7LMLxJbRok"))
}

type S struct {
	C string
	F string
	O int
}

func TestTwo(t *testing.T) {
	var f S
	fmt.Println(&f, &f.C, &f.F, &f.O)

}
