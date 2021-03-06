//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

import (
	"encoding/json"
	"fmt"
)

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4AddressBlock struct {
	ConfigObj
	NwAddress string `SNAPROUTE: "KEY"`
	NwMask    string `SNAPROUTE: "KEY"`
	NwName    string `SNAPROUTE: "KEY"`
}

func (obj IPV4AddressBlock) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4block IPV4AddressBlock
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4block); err != nil {
			fmt.Printf("### Trouble in unmarshaling IPV4Block from Json Error: %s", err)
		}
	}
	return v4block, err
}
