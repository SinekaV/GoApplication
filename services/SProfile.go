package services

import (
	"context"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileService struct{
	//profile collection,context
	ProfileCollection *mongo.Collection
	ctx context.Context
}

func(p *ProfileService)CreateProfile(profile *models.Profile){}
func(p *ProfileService)EditProfile(profile *models.Profile){}
func(p *ProfileService)DeleteProfile(profileId int){}
func(p *ProfileService)GetProfileByIdProfile(profile int){}
func(p *ProfileService)GetProfileBySearch(){}



