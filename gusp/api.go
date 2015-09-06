package gusp

import (
    "net/http"
    "net/http/cookiejar"
    "net/url"

    "github.com/PuerkitoBio/goquery"

    "github.com/snogaraleal/gu/gubase"
)

const LoginUrl = "https://login.it.gu.se/login"
const LoginCookie = "CAS3TGC"

const (
    PrefToken = "sp.token"
)


/*
 * Session
 */

type Session struct {
    Client *http.Client
}

type Utility interface {
    Dispatch(session *Session) interface{}
}

func NewSession() *Session {
    jar, err := cookiejar.New(nil)
    if err != nil {
        panic(err)
    }

    client := &http.Client{
        CheckRedirect: nil,
        Jar: jar,
    }

    return &Session{
        Client: client,
    }
}

func (session *Session) Request(utility Utility) interface{} {
    return utility.Dispatch(session)
}


/*
 * Auth
 */

type AuthUtility struct {
    Username string
    Password string
}

func (utility *AuthUtility) Dispatch(session *Session) interface{} {
    // Request login page
    resp, _ := session.Client.Get(LoginUrl)

    // Extract initial form values
    doc, _ := goquery.NewDocumentFromResponse(resp)

    formValues := url.Values{}

    form := doc.Find("form#fm1").First()
    form.Find("input").Each(func(i int, s *goquery.Selection) {
        name, _ := s.Attr("name")
        value, _ := s.Attr("value")

        formValues.Set(name, value)
    })

    // Set credentials
    formValues.Set("username", utility.Username)
    formValues.Set("password", utility.Password)

    // Attempt login
    resp, _ = session.Client.PostForm(LoginUrl, formValues)

    // Get token from cookies
    var token string

    parsedUrl, _ := url.Parse(LoginUrl)
    cookies := session.Client.Jar.Cookies(parsedUrl)
    for _, cookie := range cookies {
        if cookie.Name == LoginCookie {
            token = cookie.Value
        }
    }

    return AuthResult{token, len(token) > 0}
}


// AuthResult

type AuthResult struct {
    Token string
    Success bool
}

func (result *AuthResult) SyncPrefs() {
    if result.Success {
        // Store token
        gubase.SetPref(PrefToken, result.Token)
    } else {
        // Clean token
        gubase.SetPref(PrefToken, nil)
    }
}
