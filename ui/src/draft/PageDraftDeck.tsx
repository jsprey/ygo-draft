import React, {useState} from "react";
import {Card, Deck} from "../api/CardModel";
import {useRandomCards} from "../api/hooks/useCards";
import DeckViewer from "../deck/DeckViewer";
import {Alert, Button, Modal, Spinner} from "react-bootstrap";
import {ExportDeck} from "../deck/DeckRandomGeneratorPage";
import MultiCardDraftArea from "./MultiCardDraftArea";
import {YgoQueryClient} from "../index";
import {DraftStages} from "./DeckDraftWizard";

const componentRandomQueryID = "draft_generator"
const emptyDeck: Deck = {cards: []}

export type PageDraftDeckProps = {
    setCurrentStage: React.Dispatch<React.SetStateAction<DraftStages>>
}

function PageDraftDeck(props: PageDraftDeckProps) {
    const [myDeck, setDeck] = useState(emptyDeck)
    const [draftDeck, setDraftDeck] = useState(emptyDeck)
    const [draftSize] = useState(5)
    const [currentDraftRound, setCurrentDraftRound] = useState(1)

    const {data, isLoading, error} = useRandomCards(componentRandomQueryID, draftSize, {enabled: true, staleTime: Infinity})

    let draftCard = function(draftedCard: Card): void {
        setDraftDeck(emptyDeck)
        addCardToCurrentDeck(myDeck, setDeck, draftedCard)

        setCurrentDraftRound(currentDraftRound + 1)
        YgoQueryClient.removeQueries(["random", componentRandomQueryID])
    }

    let body;
    if (draftDeck === emptyDeck) {
        if (isLoading) {
            body = <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading Deck...</span>
            </Spinner>
        } else if (error) {
            body = <Alert variant={"danger"}>
                Could not load deck!
            </Alert>
        } else if (data) {
            setDraftDeck(data)
        }
    }

    // Abort modal used to verify the abort process.
    const [showAbortDialog, setShowAbortDialog] = useState(false)
    const handleCloseAbortDraftProcessModal = () => setShowAbortDialog(false);
    const handleAbortDraftProcess = () => {
        setDeck({} as Deck)
        setDraftDeck({} as Deck)
        props.setCurrentStage(DraftStages.Settings)
        setShowAbortDialog(false)
    }
    const handleShow = () => setShowAbortDialog(true);

    return <>
        <h1 className={"mt-3 mb-3"}>
            Deck Generation
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    disabled={isLoading}
                    onClick={() => !isLoading ? ExportDeck(myDeck) : null}>
                Export
            </Button>
            <Button className={"ml-4 object-center"}
                    variant="danger"
                    disabled={isLoading}
                    onClick={() => !isLoading ? handleShow() : null}>
                Abort Draft
            </Button>
        </h1>
        {body}
        <MultiCardDraftArea name={"Draft Area"} draftRound={currentDraftRound} cards={draftDeck.cards} draftAction={draftCard}/>
        <p className={"mt-5 text-3xl"}>Current Deck</p>
        <DeckViewer deck={myDeck}/>
        <Modal show={showAbortDialog} onHide={handleCloseAbortDraftProcessModal}>
            <Modal.Header closeButton>
                <Modal.Title>Abort Draft Process?</Modal.Title>
            </Modal.Header>
            <Modal.Body>Your currently drafted deck is going to be deleted.</Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={handleCloseAbortDraftProcessModal}>
                    No
                </Button>
                <Button variant="danger" onClick={handleAbortDraftProcess}>
                    Yes
                </Button>
            </Modal.Footer>
        </Modal>
    </>
}

function addCardToCurrentDeck(currentDeck: Deck, setCurrentDeck: React.Dispatch<React.SetStateAction<Deck>>, newCard: Card) {
    let newDeck = {cards: currentDeck.cards} as Deck
    newDeck.cards.push(newCard)
    setCurrentDeck(newDeck)
}

export default PageDraftDeck
