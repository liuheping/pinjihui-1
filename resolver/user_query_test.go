package resolver

import (
	gcontext "pinjihui.com/pinjihui/context"
	"pinjihui.com/pinjihui/schema"
	"pinjihui.com/pinjihui/repository"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"golang.org/x/net/context"
	"log"
	"testing"
	"pinjihui.com/pinjihui/service"
)

var (
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})
	ctx        context.Context
)

func init() {
	config := gcontext.LoadConfig("../")
	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	log := service.NewLogger(config)
	roleRepository := repository.NewRoleRepository(db, log)
	userRepository := repository.NewUserRepository(db, roleRepository, log)
	ctx = context.WithValue(context.Background(), "userRepository", userRepository)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					user(email:"test@1.com") {
						id
						email
						password
					}
				}
			`,
			ExpectedResult: `
				{
					"user": {
					  "id": "1",
					  "email": "test@1.com",
					  "password": "********"
					}
				}
			`,
		},
	})
}
