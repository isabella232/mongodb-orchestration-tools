// Copyright 2018 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package healthcheck

import (
	"testing"

	"github.com/timvaillancourt/go-mongodb-replset/status"
)

var (
	testMember = &status.Member{
		Id:       0,
		Name:     "localhost:27017",
		Health:   status.MemberHealthUp,
		State:    status.MemberStateRecovering,
		StateStr: "RECOVERING",
		Self:     true,
	}
	testStatus = &status.Status{
		Set:     "test",
		MyState: status.MemberStatePrimary,
		Ok:      1,
		Members: []*status.Member{
			testMember,
		},
	}
)

func TestGetSelfMemberState(t *testing.T) {
	state := getSelfMemberState(testStatus)
	if *state != testMember.State {
		t.Errorf("healthcheck.getSelfMemberState() returned %s instead of %s", *state, testMember.State)
	}
}

func TestIsMemberStateOk(t *testing.T) {
	state := getSelfMemberState(testStatus)
	if !isStateOk(state) {
		t.Errorf("healthcheck.isStateOk(\"%s\") returned false", *state)
	}

	testStatusFail := testStatus
	testStatusFail.Members[0].State = status.MemberStateRemoved
	stateFail := getSelfMemberState(testStatusFail)
	if isStateOk(stateFail) {
		t.Errorf("healthcheck.isStateOk(\"%s\") returned true", *stateFail)
	}
}
