// Copyright © 2017 Kris Nova <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//  _  ___
// | |/ / | ___  _ __   ___
// | ' /| |/ _ \| '_ \ / _ \
// | . \| | (_) | | | |  __/
// |_|\_\_|\___/|_| |_|\___|
//
// klone.go is top of the provider. This is the primary data structure.

package provider_alpha_1

import "github.com/kris-nova/klone/pkg/kloneprovider"

type KloneProvider struct {
}

func (k *KloneProvider) GetGitServer() (kloneprovider.GitServer, error) {
	srv := &GitServer{}
	return srv, nil
}
func NewKloner() (kloneprovider.KloneProvider, error) {
	return &KloneProvider{}, nil
}
