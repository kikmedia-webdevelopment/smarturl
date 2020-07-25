import { authenticationService } from 'services/authentication.service'

let retries = 0

export function handleResponse(response: Response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text)
        if (!response.ok) {
            if ([401, 403].indexOf(response.status) !== -1) {
                if (retries === 0) {
                    retries = 1
                    authenticationService.refresh()
                } else {
                    retries = 0
                    authenticationService.logout()
                    window.location.reload(true)
                }
                // auto logout if 401 unauthorized or 403 forbidden response returns from api
                
            }

            const error = (data && data.message) || response.statusText
            return Promise.reject(error)
        }

        return data
    })
}