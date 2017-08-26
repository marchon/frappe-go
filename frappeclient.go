package frappe

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Application_url string
	User            string
	Password        string
}

type FrappeInput struct {
	Doctype  string
	Resource string
}

// --------------------
// global variables
var config Config

func Connect(data Config) {
	// API configuration
	config.Application_url = data.Application_url
	config.User = data.User
	config.Password = data.Password
}

// ----------------------------------------------------------------------
// this function attempts a login using the declared credentails in the
// main method.
// returns session cookies on a successful login
func login() []*http.Cookie {
	urlData := url.Values{}
	urlData.Set("cmd", "login")
	urlData.Set("usr", config.User)
	urlData.Set("pwd", config.Password)

	url := config.Application_url
	r, err := http.PostForm(url, urlData)
	checkErr(err)
	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	verified, _ := regexp.MatchString("message", string(htmlData))
	fmt.Println(verified)
	return r.Cookies()
}

// ----------------------------------------------------------------------
// this function starts a GET request to the application for requested
// doctype.
// returns JSON (which can be parsed later to make decisions)
func Get(data FrappeInput) string {
	cookies := login()
	cookieJar, _ := cookiejar.New(nil)

	fmt.Println(config.Application_url)
	// URL for cookies to remember. i.e reply when encounter this URL
	cookieURL, _ := url.Parse(config.Application_url + "/api/resource/" + data.Doctype + "/" + data.Resource)
	cookieJar.SetCookies(cookieURL, cookies)

	client := &http.Client{
		Jar: cookieJar,
	}

	r, err := client.Get(config.Application_url + "/api/resource/" + data.Doctype + "/" + data.Resource)
	checkErr(err)
	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	return string(htmlData)
}

func Put(doc FrappeInput, data url.Values) string {
	cookies := login()
	cookieJar, _ := cookiejar.New(nil)

	// URL for cookies to remember. i.e reply when encounter this URL
	cookieURL, _ := url.Parse(config.Application_url + "/api/resource/" + doc.Doctype + "/" + doc.Resource)
	cookieJar.SetCookies(cookieURL, cookies)

	client := &http.Client{
		Jar: cookieJar,
	}
	uri := config.Application_url + "/api/resource/" + doc.Doctype + "/" + doc.Resource
	reqs, _ := http.NewRequest("PUT", uri, strings.NewReader(data.Encode()))
	reqs.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r, err := client.Do(reqs)
	checkErr(err)
	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	return string(htmlData)
}

func Post(doc FrappeInput, data url.Values) string {
	cookies := login()
	cookieJar, _ := cookiejar.New(nil)

	// URL for cookies to remember. i.e reply when encounter this URL
	cookieURL, _ := url.Parse(config.Application_url + "/api/resource/" + doc.Doctype)
	cookieJar.SetCookies(cookieURL, cookies)

	client := &http.Client{
		Jar: cookieJar,
	}
	uri := config.Application_url + "/api/resource/" + doc.Doctype
	reqs, _ := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	reqs.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(strings.NewReader(data.Encode()))
	r, err := client.Do(reqs)
	checkErr(err)
	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	return string(htmlData)
}
