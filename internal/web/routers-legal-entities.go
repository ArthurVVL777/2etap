package web

import (
	"context"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain"
	"github.com/krisch/crm-backend/internal/jwt"
	oapi "github.com/krisch/crm-backend/internal/web/olegalentities"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// initOpenAPILegalEntitiesRouters инициализирует маршруты для юридических лиц и банковских счетов.
func initOpenAPILegalEntitiesRouters(a *Web, e *echo.Echo) {
	logrus.WithField("route", "legal-entities").Info("Initializing routes for legal entities and bank accounts")

	// Middleware
	middlewares := []oapi.StrictMiddlewareFunc{
		ValidateStructMiddeware,
		AuthMiddeware(a.app, []string{}),
	}

	// Регистрация обработчиков
	handlers := oapi.NewStrictHandler(a, middlewares)
	oapi.RegisterHandlers(e, handlers)

	logrus.WithField("route", "legal-entities").Info("Routes for legal entities and bank accounts registered successfully")
}

// GetLegalEntities возвращает список всех юридических лиц.
func (a *Web) GetLegalEntities(ctx context.Context, _ oapi.GetLegalEntitiesRequestObject) (oapi.GetLegalEntitiesResponseObject, error) {
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
		e := entity // Создаем копию перед использованием
		response[i] = oapi.LegalEntityDTO{
			Uuid:      e.UUID,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
			UpdatedAt: &e.UpdatedAt,
			DeletedAt: e.DeletedAt,
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

// PutLegalEntitiesUUID обновляет юридическое лицо.
func (a *Web) PutLegalEntitiesUUID(ctx context.Context, request oapi.PutLegalEntitiesUUIDRequestObject) (oapi.PutLegalEntitiesUUIDResponseObject, error) {
	defer Span(NewSpan(ctx, "PutLegalEntitiesUUID"))()

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

	return oapi.PutLegalEntitiesUUID204Response{}, nil
}

// DeleteLegalEntitiesUUID удаляет юридическое лицо.
func (a *Web) DeleteLegalEntitiesUUID(ctx context.Context, request oapi.DeleteLegalEntitiesUUIDRequestObject) (oapi.DeleteLegalEntitiesUUIDResponseObject, error) {
	defer Span(NewSpan(ctx, "DeleteLegalEntitiesUUID"))()

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

	return oapi.DeleteLegalEntitiesUUID204Response{}, nil
}





// GetAllBankAccounts возвращает список всех банковских счетов юридического лица.
func (a *Web) GetAllBankAccounts(ctx context.Context, request oapi.GetAllBankAccountsRequestObject) (oapi.GetAllBankAccountsResponseObject, error) {
	accounts, err := a.app.LegalEntitiesService.GetAllBankAccounts(ctx, request.Params.LegalEntityId)
	if err != nil {
		return nil, err
	}

	return oapi.GetAllBankAccounts200JSONResponse(accounts), nil
}

// PostBankAccount создает новый банковский счет
func (a *Web) PostBankAccount(ctx context.Context, request oapi.PostBankAccountRequestObject) (oapi.PostBankAccountResponseObject, error) {
	bankAccount := domain.BankAccount{
		ID:            uuid.New(),
		LegalEntityID: request.Body.LegalEntityId, // ✅ Исправлено
		BIK:           *request.Body.Bik,
		BankName:      *request.Body.BankName,
		Address:       *request.Body.Address,
		CorrAccount:   request.Body.CorrAccount,
		AccountNumber: request.Body.AccountNumber,
		Currency:      *request.Body.Currency,
		Comment:       *request.Body.Comment,
	}

	account, err := a.app.LegalEntitiesService.CreateBankAccount(ctx, bankAccount)
	if err != nil {
		return nil, err
	}

	return oapi.PostBankAccount201JSONResponse(account), nil
}

// PutBankAccount обновляет банковский счет
func (a *Web) PutBankAccount(ctx context.Context, request oapi.PutBankAccountRequestObject) (oapi.PutBankAccountResponseObject, error) {
	bankAccount := domain.BankAccount{
		ID:            request.Body.Id,            // ✅ Используем ID из тела запроса
		LegalEntityID: request.Body.LegalEntityId, // ✅ Исправлено
		BIK:           *request.Body.Bik,
		BankName:      *request.Body.BankName,
		Address:       *request.Body.Address,
		CorrAccount:   *request.Body.CorrAccount,
		AccountNumber: *request.Body.AccountNumber,
		Currency:      *request.Body.Currency,
		Comment:       *request.Body.Comment,
	}

	err := a.app.LegalEntitiesService.UpdateBankAccount(ctx, bankAccount)
	if err != nil {
		return nil, err
	}

	return oapi.PutBankAccount204Response{}, nil
}

// DeleteBankAccount удаляет банковский счет
func (a *Web) DeleteBankAccount(ctx context.Context, request oapi.DeleteBankAccountRequestObject) (oapi.DeleteBankAccountResponseObject, error) {
	err := a.app.LegalEntitiesService.DeleteBankAccount(ctx, request.Uuid)
	if err != nil {
		return nil, err
	}

	return oapi.DeleteBankAccount204Response{}, nil
}