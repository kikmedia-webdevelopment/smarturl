import { authHeader, handleResponse } from 'helpers'
import config from 'config'
import { Stats } from 'models/stats'

export const statsService = {
    getStats
}

function getStats() {
    const requestOptions = {
        method: 'GET',
        headers: {
            ...authHeader(),
            'Content-Type': 'application/json',
        }
    }

    // @ts-ignore
    return fetch(`${config.apiUrl}/stats`, requestOptions)
        .then(handleResponse)
        .then((stats: Stats) => stats)
}