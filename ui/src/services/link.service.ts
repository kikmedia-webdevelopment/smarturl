import { authHeader, handleResponse } from 'helpers'
import config from 'config'
import { Link } from 'models/link'

export const linkService = {
    list,
    update,
    create,
    destroy
}

function destroy(link: Link) {
    const requestOptions = {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            ...authHeader()
        },
        body: JSON.stringify(link)
    }
    return fetch(`${config.apiUrl}/links`, requestOptions as RequestInit)
        .then(handleResponse)
        .then((response: Link[]) => {
            return response
        })
}

function list() {
    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            ...authHeader()
        }
    }

    return fetch(`${config.apiUrl}/links`, requestOptions as RequestInit)
        .then(handleResponse)
        .then((response: Link[]) => {
            return response
        })
}

function update(link: Link) {
    const requestOptions = {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
            ...authHeader()
        },
        body: JSON.stringify(link)
    }
    return fetch(`${config.apiUrl}/links`, requestOptions as RequestInit)
        .then(handleResponse)
        .then(res => res)
}

function create(link: Link) {
    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            ...authHeader()
        },
        body: JSON.stringify(link)
    }
    return fetch(`${config.apiUrl}/links`, requestOptions as RequestInit)
        .then(handleResponse)
        .then(res => res)
}