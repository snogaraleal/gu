package gusp

import (
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"

    "github.com/PuerkitoBio/goquery"

    "github.com/snogaraleal/gu/gubase"
)

const (
    LanguageSv = "100000"
    LanguageEn = "100001"
)

const (
    LoginUrl = "https://login.it.gu.se/login"
    LoginCookie = "CAS3TGC"
)

const (
    SyllabusGetUrl =
        "http://studentportal.gu.se/english/e-services/" +
        "course-syllabus/?languageId=" + LanguageEn

    SyllabusPostUrl =
        "http://studentportal.gu.se/english/e-services/" +
        "course-syllabus/syllabisearchresultview/"
)

const (
    PrefToken = "sp.token"
)

var DefaultHeaders = map[string]string {
    "User-Agent": "gucli/1.0",
    "Content-Type": "application/x-www-form-urlencoded",
}


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
    // Extract initial form values
    resp, _ := session.Client.Get(LoginUrl)
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


/*
 * Syllabus
 */

type SyllabusUtility struct {
    Query string
}

func (utility *SyllabusUtility) Dispatch(session *Session) interface{} {
    // Extract initial form values
    resp, _ := session.Client.Get(SyllabusGetUrl)
    doc, _ := goquery.NewDocumentFromResponse(resp)

    formValues := url.Values{}

    form := doc.Find("form#courseSearchForm").First()
    form.Find("input").Each(func(i int, s *goquery.Selection) {
        name, _ := s.Attr("name")
        value, _ := s.Attr("value")

        formValues.Set(name, value)
    })

    // Set query
    formValues.Set("courseQuery", utility.Query)
    formValues.Set("languageId", LanguageEn)
    formValues.Set("sLang", "en")

    // Make search request
    req, _ := http.NewRequest(
        "POST", SyllabusPostUrl, strings.NewReader(formValues.Encode()))
    for headerName, headerValue := range DefaultHeaders {
        req.Header.Add(headerName, headerValue)
    }
    resp, _ = session.Client.Do(req)
    doc, _ = goquery.NewDocumentFromResponse(resp)

    // Extract results
    message := doc.Find("#showingNumberOf").Text()

    tbody := doc.Find("#courseSyllabiSearchResultTableBody").First()
    trs := tbody.Find("tr")

    var courses []SyllabusResultCourse

    trs.Each(func(i int, s *goquery.Selection) {
        cells := s.Find("td")

        title := cells.Eq(1).Find("span").First().Text()
        code := cells.Eq(2).Text()
        level := cells.Eq(3).Text()
        anchors := cells.Eq(4).Find(".syllabi-link")

        var docs []SyllabusResultCourseDoc

        anchors.Each(func(i int, s *goquery.Selection) {
            anchorTitle, _ := s.Attr("title")
            anchorHref, _ := s.Attr("href")

            docs = append(
                docs, SyllabusResultCourseDoc{anchorTitle, anchorHref})
        })

        courses = append(
            courses, SyllabusResultCourse{title, code, level, docs})
    })

    return SyllabusResult{message, courses}
}


// SyllabusResult

type SyllabusResult struct {
    Message string

    Courses []SyllabusResultCourse
}

type SyllabusResultCourse struct {
    Title string
    Code string
    Level string

    Docs []SyllabusResultCourseDoc
}

type SyllabusResultCourseDoc struct {
    Title string
    Link string
}
