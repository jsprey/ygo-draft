import React from "react";
import {Col, Form} from "react-bootstrap";
import HelpTooltip from "../core/HelpTooltip";

export type SettingsEntryProps = {
    value: number
    setValue: React.Dispatch<React.SetStateAction<number>>
    error: string
    setError: React.Dispatch<React.SetStateAction<string>>
    title: string
    tooltip: string
    min: number
    max: number
    md: 1|2|3|4|5|6|7|8|9
}

function SettingsEntry(props: SettingsEntryProps) {
    return <Form.Group as={Col} md={props.md}>
        <Form.Label className={"flex"}>
            <div className={"self-center mr-1  dark:text-white"}>{props.title}</div>
            <HelpTooltip size={20}
                         message={props.tooltip}/>
        </Form.Label>
        <Form.Control
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
            isInvalid={props.error !== ""}
        />
        {props.error !== "" ? <div className={"text-red-600"}>{props.error}</div> : <></>}
    </Form.Group>
}


export default SettingsEntry
