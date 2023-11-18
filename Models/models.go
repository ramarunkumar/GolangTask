package models

type Event struct {
	Ev    string `json:"ev"`
	Et    string `json:"et"`
	ID    string `json:"id"`
	UID   string `json:"uid"`
	MID   string `json:"mid"`
	T     string `json:"t"`
	P     string `json:"p"`
	L     string `json:"l"`
	SC    string `json:"sc"`
	Atk1  string `json:"atrk1"`
	Atv1  string `json:"atrv1"`
	Atr1  string `json:"atrt1"`
	Atk2  string `json:"atrk2"`
	Atv2  string `json:"atrv2"`
	Atr2  string `json:"atrt2"`
	Uatk1 string `json:"uatrk1"`
	Uatv1 string `json:"uatrv1"`
	Uatr1 string `json:"uatrt1"`
	Uatk2 string `json:"uatrk2"`
	Uatv2 string `json:"uatrv2"`
	Uatr2 string `json:"uatrt2"`
	Uatk3 string `json:"uatrk3"`
	Uatv3 string `json:"uatrv3"`
	Uatr3 string `json:"uatrt3"`
}

type ConvertedEvent struct {
	Event       string               `json:"event"`
	EventType   string               `json:"event_type"`
	AppID       string               `json:"app_id"`
	UserID      string               `json:"user_id"`
	MessageID   string               `json:"message_id"`
	PageTitle   string               `json:"page_title"`
	PageURL     string               `json:"page_url"`
	BrowserLang string               `json:"browser_language"`
	ScreenSize  string               `json:"screen_size"`
	Attributes  map[string]Attribute `json:"attributes"`
	Traits      map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
