import React, {useState} from "react";
import {CardSet} from "../../api/Sets";
import SvgIconButton, {SvgIconButtonProps} from "../../core/SvgIconButton";
import classNames from "classnames";
import YgoIcon from "../../core/YgoIcon";

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
        if (filter === "") return true;

        return filter !== "" && currentSet.set_name.toLowerCase().includes(filter.toLowerCase())
    })

    let listItems = filteredSets.map((currentSet:CardSet, index:number) => {
        let svgIcons: JSX.Element[] = []
        let actionIndex = 0
        props.actionList?.forEach((cardSetReceiver, key) => {
            let svgElement = <SvgIconButton key={actionIndex++} size={key.props.size}
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

        return <div key={currentSet.set_name} className={classNames("select-none flex justify-content-between p-2 dark:text-white", index % 2 === 0 ? "bg-blue-100 dark:bg-gray-700" : "bg-blue-50 dark:bg-gray-600")}>
            {currentSet.set_name}
            <div className={"flex gap-1 ml-5"}>
                {svgIcons}
            </div>
        </div>
    })


    const iconCN = classNames("p-1", "border-bottom border-end border-top", props.isTargetList ? "stroke-red-600 hover:stroke-red-500 active:stroke-red-400 dark:stroke-red-300 dark:hover:stroke-red-400 dark:active:stroke-red-500" : "stroke-green-600 hover:stroke-green-500 active:stroke-green-400 dark:stroke-green-300 dark:hover:stroke-green-400 dark:active:stroke-green-500")
    const AllActionIcon = <YgoIcon icon={props.isTargetList ? "double-arrow-left" : "double-arrow-right"}
                                   size={30}
                                   onClick={(event) => {
                                       props.allAction()
                                       event.preventDefault()
                                   }}
                                   classNames={iconCN}/>

    return <div className={classNames(props.rootClassName)}>
        <div className={"rounded-tl rounded-tr flex-grow-1 border-start border-end border-top focus:no-border p-2 dark:text-white bg-gray-300 dark:bg-gray-700 "}>{props.title}</div>
        <div className={"flex justify-content-center"}>
            <input
                autoFocus
                className="flex-fill flex-grow-1 border-bottom border-start border-top focus:no-border pl-2 dark:text-white bg-gray-200 dark:bg-gray-600"
                placeholder="Type to filter..."
                onChange={(e) => {
                    setFilter(e.target.value)
                }}/>
            {AllActionIcon}
        </div>

        <div className={"bg-opacity-10 bg-secondary border-start border-bottom border-end"}>
            <div className={"overflow-y-auto h-56 grid grid-cols-1 auto-rows-min"}>
                {listItems}
            </div>
        </div>
    </div>
}

export default CardSelectedSetList
