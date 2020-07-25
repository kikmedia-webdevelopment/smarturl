import { authenticationService } from 'services/authentication.service';

export function authHeader() {
    const currentObject = authenticationService.currentTokenObject
    if (currentObject && currentObject.access_token) {
        return {
            Authorization: `Bearer ${currentObject.access_token}`
        }
    }
    return {}
}