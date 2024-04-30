import classNames from "classnames";
import React from "react";

export type SvgIconButtonProps = {
    size?: number | undefined
    rootClassNames?: string | undefined
    classNames?: string | undefined
    onClick?: React.MouseEventHandler<any>
    children: JSX.Element | JSX.Element[]
}

function SvgIconButton(props: SvgIconButtonProps) {
    let imageSize = props.size
    if (imageSize === 0) {
        imageSize = 16
    }

    let imageClassNames = props.classNames
    if (imageClassNames === undefined) {
        imageClassNames = "fill-blue-600 hover:fill-blue-500 active:fill-blue-400"
    }

    if (props.onClick) {
        return <button onClick={event => {
            event.preventDefault()
            if (props.onClick) {
                props.onClick(event)
            }
        }
        } className={classNames("self-center align-middle", props.rootClassNames ? props.rootClassNames : "")}>
            <svg xmlns="http://www.w3.org/2000/svg" width={imageSize} height={imageSize} fill="currentColor"
                 className={imageClassNames} viewBox="0 0 16 16">
                {props.children}
            </svg>
        </button>
    }

    return <div className={classNames("self-center align-middle", props.rootClassNames ? props.rootClassNames : "")}>
        <svg xmlns="http://www.w3.org/2000/svg" width={imageSize} height={imageSize} fill="currentColor"
             className={imageClassNames} viewBox="0 0 16 16">
            {props.children}
        </svg>
    </div>
}

export default SvgIconButton