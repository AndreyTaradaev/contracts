// маршрутизатор для http сервера API приложения Gateway
package api

import (
	"encoding/json"
	"fmt"
	"strings"

	//"gateway/censor/pkg/storage"
	db "contracts/censor/pkg/storage/oracle"
	logs "contracts/internal/log"
	tools "contracts/internal/tools"

	//"io"
	"net/http"
	"time"

	_ "contracts/censor/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	r *mux.Router
}

const timeout int = 120

// Конструктор API.
func New() (*API, error) {
	logger := logs.New()
	a := API{r: mux.NewRouter()}
	a.endpoints()
	logger.Debug("Init router http")
	return &a, nil
}

// Router возвращает маршрутизатор для использования
// в качестве аргумента HTTP-сервера.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.Use(api.headersMiddleware)
	api.r.HandleFunc("/contracts", api.contracts).Methods(http.MethodGet, http.MethodOptions)
	api.r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

// page godoc
// @Summary  Get list contracts
// @Description  Список договоров  из Биллинга
//
//	@ID		List
//
// @Tags List
// @Produce  json
// @Param  date  query int false  "начальная дата в формате YYYYMMDDHHMM (202301010101)"
// @Success 200	{array} 	model.Contract 	"Список договоров"
//
//	@Failure		500	{string}	string 	"внутренняя ошибка сервера"
//
// @Router /contracts [get]
func (api *API) contracts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	strdate := r.URL.Query().Get("date") //слово поиска
	logs.New().Debugf("argument %s", strdate)
	intdate := tools.GetIntDef(strdate, 0)
	logs.New().Debugf("argument  convert %d", intdate)
	array, err := db.GetContacts(intdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ret, err := json.Marshal(array.Get())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Endpoint", "Contracts")
	w.Header().Set("count", fmt.Sprintf("%d", array.Len()))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(ret)))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(ret)
}

func (api *API) headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.String(), "swagger") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			//logs.New().Debug("checked Path for current dir ", os.Getenv("PWD"))
		}

		reqID := r.URL.Query().Get("request_id")
		if len(reqID) == 0 {
			u := uuid.New()
			reqID = u.String()
		}
		logs.New().Debug("Header:", r.Header)
		logs.New().Debug("Url: ", r.URL)
		next.ServeHTTP(w, r)
		w.Header().Set("Request-ID", reqID)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		logstr := fmt.Sprintf("Date: %s, Host: %s, Metod: %s, Path %s,  ", time.Now().Format("02-Jan-2006 15:04:05.00"), r.RemoteAddr, r.Method, r.URL.Path)
		logs.New().Info(logstr)
	})
}
