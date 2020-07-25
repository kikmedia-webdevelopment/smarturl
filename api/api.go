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
	e.HideBanner = true
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

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

	if enableAdmin {
		// setup admin ui
		t := &Template{
			templates: template.Must(template.ParseGlob("templates/*.html")),
		}
		e.Renderer = t

		e.Static("/static", buildPath+"/static")
		g := e.Group("/admin")
		g.GET("*", func(c echo.Context) error {
			tmpl, err := template.ParseFiles(path.Join("templates", "index.html"))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing Template")
			}

			data, err := NewViewData(buildPath)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing ViewData")
			}

			res := c.Response().Writer
			if err := tmpl.Execute(res, data); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error serving Template")
			}
			return nil
		})

		a := e.Group("/api")
		a.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(config.Secret),
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
	}

	e.GET("/:id", api.getEntry)

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

	// Start server
	e.Logger.Fatal(e.Start(listenAddr))
}

func (a *API) getEntry(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusNotFound, "id not found")
	}

	entry, err := a.store.GetEntryAndIncrease(id)
	if err != nil {
		logrus.Error(err)
		return c.String(http.StatusNotFound, "id not found")
	}

	if entry.URL != "" {
		return c.Redirect(http.StatusMovedPermanently, entry.URL)
	}

	return c.String(http.StatusNotFound, "id not found")
}
