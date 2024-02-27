package grpc

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/auth"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
)

func Auth(ctx context.Context, token string) (models.User, error) {
	userRpc, err := clients.AuthGRpcClient.Auth(ctx, &auth.Token{Token: token})
	if err != nil {
		return models.User{}, err
	}

	userId, err := uuid.Parse(userRpc.Id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Token: token,

		ID:            userId,
		Login:         userRpc.Login,
		PictureUrl:    userRpc.PictureUrl,
		Email:         userRpc.Email,
		VerifiedEmail: userRpc.VerifiedEmail,
	}

	return user, nil
}

func EnrollmentAuth(ctx context.Context, token string, studyPlaceId string) (models.EnrollmentUser, error) {
	userRpc, err := clients.StudyPlacesGRpcClient.Auth(ctx, &studyplaces.EnrollmentToken{Token: token, StudyPlaceId: studyPlaceId})
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	userId, err := uuid.Parse(userRpc.Id)
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	enrollmentId, err := uuid.Parse(userRpc.Enrollment.Id)
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	enrollmentUserId, err := uuid.Parse(userRpc.Enrollment.UserId)
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	enrollmentStudyPlaceId, err := uuid.Parse(userRpc.Enrollment.StudyPlaceId)
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	enrollmentTypeId, err := uuid.Parse(userRpc.Enrollment.TypeId)
	if err != nil {
		return models.EnrollmentUser{}, err
	}

	permissions := make([]models.Permission, len(userRpc.Enrollment.Permissions))
	for i, rpcPermission := range userRpc.Enrollment.Permissions {
		permissions[i] = models.Permission(rpcPermission)
	}

	enrollment := models.Enrollment{
		Token:        token,
		ID:           enrollmentId,
		UserId:       enrollmentUserId,
		StudyPlaceId: enrollmentStudyPlaceId,
		UserName:     userRpc.Enrollment.UserName,
		Role:         models.Role(strings.ToLower(userRpc.Enrollment.Role.String())),
		TypeId:       enrollmentTypeId,
		Permissions:  models.Permissions(permissions),
		Accepted:     userRpc.Enrollment.Accepted,
		Blocked:      userRpc.Enrollment.Blocked,
	}

	user := models.EnrollmentUser{
		Token:         token,
		ID:            userId,
		Login:         userRpc.Login,
		PictureUrl:    userRpc.PictureUrl,
		Email:         userRpc.Email,
		VerifiedEmail: userRpc.VerifiedEmail,
		Enrollment:    enrollment,
	}

	return user, nil
}
