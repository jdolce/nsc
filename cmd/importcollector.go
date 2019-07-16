/*
 * Copyright 2018-2019 The NATS Authors
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"errors"
	"fmt"
	"sort"

	"github.com/nats-io/jwt"
	"github.com/nats-io/nsc/cmd/store"
)

type AccountExport struct {
	OperatorName string
	jwt.AccountClaims
}

type ByAccountName []AccountExport

func (a ByAccountName) Len() int           { return len(a) }
func (a ByAccountName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccountName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type AccountExportChoice struct {
	AccountExport
	Selection *jwt.Export
}

type AccountExportChoices []AccountExportChoice

func (aes AccountExportChoices) String() []string {
	var choices []string
	for _, c := range aes {
		choices = append(choices, c.String())
	}
	return choices
}

func (aec *AccountExportChoice) String() string {
	// label
	if aec.Selection == nil {
		return fmt.Sprintf("%s:", aec.Name)
	}
	// an actual export
	k := "->"
	if aec.Selection.IsService() {
		k = "<-"
	}

	p := ""
	if aec.Selection.TokenReq {
		p = "[!]"
	}

	if aec.Selection.Name == string(aec.Selection.Subject) {
		return fmt.Sprintf("  %s %s %s", k, aec.Selection.Subject, p)
	}

	return fmt.Sprintf("  %s [%s] %s %s", k, aec.Selection.Name, aec.Selection.Subject, p)
}

func GetAllExports() ([]AccountExport, error) {
	var exports []AccountExport

	config := GetConfig()
	if config.StoreRoot == "" {
		return nil, errors.New("no store set - `env --store <dir>`")
	}
	operators := config.ListOperators()
	for _, o := range operators {
		if o == "" {
			continue
		}

		s, err := config.LoadStore(o)
		if err != nil {
			return nil, err
		}
		accounts, err := s.ListSubContainers(store.Accounts)
		if err != nil {
			return nil, err
		}
		for _, a := range accounts {
			ac, err := s.ReadAccountClaim(a)
			if err != nil {
				return nil, err
			}

			if len(ac.Exports) == 0 {
				continue
			}

			var export AccountExport
			export.OperatorName = o
			export.Exports = ac.Exports
			export.AccountClaims = *ac
			exports = append(exports, export)
		}
	}

	sort.Sort(ByAccountName(exports))
	return exports, nil
}

func (ae *AccountExport) Choices() []AccountExportChoice {
	var choices []AccountExportChoice
	// choice without a selection is an account label
	choices = append(choices, AccountExportChoice{AccountExport: *ae})
	for _, v := range ae.Exports {
		choices = append(choices, AccountExportChoice{*ae, v})
	}
	return choices
}