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

function toBoolean(value) {
  return value === 'false' || value === 'undefined' || value === 'null' || value === '0' || value === false ?
  false : !!value
}

function toK8sObject(formData) {
  if (_.isEmpty(formData)) return {}

  const inputs = [...formData.elements]
  let body = {}

  inputs.forEach(item => {
    let value;

    if (item.type === 'checkbox') value = toBoolean(item.value);
    else if (item.type === 'number') value = parseInt(item.value);
    else value = item.value;

    if (item.name && !item.name.startsWith('!')) _.set(body, item.name, value)
  })

  return body;
}

function toFormData(form, spec) {
  if (_.isEmpty(form) || _.isEmpty(spec)) return;

  const elements = [...form.elements]

  elements.forEach((el) => {
    const value = _.get(spec, el.name);

    switch (el.name) {
      case 'components.auth.user.administrator.enabled':
        setApplicationAdmin(value);
        break;

      case 'components.auth.user.default.enabled':
        setDefaultUser(value);
        break;

      case 'components.auth.type':
        setAuthType(value);
        break;

      case 'components.auth.user.default.enabled':
      case 'components.auth.user.administrator.enabled':
      case 'global.database.logMode':
      case 'global.database.sslMode':
      case 'components.analytic.database.logMode':
      case 'components.analytic.database.sslMode':
      case 'components.messages.enabled':
      case 'global.keycloak.otp':
      case 'global.ldap.useSsl':
      case 'global.ldap.skipTls':
      case 'global.ldap.insecureSkipVerify':
        setCheckboxValueByName(el.name, value)
        break;

      default:
        el.value = value ? value : '';
        break;
    }
  })

  return form;
}
