import React, {useState} from "react";
import {CardSet, SetList, sortSets} from "../../../api/Sets";
import {useSets} from "../../../api/hooks/cards/useSets";
import CardSelectedSetList, {CardSetReceiver} from "./CardSelectedSetList";
import SvgIconButton, {SvgIconButtonProps} from "../../../core/SvgIconButton";
import SetDetailModal from "../deck/SetDetailModal";
import YgoIcon from "../../../core/YgoIcon";
import classNames from "classnames";
import Alert from "../../../core/Alert";
import Spinner from "../../../core/Spinner";

export type CardSetSelectorProps = {
    tooltip: string
    rowClass?: any
    selectedSets: CardSet[]
    setSelectedSets: React.Dispatch<React.SetStateAction<CardSet[]>>
}

const IconEye = <SvgIconButton size={25}
                               classNames={classNames("fill-secondary", "hover:fill-secondary-hover", "active:hover:fill-secondary-active")}>
    <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z"/>
    <path
        d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z"/>
</SvgIconButton>
const IconArrowRight = <SvgIconButton size={25}
                                      classNames={classNames("fill-success", "hover:fill-success-hover", "active:hover:fill-success-active")}>
    <path
        d="M0 14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v12zm4.5-6.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5a.5.5 0 0 1 0-1z"/>
</SvgIconButton>
const IconArrowLeft = <SvgIconButton size={25}
                                     classNames={classNames("fill-danger", "hover:fill-danger-hover", "active:hover:fill-danger-active")}>
    <path
        d="M16 14a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12zm-4.5-6.5H5.707l2.147-2.146a.5.5 0 1 0-.708-.708l-3 3a.5.5 0 0 0 0 .708l3 3a.5.5 0 0 0 .708-.708L5.707 8.5H11.5a.5.5 0 0 0 0-1z"/>
</SvgIconButton>

function CardSetSelector(props: CardSetSelectorProps) {
    const {data, isLoading, error} = useSets()

    const [isShowingSetCardsView, setIsShowingSetCardsView] = useState(false);
    const [expandedTooltip, setExpandedTooltip] = useState(false);
    const [currentDetailSet, setCurrentDetailSet] = useState("");
    const showSetCardsModal = (currentSet: string) => {
        setCurrentDetailSet(currentSet)
        setIsShowingSetCardsView(true);
    }

    let allSetItems: JSX.Element = <></>
    if (isLoading) {
        allSetItems = <Spinner/>
    } else if (error) {
        allSetItems = <Alert variant={'danger'}>Failed to load all sets!</Alert>
    } else if (data && data.sets) {
        let filteredSets = data.sets.filter(availableSet => {
            let selectCardSet = true
            props.selectedSets.forEach(selectedSet => {
                if (availableSet.set_name === selectedSet.set_name) {
                    selectCardSet = false
                    return
                }
            })
            return selectCardSet
        })
        filteredSets = sortSets(filteredSets)

        const actionList = new Map<React.ReactElement<SvgIconButtonProps>, CardSetReceiver>()
        actionList.set(IconEye, cardSet => showSetCardsModal(cardSet.set_code))
        actionList.set(IconArrowRight, cardSet => {
            let newSelectedCardSets: CardSet[] = []
            newSelectedCardSets.push(...props.selectedSets)
            newSelectedCardSets.push(cardSet)
            props.setSelectedSets(sortSets(newSelectedCardSets))
        })

        allSetItems = <>
            <CardSelectedSetList isTargetList={false}
                                 title={"Available Sets"}
                                 cardSets={filteredSets}
                                 rootClassName={"pr-1"}
                                 actionList={actionList}
                                 allAction={function () {
                                     let newSelectedCardSets: CardSet[]
                                     newSelectedCardSets = (data as SetList).sets
                                     props.setSelectedSets(sortSets(newSelectedCardSets))
                                 }}/>
        </>
    }

    let selectedCardsActionList = new Map<React.ReactElement<SvgIconButtonProps>, CardSetReceiver>()
    selectedCardsActionList.set(IconEye, cardSet => showSetCardsModal(cardSet.set_code))
    selectedCardsActionList.set(IconArrowLeft, cardSet => {
        let newSelectedCardSets: CardSet[] = []
        newSelectedCardSets.push(...props.selectedSets)
        newSelectedCardSets = newSelectedCardSets.filter(value => {
            return value.set_name !== cardSet.set_name
        })
        props.setSelectedSets(sortSets(newSelectedCardSets))
    })

    const CollapsedIcon = <YgoIcon icon={"help-outline"}
                                   size={18}
                                   onClick={(event) => {
                                       setExpandedTooltip(true)
                                       event.preventDefault()
                                   }}
                                   classNames={classNames("fill-dark dark:fill-light", "hover:fill-primary", "active:hover:fill-primary-hover")}/>
    const ExpandIcon = <YgoIcon icon={"help-fill"}
                                size={18}
                                onClick={(event) => {
                                    setExpandedTooltip(false)
                                    event.preventDefault()
                                }}
                                classNames={classNames("fill-primary", "hover:fill-primary-hover", "active:hover:fill-primary-active")}/>
    return <>
        <div className={classNames("flex", props.rowClass)}>
            <div className={"w-full"}>
                <div className={"w-full"}>
                    <div className={"flex mb-2"}>
                        <div className={"self-center mr-2 font-bold dark:text-white"}>Card Sets</div>
                        {expandedTooltip ? ExpandIcon : CollapsedIcon}
                    </div>
                    {expandedTooltip ? <div>
                        <Alert variant={"info"} className={"mb-2"}>{props.tooltip}</Alert>
                    </div> : <></>}
                </div>

                <div className={"w-full"}>
                    <div className={"flex w-full"}>
                        <div className={"w-2/4"}>
                            {allSetItems}
                        </div>
                        <div className={"w-2/4"}>
                            <CardSelectedSetList isTargetList={true}
                                                 title={"Selected Sets"}
                                                 cardSets={props.selectedSets}
                                                 actionList={selectedCardsActionList}
                                                 rootClassName={"pl-1"}
                                                 allAction={function () {
                                let newSelectedCardSets: CardSet[] = []
                                props.setSelectedSets(sortSets(newSelectedCardSets))
                            }}/>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        {isShowingSetCardsView ? <SetDetailModal setCode={currentDetailSet} setShow={setIsShowingSetCardsView}
                                                 isShowing={isShowingSetCardsView}/> : <></>}
    </>
}

export default CardSetSelector
