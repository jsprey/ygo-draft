import React from "react";

export type StepperStepProps = {
    stepNr: number
    stepName: string
    stepDescription: string
    isActive?: boolean
    isDone?: boolean
}

const CheckmarkSvg = <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20"  className="bi bi-check2 fill-white" viewBox="0 0 16 16">
    <path d="M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z"/>
</svg>

function StepperStep(props: StepperStepProps) {
    let stepNumberClasses = "text-xl text-white fw-bold justify-center items-center flex rounded-2 shadow-sm m-2 h-10 w-10"
    if (props.isDone) {
        stepNumberClasses += " bg-green-600"
    } else if (props.isActive) {
        stepNumberClasses += " bg-blue-600"
    } else {
        stepNumberClasses += " bg-blue-400"
    }

    return <div className={"m-2 flex"}>
        <div className={stepNumberClasses}>
            {props.isDone ? CheckmarkSvg : props.stepNr}
        </div>
        <div className={"grid grid-rows-2"}>
            <span className={props.isActive ? "text-lg fw-bold dark:text-white" : "text-lg dark:text-white"}>{props.stepName}</span>
            <span className={"text-base italic dark:text-white"}>{props.stepDescription}</span>
        </div>
    </div>
}

export default StepperStep