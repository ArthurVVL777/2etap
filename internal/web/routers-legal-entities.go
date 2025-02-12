package web

import (
	"context"

	"github.com/krisch/crm-backend/internal/jwt"
	oapi "github.com/krisch/crm-backend/internal/web/olegalentities"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// initOpenAPILegalEntitiesRouters инициализирует маршруты для юридических лиц.
func initOpenAPILegalEntitiesRouters(a *Web, e *echo.Echo) {
	logrus.WithField("route", "legal-entities").Info("Initializing routes for legal entities")

	// Middleware
	middlewares := []oapi.StrictMiddlewareFunc{
		ValidateStructMiddeware,
		AuthMiddeware(a.app, []string{}),
	}

	// Регистрация обработчиков
	handlers := oapi.NewStrictHandler(a, middlewares)
	oapi.RegisterHandlers(e, handlers)

	logrus.WithField("route", "legal-entities").Info("Routes for legal entities registered successfully")
}

// GetLegalEntities возвращает список всех юридических лиц.
func (a *Web) GetLegalEntities(ctx context.Context, request oapi.GetLegalEntitiesRequestObject) (oapi.GetLegalEntitiesResponseObject, error) {
	defer Span(NewSpan(ctx, "GetLegalEntities"))()

	claims, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	logrus.WithField("user_email", claims.Email).Info("Получение списка юридических лиц")

	entities, err := a.app.LegalEntitiesService.GetAllLegalEntities(ctx)
	if err != nil {
		logrus.Errorf("Ошибка получения юридических лиц: %v", err)
		return nil, err
	}

	response := make([]oapi.LegalEntityDTO, len(entities))
	for i, entity := range entities {
		response[i] = oapi.LegalEntityDTO{
			Uuid:      entity.UUID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: &entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
		}
	}

	return oapi.GetLegalEntities200JSONResponse(response), nil
}

// PostLegalEntities создаёт новое юридическое лицо.
func (a *Web) PostLegalEntities(ctx context.Context, request oapi.PostLegalEntitiesRequestObject) (oapi.PostLegalEntitiesResponseObject, error) {
	defer Span(NewSpan(ctx, "PostLegalEntities"))()

	claims, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	logrus.WithField("user_email", claims.Email).Info("Создание юридического лица")

	entity, err := a.app.LegalEntitiesService.CreateLegalEntity(ctx, request.Body.Name)
	if err != nil {
		logrus.Errorf("Ошибка создания юридического лица: %v", err)
		return nil, err
	}

	return oapi.PostLegalEntities201JSONResponse{
		Uuid: entity.UUID,
	}, nil
}

// PutLegalEntitiesUuid обновляет юридическое лицо.
func (a *Web) PutLegalEntitiesUuid(ctx context.Context, request oapi.PutLegalEntitiesUuidRequestObject) (oapi.PutLegalEntitiesUuidResponseObject, error) {
	defer Span(NewSpan(ctx, "PutLegalEntitiesUuid"))()

	claims, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	logrus.WithField("user_email", claims.Email).Info("Обновление юридического лица")

	err := a.app.LegalEntitiesService.UpdateLegalEntity(ctx, request.Uuid, request.Body.Name)
	if err != nil {
		logrus.Errorf("Ошибка обновления юридического лица: %v", err)
		return nil, err
	}

	return oapi.PutLegalEntitiesUuid204Response{}, nil
}

// DeleteLegalEntitiesUuid удаляет юридическое лицо.
func (a *Web) DeleteLegalEntitiesUuid(ctx context.Context, request oapi.DeleteLegalEntitiesUuidRequestObject) (oapi.DeleteLegalEntitiesUuidResponseObject, error) {
	defer Span(NewSpan(ctx, "DeleteLegalEntitiesUuid"))()

	claims, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	logrus.WithField("user_email", claims.Email).Info("Удаление юридического лица")

	err := a.app.LegalEntitiesService.DeleteLegalEntity(ctx, request.Uuid)
	if err != nil {
		logrus.Errorf("Ошибка удаления юридического лица: %v", err)
		return nil, err
	}

	return oapi.DeleteLegalEntitiesUuid204Response{}, nil
}
