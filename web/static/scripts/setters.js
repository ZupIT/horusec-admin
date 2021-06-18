/**
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

function setAuthType(authType) {
    const inputAuthType = document.getElementsByName('components.auth.type')[0];
    inputAuthType.value = authType;

    document.getElementById('keycloak-form').style.display = 'none';
    document.getElementById('ldap-form').style.display = 'none';

    if (authType !== 'horusec') document.getElementById(`${authType}-form`).style.display = 'flex';
}

function setCheckboxValueByName(name) {
    const checkbox = document.getElementsByName(name)[0]
    checkbox.value = !toBoolean(checkbox.value);
}

function setApplicationAdmin() {
    const switchInput = document.getElementsByName('components.auth.user.administrator.enabled')[0];
    switchInput.value = !toBoolean(switchInput.value);

    const form = document.getElementById('application-admin-form');
    form.style.display = toBoolean(switchInput.value) ? 'block' : 'none';
}

function setDefaultUser() {
    const switchInput = document.getElementsByName('components.auth.user.default.enabled')[0];
    switchInput.value = !toBoolean(switchInput.value);

    const form = document.getElementById('default-user-form');
    form.style.display = toBoolean(switchInput.value) ? 'block' : 'none';
}