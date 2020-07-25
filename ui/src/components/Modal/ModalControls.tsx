import React from 'react';

export type ModalControlsProps = {
    testId?: string
    children: React.ReactNode
} & typeof defaultProps;

const defaultProps = {
    testId: 'jk-ui-modal-controls',
};

export class ModalControls extends React.Component<ModalControlsProps> {
    static defaultProps = defaultProps

    render() {
        const { testId, children, ...rest } = this.props;

        return (
            <div
                {...rest}
                className="px-8 pb-8"
                data-test-id={testId}
            >
                {children}
            </div>
        )
    }
}