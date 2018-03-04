package graphqlctl

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sj14/web-demo/interfaces/web/controller/mainctl"
)

func NewGraphQLController(mainController mainctl.MainController) GraphQLController {
	return GraphQLController{mainController}
}

type GraphQLController struct {
	mainctl.MainController
}

func (interactor *GraphQLController) queryType() *graphql.Object {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"hello": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "World", nil

					},
				},
			},
		})

	return queryType
}

func (interactor *GraphQLController) Schema() graphql.Schema {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: interactor.queryType(),
		},
	)
	return schema
}

func (interactor *GraphQLController) ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
