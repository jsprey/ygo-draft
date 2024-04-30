import React, {useState} from "react";
import {Card, Deck} from "../../api/CardModel";
import {useRandomCards} from "../../api/hooks/cards/useCards";
import DeckViewer from "../../deck/DeckViewer";
import MultiCardDraftArea from "./MultiCardDraftArea";
import {YgoQueryClient} from "../../index";
import {usePrompt} from "../../api/hooks/usePromptBlocker";
import {CardFilter} from "../../api/CardFilter";
import Spinner from "../../core/Spinner";
import Alert from "../../core/Alert";
import ConfirmModal from "../../core/ConfirmModal";
import Button from "../../core/Button";

const componentRandomQueryID = "draft_generator"

export type PageDraftDeckProps = {
    isMainDraft: boolean
    deck: Deck
    setDeck: React.Dispatch<React.SetStateAction<Deck>>
    onNextClick: () => void
    onAbort: () => void
    draftSize: number
    maxRounds: number
    filter: CardFilter
}

function PageDraftDeck(props: PageDraftDeckProps) {
    usePrompt("Your unfinished deck is going to be deleted when leaving the page. Are you sure you want to leave?", true);

    const [draftDeck, setDraftDeck] = useState({cards: []} as Deck)
    const [isDrafted, setDrafted] = useState(false)
    const [currentDraftRound, setCurrentDraftRound] = useState(1)
    const [finished, setFinished] = useState(false)

    const {data, isLoading, error} = useRandomCards(componentRandomQueryID, props.draftSize, props.filter, {
        enabled: true,
        staleTime: Infinity
    })

    let handleNextClick = function (): void {
        setDraftDeck({cards: []} as Deck)
        setDrafted(false)
        setCurrentDraftRound(1)
        setFinished(false)
        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
        props.onNextClick()
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
        props.setDeck({cards: []} as Deck)
        setDraftDeck({cards: []} as Deck)
        setDrafted(false)
        setCurrentDraftRound(1)
        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
        props.onAbort()
        setShowAbortDialog(false)
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
            <Button variant={"danger"}
                    className={"ml-4 object-center"}
                    disabled={isLoading}
                    onClick={() => !isLoading ? handleShow() : null}>
                Abort Draft
            </Button>
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
