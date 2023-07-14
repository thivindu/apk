#!/bin/bash
# --------------------------------------------------------------------
# Copyright (c) 2023, WSO2 LLC. (http://wso2.com) All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# -----------------------------------------------------------------------

COMMON_ADAPTER_XDS_PORT="${COMMON_ADAPTER_XDS_PORT:-18005}"
COMMON_ADAPTER_SERVER_NAME="${COMMON_ADAPTER_SERVER_NAME:-commonadapter}"
grpc_health_probe -addr "127.0.0.1:${COMMON_ADAPTER_XDS_PORT}" \
    -tls \
    -tls-ca-cert "${COMMON_ADAPTER_PUBLIC_CERT_PATH}" \
    -tls-client-cert "${COMMON_ADAPTER_PUBLIC_CERT_PATH}" \
    -tls-client-key "${COMMON_ADAPTER_PRIVATE_KEY_PATH}" \
    -tls-server-name ${COMMON_ADAPTER_SERVER_NAME} \
    -connect-timeout=3s
    