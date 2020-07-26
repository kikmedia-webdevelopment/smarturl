import React from 'react'

export type IconName = 'trash' | 'edit'

interface Props {
    iconName: IconName
}

export class Icon extends React.PureComponent<Props> {
    renderIcon = () => {
        const { iconName } = this.props
        switch (iconName) {
            case 'trash': {
                return (
                    <React.Fragment>
                        <path stroke="none" d="M0 0h24v24H0z" />
                        <line x1="4" y1="7" x2="20" y2="7" />
                        <line x1="10" y1="11" x2="10" y2="17" />
                        <line x1="14" y1="11" x2="14" y2="17" />
                        <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
                        <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
                    </React.Fragment>
                )
            }
            case 'edit': {
                return (
                    <React.Fragment>
                        <path stroke="none" d="M0 0h24v24H0z" />
                        <path d="M9 7 h-3a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-3" />
                        <path d="M9 15h3l8.5 -8.5a1.5 1.5 0 0 0 -3 -3l-8.5 8.5v3" />
                        <line x1="16" y1="5" x2="19" y2="8" />
                    </React.Fragment>
                )
            }
        }
    }
    render() {
        return (
            <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6 stroke-current"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round"
            >
                {this.renderIcon()}
            </svg>
        )
    }
}