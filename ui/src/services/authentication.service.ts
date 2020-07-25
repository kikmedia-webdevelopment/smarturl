import { BehaviorSubject } from 'rxjs'
import config from 'config'
import { handleResponse } from 'helpers/handleResponse'
import { parseJwt } from 'helpers'

// @ts-ignore
const currentUserSubject = new BehaviorSubject(JSON.parse(localStorage.getItem('currentUser')))
const tokenSubject = new BehaviorSubject<TokenResponse | null>(null)
const refreshTokenSubject = new BehaviorSubject("")

const parsedTokenObject = localStorage.getItem('jwt')
if (parsedTokenObject) {
    tokenSubject.next(JSON.parse(parsedTokenObject))
}

const parsedRefreshToken = localStorage.getItem('refreshToken')
if (parsedRefreshToken) {
    refreshTokenSubject.next(JSON.parse(parsedRefreshToken))
}


export const authenticationService = {
    login,
    logout,
    refresh,
    currentUser: currentUserSubject.asObservable(),
    get currentUserValue() { return currentUserSubject.value },
    get currentTokenObject() { return tokenSubject.value }
}

interface TokenResponse {
    access_token: string
    refresh_token: string
}

function refresh() {
    if (tokenSubject.value) {
        const { refresh_token } = tokenSubject.value
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ refresh_token })
        }
        return fetch(`${config.apiUrl}/users/refresh`, requestOptions)
            .then(handleResponse)
            .then((tokenObject: TokenResponse) => {
                const parsedUser = parseJwt(tokenObject.access_token)
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('currentUser', JSON.stringify(parsedUser))
                localStorage.setItem('refreshToken', JSON.stringify(tokenObject.refresh_token))
                localStorage.setItem('jwt', JSON.stringify(tokenObject))
                currentUserSubject.next(parsedUser)
                tokenSubject.next(tokenObject)
                refreshTokenSubject.next(tokenObject.refresh_token)

                return tokenObject
            })

    } else {
        authenticationService.logout()
        window.location.reload(true)
    }
}

function login(email: string, password: string) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    }

    return fetch(`${config.apiUrl}/users/authenticate`, requestOptions)
        .then(handleResponse)
        .then((tokenObject: TokenResponse) => {
            const parsedUser = parseJwt(tokenObject.access_token)
            // store user details and jwt token in local storage to keep user logged in between page refreshes
            localStorage.setItem('currentUser', JSON.stringify(parsedUser))
            localStorage.setItem('refreshToken', JSON.stringify(tokenObject.refresh_token))
            localStorage.setItem('jwt', JSON.stringify(tokenObject))
            currentUserSubject.next(parsedUser)
            tokenSubject.next(tokenObject)
            refreshTokenSubject.next(tokenObject.refresh_token)

            return tokenObject
        })
}

function logout() {
    // remove user from local storage to log out
    localStorage.removeItem('currentUser')
    currentUserSubject.next(null)
}