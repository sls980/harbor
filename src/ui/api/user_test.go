// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/harbor/tests/apitests/apilib"
	"testing"
)

var testUser0002ID, testUser0003ID int
var testUser0002, testUser0003 apilib.User
var testUser0002Auth, testUser0003Auth *usrInfo

func TestUsersPost(t *testing.T) {

	fmt.Println("Testing User Add")

	assert := assert.New(t)
	apiTest := newHarborAPI()

	//case 1: register a new user without admin auth, expect 400, because self registration is on
	fmt.Println("Register user without admin auth")
	code, err := apiTest.UsersPost(testUser0002)
	if err != nil {
		t.Error("Error occured while add a test User", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 2: register a new user with admin auth, but username is empty, expect 400
	fmt.Println("Register user with admin auth, but username is empty")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 3: register a new user with admin auth, but bad username format, expect 400
	testUser0002.Username = "test@$"
	fmt.Println("Register user with admin auth, but bad username format")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 4: register a new user with admin auth, but bad userpassword format, expect 400
	testUser0002.Username = "testUser0002"
	fmt.Println("Register user with admin auth, but empty password.")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 5: register a new user with admin auth, but email is empty, expect 400
	testUser0002.Password = "testUser0002"
	fmt.Println("Register user with admin auth, but email is empty")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 6: register a new user with admin auth, but bad email format, expect 400
	testUser0002.Email = "test..."
	fmt.Println("Register user with admin auth, but bad email format")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 7: register a new user with admin auth, but userrealname is empty, expect 400
	/*
		testUser0002.Email = "testUser0002@mydomain.com"
		fmt.Println("Register user with admin auth, but user realname is empty")
		code, err = apiTest.UsersPost(testUser0002, *admin)
		if err != nil {
			t.Error("Error occured while add a user", err.Error())
			t.Log(err)
		} else {
			assert.Equal(400, code, "Add user status should be 400")
		}
	*/
	//case 8: register a new user with admin auth, but bad userrealname format, expect 400
	testUser0002.Email = "testUser0002@mydomain.com"
	testUser0002.Realname = "test$com"
	fmt.Println("Register user with admin auth, but bad user realname format")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)

	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 9: register a new user with admin auth, but bad user comment, expect 400
	testUser0002.Realname = "testUser0002"
	testUser0002.Comment = "vmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm"
	fmt.Println("Register user with admin auth, but user comment length is illegal")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Add user status should be 400")
	}

	//case 10: register a new user with admin auth, expect 201
	fmt.Println("Register user with admin auth, right parameters")
	testUser0002.Comment = "test user"
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(201, code, "Add user status should be 201")
	}

	//case 11: register duplicate user with admin auth, expect 409
	fmt.Println("Register duplicate user with admin auth")
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(409, code, "Add user status should be 409")
	}

	//case 12: register a new user with admin auth, but duplicate email, expect 409
	fmt.Println("Register user with admin auth, but duplicate email")
	testUser0002.Username = "testUsertest"
	testUser0002.Email = "testUser0002@mydomain.com"
	code, err = apiTest.UsersPost(testUser0002, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(409, code, "Add user status should be 409")
	}
}

func TestUsersGet(t *testing.T) {

	fmt.Println("Testing User Get")
	assert := assert.New(t)
	apiTest := newHarborAPI()

	testUser0002.Username = "testUser0002"
	//case 1: Get user2 with common auth, but no userid in path, expect 403

	testUser0002Auth = &usrInfo{"testUser0002", "testUser0002"}
	code, users, err := apiTest.UsersGet(testUser0002.Username, *testUser0002Auth)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Get users status should be 403")
	}
	//case 2: Get user2 with admin auth, expect 200
	code, users, err = apiTest.UsersGet(testUser0002.Username, *admin)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Get users status should be 200")
		assert.Equal(1, len(users), "Get users record should be 1 ")
		testUser0002ID = users[0].UserId
	}
}

func TestUsersGetByID(t *testing.T) {

	fmt.Println("Testing User GetByID")
	assert := assert.New(t)
	apiTest := newHarborAPI()

	//case 1: Get user2 with userID and his own auth, expect 200
	code, user, err := apiTest.UsersGetByID(testUser0002.Username, *testUser0002Auth, testUser0002ID)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Get users status should be 200")
		assert.Equal(testUser0002.Username, user.Username, "Get users username should be equal")
		assert.Equal(testUser0002.Email, user.Email, "Get users email should be equal")
	}
	//case 2: Get user2 with user3 auth, expect 403
	testUser0003.Username = "testUser0003"
	testUser0003.Email = "testUser0003@mydomain.com"
	testUser0003.Password = "testUser0003"
	testUser0003.Realname = "testUser0003"
	code, err = apiTest.UsersPost(testUser0003, *admin)
	if err != nil {
		t.Error("Error occured while add a user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(201, code, "Add user status should be 201")
	}

	testUser0003Auth = &usrInfo{"testUser0003", "testUser0003"}
	code, user, err = apiTest.UsersGetByID(testUser0002.Username, *testUser0003Auth, testUser0002ID)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Get users status should be 403")
	}
	//case 3: Get user that does not exist with user2 auth, expect 404 not found.
	code, user, err = apiTest.UsersGetByID(testUser0002.Username, *testUser0002Auth, 1000)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(404, code, "Get users status should be 404")
	}
	// Get user3ID in order to delete at the last of the test
	code, users, err := apiTest.UsersGet(testUser0003.Username, *admin)
	if err != nil {
		t.Error("Error occured while get users", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Get users status should be 200")
		assert.Equal(1, len(users), "Get users record should be 1")
		testUser0003ID = users[0].UserId
	}
}

