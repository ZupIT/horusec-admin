const toBoolean = (value) =>
  value === 'false' || value === 'undefined' || value === 'null' || value === '0' || value === false ?
  false : !!value