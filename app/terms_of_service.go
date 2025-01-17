


package app

import (
	"za-white-screen/model"
)

func (a *App) CreateTermsOfService(text, userId string) (*model.TermsOfService, *model.AppError) {
	termsOfService := &model.TermsOfService{
		Text:   text,
		UserId: userId,
	}

	if _, err := a.GetUser(userId); err != nil {
		return nil, err
	}

	return a.Srv.Store.TermsOfService().Save(termsOfService)
}

func (a *App) GetLatestTermsOfService() (*model.TermsOfService, *model.AppError) {
	return a.Srv.Store.TermsOfService().GetLatest(true)
}

func (a *App) GetTermsOfService(id string) (*model.TermsOfService, *model.AppError) {
	return a.Srv.Store.TermsOfService().Get(id, true)
}
