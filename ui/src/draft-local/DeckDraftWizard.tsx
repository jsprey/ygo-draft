import React, {useState} from "react";
import PageDraftDeck from "./PageDraftDeck";
import PageSettings from "./PageSettings";
import Stepper from "../core/stepper/Stepper";
import StepperStep from "../core/stepper/StepperStep";
import {Deck} from "../api/CardModel";
import PageOverview from "./PageOverview";
import {ExtraDeckFilter, MainDeckFilter} from "../api/CardFilter";
import {CardSet} from "../api/Sets";

export type DraftSettings = {
    mainDraftSize: number
    mainDraftRound: number
    extraDraftSize: number
    extraDraftRound: number
    selectedCardSets: CardSet[]
}

export enum DraftStages {
    Settings = 1,
    DraftMain,
    DraftExtra,
    DeckOverview
}

function DeckDraftWizard() {
    const [draftSettings, setDraftSettings] = useState<DraftSettings>({} as DraftSettings)
    const [currentStage, setCurrentStage] = useState<DraftStages>(DraftStages.Settings)
    const [deck, setDeck] = useState<Deck>({cards: []} as Deck)

    let mainDeckFilter = MainDeckFilter
    mainDeckFilter.sets = draftSettings.selectedCardSets

    let extraDeckFilter = ExtraDeckFilter
    extraDeckFilter.sets = draftSettings.selectedCardSets

    let stageBody
    switch (currentStage) {
        case DraftStages.DraftMain:
            stageBody = <PageDraftDeck isMainDraft={true} filter={mainDeckFilter} deck={deck} setDeck={setDeck} draftSize={draftSettings.mainDraftSize}
                                       maxRounds={draftSettings.mainDraftRound} setCurrentStage={setCurrentStage}/>
            break
        case DraftStages.DraftExtra:
            stageBody = <PageDraftDeck isMainDraft={false} filter={extraDeckFilter} deck={deck} setDeck={setDeck} draftSize={draftSettings.extraDraftSize}
                                       maxRounds={draftSettings.extraDraftRound} setCurrentStage={setCurrentStage}/>
            break
        case DraftStages.DeckOverview:
            stageBody = <PageOverview deck={deck}/>
            break
        default:
        case DraftStages.Settings:
            stageBody = <PageSettings setDraftSettings={setDraftSettings} submitButtonName={"Next"} onSettingsSubmit={() => setCurrentStage(DraftStages.DraftMain)}/>
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

function getCurrentStageHeader(currentStage: DraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Settings"} stepDescription={"Control the draft process"}
                     isDone={currentStage > DraftStages.Settings} isActive={currentStage === DraftStages.Settings}/>
        <StepperStep stepNr={2} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"}
                     isDone={currentStage > DraftStages.DraftMain} isActive={currentStage === DraftStages.DraftMain}/>
        <StepperStep stepNr={3} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"}
                     isDone={currentStage > DraftStages.DraftExtra} isActive={currentStage === DraftStages.DraftExtra}/>
        <StepperStep stepNr={4} stepName={"Deck Overview"} stepDescription={"Look at your finished deck"}
                     isDone={currentStage > DraftStages.DeckOverview}
                     isActive={currentStage === DraftStages.DeckOverview}/>
    </Stepper>
}

export default DeckDraftWizard
