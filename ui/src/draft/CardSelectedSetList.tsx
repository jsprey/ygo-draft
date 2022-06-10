import React, {useState} from "react";
import {CardSet} from "../api/Sets";
import {FormControl} from "react-bootstrap";
import SvgIconButton, {SvgIconButtonProps} from "../core/SvgIconButton";

export type CardSetReceiver = (cardSet: CardSet) => void;

export type CardSetListProps = {
    title: string
    cardSets: CardSet[]
    rootClassName?: any
    actionList?: Map<React.ReactElement<SvgIconButtonProps>, CardSetReceiver>
}

function CardSelectedSetList(props: CardSetListProps) {
    const [filter, setFilter] = useState<string>("")

    let listItems: JSX.Element[] = []
    props.cardSets.map(currentSet => {
        if (filter !== "" && !currentSet.set_name.includes(filter)) {
            return
        }

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

        listItems.push(<div className={"flex justify-content-between"} key={currentSet.set_code}>
            {currentSet.set_name}
            <div className={"flex gap-1 ml-5"}>
                {svgIcons}
            </div>
        </div>)
    })

    return <div className={props.rootClassName}>
        <div className={"m-0 text-center"}>{props.title}</div>
        <div className={"flex justify-content-center m-2"}><FormControl
            autoFocus
            className="flex-fill"
            placeholder="Type to filter..."
            onChange={(e) => {
                setFilter(e.target.value)
            }}
        /></div>

        <div className={"bg-opacity-10 bg-secondary"}>
            <div className={"overflow-y-auto h-56 grid p-3 grid-cols-1 gap-2 auto-rows-min"}>
                {listItems}
            </div>
        </div>
    </div>
}

export default CardSelectedSetList
