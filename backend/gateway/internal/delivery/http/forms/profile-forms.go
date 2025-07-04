package forms

import (
	"errors"
	"time"

	"github.com/google/uuid"

	time2 "quickflow/config/time"
	"quickflow/shared/models"
)

//easyjson:json
type ProfileInfo struct {
	Username      string     `json:"username,omitempty"`
	Name          string     `json:"firstname"`
	Surname       string     `json:"lastname"`
	Sex           models.Sex `json:"sex"`
	DateOfBirth   string     `json:"birth_date"`
	Bio           string     `json:"bio"`
	AvatarUrl     string     `json:"avatar_url,omitempty"`
	BackgroundUrl string     `json:"cover_url,omitempty"`
}

//easyjson:json
type ProfileForm struct {
	Id         string       `json:"id,omitempty"`
	Avatar     *models.File `json:"-"`
	Background *models.File `json:"-"`

	ProfileInfo         *ProfileInfo             `json:"profile"`
	ContactInfo         *ContactInfo             `json:"contact_info,omitempty"`
	SchoolEducation     *SchoolEducationForm     `json:"school,omitempty"`
	UniversityEducation *UniversityEducationForm `json:"university,omitempty"`
	LastSeen            string                   `json:"last_seen,omitempty"`
	IsOnline            *bool                    `json:"online,omitempty"`
	Relation            models.UserRelation      `json:"relation,omitempty"`
	ChatId              *uuid.UUID               `json:"chat_id,omitempty"`
}

func (f *ProfileForm) FormToModel() (models.Profile, error) {

	var contactInfo *models.ContactInfo
	if f.ContactInfo != nil {
		contactInfo = &models.ContactInfo{
			City:  f.ContactInfo.City,
			Email: f.ContactInfo.Email,
			Phone: f.ContactInfo.Phone,
		}
	}

	var basicInfo *models.BasicInfo
	var err error
	if f.ProfileInfo != nil {
		basicInfo, err = ProfileInfoToModel(*f.ProfileInfo)
		if err != nil {
			return models.Profile{}, err
		}
	}

	p := models.Profile{
		BasicInfo:  basicInfo,
		Avatar:     f.Avatar,
		Background: f.Background,

		SchoolEducation:     SchoolFormToModel(f.SchoolEducation),
		UniversityEducation: UniversityFormToModel(f.UniversityEducation),
		ContactInfo:         contactInfo,
	}
	if f.ProfileInfo != nil {
		p.Username = f.ProfileInfo.Username
	}

	return p, nil
}

func ModelToForm(profile models.Profile, username string, isOnline bool, relation models.UserRelation, uuid *uuid.UUID) ProfileForm {
	profileForm := ProfileForm{
		Id:                  profile.UserId.String(),
		ProfileInfo:         BasicInfoToForm(*profile.BasicInfo, username),
		SchoolEducation:     SchoolEducationToForm(profile.SchoolEducation),
		UniversityEducation: UniversityEducationToForm(profile.UniversityEducation),
		ContactInfo:         ContactInfoToForm(profile.ContactInfo),
		IsOnline:            &isOnline,
		Relation:            relation,
		ChatId:              uuid,
	}
	if !isOnline {
		profileForm.LastSeen = profile.LastSeen.Format(time2.TimeStampLayout)
	}
	return profileForm
}

//easyjson:json
type ContactInfo struct {
	City  string `json:"city,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func ContactInfoToForm(contactInfo *models.ContactInfo) *ContactInfo {
	if contactInfo == nil {
		return nil
	}

	return &ContactInfo{
		City:  contactInfo.City,
		Email: contactInfo.Email,
		Phone: contactInfo.Phone,
	}
}

func ContactInfoFormToModel(contactInfo *ContactInfo) *models.ContactInfo {
	if contactInfo == nil {
		return nil
	}

	return &models.ContactInfo{
		City:  contactInfo.City,
		Email: contactInfo.Email,
		Phone: contactInfo.Phone,
	}
}

//easyjson:json
type SchoolEducationForm struct {
	SchoolCity string `json:"school_city,omitempty"`
	SchoolName string `json:"school_name,omitempty"`
}

//easyjson:json
type UniversityEducationForm struct {
	UniversityCity    string `json:"univ_city,omitempty"`
	UniversityName    string `json:"univ_name,omitempty"`
	UniversityFaculty string `json:"faculty,omitempty"`
	GraduationYear    int    `json:"grad_year,omitempty"`
}

//easyjson:json
type Activity struct {
	LastSeen string `json:"last_seen,omitempty"`
	IsOnline bool   `json:"online,omitempty"`
}

func SchoolEducationToForm(sch *models.SchoolEducation) *SchoolEducationForm {
	if sch == nil {
		return nil
	}

	return &SchoolEducationForm{
		SchoolCity: sch.City,
		SchoolName: sch.School,
	}
}

func UniversityEducationToForm(uni *models.UniversityEducation) *UniversityEducationForm {
	if uni == nil {
		return nil
	}

	return &UniversityEducationForm{
		UniversityCity:    uni.City,
		UniversityName:    uni.University,
		UniversityFaculty: uni.Faculty,
		GraduationYear:    uni.GraduationYear,
	}
}

func SchoolFormToModel(sch *SchoolEducationForm) *models.SchoolEducation {
	if sch == nil {
		return nil
	}

	return &models.SchoolEducation{
		City:   sch.SchoolCity,
		School: sch.SchoolName,
	}
}

func UniversityFormToModel(uniForm *UniversityEducationForm) *models.UniversityEducation {
	if uniForm == nil {
		return nil
	}

	return &models.UniversityEducation{
		City:           uniForm.UniversityCity,
		University:     uniForm.UniversityName,
		Faculty:        uniForm.UniversityFaculty,
		GraduationYear: uniForm.GraduationYear,
	}
}

func BasicInfoToForm(info models.BasicInfo, username string) *ProfileInfo {
	return &ProfileInfo{
		Username:      username,
		Name:          info.Name,
		Surname:       info.Surname,
		Sex:           info.Sex,
		DateOfBirth:   info.DateOfBirth.Format(time2.DateLayout),
		Bio:           info.Bio,
		AvatarUrl:     info.AvatarUrl,
		BackgroundUrl: info.BackgroundUrl,
	}
}

func ProfileInfoToModel(info ProfileInfo) (*models.BasicInfo, error) {
	date, err := time.Parse(time2.DateLayout, info.DateOfBirth)
	if err != nil {
		return nil, errors.New("incorrect date format")
	}
	return &models.BasicInfo{
		Name:          info.Name,
		Surname:       info.Surname,
		Sex:           info.Sex,
		DateOfBirth:   date,
		Bio:           info.Bio,
		AvatarUrl:     info.AvatarUrl,
		BackgroundUrl: info.BackgroundUrl,
	}, nil
}
