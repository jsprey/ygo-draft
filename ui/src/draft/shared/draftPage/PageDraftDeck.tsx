import React, {useEffect, useState} from "react";
import {Card, Deck} from "../../../api/CardModel";
import {useRandomCards} from "../../../api/hooks/cards/useCards";
import DeckViewer from "../deck/DeckViewer";
import MultiCardDraftArea from "./MultiCardDraftArea";
import {YgoQueryClient} from "../../../index";
import {CardFilter} from "../../../api/CardFilter";
import Spinner from "../../../core/Spinner";
import Alert from "../../../core/Alert";
import ConfirmModal from "../../../core/ConfirmModal";
import Button from "../../../core/Button";
import useLocalState from "../../../api/hooks/storage/useLocalState";

const componentRandomQueryID = "draft_generator"

export type PageDraftDeckProps = {
    storagePrefix: string
    isMainDraft: boolean
    deck: Deck
    setDeck: React.Dispatch<React.SetStateAction<Deck>>
    onNextClick: () => void
    onAbort: () => void
    canAbort: boolean
    draftSize: number
    maxRounds: number
    filter: CardFilter
}

function PageDraftDeck(props: PageDraftDeckProps) {
    function getStorageKey(key:string) {
        return `${props.storagePrefix}_${key}`
    }

    const [draftDeck, setDraftDeck, resetDeckDraftDeck] = useLocalState<Deck>(getStorageKey("draft_selection"), {cards: []} as Deck)
    const [currentDraftRound, setCurrentDraftRound, resetCurrentDraftRound] = useLocalState<number>(getStorageKey("round"), 1)
    const [isDrafted, setDrafted, resetIsDrafted] = useLocalState<boolean>(getStorageKey("draft_drafted"), false)
    const [finished, setFinished, resetFinished] = useLocalState<boolean>(getStorageKey("draft_finished"), false)

    function resetStorage() {
        resetDeckDraftDeck()
        resetCurrentDraftRound()
        resetIsDrafted()
        resetFinished()
    }

    const {data, isLoading, error} = useRandomCards(componentRandomQueryID, props.draftSize, props.filter, {
        enabled: true,
        staleTime: Infinity
    })

    let handleNextClick = function (): void {
        resetStorage()
        props.onNextClick()
        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
    }

    let draftCard = function (draftedCard: Card): void {
        setDraftDeck({cards: []} as Deck)
        setDrafted(false)
        addCardToCurrentDeck(props.deck, props.setDeck, draftedCard)

        let newRound = currentDraftRound + 1
        setCurrentDraftRound(newRound)

        if (newRound > props.maxRounds) {
            setFinished(true)
            handleNextClick()
        }

        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
    }

    // Abort modal used to verify the abort process.
    const [showAbortDialog, setShowAbortDialog] = useState(false)
    let handleAbortDraftProcess = function (): void {
        resetStorage()
        props.onAbort()
        setShowAbortDialog(false)
        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
    }

    let body
    if (isLoading) {
        body = <Spinner/>
    } else if (error) {
        body = <Alert variant={'danger'}>
            Could not load deck!
        </Alert>
    } else if (!isDrafted && data) {
        setDraftDeck(data)
        setDrafted(true)
    } else if (isDrafted && data?.cards.length === 0) {
        body = <Alert variant={'danger'}>
            There are no cards that for the given filters. Abort Draft and choose different filters.
        </Alert>
    }

    const handleShow = () => setShowAbortDialog(true);

    return <div className={"mt-2"}>
        <ConfirmModal show={showAbortDialog}
                      setShow={setShowAbortDialog}
                      confirmName={"Leave"}
                      title={"Abort Draft"}
                      description={"Your currently drafted deck is going to be deleted."}
                      onConfirm={handleAbortDraftProcess}/>
        {body}
        {!finished && isDrafted ? <MultiCardDraftArea name={"Draft Area"} maxRound={props.maxRounds}
                                                      draftRound={currentDraftRound}
                                                      cards={draftDeck.cards}
                                                      draftAction={draftCard}/> : null}
        <p className={"text-3xl dark:text-white mt-2"}>Current Deck</p>
        <DeckViewer deck={props.deck} className={"mt-2"}/>
        <div className={"flex place-content-end mt-2"}>
            {props.canAbort ? <Button variant={"danger"}
                                      className={"ml-4 object-center"}
                                      disabled={isLoading}
                                      onClick={() => !isLoading ? handleShow() : null}>
                    Abort Draft
                </Button> : null}
            <Button variant={"primary"} className={"ml-4 object-center object-right"}
                    disabled={isLoading || (currentDraftRound <= props.maxRounds)}
                    onClick={() => !isLoading ? handleNextClick() : null}>
                Next
            </Button>
        </div>
    </div>
}

function addCardToCurrentDeck(currentDeck: Deck, setCurrentDeck: React.Dispatch<React.SetStateAction<Deck>>, newCard: Card) {
    let newDeck = {cards: currentDeck.cards} as Deck
    newDeck.cards.push(newCard)
    setCurrentDeck(newDeck)
}

export default PageDraftDeck
