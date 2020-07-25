import React from 'react';

export class ModalContent extends React.Component {
    render() {
        return (
            <div className="p-8 text-sm font-sans leading-normal overflow-y-auto overflow-x-hidden">
                {this.props.children}
            </div>
        )
    }
}