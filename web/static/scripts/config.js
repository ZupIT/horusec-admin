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

 function setCurrentValues() {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/config', true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.setRequestHeader('X-Horusec-Authorization', getCookie('horusec::access_token'))

    xhr.send()

    xhr.onreadystatechange = function (ev) {
        if (ev.currentTarget.status === 401) {
            window.location.href = '/view/not-authorized'
        }

        if (ev.currentTarget.status === 200 && ev.currentTarget.response) {
            const result = JSON.parse(ev.currentTarget.response)

             Object.entries(result).forEach(item => {
                const element = document.getElementById(item[0])

                if (element) {
                    element.value = item[1]

                    if (element.type === 'checkbox' && item[1]) {
                        element.checked = JSON.parse(item[1])
                    }

                    if (item[0] === 'horusec_auth_type') {
                        setAuthType(item[1], true)
                    }

                    if (item[0] === 'horusec_application_admin_data') {
                        setDataOfAdminApplication(item[1])
                    }

                    if (item[0] === 'disable_emails') {
                        setDisableEmail(item[1])
                    }
                }
            });
        }
    }
}
function setAuthType(authType, setRadioOption) {
    document.getElementById('horusec_auth_type').value = authType

     document.getElementById('keycloak-form').style.display = 'none'
     document.getElementById('ldap-form').style.display = 'none'

    if (authType !== 'horusec') {
         document.getElementById(`${authType}-form`).style.display = 'flex'
    }

    if (setRadioOption) {
        document.getElementById(`auth-${authType}`).checked = true
    }
 }
function setDataOfAdminApplication(data) {
    try {
        const appAdminData = JSON.parse(data)

        const { email, password, username } = appAdminData

        document.getElementById('admin_data_username').value = username
        document.getElementById('admin_data_email').value = email
        document.getElementById('admin_data_password').value = password
    } catch (e) {
        console.info('No admin data to set')
    }
}
function setEnableApplicationAdmin(value) {
    document.getElementById('horusec_enable_application_admin').value = value

    document.getElementById('admin_data_username').required = value == 'true' ? true : false
    document.getElementById('admin_data_email').required = value == 'true' ? true : false
    document.getElementById('admin_data_password').required = value == 'true' ? true : false

    if (value == 'false') {
        document.getElementById('horusec_application_admin_data').style.display = 'none'
    } else {
        document.getElementById('horusec_application_admin_data').style.display = 'flex'

    }
}
function setDisableEmail(value) {
    document.getElementById('disable_emails').value = value

    document.getElementById('smtp_username').required = value == 'true' ? false : true
    document.getElementById('smtp_password').required = value == 'true' ? false : true
    document.getElementById('smtp_host').required = value == 'true' ? false : true
    document.getElementById('smtp_port').required = value == 'true' ? false : true
    document.getElementById('email_from').required = value == 'true' ? false : true

    if (value === 'false') {
        document.getElementById('smtp-fields').style.display = 'flex'
    } else {
        document.getElementById('smtp-fields').style.display = 'none'
    }
}
