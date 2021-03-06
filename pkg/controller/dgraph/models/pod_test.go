/*
 * Copyright (c) 2018 VMware Inc. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import (
	"fmt"
	"testing"

	"github.com/vmware/purser/pkg/controller/dgraph"
)

// TestStorePodsInteraction ...
func TestStorePodsInteraction(t *testing.T) {
	fmt.Println("Hello World")
	err := dgraph.Open("127.0.0.1:9080")
	if err != nil {
		fmt.Println("Error while opening connection to Dgraph ", err)
	}

	err = dgraph.CreateSchema()
	if err != nil {
		fmt.Println("Error while creating schema ", err)
	}

	sourcePod := "weave:weave-scope-app-6d6b76b846-z92wk"
	destinationPods := []string{"fiaasco:ccs-billing-deployment-1-1-92-75dc8749f4-gld6q", "weave:weave-scope-agent-lbfpj"}
	interactionCounts := []float64{2.0}

	err = StorePodsInteraction(sourcePod, destinationPods, interactionCounts)
	if err != nil {
		fmt.Println("Error while building interation graph ", err)
	}
}
