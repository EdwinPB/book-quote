package internal

import (
	"embed"

	"github.com/book-quote/internal/home"
	"github.com/book-quote/public"

	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/db"
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
	"github.com/paganotoni/tailo"

	_ "github.com/lib/pq"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	// Assets is the manager for the public assets
	// it allows to watch for changes and reload the assets
	// when changes are made.
	Assets = assets.NewManager(public.Files)

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(Assets.Watch),
		gloves.WatchExtension(".go", ".css", ".js"),
	}

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envor.Get("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/book_help?sslmode=disable")

	// SessionSecret is the secret used to sign the session cookies.
	SessionSecret = envor.Get("SESSION_SECRET", "d720c059-9664-4980-8169-1158e167ae57")
	SessionName   = envor.Get("SESSION_NAME", "book_help_session")

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL)
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r server.Router) error {

	// LeapKit Middleware
	r.Use(session.Middleware(
		envor.Get("SESSION_SECRET", SessionSecret),
		envor.Get("SESSION_NAME", SessionName),
	))

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(render.AllHelpers),
		render.WithHelpers(map[string]any{
			"assetPath": Assets.PathFor,
		}),
	))

	r.HandleFunc("GET /{$}", home.Index)
	r.HandleFunc("GET /clicked/{$}", home.Clicked)

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}
