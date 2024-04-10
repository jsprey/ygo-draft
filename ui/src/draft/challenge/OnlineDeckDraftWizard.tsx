import React, {useState} from "react";
import Stepper from "../../core/stepper/Stepper";
import StepperStep from "../../core/stepper/StepperStep";
import {Deck} from "../../api/CardModel";
import {ExtraDeckFilter, MainDeckFilter} from "../../api/CardFilter";
import {DraftSettings} from "../../api/Draft";
import PageDraftDeck from "../local/PageDraftDeck";
import PageOverview from "../local/PageOverview";

export type OnlineDeckDraftWizardProps = {
    settings: DraftSettings
    onSubmit: (deck:Deck) => void
    submitName: string
    onAbort: () => void
}

export enum OnlineDraftStages {
    DraftMain,
    DraftExtra,
    DeckOverview
}

function OnlineDeckDraftWizard(props: OnlineDeckDraftWizardProps) {
    const [currentStage, setCurrentStage] = useState<OnlineDraftStages>(OnlineDraftStages.DraftMain)
    const [deck, setDeck] = useState<Deck>({cards: []} as Deck)

    const settings = props.settings

    let mainDeckFilter = MainDeckFilter
    mainDeckFilter.sets = settings.sets

    let extraDeckFilter = ExtraDeckFilter
    extraDeckFilter.sets = settings.sets

    let stageBody
    switch (currentStage) {
        case OnlineDraftStages.DraftExtra:
            stageBody = <PageDraftDeck isMainDraft={false}
                                       filter={extraDeckFilter}
                                       deck={deck}
                                       setDeck={setDeck}
                                       draftSize={settings.extra_deck_size}
                                       maxRounds={settings.extra_deck_draws}
                                       onNextClick={() => setCurrentStage(OnlineDraftStages.DeckOverview)}
                                       onAbort={() => props.onAbort()}/>
            break
        case OnlineDraftStages.DeckOverview:
            stageBody = <PageOverview deck={deck} onSubmit={props.onSubmit} submitName={props.submitName}/>
            break
        default:
        case OnlineDraftStages.DraftMain:
            stageBody = <PageDraftDeck isMainDraft={true}
                                       filter={mainDeckFilter}
                                       deck={deck}
                                       setDeck={setDeck}
                                       draftSize={settings.main_deck_size}
                                       maxRounds={settings.main_deck_draws}
                                       onNextClick={() => setCurrentStage(OnlineDraftStages.DraftExtra)}
                                       onAbort={() => props.onAbort}/>
            break
    }

    return <>
        <div className="grid grid-cols-1 pt-2 pb-3">
            {getCurrentStageHeader(currentStage)}
            <hr/>
            {stageBody}
        </div>
    </>
}

function getCurrentStageHeader(currentStage: OnlineDraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"}
                     isDone={currentStage > OnlineDraftStages.DraftMain} isActive={currentStage === OnlineDraftStages.DraftMain}/>
        <StepperStep stepNr={2} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"}
                     isDone={currentStage > OnlineDraftStages.DraftExtra} isActive={currentStage === OnlineDraftStages.DraftExtra}/>
        <StepperStep stepNr={3} stepName={"Deck Submit"} stepDescription={"Submit your finished deck"}
                     isDone={currentStage > OnlineDraftStages.DeckOverview}
                     isActive={currentStage === OnlineDraftStages.DeckOverview}/>
    </Stepper>
}

export default OnlineDeckDraftWizard
