import React from 'react'
import { Route, Redirect, RouteProps } from 'react-router-dom'
import { authenticationService } from 'services'

interface Props extends RouteProps {
    
}

export const PrivateRoute: React.FC<Props> = ({ component: Component, ...rest }) => {
    return (
        <Route {...rest} render={props => {
            const currentUser = authenticationService.currentUserValue
            if (!currentUser) {
                // not logged in
                return <Redirect to={{ pathname: '/admin/login', state: { from: props.location }}} />
            }

            // authed
            // @ts-ignore
            return <Component {...props} />
        }} />
    )
}