/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package path implements mep server api plans
package plans

import (
	"encoding/json"

	"github.com/apache/servicecomb-service-center/pkg/log"

	"mepserver/common/arch/workspace"
	"mepserver/common/extif/backend"
	"mepserver/common/extif/dns"
	"mepserver/common/util"
	"mepserver/mm5/models"
)

type DNSRuleGet struct {
	workspace.TaskBase
	AppInstanceId string      `json:"appInstanceId,in"`
	DNSRuleId     string      `json:"dnsRuleId,in"`
	HttpRsp       interface{} `json:"httpRsp,out"`
}

func (t *DNSRuleGet) OnRequest(inputData string) workspace.TaskCode {
	log.Debugf("query request arrived to fetch dns rule %s for appId %s.", t.DNSRuleId, t.AppInstanceId)

	if len(t.DNSRuleId) == 0 {
		log.Errorf(nil, "invalid dns id on query request")
		t.SetFirstErrorCode(util.ParseInfoErr, "invalid query request")
		return workspace.TaskFinish
	}

	dnsRlEntry, err := backend.GetRecord(util.EndDNSRuleKeyPath + t.AppInstanceId + "/" + t.DNSRuleId)
	if err != 0 {
		log.Errorf(nil, "get dns rules from data-store failed")
		t.SetFirstErrorCode(workspace.ErrCode(err), "dns rule retrieval failed")
		return workspace.TaskFinish
	}

	dnsRlInStore := &dns.RuleEntry{}
	jsonErr := json.Unmarshal(dnsRlEntry, dnsRlInStore)
	if jsonErr != nil {
		log.Errorf(nil, "failed to parse the dns entry from data-store")
		t.SetFirstErrorCode(util.OperateDataWithEtcdErr, "parse dns rules from data-store failed")
		return workspace.TaskFinish
	}
	t.HttpRsp = models.NewDnsConfigRule(
		t.DNSRuleId,
		dnsRlInStore.DomainName,
		dnsRlInStore.IpAddressType,
		dnsRlInStore.IpAddress,
		dnsRlInStore.TTL,
		dnsRlInStore.State)
	return workspace.TaskFinish
}
