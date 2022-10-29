package iarapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	testClient http.Client

	subscriberInfoGood = SubscriberInfo{
		Id:       0,
		StatusID: 1,
		Name:     "TestSubscriber",
		IsActive: true,
	}

	memberInfoGood = MemberInfo{
		Id:           123456,
		SubscriberId: 654321,
		FirstName:    "Test",
		LastName:     "User",
	}

	incidentListGood = IncidentList{
		{
			Id:           12345678,
			SubscriberId: 654321,
			ArrivedOn:    "today",
		},
	}

	messageListGood = MessageList{
		{
			Id:           "87654321",
			MessageId:    1098765,
			SubscriberId: 654321,
			Message:      "HelloWorld!",
		},
	}

	dispatcherListGood = Dispatchers{
		SubscriberId: 654321,
		Dispatchers: []*Dispatcher{
			{
				DispatcherId:   555,
				DispatcherName: "Test Dispatcher",
			},
		},
	}

	responseCodesGood = ResponderCodes{
		ResponseCodes: []*ResponderCode{
			{
				Id:             1,
				SubscriberId:   654321,
				IsTelephoneKey: false,
				IsDefaultKey:   true,
			},
		},
		TelephoneKeys: []*ResponderCode{
			{
				Id:             11,
				SubscriberId:   654321,
				IsTelephoneKey: true,
				IsDefaultKey:   true,
			},
		},
	}

	onDutyAtCodesGood = OnDutyAtCodeList{
		{
			Id:           "444",
			SubscriberId: 654321,
			KeyEntry:     "11",
		},
	}

	responderListGood = ResponderList{
		{
			Id:           "321",
			SubscriberId: 654321,
			RespondingTo: "station",
		},
	}

	apparatusListGood = ApparatusList{
		{},
	}
)

func TestMain(m *testing.M) {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/Subscriber", subscriberHandler)
	http.HandleFunc("/Member", memberHandler)
	http.HandleFunc("/IncidentList", incidentListHandler)
	http.HandleFunc("/MessageList", messageListHandler)
	http.HandleFunc("/DispatcherContent/AssociatedDispatchers", dispatchersHandler)
	http.HandleFunc("/ResponderCodes", responderCodesHandler)
	http.HandleFunc("/OnDutyAtCodes", onDutyAtCodesHandler)
	http.HandleFunc("/ResponderList", responderListHandler)
	http.HandleFunc("/ApparatusList", apparatusListHandler)
	http.HandleFunc("/SearchIncidents", searchIncidentsHandler)

	ts := httptest.NewTLSServer(nil)
	defer ts.Close()

	testClient = *ts.Client()
	loginUrl = ts.URL + "/login"
	apiBase = ts.URL
	apiHost = ts.URL

	os.Exit(m.Run())
}

func TestClient_login(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	type args struct {
		agency   string
		user     string
		password string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "good",
			args:   args{agency: "good", user: "good", password: "good"},
			fields: fields{httpClient: testClient},
		},
		{
			name:    "bad agency",
			args:    args{agency: "", user: "good", password: "good"},
			fields:  fields{httpClient: testClient},
			wantErr: true,
		},
		{
			name:    "bad credentials",
			args:    args{agency: "good", user: "good", password: "bad"},
			fields:  fields{httpClient: testClient},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			if err := c.login(tt.args.agency, tt.args.user, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Client.login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Subscriber(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *SubscriberInfo
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &subscriberInfoGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.Subscriber()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Subscriber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Subscriber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Member(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *MemberInfo
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &memberInfoGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.Member()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Member() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Member() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Incidents(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *IncidentList
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &incidentListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.Incidents()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Incidents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Incidents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Messages(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *MessageList
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &messageListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.Messages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Messages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Messages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Dispatchers(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *Dispatchers
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &dispatcherListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.Dispatchers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Dispatchers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Dispatchers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ResponderCodes(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *ResponderCodes
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &responseCodesGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.ResponderCodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ResponderCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ResponderCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_OnDutyAtCodes(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *OnDutyAtCodeList
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &onDutyAtCodesGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.OnDutyAtCodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.OnDutyAtCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.OnDutyAtCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ResponderList(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *ResponderList
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &responderListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.ResponderList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ResponderList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ResponderList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ApparatusList(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    *ApparatusList
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			want:   &apparatusListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.ApparatusList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ApparatusList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ApparatusList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchIncidents(t *testing.T) {
	type fields struct {
		httpClient http.Client
	}

	type args struct {
		isr *IncidentSearchRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IncidentList
		wantErr bool
	}{
		{
			name:   "good",
			fields: fields{httpClient: testClient},
			args: args{isr: &IncidentSearchRequest{
				StartTime: time.Now().Add(-24 * time.Hour),
				EndTime:   time.Now(),
				PageSize:  100,
			}},
			want: &incidentListGood,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			got, err := c.SearchIncidents(tt.args.isr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchIncidents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SearchIncidents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	req := new(LoginRequest)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reply := new(LoginReply)

	if req.Agency == "good" && req.User == "good" && req.Password == "good" {
		reply.Message = "Login to iamresponding.com/ successful"
	}

	sendResponse(w, r, reply)
}

func subscriberHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &subscriberInfoGood)
}

func memberHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &memberInfoGood)
}

func incidentListHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &incidentListGood)
}

func messageListHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &messageListGood)
}

func dispatchersHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &dispatcherListGood)
}

func responderCodesHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &responseCodesGood)
}

func onDutyAtCodesHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &onDutyAtCodesGood)
}

func responderListHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &responderListGood)
}

func apparatusListHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &apparatusListGood)
}

func searchIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, r, &incidentListGood)
}

func sendResponse(w http.ResponseWriter, r *http.Request, t interface{}) {
	defer r.Body.Close()

	data, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
