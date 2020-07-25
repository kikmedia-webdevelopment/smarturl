import React from 'react';

interface Props extends React.HTMLProps<HTMLLabelElement> {}

export class Label extends React.Component<Props, {}> {
    render() {
        const { ...props } = this.props
        return (
            <label
                className="block text-sm font-medium leading-6 text-gray-700"
                {...props}
            />
        )
    }
}