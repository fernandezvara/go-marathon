/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import (
	"testing"
)

func TestGroups(t *testing.T) {
	NewFakeMarathonEndpoint()
	groups, err := test_client.Groups()
	AssertOnError(err, t)
	AssertOnNull(groups, t)
	AssertOnInteger(len(groups.Groups), 1, t)
	group := groups.Groups[0]
	AssertOnString(group.ID, FAKE_GROUP_NAME, t)
}

func TestGroup(t *testing.T) {
	NewFakeMarathonEndpoint()
	group, err := test_client.Group(FAKE_GROUP_NAME)
	AssertOnError(err, t)
	AssertOnNull(group, t)
	AssertOnInteger(len(group.Apps), 1, t)
	AssertOnString(group.ID, FAKE_GROUP_NAME, t)
}
