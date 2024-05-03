import React, {LegacyRef, ReactNode} from "react";
import classNames from "classnames";

type ButtonProps = {
    children?: ReactNode | undefined
    className?: string
    variant: "primary" | "secondary" | "danger" | "success" | "neutral"
    onClick?: () => void
    ref?: LegacyRef<HTMLButtonElement>
    disabled?: boolean
};

function Button(props: ButtonProps) {
    let rootCN = classNames("p-2 rounded border", props.className ? props.className : "")
    const attributes: React.DOMAttributes<HTMLButtonElement> = {}
    const classAttributes: React.ClassAttributes<HTMLButtonElement> = {}

    if (props.disabled) {
        rootCN = classNames(rootCN, "text-dark-3 dark:text-light-3", "bg-light-1 dark:bg-dark-1 border border-border", "cursor-not-allowed")
    } else {
        attributes.onClick = event => {
            event.preventDefault()
            if (props.onClick) {
                props.onClick()
            }
        }

        if (props.ref) {
            classAttributes.ref = props.ref
        }

        switch (props.variant) {
            case "danger":
                rootCN = classNames(rootCN, "text-light", "bg-danger border border-danger-hover", "hover:bg-danger-hover border-danger-active", "active:bg-danger-active border-danger-dark")
                break;
            case "success":
                rootCN = classNames(rootCN, "text-light", "bg-success border border-success-hover", "hover:bg-success-hover border-success-active", "active:bg-success-active border-success-dark")
                break;
            case "primary":
                rootCN = classNames(rootCN, "text-light", "bg-primary border border-primary-hover", "hover:bg-primary-hover border-primary-active", "active:bg-primary-active border-primary-dark")
                break;
            case "secondary":
                rootCN = classNames(rootCN, "text-light", "bg-secondary border border-secondary-hover", "hover:bg-secondary-hover border-secondary-active", "active:bg-secondary-active border-secondary-dark")
                break;
            case "neutral":
                rootCN = classNames(rootCN, "text-dark dark:text-light", "bg-light dark:bg-dark border border-border", "hover:bg-light-1 dark:hover:bg-dark-1", "active:bg-light-2 dark:active:bg-dark-2")
                break;
        }
    }

    return <button className={rootCN} {...attributes} {...classAttributes}>
        {props.children}
    </button>
}

export default Button
