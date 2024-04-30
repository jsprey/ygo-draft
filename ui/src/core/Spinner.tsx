import React from "react";
import classNames from "classnames";

type SpinnerProps = {
    className?: string
};

function Spinner(props: SpinnerProps) {
    return <svg className={classNames("h-5 w-5", "animate-spin", props.className ? props.className : "")}/>
}

export default Spinner
