package JiraIssue

//
//func TestIssue_Create(t *testing.T) {
//	const path = "/rest/api/2/issue/"
//	var (
//		jira   Issue
//		domain Atlassian.Domain
//	)
//
//	s := "testdomain"
//	if err := domain.Set(&s); err != nil {
//		t.Fatal(err)
//	}
//	req, err := jira.Create(&domain)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if req.Method != http.MethodPost {
//		t.Fatal("method error")
//	}
//	if v := req.Header.Get("Content-Type"); v != "" {
//		t.Fatalf("content-type should be empty: '%s'", v)
//	}
//	if req.Header.Get("Authorization") != "" {
//		t.Fatal("authorization should be blank at this point")
//	}
//	if req.URL.String() != Atlassian.JiraUrlFactory(Atlassian.JiraUrlPattern, domain.Get(), path) {
//		t.Fatalf("url error. Got: %s", req.URL.String())
//	}
//	var body []byte
//	if _, err := req.Body.Read(body); err != nil {
//		t.Fatal(err)
//	}
//	defer func() { _ = req.Body.Close() }()
//	expected := jira.Marshall()
//	if !bytes.Equal(body, expected) {
//		t.Fatalf("body does not match expected value\n"+
//			"\tbody    : '%s'\n"+
//			"\texpected: '%s'\n", body, expected)
//	}
//}
