<p></p>
<p></p>
<p align="center" margin="20 0"><a href="https://horusec.io/"><img src="assets/horusec_logo.png" alt="logo_header" width="65%" style="max-width:100%;"/></a></p>
<p></p>
<p></p>

<p align="center">
    <a href="https://github.com/ZupIT/horusec-engine/pulse" alt="activity">
        <img src="https://img.shields.io/github/commit-activity/m/ZupIT/horusec-admin?label=activity"/></a>
    <a href="https://github.com/ZupIT/horusec-admin/graphs/contributors" alt="contributors">
        <img src="https://img.shields.io/github/contributors/ZupIT/horusec-admin?label=contributors"/></a>
    <a href="https://github.com/ZupIT/horusec-admin/actions/workflows/lint.yml" alt="lint">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-admin/Lint?label=lint"/></a>
    <a href="https://github.com/ZupIT/horusec-admin/actions/workflows/tests.yml" alt="tests">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-admin/Test?label=test"/></a>
    <a href="https://github.com/ZupIT/horusec-admin/actions/workflows/security.yml" alt="security">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-admin/Security?label=security"/></a>
    <a href="https://github.com/ZupIT/horusec-admin/actions/workflows/coverage.yml" alt="coverage">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-admin/Coverage?label=coverage"/></a>
    <a href="https://opensource.org/licenses/Apache-2.0" alt="license">
        <img src="https://img.shields.io/badge/license-Apache%202-blue"/></a>
</p>

# Horusec-Admin
This project contains 

# Running local
For run this project locally is necessary to communicate with an API server of a Kubernetes cluster.

To point to the same cluster as configured for your `kubectl`, just execute:
```bash
NAMESPACE=default KUBECONFIG=~/.kube/config go run ./cmd/app
```