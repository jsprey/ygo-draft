import React, {useState} from "react";
import {CardSet} from "../api/Sets";
import {FormControl} from "react-bootstrap";
import SvgIconButton, {SvgIconButtonProps} from "../core/SvgIconButton";

const IconSelectAll = <SvgIconButton size={25}
                                     classNames={"fill-red-600 stroke-green-600 hover:stroke-green-500 active:stroke-green-400"}>
    <path
        d="M3.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L9.293 8 3.646 2.354a.5.5 0 0 1 0-.708z"/>
    <path
        d="M7.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L13.293 8 7.646 2.354a.5.5 0 0 1 0-.708z"/>
</SvgIconButton>
const IconRemoveAll = <SvgIconButton size={25}
                                     classNames={"stroke-red-600 hover:stroke-red-500 active:stroke-red-400"}>
    <path
        d="M8.354 1.646a.5.5 0 0 1 0 .708L2.707 8l5.647 5.646a.5.5 0 0 1-.708.708l-6-6a.5.5 0 0 1 0-.708l6-6a.5.5 0 0 1 .708 0z"/>
    <path
        d="M12.354 1.646a.5.5 0 0 1 0 .708L6.707 8l5.647 5.646a.5.5 0 0 1-.708.708l-6-6a.5.5 0 0 1 0-.708l6-6a.5.5 0 0 1 .708 0z"/>
</SvgIconButton>

export type CardSetReceiver = (cardSet: CardSet) => void;
export type CardAllSetReceiver = () => void;

export type CardSetListProps = {
    title: string
    cardSets: CardSet[]
    rootClassName?: any
    isTargetList: boolean
    allAction: CardAllSetReceiver
    actionList?: Map<React.ReactElement<SvgIconButtonProps>, CardSetReceiver>
}

function CardSelectedSetList(props: CardSetListProps) {
    const [filter, setFilter] = useState<string>("")

    const filteredSets = props.cardSets.filter((currentSet) => {
        if (filter == "") return true;

        return filter !== "" && !currentSet.set_name.toLowerCase().includes(filter.toLowerCase())
    })

    let listItems = filteredSets.map(currentSet => {
        let svgIcons: JSX.Element[] = []
        let index = 0
        props.actionList?.forEach((cardSetReceiver, key) => {
            let svgElement = <SvgIconButton key={index++} size={key.props.size}
                                            classNames={key.props.classNames}
                                            rootClassNames={key.props.rootClassNames}
                                            onClick={() => {
                                                cardSetReceiver(currentSet)
                                            }
                                            }>
                {key.props.children}
            </SvgIconButton>

            svgIcons.push(svgElement)
        })

        return <div key={currentSet.set_name} className={"select-none flex justify-content-between  dark:text-white"}>
            {currentSet.set_name}
            <div className={"flex gap-1 ml-5"}>
                {svgIcons}
            </div>
        </div>
    })

    let allActionIcon = <div className={"border-1 rounded-2 p-1 ml-1"} onClick={() => props.allAction()}>
        {props.isTargetList ? IconRemoveAll : IconSelectAll}
    </div>

    return <div className={props.rootClassName}>
        <div className={"m-0 text-center dark:text-white"}>{props.title}</div>
        <div className={"flex justify-content-center m-2"}>
            <FormControl
                autoFocus
                className="flex-fill"
                placeholder="Type to filter..."
                onChange={(e) => {
                    setFilter(e.target.value)
                }}/>
            {allActionIcon}
        </div>

        <div className={"bg-opacity-10 bg-secondary"}>
            <div className={"overflow-y-auto h-56 grid p-3 grid-cols-1 gap-2 auto-rows-min"}>
                {listItems}
            </div>
        </div>
    </div>
}

export default CardSelectedSetList
