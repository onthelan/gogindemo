// handlers.users_test.go

func getLonginPOSTPayload() string {
  params := url.Values{}
  params.Add("username", "user1")
  params.Add("password", "pass1")

  return params.Encode()
}

func getRegistrationPOSTPayload() string {
  params := url.Values{}
  params.Add("username", "u1")
  params.Add("password", "p1")

  return params.Encode()
}
