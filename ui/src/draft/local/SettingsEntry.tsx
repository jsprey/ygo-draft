import React, {useState} from "react";
import YgoIcon from "../../core/YgoIcon";
import classNames from "classnames";
import Alert from "../../core/Alert";

export type SettingsEntryProps = {
    value: number
    setValue: React.Dispatch<React.SetStateAction<number>>
    error: string
    setError: React.Dispatch<React.SetStateAction<string>>
    title: string
    tooltip: string
    min: number
    max: number
    className?: string
}

function SettingsEntry(props: SettingsEntryProps) {
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
    
    return <div className={classNames(props.className ? props.className : "")}>
        <div className={"w-full"}>
            <div className={"flex mb-2"}>
                <div className={"self-center font-bold mr-2 dark:text-light"}>{props.title}</div>
                {expanded ? ExpandIcon : CollapsedIcon}
            </div>
            {expanded ? <div>
                <Alert variant={"info"} className={"mb-2"}>{props.tooltip}</Alert>
            </div> : <></>}
        </div>
        <input
            required
            value={props.value}
            type="number"
            onChange={event => {
                if (isNaN(parseInt(event.target.value))) {
                    props.setValue(0)
                    props.setError(`Value needs to be between [${props.min} - ${props.max}].`)
                } else {
                    const setValue = parseInt(event.target.value)
                    if (setValue < props.min || setValue > props.max) {
                        props.setValue(parseInt(event.target.value))
                        props.setError(`Value needs to be between [${props.min} - ${props.max}].`)
                    } else {
                        props.setValue(parseInt(event.target.value))
                        props.setError("")
                    }
                }
            }
            }
        />
        {props.error !== "" ? <div className={"text-danger dark:text-danger"}>{props.error}</div> : <></>}
    </div>
}


export default SettingsEntry
