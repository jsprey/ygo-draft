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

        return <div key={currentSet.set_name} className={classNames("select-none flex justify-between p-2 text-dark dark:text-light", index % 2 === 0 ? "bg-light-1 dark:bg-dark-1" : "bg-light-2 dark:bg-dark-2")}>
            {currentSet.set_name}
            <div className={"flex gap-2 ml-5"}>
                {svgIcons}
            </div>
        </div>
    })


    const iconCN = classNames("p-1", "border-b border-r border-t border-dark dark:border-light", props.isTargetList ? "stroke-danger hover:stroke-danger-hover active:stroke-danger-active" : "stroke-success hover:stroke-success-hover active:stroke-success-active")
    const AllActionIcon = <YgoIcon icon={props.isTargetList ? "double-arrow-left" : "double-arrow-right"}
                                   size={30}
                                   onClick={(event) => {
                                       props.allAction()
                                       event.preventDefault()
                                   }}
                                   classNames={iconCN}/>

    return <div className={classNames(props.rootClassName)}>
        <div className={"rounded-tl rounded-tr flex-grow border-l border-r border-t focus:no-border p-2 text-dark dark:text-light bg-light-3 dark:bg-dark-3 border-dark dark:border-light"}>{props.title}</div>
        <div className={"flex justify-content-center"}>
            <input
                autoFocus
                className="flex-fill flex-grow border outline-none pl-2 text-dark dark:text-light bg-gray-200 dark:bg-gray-600 border-dark dark:border-light"
                placeholder="Type to filter..."
                onChange={(e) => {
                    setFilter(e.target.value)
                }}/>
            {AllActionIcon}
        </div>

        <div className={"bg-light-1 dark:bg-dark-1 border-l border-b border-r border-dark dark:border-light"}>
            <div className={"overflow-y-auto h-56 grid grid-cols-1 auto-rows-min"}>
                {listItems}
            </div>
        </div>
    </div>
}

export default CardSelectedSetList
