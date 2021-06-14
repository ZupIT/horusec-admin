function jwtBuilder() {
  return {
    secreteName: "",
  };
}

function brokerBuilder() {
  return {
    enabled: true,
    host: "",
    port: 0,
    secretName: "",
  };
}

function databaseBuilder() {
  return {
    host: "",
    port: 0,
    dialect: "",
    sslMode: false,
    logMode: false,
    secretName: "",
  };
}

function keycloackClientsBuilder() {
  return {
    public: {
      id: "",
      secret: "",
    },
    confidential: {
      id: "",
      secret: "",
    },
  };
}

function keycloackBuilder() {
  return {
    publicURL: "",
    internalURL: "",
    realm: "",
    otp: false,
    clients: keycloackClientsBuilder,
  };
}

function administratorBuilder() {
  return {
    enabled: false,
    username: "",
    email: "",
    password: "",
  };
}

function ingressBuilder() {
  return {
    enabled: true,
    scheme: "",
    host: "",
    protocols: [""],
  };
}

function componentsBuilder() {
  return {
    account: {},
    analytic: {},
    api: {},
    auth: {},
    manager: {},
  };
}
