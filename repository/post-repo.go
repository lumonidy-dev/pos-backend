package repository

import (
	"context"
	"log"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
)

// PostRepository es una interfaz que define las operaciones que se pueden realizar con un post
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

const (
	projectId      string = "lumo-pos"
	collectionName string = "posts"
)

// NewPostReposity retorna una nueva instancia de PostRepository
func NewPostReposity() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Error al crear el cliente de Firestore: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Body":  post.Body,
	})

	if err != nil {
		log.Fatalf("Error al añadir un documento: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Error al crear el cliente de Firestore: %v", err)
		return nil, err
	}

	defer client.Close()

	var post []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Falló iterando los documentos: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(string),
			Title: doc.Data()["Title"].(string),
			Body:  doc.Data()["Body"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
