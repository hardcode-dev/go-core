// Package main - демонстрация аутентификации пользователя при помощи Яндекс.OAuth.
// Для использования внешнего поставщика аутентификации OAuth нужно предварительно
// зарегистрировать приложение у поставщика и указать запрашиваемые разрешения.
// Наше приложение запрашивает имя, адрес почты и дату рождения пользователя.
// Регистрация занимает несколько секунд и осуществляется по ссылке "https://oauth.yandex.ru/client/new".
// После успешной регистрации наше приложение получило уникальный идентификатор и пароль доступа.
package main

// Импорт пакетов. Пакеты не из стандартной библиотеки загружаются с помощью команды "go get".
import (
	"encoding/json" // стандартная библиотека (пакеты установлены вместе с языком)
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2" // внешний пакет, будет загружен из Интернета
)

// Пользовательский тип данных - структура.
type user struct {
	ID           string `json:"client_id"`
	Login        string `json:"login"`
	First        string `json:"first_name"`
	Last         string `json:"last_name"`
	RealName     string `json:"real_name"`
	DefaultEmail string `json:"default_email"`
}

// Глобальные переменные (не лучшая практика, но для небольшого приложения можно использовать).
var (
	// Настройки подключения к Яндекс.OAuth.
	oauthStateString  = "pseudo-random"
	yandexOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",   // ссылка для перехода при успешной аутентификации
		ClientID:     "da31dbdf3e9b43479fb27462b02778e8", // ID Яндекс.OAuth
		ClientSecret: "39c7981defca427f855e000d72621dee", // пароль Яндекс.OAuth
		Endpoint:     endpoint,                           // адреса серверов Яндекс.OAuth
	}
	endpoint = oauth2.Endpoint{
		AuthURL:  "https://oauth.yandex.ru/authorize",
		TokenURL: "https://oauth.yandex.ru/token",
	}
)

// Функция main() - точка входа в программу (как во многих Си-подобных языках).
func main() {
	// Сопоставление ссылок обработчикам.
	// При переходе на http://localhost:8080/URL будет запущена функция, сопоставленная с URL.
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleYandexLogin)
	http.HandleFunc("/callback", handleCallback)
	// Запуск высокопроизводительного многопользовательского HTTP-сервера одной строкой!
	http.ListenAndServe(":8080", nil)
}

// Обработчик ссылки "/" - показывает ссылку с переходом на страницу Яндекс.OAuth
// и перенаправляет на "/login"
func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlPage = `
	<html>
		<body>
			<a href="/login">Вход с помощью Яндекс</a>
		</body>
	</html>
	`
	fmt.Fprintf(w, htmlPage)
}

// Обработчик "/login" - перенаправляет на сайт Яндекс.OAuth.
func handleYandexLogin(w http.ResponseWriter, r *http.Request) {
	url := yandexOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Обработчик, вызываемый при успешной аутентификации.
// Печатает на странице JSON с данными профиля пользователя.
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Параметры для получения токена переданы в URL при переадресации от Яндекса.
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

// Функция получает токен доступа, на основе полученной на предыдущем этапе ссылки.
// Используя токен доступа, функция запрашивает у сервера Яндекс.OAuth данные пользователя.
// Сервер присылает JSON с разрешёнными для приложения данными.
// Функция выводит результат в консоль и возвращает данные выше по стеку.
func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := yandexOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://login.yandex.ru/info?format=json", nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Authorization", "OAuth "+token.AccessToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var u user

	err = json.Unmarshal([]byte(contents), &u)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %s", err.Error())
	}

	log.Printf("User Info: %+v\n", u)

	return contents, nil
}
