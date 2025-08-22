import React from "react";
import PageDraftDeck from "../shared/draftPage/PageDraftDeck";
import PageSettings from "../shared/settingsPage/PageSettings";
import Stepper from "../../core/stepper/Stepper";
import StepperStep from "../../core/stepper/StepperStep";
import {Deck} from "../../api/CardModel";
import LocaleDraftWizardResultPage from "./LocaleDraftWizardResultPage";
import {ExtraDeckFilter, MainDeckFilter} from "../../api/CardFilter";
import {DraftSettings} from "../../api/Draft";
import {ExportDeck} from "../shared/deck/DeckRandomGeneratorPage";
import useLocalState from "../../api/hooks/storage/useLocalState";

export enum LocalDraftStages {
    Settings = 1,
    DraftMain,
    DraftExtra,
    DeckOverview
}

function LocaleDraftWizard() {
    const [draftSettings, setDraftSettings, resetDraftSettings] = useLocalState<DraftSettings>("deck_draft_wizard_settings", {} as DraftSettings)
    const [currentStage, setCurrentStage, resetCurrentStage] = useLocalState<LocalDraftStages>("deck_draft_wizard_stage", LocalDraftStages.Settings)
    const [deck, setDeck, resetDeck] = useLocalState<Deck>("deck_draft_wizard_deck", {cards: []} as Deck)

    let mainDeckFilter = MainDeckFilter
    mainDeckFilter.sets = draftSettings.sets

    let extraDeckFilter = ExtraDeckFilter
    extraDeckFilter.sets = draftSettings.sets

    function onAbortDraft() {
        resetDeck()
        resetDraftSettings()
        resetCurrentStage()
        setCurrentStage(LocalDraftStages.Settings)
    }

    let stageBody
    switch (currentStage) {
        case LocalDraftStages.DraftMain:
            stageBody = <PageDraftDeck isMainDraft={true}
                                       storagePrefix={"locale_main"}
                                       filter={mainDeckFilter}
                                       deck={deck}
                                       canAbort={true}
                                       setDeck={setDeck}
                                       draftSize={draftSettings.main_deck_size}
                                       maxRounds={draftSettings.main_deck_draws}
                                       onNextClick={() => setCurrentStage(LocalDraftStages.DraftExtra)}
                                       onAbort={onAbortDraft}/>
            break
        case LocalDraftStages.DraftExtra:
            stageBody = <PageDraftDeck isMainDraft={false}
                                       storagePrefix={"locale_extra"}
                                       filter={extraDeckFilter}
                                       deck={deck}
                                       canAbort={true}
                                       setDeck={setDeck}
                                       draftSize={draftSettings.extra_deck_size}
                                       maxRounds={draftSettings.extra_deck_draws}
                                       onNextClick={() => setCurrentStage(LocalDraftStages.DeckOverview)}
                                       onAbort={onAbortDraft}/>
            break
        case LocalDraftStages.DeckOverview:
            stageBody = <LocaleDraftWizardResultPage onAbort={onAbortDraft} deck={deck}/>
            break
        default:
        case LocalDraftStages.Settings:
            stageBody = <PageSettings submitButtonName={"Next"}
                                      local={true}
                                      onSettingsSubmit={(settings: DraftSettings) => {
                                          setDraftSettings(settings)
                                          setCurrentStage(LocalDraftStages.DraftMain)
                                      }}/>
            break
    }

    return <>
        <div className="grid grid-cols-1">
            {getCurrentStageHeader(currentStage)}
            {stageBody}
        </div>
    </>
}

function getCurrentStageHeader(currentStage: LocalDraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Settings"} stepDescription={"Control the draft process"}
                     isDone={currentStage > LocalDraftStages.Settings}
                     isActive={currentStage === LocalDraftStages.Settings}/>
        <StepperStep stepNr={2} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"}
                     isDone={currentStage > LocalDraftStages.DraftMain}
                     isActive={currentStage === LocalDraftStages.DraftMain}/>
        <StepperStep stepNr={3} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"}
                     isDone={currentStage > LocalDraftStages.DraftExtra}
                     isActive={currentStage === LocalDraftStages.DraftExtra}/>
        <StepperStep stepNr={4} stepName={"Deck Overview"} stepDescription={"Look at your finished deck"}
                     isDone={currentStage > LocalDraftStages.DeckOverview}
                     isActive={currentStage === LocalDraftStages.DeckOverview}/>
    </Stepper>
}

export default LocaleDraftWizard
