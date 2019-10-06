package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

var (
	schema graphql.Schema
)

func SchemaInit() {

	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorials": &graphql.Field{

					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					Type: authorType,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial by its id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)

				if ok {

					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							return tutorial, nil
						}
					}
				}

				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	var err error
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err = graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Println(err.Error())
	}

	// 	query := `
	//     {
	//         list {
	//             id
	//             title
	//             comments {
	//                 body
	//             }
	//             author {
	//                 Name
	//                 Tutorials
	//             }
	//         }
	//     }
	// `
	// 	params := graphql.Params{Schema: schema, RequestString: query}
	// 	r := graphql.Do(params)
	// 	if len(r.Errors) > 0 {
	// 		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	// 	}
	// 	rJSON, _ := json.Marshal(r)
	// 	fmt.Printf("%s \n", rJSON)
}
