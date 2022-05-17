export type StepperProps= {
    children: JSX.Element[]
}

function Stepper(props: StepperProps) {
    return <div className={"flex justify-evenly mt-3"}>
        {props.children}
    </div>
}

export default Stepper