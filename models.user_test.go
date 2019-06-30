// models.user_test.go

func TestInvalidUserRegistration(t *testing.T) {
  saveLists()

  u, err := registerNewUser("user1", "pass1")

  if err == nil || u != nil {
    t.fail()
  }

  u, err = registerNewUser("newuser", "")

  if err == nil || u != nil {
    t.Fail()
  }

  restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
    saveLists()

    u, err := registerNewUser ("newuser", "newpass")

    if err != nil || u.Username == "" {
      t.Fail()
    }

    restoreLists()
}

func TestUsernameAvailability(t *testing.T) {
  saveLists()

  if !isUsernameAvailable("newuser") {
    t.Fail()
  }

  if isUsernameAvailable("user1") {
    t.fail()
  }

  registerNewUser("newuser", "newpass")

  if isUsernameAvailable("newuser") {
    t.fail()
  }

  restoreLists()
}
