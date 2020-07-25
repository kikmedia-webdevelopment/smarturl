import React from 'react';

interface Props {
    label: string
}

export class ModalHeader extends React.Component<Props> {
    render() {
        return (
            <div className="flex flex-shrink-0 items-center py-6 px-8 rounded-t-sm border-b border-solid border-gray-300">
                <h1 className="flex-grow-1 m-0 text-black font-sans font-bold text-base overflow-hidden leading-normal truncate">{this.props.label}</h1>
            </div>
        )
    }
}