import React from "react";
import YgoIcon from "../YgoIcon";
import classNames from "classnames";

export type StepperStepProps = {
    stepNr: number
    stepName: string
    stepDescription: string
    isActive?: boolean
    isDone?: boolean
}

function StepperStep(props: StepperStepProps) {
    let stepNumberClasses = classNames("h-14 w-14", "flex justify-center items-center", "text-xl font-bold", "rounded-tl-lg rounded-bl-lg shadow-sm", "border-l border-t border-b border-border")
    if (props.isDone) {
        stepNumberClasses = classNames(stepNumberClasses, "bg-success text-light")
    } else if (props.isActive) {
        stepNumberClasses = classNames(stepNumberClasses, "bg-primary text-light")
    } else {
        stepNumberClasses = classNames(stepNumberClasses, "bg-primary-light text-dark")
    }

    return <div className={"m-2 flex"}>
        <div className={stepNumberClasses}>
            {props.isDone ? <YgoIcon icon={"checkmark"} size={20}/> : props.stepNr}
        </div>
        <div className={"h-14 grid grid-rows-2 text-dark dark:text-light"}>
            <span className={classNames("text-lg pl-2 pr-2 bg-light dark:bg-dark rounded-tr-lg border-t border-r border-border", props.isActive ? "font-bold" : "")}>{props.stepName}</span>
            <span className={"truncate bg-light-1 dark:bg-dark-1 pl-2 pr-2 text-base italic text-dark-1 dark:text-light-1 rounded-br-lg border-b border-r border-border"}>{props.stepDescription}</span>
        </div>
    </div>
}

export default StepperStep