import React from 'react'
import { canUseDom } from 'helpers'
import ReactDOM from 'react-dom'

type UnknownProps = Record<string, any>

let portalContainer: HTMLDivElement

export class Portal extends React.Component<UnknownProps> {
    el!: HTMLDivElement
    
    constructor(props: UnknownProps) {
        super(props)

        // This fixes SSR
        if (!canUseDom) return

        if (!portalContainer) {
            portalContainer = document.createElement('div')
            portalContainer.setAttribute('jk-portal-container', '')
            document.body.appendChild(portalContainer)
        }

        this.el = document.createElement('div')
        portalContainer.appendChild(this.el)
    }

    componentWillUnmount() {
        portalContainer.removeChild(this.el)
    }

    render() {
        // This fixes SSR
        if (!canUseDom) return null
        return ReactDOM.createPortal(this.props.children, this.el)
    }
}