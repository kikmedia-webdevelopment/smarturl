import React from 'react'
import cn from 'classnames'

type IconType = {}

export type ButtonProps = {
    icon?: IconType;
    indicateDropdown?: boolean;
    onClick?: React.MouseEventHandler;
    isFullWidth?: boolean;
    onBlur?: React.FocusEventHandler;
    loading?: boolean;
    disabled?: boolean;
    testId?: string;
    buttonType?: 'primary' | 'positive' | 'negative' | 'warning' | 'muted' | 'naked';
    type?: 'button' | 'submit' | 'reset';
    size?: 'small' | 'large';
    href?: string;
    style?: React.CSSProperties;
    className?: string;
    children?: React.ReactNode;
    isActive?: boolean;
} & typeof defaultProps;

const defaultProps = {
    loading: false,
    isFullWidth: false,
    indicateDropdown: false,
    disabled: false,
    testId: 'jk-ui-button',
    buttonType: 'primary',
    type: 'button',
}

export class Button extends React.Component<ButtonProps> {
    static defaultProps = defaultProps

    render() {
        const {
            className,
            children,
            icon,
            buttonType,
            size,
            isFullWidth,
            onBlur,
            testId,
            onClick,
            loading,
            disabled,
            indicateDropdown,
            href,
            type,
            isActive,
            ...otherProps
        } = this.props

        const classNames = cn(
            'box-border h-10 inline-block p-0 rounded-sm font-sans text-sm overflow-hidden transition-all',
            'text-white bg-secondary border border-transparent hover:bg-brand'
            )

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const Element: any = href ? 'a' : 'button';

        return (
            <Element
                onBlur={(e: React.FocusEvent) => {
                    if (onBlur && !disabled) {
                        onBlur(e);
                    }
                }}
                onClick={(e: React.MouseEvent) => {
                    if (onClick && !disabled && !loading) {
                        onClick(e);
                    }
                }}
                data-test-id={testId}
                disabled={disabled}
                href={!disabled ? href : null}
                type={type}
                className={classNames}
                {...otherProps}
            >
                {children && <span className="px-5 py-3">{children}</span>}
            </Element>
        )
    }
}