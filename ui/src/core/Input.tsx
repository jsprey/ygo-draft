import React, {ChangeEvent, HTMLInputTypeAttribute, ReactNode, useState} from "react";
import classNames from "classnames";
import YgoIcon from "./YgoIcon";
import Alert from "./Alert";

type InputProps = {
    rootClassNames?: string
    inputClassNames?: string
    disabled?: boolean
    children?: ReactNode | undefined
    onChange?: (value: ChangeEvent<HTMLInputElement>) => void
    isValid?: boolean | undefined
    type?: HTMLInputTypeAttribute | undefined;
    value?: string | ReadonlyArray<string> | number | undefined;
    label?: string | undefined;
    tooltip?: string | undefined;
};

function Input(props: InputProps) {
    const [expanded, setExpanded] = useState<boolean>(false)

    const CollapsedIcon = <YgoIcon icon={"help-outline"}
                                   size={18}
                                   onClick={(event) => {
                                       setExpanded(true)
                                       event.preventDefault()
                                   }}
                                   classNames={classNames("fill-dark dark:fill-light", "hover:fill-primary", "active:hover:fill-primary-hover")}/>
    const ExpandIcon = <YgoIcon icon={"help-fill"}
                                size={18}
                                onClick={(event) => {
                                    setExpanded(false)
                                    event.preventDefault()
                                }}
                                classNames={classNames("fill-primary", "hover:fill-primary-hover", "active:hover:fill-primary-active")}/>

    const inputAttributes: React.InputHTMLAttributes<HTMLInputElement> = {}
    if (props.onChange) {
        inputAttributes.onChange = props.onChange
    }

    let inputCN = classNames("p-1 pl-2", "rounded-lg", "bg-light-1 dark:bg-dark-1", "text-lg dark:text-light", "border-2", props.inputClassNames ? props.inputClassNames : "")
    if (props.isValid !== undefined && !props.isValid) {
        inputCN = classNames(inputCN, "outline-danger-dark border-danger dark:border-danger text-danger")
    } else {
        inputCN = classNames(inputCN, "border-border")
    }

    if (props.type === "checkbox") {
        inputCN = classNames(inputCN, "h-[20px] w-[20px] accent-primary")
    } else {
        inputCN = classNames(inputCN, "h-[3rem]")
    }

    if (props.value) {
        inputAttributes.value = props.value
    }

    if (props.type) {
        inputAttributes.type = props.type
    }

    function getTooltipContent() {
        if (props.tooltip === undefined || !expanded) {
            return null
        }

        return <div>
            <Alert variant={"info"} className={"mb-2"}>{props.tooltip}</Alert>
        </div>
    }

    function getTooltipIcon(): ReactNode {
        if (props.tooltip === undefined) {
            return null
        }

        return expanded ? ExpandIcon : CollapsedIcon
    }

    if (props.type === "checkbox") {
        return <div
            className={classNames(props.rootClassNames ? props.rootClassNames : "")}>
            <div className={"w-full"}>
                <div className={classNames("flex items-center", expanded ? "mb-2" : "")}>
                    <input className={inputCN}
                           {...inputAttributes}
                           disabled={props.disabled ? props.disabled : false}>
                    </input>
                    <div className={"font-bold mr-2 ml-2 dark:text-light"}>{props.label}</div>
                    {getTooltipIcon()}
                </div>
                {getTooltipContent()}
            </div>
        </div>
    }

    return <div className={classNames(props.rootClassNames ? props.rootClassNames : "")}>
        {props.label !== undefined ? <div className={"self-center w-full"}>
                <div className={"flex mb-2 items-center"}>
                    <div className={" font-bold mr-2 dark:text-light"}>{props.label}</div>
                    {getTooltipIcon()}
                </div>
                {getTooltipContent()}
            </div> : null}
        <input className={inputCN}
               {...inputAttributes}
               disabled={props.disabled ? props.disabled : false}>
        </input>
    </div>
}

export default Input
