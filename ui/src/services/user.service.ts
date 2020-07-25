import { authHeader, handleResponse } from 'helpers'
import config from 'config'

export const userService = {
    getAll
}

function getAll() {
    const requestOptions = { method: 'GET', headers: authHeader() }
    // @ts-ignore
    return fetch(`${config.apiUrl}/users`, requestOptions).then(handleResponse)
}