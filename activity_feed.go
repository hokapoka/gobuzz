package gobuzz

import(
	"time"
	"strings"
)

type ActivityFeed struct{
	Kind string
	Links map[string][]link
	Title string
	Updated *jTime
	Id string
	Items []Activity
}

func GetTime(c string) *time.Time{
	t, _ := time.Parse( "2006-01-02T15:04:05", strings.Split(c, ".", -1)[0])
	return t
}


type link struct{
	Href string
	Type string
	Count int
}

type Activity struct{
	Kind string
	Title string
	Published *jTime
	Updated *jTime
	Id string
	Links map[string][]link
	Actor Actor
	Verbs []string
	Object Object
	Source map[string]string
	Visibility map[string][]Visibility
}

type Actor struct{
	Id string
	Name string
	ProfileUrl string
	ThumbnailUrl string

}

type Object struct{
	Type string
	Content string
	Links map[string][]link

}

type Visibility struct{
	Id string
	Title string
}


