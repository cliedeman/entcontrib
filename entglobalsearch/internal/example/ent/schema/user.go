// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"entgo.io/contrib/entglobalsearch"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User defines the todo type schema.
type User struct {
	ent.Schema
}

// Fields returns user fields.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Bool("active"),
	}
}

// Annotations of the schema.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{entglobalsearch.Filter("user.Active(true)")}
}
