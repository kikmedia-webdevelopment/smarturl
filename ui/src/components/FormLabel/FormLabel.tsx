import React from 'react'

export type FormLabelProps = {
    htmlFor: string;
    children: React.ReactNode;
    testId?: string;
    requiredText?: string;
    required?: boolean;
} & typeof defaultProps;

const defaultProps = {
    testId: 'jk-ui-form-label',
    requiredText: 'required',
    required: false,
}

export class FormLabel extends React.Component<FormLabelProps> {
    static defaultProps = defaultProps

    render() {
        const {
            children,
            testId,
            htmlFor,
            requiredText,
            required,
            ...otherProps
        } = this.props

        return (
            <label
                className="inline-block text-black font-semibold mb-2"
                data-test-id={testId}
                htmlFor={htmlFor}
                {...otherProps}
            >
                {children}
                {required && !!requiredText.length && (
                    <span className="text-gray-700 font-normal ml-1">
                        ({requiredText})
                    </span>
                )}
            </label>
        )
    }
}