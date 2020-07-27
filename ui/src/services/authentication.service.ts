import { BehaviorSubject } from 'rxjs'
import { parseJwt } from 'helpers'
import api from './api.service'

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

export interface TokenResponse {
    access_token: string
    refresh_token: string
}

function refresh() {
    if (tokenSubject.value) {
        const { refresh_token } = tokenSubject.value
        return api.post<TokenResponse>('/users/refresh', { 'refresh_token': refresh_token})
            .then((res) => {
                const tokenObject = res.data
                const parsedUser = parseJwt(tokenObject.access_token)

                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('currentUser', JSON.stringify(parsedUser))
                localStorage.setItem('refreshToken', JSON.stringify(tokenObject.refresh_token))
                localStorage.setItem('jwt', JSON.stringify(tokenObject))
                currentUserSubject.next(parsedUser)
                tokenSubject.next(tokenObject)
                refreshTokenSubject.next(tokenObject.refresh_token)

                return Promise.resolve(true)
            })
    } else {
        authenticationService.logout()
        window.location.reload(true)
        return Promise.reject(false)
    }
}

async function login(email: string, password: string) {
    return api.post<TokenResponse>('/users/authenticate', { email, password })
        .then(response => {
            const tokenObject = response.data

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