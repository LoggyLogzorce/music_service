package user

import (
	"music_service/internal/context"
	"net/http"
	"time"
)

func SetHeaders(ctx *context.Context) {
	ctx.Response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response.Header().Set("Expires", time.Now().Format(http.TimeFormat))
	ctx.Response.Header().Set("Pragma", "no-cache")
}

type Handler struct {
}

func (*Handler) LoginPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/login.html")
	return
}

func (*Handler) RegisterPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/register.html")
	return
}

func (*Handler) HomePage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/index.html")
	return
}

func (*Handler) ChartsPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/charts.html")
	return
}

func (*Handler) Podcasts(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/podcasts.html")
	return
}

func (*Handler) MyCollectionPage(ctx *context.Context) {
	SetHeaders(ctx)
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/my_collection.html")
	return
}

func (*Handler) AboutPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/about.html")
	return
}

func (*Handler) VacanciesPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/vacancies.html")
	return
}

func (*Handler) PressPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/press.html")
	return
}

func (*Handler) ContactPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/contact.html")
	return
}

func (*Handler) ReferencePage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/reference.html")
	return
}

func (*Handler) TermsOfUsePage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/terms_of_use.html")
	return
}

func (*Handler) ConfidentialityPage(ctx *context.Context) {
	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/confidentiality.html")
	return
}
