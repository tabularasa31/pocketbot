package server

import (
	"PocketBot/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server          *http.Server
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string
}

func NewAuthorizationServer(pocketClient *pocket.Client, tokenRepository repository.TokenRepository, redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{pocketClient: pocketClient, tokenRepository: tokenRepository, redirectURL: redirectURL}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":80",
		Handler: s,
	}
	log.Println("Авторизация сервера START")

	return s.server.ListenAndServe()
}

// Обработчик HTTP запросов
func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Если метод не GET, то выходим из обработчика
	if r.Method != http.MethodGet {
		log.Println("Method not GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Получаем чат айди
	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Конвертим чат айди в формат int64
	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Запрашивает реквест токен из нашей бд
	requestToken, err := s.tokenRepository.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Получаем аксесс токен
	authResp, err := s.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Записываем аксесс токен в нашу бд
	err = s.tokenRepository.Save(chatID, authResp.AccessToken, repository.AccessTokens)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Отдаем http запрос, в header которого добавлена инфа о нашей редирект ссылке
	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)

}
