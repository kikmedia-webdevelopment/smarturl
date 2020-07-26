import React from 'react'

interface Props {
    onClick?: (event: React.MouseEvent<HTMLDivElement, MouseEvent>) => void
}

export class MenuItem extends React.PureComponent<Props> {

    handleClick = (event: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
        if (this.props.onClick) {
            this.props.onClick(event)
        }
    }

    render() {
        const { children } = this.props
        return (
            <div
                className="cursor-pointer flex outline-none no-underline text-black items-center h-10"
                role="menuitem"
                tabIndex={0}
                onClick={this.handleClick}
            >
                <span
                    className="flex-1 ml-4 mr-4 font-sans text-sm"
                >
                    {children}
                </span>
            </div>
        )
    }
}