import React, {ChangeEvent, HTMLInputTypeAttribute, ReactNode} from "react";
import classNames from "classnames";

type InputProps = {
    className?: string
    disabled?: boolean
    children?: ReactNode | undefined
    onChange?: (value: ChangeEvent<HTMLInputElement>) => void
    isValid?: boolean
    type?: HTMLInputTypeAttribute | undefined;
    value?: string | ReadonlyArray<string> | number | undefined;
};

function Input(props: InputProps) {
    const inputAttributes: React.InputHTMLAttributes<HTMLInputElement> = {}
    if (props.onChange) {
        inputAttributes.onChange = props.onChange
    }

    let rootCN = classNames("p-1", "rounded-lg", "bg-light-1 dark:bg-dark-1 dark:text-light", "outline-none border-2", props.className ? props.className : "")
    if (props.isValid) {
        rootCN = classNames(rootCN, "border-border")
    } else {
        rootCN = classNames(rootCN, "border-danger dark:border-danger text-danger")
    }

    if (props.value) {
        inputAttributes.value = props.value
    }

    if (props.type) {
        inputAttributes.type = props.type
    }

    return <input className={rootCN}
                  {...inputAttributes}
                  disabled={props.disabled ? props.disabled : false}>
    </input>
}

export default Input
