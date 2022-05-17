import React, {useState} from "react";
import PageDraftDeck from "./PageDraftDeck";
import PageSettings from "./PageSettings";
import Stepper from "../core/stepper/Stepper";
import StepperStep from "../core/stepper/StepperStep";

export enum DraftStages {
    Settings=1,
    DraftMain,
    DraftExtra,
    DeckOverview
}

function DeckDraftWizard() {
    const [currentStage, setCurrentStage] = useState<DraftStages>(DraftStages.Settings)

    let stageBody
    switch (currentStage) {
        case DraftStages.DraftMain:
            stageBody = <PageDraftDeck setCurrentStage={setCurrentStage}/>
            break
        case DraftStages.DraftExtra:
            stageBody = <PageDraftDeck setCurrentStage={setCurrentStage}/>
            break
        case DraftStages.DeckOverview:
            stageBody = <></>
            break
        default:
        case DraftStages.Settings:
            stageBody = <PageSettings setCurrentStage={setCurrentStage}/>
            break
    }

    return <>
        <div className="grid grid-cols-1">
            {getCurrentStageHeader(currentStage)}
            <hr/>
            {stageBody}
        </div>
    </>
}

function getCurrentStageHeader(currentStage: DraftStages): JSX.Element {
    return <Stepper>
        <StepperStep stepNr={1} stepName={"Settings"} stepDescription={"Control the draft process"} isDone={currentStage > DraftStages.Settings} isActive={currentStage === DraftStages.Settings}/>
        <StepperStep stepNr={2} stepName={"Draft: Main"} stepDescription={"Draft cards for your main deck"} isDone={currentStage > DraftStages.DraftMain} isActive={currentStage === DraftStages.DraftMain}/>
        <StepperStep stepNr={3} stepName={"Draft: Extra"} stepDescription={"Draft cards for your extra deck"} isDone={currentStage > DraftStages.DraftExtra} isActive={currentStage === DraftStages.DraftExtra}/>
        <StepperStep stepNr={4} stepName={"Deck Overview"} stepDescription={"Look at your finished deck"} isDone={currentStage > DraftStages.DeckOverview} isActive={currentStage === DraftStages.DeckOverview}/>
    </Stepper>
}

export default DeckDraftWizard
