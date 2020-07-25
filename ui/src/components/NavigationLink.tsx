import React from 'react'
import { NavLink, NavLinkProps } from 'react-router-dom'

interface Props extends NavLinkProps {
    to: string
}

export class NavigationLink extends React.Component<Props> {
    render() {
        const { to, children } = this.props
        return (
            <NavLink to={to}
                className="ml-4 px-3 py-2 rounded-md text-sm font-medium text-black hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700"
                activeClassName="bg-gray-900 text-white"
            >
                {children}
            </NavLink>
        )
    }
}