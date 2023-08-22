import React, {useState} from "react";
import {Alert, Col, Form, Placeholder, Row} from "react-bootstrap";
import HelpTooltip from "../core/HelpTooltip";
import {CardSet, SetList, sortSets} from "../api/Sets";
import {useSets} from "../api/hooks/useSets";
import CardSelectedSetList, {CardSetReceiver} from "./CardSelectedSetList";
import SvgIconButton, {SvgIconButtonProps} from "../core/SvgIconButton";
import SetDetailModal from "./SetDetailModal";

export type CardSetSelectorProps = {
    tooltip: string
    rowClass?: any
    selectedSets: CardSet[]
    setSelectedSets: React.Dispatch<React.SetStateAction<CardSet[]>>
}

const IconEye = <SvgIconButton size={25}>
    <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z"/>
    <path
        d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z"/>
</SvgIconButton>
const IconArrowRight = <SvgIconButton size={25}
                                      classNames={"fill-green-600 hover:fill-green-500 active:fill-green-400"}>
    <path
        d="M0 14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v12zm4.5-6.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5a.5.5 0 0 1 0-1z"/>
</SvgIconButton>
const IconArrowLeft = <SvgIconButton size={25} classNames={"fill-red-600 hover:fill-red-500 active:fill-red-400"}>
    <path
        d="M16 14a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12zm-4.5-6.5H5.707l2.147-2.146a.5.5 0 1 0-.708-.708l-3 3a.5.5 0 0 0 0 .708l3 3a.5.5 0 0 0 .708-.708L5.707 8.5H11.5a.5.5 0 0 0 0-1z"/>
</SvgIconButton>

function CardSetSelector(props: CardSetSelectorProps) {
    const {data, isLoading, error} = useSets()

    const [isShowingSetCardsView, setIsShowingSetCardsView] = useState(false);
    const [currentDetailSet, setCurrentDetailSet] = useState("");
    const showSetCardsModal = (currentSet:string) => {
        setCurrentDetailSet(currentSet)
        setIsShowingSetCardsView(true);
    }

    let allSetItems: JSX.Element = <></>
    if (isLoading) {
        allSetItems = <Placeholder animation="glow">
            <Placeholder xs={12} bg="primary"/>
            <Placeholder xs={12} bg="primary"/>
            <Placeholder xs={12} bg="primary"/>
            <Placeholder xs={12} bg="primary"/>
        </Placeholder>
    } else if (error) {
        allSetItems = <Alert variant={"danger"}>Failed to load all sets!</Alert>
    } else if (data && data.sets) {
        let availableSets: CardSet[] = sortSets(data.sets)
        availableSets = availableSets.filter(availableSet => {
            let selectCardSet = true
            props.selectedSets.forEach(selectedSet => {
                if (availableSet.set_name === selectedSet.set_name) {
                    selectCardSet = false
                    return
                }
            })
            return selectCardSet
        })

        const actionList = new Map<React.ReactElement<SvgIconButtonProps>, CardSetReceiver>()
        actionList.set(IconEye, cardSet => showSetCardsModal(cardSet.set_code))
        actionList.set(IconArrowRight, cardSet => {
            let newSelectedCardSets: CardSet[] = []
            newSelectedCardSets.push(...props.selectedSets)
            newSelectedCardSets.push(cardSet)
            props.setSelectedSets(sortSets(newSelectedCardSets))
        })

        allSetItems = <>
            <CardSelectedSetList isTargetList={false} title={"Available Sets"} cardSets={availableSets}
                                 actionList={actionList} allAction={function () {
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

    return <>
        <Row className={props.rowClass}><Form.Group as={Col}>
            <Form.Label className={"flex"}>
                <div className={"self-center mr-1"}>Select Packs for your drafting phases</div>
                <HelpTooltip size={20}
                             message={props.tooltip}/>
            </Form.Label>

            <div>
                <div className={"grid grid-cols-2 gap-3"}>
                    {allSetItems}
                    <CardSelectedSetList isTargetList={true} title={"Selected Sets"} cardSets={props.selectedSets}
                                         actionList={selectedCardsActionList} allAction={function () {
                        let newSelectedCardSets: CardSet[] = []
                        props.setSelectedSets(sortSets(newSelectedCardSets))
                    }}/>
                </div>
            </div>

        </Form.Group>
        </Row>
        {isShowingSetCardsView ? <SetDetailModal setCode={currentDetailSet} setShow={setIsShowingSetCardsView} isShowing={isShowingSetCardsView}/> : <></>}
    </>
}

export default CardSetSelector
