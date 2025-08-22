import React from "react";
import classNames from "classnames";
import Input from "../../../core/Input";

export type SettingsEntryProps = {
    value: number
    setValue: React.Dispatch<React.SetStateAction<number>>
    error: string
    setError: React.Dispatch<React.SetStateAction<string>>
    label: string
    tooltip: string
    min: number
    max: number
    className?: string
}

function SettingsEntry(props: SettingsEntryProps) {
    return <div className={classNames(props.className ? props.className : "")}>
        <Input isValid={props.error === ""}
               value={props.value}
               label={props.label}
               tooltip={props.tooltip}
               rootClassNames={"w-full"}
               inputClassNames={"w-full"}
               type={"number"}
               onChange={event => {
                   if (event.target.value === "" || isNaN(parseInt(event.target.value))) {
                       props.setValue(props.min)
                       props.setError("")
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
               }}/>
        {props.error !== "" ? <div className={"text-danger dark:text-danger"}>{props.error}</div> : <></>}
    </div>
}


export default SettingsEntry
