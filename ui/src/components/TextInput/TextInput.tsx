import React from 'react'
import cn from 'classnames'

export type TextInputProps = {
    width?: 'small' | 'medium' | 'large' | 'full';
    isReadOnly?: boolean;
    type?:
    | 'text'
    | 'password'
    | 'email'
    | 'number'
    | 'search'
    | 'url'
    | 'date'
    | 'time';
    name?: string;
    id?: string;
    className?: string;
    withCopyButton?: boolean;
    testId?: string;
    onCopy?: (value: string) => void;
    value?: string;
    inputRef?: React.RefObject<HTMLInputElement>;
    error?: boolean;
    willBlurOnEsc: boolean;
} & JSX.IntrinsicElements['input'] &
    typeof defaultProps;

export interface TextInputState {
    value?: string
}

const defaultProps = {
    testId: 'jk-ui-text-input',
    disabled: false,
    isReadOnly: false,
    required: false,
    width: 'full',
    willBlurOnEsc: true,
}

export class TextInput extends React.Component<TextInputProps, TextInputState> {
    static defaultProps = defaultProps

    constructor(props: TextInputProps) {
        super(props)

        this.state = {
            value: props.value
        }
    }

    UNSAFE_componentWillReceiveProps(nextProps: TextInputProps) {
        if (this.props.value !== nextProps.value) {
           this.setState({
                value: nextProps.value
            })
        }
    }

    handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { disabled, isReadOnly, onChange } = this.props
        if (disabled || isReadOnly) return;

        if (onChange) {
            onChange(e);
        }
        this.setState({ value: e.target.value });
    }

    handleFocus = (e: React.FocusEvent) => {
        if (this.props.disabled) {
            (e.target as HTMLInputElement).select();
        }
    }

    handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        const ESC = 27;

        if (this.props.onKeyDown) {
            this.props.onKeyDown(e);
        }

        if (e.keyCode === ESC && this.props.willBlurOnEsc) {
            e.currentTarget.blur();
        }
    }


    render() {
        const {
            className,
            withCopyButton,
            placeholder,
            maxLength,
            disabled,
            required,
            isReadOnly,
            onChange,
            testId,
            onBlur,
            onCopy,
            error,
            width,
            value,
            type,
            name,
            id,
            inputRef,
            willBlurOnEsc,
            ...otherProps
        } = this.props;

        const classNames = cn(
            'outline-none bg-white border border-solid border-gray-300 max-h-10 text-black font-sans text-sm p-2 m-0 w-full focus:outline-none focus:shadow-outline',
            {
                'cursor-not-allowed opacity-75 bg-gray-200': disabled
            }
        )

        return (
            <div
                className="flex w-full"
            >
                <input
                    onKeyDown={this.handleKeyDown}
                    aria-label={name}
                    className={classNames}
                    id={id}
                    name={name}
                    required={required}
                    placeholder={placeholder}
                    maxLength={maxLength}
                    data-test-id={testId}
                    disabled={disabled}
                    onBlur={onBlur}
                    onFocus={this.handleFocus}
                    onChange={this.handleChange}
                    value={this.state.value}
                    type={type}
                    ref={inputRef}
                    {...otherProps}
                />
            </div>
        )
    }
}