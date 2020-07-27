import { authHeader, handleResponse } from 'helpers'
import config from 'config'

export const userService = {
    getAll,
    updatePassword
}

function getAll() {
    const requestOptions = { method: 'GET', headers: authHeader() }
    // @ts-ignore
    return fetch(`${config.apiUrl}/users`, requestOptions).then(handleResponse)
}

function updatePassword(params: {
        currentPassword: string
        newPassword: string
        newPasswordConfirmed: string
}) {
    const requestOptions = {
        method: 'POST',
        headers: { ...authHeader(), 'Content-Type': 'application/json',},
        body: JSON.stringify(params)
    }
    // @ts-ignore
    return fetch(`${config.apiUrl}/users/change-password`, requestOptions).then(handleResponse)
}