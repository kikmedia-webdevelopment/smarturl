import React from 'react'

export interface TabFocusTrapProps {
    children: React.ReactNode
}

export class TabFocusTrap extends React.Component<TabFocusTrapProps> {
    render() {
        const { children, ...otherProps } = this.props;


        return (
            <span tabIndex={-1} className="" {...otherProps}>
                {children}
            </span>
        );
    }
}