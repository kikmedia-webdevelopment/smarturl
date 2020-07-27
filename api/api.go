package api

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/juliankoehn/mchurl/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

type API struct {
	db     *gorm.DB
	echo   *echo.Echo
	config *config.Configuration
}

// New starts a new Web-Service
func New(db *gorm.DB, config *config.Configuration) *API {
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
		echo:   e,
		config: config,
		db:     db,
	}

	e.GET("/:id", api.loadLink)

	if enableAdmin {
		var templatePath string
		// used for testing
		if strings.HasSuffix(os.Args[0], ".test") {
			templatePath = "../templates/*.html"
		} else {
			templatePath = "templates/*.html"
		}
		// setup admin ui
		t := &Template{
			templates: template.Must(template.ParseGlob(templatePath)),
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

	}
	e.GET("/qrcode/:target/:dimension/:level", api.GetQRCode)

	if config.Web.Redirect != "" {
		e.Any("/", func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, config.Web.Redirect)
		})
	}

	return api
}

// Starts the API Server
func (a *API) Start() {
	var listenAddr string

	if a.config.Web.ListenAddr != "" {
		listenAddr = a.config.Web.ListenAddr
	} else {
		logrus.Info("missing ListenAddr in Config")
		listenAddr = ":1323"
	}

	if a.config.Web.UseTLS {
		logrus.Info("Starting TLS Server")
		go func(c *echo.Echo) {
			a.echo.Logger.Fatal(a.echo.StartAutoTLS(":443"))
		}(a.echo)
	}

	a.echo.Logger.Fatal(a.echo.Start(listenAddr))
}
