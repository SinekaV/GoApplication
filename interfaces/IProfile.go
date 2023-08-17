// 5 method Signatures.This interface will provide the required methods
// Create Profile--EditOne profile--EditMany--Delete profile(smooth deletion)--Fetch profile
package interfaces

import "rest-api/models"

type IProfile interface{
	CreateProfile(profile *models.Profile)
	EditProfile(profile *models.Profile)
	DeleteProfile(profileId int)
	GetProfileById(profile int)
	GetProfileBySearch()
}