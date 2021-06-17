function toBoolean(value) {
  return value === 'false' || value === 'undefined' || value === 'null' || value === '0' || value === false ?
  false : !!value
}

function toK8sObject(formData) {
  if (_.isEmpty(formData)) return {}

  const inputs = [...formData.elements]
  let body = {}

  console.log(inputs)

  inputs.forEach(item => {
    const value = item.type === 'checkbox' ? toBoolean(item.value) : item.value

    if (item.name && !item.name.startsWith('!')) _.set(body, item.name, value)
  })

  return body;
}