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

function saveData(body) {
    try {
        const xhr = new XMLHttpRequest();
        xhr.open('PATCH', '/api/config', true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.setRequestHeader('X-Horusec-Authorization', getCookie('horusec::access_token'))

        xhr.send(JSON.stringify(body))

        xhr.onreadystatechange = function (ev) {
            if (ev.currentTarget.status === 204) {
                showSnackBar('success')
            } else if (ev.currentTarget.status === 401) {
                window.location.href = '/view/not-authorized'
            } else {
                showSnackBar('error')
            }
        }
    } catch (e) {
        console.error(error)
        showSnackBar('error')
    }
}

function getCurrentData() {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/config', true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.setRequestHeader('X-Horusec-Authorization', getCookie('horusec::access_token'))

    xhr.send()

    return new Promise((resolve) => {
        xhr.onreadystatechange = function (ev) {
            if (ev.currentTarget.status === 401) {
                window.location.href = '/view/not-authorized'
            }

            if (ev.currentTarget.status === 200 && ev.currentTarget.response) {
                resolve(JSON.parse(ev.currentTarget.response))
            }
        }
    })
}
