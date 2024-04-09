import React from "react";
import classNames from "classnames";
import './DraftBattleBadge.css'

type DraftBattleBadgeProps = {
    playerNameOne: string
    playerNameTwo: string
};

function DraftBattleBadge(props: DraftBattleBadgeProps) {
    const height: string = "5rem"
    const letterWidth: string = "2.5rem"
    let diagonalFactor: number = 0.05

    let clipLeft = `polygon(0 0, 100% 0, ${100 - (diagonalFactor * 100)}% 100%, 0% 100%)`
    let clipRight = `polygon(${diagonalFactor * 100}% 0, 100% 0, 100% 100%, 0% 100%)`
    let leftBannerMarginRight = `calc(50% * ${-diagonalFactor / 2})`
    let rightBannerMarginLeft = `calc(50% * ${-diagonalFactor / 2})`

    let leftBannerCN = classNames("rounded-tl rounded-bl flex-grow-1", "bg-gradient-to-r", "from-blue-500 from-0%", "via-blue-400 via-60%", "to-blue-100 to-90%")
    let rightBannerCN = classNames("rounded-tr rounded-br flex-grow-1", "bg-gradient-to-r", "from-red-100 from-0%", "via-red-400 via-40%", "to-red-500 to-100%")
    let vsTextCN = classNames("flex drop-shadow font-comic text-6xl text-white align-self-center align-items-center justify-content-center select-none")
    let nameTextCN = classNames("flex align-self-center", "font-comic", "text-5xl tracking-wider select-none", "text-white strokeBlack", "truncate")

    return <div className={"w-100"} style={{height: height}}>
        <div className={"inline-flex w-100"} style={{height: height}}>
            <div style={{width: "calc(45%)", marginRight: "-45%", zIndex: 10}}
                 className={classNames(nameTextCN, "pl-5")}>
                {props.playerNameOne}
            </div>
            <div style={{
                marginRight: leftBannerMarginRight,
                height: height,
                width: "50%",
                clipPath: clipLeft
            }}
                 className={leftBannerCN}>
            </div>
            <div style={{width: letterWidth, marginLeft: "-2.5rem", zIndex: 10}}
                 className={classNames(vsTextCN, "strokeLeft")}>
                V
            </div>
            <div style={{width: letterWidth, marginRight: "-2.5rem", zIndex: 10}}
                 className={classNames(vsTextCN, "strokeRight pl-2")}>
                S
            </div>
            <div style={{
                marginLeft: rightBannerMarginLeft,
                height: height,
                width: "50%",
                clipPath: clipRight
            }}
                 className={rightBannerCN}>
            </div>
            <div style={{width: "calc(45%)", marginLeft: "-45%", zIndex: 10}}
                 className={classNames(nameTextCN, "pr-5 justify-content-end")}>
                {props.playerNameTwo}
            </div>
        </div>
    </div>
}

export default DraftBattleBadge
