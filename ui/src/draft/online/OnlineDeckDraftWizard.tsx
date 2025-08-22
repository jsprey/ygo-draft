import React from "react";
import {Deck} from "../../api/CardModel";
import {ExtraDeckFilter, MainDeckFilter} from "../../api/CardFilter";
import {DraftSettings} from "../../api/Draft";
import PageDraftDeck from "../shared/draftPage/PageDraftDeck";
import useLocalState from "../../api/hooks/storage/useLocalState";
import OnlineDeckDraftWizardResultPage from "./OnlineDeckDraftWizardResultPage";
import StepperStep from "../../core/stepper/StepperStep";
import Stepper from "../../core/stepper/Stepper";

export type OnlineDeckDraftWizardProps = {
    settings: DraftSettings
    draftID: string
    roundID: string
    onAbort: () => void
}

export enum OnlineDraftStages {
    DraftMain,
    DraftExtra,
    DeckOverview
}

export const OnlineDeckDraftWizardStoragePrefix = "online_deck_wizard"

function OnlineDeckDraftWizard(props: OnlineDeckDraftWizardProps) {
    function getStorageKey(key: string) {
        return `${OnlineDeckDraftWizardStoragePrefix}_`
    }

    const [currentStage, setCurrentStage] = useLocalState<OnlineDraftStages>(getStorageKey("stage"), OnlineDraftStages.DraftMain)
    const [deck, setDeck] = useLocalState<Deck>(getStorageKey("deck"), {cards: []} as Deck)

    const settings = props.settings

    let mainDeckFilter = MainDeckFilter
    mainDeckFilter.sets = settings.sets

    let extraDeckFilter = ExtraDeckFilter
    extraDeckFilter.sets = settings.sets

    let stageBody
    switch (currentStage) {
        case OnlineDraftStages.DraftExtra:
            stageBody = <PageDraftDeck isMainDraft={false}
                                       storagePrefix={getStorageKey("extra")}
                                       filter={extraDeckFilter}
                                       deck={deck}
                                       setDeck={setDeck}
                                       canAbort={false}
                                       draftSize={settings.extra_deck_size}
                                       maxRounds={settings.extra_deck_draws}
                                       onNextClick={() => setCurrentStage(OnlineDraftStages.DeckOverview)}
                                       onAbort={() => props.onAbort()}/>
            break
        case OnlineDraftStages.DeckOverview:
            stageBody = <OnlineDeckDraftWizardResultPage deck={deck}
                                                         draftID={props.draftID}
                                                         roundID={props.roundID}
                                                         onAbort={props.onAbort}/>
            break
        default:
        case OnlineDraftStages.DraftMain:
            stageBody = <PageDraftDeck isMainDraft={true}
                                       storagePrefix={getStorageKey("main")}
                                       filter={mainDeckFilter}
                                       deck={deck}
                                       setDeck={setDeck}
                                       canAbort={false}
                                       draftSize={settings.main_deck_size}
                                       maxRounds={settings.main_deck_draws}
                                       onNextClick={() => setCurrentStage(OnlineDraftStages.DraftExtra)}
                                       onAbort={() => {
                                           props.onAbort()
                                       }}/>
            break
    }

    return <>
        <div className="grid grid-cols-1 pb-3">
            {getCurrentStageHeader(currentStage)}
            <span className={"mb-2"}></span>
            {stageBody}
        </div>
    </>
}

function getCurrentStageHeader(currentStage: OnlineDraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"}
                     isDone={currentStage > OnlineDraftStages.DraftMain}
                     isActive={currentStage === OnlineDraftStages.DraftMain}/>
        <StepperStep stepNr={2} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"}
                     isDone={currentStage > OnlineDraftStages.DraftExtra}
                     isActive={currentStage === OnlineDraftStages.DraftExtra}/>
        <StepperStep stepNr={3} stepName={"Deck Submit"} stepDescription={"Submit your finished deck"}
                     isDone={currentStage > OnlineDraftStages.DeckOverview}
                     isActive={currentStage === OnlineDraftStages.DeckOverview}/>
    </Stepper>
}

export default OnlineDeckDraftWizard
