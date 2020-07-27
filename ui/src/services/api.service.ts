import axios, { AxiosError, AxiosRequestConfig } from 'axios'
import config from 'config'
import { authHeader } from 'helpers'
import { authenticationService } from './authentication.service'

const api = axios.create({
    baseURL: config.apiUrl
})

api.interceptors.request.use((config: AxiosRequestConfig) => {
    const authheaders = authHeader()

    config.headers = {
        ...config.headers,
        ...authheaders
    }
    return config
}, (error: AxiosError) => {
    return Promise.reject(error)
})

api.interceptors.response.use((response) => {
    return response
}, (error: AxiosError) => {
        let retryErrorType = null

    return new Promise(async (resolve, reject) => {
        let { response, config } = error

        const status = response ? response.status : null
        if (!response) {
            retryErrorType = 'Connection'
            // network error
            return reject(retryErrorType)
        } else {
            
        }

        if (response.status >= 500 && response.status < 600) {
            retryErrorType = `Server ${response.status}`
            return reject(retryErrorType)
            // 5** errors are server related
        }

        if (status === 401) {
            // reauth
            const res = await authenticationService.refresh()
            if (res) {
                return api(config)

            }
            return reject()
        }
        return reject()
    })
})

export default api