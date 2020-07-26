import React from 'react'
import { MenuGroup } from './MenuGroup'
import { MenuItem } from './MenuItem'

export class Menu extends React.PureComponent {
    static Group = MenuGroup
    static Item = MenuItem

    menuRef: React.RefObject<HTMLElement>

    constructor(props: any) {
        super(props)

        this.menuRef = React.createRef<HTMLElement>()
    }

    componentDidMount() {

    }

    render() {
        const { children } = this.props
        return (
            <nav
                className="py-1 rounded-md bg-white shadow-xs"
                role="menu"
                ref={this.menuRef}
            > 
                {children}
            </nav>
        )
    }
}