import React, {ChangeEvent, ReactNode} from "react";
import classNames from "classnames";

type InputProps = {
    className?: string
    disabled?: boolean
    children?: ReactNode | undefined
    onValueChanged?: (value: ChangeEvent<HTMLInputElement>) => void
    isValid?: boolean
};

function Input(props: InputProps) {
    const inputAttributes: React.InputHTMLAttributes<HTMLInputElement> = {}
    if (props.onValueChanged) {
        inputAttributes.onChange = props.onValueChanged
    }

    let rootCN = classNames("p-1", "rounded-lg", "bg-light dark:bg-dark", "outline-none border", props.className ? props.className : "")
    if (!props.isValid) {
        rootCN = classNames(rootCN, "border-danger dark:border-danger text-danger-dark")
    } else {
        rootCN = classNames(rootCN, "border-light-3 dark:border-dark-3")
    }

    return <input className={rootCN}
                  {...inputAttributes}
                  disabled={props.disabled ? props.disabled : false}>
    </input>
}

export default Input
