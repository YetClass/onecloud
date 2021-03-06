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

// Code generated by model-api-gen. DO NOT EDIT.

package cloudid

import (
	"yunion.io/x/onecloud/pkg/apis"
)

// SCloudaccount is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudaccount.
type SCloudaccount struct {
	apis.SStandaloneResourceBase
	apis.SDomainizedResourceBase
	Provider         string `json:"provider"`
	Brand            string `json:"brand"`
	IamLoginUrl      string `json:"iam_login_url"`
	IsSupportCloudId *bool  `json:"is_support_cloud_id,omitempty"`
}

// SCloudaccountResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudaccountResourceBase.
type SCloudaccountResourceBase struct {
	// 云账号Id
	CloudaccountId string `json:"cloudaccount_id"`
}

// SCloudgroup is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudgroup.
type SCloudgroup struct {
	apis.SStatusInfrasResourceBase
	Provider string `json:"provider"`
}

// SCloudgroupJointsBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudgroupJointsBase.
type SCloudgroupJointsBase struct {
	apis.SJointResourceBase
	// 用户组Id
	CloudgroupId string `json:"cloudgroup_id"`
}

// SCloudgroupPolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudgroupPolicy.
type SCloudgroupPolicy struct {
	SCloudgroupJointsBase
	SCloudpolicyResourceBase
}

// SCloudgroupUser is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudgroupUser.
type SCloudgroupUser struct {
	SCloudgroupJointsBase
	SClouduserResourceBase
}

// SCloudgroupcache is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudgroupcache.
type SCloudgroupcache struct {
	apis.SStatusStandaloneResourceBase
	apis.SExternalizedResourceBase
	SCloudaccountResourceBase
	// 用户组Id
	CloudgroupId string `json:"cloudgroup_id"`
}

// SCloudpolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudpolicy.
type SCloudpolicy struct {
	apis.SStatusStandaloneResourceBase
	apis.SExternalizedResourceBase
	// 权限类型
	// | 权限类型      |  说明                |
	// |---------------|----------------------|
	// | system        | 平台内置权限         |
	// | custom        | 目前不支持           |
	PolicyType string `json:"policy_type"`
	// 平台
	Provider string `json:"provider"`
}

// SCloudpolicyResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudpolicyResourceBase.
type SCloudpolicyResourceBase struct {
	// 权限Id
	CloudpolicyId string `json:"cloudpolicy_id"`
}

// SCloudprovider is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudprovider.
type SCloudprovider struct {
	apis.SStandaloneResourceBase
	Provider       string `json:"provider"`
	CloudaccountId string `json:"cloudaccount_id"`
}

// SCloudproviderResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SCloudproviderResourceBase.
type SCloudproviderResourceBase struct {
	// 子订阅Id
	CloudproviderId string `json:"cloudprovider_id"`
}

// SClouduser is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SClouduser.
type SClouduser struct {
	apis.SStatusUserResourceBase
	apis.SExternalizedResourceBase
	SCloudproviderResourceBase
	SCloudaccountResourceBase
	Secret string `json:"secret"`
	// 是否可以控制台登录
	IsConsoleLogin *bool `json:"is_console_login,omitempty"`
	// 手机号码
	MobilePhone string `json:"mobile_phone"`
	// 邮箱地址
	Email string `json:"email"`
}

// SClouduserJointsBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SClouduserJointsBase.
type SClouduserJointsBase struct {
	apis.SJointResourceBase
	SClouduserResourceBase
}

// SClouduserPolicy is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SClouduserPolicy.
type SClouduserPolicy struct {
	SClouduserJointsBase
	SCloudpolicyResourceBase
}

// SClouduserResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudid/models.SClouduserResourceBase.
type SClouduserResourceBase struct {
	ClouduserId string `json:"clouduser_id"`
}
