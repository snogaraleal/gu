package gusgs

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

const BaseUrl = "https://marknad.sgsstudentbostader.se"

var DefaultHeaders = map[string]string {
    "User-Agent": "gucli/1.0",
    "Accept": "application/json",
    "Content-Type": "application/x-www-form-urlencoded",
}

var DefaultMarkets = map[string]int {
    "direkt": 100,
    "torget": 101,
    "fortur": 102,
    "sistam": 103,
}


/*
 * Base
 */

type Utility interface {
    GetPath() (string, string)
    GetData() url.Values
    GetResp(response io.Reader) interface{}
}

func Request(utility Utility) interface{} {
    meth, path := utility.GetPath()
    data := utility.GetData().Encode()

    var body io.Reader

    // Get URL and body
    path = BaseUrl + path
    if meth == "GET" {
        path = fmt.Sprintf("%s?%s", path, data)
        body = nil
    } else {
        body = strings.NewReader(data)
    }

    // Prepare request
    req, _ := http.NewRequest(meth, path, body)
    for headerName, headerValue := range DefaultHeaders {
        req.Header.Add(headerName, headerValue)
    }

    // Send request
    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()

    // Process response
    return utility.GetResp(resp.Body)
}


/*
 * Search
 */

type SearchUtility struct {
    Market int
}

func (utility *SearchUtility) GetPath() (string, string) {
    return "POST", "/API/Service/SearchServiceHandler.ashx"
}

func (utility *SearchUtility) GetData() url.Values {
    values := url.Values{}

    search, _ := json.Marshal(utility.GetSearchData())
    values.Set("Parm1", string(search))
    values.Set("CallbackParmCount", "1")
    values.Set("CallbackMethod", "PostObjectSearch")

    return values
}

func (utility *SearchUtility) GetResp(response io.Reader) interface{} {
    var data map[string]interface{}

    dec := json.NewDecoder(response)
    dec.Decode(&data)

    return SearchResultFromData(data)
}

func (utility *SearchUtility) GetSearchData() map[string]interface{} {
    data := map[string]interface{}{
        "CompanyNo": 1,
        "SyndicateNo": 1,
        "ObjectMainGroupNo": 1,

        "Advertisements": []map[string]interface{}{
            {"No": -1},
        },

        "AreaLimit": map[string]interface{}{
            "Min": 0,
            "Max": 150,
        },

        "RentLimit": map[string]interface{}{
            "Min": 0,
            "Max": 15000,
        },

        "ReturnParameters": []string{
            "ObjectNo",
            "FirstEstateImageUrl",
            "Street",
            "SeekAreaDescription",
            "PlaceName",
            "ObjectSubDescription",
            "ObjectArea",
            "RentPerMonth",
            "MarketPlaceDescription",
            "CountInterest",
            "FirstInfoTextShort",
            "FirstInfoText",
            "EndPeriodMP",
            "FreeFrom",
            "SeekAreaUrl",
            "Latitude",
            "Longitude",
            "BoardNo",
        },

        "SortOrder": "CompanyNo asc,SeekAreaDescription asc,StreetName asc",

        "Page": 1,
        "Take": 10,
    }

    if utility.Market != 0 {
        data["MarketPlaces"] = []map[string]interface{}{
            {"No": utility.Market},
        }
    }

    return data
}


// SearchResult

type SearchResult struct {
    TotalCount int
    Items []SearchResultItem
}

func SearchResultFromData(data map[string]interface{}) SearchResult {
    items := []SearchResultItem{}
    itemsData := data["Result"].([]interface{})
    for _, itemData := range itemsData {
        itemData := itemData.(map[string]interface{})
        items = append(items, SearchResultItemFromData(itemData))
    }

    return SearchResult{int(data["TotalCount"].(float64)), items}
}


// SearchResultItem

type SearchResultItem struct {
    SeekArea string
    Address string

    Description string
    Area string
    Floor string

    LastDay string
    FreeFrom string

    Rent float64

    Properties []SearchResultItemProperty
}

func SearchResultItemFromData(data map[string]interface{}) SearchResultItem {
    item := SearchResultItem{}

    // Details
    item.SeekArea = data["SeekAreaDescription"].(string)
    item.Address = data["Street"].(string)

    item.Description = data["ObjectTypeDescription"].(string)
    item.Area = data["ObjectArea"].(string)
    item.Floor = data["ObjectFloor"].(string)

    item.LastDay = data["EndPeriodMPDateString"].(string)
    item.FreeFrom = data["FreeFrom"].(string)

    // Rent
    rent := data["RentPerMonth"]
    switch rent.(type) {
    case float64:
        item.Rent = rent.(float64)
    case string:
        item.Rent, _ = strconv.ParseFloat(rent.(string), 64)
    }

    // Properties
    itemProperties := []SearchResultItemProperty{}
    itemPropertiesData := data["Properties"].([]interface{})
    for _, itemPropertyData := range itemPropertiesData {
        itemPropertyData := itemPropertyData.(map[string]interface{})
        itemProperties = append(
            itemProperties,
            SearchResultItemPropertyFromData(itemPropertyData))
    }
    item.Properties = itemProperties

    return item
}


// SearchResultItemProperty

type SearchResultItemProperty struct {
    Description string
}

func SearchResultItemPropertyFromData(
    data map[string]interface{}) SearchResultItemProperty {

    return SearchResultItemProperty{data["PropertyDescription"].(string)}
}


/*
 * Auth
 */

type AuthUtility struct {
    username string
    password string
}

func (utility *AuthUtility) GetPath() (string, string) {
    return "GET", "/API/Service/AuthorizationServiceHandler.ashx"
}

func (utility *AuthUtility) GetData() url.Values {
    values := url.Values{}

    values.Set("SyndicateNo", "1")
    values.Set("SyndicateGroupNo", "1")
    values.Set("Method", "APILoginSGS")

    values.Set("username", utility.username)
    values.Set("password", utility.password)

    return values
}

func (utility *AuthUtility) GetResp(response io.Reader) interface{} {
    var data map[string]interface{}

    dec := json.NewDecoder(response)
    dec.Decode(&data)

    return AuthResultFromData(data)
}


// AuthResult

type AuthResult struct {
    token string
}

func AuthResultFromData(data map[string]interface{}) AuthResult {
    return AuthResult{data["SecurityTokenId"].(string)}
}
