import classNames from "classnames";

export type StepperProps = {
    children: JSX.Element[]
    className?: string
}

function Stepper(props: StepperProps) {
    return <div className={classNames("flex justify-evenly", props.className ? props.className : "")}>
        {props.children}
    </div>
}

export default Stepper