func TestUsersPut(t *testing.T) {
	fmt.Println("Testing User Put")
	assert := assert.New(t)
	apiTest := newHarborAPI()
	var profile apilib.UserProfile
	//case 1: change user2 profile with user3 auth
	code, err := apiTest.UsersPut(testUser0002ID, profile, *testUser0003Auth)
	if err != nil {
		t.Error("Error occured while change user profile", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Change user profile status should be 403")
	}
	//case 2: change user2 profile with user2 auth, but bad parameters format.
	code, err = apiTest.UsersPut(testUser0002ID, profile, *testUser0002Auth)
	if err != nil {
		t.Error("Error occured while change user profile", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Change user profile status should be 400")
	}
	//case 3: change user2 profile with user2 auth, but duplicate email.
	profile.Realname = "test user"
	profile.Email = "testUser0003@mydomain.com"
	profile.Comment = "change profile"
	code, err = apiTest.UsersPut(testUser0002ID, profile, *testUser0002Auth)
	if err != nil {
		t.Error("Error occured while change user profile", err.Error())
		t.Log(err)
	} else {
		assert.Equal(409, code, "Change user profile status should be 409")
	}
	//case 4: change user2 profile with user2 auth, right parameters format.
	profile.Realname = "test user"
	profile.Email = "testUser0002@vmware.com"
	profile.Comment = "change profile"
	code, err = apiTest.UsersPut(testUser0002ID, profile, *testUser0002Auth)
	if err != nil {
		t.Error("Error occured while change user profile", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Change user profile status should be 200")
		testUser0002.Email = profile.Email
	}
}

func TestUsersToggleAdminRole(t *testing.T) {
	fmt.Println("Testing Toggle User Admin Role")
	assert := assert.New(t)
	apiTest := newHarborAPI()
	//case 1: toggle user2 admin role without admin auth
	code, err := apiTest.UsersToggleAdminRole(testUser0002ID, *testUser0002Auth, int32(1))
	if err != nil {
		t.Error("Error occured while toggle user admin role", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Toggle user admin role status should be 403")
	}
	//case 2: toggle user2 admin role with admin auth
	code, err = apiTest.UsersToggleAdminRole(testUser0002ID, *admin, int32(1))
	if err != nil {
		t.Error("Error occured while toggle user admin role", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Toggle user admin role status should be 200")
	}
}
func TestUsersUpdatePassword(t *testing.T) {
	fmt.Println("Testing Update User Password")
	assert := assert.New(t)
	apiTest := newHarborAPI()
	password := apilib.Password{OldPassword: "", NewPassword: ""}
	//case 1: update user2 password with user3 auth
	code, err := apiTest.UsersUpdatePassword(testUser0002ID, password, *testUser0003Auth)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Update user password status should be 403")
	}
	//case 2: update user2 password with admin auth, but oldpassword is empty
	code, err = apiTest.UsersUpdatePassword(testUser0002ID, password, *admin)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Update user password status should be 400")
	}
	//case 3: update user2 password with admin auth, but oldpassword is wrong
	password.OldPassword = "000"
	code, err = apiTest.UsersUpdatePassword(testUser0002ID, password, *admin)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Update user password status should be 403")
	}
	//case 4: update user2 password with admin auth, but newpassword is empty
	password.OldPassword = "testUser0002"
	code, err = apiTest.UsersUpdatePassword(testUser0002ID, password, *admin)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(400, code, "Update user password status should be 400")
	}
	//case 5: update user2 password with admin auth, right parameters
	password.NewPassword = "TestUser0002"
	code, err = apiTest.UsersUpdatePassword(testUser0002ID, password, *admin)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Update user password status should be 200")
		testUser0002.Password = password.NewPassword
		testUser0002Auth.Passwd = password.NewPassword
		//verify the new password takes effect
		code, user, err := apiTest.UsersGetByID(testUser0002.Username, *testUser0002Auth, testUser0002ID)
		if err != nil {
			t.Error("Error occured while get users", err.Error())
			t.Log(err)
		} else {
			assert.Equal(200, code, "Get users status should be 200")
			assert.Equal(testUser0002.Username, user.Username, "Get users username should be equal")
			assert.Equal(testUser0002.Email, user.Email, "Get users email should be equal")
		}

	}
	//case 6: update user2 password setting the new password same as the old
	password.OldPassword = password.NewPassword
	code, err = apiTest.UsersUpdatePassword(testUser0002ID, password, *admin)
	if err != nil {
		t.Error("Error occured while update user password", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "When new password is same as old, update user password status should be 200")
	}
}

func TestUsersDelete(t *testing.T) {

	fmt.Println("Testing User Delete")
	assert := assert.New(t)
	apiTest := newHarborAPI()

	//case 1:delete user without admin auth
	code, err := apiTest.UsersDelete(testUser0002ID, *testUser0003Auth)
	if err != nil {
		t.Error("Error occured while delete test user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Delete test user status should be 403")
	}
	//case 2: delete user with admin auth, user2 has already been toggled to admin, but can not delete himself
	code, err = apiTest.UsersDelete(testUser0002ID, *testUser0002Auth)
	if err != nil {
		t.Error("Error occured while delete test user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(403, code, "Delete test user status should be 403")
	}
	//case 3: delete user with admin auth
	code, err = apiTest.UsersDelete(testUser0002ID, *admin)
	if err != nil {
		t.Error("Error occured while delete test user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Delete test user status should be 200")
	}
	//delete user3 with admin auth
	code, err = apiTest.UsersDelete(testUser0003ID, *admin)
	if err != nil {
		t.Error("Error occured while delete test user", err.Error())
		t.Log(err)
	} else {
		assert.Equal(200, code, "Delete test user status should be 200")
	}
}
