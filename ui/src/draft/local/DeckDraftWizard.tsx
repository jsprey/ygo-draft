import React, {useState} from "react";
import PageDraftDeck from "./PageDraftDeck";
import PageSettings from "./PageSettings";
import Stepper from "../../core/stepper/Stepper";
import StepperStep from "../../core/stepper/StepperStep";
import {Deck} from "../../api/CardModel";
import PageOverview from "./PageOverview";
import {ExtraDeckFilter, MainDeckFilter} from "../../api/CardFilter";
import {DraftSettings} from "../../api/Draft";

export enum LocalDraftStages {
    Settings = 1,
    DraftMain,
    DraftExtra,
    DeckOverview
}

function DeckDraftWizard() {
    const [draftSettings, setDraftSettings] = useState<DraftSettings>({} as DraftSettings)
    const [currentStage, setCurrentStage] = useState<LocalDraftStages>(LocalDraftStages.Settings)
    const [deck, setDeck] = useState<Deck>({cards: []} as Deck)

    let mainDeckFilter = MainDeckFilter
    mainDeckFilter.sets = draftSettings.sets

    let extraDeckFilter = ExtraDeckFilter
    extraDeckFilter.sets = draftSettings.sets

    let stageBody
    switch (currentStage) {
        case LocalDraftStages.DraftMain:
            stageBody = <PageDraftDeck isMainDraft={true} filter={mainDeckFilter} deck={deck} setDeck={setDeck} draftSize={draftSettings.main_deck_size}
                                       maxRounds={draftSettings.main_deck_draws} setCurrentStage={setCurrentStage}/>
            break
        case LocalDraftStages.DraftExtra:
            stageBody = <PageDraftDeck isMainDraft={false} filter={extraDeckFilter} deck={deck} setDeck={setDeck} draftSize={draftSettings.extra_deck_size}
                                       maxRounds={draftSettings.extra_deck_draws} setCurrentStage={setCurrentStage}/>
            break
        case LocalDraftStages.DeckOverview:
            stageBody = <PageOverview deck={deck}/>
            break
        default:
        case LocalDraftStages.Settings:
            stageBody = <PageSettings submitButtonName={"Next"} onSettingsSubmit={(settings: DraftSettings) => {
                setDraftSettings(settings)
                setCurrentStage(LocalDraftStages.DraftMain)
            }}/>
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

function getCurrentStageHeader(currentStage: LocalDraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Settings"} stepDescription={"Control the draft process"}
                     isDone={currentStage > LocalDraftStages.Settings} isActive={currentStage === LocalDraftStages.Settings}/>
        <StepperStep stepNr={2} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"}
                     isDone={currentStage > LocalDraftStages.DraftMain} isActive={currentStage === LocalDraftStages.DraftMain}/>
        <StepperStep stepNr={3} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"}
                     isDone={currentStage > LocalDraftStages.DraftExtra} isActive={currentStage === LocalDraftStages.DraftExtra}/>
        <StepperStep stepNr={4} stepName={"Deck Overview"} stepDescription={"Look at your finished deck"}
                     isDone={currentStage > LocalDraftStages.DeckOverview}
                     isActive={currentStage === LocalDraftStages.DeckOverview}/>
    </Stepper>
}

export default DeckDraftWizard
