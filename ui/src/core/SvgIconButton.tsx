export type SvgIconButtonProps = {
    size?: number
    rootClassNames?: string
    classNames?: string
    onClick?: React.MouseEventHandler<SVGSVGElement>
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

    return <div className={props.rootClassNames ? "self-center ".concat(props?.rootClassNames) : "self-center"}>
        <svg onClick={props.onClick !== undefined ? props.onClick : function () {}} xmlns="http://www.w3.org/2000/svg" width={imageSize} height={imageSize} fill="currentColor"
             className={imageClassNames} viewBox="0 0 16 16">
            {props.children}
        </svg>
    </div>
}

export default SvgIconButton