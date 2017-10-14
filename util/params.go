package util

type PageInfo struct {
    PageSize string
    PageIndex string
}

type Actor struct {
    PageInfo
}

type Channel struct {

}

type Class struct {
    PageInfo
}

type MovieInfo struct {
    MovieID string
}

type Movies struct {
    PageInfo
    Type string
    ID string
    Data string
}

type Login struct {
	Token string
}