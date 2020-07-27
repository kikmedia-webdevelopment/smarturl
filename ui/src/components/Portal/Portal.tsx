import React from 'react'
import { canUseDom } from 'helpers'
import ReactDOM from 'react-dom'

type UnknownProps = Record<string, any>


const createContainer = (): HTMLDivElement => {
    const container = document.createElement('div')
    container.setAttribute('class', 'jk-portal')
    return container
}
const getBody = (): HTMLElement => {
    return document && document.body
}

const getPortalParent = () => {
    const parentElement = document.querySelector(
        'body > .jk-portal-container'
    )
    if (!parentElement) {
        const parent = document.createElement('div');
        parent.setAttribute('class', 'jk-portal-container');
        parent.setAttribute('style', `display: flex;`);
        getBody().appendChild(parent)
        return parent
    }
    return parentElement
}

type State = {
    container?: HTMLElement
    portalIsMounted: boolean
}

export class Portal extends React.Component<UnknownProps, State> {
    el!: HTMLDivElement
    
    constructor(props: UnknownProps) {
        super(props)

        this.state = {
            container: canUseDom ? createContainer() : undefined,
            portalIsMounted: false
        }
    }

    componentDidMount() {
        const { container } = this.state
        if (container) {
            getPortalParent().appendChild(container)
        } else {
            // SSR path
            const newContainer = createContainer()
            this.setState({
                container: newContainer
            })
        }

        this.setState({
            portalIsMounted: true
        })
    }

    componentWillUnmount() {
        const { container } = this.state
        if (container) {
            getPortalParent().removeChild(container)
            const portals = !!document.querySelector(
                'body > .jk-portal-container > .jk-portal',
            )
            if (!portals) {
                getBody().removeChild(getPortalParent());
            }
        }
    }

    render() {
        const { container, portalIsMounted } = this.state

        return container && portalIsMounted
            ? ReactDOM.createPortal(this.props.children, container)
            : null
    }
}