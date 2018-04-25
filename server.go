package main

import (
	gcontext "pinjihui.com/pinjihui/context"
	h "pinjihui.com/pinjihui/handler"
	"pinjihui.com/pinjihui/resolver"
	"pinjihui.com/pinjihui/schema"
	"pinjihui.com/pinjihui/repository"
	"log"
	"net/http"

	"pinjihui.com/pinjihui/loader"
	graphql "github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
	"pinjihui.com/pinjihui/service"
)

func main() {
	config := gcontext.LoadConfig(".")

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	ctx := context.Background()
	log := service.NewLogger(config)
	roleRepository := repository.NewRoleRepository(db, log)
	userRepository := repository.NewUserRepository(db, roleRepository, log)
	authService := service.NewAuthService(config, log)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "roleRepository", roleRepository)
	ctx = context.WithValue(ctx, "userRepository", userRepository)
	ctx = context.WithValue(ctx, "authService", authService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
