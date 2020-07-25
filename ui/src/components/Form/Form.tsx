import React from 'react'
import {uid} from 'react-uid'

export type FormProps = {
    onSubmit?: React.FormEventHandler
    testId?: string
    children: React.ReactChild | React.ReactNodeArray
} & typeof defaultProps

const defaultProps = {
    testId: 'jk-ui-form'
}

export class Form extends React.Component<FormProps> {
    static defaultProps = defaultProps

    handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();
        if (this.props.onSubmit) {
            this.props.onSubmit(event);
        }
    }

    render() {
        const { testId, children, ...otherProps } = this.props
        return (
            <form
                data-test-id={testId}
                onSubmit={this.handleSubmit}
                {...otherProps}
            >
                {React.Children.map(children, child => {
                    if (child) {
                        return <div key={uid("any")} className="block mb-4">{child}</div>;
                    }
                    return null;
                })}
            </form>
        )
    }
}