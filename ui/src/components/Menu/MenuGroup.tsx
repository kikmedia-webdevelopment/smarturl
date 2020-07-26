import React from 'react'

interface Props {
    title?: string
}

export class MenuGroup extends React.PureComponent<Props> {
    render() {
        const { title, children } = this.props
        return (
            <div className="border-b- border-solid border-gray-300 py-2">
                {title && (
                    <h2
                        className="text-gray-700 leading-none uppercase text-xs my-2 mx-4 font-normal font-sans"
                    >{title}</h2>
                )}
                {children}
            </div>
        )
    }
}