package database

import (
	"context"
	"ecolink-core/internal/models"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreDB struct {
	client *firestore.Client
	ctx    context.Context
}

func NewFirestoreDB(projectID, credentialsPath string) (*FirestoreDB, error) {
	ctx := context.Background()

	var client *firestore.Client
	var err error

	if credentialsPath != "" {
		client, err = firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsPath))
	} else {
		client, err = firestore.NewClient(ctx, projectID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %v", err)
	}

	return &FirestoreDB{
		client: client,
		ctx:    ctx,
	}, nil
}

func (db *FirestoreDB) SaveLink(link *models.Link) error {
	_, err := db.client.Collection("links").Doc(link.Code).Set(db.ctx, map[string]interface{}{
		"url":        link.URL,
		"code":       link.Code,
		"user_id":    link.UserID,
		"clicks":     link.Clicks,
		"created_at": link.CreatedAt,
		"updated_at": time.Now(),
	})
	return err
}

func (db *FirestoreDB) GetLink(code string) (*models.Link, error) {
	doc, err := db.client.Collection("links").Doc(code).Get(db.ctx)
	if err != nil {
		return nil, err
	}

	data := doc.Data()
	link := &models.Link{
		URL:       data["url"].(string),
		Code:      data["code"].(string),
		UserID:    data["user_id"].(string),
		Clicks:    int(data["clicks"].(int64)),
		CreatedAt: data["created_at"].(time.Time),
	}

	return link, nil
}

func (db *FirestoreDB) GetUserLinks(userID string) ([]*models.Link, error) {
	iter := db.client.Collection("links").Where("user_id", "==", userID).OrderBy("created_at", firestore.Desc).Documents(db.ctx)

	var links []*models.Link
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		data := doc.Data()
		link := &models.Link{
			URL:       data["url"].(string),
			Code:      data["code"].(string),
			UserID:    data["user_id"].(string),
			Clicks:    int(data["clicks"].(int64)),
			CreatedAt: data["created_at"].(time.Time),
		}
		links = append(links, link)
	}

	return links, nil
}

func (db *FirestoreDB) IncrementClicks(code string) error {
	_, err := db.client.Collection("links").Doc(code).Update(db.ctx, []firestore.Update{
		{Path: "clicks", Value: firestore.Increment(1)},
		{Path: "updated_at", Value: time.Now()},
	})
	return err
}

func (db *FirestoreDB) DeleteLink(code string) error {
	_, err := db.client.Collection("links").Doc(code).Delete(db.ctx)
	return err
}

func (db *FirestoreDB) SaveUser(user *models.User) error {
	_, err := db.client.Collection("users").Doc(user.ID).Set(db.ctx, map[string]interface{}{
		"id":         user.ID,
		"google_id":  user.GoogleID,
		"name":       user.Name,
		"email":      user.Email,
		"picture":    user.Picture,
		"provider":   user.Provider,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
	return err
}

func (db *FirestoreDB) GetUser(id string) (*models.User, error) {
	doc, err := db.client.Collection("users").Doc(id).Get(db.ctx)
	if err != nil {
		return nil, err
	}

	data := doc.Data()
	user := &models.User{
		ID:        data["id"].(string),
		GoogleID:  data["google_id"].(string),
		Name:      data["name"].(string),
		Email:     data["email"].(string),
		Picture:   data["picture"].(string),
		Provider:  data["provider"].(string),
		CreatedAt: data["created_at"].(time.Time),
		UpdatedAt: data["updated_at"].(time.Time),
	}

	return user, nil
}

func (db *FirestoreDB) GetUserByGoogleID(googleID string) (*models.User, error) {
	iter := db.client.Collection("users").Where("google_id", "==", googleID).Limit(1).Documents(db.ctx)

	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}

	data := doc.Data()
	user := &models.User{
		ID:        data["id"].(string),
		GoogleID:  data["google_id"].(string),
		Name:      data["name"].(string),
		Email:     data["email"].(string),
		Picture:   data["picture"].(string),
		Provider:  data["provider"].(string),
		CreatedAt: data["created_at"].(time.Time),
		UpdatedAt: data["updated_at"].(time.Time),
	}

	return user, nil
}

func (db *FirestoreDB) Close() error {
	return db.client.Close()
}
