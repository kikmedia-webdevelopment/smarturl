package api

import (
	"html/template"
	"net/http"
	"path"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/stores"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

type API struct {
	store  *stores.Store
	echo   *echo.Echo
	config *config.Configuration
}

// New starts a new Web-Service
func New(store *stores.Store, config *config.Configuration) {
	buildPath := path.Clean("ui/build")
	enableAdmin := true

	e := echo.New()

	if config.Web.UseTLS {
		e.AutoTLSManager.Cache = autocert.DirCache(".cache")
	}

	e.HideBanner = true
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"Authorization",
		},
	}))

	if config.Web.Debug {
		e.Debug = true
	}

	api := &API{
		store:  store,
		echo:   e,
		config: config,
	}

	e.GET("/:id", api.loadLink)

	if enableAdmin {
		// setup admin ui
		t := &Template{
			templates: template.Must(template.ParseGlob("templates/*.html")),
		}
		e.Renderer = t

		e.Static("/static", buildPath+"/static")
		e.GET("/admin", api.serveAdmin)
		e.GET("/admin/*", api.serveAdmin)

		//e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		//	Level: 5,
		//}))

		a := e.Group("/api")
		a.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.JWT.Secret),
			Skipper: func(c echo.Context) bool {
				// skip auth route
				if c.Path() == "/api/users/authenticate" {
					return true
				}
				if c.Path() == "/api/users/refresh" {
					return true
				}
				return false
			},
		}))
		a.POST("/users/authenticate", api.Login)
		a.POST("/users/refresh", api.RefreshToken)
		a.GET("/links", api.LinksList)
		a.PATCH("/links", api.LinkUpdate)
		a.POST("/links", api.LinkCreate)
		a.DELETE("/links", api.LinkDelete)

		a.GET("/stats", api.listStats)

		a.GET("/qrcode/:target", api.GetQRCode)
		a.POST("/qrcode/:target", api.GetQRCode)
	}

	var listenAddr string

	if config.Web.ListenAddr != "" {
		listenAddr = config.Web.ListenAddr
	} else {
		logrus.Info("missing ListenAddr in Config")
		listenAddr = ":1323"
	}

	if config.Web.Redirect != "" {
		e.Any("/", func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, config.Web.Redirect)
		})
	}

	if config.Web.UseTLS {
		logrus.Info("Starting TLS Server")
		go func(c *echo.Echo) {
			e.Logger.Fatal(e.StartAutoTLS(":443"))
		}(e)
	}

	// Start server
	e.Logger.Fatal(e.Start(listenAddr))
}
