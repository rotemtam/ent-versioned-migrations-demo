package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rotemtam/ent-versioned-migrations-demo/ent"
)

func main() {
	client, err := ent.Open(dialect.MySQL, "root:pass@tcp(localhost:3306)/db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := seed(context.Background(), client); err != nil {
		log.Fatalf("failed seeding db: %v", err)
	}
	for _, user := range client.User.Query().
		WithBlogPosts().
		AllX(context.Background()) {
		log.Printf("%q by %s", user.Edges.BlogPosts[0].Title, user.Name)
	}
}

func seed(ctx context.Context, client *ent.Client) error {
	for _, name := range []string{"rotemtam", "a8m", "jcl"} {
		client.User.Create().
			SetName(name).
			SetEmail(name + "@gmail.com").
			AddBlogPosts(
				client.Blog.Create().SetTitle(name + "'s first post").SetBody("content").SaveX(ctx),
			).
			SaveX(ctx)
	}
	return nil
}
