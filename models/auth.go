package models

import (
	"errors"

	"github.com/google/uuid"
)

type Role string

var (
	RoleTeacher  Role = "teacher"
	RoleStudent  Role = "student"
	RolePersonal Role = "personal"
	RoleAdmin    Role = "admin"
)

type Permission string
type Permissions []Permission

var (
	PermissionAdmin       Permission = "admin"
	PermissionEnrollments Permission = "enrollments"
	PermissionRegistry    Permission = "registry"
	PermissionSchedule    Permission = "schedule"
)

type PreferenceGroup string

var (
	PreferenceGroupWebsite  PreferenceGroup = "website"
	PreferenceGroupSchedule PreferenceGroup = "schedule"
	PreferenceGroupJournal  PreferenceGroup = "journal"
)

var PreferenceGroupMap = map[PreferenceGroup]bool{
	PreferenceGroupWebsite:  true,
	PreferenceGroupSchedule: true,
	PreferenceGroupJournal:  true,
}

func (g PreferenceGroup) IsValid() bool {
	return PreferenceGroupMap[g]
}

type User struct {
	Token string

	ID            uuid.UUID
	Login         string
	PictureUrl    string
	Email         string
	VerifiedEmail bool
}

type EnrollmentUser struct {
	Token string

	ID            uuid.UUID
	Login         string
	PictureUrl    string
	Email         string
	VerifiedEmail bool
	Enrollment    Enrollment
}

type Enrollment struct {
	Token string

	ID           uuid.UUID
	UserId       uuid.UUID
	StudyPlaceId uuid.UUID
	UserName     string
	Role         Role
	TypeId       uuid.UUID
	Permissions  Permissions
	Accepted     bool
	Blocked      bool
}

func (e Permissions) Assert(permissions ...Permission) error {
	for _, permission := range permissions {
		found := false

		for _, p := range e {
			if p == PermissionAdmin {
				return nil
			}

			if p == permission {
				found = true
				break
			}
		}

		if !found {
			return errors.New("no permission")
		}
	}

	return nil
}
