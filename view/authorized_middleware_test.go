package view

import (
	"context"
	"isjhar/template/echo-golang/data/repositories"
	"isjhar/template/echo-golang/domain/entities"
	usecases "isjhar/template/echo-golang/domain/use-cases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
)

func TestAuthorizedUser(t *testing.T) {
	ctx := context.Background()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := AuthorizedUser("header")(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	generatePairTokenUseCase := usecases.GeneratePairTokenUseCase{
		JwtRepository: repositories.JwtRepository{},
	}
	pairToken, err := generatePairTokenUseCase.Execute(ctx, entities.User{
		ID:       1,
		Username: "admin",
		Name:     null.StringFrom("admin"),
	})
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderAuthorization, "bearer "+pairToken.AccessToken)

	h(c)

	if http.StatusOK != rec.Code {
		t.Fatalf("Call error")
	}
}
