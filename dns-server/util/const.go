/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
package util

const (
	DbStringExceptions    = "/.\\"
	DefaultDnsPort        = 53
	DefaultManagementPort = 80
	DefaultConnTimeout    = 2
	MinConnTimeout        = 2
	MaxConnTimeout        = 50
	MaxDbNameLength       = 256
	MaxPortNumber         = 65535
	DefaultTTL            = 30
	DNSUDPPacketSize      = 65535
	ForwardRetryCount     = 3
	DefaultIP             = "0.0.0.0"
)

const MaxDnsQuestionLength = 254
