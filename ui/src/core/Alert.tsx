import React, {ReactNode} from "react";
import classNames from "classnames";

type AlertProps = {
    children?: ReactNode | undefined
    className?: string
    variant: 'danger' | 'success' | 'info' | 'warning'
};

function Alert(props: AlertProps) {
    let rootCN = classNames("px-4 py-3 rounded relative", props.className ? props.className : "")
    switch (props.variant) {
        case "danger":
            rootCN = classNames(rootCN, "bg-danger-light border border-danger-dark text-danger-dark")
            break;
        case "success":
            rootCN = classNames(rootCN, "bg-success-light border border-success-dark text-success-dark")
            break;
        case "info":
            rootCN = classNames(rootCN, "bg-primary-light border border-primary-dark text-primary-dark")
            break;
        case "warning":
            rootCN = classNames(rootCN, "bg-warning-light border border-warning-dark text-warning-dark")
            break;
    }

    return <div className={classNames(rootCN)} role="alert">
        <span className="block sm:inline">
            {props.children}
        </span>
        <span className="absolute top-0 bottom-0 right-0 px-4 py-3">
  </span>
    </div>
}

export default Alert
